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
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/txmgr"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// Sender uses txmgr to send transactions a specific destination chain.
type onSubmitFunc func(context.Context, *ethtypes.Transaction, *ethclient.Receipt, xchain.Submission)

// Sender uses txmgr to send transactions to a specific destination chain.
type Sender struct {
	network      netconf.ID
	txMgr        txmgr.TxManager
	gasEstimator gasEstimator
	abi          *abi.ABI
	chain        netconf.Chain
	gasToken     tokens.Token
	chainNames   map[xchain.ChainVersion]string
	ethCl        ethclient.Client
	onSubmit     onSubmitFunc
}

// NewSender returns a new sender.
func NewSender(
	network netconf.ID,
	chain netconf.Chain,
	rpcClient ethclient.Client,
	privateKey ecdsa.PrivateKey,
	chainNames map[xchain.ChainVersion]string,
	onSubmit onSubmitFunc,
) (Sender, error) {
	const receiptPollFreq = 3 // Query receipts every 1/3 of the block time
	cfg, err := txmgr.NewConfig(
		txmgr.NewCLIConfig(
			chain.ID,
			chain.BlockPeriod/receiptPollFreq,
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

	meta, ok := evmchain.MetadataByID(chain.ID)
	if !ok {
		return Sender{}, errors.New("chain metadata not found", "chain_id", chain.ID)
	}

	return Sender{
		network:      network,
		txMgr:        txMgr,
		gasEstimator: newGasEstimator(network),
		abi:          &parsedAbi,
		chain:        chain,
		gasToken:     meta.NativeToken,
		chainNames:   chainNames,
		ethCl:        rpcClient,
		onSubmit:     onSubmit,
	}, nil
}

// SendAsync sends the submission to the destination chain asynchronously.
// It returns a channel that will receive an error if the submission fails or nil when it succeeds.
// Nonces are however reserved synchronously, so ordering of submissions
// is preserved.
func (s Sender) SendAsync(ctx context.Context, sub xchain.Submission) <-chan error {
	// Helper function to return error "synchronously".
	returnErr := func(err error) chan error {
		resp := make(chan error, 1)
		resp <- err

		return resp
	}

	if s.txMgr == nil {
		return returnErr(errors.New("tx mgr not found [BUG]", "dest_chain_id", sub.DestChainID))
	} else if sub.DestChainID != s.chain.ID {
		return returnErr(errors.New("unexpected destination chain [BUG]",
			"got", sub.DestChainID, "expect", s.chain.ID))
	}

	// Get some info for logging
	var startOffset uint64
	if len(sub.Msgs) > 0 {
		startOffset = sub.Msgs[0].StreamOffset
	}

	dstChain := s.chain.Name
	srcChain := s.chainNames[sub.AttHeader.ChainVersion]

	// Request attributes added to context (for downstream logging) and manually added to errors (for upstream logging).
	reqAttrs := []any{
		"req_id", randomHex7(),
		"src_chain", srcChain,
	}

	ctx = log.WithCtx(ctx, reqAttrs...)
	log.Debug(ctx, "Received submission",
		"attest_offset", sub.AttHeader.AttestOffset,
		"start_msg_offset", startOffset,
		"msgs", len(sub.Msgs),
	)

	txData, err := xchain.EncodeXSubmit(xchain.SubmissionToBinding(sub))
	if err != nil {
		return returnErr(err)
	}

	// Reserve a nonce here to ensure correctly ordered submissions.
	nonce, err := s.txMgr.ReserveNextNonce(ctx)
	if err != nil {
		return returnErr(err)
	}

	estimatedGas := s.gasEstimator(s.chain.ID, sub.Msgs)

	candidate := txmgr.TxCandidate{
		TxData:   txData,
		To:       &s.chain.PortalAddress,
		GasLimit: estimatedGas,
		Value:    big.NewInt(0),
		Nonce:    &nonce,
	}

	asyncResp := make(chan error, 1) // Actual async response populated by goroutine below.
	go func() {
		tx, rec, err := s.txMgr.Send(ctx, candidate)
		if err != nil {
			asyncResp <- errors.Wrap(err, "failed to send tx", reqAttrs...)
			return
		}

		submissionTotal.WithLabelValues(srcChain, dstChain).Inc()
		msgTotal.WithLabelValues(srcChain, dstChain).Add(float64(len(sub.Msgs)))
		gasEstimated.WithLabelValues(dstChain).Observe(float64(estimatedGas))

		receiptAttrs := []any{
			"valset_id", sub.ValidatorSetID,
			"status", rec.Status,
			"nonce", tx.Nonce(),
			"height", rec.BlockNumber.Uint64(),
			"gas_used", rec.GasUsed,
			"tx_hash", rec.TxHash,
		}

		if s.onSubmit != nil {
			go s.onSubmit(ctx, tx, rec, sub)
		}

		const statusReverted = 0
		if rec.Status == statusReverted {
			// Try and get debug information of the reverted transaction
			resp, err := s.ethCl.CallContract(ctx, callFromTx(s.txMgr.From(), tx), rec.BlockNumber)

			errAttrs := slices.Concat(receiptAttrs, reqAttrs, []any{
				"call_resp", hexutil.Encode(resp),
				"call_err", err,
				"gas_limit", tx.Gas(),
			})

			revertedSubmissionTotal.WithLabelValues(srcChain, dstChain).Inc()

			asyncResp <- errors.New("submission reverted", errAttrs...)

			return
		}

		log.Info(ctx, "Sent submission", receiptAttrs...)
		asyncResp <- nil
	}()

	return asyncResp
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
