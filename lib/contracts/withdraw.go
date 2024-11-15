package contracts

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"

	"cosmossdk.io/math"
)

type WithWithdrawThreshold struct {
	Name        string
	OnlyOmniEVM bool
	NotOmniEVM  bool
	Address     common.Address
	Thresholds  WithdrawThresholds
}

func ToWithdraw(ctx context.Context, network netconf.ID) ([]WithWithdrawThreshold, error) {
	// GasPumps will not deployed initially on mainnet
	// TODO: remove this when mainnet GasPumps
	if network == netconf.Mainnet {
		return []WithWithdrawThreshold{}, nil
	}

	addrs, err := GetAddresses(ctx, network)
	if err != nil {
		return nil, err
	}

	return []WithWithdrawThreshold{
		{
			Name:        "gas-pump",
			Address:     addrs.GasPump,
			OnlyOmniEVM: false,
			NotOmniEVM:  true,
			Thresholds:  WithdrawThresholds{minEther: 10},
		},
	}, nil
}

type WithdrawThresholds struct {
	minEther float64
}

func (t WithdrawThresholds) MinBalance() *big.Int {
	gwei := t.minEther * params.GWei

	if gwei < 1 {
		panic("ether float64 must be greater than 1 Gwei")
	}

	return math.NewInt(params.GWei).MulRaw(int64(gwei)).BigInt()
}
