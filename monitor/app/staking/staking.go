package staking

import (
	"context"
	"math/big"
	"sort"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/cchain/queryutil"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

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
			instrEffRewards(ctx, cprov, allDelegations)
			instrSupplies(ctx, cprov, network, ethCls)
		}
	}
}

// instrSupplies instruments the supplies of OMNI token on the L1 (without the bridge balances)
// and on the consensus chain.
func instrSupplies(ctx context.Context, cprov cchain.Provider, network netconf.Network, ethCls map[uint64]ethclient.Client) {
	response, err := cprov.QueryClients().Bank.TotalSupply(ctx, &types.QueryTotalSupplyRequest{})
	if err != nil {
		log.Warn(ctx, "Failed to get total supply (will retry)", err)
		return
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
		log.Warn(ctx, "Failed to get Omni bindings (will retry)", err)
		return
	}
	l1TokenSupplyWei, err := l1Token.TotalSupply(nil)
	if err != nil {
		log.Warn(ctx, "Failed to get L1 total supply (will retry)", err)
		return
	}
	l1BridgeBalanceWei, err := l1Token.BalanceOf(nil, addrs.L1Bridge)
	if err != nil {
		log.Warn(ctx, "Failed to get L1 bridge balance (will retry)", err)
		return
	}

	l1TotalSupplyWei := new(big.Int).Sub(l1TokenSupplyWei, l1BridgeBalanceWei).Uint64()
	eChainSupply.Set(toGwei(float64(l1TotalSupplyWei)))
}

// instrEffRewards instruments effective staking rewards.
func instrEffRewards(ctx context.Context, cprov cchain.Provider, allDelegations []queryutil.DelegationBalance) {
	// Collect data during multiple blocks.
	const blocks = uint64(30)
	rewards, ok, err := queryutil.AvgRewardsRate(ctx, cprov, allDelegations, blocks)
	if err != nil {
		log.Warn(ctx, "Failed to get rewards rate (will retry)", err)
		return
	}

	if !ok {
		return
	}

	rewardsF64, err := rewards.Float64()
	if err != nil {
		log.Warn(ctx, "Failed to convert rewards rate to float64 [BUG]", err)
		return
	}

	rewardsAvg.Set(rewardsF64)
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
