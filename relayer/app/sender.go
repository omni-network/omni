package relayer

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	miningTimeout = 1 * time.Minute
)

type Portal struct {
	Session   bindings.OmniPortalSession
	RPCClient *ethclient.Client
}

var _ Sender = (*SenderService)(nil)

type SenderService struct {
	Portal map[uint64]Portal
}

// NewSenderService creates a new sender service.
func NewSenderService(chains []netconf.Chain, privateKey ecdsa.PrivateKey) (SenderService, error) {
	portal := make(map[uint64]Portal)
	for _, chain := range chains {
		rpcClient, err := ethclient.Dial(chain.RPCURL)
		if err != nil {
			return SenderService{}, errors.Wrap(err, "dial rpc", "url", chain.RPCURL)
		}

		p, err := NewPortal(chain, rpcClient, privateKey)
		if err != nil {
			return SenderService{}, err
		}

		portal[chain.ID] = p
	}

	return SenderService{
		Portal: portal,
	}, nil
}

func NewPortal(chain netconf.Chain, rpcClient *ethclient.Client, privateKey ecdsa.PrivateKey) (Portal, error) {
	contract, err := bindings.NewOmniPortal(common.HexToAddress(chain.PortalAddress), rpcClient)
	if err != nil {
		return Portal{}, err
	}

	transactor, err := bind.NewKeyedTransactorWithChainID(&privateKey, big.NewInt(int64(chain.ID)))
	if err != nil {
		return Portal{}, errors.Wrap(err, "new transactor")
	}

	session := bindings.OmniPortalSession{
		Contract:     contract,
		TransactOpts: *transactor,
		CallOpts: bind.CallOpts{
			From: crypto.PubkeyToAddress(privateKey.PublicKey),
		},
	}

	return Portal{
		Session:   session,
		RPCClient: rpcClient,
	}, nil
}

func (s SenderService) SendTransaction(ctx context.Context, submission xchain.Submission) error {
	xChainSubmission := TranslateSubmission(submission)
	destChainID := submission.DestChainID()

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
		"tx_hash", tx.Hash().Hex(),
		"nonce", tx.Nonce(),
		"gas_price", tx.GasPrice(),
	)

	waitCtx, cancel := context.WithTimeout(ctx, miningTimeout)
	receipt, err := bind.WaitMined(waitCtx, portal.RPCClient, tx)
	defer cancel()

	if ctx.Err() != nil {
		// shutdown
		return errors.Wrap(ctx.Err(), "ctx error")
	} else if waitCtx.Err() != nil {
		// todo(Lazar): handle error increase gas price and retry
		return errors.Wrap(waitCtx.Err(), "wait mined")
	} else if err != nil {
		return errors.Wrap(err, "wait mined")
	} else if receipt.Status == ethtypes.ReceiptStatusFailed {
		return errors.New("submission tx failed", tx.Hash().Hex())
	}

	// todo(Lazar): handle success case, metrics and cache

	return nil
}
