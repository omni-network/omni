package contracts

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// PackInitCode packs the init code for a contract deployment.
func PackInitCode(abi *abi.ABI, bytecodeHex string, params ...any) ([]byte, error) {
	input, err := abi.Pack("", params...)
	if err != nil {
		return nil, errors.Wrap(err, "pack init code")
	}

	bytecode, err := hexutil.Decode(bytecodeHex)
	if err != nil {
		return nil, errors.Wrap(err, "decode bytecode")
	}

	return append(bytecode, input...), nil
}

func IsEmptyAddress(addr common.Address) bool {
	return addr == common.Address{}
}

func IsDeployed(ctx context.Context, client ethclient.Client, addr common.Address) (bool, error) {
	code, err := client.CodeAt(ctx, addr, nil)
	if err != nil {
		return false, errors.Wrap(err, "code at", "address", addr)
	}

	if len(code) == 0 {
		return false, nil
	}

	return true, nil
}
