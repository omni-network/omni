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
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
)

func deployContract(ctx context.Context, chainID uint64, client *ethclient.Client, privKey *ecdsa.PrivateKey,
) (common.Address, *bindings.OmniPortal, *bind.TransactOpts, error) {
	txOpts, err := newTxOpts(ctx, privKey, chainID)
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	addr, _, _, err := bindings.DeployOmniPortal(txOpts, client)
	if err != nil {
		return common.Address{}, nil, nil, errors.Wrap(err, "deploy portal contract")
	}

	contract, err := bindings.NewOmniPortal(addr, client)
	if err != nil {
		return common.Address{}, nil, nil, errors.Wrap(err, "new portal contract")
	}

	return addr, contract, txOpts, nil
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
