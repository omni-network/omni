package staking

import (
	"context"
	"sort"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/cchain/queryutil"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/params"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
)

func MonitorForever(ctx context.Context, cprov cchain.Provider, network netconf.Network, ethCls map[uint64]ethclient.Client) {
	ticker := time.NewTicker(time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			allDelegations, err := queryutil.AllDelegations(ctx, cprov)
			if err != nil {
				log.Warn(ctx, "Failed to query all delegations", err)
				continue
			}

			instrStakeSizes(allDelegations)

			if err := instrEffRewards(ctx, cprov, allDelegations); err != nil {
				log.Warn(ctx, "Effective rewards intrumentation failed", err)
			}

			if err := instrSupplies(ctx, cprov, network, ethCls); err != nil {
				log.Warn(ctx, "Token supply intrumentation failed", err)
			}
		}
	}
}

// instrSupplies instruments the supplies of OMNI token on the L1 (without the bridge balances)
// and on the consensus chain.
func instrSupplies(ctx context.Context, cprov cchain.Provider, network netconf.Network, ethCls map[uint64]ethclient.Client) error {
	response, err := cprov.QueryClients().Bank.TotalSupply(ctx, &types.QueryTotalSupplyRequest{})
	if err != nil {
		return errors.Wrap(err, "total supply query")
	}
	var cosmosSupplyWei uint64
	for _, coin := range response.Supply {
		if coin.Denom == sdk.DefaultBondDenom {
			cosmosSupplyWei += coin.Amount.Uint64()
		}
	}

	cChainSupply.Set(toGwei(float64(cosmosSupplyWei)))

	addrs, err := contracts.GetAddresses(ctx, network.ID)
	if err != nil {
		panic(err)
	}

	ethChainID := netconf.EthereumChainID(network.ID)
	l1Client := ethCls[ethChainID]
	l1Token, err := bindings.NewOmni(addrs.Token, l1Client)
	if err != nil {
		return errors.Wrap(err, "contract bindings")
	}

	callOpts := &bind.CallOpts{Context: ctx}

	l1TokenSupplyWei, err := l1Token.TotalSupply(callOpts)
	if err != nil {
		return errors.Wrap(err, "l1 token supply")
	}
	eChainSupply.Set(toGwei(float64(l1TokenSupplyWei.Uint64())))

	l1BridgeBalanceWei, err := l1Token.BalanceOf(callOpts, addrs.L1Bridge)
	if err != nil {
		return errors.Wrap(err, "l1 bridge balance")
	}
	bridgeBalance.Set(toGwei(float64(l1BridgeBalanceWei.Uint64())))

	return nil
}

// instrEffRewards instruments effective staking rewards.
func instrEffRewards(ctx context.Context, cprov cchain.Provider, allDelegations []queryutil.DelegationBalance) error {
	// Collect data during multiple blocks.
	const blocks = uint64(30)
	rewards, ok, err := queryutil.AvgRewardsRate(ctx, cprov, allDelegations, blocks)
	if err != nil {
		return errors.Wrap(err, "avg rewards")
	}

	if !ok {
		return nil
	}

	rewardsF64, err := rewards.Float64()
	if err != nil {
		return errors.Wrap(err, "rewards to flaot64 conversion")
	}

	rewardsAvg.Set(rewardsF64)

	return nil
}

// instrStakeSizes delegations instruments delegations data.
func instrStakeSizes(allDelegations []queryutil.DelegationBalance) {
	delegatorsTotal := float64(len(allDelegations))
	delegatorsCount.Set(delegatorsTotal)

	var totalStake float64
	var stakes []float64
	for _, del := range allDelegations {
		stake := float64(del.Balance.Amount.Uint64())
		totalStake += stake
		stakes = append(stakes, stake)
	}

	avgStakeWei := totalStake / delegatorsTotal
	stakeAvg.Set(toGwei(avgStakeWei))

	l := len(stakes)
	if l == 0 {
		return
	}
	sort.Slice(stakes, func(i, j int) bool {
		return stakes[i] < stakes[j]
	})
	medianVal := stakes[l/2+l%2]
	stakeMedian.Set(toGwei(medianVal))
}

func toGwei(wei float64) float64 {
	return wei / params.GWei
}
