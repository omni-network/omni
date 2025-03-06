package queryutil

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/forkjoin"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

var yearMillis = math.LegacyNewDec(365 * 24 * time.Hour.Milliseconds())

// AvgInflationRate returns the average inflation for all delegations over the given number of blocks
// or true if all delegations changed (couldn't calculate inflation).
func AvgInflationRate(ctx context.Context, cprov cchain.Provider, waitBlocks uint64) (math.LegacyDec, bool, error) {
	delegators, err := allDelegators(ctx, cprov)
	if err != nil {
		return math.LegacyDec{}, false, errors.Wrap(err, "get delegators")
	}

	result, cancel := forkjoin.NewWithInputs(ctx, func(ctx context.Context, addr sdk.AccAddress) ([]math.LegacyDec, error) {
		infl, changed, err := DelegatorInflationRates(ctx, cprov, addr, waitBlocks)
		if changed {
			return nil, nil
		}

		return infl, err
	}, delegators, forkjoin.WithWorkers(4)) // Don't overload the API
	defer cancel()

	inflations, err := result.Flatten()
	if err != nil {
		return math.LegacyDec{}, false, errors.Wrap(err, "forkjoin")
	}

	sum, length := math.LegacyZeroDec(), math.LegacyZeroDec()
	for _, infls := range inflations {
		for _, infl := range infls {
			sum = sum.Add(infl)
			length = length.Add(math.LegacyOneDec())
		}
	}

	if length.IsZero() {
		return math.LegacyDec{}, true, errors.New("zero delegations")
	}

	return sum.Quo(length), false, nil
}

// DelegatorInflationRates returns the inflation rate per delegation for the given delegator over the given number of blocks,
// or true if the delegation changed (couldn't calculate inflation).
func DelegatorInflationRates(ctx context.Context, cprov cchain.Provider, delegator sdk.AccAddress, waitBlocks uint64) ([]math.LegacyDec, bool, error) {
	rewards0, height0, timestamp0, err := getDelegationRewards(ctx, cprov, delegator)
	if err != nil {
		return nil, false, err
	}

	if err := waitUntil(ctx, cprov, height0+waitBlocks); err != nil {
		return nil, false, err
	}

	rewards1, _, timestamp1, err := getDelegationRewards(ctx, cprov, delegator)
	if err != nil {
		return nil, false, err
	} else if len(rewards0) != len(rewards1) {
		return nil, true, errors.New("delegations mismatch") // Staking actions occurred
	}

	milliDelta := math.LegacyNewDec(timestamp1.Sub(timestamp0).Milliseconds())

	var resp []math.LegacyDec
	for i := range len(rewards0) {
		rew0 := rewards0[i]
		rew1 := rewards1[i]

		if !rew0.Delegation.Balance.Equal(rew1.Delegation.Balance) {
			return nil, true, errors.New("delegation balance mismatch")
		}

		rewardDelta := rew1.Rewards.Sub(rew0.Rewards)
		rewardsPerYear := rewardDelta.Mul(yearMillis).Quo(milliDelta)
		stake := rew0.Delegation.Balance.Amount.ToLegacyDec()
		rewardsAPY := rewardsPerYear.Quo(stake)

		resp = append(resp, rewardsAPY)
	}

	return resp, false, nil
}

func allDelegators(ctx context.Context, cprov cchain.Provider) ([]sdk.AccAddress, error) {
	vals, err := cprov.SDKValidators(ctx)
	if err != nil {
		return nil, err
	}

	uniq := make(map[string]sdk.AccAddress)
	for _, val := range vals {
		resp, err := cprov.QueryClients().Staking.ValidatorDelegations(ctx, &stakingtypes.QueryValidatorDelegationsRequest{
			ValidatorAddr: val.OperatorAddress,
		})
		if err != nil {
			return nil, errors.Wrap(err, "query validator delegations")
		}

		for _, del := range resp.DelegationResponses {
			addr, err := sdk.AccAddressFromBech32(del.Delegation.DelegatorAddress)
			if err != nil {
				return nil, errors.Wrap(err, "parse delegator address")
			}

			uniq[del.Delegation.DelegatorAddress] = addr
		}
	}

	var resp []sdk.AccAddress
	for _, addr := range uniq {
		resp = append(resp, addr)
	}

	return resp, nil
}

func waitUntil(ctx context.Context, cprov cchain.Provider, target uint64) error {
	for {
		status, err := cprov.NodeStatus(ctx)
		if err != nil {
			return errors.Wrap(err, "get block")
		}
		height := status.Height

		if height >= target {
			return nil
		}

		time.Sleep(2 * time.Second)
	}
}

// delegationReward contains a delegation and its distribution module accrued rewards.
type delegationReward struct {
	Delegation stakingtypes.DelegationResponse
	Rewards    math.LegacyDec
}

// getDelegationRewards returns the current rewards-per-delegation (and height) for the given delegator.
func getDelegationRewards(ctx context.Context, cprov cchain.Provider, delegator sdk.AccAddress) ([]delegationReward, uint64, time.Time, error) {
	status, err := cprov.NodeStatus(ctx)
	if err != nil {
		return nil, 0, time.Time{}, errors.Wrap(err, "node status")
	}
	timestamp := *status.Timestamp
	height := status.Height

	resp, err := cprov.QueryClients().Staking.DelegatorDelegations(ctx, &stakingtypes.QueryDelegatorDelegationsRequest{
		DelegatorAddr: delegator.String(),
	})
	if err != nil {
		return nil, 0, timestamp, errors.Wrap(err, "query delegator delegations")
	} else if len(resp.DelegationResponses) == 0 {
		return nil, 0, timestamp, errors.New("no delegations")
	}

	var delegationRewards []delegationReward
	for _, del := range resp.DelegationResponses {
		rewardResp, err := cprov.QueryClients().Distribution.DelegationRewards(ctx, &distrtypes.QueryDelegationRewardsRequest{
			DelegatorAddress: del.Delegation.DelegatorAddress,
			ValidatorAddress: del.Delegation.ValidatorAddress,
		})
		if err != nil {
			return nil, 0, timestamp, errors.Wrap(err, "query delegation rewards")
		} else if len(rewardResp.Rewards) != 1 {
			continue // This is expected if delegation was processed in same block.
		}

		delegationRewards = append(delegationRewards, delegationReward{
			Delegation: del,
			Rewards:    rewardResp.Rewards[0].Amount,
		})
	}

	return delegationRewards, height, timestamp, nil
}
