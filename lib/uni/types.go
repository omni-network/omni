// Package uni provides universal generic blockchain type abstractions.
package uni

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/gagliardetto/solana-go"
)

type Kind int

const (
	KindUnknown Kind = iota
	KindEth          // eth/evm
	KindSol          // solana
)

func IsSolChain(chainID uint64) bool {
	return chainID == 350 || chainID == 351 || chainID == 352 // Hardcoded to avoid dependency on evmchain.
}

type Address struct {
	kind Kind
	eth  common.Address
	sol  solana.PublicKey
}

func EthAddress(eth common.Address) Address {
	return Address{
		kind: KindEth,
		eth:  eth,
	}
}

func SolAddress(sol solana.PublicKey) Address {
	return Address{
		kind: KindSol,
		sol:  sol,
	}
}

func (a Address) Eth() common.Address {
	return a.eth
}

func (a Address) Sol() solana.PublicKey {
	return a.sol
}

func (a Address) Equals(o Address) bool {
	return a == o
}

func (a Address) EqualsEth(address common.Address) bool {
	return a.IsEth() && a.eth == address
}

func (a Address) IsEth() bool {
	return a.kind == KindEth
}

func (a Address) IsSol() bool {
	return a.kind == KindSol
}

func (a Address) IsZero() bool {
	return a.eth == common.Address{} && a.sol == solana.PublicKey{}
}
