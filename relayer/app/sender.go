package relayer

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"slices"
	"strings"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/txmgr"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Sender uses txmgr to send transactions to the destination chain.
type Sender struct {
	txMgr      txmgr.TxManager
	portal     common.Address
	abi        *abi.ABI
	chain      netconf.Chain
	chainNames map[uint64]string
	rpcClient  ethclient.Client
}

// NewSender creates a new sender that uses txmgr to send transactions to the destination chain.
func NewSender(chain netconf.Chain, rpcClient ethclient.Client,
	privateKey ecdsa.PrivateKey, chainNames map[uint64]string) (Sender, error) {
	// we want to query receipts every 1/3 of the block time
	cfg, err := txmgr.NewConfig(txmgr.NewCLIConfig(
		chain.ID,
		chain.BlockPeriod/3,
		txmgr.DefaultSenderFlagValues,
	),
		&privateKey,
		rpcClient,
	)
	if err != nil {
		return Sender{}, err
	}

	txMgr, err := txmgr.NewSimple(chain.Name, cfg)
	if err != nil {
		return Sender{}, errors.Wrap(err, "create tx mgr")
	}

	// Create ABI
	parsedAbi, err := abi.JSON(strings.NewReader(bindings.OmniPortalMetaData.ABI))
	if err != nil {
		return Sender{}, errors.Wrap(err, "parse abi error")
	}

	return Sender{
		txMgr:      txMgr,
		portal:     chain.PortalAddress,
		abi:        &parsedAbi,
		chain:      chain,
		chainNames: chainNames,
		rpcClient:  rpcClient,
	}, nil
}

// SendTransaction sends the submission to the destination chain.
func (o Sender) SendTransaction(ctx context.Context, submission xchain.Submission) error {
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

	// Request attributes added to context (for downstream logging) and manually added to errors (for upstream logging).
	reqAttrs := []any{
		"req_id", randomHex7(),
		"src_chain", srcChain,
	}

	ctx = log.WithCtx(ctx, reqAttrs...)
	log.Debug(ctx, "Received submission",
		"start_offset", startOffset,
		"msgs", len(submission.Msgs),
	)

	txData, err := o.getXSubmitBytes(submissionToBinding(submission))
	if err != nil {
		return err
	}

	candidate := txmgr.TxCandidate{
		TxData: txData,
		To:     &o.portal,
		Value:  big.NewInt(0),
	}

	tx, rec, err := o.txMgr.Send(ctx, candidate)
	if err != nil {
		return errors.Wrap(err, "failed to send tx", reqAttrs...)
	}

	submissionTotal.WithLabelValues(srcChain, dstChain).Inc()
	msgTotal.WithLabelValues(srcChain, dstChain).Add(float64(len(submission.Msgs)))

	receiptAttrs := []any{
		"valset_id", submission.ValidatorSetID,
		"status", rec.Status,
		"nonce", tx.Nonce(),
		"gas_used", rec.GasUsed,
		"tx_hash", rec.TxHash,
	}

	if rec.Status == 0 {
		// Try and get debug information of the reverted transaction
		resp, err := o.rpcClient.CallContract(ctx, ethereum.CallMsg{
			From:          o.txMgr.From(),
			To:            tx.To(),
			Gas:           tx.Gas(),
			GasPrice:      tx.GasPrice(),
			GasFeeCap:     tx.GasFeeCap(),
			GasTipCap:     tx.GasTipCap(),
			Value:         tx.Value(),
			Data:          tx.Data(),
			AccessList:    tx.AccessList(),
			BlobGasFeeCap: tx.BlobGasFeeCap(),
			BlobHashes:    tx.BlobHashes(),
		}, rec.BlockNumber)

		errAttrs := slices.Concat(receiptAttrs, reqAttrs, []any{
			"call_resp", hexutil.Encode(resp),
			"call_err", err,
		})

		revertedSubmissionTotal.WithLabelValues(srcChain, dstChain).Inc()

		return errors.New("submission reverted", errAttrs...)
	}

	log.Info(ctx, "Sent submission", receiptAttrs...)

	return nil
}

// getXSubmitBytes returns the byte representation of the xsubmit function call.
func (o Sender) getXSubmitBytes(sub bindings.XSubmission) ([]byte, error) {
	bytes, err := o.abi.Pack("xsubmit", sub)
	if err != nil {
		return nil, errors.Wrap(err, "pack xsubmit")
	}

	return bytes, nil
}
