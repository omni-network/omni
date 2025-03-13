package app

import (
	"bytes"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/solver/targets"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var (
	middlemanABI          = mustGetABI(bindings.SolverNetMiddlemanMetaData)
	execAndTransferMethod = mustMethod(middlemanABI, "executeAndTransfer")
)

type callAllowFunc func(chainID uint64, target common.Address, calldata []byte) bool

func newCallAllower(network netconf.ID, middlemanAddr common.Address) callAllowFunc {
	return func(chainID uint64, target common.Address, calldata []byte) bool {
		if !targets.IsRestricted(network) {
			return true
		}

		// flowgen can bridge to itself
		if target == eoa.MustAddress(network, eoa.RoleFlowgen) {
			return true
		}

		if target == middlemanAddr {
			proxiedTarget, _, err := parseMiddlemanCall(calldata)
			if err != nil {
				return false
			}

			target = proxiedTarget
		}

		_, ok := targets.Get(chainID, target)

		return ok
	}
}

// parseMiddlemanCall parses a call to the middleman contract, and returns proxied target and call data.
func parseMiddlemanCall(data []byte) (common.Address, []byte, error) {
	methodID := data[:4]

	// executeAndTransfer is only supported method
	method := execAndTransferMethod
	if !bytes.Equal(methodID, method.ID) {
		return common.Address{}, nil, errors.New("unknown method", "method", hexutil.Encode(methodID))
	}

	// decode call arguments
	unpacked, err := method.Inputs.Unpack(data[4:])
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "unpack call args")
	}

	wrap := struct {
		Token  common.Address
		To     common.Address
		Target common.Address
		Data   []byte
	}{}
	if err := method.Inputs.Copy(&wrap, unpacked); err != nil {
		return common.Address{}, nil, errors.Wrap(err, "copy call args")
	}

	return wrap.Target, wrap.Data, nil
}

func mustGetABI(metadata *bind.MetaData) *abi.ABI {
	abi, err := metadata.GetAbi()
	if err != nil {
		panic(err)
	}

	return abi
}

func mustMethod(abi *abi.ABI, name string) abi.Method {
	method, ok := abi.Methods[name]
	if !ok {
		panic("method not found")
	}

	return method
}
