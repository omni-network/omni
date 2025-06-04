package usdt0

import (
	"github.com/omni-network/omni/lib/evmchain"

	"github.com/ethereum/go-ethereum/common"
)

var (
	// oftByChain maps chain ID to OFT adapter address (the contract that sends / receives tokens).
	oftByChain = map[uint64]common.Address{
		evmchain.IDEthereum:    addr("0x6C96dE32CEa08842dcc4058c14d3aaAD7Fa41dee"),
		evmchain.IDOptimism:    addr("0xF03b4d9AC1D5d1E7c4cEf54C2A313b9fe051A0aD"),
		evmchain.IDArbitrumOne: addr("0x14E4A1B13bf7F943c8ff7C51fb60FA964A298D92"),
		evmchain.IDHyperEVM:    addr("0x904861a24F30EC96ea7CFC3bE9EA4B476d237e98"),
	}
)

func OFTByChain(chainID uint64) common.Address {
	if addr, ok := oftByChain[chainID]; ok {
		return addr
	}

	return common.Address{}
}

func addr(hex string) common.Address {
	return common.HexToAddress(hex)
}
