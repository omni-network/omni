package main

import (
	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type ChainName string

const (
	ArbSepolia ChainName = "arb-sepolia"
	Holesky    ChainName = "holesky"
	OpSepolia  ChainName = "op-sepolia"
	OmniOmega  ChainName = "omni-omega"
)

func (c ChainName) Validate() error {
	switch c {
	case ArbSepolia, Holesky, OpSepolia, OmniOmega:
		return nil
	default:
		return errors.New("unknown chain", "chain", c)
	}
}

func (c ChainName) VerifierURL() string {
	return verifierURLs[c]
}

func (c ChainName) RPCURL() string {
	return rpcURLs[c]
}

var (
	// verifierURLs maps chain names to their respective verifier URLs.
	verifierURLs = map[ChainName]string{
		ArbSepolia: "https://api-sepolia.arbiscan.io/api",
		Holesky:    "https://api-holesky.etherscan.io/api",
		OpSepolia:  "https://api-sepolia-optimistic.etherscan.io/api",
		OmniOmega:  "https://api.routescan.io/v2/network/testnet/evm/164_4/etherscan/api",
	}

	// rpcURLs maps chain names to their respective RPC URLs.
	rpcURLs = map[ChainName]string{
		ArbSepolia: "https://sepolia-rollup.arbitrum.io/rpc",
		Holesky:    "https://ethereum-holesky-rpc.publicnode.com",
		OpSepolia:  "https://sepolia.optimism.io",
		OmniOmega:  "https://omega.omni.network",
	}

	// create3ABI is the ABI for the Create3 contract.
	create3ABI = mustGetABI(bindings.Create3MetaData)
)

func mustGetABI(metadata *bind.MetaData) *abi.ABI {
	abi, err := metadata.GetAbi()
	if err != nil {
		panic(err)
	}

	return abi
}
