// Package uni provides universal generic blockchain type abstractions.
package uni

import (
	"encoding/json"

	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"

	"github.com/gagliardetto/solana-go"
)

// Kind represents the type of blockchain execution environment.
type Kind int

const (
	KindEVM Kind = iota // ethereum
	KindSVM             // solana
)

// IsSVMChain return true if the given chain ID is one of the SVM chains.
// Note that evmchain.IsSVMChain is equivalent and preferred, but this avoids a dependency on the evmchain->tokens package.
func IsSVMChain(chainID uint64) bool {
	return chainID == 350 || chainID == 351 || chainID == 352 // Hardcoded to avoid dependency on evmchain.
}

// Address represents a universal blockchain address.
// Note the zero value defaults to zero evm address.
type Address struct {
	kind Kind
	evm  common.Address
	svm  solana.PublicKey
}

func (a Address) MarshalJSON() ([]byte, error) {
	var b []byte
	var err error
	if a.IsEVM() {
		b, err = json.Marshal(a.evm)
	} else {
		b, err = json.Marshal(a.svm)
	}

	if err != nil {
		return nil, errors.Wrap(err, "marshal address")
	}

	return b, nil
}

func (a *Address) UnmarshalJSON(bz []byte) error {
	var evm common.Address
	if err := json.Unmarshal(bz, &evm); err == nil {
		a.kind = KindEVM
		a.evm = evm

		return nil
	}

	var svm solana.PublicKey
	if err := json.Unmarshal(bz, &svm); err == nil {
		a.kind = KindSVM
		a.svm = svm

		return nil
	}

	return errors.New("invalid address format", "address", string(bz))
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

// Bytes32 returns the address as a 32-byte array.
func (a Address) Bytes32() [32]byte {
	if a.IsEVM() {
		var bz [32]byte
		copy(bz[12:], a.evm.Bytes())

		return bz
	}

	return a.svm
}
