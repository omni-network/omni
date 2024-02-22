package send

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"strings"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/test/e2e/netman"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Manager struct {
	txSender map[uint64]Sender
}

// XCallsArgs represents the arguments for the xcall function.
type XCallArgs struct {
	DestChainID uint64
	Address     common.Address
	Data        []byte
	GasLimit    uint64
}

func New(ctx context.Context, portals map[uint64]netman.Portal, privateKey *ecdsa.PrivateKey) (Manager, error) {
	manager := Manager{
		txSender: make(map[uint64]Sender),
	}

	// deploy tx sender for each portal
	for _, portal := range portals {
		err := manager.makeInstance(ctx, portal, privateKey)
		if err != nil {
			return Manager{}, errors.Wrap(err, "deploy tx sender")
		}
	}

	return manager, nil
}

func (s Manager) makeInstance(ctx context.Context, portal netman.Portal, privateKey *ecdsa.PrivateKey) error {
	chain := portal.Chain

	if _, ok := s.txSender[chain.ID]; ok {
		return errors.New("tx sender already exists", "chain", chain.ID)
	}

	parsedAbi, err := abi.JSON(strings.NewReader(bindings.OmniPortalMetaData.ABI))
	if err != nil {
		return errors.Wrap(err, "parse abi error")
	}

	txSender, err := newTxSender(
		ctx,
		portal,
		privateKey,
		parsedAbi,
	)
	if err != nil {
		return errors.Wrap(err, "deploy tx sender", "chain", chain.ID)
	}

	s.txSender[chain.ID] = txSender

	return nil
}

func (s Manager) SendXCall(ctx context.Context, args XCallArgs, value *big.Int, sourceChainID uint64, address common.Address) (*types.Receipt, error) {
	txSender := s.txSender[sourceChainID]
	bytes, err := xCallBytes(ctx, txSender, args.DestChainID, args.Address, args.Data, args.GasLimit)
	if err != nil {
		return nil, errors.Wrap(err, "get xsubmit bytes")
	}

	receipt, err := txSender.sendTransaction(ctx, address, bytes, value)
	if err != nil {
		return nil, errors.Wrap(err, "send xsubmit")
	}

	return receipt, nil
}

// getXCallBytes returns the byte representation of the xcall function call.
func xCallBytes(ctx context.Context, txSender Sender, destChainID uint64, to common.Address, data []byte, gasLimit uint64) ([]byte, error) {
	log.Info(ctx, "Packing xcall", "destChainID", destChainID, "to", to.String(), "data", data, "gasLimit", gasLimit)
	bytes, err := txSender.abi.Pack("xcall", destChainID, to, data) // TODO: add gas limit?
	if err != nil {
		return nil, errors.Wrap(err, "pack xcall")
	}

	return bytes, nil
}
