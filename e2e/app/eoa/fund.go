package eoa

import (
	"math/big"

	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/params"

	"cosmossdk.io/math"
)

var (
	// thresholdTiny is used for EOAs which are rarely used, mostly to deploy a handful of contracts per network.
	thresholdTiny = FundThresholds{
		minEther:    0.001,
		targetEther: 0.01,
	}

	// thresholdSmall is used by EOAs that deploy contracts or perform actions a couple times per week/month.
	//nolint:unused // Might be used in future.
	thresholdSmall = FundThresholds{
		minEther:    0.1,
		targetEther: 1.0,
	}

	// thresholdMedium is used by EOAs that regularly perform actions and need enough balance
	// to last a weekend without topping up even if fees are spiking.
	thresholdMedium = FundThresholds{
		minEther:    0.5,
		targetEther: 5,
	}

	// thresholdLarge is used by EOAs that constantly perform actions and need enough balance
	// to last a weekend without topping up even if fees are spiking.
	thresholdLarge = FundThresholds{
		minEther:    5,
		targetEther: 20, // TODO(corver): Increase along with e2e/app#saneMaxEther
	}

	defaultThresholdsByRole = map[Role]FundThresholds{
		RoleRelayer:         thresholdLarge,  // Relayer needs a ton of balance.
		RoleMonitor:         thresholdMedium, // Dynamic Fee updates every few hours.
		RoleCreate3Deployer: thresholdTiny,   // Only 1 contract per chain
		RoleAdmin:           thresholdTiny,   // Rarely used
		RoleDeployer:        thresholdTiny,   // Protected chains are only deployed once
		RoleTester:          thresholdLarge,  // Tester funds pingpongs, validator updates, etc.
	}

	ephemeralOverrides = map[Role]FundThresholds{
		RoleDeployer: thresholdMedium, // Ephemeral chains are deployed often and fees can spike by a lot
	}
)

func GetFundThresholds(network netconf.ID, role Role) (FundThresholds, bool) {
	if network.IsEphemeral() {
		if resp, ok := ephemeralOverrides[role]; ok {
			return resp, true
		}
	}

	resp, ok := defaultThresholdsByRole[role]

	return resp, ok
}

type FundThresholds struct {
	minEther    float64
	targetEther float64
}

func (t FundThresholds) MinBalance() *big.Int {
	gwei := t.minEther * params.GWei

	if gwei < 1 {
		panic("ether float64 must be greater than 1 Gwei")
	}

	return math.NewInt(params.GWei).MulRaw(int64(gwei)).BigInt()
}

func (t FundThresholds) TargetBalance() *big.Int {
	gwei := t.targetEther * params.GWei
	if gwei < 1 {
		panic("ether float64 must be greater than 1 Gwei")
	}

	return math.NewInt(params.GWei).MulRaw(int64(gwei)).BigInt()
}
