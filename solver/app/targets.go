package app

import (
	"bytes"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"

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

var (
	// targetsRestricted maps each network to whether targets should be restricted to the allowed set.
	targetsRestricted = map[netconf.ID]bool{
		netconf.Staging: true,
		netconf.Omega:   true,
		netconf.Mainnet: true,
	}

	// allowedTargets maps chain id to a set of allowed target addresses.
	allowedTargets = map[uint64]map[common.Address]bool{
		evmchain.IDSepolia: {
			common.HexToAddress("0x77F170Dcd0439c0057055a6D7e5A1Eb9c48cCD2a"): true, // wstETH vault 1
			common.HexToAddress("0x1BAe55e4774372F6181DaAaB4Ca197A8D9CC06Dd"): true, // wstETH vault 2
			common.HexToAddress("0x6415D3B5fc615D4a00C71f4044dEc24C141EBFf8"): true, // wstETH vault 3
		},
		evmchain.IDHolesky: {
			common.HexToAddress("0xd88dDf98fE4d161a66FB836bee4Ca469eb0E4a75"): true, // wstETH vault 1
			common.HexToAddress("0xa4c81649c79f8378a4409178E758B839F1d57a54"): true, // wstETH vault 2
		},
		evmchain.IDOmniStaging: {
			common.HexToAddress(predeploys.Staking): true,
		},
		evmchain.IDOmniOmega: {
			common.HexToAddress(predeploys.Staking): true,
		},
		evmchain.IDOmniMainnet: {
			common.HexToAddress(predeploys.Staking): true,
		},
	}
)

func newCallAllower(network netconf.ID, middlemanAddr common.Address) callAllowFunc {
	return func(chainID uint64, target common.Address, calldata []byte) bool {
		if !targetsRestricted[network] {
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

		if _, ok := allowedTargets[chainID]; !ok {
			return false
		}

		if allowed, ok := allowedTargets[chainID][target]; !ok || !allowed {
			return false
		}

		return true
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
