package relayer

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var _ Sender = (*SimpleSender)(nil)

type SimpleSender struct {
	Clients  map[uint64]*ethclient.Client
	Sessions map[uint64]bindings.OmniPortalSession
}

// NewSenderService creates a new sender service.
func NewSenderService(chains []netconf.Chain, rpcClientPerChain map[uint64]*ethclient.Client,
	privateKey ecdsa.PrivateKey,
) (SimpleSender, error) {
	sessions := make(map[uint64]bindings.OmniPortalSession)
	for _, chain := range chains {
		session, err := NewSession(chain, rpcClientPerChain[chain.ID], privateKey)
		if err != nil {
			return SimpleSender{}, err
		}

		sessions[chain.ID] = session
	}

	return SimpleSender{
		Sessions: sessions,
		Clients:  rpcClientPerChain,
	}, nil
}

func NewSession(chain netconf.Chain, rpcClient *ethclient.Client,
	privateKey ecdsa.PrivateKey) (bindings.OmniPortalSession, error) {
	contract, err := bindings.NewOmniPortal(common.HexToAddress(chain.PortalAddress), rpcClient)
	if err != nil {
		return bindings.OmniPortalSession{}, err
	}

	transactor, err := bind.NewKeyedTransactorWithChainID(&privateKey, big.NewInt(int64(chain.ID)))
	if err != nil {
		return bindings.OmniPortalSession{}, errors.Wrap(err, "new transactor")
	}

	session := bindings.OmniPortalSession{
		Contract:     contract,
		TransactOpts: *transactor,
		CallOpts: bind.CallOpts{
			From: crypto.PubkeyToAddress(privateKey.PublicKey),
		},
	}

	return session, nil
}

func (s SimpleSender) SendTransaction(ctx context.Context, submission xchain.Submission) error {
	xChainSubmission := TranslateSubmission(submission)

	session, ok := s.Sessions[submission.DestChainID]
	if !ok {
		return errors.New("session not found", "destChainID", submission.DestChainID)
	}

	tx, err := session.Xsubmit(xChainSubmission)
	if err != nil {
		// todo(Lazar): handle error
		return err
	}

	log.Info(ctx, "Submitted_tx",
		"tx_hash", tx.Hash().Hex(),
		"nonce", tx.Nonce(),
		"gas_price", tx.GasPrice(),
	)

	// todo(Lazar): handle error
	if ctx.Err() != nil {
		// shutdown
		return errors.Wrap(ctx.Err(), "ctx error")
	}

	// todo(Lazar): handle success case, metrics and cache

	return nil
}
