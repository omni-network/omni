package eoa

import (
	"math/big"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/umath"
)

const (
	hotDynamicMultiplier  = 2
	coldDynamicMultiplier = 1
)

// dynamicThreshold list roles we use to sum their funding thresholds and a multiplier
// that increases the sum to calculate the dynamic threshold. Used by funding roles.
type dynamicThreshold struct {
	Multiplier float64
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
			minEther:    0.1,
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

	ethOnly = map[Role]bool{
		RoleFlowgen: true, // Flowgen is only used on ETH chains
	}
)

func GetFundThresholds(token tokens.Token, network netconf.ID, role Role) (FundThresholds, bool) {
	thresh, ok := getThreshold(network, role)
	if !ok {
		return FundThresholds{}, false
	}

	if token != tokens.ETH && ethOnly[role] {
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
	return umath.EtherToWei(t.minEther)
}

func (t FundThresholds) TargetBalance() *big.Int {
	return umath.EtherToWei(t.targetEther)
}

func convert(threshold FundThresholds, token tokens.Token) (FundThresholds, error) {
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
func multipleSum(network netconf.ID, multiplier float64, roles []Role) FundThresholds {
	var sum FundThresholds
	for _, role := range roles {
		thresh, ok := getThreshold(network, role)
		if !ok {
			continue
		}

		sum.minEther += thresh.minEther * multiplier
		sum.targetEther += thresh.targetEther * multiplier
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
