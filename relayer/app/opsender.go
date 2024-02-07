package relayer

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"strings"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
	"github.com/omni-network/omni/relayer/txmgr"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// OpSender uses txmgr to send transactions to the destination chain.
type OpSender struct {
	txMgr  txmgr.TxManager
	portal common.Address
	abi    *abi.ABI
}

// NewOpSender creates a new sender that uses txmgr to send transactions to the destination chain.
func NewOpSender(ctx context.Context, chain netconf.Chain, rpcClient *ethclient.Client,
	privateKey ecdsa.PrivateKey) (OpSender, error) {

	cfg, err := txmgr.NewConfig(ctx, txmgr.NewCLIConfig(chain.RPCURL, txmgr.DefaultSenderFlagValues),
		&privateKey, rpcClient)
	if err != nil {
		return OpSender{}, err
	}

	txMgr, err := initTxMgr(cfg)
	if err != nil {
		return OpSender{}, err
	}

	// Create ABI
	parsedAbi, err := abi.JSON(strings.NewReader(bindings.OmniPortalMetaData.ABI))
	if err != nil {
		return OpSender{}, errors.Wrap(err, "parse abi error")
	}

	return OpSender{
		txMgr:  txMgr,
		portal: common.HexToAddress(chain.PortalAddress),
		abi:    &parsedAbi,
	}, nil
}

// SendTransaction sends the submission to the destination chain.
func (o OpSender) SendTransaction(ctx context.Context, submission xchain.Submission) error {
	if o.txMgr == nil {
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

	txData, err := o.getXSubmitBytes(SubmissionToBinding(submission))
	if err != nil {
		return err
	}

	candidate := txmgr.TxCandidate{
		TxData:   txData,
		To:       &o.portal,
		GasLimit: gasLimit,
		Value:    big.NewInt(0),
	}

	rec, err := o.txMgr.Send(ctx, candidate)
	if err != nil {
		return errors.Wrap(err, "failed to send tx")
	}

	log.Info(ctx, "Sent submission transaction",
		"dest_chain_id", submission.DestChainID,
		"block_height", submission.BlockHeader.BlockHeight,
		"source_chain_id", submission.BlockHeader.SourceChainID,
		"status", rec.Status,
		"gas_used", rec.GasUsed,
		"tx_hash", rec.TxHash)

	return nil
}

// initTxMgr creates a new txmgr.TxManager from the given config.
func initTxMgr(cfg txmgr.Config) (txmgr.TxManager, error) {
	txMgr, err := txmgr.NewSimpleTxManagerFromConfig("op-relayer", cfg)
	if err != nil {
		return nil, errors.New("failed to create tx mgr", "error", err)
	}

	return txMgr, nil
}

// getXSubmitBytes returns the byte representation of the xsubmit function call.
func (o OpSender) getXSubmitBytes(sub bindings.XTypesSubmission) ([]byte, error) {
	bytes, err := o.abi.Pack("xsubmit", sub)
	if err != nil {
		return nil, errors.Wrap(err, "pack xsubmit")
	}

	return bytes, nil
}
