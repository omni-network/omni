package backend

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/txmgr"
	"github.com/omni-network/omni/test/e2e/types"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type Backend interface {
	bind.ContractBackend
	bind.DeployBackend
	ethereum.BlockNumberReader
	ethereum.ChainStateReader

	// Send provides access to underlying txmgr directly.
	Send(ctx context.Context, candidate txmgr.TxCandidate) (*ethtypes.Transaction, *ethtypes.Receipt, error)
}

var _ Backend = backend{}

type backend struct {
	ethclient.Client

	from  common.Address
	txMgr txmgr.TxManager
	name  string
}

func newBackend(chain types.EVMChain, rpcAddr string, privateKey *ecdsa.PrivateKey) (backend, error) {
	ethCl, err := ethclient.Dial(chain.Name, rpcAddr)
	if err != nil {
		return backend{}, errors.Wrap(err, "dial")
	}

	txMgr, err := newTxMgr(ethCl, chain, privateKey)
	if err != nil {
		return backend{}, errors.Wrap(err, "deploy tx txMgr")
	}

	return backend{
		Client: ethCl,
		from:   crypto.PubkeyToAddress(privateKey.PublicKey),
		txMgr:  txMgr,
		name:   chain.Name,
	}, nil
}

func (b backend) Send(ctx context.Context, candidate txmgr.TxCandidate) (*ethtypes.Transaction, *ethtypes.Receipt, error) {
	return b.txMgr.Send(ctx, candidate)
}

// SendTransaction intercepts the tx that bindings generates, strips fields, and passes
// it to the txmgr for reliable broadcasting.
func (b backend) SendTransaction(ctx context.Context, in *ethtypes.Transaction) error {
	candidate := txmgr.TxCandidate{
		TxData:   in.Data(),
		To:       in.To(),
		GasLimit: in.Gas(),
		Value:    in.Value(),
	}

	ctx = log.WithCtx(ctx, "req_id", randomHex7(), "chain", b.name)

	out, resp, err := b.txMgr.Send(ctx, candidate)
	if err != nil {
		return errors.Wrap(err, "txmgr send tx")
	}

	// TODO(corver): Maybe remove this as it is very noisy.
	log.Debug(ctx, "Backend sent tx",
		"nonce", out.Nonce(),
		"gas_used", resp.GasUsed,
		"status", resp.Status,
		"height", resp.BlockNumber.Uint64(),
	)

	*in = *out

	return nil
}

func randomHex7() string {
	bytes := make([]byte, 4)
	_, _ = rand.Read(bytes)
	hexString := hex.EncodeToString(bytes)

	// Trim the string to 7 characters
	if len(hexString) > 7 {
		hexString = hexString[:7]
	}

	return hexString
}
