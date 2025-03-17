package contracts

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
)

// Contract defines a contract to monitor.
type Contract struct {
	Name               string
	Address            common.Address
	FundThresholds     *FundThresholds
	WithdrawThresholds *WithdrawThresholds
	IsDeployedOn       func(chainID uint64, network netconf.ID) bool
	Tokens             func(chainID uint64, network netconf.ID) []Token
}

type Token struct {
	Symbol  string
	Address common.Address
}

// ToMonitor returns all contracts for the given network relevant to the monitor.
func ToMonitor(ctx context.Context, network netconf.ID) ([]Contract, error) {
	addrs, err := GetAddresses(ctx, network)
	if err != nil {
		return nil, err
	}

	// GasStation funds user GasPump requests, and needs a large OMNI balance. It is only deployed on OmniEVM.
	gasStation := Contract{
		Name:           "gas-station",
		Address:        addrs.GasStation,
		FundThresholds: &FundThresholds{minEther: 10, targetEther: 100},
		IsDeployedOn:   isOmni,
	}

	// GasPump is collects ETH from users. It is only deployed on ETH chains.
	gasPump := Contract{
		Name:               "gas-pump",
		Address:            addrs.GasPump,
		WithdrawThresholds: &WithdrawThresholds{maxEther: 10},
		IsDeployedOn:       isNotOmni,
	}

	// Staking contract collects validator deposits. It is only deployed on OmniEVM.
	staking := Contract{
		Name:         "staking",
		Address:      common.HexToAddress(predeploys.Staking),
		IsDeployedOn: isOmni,
	}

	// NativeBridge collects & spends native OMNI. It is only deployed on OmniEVM.
	nativeBridge := Contract{
		Name:         "native-bridge",
		Address:      common.HexToAddress(predeploys.OmniBridgeNative),
		IsDeployedOn: isOmni,
	}

	// L1Bridge collects & spends OMNI ERC20 on Ethereum. It is only deployed on Ethereum.
	l1Bridge := Contract{
		Name:         "l1-bridge",
		Address:      addrs.L1Bridge,
		IsDeployedOn: isEthereum,
		Tokens: func(chainID uint64, network netconf.ID) []Token {
			if !netconf.IsEthereumChain(network, chainID) {
				return nil
			}

			return []Token{{Symbol: "OMNI", Address: addrs.Token}}
		},
	}

	return []Contract{
		gasStation,
		gasPump,
		staking,
		nativeBridge,
		l1Bridge,
	}, nil
}

func isOmni(chainID uint64, network netconf.ID) bool {
	return chainID == network.Static().OmniExecutionChainID
}

func isNotOmni(chainID uint64, network netconf.ID) bool {
	return chainID != network.Static().OmniExecutionChainID && chainID != network.Static().OmniConsensusChainIDUint64()
}

func isEthereum(chainID uint64, network netconf.ID) bool {
	return netconf.IsEthereumChain(network, chainID)
}

// ToFund returns all fundable contracts for the given network.
func ToFund(ctx context.Context, network netconf.ID) ([]Contract, error) {
	contracts, err := ToMonitor(ctx, network)
	if err != nil {
		return nil, err
	}

	var fundContracts []Contract
	for _, contract := range contracts {
		if contract.FundThresholds != nil {
			fundContracts = append(fundContracts, contract)
		}
	}

	return fundContracts, nil
}

// FundThresholds defines the thresholds for funding a contract.
type FundThresholds struct {
	minEther    float64
	targetEther float64
}

// MinBalance returns the minimum balance required for funding a contract.
func (t FundThresholds) MinBalance() *big.Int {
	return bi.Ether(t.minEther)
}

// TargetBalance returns the target balance to fund a contract to.
func (t FundThresholds) TargetBalance() *big.Int {
	return bi.Ether(t.targetEther)
}

// WithdrawThresholds defines the thresholds for withdrawing from a contract.
type WithdrawThresholds struct {
	maxEther float64
}

// MaxBalance returns the max balance a contract can have before a withdrawal.
func (t WithdrawThresholds) MaxBalance() *big.Int {
	return bi.Ether(t.maxEther)
}
