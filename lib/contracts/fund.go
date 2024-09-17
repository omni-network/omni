package contracts

import (
	"math/big"

	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"

	"cosmossdk.io/math"
)

type WithFundThreshold struct {
	Name        string
	OnlyOmniEVM bool
	Address     common.Address
	Thresholds  FundThresholds
}

func ToFund(network netconf.ID) []WithFundThreshold {
	return []WithFundThreshold{
		{
			Name:        "gas-station",
			Address:     GasStation(network),
			OnlyOmniEVM: true,
			Thresholds:  FundThresholds{minEther: 200, targetEther: 1000}, // GasStation funds user GasPump requests, and needs a large OMNI balance.
		},
	}
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
