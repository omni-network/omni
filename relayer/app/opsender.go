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
	"github.com/omni-network/omni/lib/txmgr"
	t "github.com/omni-network/omni/lib/txmgr"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// OpSender uses txmgr to send transactions to the destination chain.
type OpSender struct {
	txMgr      txmgr.TxManager
	portal     common.Address
	abi        *abi.ABI
	chain      netconf.Chain
	chainNames map[uint64]string
}

// NewOpSender creates a new sender that uses txmgr to send transactions to the destination chain.
func NewOpSender(ctx context.Context, chain netconf.Chain, rpcClient *ethclient.Client,
	privateKey ecdsa.PrivateKey, chainNames map[uint64]string) (OpSender, error) {
	// we want to query receipts every 1/3 of the block time
	cfg, err := txmgr.NewConfig(ctx, txmgr.NewCLIConfig(
		chain.RPCURL,
		chain.BlockPeriod/3,
		txmgr.DefaultSenderFlagValues,
	),
		&privateKey,
		rpcClient,
	)
	if err != nil {
		return OpSender{}, err
	}

	txMgr, err := initTxMgr(cfg, chain.Name)
	if err != nil {
		return OpSender{}, err
	}

	// Create ABI
	parsedAbi, err := abi.JSON(strings.NewReader(bindings.OmniPortalMetaData.ABI))
	if err != nil {
		return OpSender{}, errors.Wrap(err, "parse abi error")
	}

	return OpSender{
		txMgr:      txMgr,
		portal:     common.HexToAddress(chain.PortalAddress),
		abi:        &parsedAbi,
		chain:      chain,
		chainNames: chainNames,
	}, nil
}

// SendTransaction sends the submission to the destination chain.
func (o OpSender) SendTransaction(ctx context.Context, submission xchain.Submission) error {
	if o.txMgr == nil {
		return errors.New("tx mgr not found", "dest_chain_id", submission.DestChainID)
	} else if submission.DestChainID != o.chain.ID {
		return errors.New("unexpected destination chain [BUG]",
			"got", submission.DestChainID, "expect", o.chain.ID)
	}

	// Get some info for logging
	var startOffset uint64
	if len(submission.Msgs) > 0 {
		startOffset = submission.Msgs[0].StreamOffset
	}

	dstChain := o.chain.Name
	srcChain := o.chainNames[submission.BlockHeader.SourceChainID]

	ctx = log.WithCtx(ctx, "req_id", randomHex7())
	log.Debug(ctx, "Received submission",
		"dest_chain", dstChain,
		"height", submission.BlockHeader.BlockHeight,
		"source_chain", srcChain,
		"start_offset", startOffset,
		"msgs", len(submission.Msgs),
	)

	const gasLimit = 1_000_000 // TODO(lazar): make configurable

	txData, err := o.getXSubmitBytes(t.SubmissionToBinding(submission))
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
		"dest_chain", dstChain,
		"height", submission.BlockHeader.BlockHeight,
		"source_chain", submission.BlockHeader.SourceChainID,
		"status", rec.Status,
		"gas_used", rec.GasUsed,
		"tx_hash", rec.TxHash)

	submissionTotal.WithLabelValues(srcChain, dstChain).Inc()
	msgTotal.WithLabelValues(srcChain, dstChain).Add(float64(len(submission.Msgs)))

	return nil
}

// initTxMgr creates a new txmgr.TxManager from the given config.
func initTxMgr(cfg txmgr.Config, chainName string) (txmgr.TxManager, error) {
	txMgr, err := txmgr.NewSimpleTxManagerFromConfig(chainName, cfg)
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
