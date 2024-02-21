package txsender

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/txmgr"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Sender struct {
	txMgr     txmgr.TxManager
	portal    common.Address
	abi       *abi.ABI
	chainName string
	chainID   uint64
}

const (
	gasLimit = 1000000
	interval = 3
)

func DeployTxSender(
	ctx context.Context,
	rpcCurl string,
	blockPeriod time.Duration,
	chainID uint64,
	portalAddress common.Address,
	rpcClient *ethclient.Client,
	privateKey *ecdsa.PrivateKey,
	chainName string,
	abi abi.ABI,
) (Sender, error) {
	// creates our new CLI config for our tx manager
	cliConfig := txmgr.NewCLIConfig(
		rpcCurl,
		blockPeriod/interval, // we want to query receipts at block period / interval
		txmgr.DefaultSenderFlagValues,
	)

	// get the config for our tx manager
	cfg, err := txmgr.NewConfig(ctx, cliConfig, privateKey, rpcClient)
	if err != nil {
		return Sender{}, errors.New("create tx mgr config", "error", err)
	}

	// create a simple tx manager from our config
	txMgr, err := txmgr.NewSimpleTxManagerFromConfig(chainName, cfg)
	if err != nil {
		return Sender{}, errors.New("create tx mgr", "error", err)
	}

	return Sender{
		txMgr:     txMgr,
		portal:    portalAddress,
		abi:       &abi,
		chainName: chainName,
		chainID:   chainID,
	}, nil
}

func (s Sender) sendTransaction(
	ctx context.Context,
	destChainID uint64,
	data []byte,
	value *big.Int,
) error {
	if s.txMgr == nil {
		return errors.New("tx mgr not found", "dest_chain_id", destChainID)
	}

	if destChainID == s.chainID {
		return errors.New("unexpected destination chain [BUG]",
			"got", destChainID, "expect", s.chainID)
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
