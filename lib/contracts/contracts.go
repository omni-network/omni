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

// ActionType defines the type of action to perform on a contract.
type ActionType uint8

const (
	Fund ActionType = iota
	Withdraw
	Monitor
)

// Contract defines a contract to monitor.
type Contract struct {
	Name               string
	Address            common.Address
	OnlyOmniEVM        bool
	NotOmniEVM         bool
	ActionType         ActionType
	FundThresholds     *FundThresholds
	WithdrawThresholds *WithdrawThresholds
}

// AllContracts returns all contracts for the given network relevant to the monitor.
func AllContracts(ctx context.Context, network netconf.ID) ([]Contract, error) {
	addrs, err := GetAddresses(ctx, network)
	if err != nil {
		return nil, err
	}

	return []Contract{
		// Funded contracts
		{
			Name:           "gas-station",
			Address:        addrs.GasStation,
			OnlyOmniEVM:    true,
			NotOmniEVM:     false,
			ActionType:     Fund,
			FundThresholds: &FundThresholds{minEther: 200, targetEther: 1000}, // GasStation funds user GasPump requests, and needs a large OMNI balance.
		},
		// Withdrawal contracts
		{
			Name:               "gas-pump",
			Address:            addrs.GasPump,
			OnlyOmniEVM:        false,
			NotOmniEVM:         true,
			ActionType:         Withdraw,
			WithdrawThresholds: &WithdrawThresholds{maxEther: 10},
		},
		// Monitoring contracts
		{
			Name:        "staking",
			Address:     common.HexToAddress(predeploys.Staking),
			OnlyOmniEVM: true,
			NotOmniEVM:  false,
			ActionType:  Monitor,
		},
		{
			Name:        "nativeBridge",
			Address:     common.HexToAddress(predeploys.OmniBridgeNative),
			OnlyOmniEVM: true,
			NotOmniEVM:  false,
			ActionType:  Monitor,
		},
	}, nil
}

// ToFund returns all fundable contracts for the given network.
func ToFund(ctx context.Context, network netconf.ID) ([]Contract, error) {
	// GasStation will not deployed initially on mainnet
	// TODO: remove this when mainnet GasStation
	if network == netconf.Mainnet {
		return []Contract{}, nil
	}

	contracts, err := AllContracts(ctx, network)
	if err != nil {
		return nil, err
	}

	var fundContracts []Contract
	for _, contract := range contracts {
		if contract.ActionType == Fund {
			fundContracts = append(fundContracts, contract)
		}
	}

	return fundContracts, nil
}

// ToWithdraw returns all withdrawable contracts for the given network.
func ToWithdraw(ctx context.Context, network netconf.ID) ([]Contract, error) {
	// GasPumps will not deployed initially on mainnet
	// TODO: remove this when mainnet GasPumps
	if network == netconf.Mainnet {
		return []Contract{}, nil
	}

	contracts, err := AllContracts(ctx, network)
	if err != nil {
		return nil, err
	}

	var withdrawContracts []Contract
	for _, contract := range contracts {
		if contract.ActionType == Withdraw {
			withdrawContracts = append(withdrawContracts, contract)
		}
	}

	return withdrawContracts, nil
}

// ToMonitor returns all monitorable contracts for the given network.
func ToMonitor(ctx context.Context, network netconf.ID) ([]Contract, error) {
	contracts, err := AllContracts(ctx, network)
	if err != nil {
		return nil, err
	}

	var monitorContracts []Contract
	for _, contract := range contracts {
		if contract.ActionType == Monitor {
			monitorContracts = append(monitorContracts, contract)
		}
	}

	return monitorContracts, nil
}

// FundThresholds defines the thresholds for funding a contract.
type FundThresholds struct {
	minEther    float64
	targetEther float64
}

// MinBalance returns the minimum balance required for funding a contract.
func (t FundThresholds) MinBalance() *big.Int {
	gwei := t.minEther * params.GWei

	if gwei < 1 {
		panic("ether float64 must be greater than 1 Gwei")
	}

	return math.NewInt(params.GWei).MulRaw(int64(gwei)).BigInt()
}

// TargetBalance returns the target balance to fund a contract to.
func (t FundThresholds) TargetBalance() *big.Int {
	gwei := t.targetEther * params.GWei
	if gwei < 1 {
		panic("ether float64 must be greater than 1 Gwei")
	}

	return math.NewInt(params.GWei).MulRaw(int64(gwei)).BigInt()
}

// WithdrawThresholds defines the thresholds for withdrawing from a contract.
type WithdrawThresholds struct {
	maxEther float64
}

// MaxBalance returns the max balance a contract can have before a withdrawal.
func (t WithdrawThresholds) MaxBalance() *big.Int {
	gwei := t.maxEther * params.GWei

	if gwei < 1 {
		panic("ether float64 must be greater than 1 Gwei")
	}

	return math.NewInt(params.GWei).MulRaw(int64(gwei)).BigInt()
}
