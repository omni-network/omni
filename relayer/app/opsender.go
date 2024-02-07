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
	txMgrs  map[uint64]txmgr.TxManager
	portals map[uint64]common.Address
	abi     *abi.ABI
}

// NewOpSender creates a new sender that uses txmgr to send transactions to the destination chain.
// TODO(corver): Change to single destination chain.
func NewOpSender(ctx context.Context, chains []netconf.Chain, rpcClientPerChain map[uint64]*ethclient.Client,
	privateKey ecdsa.PrivateKey) (OpSender, error) {
	txMgrs := make(map[uint64]txmgr.TxManager)
	portals := make(map[uint64]common.Address)

	for _, chain := range chains {
		cfg, err := txmgr.NewConfig(ctx, txmgr.NewCLIConfig(chain.RPCURL, txmgr.DefaultSenderFlagValues),
			&privateKey, rpcClientPerChain[chain.ID])
		if err != nil {
			return OpSender{}, err
		}

		txMgr, err := initTxMgr(cfg)
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

// SendTransaction sends the submission to the destination chain.
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

	txData, err := o.getXSubmitBytes(SubmissionToBinding(submission))
	if err != nil {
		return err
	}

	candidate := txmgr.TxCandidate{
		TxData:   txData,
		To:       &to,
		GasLimit: gasLimit,
		Value:    big.NewInt(0),
	}

	rec, err := txMgr.Send(ctx, candidate)
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
