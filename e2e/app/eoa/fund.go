package eoa

import (
	"math/big"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"

	"github.com/ethereum/go-ethereum/params"

	"cosmossdk.io/math"
)

var (
	// tokenConversion defines conversion rate for fund threshold amounts.
	tokenConversion = map[tokens.Token]float64{
		tokens.OMNI: 1000,
	}

	// thresholdTiny is used for EOAs which are rarely used, mostly to deploy a handful of contracts per network.
	thresholdTiny = FundThresholds{
		minETH:    0.001,
		targetETH: 0.01,
	}

	// thresholdSmall is used by EOAs that deploy contracts or perform actions a couple times per week/month.
	//nolint:unused // Might be used in future.
	thresholdSmall = FundThresholds{
		minETH:    0.1,
		targetETH: 1.0,
	}

	// thresholdMedium is used by EOAs that regularly perform actions and need enough balance
	// to last a weekend without topping up even if fees are spiking.
	thresholdMedium = FundThresholds{
		minETH:    0.5,
		targetETH: 2,
	}

	// thresholdBig is used for funding actions to the medium EOAs.
	thresholdBig = FundThresholds{
		minETH:    thresholdMedium.targetETH,
		targetETH: thresholdMedium.targetETH * 3,
	}

	// thresholdLarge is used by EOAs that constantly perform actions and need enough balance
	// to last a weekend without topping up even if fees are spiking.
	thresholdLarge = FundThresholds{
		minETH:    10,
		targetETH: 50,
	}

	defaultThresholdsByRole = map[Role]FundThresholds{
		RoleRelayer:         thresholdMedium, // Relayer needs sufficient balance to operator for 2 weeks
		RoleMonitor:         thresholdMedium, // Dynamic Fee updates every few hours.
		RoleCreate3Deployer: thresholdTiny,   // Only 1 contract per chain
		RoleManager:         thresholdTiny,   // Rarely used
		RoleUpgrader:        thresholdTiny,   // Rarely used
		RoleDeployer:        thresholdTiny,   // Protected chains are only deployed once
		RoleFunder:          thresholdBig,    // Used to fund medium and smaller accounts.
		RoleSafe:            thresholdLarge,  // Used to fund funder.
		RoleTester:          thresholdLarge,  // Tester funds pingpongs, validator updates, etc.
	}

	ephemeralOverrides = map[Role]FundThresholds{
		RoleDeployer: thresholdMedium, // Ephemeral chains are deployed often and fees can spike by a lot
	}
)

func GetFundThresholds(token tokens.Token, network netconf.ID, role Role) (FundThresholds, bool) {
	if network.IsEphemeral() {
		if resp, ok := ephemeralOverrides[role]; ok {
			return resp, true
		}
	}

	if network == netconf.Mainnet && role == RoleTester {
		return FundThresholds{}, false
	}

	resp, ok := defaultThresholdsByRole[role]

	if ok && token != tokens.ETH {
		conv, err := convert(resp, token)
		if err != nil {
			panic(err)
		}

		return conv, true
	}

	return resp, ok
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
