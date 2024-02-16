package netman

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
)

func deployOmniContracts(ctx context.Context, chainID uint64, client *ethclient.Client, privKey *ecdsa.PrivateKey,
	valSetID uint64, validators []bindings.Validator,
) (common.Address, *bindings.OmniPortal, *bind.TransactOpts, error) {
	txOpts, err := newTxOpts(ctx, privKey, chainID)
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	// TODO: make these configurable
	owner := txOpts.From
	fee := new(big.Int).SetUint64(params.GWei)

	proxyAdmin, _, _, err := bindings.DeployProxyAdmin(txOpts, client)
	if err != nil {
		return common.Address{}, nil, nil, errors.Wrap(err, "deploy proxy admin")
	}

	feeOracle, err := deployFeeOracleV1(ctx, txOpts, client, proxyAdmin, owner, fee)
	if err != nil {
		return common.Address{}, nil, nil, errors.Wrap(err, "deploy fee oracle")
	}

	portal, err := deployPortal(ctx, txOpts, client, proxyAdmin, owner, feeOracle, valSetID, validators)
	if err != nil {
		return common.Address{}, nil, nil, errors.Wrap(err, "deploy portal")
	}

	contract, err := bindings.NewOmniPortal(portal, client)
	if err != nil {
		return common.Address{}, nil, nil, errors.Wrap(err, "new portal contract")
	}

	return portal, contract, txOpts, nil
}

func deployFeeOracleV1(ctx context.Context, txOpts *bind.TransactOpts, client *ethclient.Client,
	proxyAdmin common.Address, owner common.Address, fee *big.Int,
) (common.Address, error) {
	impl, tx, _, err := bindings.DeployFeeOracleV1(txOpts, client)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "deploy fee oracle impl")
	}

	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil || receipt.Status != types.ReceiptStatusSuccessful {
		return common.Address{}, errors.Wrap(err, "wait fee oracle impl mined")
	}

	abi, err := bindings.FeeOracleV1MetaData.GetAbi()
	if err != nil {
		return common.Address{}, errors.Wrap(err, "get fee oracle v1 abi")
	}

	enc, err := abi.Pack("initialize", owner, fee)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "pack fee oracle v1 initializer")
	}

	proxy, _, _, err := bindings.DeployTransparentUpgradeableProxy(txOpts, client, impl, proxyAdmin, enc)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "deploy fee oracle proxy")
	}

	receipt, err = bind.WaitMined(ctx, client, tx)
	if err != nil || receipt.Status != types.ReceiptStatusSuccessful {
		return common.Address{}, errors.Wrap(err, "wait mined fee oracle proxy")
	}

	return proxy, nil
}

func deployPortal(ctx context.Context, txOpts *bind.TransactOpts, client *ethclient.Client,
	proxyAdmin common.Address, owner common.Address, feeOracle common.Address, valSetID uint64,
	validators []bindings.Validator,
) (common.Address, error) {
	impl, tx, _, err := bindings.DeployOmniPortal(txOpts, client)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "deploy portal impl")
	}

	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil || receipt.Status != types.ReceiptStatusSuccessful {
		return common.Address{}, errors.Wrap(err, "wait mined portal proxy")
	}

	abi, err := bindings.OmniPortalMetaData.GetAbi()
	if err != nil {
		return common.Address{}, errors.Wrap(err, "get portal abi")
	}

	enc, err := abi.Pack("initialize", owner, feeOracle, valSetID, validators)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "pack portal initializer")
	}

	proxy, tx, _, err := bindings.DeployTransparentUpgradeableProxy(txOpts, client, impl, proxyAdmin, enc)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "deploy portal proxy")
	}

	receipt, err = bind.WaitMined(ctx, client, tx)
	if err != nil || receipt.Status != types.ReceiptStatusSuccessful {
		return common.Address{}, errors.Wrap(err, "wait mined portal proxy")
	}

	return proxy, nil
}

func newTxOpts(ctx context.Context, privKey *ecdsa.PrivateKey, chainID uint64) (*bind.TransactOpts, error) {
	txOpts, err := bind.NewKeyedTransactorWithChainID(privKey, big.NewInt(int64(chainID)))
	if err != nil {
		return nil, errors.Wrap(err, "keyed tx ops")
	}

	txOpts.Context = ctx

	return txOpts, nil
}

func logBalance(ctx context.Context, client *ethclient.Client, privkey *ecdsa.PrivateKey, name string) error {
	addr := crypto.PubkeyToAddress(privkey.PublicKey)

	b, err := client.BalanceAt(ctx, addr, nil)
	if err != nil {
		return errors.Wrap(err, "get balance")
	}

	bf, _ := b.Float64()
	bf /= params.Ether

	log.Info(ctx, "Provided public chain key balance",
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
