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
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// Sender uses txmgr to send transactions to the destination chain.
type Sender struct {
	network    netconf.ID
	txMgr      txmgr.TxManager
	portal     common.Address
	abi        *abi.ABI
	chain      netconf.Chain
	chainNames map[xchain.ChainVersion]string
	rpcClient  ethclient.Client
}

// NewSender creates a new sender that uses txmgr to send transactions to the destination chain.
func NewSender(
	network netconf.ID,
	chain netconf.Chain,
	rpcClient ethclient.Client,
	privateKey ecdsa.PrivateKey,
	chainNames map[xchain.ChainVersion]string,
) (Sender, error) {
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
		network:    network,
		txMgr:      txMgr,
		portal:     chain.PortalAddress,
		abi:        &parsedAbi,
		chain:      chain,
		chainNames: chainNames,
		rpcClient:  rpcClient,
	}, nil
}

// SendTransaction sends the submission to the destination chain.
func (o Sender) SendTransaction(ctx context.Context, sub xchain.Submission) error {
	if o.txMgr == nil {
		return errors.New("tx mgr not found", "dest_chain_id", sub.DestChainID)
	} else if sub.DestChainID != o.chain.ID {
		return errors.New("unexpected destination chain [BUG]",
			"got", sub.DestChainID, "expect", o.chain.ID)
	}

	// Get some info for logging
	var startOffset uint64
	if len(sub.Msgs) > 0 {
		startOffset = sub.Msgs[0].StreamOffset
	}

	dstChain := o.chain.Name
	srcChain := o.chainNames[sub.BlockHeader.ChainVersion()]

	// Request attributes added to context (for downstream logging) and manually added to errors (for upstream logging).
	reqAttrs := []any{
		"req_id", randomHex7(),
		"src_chain", srcChain,
	}

	ctx = log.WithCtx(ctx, reqAttrs...)
	log.Debug(ctx, "Received submission",
		"block_offset", sub.BlockHeader.BlockOffset,
		"start_msg_offset", startOffset,
		"msgs", len(sub.Msgs),
	)

	txData, err := o.getXSubmitBytes(submissionToBinding(sub))
	if err != nil {
		return err
	}

	// Reserve a nonce here to ensure correctly ordered submissions.
	nonce, err := o.txMgr.ReserveNextNonce(ctx)
	if err != nil {
		return err
	}

	estimatedGas := estimateGas(sub.Msgs)
	if sub.BlockHeader.SourceChainID == o.network.Static().OmniConsensusChainIDUint64() {
		estimatedGas = consensusGasLimit(o.network)
	}

	candidate := txmgr.TxCandidate{
		TxData:   txData,
		To:       &o.portal,
		GasLimit: estimatedGas,
		Value:    big.NewInt(0),
		Nonce:    &nonce,
	}

	tx, rec, err := o.txMgr.Send(ctx, candidate)
	if err != nil {
		return errors.Wrap(err, "failed to send tx", reqAttrs...)
	}

	submissionTotal.WithLabelValues(srcChain, dstChain).Inc()
	msgTotal.WithLabelValues(srcChain, dstChain).Add(float64(len(sub.Msgs)))
	gasEstimated.WithLabelValues(dstChain).Observe(float64(estimatedGas))

	receiptAttrs := []any{
		"valset_id", sub.ValidatorSetID,
		"status", rec.Status,
		"nonce", tx.Nonce(),
		"gas_used", rec.GasUsed,
		"tx_hash", rec.TxHash,
	}

	if rec.Status == 0 {
		// Try and get debug information of the reverted transaction
		resp, err := o.rpcClient.CallContract(ctx, callFromTx(o.txMgr.From(), tx), rec.BlockNumber)

		errAttrs := slices.Concat(receiptAttrs, reqAttrs, []any{
			"call_resp", hexutil.Encode(resp),
			"call_err", err,
			"gas_limit", estimatedGas,
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

func callFromTx(from common.Address, tx *ethtypes.Transaction) ethereum.CallMsg {
	resp := ethereum.CallMsg{
		From:          from,
		To:            tx.To(),
		Gas:           tx.Gas(),
		Value:         tx.Value(),
		Data:          tx.Data(),
		AccessList:    tx.AccessList(),
		BlobGasFeeCap: tx.BlobGasFeeCap(),
		BlobHashes:    tx.BlobHashes(),
	}

	// Either populate gas price or gas caps (not both).
	if tx.GasPrice() != nil && tx.GasPrice().Sign() != 0 {
		resp.GasPrice = tx.GasPrice()
	} else {
		resp.GasFeeCap = tx.GasFeeCap()
		resp.GasTipCap = tx.GasTipCap()
	}

	return resp
}
