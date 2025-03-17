package contract

import (
	"context"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/umath"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
		for _, contract := range allContracts {
			if !contract.IsDeployedOn(chain.ID, network.ID) {
				continue
			}

			go monitorContractForever(ctx, contract, chain, network.ID, rpcClients[chain.ID])
		}
	}

	return nil
}

// monitorContractForever blocks and periodically monitors the contract for the given chain.
func monitorContractForever(
	ctx context.Context,
	contract contracts.Contract,
	chain netconf.Chain,
	network netconf.ID,
	client ethclient.Client,
) {
	ctx = log.WithCtx(ctx,
		"chain", chain.Name,
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
			err := monitorContractOnce(ctx, contract, chain, network, client)
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
	chain netconf.Chain,
	network netconf.ID,
	client ethclient.Client,
) error {
	balance, err := client.BalanceAt(ctx, contract.Address, nil)
	if err != nil {
		return err
	}

	// Convert to ether units
	balanceEth := umath.ToEtherF64(balance)

	// Always set the balance metric
	contractBalance.WithLabelValues(chain.Name, contract.Name).Set(balanceEth)

	// Handle funding threshold checks, if any
	if contract.FundThresholds != nil {
		var isLow float64
		if umath.LTE(balance, contract.FundThresholds.MinBalance()) {
			isLow = 1
		}

		contractBalanceLow.WithLabelValues(chain.Name, contract.Name).Set(isLow)
	}

	// Handle withdrawal threshold checks, if any
	if contract.WithdrawThresholds != nil {
		var isHigh float64
		if umath.GTE(balance, contract.WithdrawThresholds.MaxBalance()) {
			isHigh = 1
		}

		contractBalanceHigh.WithLabelValues(chain.Name, contract.Name).Set(isHigh)
	}

	// Monitor token balances, if any
	if contract.Tokens != nil {
		for _, t := range contract.Tokens(chain.ID, network) {
			token, err := bindings.NewIERC20(t.Address, client)
			if err != nil {
				return err
			}

			balance, err := token.BalanceOf(&bind.CallOpts{Context: ctx}, contract.Address)
			if err != nil {
				return err
			}

			balanceEth := umath.ToEtherF64(balance)
			contractTokenBalance.WithLabelValues(chain.Name, contract.Name, t.Symbol, t.Address.Hex()).Set(balanceEth)
		}
	}

	return nil
}
