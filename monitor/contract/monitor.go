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

	toFund, err := contracts.ToFund(ctx, network.ID)
	if err != nil {
		log.Error(ctx, "Failed to get contract addreses to monitor for funding - skipping monitoring", err)
		return nil
	}

	toWithdraw, err := contracts.ToWithdraw(ctx, network.ID)
	if err != nil {
		log.Error(ctx, "Failed to get contract addreses to monitor for withdrawals - skipping monitoring", err)
		return nil
	}

	for _, chain := range network.EVMChains() {
		isOmniEVM := chain.ID == network.ID.Static().OmniExecutionChainID

		// Monitor funding contracts
		for _, contract := range toFund {
			if (contract.OnlyOmniEVM && !isOmniEVM) || (contract.NotOmniEVM && isOmniEVM) {
				continue
			}

			go monitorFundingContractForever(ctx, contract, chain.Name, rpcClients[chain.ID])
		}

		// Monitor withdraw contracts
		for _, contract := range toWithdraw {
			if (contract.OnlyOmniEVM && !isOmniEVM) || (contract.NotOmniEVM && isOmniEVM) {
				continue
			}

			go monitorWithdrawContractForever(ctx, contract, chain.Name, rpcClients[chain.ID])
		}
	}

	return nil
}

// monitorFundingContractForever blocks and periodically monitors funding the contract for the given chain.
func monitorFundingContractForever(
	ctx context.Context,
	contract contracts.WithFundThreshold,
	chainName string,
	client ethclient.Client,
) {
	ctx = log.WithCtx(ctx,
		"chain", chainName,
		"name", contract.Name,
		"address", contract.Address,
	)

	log.Info(ctx, "Monitoring account for funding")

	ticker := time.NewTicker(time.Second * 30)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := monitorFundingContractOnce(ctx, contract, chainName, client)
			if ctx.Err() != nil {
				return
			} else if err != nil {
				log.Warn(ctx, "Monitoring contract for funding failed (will retry)", err)

				continue
			}
		}
	}
}

// monitorFundingContractOnce monitors funding the contract for the given chain.
func monitorFundingContractOnce(
	ctx context.Context,
	contract contracts.WithFundThreshold,
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

	contractBalance.WithLabelValues(chainName, contract.Name).Set(balanceEth)

	var isLow float64
	if balance.Cmp(contract.Thresholds.MinBalance()) <= 0 {
		isLow = 1
	}

	contractBalanceLow.WithLabelValues(chainName, contract.Name).Set(isLow)

	return nil
}

// monitorWithdrawContractForever blocks and periodically monitors the contract for withdrawal needs.
func monitorWithdrawContractForever(
	ctx context.Context,
	contract contracts.WithWithdrawThreshold,
	chainName string,
	client ethclient.Client,
) {
	ctx = log.WithCtx(ctx,
		"chain", chainName,
		"name", contract.Name,
		"address", contract.Address,
	)

	log.Info(ctx, "Monitoring account for withdrawal")

	ticker := time.NewTicker(time.Second * 30)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := monitorWithdrawContractOnce(ctx, contract, chainName, client)
			if ctx.Err() != nil {
				return
			} else if err != nil {
				log.Warn(ctx, "Monitoring contract for withdrawal failed (will retry)", err)
				continue
			}
		}
	}
}

// monitorWithdrawContractOnce monitors contract for withdrawal needs.
func monitorWithdrawContractOnce(
	ctx context.Context,
	contract contracts.WithWithdrawThreshold,
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

	contractBalance.WithLabelValues(chainName, contract.Name).Set(balanceEth)

	var isHigh float64
	if balance.Cmp(contract.Thresholds.MinBalance()) >= 0 {
		isHigh = 1
	}

	contractBalanceHigh.WithLabelValues(chainName, contract.Name).Set(isHigh)

	return nil
}
