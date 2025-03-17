package queryutil

import (
	"context"
	"math/big"
	"time"

	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/forkjoin"
	"github.com/omni-network/omni/lib/umath"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

var yearMillis = math.LegacyNewDec(365 * 24 * time.Hour.Milliseconds())

// AvgRewardsRate returns the average staking rewards for all delegations over the given number of blocks.
func AvgRewardsRate(ctx context.Context, cprov cchain.Provider, delegations []DelegationBalance, waitBlocks uint64) (math.LegacyDec, error) {
	var delegators []sdk.AccAddress
	for _, del := range delegations {
		delegators = append(delegators, del.DelegatorAddress)
	}

	result, cancel := forkjoin.NewWithInputs(ctx, func(ctx context.Context, addr sdk.AccAddress) ([]math.LegacyDec, error) {
		return DelegatorInflationRates(ctx, cprov, addr, waitBlocks)
	}, delegators, forkjoin.WithWorkers(4)) // Don't overload the API
	defer cancel()

	inflations, err := result.Flatten()
	if err != nil {
		return math.LegacyDec{}, errors.Wrap(err, "forkjoin")
	}

	sum, length := math.LegacyZeroDec(), math.LegacyZeroDec()
	for _, infls := range inflations {
		for _, infl := range infls {
			sum = sum.Add(infl)
			length = length.Add(math.LegacyOneDec())
		}
	}

	if length.IsZero() {
		return math.LegacyDec{}, errors.New("zero delegations")
	}

	return sum.Quo(length), nil
}

// DelegatorInflationRates returns the inflation rate per delegation for the given delegator over the given number of blocks.
func DelegatorInflationRates(ctx context.Context, cprov cchain.Provider, delegator sdk.AccAddress, waitBlocks uint64) ([]math.LegacyDec, error) {
	rewards0, height0, timestamp0, err := getDelegationRewards(ctx, cprov, delegator)
	if err != nil {
		return nil, err
	}

	if err := waitUntil(ctx, cprov, height0+waitBlocks); err != nil {
		return nil, err
	}

	rewards1, _, timestamp1, err := getDelegationRewards(ctx, cprov, delegator)
	if err != nil {
		return nil, err
	} else if len(rewards0) != len(rewards1) {
		return nil, errors.New("delegations mismatch") // Staking actions occurred
	}

	milliDelta := math.LegacyNewDec(timestamp1.Sub(timestamp0).Milliseconds())

	var resp []math.LegacyDec
	for i := range len(rewards0) {
		rew0 := rewards0[i]
		rew1 := rewards1[i]

		if !rew0.Delegation.Balance.Equal(rew1.Delegation.Balance) {
			return nil, errors.New("delegation balance mismatch")
		}

		rewardDelta := rew1.Rewards.Sub(rew0.Rewards)
		rewardsPerYear := rewardDelta.Mul(yearMillis).Quo(milliDelta)
		stake := rew0.Delegation.Balance.Amount.ToLegacyDec()
		rewardsAPY := rewardsPerYear.Quo(stake)

		resp = append(resp, rewardsAPY)
	}

	return resp, nil
}

// DelegationBalance represents the total delegation balance of a delegator.
type DelegationBalance struct {
	DelegatorAddress sdk.AccAddress
	Balance          *big.Int
}

// AllDelegations returns delegation balances of each unique delegator.
func AllDelegations(ctx context.Context, cprov cchain.Provider) ([]DelegationBalance, error) {
	vals, err := cprov.SDKValidators(ctx)
	if err != nil {
		return nil, err
	}

	uniq := make(map[string]DelegationBalance)
	for _, val := range vals {
		if val.Jailed || val.IsUnbonded() {
			continue // Only use delegations from bonded unjailed validators
		}
		request := &stakingtypes.QueryValidatorDelegationsRequest{ValidatorAddr: val.OperatorAddress}
		for {
			resp, err := cprov.QueryClients().Staking.ValidatorDelegations(ctx, request)
			if err != nil {
				return nil, errors.Wrap(err, "query validator delegations")
			}

			for _, del := range resp.DelegationResponses {
				if del.Balance.Denom != sdk.DefaultBondDenom {
					continue
				}
				addr, err := sdk.AccAddressFromBech32(del.Delegation.DelegatorAddress)
				if err != nil {
					return nil, errors.Wrap(err, "parse delegator address")
				}
				if delegation, ok := uniq[del.Delegation.DelegatorAddress]; ok {
					delegation.Balance = umath.Add(delegation.Balance, del.Balance.Amount.BigInt())
					uniq[del.Delegation.DelegatorAddress] = delegation
				} else {
					uniq[del.Delegation.DelegatorAddress] =
						DelegationBalance{
							DelegatorAddress: addr,
							Balance:          del.Balance.Amount.BigInt(),
						}
				}
			}
			if len(resp.Pagination.NextKey) == 0 {
				break
			}
			request.Pagination = &query.PageRequest{Key: resp.Pagination.NextKey}
		}
	}

	var resp []DelegationBalance
	for _, del := range uniq {
		resp = append(resp, del)
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
