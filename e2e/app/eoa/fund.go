package eoa

import (
	"math/big"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"
)

const (
	hotDynamicMultiplier  = 2
	coldDynamicMultiplier = 1
)

// dynamicThreshold list roles we use to sum their funding thresholds and a multiplier
// that increases the sum to calculate the dynamic threshold. Used by funding roles.
type dynamicThreshold struct {
	Multiplier uint64
	Roles      []Role
}

var (
	// tokenConversion defines conversion rate for fund threshold amounts.
	tokenConversion = map[tokens.Asset]float64{
		tokens.OMNI: 500,
		tokens.ETH:  1,
	}

	// thresholdTiny is used for EOAs which are rarely used, mostly to deploy a handful of contracts per network.
	thresholdTiny = FundThresholds{
		minEther:    0.001,
		targetEther: 0.01,
	}

	// thresholdSmall is used for EOAs which are used sometimes, mostly to make small test transactions per network.
	thresholdSmall = FundThresholds{
		minEther:    0.02,
		targetEther: .2,
	}

	// thresholdMedium is used by EOAs that regularly perform actions and need enough balance
	// to last a weekend without topping up even if fees are spiking.
	thresholdMedium = FundThresholds{
		minEther:    0.5,
		targetEther: 2,
	}

	// thresholdLarge is used by EOAs that constantly perform actions and need enough balance
	// to last a weekend without topping up even if fees are spiking.
	thresholdLarge = FundThresholds{
		minEther:    10,
		targetEther: 50,
	}

	staticThresholdsByRole = map[Role]FundThresholds{
		RoleRelayer:         thresholdMedium, // Relayer needs sufficient balance to operator for 2 weeks
		RoleMonitor:         thresholdMedium, // Dynamic Fee updates every few hours.
		RoleCreate3Deployer: thresholdTiny,   // Only 1 contract per chain
		RoleManager:         thresholdTiny,   // Rarely used
		RoleUpgrader:        thresholdTiny,   // Rarely used
		RoleDeployer:        thresholdTiny,   // Protected chains are only deployed once
		RoleTester:          thresholdLarge,  // Tester funds pingpongs, validator updates, etc, on non-mainnet.
		RoleXCaller:         thresholdSmall,  // XCaller funds used for sending xmsgs across networks.

		// Enough funds to fill orders, restricted to supported targets (to be implemented)
		RoleSolver: {
			minEther:    1,
			targetEther: 3,
		},

		// Needs enough to cover gas, and bridge eth between chains
		RoleFlowgen: {
			minEther:    0.01,
			targetEther: 1,
		},
	}

	dynamicThresholdsByRole = map[Role]dynamicThreshold{
		RoleHot: {
			Multiplier: hotDynamicMultiplier,
			Roles:      []Role{RoleRelayer, RoleMonitor, RoleFlowgen, RoleSolver, RoleCreate3Deployer, RoleManager, RoleUpgrader, RoleDeployer, RoleTester, RoleXCaller},
		},
		RoleCold: {
			Multiplier: coldDynamicMultiplier,
			Roles:      []Role{RoleHot},
		},
	}

	ephemeralOverrides = map[Role]FundThresholds{
		RoleDeployer: thresholdMedium, // Ephemeral chains are deployed often and fees can spike by a lot
	}

	nativeTokens = map[tokens.Asset]bool{
		tokens.ETH:  true,
		tokens.OMNI: true,
	}
)

func GetFundThresholds(token tokens.Asset, network netconf.ID, role Role) (FundThresholds, bool) {
	thresh, ok := getThreshold(token, network, role)
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
	minEther    float64
	targetEther float64
}

func (t FundThresholds) MinBalance() *big.Int {
	return bi.Ether(t.minEther)
}

func (t FundThresholds) TargetBalance() *big.Int {
	return bi.Ether(t.targetEther)
}

func convert(threshold FundThresholds, token tokens.Asset) (FundThresholds, error) {
	conversion, ok := tokenConversion[token]
	if !ok {
		return FundThresholds{}, errors.New("fund conversion", "token", token.String())
	}

	return FundThresholds{
		minEther:    threshold.minEther * conversion,
		targetEther: threshold.targetEther * conversion,
	}, nil
}

// multipleSum returns a function that calculates the sum of the thresholds for the given roles and multiplier.
func multipleSum(token tokens.Asset, network netconf.ID, multiplier uint64, roles []Role) FundThresholds {
	minSum, targetSum := bi.Zero(), bi.Zero()
	for _, role := range roles {
		thresh, ok := getThreshold(token, network, role)
		if !ok {
			continue
		}

		minSum = bi.Add(minSum, bi.MulRaw(thresh.MinBalance(), multiplier))
		targetSum = bi.Add(targetSum, bi.MulRaw(thresh.TargetBalance(), multiplier))
	}

	return FundThresholds{
		minEther:    bi.ToEtherF64(minSum),
		targetEther: bi.ToEtherF64(targetSum),
	}
}

func getThreshold(token tokens.Asset, network netconf.ID, role Role) (FundThresholds, bool) {
	if !nativeTokens[token] {
		// Only native tokenmeta are supported by default.
		return FundThresholds{}, false
	}

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
		return multipleSum(token, network, dynamic.Multiplier, dynamic.Roles), true
	}

	thresh, ok := staticThresholdsByRole[role]

	return thresh, ok
}
