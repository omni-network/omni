package relayer

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/omni-network/omni/lib/xchain"
)

type Portal struct {
	Session bindings.OmniPortalSession
}

var _ Sender = (*SenderService)(nil)

type SenderService struct {
	Portal     map[uint64]*Portal
	PrivateKey *ecdsa.PrivateKey
	RpcClient  *ethclient.Client
}

func NewPortal(portalCfg PortalConfig, rpcClient *ethclient.Client, privateKey ecdsa.PrivateKey) (Portal, error) {
	contract, err := bindings.NewOmniPortal(common.HexToAddress(portalCfg.contractAddress), rpcClient)
	if err != nil {
		return Portal{}, err
	}

	transactor, err := bind.NewKeyedTransactorWithChainID(&privateKey, big.NewInt(int64(portalCfg.chainID)))
	if err != nil {
		return Portal{}, err
	}

	session := bindings.OmniPortalSession{
		Contract:     contract,
		TransactOpts: *transactor,
		CallOpts: bind.CallOpts{
			From: crypto.PubkeyToAddress(privateKey.PublicKey),
		},
	}

	return Portal{
		Session: session,
	}, nil
}

func (s SenderService) SendTransaction(ctx context.Context, submission xchain.Submission) error {
	xChainSubmission := translateSubmission(submission)

	// todo(Lazar): add destChainID as top level property in xchain.Submission
	if len(submission.Msgs) == 0 {
		return nil
	}

	destChainID := submission.Msgs[0].DestChainID
	portal, ok := s.Portal[destChainID]
	if !ok {
		return errors.New("portal not found", "destChainID", destChainID)
	}

	tx, err := portal.Session.Xsubmit(xChainSubmission)
	if err != nil {
		// todo(Lazar): handle error
		return err
	}

	log.Info(ctx, "submitted_tx",
		"txHash", tx.Hash().Hex(),
		"nonce", tx.Nonce(),
		"gas_price", tx.GasPrice(),
	)

	waitCtx, cancel := context.WithTimeout(ctx, 1*time.Minute)
	receipt, err := bind.WaitMined(waitCtx, s.RpcClient, tx)
	cancel()
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			// todo(Lazar): handle error increase gas price and retry
		}

		return errors.Wrap(err, "submission tx not mined (tx=%s): %w", tx.Hash().Hex())
	}

	if receipt.Status == ethtypes.ReceiptStatusFailed {
		return errors.New("submission tx failed (tx=%s)", tx.Hash().Hex())
	}

	// todo(Lazar): handle success case, metrics and cache

	return nil
}

func translateSubmission(submission xchain.Submission) bindings.XChainSubmission {
	chainSubmission := bindings.XChainSubmission{
		AttestationRoot: submission.AttestationRoot,
		BlockHeader: bindings.XChainBlockHeader{
			SourceChainId: submission.BlockHeader.SourceChainID,
			BlockHeight:   submission.BlockHeader.BlockHeight,
			BlockHash:     submission.BlockHeader.BlockHash,
		},
		Proof:      submission.Proof,
		ProofFlags: submission.ProofFlags,
	}

	signatures := make([]bindings.XChainSigTuple, len(submission.Signatures))
	for i, sig := range submission.Signatures {
		signatures[i] = bindings.XChainSigTuple{
			ValidatorPubKey: sig.ValidatorPubKey[:],
			Signature:       sig.Signature[:],
		}
	}

	chainSubmission.Signatures = signatures

	msgs := make([]bindings.XChainMsg, len(submission.Msgs))
	for i, msg := range submission.Msgs {
		msgs[i] = bindings.XChainMsg{
			SourceChainId: msg.SourceChainID,
			DestChainId:   msg.DestChainID,
			StreamOffset:  msg.StreamOffset,
			Sender:        common.BytesToAddress(msg.SourceMsgSender[:]),
			To:            common.BytesToAddress(msg.DestAddress[:]),
			Data:          msg.Data,
			GasLimit:      msg.DestGasLimit,
			TxHash:        msg.TxHash,
		}
	}

	chainSubmission.Msgs = msgs

	return chainSubmission
}
