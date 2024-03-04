package backend

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"math/big"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/txmgr"

	k1 "github.com/cometbft/cometbft/crypto/secp256k1"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// Backend represents an ethclient with one or more private keys (accounts).
// It can be used to send transactions and sign data on behalf of these accounts.
// It is designed to be used with bindings based contracts.
// It improves on normal bindings transactor by using txmgr for reliable sending.
type Backend interface {
	ethclient.Client

	// AddAccount adds a new account to the backend, it returns the address of the account for convenience.
	AddAccount(privkey *ecdsa.PrivateKey) (common.Address, error)

	// BindOpts returns a new TransactOpts for interacting with bindings based contracts for the provided account.
	// The TransactOpts are partially stubbed, since txmgr handles nonces and signing.
	//
	// Do not cache or store the TransactOpts, as they are not safe for concurrent use (pointer).
	// Rather create a new TransactOpts for each transaction.
	BindOpts(ctx context.Context, account common.Address) (*bind.TransactOpts, error)

	// Send provides access to underlying txmgr directly of the provided account.
	Send(ctx context.Context, account common.Address, candidate txmgr.TxCandidate) (*ethtypes.Transaction, *ethtypes.Receipt, error)

	// WaitMined waits for the given transaction to be mined and returns the receipt.
	WaitMined(ctx context.Context, tx *ethtypes.Transaction) (*ethtypes.Receipt, error)

	// Sign signs the given input with the account's private key.
	// It returns a 65-byte Ethereum signature in the format [R || S || V].
	Sign(from common.Address, input [32]byte) ([65]byte, error)

	// Chain returns the chainName and chainID of the chain.
	Chain() (string, uint64)
}

var _ Backend = (*backend)(nil)

type account struct {
	from       common.Address
	privateKey *ecdsa.PrivateKey
	txMgr      txmgr.TxManager
}

type backend struct {
	ethclient.Client

	accounts    map[common.Address]account
	chainName   string
	chainID     uint64
	blockPeriod time.Duration
}

func NewBackend(chainName string, chainID uint64, blockPeriod time.Duration, ethCl ethclient.Client, privateKeys ...*ecdsa.PrivateKey) (Backend, error) {
	return newBackend(chainName, chainID, blockPeriod, ethCl, privateKeys...)
}

func newBackend(chainName string, chainID uint64, blockPeriod time.Duration, ethCl ethclient.Client, privateKeys ...*ecdsa.PrivateKey) (*backend, error) {
	accounts := make(map[common.Address]account)
	for _, pk := range privateKeys {
		txMgr, err := newTxMgr(ethCl, chainName, chainID, blockPeriod, pk)
		if err != nil {
			return nil, errors.Wrap(err, "deploy tx txMgr")
		}

		addr := crypto.PubkeyToAddress(pk.PublicKey)
		accounts[addr] = account{
			from:       addr,
			privateKey: pk,
			txMgr:      txMgr,
		}
	}

	return &backend{
		Client:      ethCl,
		accounts:    accounts,
		chainName:   chainName,
		chainID:     chainID,
		blockPeriod: blockPeriod,
	}, nil
}

func (b *backend) AddAccount(privkey *ecdsa.PrivateKey) (common.Address, error) {
	txMgr, err := newTxMgr(b.Client, b.chainName, b.chainID, b.blockPeriod, privkey)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "deploy tx txMgr")
	}

	addr := crypto.PubkeyToAddress(privkey.PublicKey)

	b.accounts[addr] = account{
		from:       addr,
		privateKey: privkey,
		txMgr:      txMgr,
	}

	return addr, nil
}

func (b *backend) Chain() (string, uint64) {
	return b.chainName, b.chainID
}

func (b *backend) Sign(from common.Address, input [32]byte) ([65]byte, error) {
	acc, ok := b.accounts[from]
	if !ok {
		return [65]byte{}, errors.New("unknown from address", "from", from)
	}

	pk := k1.PrivKey(crypto.FromECDSA(acc.privateKey))

	return k1util.Sign(pk, input)
}

func (b *backend) Send(ctx context.Context, from common.Address, candidate txmgr.TxCandidate) (*ethtypes.Transaction, *ethtypes.Receipt, error) {
	acc, ok := b.accounts[from]
	if !ok {
		return nil, nil, errors.New("unknown from address", "from", from)
	}

	return acc.txMgr.Send(ctx, candidate)
}

func (b *backend) WaitMined(ctx context.Context, tx *ethtypes.Transaction) (*ethtypes.Receipt, error) {
	rec, err := bind.WaitMined(ctx, b, tx)
	if err != nil {
		return nil, errors.Wrap(err, "wait mined", "chain", b.chainName)
	} else if rec.Status == ethtypes.ReceiptStatusSuccessful {
		return rec, nil
	}

	return rec, nil
}

// BindOpts returns a new TransactOpts for interacting with bindings based contracts for the provided account.
// The TransactOpts are partially stubbed, since txmgr handles nonces and signing.
//
// Do not cache or store the TransactOpts, as they are not safe for concurrent use (pointer).
// Rather create a new TransactOpts for each transaction.
func (b *backend) BindOpts(ctx context.Context, from common.Address) (*bind.TransactOpts, error) {
	if header, err := b.HeaderByNumber(ctx, nil); err != nil {
		return nil, errors.Wrap(err, "header by number")
	} else if header.BaseFee == nil {
		return nil, errors.New("only dynamic transaction backends supported")
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
		Context: ctx,
	}, nil
}

// SendTransaction intercepts the tx that bindings generates, extracts the from address (assuming the
// backendStubSigner was used), the strips fields, and passes it to the txmgr for reliable broadcasting.
func (b *backend) SendTransaction(ctx context.Context, in *ethtypes.Transaction) error {
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

var _ ethtypes.Signer = backendStubSigner{}

// backendStubSigner is a stub signer that is used by bindings to sign transactions for a backend.
// It just encodes the from address into the signature, since there is no other way to pass the from
// address from TxOpts to the backend. The backend then extracts the from address,
// gets the account txmgr, and sends the transaction.
type backendStubSigner struct {
	ethtypes.Signer
}

func (f backendStubSigner) ChainID() *big.Int {
	return new(big.Int)
}

func (f backendStubSigner) Sender(tx *ethtypes.Transaction) (common.Address, error) {
	// Convert R,S,V into a 20 byte signature into:
	// R: 8 bytes uint64
	// S: 8 bytes uint64
	// V: 4 bytes uint32

	v, r, s := tx.RawSignatureValues()

	if len(r.Bytes()) != 20 {
		return common.Address{}, errors.New("invalid r length", "length", len(r.Bytes()))
	}
	if s.Uint64() != 0 {
		return common.Address{}, errors.New("non-empty s [BUG]", "length", len(s.Bytes()))
	}
	if v.Uint64() != 0 {
		return common.Address{}, errors.New("non-empty v [BUG]", "length", len(v.Bytes()))
	}

	addr := make([]byte, 20)
	copy(addr, r.Bytes())

	return common.Address(addr), nil
}

//nolint:nonamedreturns // Ambiguous return values
func (f backendStubSigner) SignatureValues(_ *ethtypes.Transaction, sig []byte) (r, s, v *big.Int, err error) {
	if len(sig) != len(common.Address{}) {
		return nil, nil, nil, errors.New("invalid from signature length", "length", len(sig))
	}

	// Set the 20 byte signature (from address) as R
	r = new(big.Int).SetBytes(sig)
	s = new(big.Int) // 0
	v = new(big.Int) // 0

	return r, s, v, nil
}
