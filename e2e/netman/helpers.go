package netman

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
)

func deployOmniContracts(ctx context.Context, txOpts *bind.TransactOpts, backend *ethbackend.Backend,
	valSetID uint64, validators []bindings.Validator,
) (common.Address, *bindings.OmniPortal, error) {
	// TODO: make these configurable
	owner := txOpts.From
	fee := new(big.Int).SetUint64(params.GWei)

	proxyAdmin, err := DeployProxyAdmin(ctx, txOpts, backend, owner)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "deploy proxy admin")
	}

	feeOracle, err := deployFeeOracleV1(ctx, txOpts, backend, proxyAdmin, owner, fee)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "deploy fee oracle")
	}

	portal, err := deployPortal(ctx, txOpts, backend, proxyAdmin, owner, feeOracle, valSetID, validators)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "deploy portal")
	}

	contract, err := bindings.NewOmniPortal(portal, backend)
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "new portal")
	}

	return portal, contract, nil
}

func DeployProxyAdmin(ctx context.Context, txOpts *bind.TransactOpts, backend *ethbackend.Backend, owner common.Address,
) (common.Address, error) {
	proxyAdmin, tx, _, err := bindings.DeployProxyAdmin(txOpts, backend, owner)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "deploy proxy admin")
	}

	receipt, err := bind.WaitMined(ctx, backend, tx)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "wait mined proxy admin")
	} else if receipt.Status != types.ReceiptStatusSuccessful {
		return common.Address{}, errors.New("deploy proxy failed")
	}

	return proxyAdmin, nil
}

func deployFeeOracleV1(ctx context.Context, txOpts *bind.TransactOpts, backend *ethbackend.Backend,
	proxyAdmin common.Address, owner common.Address, fee *big.Int,
) (common.Address, error) {
	impl, tx, _, err := bindings.DeployFeeOracleV1(txOpts, backend)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "deploy fee oracle impl")
	}

	receipt, err := bind.WaitMined(ctx, backend, tx)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "wait mined fee oracle impl")
	} else if receipt.Status != types.ReceiptStatusSuccessful {
		return common.Address{}, errors.New("deploy fee oracle imple failed")
	}

	abi, err := bindings.FeeOracleV1MetaData.GetAbi()
	if err != nil {
		return common.Address{}, errors.Wrap(err, "get fee oracle abi")
	}

	enc, err := abi.Pack("initialize", owner, fee)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "encode fee oracle initializer")
	}

	proxy, tx, _, err := bindings.DeployTransparentUpgradeableProxy(txOpts, backend, impl, proxyAdmin, enc)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "deploy fee oracle proxy")
	}

	receipt, err = bind.WaitMined(ctx, backend, tx)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "wait mined transparent proxy")
	} else if receipt.Status != types.ReceiptStatusSuccessful {
		return common.Address{}, errors.New("deploy fee transparent proxy failed")
	}

	return proxy, nil
}

func deployPortal(ctx context.Context, txOpts *bind.TransactOpts, backend *ethbackend.Backend,
	proxyAdmin common.Address, owner common.Address, feeOracle common.Address, valSetID uint64,
	validators []bindings.Validator,
) (common.Address, error) {
	impl, tx, _, err := bindings.DeployOmniPortal(txOpts, backend)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "deploy portal impl")
	}

	receipt, err := bind.WaitMined(ctx, backend, tx)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "wait mined portal")
	} else if receipt.Status != types.ReceiptStatusSuccessful {
		return common.Address{}, errors.New("deploy portal failed")
	}

	abi, err := bindings.OmniPortalMetaData.GetAbi()
	if err != nil {
		return common.Address{}, errors.Wrap(err, "get portal abi")
	}

	enc, err := abi.Pack("initialize", owner, feeOracle, valSetID, validators)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "encode portal initializer")
	}

	proxy, tx, _, err := bindings.DeployTransparentUpgradeableProxy(txOpts, backend, impl, proxyAdmin, enc)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "deploy portal proxy")
	}

	receipt, err = bind.WaitMined(ctx, backend, tx)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "wait mined upgradable proxy")
	} else if receipt.Status != types.ReceiptStatusSuccessful {
		return common.Address{}, errors.New("deploy upgradable proxy failed")
	}

	return proxy, nil
}

func logBalance(ctx context.Context, backend *ethbackend.Backend, chain string, addr common.Address, name string,
) error {
	b, err := backend.BalanceAt(ctx, addr, nil)
	if err != nil {
		return errors.Wrap(err, "get balance")
	}

	bf, _ := b.Float64()
	bf /= params.Ether

	log.Info(ctx, "Provided public chain key balance",
		"chain", chain,
		"address", addr.Hex(),
		"balance", fmt.Sprintf("%.2f", bf),
		"key_name", name,
	)

	return nil
}

func mustHexToKey(privKeyHex string) *ecdsa.PrivateKey {
	privKey, err := crypto.HexToECDSA(strings.TrimPrefix(privKeyHex, "0x"))
	if err != nil {
		panic(err)
	}

	return privKey
}
