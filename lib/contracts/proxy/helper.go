package proxy

import (
	"context"

	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"

	"github.com/ethereum/go-ethereum/common"
)

var (
	proxyImplementationSlot = common.HexToHash("0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc")
	proxyAdminSlot          = common.HexToHash("0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103")
)

func Impl(ctx context.Context, backend *ethbackend.Backend, addr common.Address) (common.Address, error) {
	code, err := backend.CodeAt(ctx, addr, nil)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "code at", "address", addr)
	}

	if len(code) == 0 {
		return common.Address{}, errors.New("no proxy", "address", addr)
	}

	impl, err := backend.StorageAt(ctx, addr, proxyImplementationSlot, nil)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "storage at")
	}

	return cast.MustBytesToAddress(impl), nil
}

func Admin(ctx context.Context, backend *ethbackend.Backend, addr common.Address) (common.Address, error) {
	code, err := backend.CodeAt(ctx, addr, nil)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "code at", "address", addr)
	}

	if len(code) == 0 {
		return common.Address{}, errors.New("no proxy", "address", addr)
	}

	admin, err := backend.StorageAt(ctx, addr, proxyAdminSlot, nil)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "storage at")
	}

	if len(admin) == 0 {
		return common.Address{}, errors.New("no admin", "address", addr)
	}

	return cast.MustBytesToAddress(admin), nil
}
