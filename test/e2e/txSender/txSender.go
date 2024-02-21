package txsender

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/txmgr"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TxSender struct {
	txMgr     txmgr.TxManager
	portal    common.Address
	abi       *abi.ABI
	chain     netconf.Chain
	chainName string
}

const (
	gasLimit = 1000000
	interval = 3
)

func NewTxSender(
	ctx context.Context,
	chain netconf.Chain,
	rpcClient *ethclient.Client,
	privateKey *ecdsa.PrivateKey,
	chainName string,
	abi abi.ABI,
) (TxSender, error) {
	// creates our new CLI config for our tx manager
	cliConfig := txmgr.NewCLIConfig(
		chain.RPCURL,
		chain.BlockPeriod/interval, // we want to query receipts at block period / interval
		txmgr.DefaultSenderFlagValues,
	)

	// get the config for our tx manager
	cfg, err := txmgr.NewConfig(ctx, cliConfig, privateKey, rpcClient)
	if err != nil {
		return TxSender{}, err
	}

	// create a simple tx manager from our config
	txMgr, err := txmgr.NewSimpleTxManagerFromConfig(chainName, cfg)
	if err != nil {
		return TxSender{}, errors.New("failed to create tx mgr", "error", err)
	}

	// return our new TxSender
	return TxSender{
		txMgr:     txMgr,
		portal:    common.HexToAddress(chain.PortalAddress),
		abi:       &abi,
		chain:     chain,
		chainName: chainName,
	}, nil
}

func (s TxSender) sendTransaction(
	ctx context.Context,
	destChainID uint64,
	data []byte,
	value *big.Int,
) error {
	if s.txMgr == nil {
		return errors.New("tx mgr not found", "dest_chain_id", destChainID)
	}

	if destChainID == s.chain.ID {
		return errors.New("unexpected destination chain [BUG]",
			"got", destChainID, "expect", s.chain.ID)
	}

	// make it "optional" to pass in the send value
	if value == nil {
		value = big.NewInt(0)
	}

	candidate := txmgr.TxCandidate{
		TxData:   data,
		To:       &s.portal,
		GasLimit: gasLimit,
		Value:    value,
	}

	rec, err := s.txMgr.Send(ctx, candidate)
	if err != nil {
		return errors.Wrap(err, "failed to send tx")
	}

	log.Hex7("Sent tx", rec.TxHash.Bytes())

	return nil
}
