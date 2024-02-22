package xtx

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/txmgr"
	"github.com/omni-network/omni/test/e2e/netman"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Sender struct {
	txMgr     txmgr.TxManager
	portal    common.Address
	abi       *abi.ABI
	chainName string
	chainID   uint64
}

const (
	gasLimit = 1_000_000
	interval = 3
)

func NewTxSender(ctx context.Context, portal netman.Portal, privateKey *ecdsa.PrivateKey, abi abi.ABI) (Sender, error) {
	// creates our new CLI config for our tx manager
	cliConfig := txmgr.NewCLIConfig(
		portal.RPCURL,
		portal.Chain.BlockPeriod/interval, // we want to query receipts at block period / interval
		txmgr.DefaultSenderFlagValues,
	)

	// get the config for our tx manager
	cfg, err := txmgr.NewConfig(ctx, cliConfig, privateKey, portal.Client)
	if err != nil {
		return Sender{}, errors.New("create tx mgr config", "error", err)
	}

	// create a simple tx manager from our config
	txMgr, err := txmgr.NewSimpleTxManagerFromConfig(portal.Chain.Name, cfg)
	if err != nil {
		return Sender{}, errors.New("create tx mgr", "error", err)
	}

	return Sender{
		txMgr:     txMgr,
		portal:    portal.DeployInfo.PortalAddress,
		abi:       &abi,
		chainName: portal.Chain.Name,
		chainID:   portal.Chain.ID,
	}, nil
}

func (s Sender) sendTransaction(
	ctx context.Context,
	destChainID uint64,
	data []byte,
	value *big.Int,
) (*types.Receipt, error) {
	if s.txMgr == nil {
		return nil, errors.New("tx mgr not found", "dest_chain_id", destChainID)
	}

	if destChainID == s.chainID {
		return nil, errors.New("unexpected destination chain [BUG]",
			"got", destChainID, "expect", s.chainID)
	}

	// make it "optional" to pass in the send value
	if value == nil {
		value = big.NewInt(0)
	}

	candidate := txmgr.TxCandidate{
		TxData:   data,
		To:       &s.portal,
		GasLimit: uint64(gasLimit),
		Value:    value,
	}

	rec, err := s.txMgr.Send(ctx, candidate)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send tx")
	}

	log.Info(ctx, "Sent tx", "status", rec.Status, "gas_used", rec.GasUsed, "tx_hash", rec.TxHash.String())

	log.Hex7("Sent tx", rec.TxHash.Bytes())

	return rec, nil
}
