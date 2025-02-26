package eoa

import (
	"math"
	"math/big"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"

	"github.com/ethereum/go-ethereum/params"
)

const (
	hotDynamicMultiplier  = 2
	coldDynamicMultiplier = 1
)

// dynamicThreshold list roles we use to sum their funding thresholds and a multiplier
// that increases the sum to calculate the dynamic threshold. Used by funding roles.
type dynamicThreshold struct {
	Multiplier int
	Roles      []Role
}

var (
	// tokenConversion defines conversion rate for fund threshold amounts.
	tokenConversion = map[tokens.Token]int{
		tokens.OMNI: 500,
		tokens.ETH:  1,
	}

	// thresholdTiny is used for EOAs which are rarely used, mostly to deploy a handful of contracts per network.
	thresholdTiny = FundThresholds{
		minGwei:    gwei(0.001),
		targetGwei: gwei(0.01),
	}

	// thresholdSmall is used for EOAs which are used sometimes, mostly to make small test transactions per network.
	thresholdSmall = FundThresholds{
		minGwei:    gwei(0.02),
		targetGwei: gwei(.2),
	}

	// thresholdMedium is used by EOAs that regularly perform actions and need enough balance
	// to last a weekend without topping up even if fees are spiking.
	thresholdMedium = FundThresholds{
		minGwei:    gwei(0.5),
		targetGwei: gwei(2),
	}

	// thresholdLarge is used by EOAs that constantly perform actions and need enough balance
	// to last a weekend without topping up even if fees are spiking.
	thresholdLarge = FundThresholds{
		minGwei:    gwei(10),
		targetGwei: gwei(50),
	}

	staticThresholdsByRole = map[Role]FundThresholds{
		RoleRelayer:         thresholdMedium, // Relayer needs sufficient balance to operator for 2 weeks
		RoleMonitor:         thresholdMedium, // Dynamic Fee updates every few hours.
		RoleCreate3Deployer: thresholdSmall,  // Only 1 contract per chain (increased from tiny for sepolia)
		RoleManager:         thresholdTiny,   // Rarely used
		RoleUpgrader:        thresholdTiny,   // Rarely used
		RoleDeployer:        thresholdTiny,   // Protected chains are only deployed once
		RoleTester:          thresholdLarge,  // Tester funds pingpongs, validator updates, etc, on non-mainnet.
		RoleXCaller:         thresholdSmall,  // XCaller funds used for sending xmsgs across networks.
		RoleSolver:          thresholdLarge,  // Needs significant funds to fill orders (TODO: higher threshold?)
	}

	dynamicThresholdsByRole = map[Role]dynamicThreshold{
		RoleHot: {
			Multiplier: hotDynamicMultiplier,
			Roles:      []Role{RoleRelayer, RoleMonitor, RoleCreate3Deployer, RoleManager, RoleUpgrader, RoleDeployer, RoleTester, RoleXCaller},
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
	minGwei    uint64
	targetGwei uint64
}

func (t FundThresholds) MinBalance() *big.Int {
	return new(big.Int).Mul(big.NewInt(params.GWei), new(big.Int).SetUint64(t.minGwei))
}

func (t FundThresholds) TargetBalance() *big.Int {
	return new(big.Int).Mul(big.NewInt(params.GWei), new(big.Int).SetUint64(t.targetGwei))
}

func convert(threshold FundThresholds, token tokens.Token) (FundThresholds, error) {
	conversion, ok := tokenConversion[token]
	if !ok {
		return FundThresholds{}, errors.New("fund conversion", "token", token.String())
	}

	return FundThresholds{
		minGwei:    threshold.minGwei * uint64(conversion),
		targetGwei: threshold.targetGwei * uint64(conversion),
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

		sum.minGwei += thresh.minGwei * uint64(multiplier)
		sum.targetGwei += thresh.targetGwei * uint64(multiplier)
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

func gwei(eth float64) uint64 {
	g := eth * params.GWei

	_, dec := math.Modf(g)
	if dec != 0 {
		panic("ether float64 must be an int multiple of GWei")
	}

	return uint64(g)
}
