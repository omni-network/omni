package ethbackend

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"math/big"
	"sync"
	"time"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/fireblocks"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/tracer"
	"github.com/omni-network/omni/lib/txmgr"
	"github.com/omni-network/omni/lib/umath"

	k1 "github.com/cometbft/cometbft/crypto/secp256k1"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"golang.org/x/sync/semaphore"
)

// mempoolLimit is the maximum number of transactions that can be in-flight at once.
const mempoolLimit = 16

type account struct {
	from       common.Address
	sema       *semaphore.Weighted
	privateKey *ecdsa.PrivateKey // Either local private key is set,
	fireCl     fireblocks.Client // or, Fireblocks is used
	txMgr      txmgr.TxManager
}

type Backend struct {
	ethclient.Client

	mu          sync.RWMutex
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
			sema:   semaphore.NewWeighted(mempoolLimit),
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

// NewDevBackend returns a backend with all pre-funded anvil dev accounts.
func NewDevBackend(chainName string, chainID uint64, blockPeriod time.Duration, ethCl ethclient.Client) (*Backend, error) {
	return NewBackend(chainName, chainID, blockPeriod, ethCl, append(eoa.DevPrivateKeys(), anvil.DevPrivateKeys()...)...)
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
			sema:       semaphore.NewWeighted(mempoolLimit),
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

	b.mu.Lock()
	defer b.mu.Unlock()

	b.accounts[addr] = account{
		from:       addr,
		privateKey: privkey,
		txMgr:      txMgr,
		sema:       semaphore.NewWeighted(mempoolLimit),
	}

	return addr, nil
}

func (b *Backend) Chain() (string, uint64) {
	return b.chainName, b.chainID
}

func (b *Backend) Sign(ctx context.Context, from common.Address, input [32]byte) ([65]byte, error) {
	b.mu.RLock()
	acc, ok := b.accounts[from]
	b.mu.RUnlock()

	if !ok {
		return [65]byte{}, errors.New("unknown from address", "from", from)
	} else if acc.privateKey == nil {
		return acc.fireCl.Sign(ctx, input, from)
	}

	pk := k1.PrivKey(crypto.FromECDSA(acc.privateKey))

	return k1util.Sign(pk, input)
}

func (b *Backend) PublicKey(from common.Address) (*ecdsa.PublicKey, error) {
	b.mu.RLock()
	acc, ok := b.accounts[from]
	b.mu.RUnlock()

	if !ok {
		return nil, errors.New("unknown from address", "from", from)
	}

	return &acc.privateKey.PublicKey, nil
}

func (b *Backend) Send(ctx context.Context, from common.Address, candidate txmgr.TxCandidate) (*ethtypes.Transaction, *ethclient.Receipt, error) {
	ctx, span := tracer.Start(ctx, "ethbackend/send")
	defer span.End()

	b.mu.RLock()
	acc, ok := b.accounts[from]
	b.mu.RUnlock()

	if !ok {
		return nil, nil, errors.New("unknown from address", "from", from)
	}

	if err := acc.sema.Acquire(ctx, 1); err != nil {
		return nil, nil, errors.Wrap(err, "acquire semaphore")
	}
	defer acc.sema.Release(1)

	return acc.txMgr.Send(ctx, candidate)
}

// WaitMined waits for the transaction to be mined and asserts the receipt is successful.
func (b *Backend) WaitMined(ctx context.Context, tx *ethtypes.Transaction) (*ethclient.Receipt, error) {
	ctx, span := tracer.Start(ctx, "ethbackend/wait_mined")
	defer span.End()

	rec, err := waitMinedHash(ctx, b, tx.Hash())
	if err != nil {
		return nil, errors.Wrap(err, "wait mined", "chain", b.chainName, "tx", tx.Hash())
	} else if rec.Status != ethtypes.ReceiptStatusSuccessful {
		return rec, errors.New("receipt status unsuccessful", "status", rec.Status, "tx", tx.Hash(), "chain", b.chainName)
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

	b.mu.RLock()
	_, ok := b.accounts[from]
	b.mu.RUnlock()

	if !ok {
		return nil, errors.New("unknown from address", "from", from)
	}

	// Stub nonce and signer since txmgr will handle this.
	// Bindings will estimate gas.
	return &bind.TransactOpts{
		From:  from,
		Nonce: bi.One(),
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

type nonceKey struct{}

// WithReservedNonce returns a copy of the context including the next reserved nonce (also returned).
// This is useful to ensure sequential nonces with parallel transaction sending/mining..
func (b *Backend) WithReservedNonce(ctx context.Context, from common.Address) (context.Context, uint64, error) {
	b.mu.RLock()
	acc, ok := b.accounts[from]
	b.mu.RUnlock()
	if !ok {
		return nil, 0, errors.New("unknown from address", "from", from)
	}

	nonce, err := acc.txMgr.ReserveNextNonce(ctx)
	if err != nil {
		return nil, 0, errors.Wrap(err, "reserve next nonce")
	}

	return context.WithValue(ctx, nonceKey{}, nonce), nonce, nil
}

// SendTransaction intercepts the tx that bindings generates, extracts the from address (if the
// backendStubSigner was used), the strips fields, and passes it to the txmgr for reliable broadcasting.
func (b *Backend) SendTransaction(ctx context.Context, in *ethtypes.Transaction) error {
	ctx, span := tracer.Start(ctx, "ethbackend/send_transaction")
	defer span.End()

	if !(backendStubSigner{}).IsStub(in) {
		return b.Client.SendTransaction(ctx, in)
	}

	from, err := backendStubSigner{}.Sender(in)
	if err != nil {
		return errors.Wrap(err, "from signer sender")
	}

	b.mu.RLock()
	acc, ok := b.accounts[from]
	b.mu.RUnlock()

	if !ok {
		return errors.New("unknown from address", "from", from)
	}

	if err := acc.sema.Acquire(ctx, 1); err != nil {
		return errors.Wrap(err, "acquire semaphore")
	}
	defer acc.sema.Release(1)

	// Get optional nonce from context. Txmgr handles nonce if nil.
	var noncePtr *uint64
	if nonce, ok := ctx.Value(nonceKey{}).(uint64); ok {
		noncePtr = &nonce
	}

	candidate := txmgr.TxCandidate{
		TxData:   in.Data(),
		To:       in.To(),
		GasLimit: acc.txMgr.BumpGasLimit(in.Gas()), // If gas already estimated, just bump that (instead of re-estimating).
		Value:    in.Value(),
		Nonce:    noncePtr,
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
	progress, syncing, err := b.ProgressIfSyncing(ctx)
	if ethclient.IsErrMethodNotAvailable(err) {
		return nil // Assume synced if method not available.
	} else if err != nil {
		return err
	} else if !syncing {
		return nil
	} else if !progress.Done() {
		return errors.New("backend not synced",
			"lag", umath.SubtractOrZero(progress.HighestBlock, progress.CurrentBlock),
			"indexing", progress.TxIndexRemainingBlocks,
		)
	}

	return nil
}

// WaitConfirmed is similar to WaitMined, except that it also waits for the (first) receipt block to be "confirmed".
// "Confirmed" depends on the evmchain.Reorgs value, if false, it is just calls WaitMined and returns the receipt.
// Otherwise, it also waits for "safe" confirmation (this is very slow up to 6min).
// It returns an error if the (first) receipt is reorged out (even though the tx may still be included
// in a different block).
func (b *Backend) WaitConfirmed(ctx context.Context, tx *ethtypes.Transaction) (*ethclient.Receipt, error) {
	ctx, span := tracer.Start(ctx, "ethbackend/wait_confirmed")
	defer span.End()

	rec, err := b.WaitMined(ctx, tx)
	if err != nil {
		// Return receipt with error, in case it was mined but reverted
		return rec, errors.Wrap(err, "wait mined")
	}

	meta, ok := evmchain.MetadataByID(b.chainID)
	if !ok {
		return nil, errors.New("unknown chain id", "chain", b.chainName)
	} else if !meta.Reorgs {
		return rec, nil // No need to confirm if reorgs are not expected.
	}

	if err := b.confirmReceipt(ctx, rec); err != nil {
		return nil, errors.Wrap(err, "confirm receipt")
	}

	return rec, nil
}

func (b *Backend) confirmReceipt(ctx context.Context, rec *ethclient.Receipt) error {
	ticker := time.NewTicker(b.blockPeriod * 4)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return errors.Wrap(ctx.Err(), "confirm receipt timeout")
		case <-ticker.C:
			safeHead, err := b.HeaderByType(ctx, ethclient.HeadSafe)
			if err != nil {
				continue // Ignore RPC errors, just retry, same as in WaitMined.
			}

			// Check if the receipt's block number is safe.
			if bi.LT(safeHead.Number, rec.BlockNumber) {
				continue // Receipt block is not safe yet.
			}

			// Walk up the safe chain to confirm the receipt's block hash.
			safe := safeHead
			for {
				if safe.Hash() == rec.BlockHash || safe.ParentHash == rec.BlockHash {
					return nil // Receipt block hash is confirmed as safe.
				}

				if bi.EQ(safe.Number, rec.BlockNumber) {
					return errors.New("receipt block reorged", "height", rec.BlockNumber, "actual_block_hash", safe.Hash(), "receipt_block_hash", rec.BlockHash)
				}

				// Move to the parent block.
				safe, err = b.HeaderByHash(ctx, safe.ParentHash)
				if err != nil {
					// Ignore RPC errors, just retry, same as in WaitMined.
					break
				}
			}
		}
	}
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
	return bi.Zero()
}

func (backendStubSigner) IsStub(tx *ethtypes.Transaction) bool {
	v, r, s := tx.RawSignatureValues()

	return bi.IsZero(v) && bi.IsZero(s) && len(r.Bytes()) <= common.AddressLength
}

func (backendStubSigner) Sender(tx *ethtypes.Transaction) (common.Address, error) {
	// Convert R,S,V into a 20 byte signature into:
	// R: 8 bytes uint64
	// S: 8 bytes uint64
	// V: 4 bytes uint32

	v, r, s := tx.RawSignatureValues()

	if len(r.Bytes()) > common.AddressLength {
		return common.Address{}, errors.New("invalid r length", "length", len(r.Bytes()))
	}
	if !bi.IsZero(s) {
		return common.Address{}, errors.New("non-empty s [BUG]", "length", len(s.Bytes()))
	}
	if !bi.IsZero(v) {
		return common.Address{}, errors.New("non-empty v [BUG]", "length", len(v.Bytes()))
	}

	var addr [common.AddressLength]byte
	// big.Int.Bytes() truncates leading zeros, so we need to left pad the address to 20 bytes.
	pad := common.AddressLength - len(r.Bytes())
	copy(addr[pad:], r.Bytes())

	return addr, nil
}

//nolint:nonamedreturns // Ambiguous return values
func (backendStubSigner) SignatureValues(_ *ethtypes.Transaction, sig []byte) (r, s, v *big.Int, err error) {
	if len(sig) != len(common.Address{}) {
		return nil, nil, nil, errors.New("invalid from signature length", "length", len(sig))
	}

	// Set the 20 byte signature (from address) as R
	r = new(big.Int).SetBytes(sig)
	s = bi.Zero() // 0
	v = bi.Zero() // 0

	return r, s, v, nil
}

// waitMinedHash waits for a transaction with the provided hash to be mined on the blockchain.
// It stops waiting when the context is canceled.
// This is a copy of bind.WaitMinedHash returning omni custom Receipt type.
func waitMinedHash(ctx context.Context, b *Backend, hash common.Hash) (*ethclient.Receipt, error) {
	ticker := time.NewTimer(0) // Check immediately
	defer ticker.Stop()

	ctx = log.WithCtx(ctx, "tx", hash.Hex(), "chain", b.chainName)

	var logFailedOnce bool
	for {
		select {
		case <-ctx.Done():
			return nil, errors.Wrap(ctx.Err(), "wait mined timeout")
		case <-ticker.C:
			ticker.Reset(time.Second) // Then every 1 second

			receipt, err := b.TxReceipt(ctx, hash)
			if err == nil {
				return receipt, nil
			} else if errors.Is(err, ethereum.NotFound) {
				continue
			} else if !logFailedOnce {
				// Only log this once, avoid spamming the logs.
				log.DebugErr(ctx, "Failed querying receipt", err)
				logFailedOnce = true
			}
		}
	}
}
