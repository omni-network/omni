package contract

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/params"
)

// StartMonitoring starts the monitoring goroutines.
func StartMonitoring(ctx context.Context, network netconf.Network, endpoints xchain.RPCEndpoints, rpcClients map[uint64]ethclient.Client) error {
	log.Info(ctx, "Monitoring contracts")

	// If staging, we UseStagingOmniRPC(), to set rpc from which staging
	// addrs are derived (Create3 salt is derivative of first block hash)
	if network.ID == netconf.Staging {
		omniEVM, ok := network.Chain(netconf.Staging.Static().OmniExecutionChainID)
		if !ok {
			return errors.New("network missing omniEVM chain")
		}

		omniEVMRPC, err := endpoints.ByNameOrID(omniEVM.Name, omniEVM.ID)
		if err != nil {
			return err
		}

		contracts.UseStagingOmniRPC(omniEVMRPC)
	}

	allContracts, err := contracts.ToMonitor(ctx, network.ID)
	if err != nil {
		log.Error(ctx, "Failed to get contract addreses to monitor - skipping monitoring", err)
		return nil
	}

	for _, chain := range network.EVMChains() {
		isOmniEVM := chain.ID == network.ID.Static().OmniExecutionChainID

		for _, contract := range allContracts {
			if (contract.OnlyOmniEVM && !isOmniEVM) || (contract.NotOmniEVM && isOmniEVM) {
				continue
			}

			go monitorContractForever(ctx, contract, chain.Name, rpcClients[chain.ID])
		}
	}

	return nil
}

// monitorContractForever blocks and periodically monitors the contract for the given chain.
func monitorContractForever(
	ctx context.Context,
	contract contracts.Contract,
	chainName string,
	client ethclient.Client,
) {
	ctx = log.WithCtx(ctx,
		"chain", chainName,
		"name", contract.Name,
		"address", contract.Address,
	)

	log.Info(ctx, "Monitoring contract")

	ticker := time.NewTicker(time.Second * 30)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := monitorContractOnce(ctx, contract, chainName, client)
			if ctx.Err() != nil {
				return
			} else if err != nil {
				log.Warn(ctx, "Monitoring contract failed (will retry)", err)
				continue
			}
		}
	}
}

// monitorContractOnce monitors the contract for the given chain.
func monitorContractOnce(
	ctx context.Context,
	contract contracts.Contract,
	chainName string,
	client ethclient.Client,
) error {
	balance, err := client.BalanceAt(ctx, contract.Address, nil)
	if err != nil {
		return err
	}

	// Convert to ether units
	bf, _ := balance.Float64()
	balanceEth := bf / params.Ether

	// Always set the balance metric
	contractBalance.WithLabelValues(chainName, contract.Name).Set(balanceEth)

	// Handle funding threshold checks, if any
	if contract.FundThresholds != nil {
		var isLow float64
		if balance.Cmp(contract.FundThresholds.MinBalance()) <= 0 {
			isLow = 1
		}

		contractBalanceLow.WithLabelValues(chainName, contract.Name).Set(isLow)
	}

	// Handle withdrawal threshold checks, if any
	if contract.WithdrawThresholds != nil {
		var isHigh float64
		if balance.Cmp(contract.WithdrawThresholds.MaxBalance()) >= 0 {
			isHigh = 1
		}

		contractBalanceHigh.WithLabelValues(chainName, contract.Name).Set(isHigh)
	}

	return nil
}
