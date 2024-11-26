package bindings

import (
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func (*SolveOutbox) ParseError(revert []byte) error {
	abi, err := SolveOutboxMetaData.GetAbi()
	if err != nil {
		return errors.Wrap(err, "parse abi")
	}

	return Error(abi, revert)
}

func (*SolveInbox) ParseError(revert []byte) error {
	abi, err := SolveInboxMetaData.GetAbi()
	if err != nil {
		return errors.Wrap(err, "parse abi")
	}

	return Error(abi, revert)
}

func Error(abi *abi.ABI, revert []byte) error {
	if revert == nil {
		return nil
	}

	for name, error := range abi.Errors {
		if _, err := error.Unpack(revert); err != nil {
			continue
		}

		return errors.New("custom error", "name", name)
	}

	return errors.New("unknown custom error", "selector", hexutil.Encode(revert[:4]))
}
