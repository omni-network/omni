package txmgr

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type DefaultFlagValues struct {
	NumConfirmations          uint64
	SafeAbortNonceTooLowCount uint64
	FeeLimitMultiplier        uint64
	FeeLimitThresholdGwei     float64
	ResubmissionTimeout       time.Duration
	NetworkTimeout            time.Duration
	TxSendTimeout             time.Duration
	TxNotInMempoolTimeout     time.Duration
}

type CLIConfig struct {
	ChainID                   uint64
	NumConfirmations          uint64
	SafeAbortNonceTooLowCount uint64
	FeeLimitMultiplier        uint64
	FeeLimitThresholdGwei     float64
	MinBaseFeeGwei            float64
	MinTipCapGwei             float64
	ResubmissionTimeout       time.Duration
	ReceiptQueryInterval      time.Duration
	NetworkTimeout            time.Duration
	TxSendTimeout             time.Duration
	TxNotInMempoolTimeout     time.Duration
}

var (
	//nolint:gochecknoglobals // should be configurable
	DefaultSenderFlagValues = DefaultFlagValues{
		NumConfirmations:          uint64(1),
		SafeAbortNonceTooLowCount: uint64(3),
		FeeLimitMultiplier:        uint64(5),
		FeeLimitThresholdGwei:     100.0,
		ResubmissionTimeout:       48 * time.Second,
		NetworkTimeout:            30 * time.Second,
		TxSendTimeout:             20 * time.Minute,
		TxNotInMempoolTimeout:     2 * time.Minute,
	}
)

func NewCLIConfig(chainID uint64, interval time.Duration, defaults DefaultFlagValues) CLIConfig {
	return CLIConfig{
		NumConfirmations:          defaults.NumConfirmations,
		SafeAbortNonceTooLowCount: defaults.SafeAbortNonceTooLowCount,
		FeeLimitMultiplier:        defaults.FeeLimitMultiplier,
		FeeLimitThresholdGwei:     defaults.FeeLimitThresholdGwei,
		ResubmissionTimeout:       defaults.ResubmissionTimeout,
		NetworkTimeout:            defaults.NetworkTimeout,
		TxSendTimeout:             defaults.TxSendTimeout,
		TxNotInMempoolTimeout:     defaults.TxNotInMempoolTimeout,
		ReceiptQueryInterval:      interval,
		ChainID:                   chainID,
	}
}

func (m CLIConfig) Check() error {
	if m.NumConfirmations == 0 {
		return errors.New("numConfirmations must not be 0")
	}
	if m.NetworkTimeout == 0 {
		return errors.New("must provide NetworkTimeout")
	}
	if m.FeeLimitMultiplier == 0 {
		return errors.New("must provide FeeLimitMultiplier")
	}
	if m.MinBaseFeeGwei < m.MinTipCapGwei {
		return errors.New("minBaseFee smaller than minTipCap",
			m.MinBaseFeeGwei, m.MinTipCapGwei)
	}
	if m.ResubmissionTimeout == 0 {
		return errors.New("must provide ResubmissionTimeout")
	}
	if m.ReceiptQueryInterval == 0 {
		return errors.New("must provide ReceiptQueryInterval")
	}
	if m.TxNotInMempoolTimeout == 0 {
		return errors.New("must provide TxNotInMempoolTimeout")
	}
	if m.SafeAbortNonceTooLowCount == 0 {
		return errors.New("safeAbortNonceTooLowCount must not be 0")
	}

	return nil
}

func externalSignerFn(external ExternalSigner, address common.Address, chainID uint64) SignerFn {
	signer := types.LatestSignerForChainID(big.NewInt(int64(chainID)))
	return func(ctx context.Context, from common.Address, tx *types.Transaction) (*types.Transaction, error) {
		if address != from {
			return nil, bind.ErrNotAuthorized
		}

		sig, err := external(ctx, signer.Hash(tx), from)
		if err != nil {
			return nil, errors.Wrap(err, "external sign")
		}

		res, err := tx.WithSignature(signer, sig[:])
		if err != nil {
			return nil, errors.Wrap(err, "set signature")
		}

		return res, nil
	}
}

// privateKeySignerFn returns a SignerFn that signs transactions with the given private key.
func privateKeySignerFn(key *ecdsa.PrivateKey, chainID uint64) SignerFn {
	from := crypto.PubkeyToAddress(key.PublicKey)
	signer := types.LatestSignerForChainID(big.NewInt(int64(chainID)))

	return func(_ context.Context, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		if address != from {
			return nil, bind.ErrNotAuthorized
		}
		signature, err := crypto.Sign(signer.Hash(tx).Bytes(), key)
		if err != nil {
			return nil, errors.Wrap(err, "could not sign transaction")
		}
		res, err := tx.WithSignature(signer, signature)
		if err != nil {
			return nil, errors.Wrap(err, "could not sign transaction")
		}

		return res, nil
	}
}

// ExternalSigner is a function that signs a transaction with a given address and returns the signature.
type ExternalSigner func(context.Context, common.Hash, common.Address) ([65]byte, error)

// SignerFn is a generic transaction signing function. It may be a remote signer so it takes a context.
// It also takes the address that should be used to sign the transaction with.
type SignerFn func(context.Context, common.Address, *types.Transaction) (*types.Transaction, error)

// SignerFactory creates a SignerFn that is bound to a specific chainID.
type SignerFactory func(chainID *big.Int) SignerFn

// Config houses parameters for altering the behavior of a simple.
type Config struct {
	Backend ethclient.Client

	// ResubmissionTimeout is the interval at which, if no previously
	// published transaction has been mined, the new tx with a bumped gas
	// price will be published. Only one publication at MaxGasPrice will be
	// attempted.
	ResubmissionTimeout time.Duration

	// The multiplier applied to fee suggestions to put a hard limit on fee increases.
	FeeLimitMultiplier uint64

	// Minimum threshold (in Wei) at which the FeeLimitMultiplier takes effect.
	// On low-fee networks, like test networks, this allows for arbitrary fee bumps
	// below this threshold.
	FeeLimitThreshold *big.Int

	// Minimum base fee (in Wei) to assume when determining tx fees.
	MinBaseFee *big.Int

	// Minimum tip cap (in Wei) to enforce when determining tx fees.
	MinTipCap *big.Int

	// ChainID is the chain ID of the L1 chain.
	ChainID *big.Int

	// TxSendTimeout is how long to wait for sending a transaction.
	// By default, it is unbounded. If set, this is recommended to be at least 20 minutes.
	TxSendTimeout time.Duration

	// TxNotInMempoolTimeout is how long to wait before aborting a transaction doSend if the transaction does not
	// make it to the mempool. If the tx is in the mempool, TxSendTimeout is used instead.
	TxNotInMempoolTimeout time.Duration

	// NetworkTimeout is the allowed duration for a single network request.
	// This is intended to be used for network requests that can be replayed.
	// todo(lazar): this should be handled by eth client
	NetworkTimeout time.Duration

	// RequireQueryInterval is the interval at which the tx manager will
	// query the backend to check for confirmations after a tx at a
	// specific gas price has been published.
	ReceiptQueryInterval time.Duration

	// NumConfirmations specifies how many blocks are need to consider a
	// transaction confirmed.
	NumConfirmations uint64

	// SafeAbortNonceTooLowCount specifies how many ErrNonceTooLow observations
	// are required to give up on a tx at a particular nonce without receiving
	// confirmation.
	SafeAbortNonceTooLowCount uint64

	// Signer is used to sign transactions when the gas price is increased.
	Signer SignerFn

	From common.Address
}

// NewConfigWithSigner returns a new txmgr config from the given CLI config and external signer.
func NewConfigWithSigner(cfg CLIConfig, external ExternalSigner, from common.Address, client ethclient.Client) (Config, error) {
	signer := externalSignerFn(external, from, cfg.ChainID)

	return newConfig(cfg, signer, from, client)
}

// NewConfig returns a new txmgr config from the given CLI config and private key.
func NewConfig(cfg CLIConfig, privateKey *ecdsa.PrivateKey, client ethclient.Client) (Config, error) {
	signer := privateKeySignerFn(privateKey, cfg.ChainID)
	addr := crypto.PubkeyToAddress(privateKey.PublicKey)

	return newConfig(cfg, signer, addr, client)
}

func newConfig(cfg CLIConfig, signer SignerFn, from common.Address, client ethclient.Client) (Config, error) {
	if err := cfg.Check(); err != nil {
		return Config{}, errors.New("invalid config", err)
	}

	feeLimitThreshold, err := GweiToWei(cfg.FeeLimitThresholdGwei)
	if err != nil {
		return Config{}, errors.Wrap(err, "invalid fee limit threshold")
	}

	minBaseFee, err := GweiToWei(cfg.MinBaseFeeGwei)
	if err != nil {
		return Config{}, errors.Wrap(err, "invalid min base fee")
	}

	minTipCap, err := GweiToWei(cfg.MinTipCapGwei)
	if err != nil {
		return Config{}, errors.Wrap(err, "invalid min tip cap")
	}

	chainID := big.NewInt(int64(cfg.ChainID))

	return Config{
		Backend:                   client,
		ResubmissionTimeout:       cfg.ResubmissionTimeout,
		FeeLimitMultiplier:        cfg.FeeLimitMultiplier,
		FeeLimitThreshold:         feeLimitThreshold,
		MinBaseFee:                minBaseFee,
		MinTipCap:                 minTipCap,
		ChainID:                   chainID,
		TxSendTimeout:             cfg.TxSendTimeout,
		TxNotInMempoolTimeout:     cfg.TxNotInMempoolTimeout,
		NetworkTimeout:            cfg.NetworkTimeout,
		ReceiptQueryInterval:      cfg.ReceiptQueryInterval,
		NumConfirmations:          cfg.NumConfirmations,
		SafeAbortNonceTooLowCount: cfg.SafeAbortNonceTooLowCount,
		Signer:                    signer,
		From:                      from,
	}, nil
}

func (m Config) Check() error {
	if m.Backend == nil {
		return errors.New("must provide the backend")
	}
	if m.NumConfirmations == 0 {
		return errors.New("numConfirmations must not be 0")
	}
	if m.NetworkTimeout == 0 {
		return errors.New("must provide NetworkTimeout")
	}
	if m.FeeLimitMultiplier == 0 {
		return errors.New("must provide FeeLimitMultiplier")
	}
	if m.MinBaseFee != nil && m.MinTipCap != nil && m.MinBaseFee.Cmp(m.MinTipCap) == -1 {
		return errors.New("minBaseFee smaller than minTipCap",
			m.MinBaseFee, m.MinTipCap)
	}
	if m.ResubmissionTimeout == 0 {
		return errors.New("must provide ResubmissionTimeout")
	}
	if m.ReceiptQueryInterval == 0 {
		return errors.New("must provide ReceiptQueryInterval")
	}
	if m.TxNotInMempoolTimeout == 0 {
		return errors.New("must provide TxNotInMempoolTimeout")
	}
	if m.SafeAbortNonceTooLowCount == 0 {
		return errors.New("safeAbortNonceTooLowCount must not be 0")
	}
	if m.Signer == nil {
		return errors.New("must provide the Signer")
	}
	if m.ChainID == nil {
		return errors.New("must provide the chainID")
	}

	return nil
}
