// Package uni provides universal generic blockchain type abstractions.
package uni

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/gagliardetto/solana-go"
)

// Kind represents the type of blockchain execution environment.
type Kind int

const (
	KindUnknown Kind = iota
	KindEVM          // ethereum
	KindSVM          // solana
)

func IsSVMChain(chainID uint64) bool {
	return chainID == 350 || chainID == 351 || chainID == 352 // Hardcoded to avoid dependency on evmchain.
}

type Address struct {
	kind Kind
	evm  common.Address
	svm  solana.PublicKey
}

func EVMAddress(evm common.Address) Address {
	return Address{
		kind: KindEVM,
		evm:  evm,
	}
}

func SVMAddress(svm solana.PublicKey) Address {
	return Address{
		kind: KindSVM,
		svm:  svm,
	}
}

func (a Address) EVM() common.Address {
	return a.evm
}

func (a Address) SVM() solana.PublicKey {
	return a.svm
}

func (a Address) Equals(o Address) bool {
	return a == o
}

func (a Address) EqualsEVM(address common.Address) bool {
	return a.IsEVM() && a.evm == address
}

func (a Address) IsEVM() bool {
	return a.kind == KindEVM
}

func (a Address) IsSVM() bool {
	return a.kind == KindSVM
}

func (a Address) IsZero() bool {
	return a.evm == common.Address{} && a.svm == solana.PublicKey{}
}

func (a Address) String() string {
	if a.IsEVM() {
		return a.evm.String()
	}

	return a.svm.String()
}
