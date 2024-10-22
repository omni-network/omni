package eoa

import (
	"math/big"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"

	"github.com/ethereum/go-ethereum/params"

	"cosmossdk.io/math"
)

const (
	hotDynamicMultiplier  = 2
	coldDynamicMultiplier = 1
)

type dynamicThreshold struct {
	Multiplier int
	Roles      []Role
}

var (
	// tokenConversion defines conversion rate for fund threshold amounts.
	tokenConversion = map[tokens.Token]float64{
		tokens.OMNI: 500,
		tokens.ETH:  1,
	}

	// thresholdTiny is used for EOAs which are rarely used, mostly to deploy a handful of contracts per network.
	thresholdTiny = FundThresholds{
		minETH:    0.001,
		targetETH: 0.01,
	}

	// thresholdMedium is used by EOAs that regularly perform actions and need enough balance
	// to last a weekend without topping up even if fees are spiking.
	thresholdMedium = FundThresholds{
		minETH:    0.5,
		targetETH: 2,
	}

	// thresholdLarge is used by EOAs that constantly perform actions and need enough balance
	// to last a weekend without topping up even if fees are spiking.
	thresholdLarge = FundThresholds{
		minETH:    10,
		targetETH: 50,
	}

	staticThresholdsByRole = map[Role]FundThresholds{
		RoleRelayer:         thresholdMedium, // Relayer needs sufficient balance to operator for 2 weeks
		RoleMonitor:         thresholdMedium, // Dynamic Fee updates every few hours.
		RoleCreate3Deployer: thresholdTiny,   // Only 1 contract per chain
		RoleManager:         thresholdTiny,   // Rarely used
		RoleUpgrader:        thresholdTiny,   // Rarely used
		RoleDeployer:        thresholdTiny,   // Protected chains are only deployed once
		RoleTester:          thresholdLarge,  // Tester funds pingpongs, validator updates, etc, on non-mainnet.
	}

	dynamicThresholdsByRole = map[Role]dynamicThreshold{
		RoleHot: {
			Multiplier: hotDynamicMultiplier,
			Roles:      []Role{RoleRelayer, RoleMonitor, RoleCreate3Deployer, RoleManager, RoleUpgrader, RoleDeployer, RoleTester},
		},
		RoleCold: {
			Multiplier: coldDynamicMultiplier,
			Roles:      []Role{RoleHot},
		},
	}

	ephemeralOverrides = map[Role]FundThresholds{
		RoleDeployer: thresholdMedium, // Ephemeral chains are deployed often and fees can spike by a lot
	}
)

func GetFundThresholds(token tokens.Token, network netconf.ID, role Role) (FundThresholds, bool) {
	thresh, ok := getThreshold(network, role)
	if !ok {
		return FundThresholds{}, false
	}

	// Convert thresholds to the token's denomination.
	conv, err := convert(thresh, token)
	if err != nil {
		panic(err)
	}

	return conv, true
}

type FundThresholds struct {
	minETH    float64
	targetETH float64
}

func (t FundThresholds) MinBalance() *big.Int {
	gwei := t.minETH * params.GWei

	if gwei < 1 {
		panic("ether float64 must be greater than 1 Gwei")
	}

	return math.NewInt(params.GWei).MulRaw(int64(gwei)).BigInt()
}

func (t FundThresholds) TargetBalance() *big.Int {
	gwei := t.targetETH * params.GWei
	if gwei < 1 {
		panic("ether float64 must be greater than 1 Gwei")
	}

	return math.NewInt(params.GWei).MulRaw(int64(gwei)).BigInt()
}

func convert(threshold FundThresholds, token tokens.Token) (FundThresholds, error) {
	conversion, ok := tokenConversion[token]
	if !ok {
		return FundThresholds{}, errors.New("fund conversion", "token", token.String())
	}

	return FundThresholds{
		minETH:    threshold.minETH * conversion,
		targetETH: threshold.targetETH * conversion,
	}, nil
}

// multipleSum returns a function that calculates the sum of the thresholds for the given roles and multiplier.
func multipleSum(network netconf.ID, multiplier int, roles []Role) FundThresholds {
	var sum FundThresholds
	for _, role := range roles {
		thresh, ok := getThreshold(network, role)
		if !ok {
			continue
		}

		sum.minETH += thresh.minETH * float64(multiplier)
		sum.targetETH += thresh.targetETH * float64(multiplier)
	}

	return sum
}

func getThreshold(network netconf.ID, role Role) (FundThresholds, bool) {
	if _, ok := AccountForRole(network, role); !ok {
		// Skip roles that don't have an account.
		return FundThresholds{}, false
	}

	if network.IsEphemeral() {
		override, ok := ephemeralOverrides[role]
		if ok {
			return override, true
		}
	}

	dynamic, ok := dynamicThresholdsByRole[role]
	if ok {
		return multipleSum(network, dynamic.Multiplier, dynamic.Roles), true
	}

	thresh, ok := staticThresholdsByRole[role]

	return thresh, ok
}
