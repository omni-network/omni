package relayer

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/txmgr/metrics"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethlog "github.com/ethereum/go-ethereum/log"
	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/netconf"

	opcrypto "github.com/ethereum-optimism/optimism/op-service/crypto"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/omni-network/omni/lib/xchain"
)

var _ Sender = (*OpSender)(nil)

type OpSender struct {
	txMgrs  map[uint64]txmgr.TxManager
	portals map[uint64]common.Address
	abi     *abi.ABI
}

func NewOpSender(ctx context.Context, chains []netconf.Chain, rpcClientPerChain map[uint64]*ethclient.Client,
	privateKey *ecdsa.PrivateKey) (OpSender, error) {
	txMgrs := make(map[uint64]txmgr.TxManager)
	portals := make(map[uint64]common.Address)

	l := WrapLogger(ctx)

	for _, chain := range chains {
		cfg, err := NewTxMgrConfig(ctx, txmgr.NewCLIConfig(chain.RPCURL, txmgr.DefaultBatcherFlagValues),
			privateKey, rpcClientPerChain[chain.ID])
		if err != nil {
			return OpSender{}, err
		}

		txMgr, err := initTxMgr(cfg, l)
		if err != nil {
			return OpSender{}, err
		}

		txMgrs[chain.ID] = txMgr
		portals[chain.ID] = common.HexToAddress(chain.PortalAddress)
	}

	// Create ABI
	parsedAbi, err := abi.JSON(strings.NewReader(bindings.OmniPortalMetaData.ABI))
	if err != nil {
		return OpSender{}, errors.Wrap(err, "parse abi error")
	}

	return OpSender{
		txMgrs:  txMgrs,
		portals: portals,
		abi:     &parsedAbi,
	}, nil
}

func (o OpSender) SendTransaction(ctx context.Context, submission xchain.Submission) error {
	txMgr, ok := o.txMgrs[submission.DestChainID]
	if !ok {
		return errors.New("tx mgr not found", "dest_chain_id", submission.DestChainID)
	}

	// Get some info for logging
	var startOffset uint64
	if len(submission.Msgs) > 0 {
		startOffset = submission.Msgs[0].StreamOffset
	}

	log.Debug(ctx, "Sending submission transaction",
		"dest_chain_id", submission.DestChainID,
		"block_height", submission.BlockHeader.BlockHeight,
		"source_chain_id", submission.BlockHeader.SourceChainID,
		"start_offset", startOffset,
		"msgs", len(submission.Msgs),
	)

	const gasLimit = 1_000_000 // TODO(lazar): make configurable

	to, ok := o.portals[submission.DestChainID]
	if !ok {
		return errors.New("portal not found", "dest_chain_id", submission.DestChainID)
	}

	txData, err := o.GetXSubmitBytes(TranslateSubmission(submission))
	if err != nil {
		return err
	}

	candidate := txmgr.TxCandidate{
		TxData:   txData,
		To:       &to,
		GasLimit: gasLimit,
		Value:    big.NewInt(0), // todo(lazar); is this right?
	}

	rec, err := txMgr.Send(ctx, candidate)
	if err != nil {
		return errors.Wrap(err, "failed to send tx")
	}

	log.Debug(ctx, "Sent submission transaction",
		"dest_chain_id", submission.DestChainID,
		"block_height", submission.BlockHeader.BlockHeight,
		"source_chain_id", submission.BlockHeader.SourceChainID,
		"status", rec.Status,
		"gas_used", rec.GasUsed,
		"tx_hash", rec.TxHash)

	return nil
}

// NewTxMgrConfig - creates a new txmgr config from the given CLI config and private key. This is taken and modified from op
func NewTxMgrConfig(ctx context.Context, cfg txmgr.CLIConfig,
	privateKey *ecdsa.PrivateKey, client *ethclient.Client) (txmgr.Config, error) {
	if err := cfg.Check(); err != nil {
		return txmgr.Config{}, fmt.Errorf("invalid config: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, cfg.NetworkTimeout)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, cfg.NetworkTimeout)
	defer cancel()
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return txmgr.Config{}, errors.Wrap(err, "could not dial fetch L1 chain ID")
	}

	signer := func(chainID *big.Int) opcrypto.SignerFn {
		s := opcrypto.PrivateKeySignerFn(privateKey, chainID)
		return func(_ context.Context, addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return s(addr, tx)
		}
	}

	feeLimitThreshold, err := eth.GweiToWei(cfg.FeeLimitThresholdGwei)
	if err != nil {
		return txmgr.Config{}, errors.Wrap(err, "invalid fee limit threshold")
	}

	minBaseFee, err := eth.GweiToWei(cfg.MinBaseFeeGwei)
	if err != nil {
		return txmgr.Config{}, errors.Wrap(err, "invalid min base fee")
	}

	minTipCap, err := eth.GweiToWei(cfg.MinTipCapGwei)
	if err != nil {
		return txmgr.Config{}, errors.Wrap(err, "invalid min tip cap")
	}

	return txmgr.Config{
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
		Signer:                    signer(chainID),
		From:                      crypto.PubkeyToAddress(privateKey.PublicKey),
	}, nil
}

func initTxMgr(cfg txmgr.Config, logger ethlog.Logger) (txmgr.TxManager, error) {
	// todo(lazar): metrics
	m := metrics.NoopTxMetrics{}
	txMgr, err := txmgr.NewSimpleTxManagerFromConfig("op-relayer", logger, &m, cfg)
	if err != nil {
		return nil, errors.New("failed to create tx mgr", "error", err)
	}

	return txMgr, nil
}

func (o OpSender) GetXSubmitBytes(xsub bindings.XTypesSubmission) ([]byte, error) {
	bytes, err := o.abi.Pack("xsubmit", xsub)
	if err != nil {
		return nil, errors.Wrap(err, "pack xsubmit")
	}

	return bytes, nil
}
