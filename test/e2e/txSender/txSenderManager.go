package txsender

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"strings"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TxSenderManager struct {
	txSender map[uint64]TxSender
	abi      *abi.ABI
}

func NewTxSenderManager() (TxSenderManager, error) {
	// create ABI interface
	parsedAbi, err := abi.JSON(strings.NewReader(bindings.OmniPortalMetaData.ABI))
	if err != nil {
		return TxSenderManager{}, errors.Wrap(err, "parse abi error")
	}

	return TxSenderManager{
		txSender: make(map[uint64]TxSender),
		abi:      &parsedAbi,
	}, nil
}

func (s TxSenderManager) CreateNewTxSender(ctx context.Context,
	chain netconf.Chain,
	rpcClient *ethclient.Client,
	privateKey *ecdsa.PrivateKey,
	chainName string,
) error {
	txSender, err := NewTxSender(ctx, chain, rpcClient, privateKey, chainName, *s.abi)
	if err != nil {
		return err
	}

	s.txSender[chain.ID] = txSender

	return nil
}

func (s TxSenderManager) SendXCallTransaction(
	ctx context.Context,
	msg xchain.Msg,
	value *big.Int,
	sourceChainID uint64,
) error {
	txSender := s.txSender[sourceChainID]
	bytes, err := s.getXCallBytes(MsgToBindings(msg))
	if err != nil {
		return errors.Wrap(err, "get xsubmit bytes")
	}

	return txSender.sendTransaction(ctx, msg.DestChainID, bytes, value)
}

// getXCallBytes returns the byte representation of the xcall function call.
func (s TxSenderManager) getXCallBytes(sub bindings.XTypesMsg) ([]byte, error) {
	bytes, err := s.abi.Pack("xcall", sub)
	if err != nil {
		return nil, errors.Wrap(err, "pack xcall")
	}

	return bytes, nil
}
