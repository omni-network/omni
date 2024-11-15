package contracts

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"

	"cosmossdk.io/math"
)

type Contract struct {
	Name               string
	Address            common.Address
	OnlyOmniEVM        bool
	NotOmniEVM         bool
	FundThresholds     *FundThresholds
	WithdrawThresholds *WithdrawThresholds
}

func ToFund(ctx context.Context, network netconf.ID) ([]Contract, error) {
	// GasStation will not deployed initially on mainnet
	// TODO: remove this when mainnet GasStation
	if network == netconf.Mainnet {
		return []Contract{}, nil
	}

	addrs, err := GetAddresses(ctx, network)
	if err != nil {
		return nil, err
	}

	return []Contract{
		{
			Name:           "gas-station",
			Address:        addrs.GasStation,
			OnlyOmniEVM:    true,
			NotOmniEVM:     false,
			FundThresholds: &FundThresholds{minEther: 200, targetEther: 1000}, // GasStation funds user GasPump requests, and needs a large OMNI balance.
		},
	}, nil
}

func ToWithdraw(ctx context.Context, network netconf.ID) ([]Contract, error) {
	// GasPumps will not deployed initially on mainnet
	// TODO: remove this when mainnet GasPumps
	if network == netconf.Mainnet {
		return []Contract{}, nil
	}

	addrs, err := GetAddresses(ctx, network)
	if err != nil {
		return nil, err
	}

	return []Contract{
		{
			Name:               "gas-pump",
			Address:            addrs.GasPump,
			OnlyOmniEVM:        false,
			NotOmniEVM:         true,
			WithdrawThresholds: &WithdrawThresholds{maxEther: 10},
		},
	}, nil
}

func ToMonitor() ([]Contract, error) {
	return []Contract{
		{
			Name:        "staking",
			Address:     common.HexToAddress(predeploys.Staking),
			OnlyOmniEVM: true,
			NotOmniEVM:  false,
		},
		{
			Name:        "nativeBridge",
			Address:     common.HexToAddress(predeploys.OmniBridgeNative),
			OnlyOmniEVM: true,
			NotOmniEVM:  false,
		},
	}, nil
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

type WithdrawThresholds struct {
	maxEther float64
}

func (t WithdrawThresholds) MaxBalance() *big.Int {
	gwei := t.maxEther * params.GWei

	if gwei < 1 {
		panic("ether float64 must be greater than 1 Gwei")
	}

	return math.NewInt(params.GWei).MulRaw(int64(gwei)).BigInt()
}
