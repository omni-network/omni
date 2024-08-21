// Package routerecon defines functions to reconcile routescan cross-transactions verifying the data
// matches expected on-chain values.
package routerecon

import (
	"context"
	"strings"
	"time"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func ReconForever(ctx context.Context, network netconf.Network, xprov xchain.Provider, ethCls map[uint64]ethclient.Client) {
	if network.ID != netconf.Omega {
		return
	}

	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			crossTx, err := paginateLatestCrossTx(ctx)
			if err != nil {
				reconFailure.Inc()
				log.Warn(ctx, "RouteRecon failed fetching latest cross transaction (will retry)", err)

				continue
			}

			if err := reconOnce(ctx, network, xprov, ethCls, crossTx); err != nil {
				reconFailure.Inc()
				log.Warn(ctx, "RouteRecon failed (will retry)", err, "id", crossTx.ID)

				continue
			}

			reconSuccess.Inc()
			log.Info(ctx, "RouteRecon success", "id", crossTx.ID)
		}
	}
}

func reconOnce(
	ctx context.Context,
	network netconf.Network,
	xprov xchain.Provider,
	ethCls map[uint64]ethclient.Client,
	crossTx crossTxJSON,
) error {
	relayerAddr, ok := eoa.Address(network.ID, eoa.RoleRelayer)
	if !ok {
		return errors.New("relayer address not found")
	}

	src, err := fetchSource(ctx, network, xprov, ethCls, crossTx.SrcChainID, crossTx.SrcTxHash, crossTx.SrcBlockNumber)
	if err != nil {
		return errors.Wrap(err, "fetch source")
	}

	dst, err := fetchSource(ctx, network, xprov, ethCls, crossTx.DstChainID, crossTx.DstTxHash, crossTx.DstBlockNumber)
	if err != nil {
		return errors.Wrap(err, "fetch destination")
	}

	msgID, err := crossTx.MsgID()
	if err != nil {
		return errors.Wrap(err, "extract msg id")
	}

	msg, err := msgByID(src.XBlock.Msgs, msgID)
	if err != nil {
		return err
	}

	receipt, err := receiptByID(dst.XBlock.Receipts, msgID)
	if err != nil {
		return err
	} else if receipt.RelayerAddress != relayerAddr {
		return errors.New("receipt relayer mismatch", "got", receipt.RelayerAddress, "want", relayerAddr)
	}

	if err := verifySrc(crossTx, src, msg); err != nil {
		return errors.Wrap(err, "verify source")
	}

	if err := verifyDst(crossTx, dst, receipt, relayerAddr); err != nil {
		return errors.Wrap(err, "verify destination")
	}

	if crossTx.SrcTimestamp.After(crossTx.DstTimestamp) {
		return errors.New("source timestamp after destination", "src", crossTx.SrcTimestamp, "dst", crossTx.DstTimestamp)
	}

	return nil
}

func verifyDst(crossTx crossTxJSON, dst source, receipt xchain.Receipt, relayerAddr common.Address) error {
	if common.HexToHash(crossTx.DstBlockHash) != dst.XBlock.BlockHash {
		return errors.New("block hash mismatch", "got", crossTx.DstBlockHash, "want", dst.XBlock.BlockHash)
	}

	if common.HexToHash(crossTx.DstTxHash) != receipt.TxHash {
		return errors.New("tx hash mismatch", "got", crossTx.SrcTxHash, "want", receipt.TxHash)
	}

	if !crossTx.DstTimestamp.Equal(dst.XBlock.Timestamp) {
		return errors.New("timestamp mismatch", "got", crossTx.DstTimestamp, "want", dst.XBlock.Timestamp)
	}

	if common.HexToAddress(crossTx.Data.Relayer) != dst.Sender {
		return errors.New("relayer not destination tx sender", "relayer", crossTx.From, "sender", dst.Sender)
	}

	if common.HexToAddress(crossTx.Data.Relayer) != relayerAddr {
		return errors.New("relayer mismatch", "got", crossTx.Data.Relayer, "want", relayerAddr)
	}

	if crossTx.Data.GasUsed != receipt.GasUsed {
		return errors.New("gas used mismatch", "got", crossTx.Data.GasUsed, "want", receipt.GasUsed)
	}

	var success bool
	if len(crossTx.Data.Error) == 0 {
		success = true
	}
	if success != receipt.Success {
		return errors.New("success/error mismatch", "got_error", crossTx.Data.Error.String(), "want_success", receipt.Success)
	}

	switch strings.ToLower(crossTx.Data.ConfirmationLevel) {
	case "finalized":
		if receipt.ConfLevel != xchain.ConfFinalized {
			return errors.New("conf level mismatch", "got", crossTx.Data.ConfirmationLevel, "want", receipt.ConfLevel)
		}
	case "latest":
		if receipt.ConfLevel != xchain.ConfLatest {
			return errors.New("conf level mismatch", "got", crossTx.Data.ConfirmationLevel, "want", receipt.ConfLevel)
		}

		if receipt.MsgID.ConfLevel() == xchain.ConfFinalized {
			return errors.New("finalized shard delivered by fuzzy attestation", "level", crossTx.Data.ConfirmationLevel, "shard", receipt.MsgID.ShardID.Label())
		}
	default:
		return errors.New("unknown confirmation level", "level", crossTx.Data.ConfirmationLevel)
	}

	return nil
}

func verifySrc(crossTx crossTxJSON, src source, msg xchain.Msg) error {
	if common.HexToHash(crossTx.SrcBlockHash) != src.XBlock.BlockHash {
		return errors.New("block hash mismatch", "got", crossTx.SrcBlockHash, "want", src.XBlock.BlockHash)
	}

	if common.HexToHash(crossTx.SrcTxHash) != msg.TxHash {
		return errors.New("tx hash mismatch", "got", crossTx.SrcTxHash, "want", msg.TxHash)
	}

	if !crossTx.SrcTimestamp.Equal(src.XBlock.Timestamp) {
		return errors.New("timestamp mismatch", "got", crossTx.SrcTimestamp, "want", src.XBlock.Timestamp)
	}

	if common.HexToAddress(crossTx.From) != msg.SourceMsgSender {
		return errors.New("sender/from mismatch", "got", crossTx.From, "want", src.Sender)
	}

	if common.HexToAddress(crossTx.To) != msg.DestAddress {
		return errors.New("destination/to mismatch", "got", crossTx.To, "want", msg.DestAddress)
	}

	if crossTx.Data.GasLimit != msg.DestGasLimit {
		return errors.New("gas limit mismatch", "got", crossTx.Data.GasLimit, "want", msg.DestGasLimit)
	}

	return nil
}

type source struct {
	Sender common.Address
	Chain  netconf.Chain
	Tx     *types.Transaction
	XBlock xchain.Block
}

func fetchSource(
	ctx context.Context,
	network netconf.Network,
	xprov xchain.Provider,
	ethCls map[uint64]ethclient.Client,
	chainID chainID,
	txHash string,
	blockNumber uint64,
) (source, error) {
	cID, err := chainID.ID()
	if err != nil {
		return source{}, err
	}

	chain, ok := network.Chain(cID)
	if !ok {
		return source{}, errors.New("unknown chain id", "id", cID)
	}

	tx, err := getTx(ctx, ethCls[chain.ID], txHash)
	if err != nil {
		return source{}, err
	}

	xBlock, ok, err := xprov.GetBlock(ctx, xchain.ProviderRequest{
		ChainID:   chain.ID,
		Height:    blockNumber,
		ConfLevel: xchain.ConfLatest,
	})
	if err != nil {
		return source{}, err
	} else if !ok {
		return source{}, errors.New("block not found")
	}

	sender, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
	if err != nil {
		return source{}, errors.Wrap(err, "extract source sender")
	}

	return source{
		Sender: sender,
		Chain:  chain,
		Tx:     tx,
		XBlock: xBlock,
	}, nil
}

func receiptByID(receipts []xchain.Receipt, id xchain.MsgID) (xchain.Receipt, error) {
	for _, receipt := range receipts {
		if receipt.MsgID == id {
			return receipt, nil
		}
	}

	return xchain.Receipt{}, errors.New("receipt not found", "id", id)
}

func msgByID(msgs []xchain.Msg, id xchain.MsgID) (xchain.Msg, error) {
	for _, msg := range msgs {
		if msg.MsgID == id {
			return msg, nil
		}
	}

	return xchain.Msg{}, errors.New("msg not found", "id", id)
}

func getTx(ctx context.Context, ethCl ethclient.Client, hash string) (*types.Transaction, error) {
	if ethCl == nil {
		return nil, errors.New("missing eth client")
	}

	tx, isPending, err := ethCl.TransactionByHash(ctx, common.HexToHash(hash))
	if err != nil {
		return nil, err
	} else if isPending {
		return nil, errors.New("pending tx")
	}

	return tx, nil
}
