package keeper

import (
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/holiman/uint256"
)

var (
	redenomABI     = mustGetABI(bindings.RedenomMetaData)
	submittedEvent = mustGetEvent(redenomABI, "Submitted")
)

// mustGetABI returns the metadata's ABI as an abi.ABI type.
// It panics on error.
func mustGetABI(metadata *bind.MetaData) *abi.ABI {
	abi, err := metadata.GetAbi()
	if err != nil {
		panic(err)
	}

	return abi
}

// mustGetEvent returns the event with the given name from the ABI.
// It panics if the event is not found.
func mustGetEvent(abi *abi.ABI, name string) abi.Event {
	event, ok := abi.Events[name]
	if !ok {
		panic("event not found")
	}

	return event
}

// incHash returns the next hash, in lexicographical order (a.k.a plus one).
// Note it rolls over for common.MaxHash, returning zero hash.
func incHash(h common.Hash) common.Hash {
	return new(uint256.Int).AddUint64(
		new(uint256.Int).SetBytes32(h[:]),
		1,
	).Bytes32()
}

// calcMint returns the amount to mint to increase the balance to the multiplier.
func calcMint(balance *uint256.Int, multiplier int) (*big.Int, error) {
	if balance == nil {
		return nil, errors.New("nil balance")
	}

	if multiplier <= 0 {
		return nil, errors.New("invalid multiplier", "multiplier", multiplier)
	}

	if balance.IsZero() {
		return bi.Zero(), nil
	}

	return bi.MulRaw(balance.ToBig(), multiplier-1), nil
}
