package ethbackend

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"math/big"
	"time"

	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/fireblocks"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/txmgr"
	"github.com/omni-network/omni/lib/umath"

	k1 "github.com/cometbft/cometbft/crypto/secp256k1"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type account struct {
	from       common.Address
	privateKey *ecdsa.PrivateKey // Either local private key is set,
	fireCl     fireblocks.Client // or, Fireblocks is used
	txMgr      txmgr.TxManager
}

type Backend struct {
	ethclient.Client

	accounts    map[common.Address]account
	chainName   string
	chainID     uint64
	blockPeriod time.Duration
}

// NewFireBackend returns a backend that supports all accounts supported by the configured fireblocks client.
// Note that private keys can still be added via AddAccount.
func NewFireBackend(ctx context.Context, chainName string, chainID uint64, blockPeriod time.Duration, ethCl ethclient.Client, fireCl fireblocks.Client) (*Backend, error) {
	accs, err := fireCl.Accounts(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "fireblocks accounts")
	}

	accounts := make(map[common.Address]account)
	for addr := range accs {
		txMgr, err := newFireblocksTxMgr(ethCl, chainName, chainID, blockPeriod, addr, fireCl)
		if err != nil {
			return nil, errors.Wrap(err, "new txmgr")
		}

		accounts[addr] = account{
			from:   addr,
			fireCl: fireCl,
			txMgr:  txMgr,
		}
	}

	return &Backend{
		Client:      ethCl,
		accounts:    accounts,
		chainName:   chainName,
		chainID:     chainID,
		blockPeriod: blockPeriod,
	}, nil
}

// NewAnvilBackend returns a backend with all pre-funded anvil dev accounts.
func NewAnvilBackend(chainName string, chainID uint64, blockPeriod time.Duration, ethCl ethclient.Client) (*Backend, error) {
	return NewBackend(chainName, chainID, blockPeriod, ethCl, anvil.DevPrivateKeys()...)
}

// NewBackend returns a new backend backed by in-memory private keys.
func NewBackend(chainName string, chainID uint64, blockPeriod time.Duration, ethCl ethclient.Client, privateKeys ...*ecdsa.PrivateKey) (*Backend, error) {
	accounts := make(map[common.Address]account)
	for _, pk := range privateKeys {
		txMgr, err := newTxMgr(ethCl, chainName, chainID, blockPeriod, pk)
		if err != nil {
			return nil, errors.Wrap(err, "new txmgr")
		}

		addr := crypto.PubkeyToAddress(pk.PublicKey)
		accounts[addr] = account{
			from:       addr,
			privateKey: pk,
			txMgr:      txMgr,
		}
	}

	return &Backend{
		Client:      ethCl,
		accounts:    accounts,
		chainName:   chainName,
		chainID:     chainID,
		blockPeriod: blockPeriod,
	}, nil
}

// AddAccount adds a in-memory private key account to the backend.
// Note this can be called even if other accounts are fireblocks based.
func (b *Backend) AddAccount(privkey *ecdsa.PrivateKey) (common.Address, error) {
	txMgr, err := newTxMgr(b.Client, b.chainName, b.chainID, b.blockPeriod, privkey)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "new txmgr")
	}

	addr := crypto.PubkeyToAddress(privkey.PublicKey)

	b.accounts[addr] = account{
		from:       addr,
		privateKey: privkey,
		txMgr:      txMgr,
	}

	return addr, nil
}

func (b *Backend) Chain() (string, uint64) {
	return b.chainName, b.chainID
}

func (b *Backend) Sign(ctx context.Context, from common.Address, input [32]byte) ([65]byte, error) {
	acc, ok := b.accounts[from]
	if !ok {
		return [65]byte{}, errors.New("unknown from address", "from", from)
	} else if acc.privateKey == nil {
		return acc.fireCl.Sign(ctx, input, from)
	}

	pk := k1.PrivKey(crypto.FromECDSA(acc.privateKey))

	return k1util.Sign(pk, input)
}

func (b *Backend) PublicKey(from common.Address) (*ecdsa.PublicKey, error) {
	acc, ok := b.accounts[from]
	if !ok {
		return nil, errors.New("unknown from address", "from", from)
	}

	return &acc.privateKey.PublicKey, nil
}

func (b *Backend) Send(ctx context.Context, from common.Address, candidate txmgr.TxCandidate) (*ethtypes.Transaction, *ethtypes.Receipt, error) {
	acc, ok := b.accounts[from]
	if !ok {
		return nil, nil, errors.New("unknown from address", "from", from)
	}

	return acc.txMgr.Send(ctx, candidate)
}

// WaitMined waits for the transaction to be mined and asserts the receipt is successful.
func (b *Backend) WaitMined(ctx context.Context, tx *ethtypes.Transaction) (*ethtypes.Receipt, error) {
	rec, err := bind.WaitMined(ctx, b, tx)
	if err != nil {
		return nil, errors.Wrap(err, "wait mined", "chain", b.chainName)
	} else if rec.Status != ethtypes.ReceiptStatusSuccessful {
		return rec, errors.New("receipt status unsuccessful", "status", rec.Status, "tx", tx.Hash())
	}

	return rec, nil
}

// BindOpts returns a new TransactOpts for interacting with bindings based contracts for the provided account.
// The TransactOpts are partially stubbed, since txmgr handles nonces and signing.
//
// Do not cache or store the TransactOpts, as they are not safe for concurrent use (pointer).
// Rather create a new TransactOpts for each transaction.
func (b *Backend) BindOpts(ctx context.Context, from common.Address) (*bind.TransactOpts, error) {
	if header, err := b.HeaderByNumber(ctx, nil); err != nil {
		return nil, errors.Wrap(err, "header by number")
	} else if header.BaseFee == nil {
		return nil, errors.New("only dynamic transaction Backends supported")
	}

	_, ok := b.accounts[from]
	if !ok {
		return nil, errors.New("unknown from address", "from", from)
	}

	// Stub nonce and signer since txmgr will handle this.
	// Bindings will estimate gas.
	return &bind.TransactOpts{
		From:  from,
		Nonce: big.NewInt(1),
		Signer: func(from common.Address, tx *ethtypes.Transaction) (*ethtypes.Transaction, error) {
			resp, err := tx.WithSignature(backendStubSigner{}, from[:])
			if err != nil {
				return nil, errors.Wrap(err, "with signature")
			}

			return resp, nil
		},
		Context: log.WithCtx(ctx, "dest_chain", b.chainName, "from_addr", from.Hex()[:10]),
	}, nil
}

// SendTransaction intercepts the tx that bindings generates, extracts the from address (assuming the
// backendStubSigner was used), the strips fields, and passes it to the txmgr for reliable broadcasting.
func (b *Backend) SendTransaction(ctx context.Context, in *ethtypes.Transaction) error {
	from, err := backendStubSigner{}.Sender(in)
	if err != nil {
		return errors.Wrap(err, "from signer sender")
	}

	acc, ok := b.accounts[from]
	if !ok {
		return errors.New("unknown from address", "from", from)
	}

	candidate := txmgr.TxCandidate{
		TxData:   in.Data(),
		To:       in.To(),
		GasLimit: in.Gas(),
		Value:    in.Value(),
	}

	ctx = log.WithCtx(ctx, "req_id", randomHex7(), "chain", b.chainName)

	out, resp, err := acc.txMgr.Send(ctx, candidate)
	if err != nil {
		return errors.Wrap(err, "txmgr send tx")
	}

	// TODO(corver): Maybe remove this as it is very noisy.
	log.Debug(ctx, "Backend sent tx",
		"nonce", out.Nonce(),
		"gas_used", resp.GasUsed,
		"status", resp.Status,
		"height", resp.BlockNumber.Uint64(),
		"tx", out.Hash(),
	)

	*in = *out //nolint:govet // Copy lock (caches) isn't a problem since we are overwriting the object.

	return nil
}

// EnsureSynced returns an error if the backend is not synced.
func (b *Backend) EnsureSynced(ctx context.Context) error {
	syncing, err := b.SyncProgress(ctx)
	if ethclient.IsErrMethodNotAvailable(err) {
		return nil // Assume synced if method not available.
	} else if err != nil {
		return err
	} else if syncing == nil {
		return nil // Syncing is nil if node is not syncing.
	} else if !syncing.Done() {
		return errors.New("backend not synced",
			"lag", umath.SubtractOrZero(syncing.HighestBlock, syncing.CurrentBlock),
			"indexing", syncing.TxIndexRemainingBlocks,
		)
	}

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

var _ ethtypes.Signer = backendStubSigner{}

// backendStubSigner is a stub signer that is used by bindings to sign transactions for a Backend.
// It just encodes the from address into the signature, since there is no other way to pass the from
// address from TxOpts to the Backend. The Backend then extracts the from address,
// gets the account txmgr, and sends the transaction.
type backendStubSigner struct {
	ethtypes.Signer
}

func (backendStubSigner) ChainID() *big.Int {
	return new(big.Int)
}

func (backendStubSigner) Sender(tx *ethtypes.Transaction) (common.Address, error) {
	// Convert R,S,V into a 20 byte signature into:
	// R: 8 bytes uint64
	// S: 8 bytes uint64
	// V: 4 bytes uint32

	v, r, s := tx.RawSignatureValues()

	addrLen := len(common.Address{})

	if len(r.Bytes()) > addrLen {
		return common.Address{}, errors.New("invalid r length", "length", len(r.Bytes()))
	}
	if s.Uint64() != 0 {
		return common.Address{}, errors.New("non-empty s [BUG]", "length", len(s.Bytes()))
	}
	if v.Uint64() != 0 {
		return common.Address{}, errors.New("non-empty v [BUG]", "length", len(v.Bytes()))
	}

	addr := make([]byte, addrLen)
	// big.Int.Bytes() truncates leading zeros, so we need to left pad the address to 20 bytes.
	pad := addrLen - len(r.Bytes())
	copy(addr[pad:], r.Bytes())

	return common.Address(addr), nil
}

//nolint:nonamedreturns // Ambiguous return values
func (backendStubSigner) SignatureValues(_ *ethtypes.Transaction, sig []byte) (r, s, v *big.Int, err error) {
	if len(sig) != len(common.Address{}) {
		return nil, nil, nil, errors.New("invalid from signature length", "length", len(sig))
	}

	// Set the 20 byte signature (from address) as R
	r = new(big.Int).SetBytes(sig)
	s = new(big.Int) // 0
	v = new(big.Int) // 0

	return r, s, v, nil
}
