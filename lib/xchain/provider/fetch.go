package provider

import (
	"context"
	"math/big"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// GetEmittedCursor returns the emitted cursor for the destination chain on the source chain,
// or false if not available, or an error. Calls the source chain portal OutXStreamOffset method.
func (p *Provider) GetEmittedCursor(ctx context.Context, sourceChainID uint64, destinationChainID uint64,
) (xchain.StreamCursor, bool, error) {
	chain, rpcClient, err := p.getChain(sourceChainID)
	if err != nil {
		return xchain.StreamCursor{}, false, err
	}

	height, err := rpcClient.BlockNumber(ctx)
	if err != nil {
		return xchain.StreamCursor{}, false, errors.Wrap(err, "get block number")
	}

	caller, err := bindings.NewOmniPortalCaller(common.HexToAddress(chain.PortalAddress), rpcClient)
	if err != nil {
		return xchain.StreamCursor{}, false, errors.Wrap(err, "new caller")
	}

	opts := &bind.CallOpts{Context: ctx, BlockNumber: big.NewInt(int64(height))}
	offset, err := caller.OutXStreamOffset(opts, destinationChainID)
	if err != nil {
		return xchain.StreamCursor{}, false, errors.Wrap(err, "call inXStreamOffset")
	}

	if offset == 0 {
		return xchain.StreamCursor{}, false, nil
	}

	return xchain.StreamCursor{
		StreamID: xchain.StreamID{
			SourceChainID: sourceChainID,
			DestChainID:   destinationChainID,
		},
		Offset:            offset,
		SourceBlockHeight: height,
	}, true, nil
}

// GetSubmittedCursor returns the submitted cursor for the source chain on the destination chain,
// or false if not available, or an error. Calls the destination chain portal InXStreamOffset method.
func (p *Provider) GetSubmittedCursor(ctx context.Context, destChainID uint64, sourceChainID uint64,
) (xchain.StreamCursor, bool, error) {
	chain, rpcClient, err := p.getChain(destChainID)
	if err != nil {
		return xchain.StreamCursor{}, false, err
	}

	caller, err := bindings.NewOmniPortalCaller(common.HexToAddress(chain.PortalAddress), rpcClient)
	if err != nil {
		return xchain.StreamCursor{}, false, errors.Wrap(err, "new caller")
	}

	offset, err := caller.InXStreamOffset(&bind.CallOpts{Context: ctx}, sourceChainID)
	if err != nil {
		return xchain.StreamCursor{}, false, errors.Wrap(err, "call inXStreamOffset")
	}

	if offset == 0 {
		return xchain.StreamCursor{}, false, nil
	}

	blockHeight, err := caller.InXStreamBlockHeight(&bind.CallOpts{Context: ctx}, sourceChainID)
	if err != nil {
		return xchain.StreamCursor{}, false, errors.Wrap(err, "call inXStreamBlockHeight")
	}

	return xchain.StreamCursor{
		StreamID: xchain.StreamID{
			SourceChainID: sourceChainID,
			DestChainID:   destChainID,
		},
		Offset:            offset,
		SourceBlockHeight: blockHeight,
	}, true, nil
}

// GetBlock returns the XBlock for the provided chain and height, or false if not available yet (not finalized),
// or an error.
func (p *Provider) GetBlock(ctx context.Context, chainID uint64, height uint64) (xchain.Block, bool, error) {
	if chainID == cChainID {
		return p.cprov.XBlock(ctx, height)
	}

	chain, rpcClient, err := p.getChain(chainID)
	if err != nil {
		return xchain.Block{}, false, err
	}

	finalisedHeader, err := rpcClient.HeaderByType(ctx, ethclient.HeadType(chain.FinalizationStrat))
	if err != nil {
		return xchain.Block{}, false, err
	}

	// ignore if our height is greater than the finalized height
	if height > finalisedHeader.Number.Uint64() {
		return xchain.Block{}, false, nil
	}

	// check if we can reuse the header
	header := finalisedHeader
	if height != finalisedHeader.Number.Uint64() {
		// fetch the block header for the given height
		header, err = rpcClient.HeaderByNumber(ctx, big.NewInt(int64(height)))
		if err != nil {
			return xchain.Block{}, false, errors.Wrap(err, "could not get header by number")
		}
	}

	// Filter xmsgs logs.
	xmsgs, err := p.getXMsgLogs(ctx, chainID, height)
	if err != nil {
		return xchain.Block{}, false, err
	}

	// Filter xreceipts logs.
	receipts, err := p.getXReceiptLogs(ctx, chainID, height)
	if err != nil {
		return xchain.Block{}, false, err
	}

	return xchain.Block{
		BlockHeader: xchain.BlockHeader{
			SourceChainID: chainID,
			BlockHeight:   height,
			BlockHash:     header.Hash(),
		},
		Msgs:      xmsgs,
		Receipts:  receipts,
		Timestamp: time.Unix(int64(header.Time), 0),
	}, true, nil
}

func (p *Provider) getXReceiptLogs(ctx context.Context, chainID uint64, height uint64) ([]xchain.Receipt, error) {
	chain, rpcClient, err := p.getChain(chainID)
	if err != nil {
		return nil, err
	}

	filterer, err := bindings.NewOmniPortalFilterer(common.HexToAddress(chain.PortalAddress), rpcClient)
	if err != nil {
		return nil, errors.Wrap(err, "new filterer")
	}

	filterOpts := bind.FilterOpts{
		Start:   height,
		End:     &height,
		Context: ctx,
	}

	iter, err := filterer.FilterXReceipt(&filterOpts, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "filter receipts logs")
	}

	var receipts []xchain.Receipt
	for iter.Next() {
		e := iter.Event
		receipts = append(receipts, xchain.Receipt{
			MsgID: xchain.MsgID{
				StreamID: xchain.StreamID{
					SourceChainID: e.SourceChainId,
					DestChainID:   chain.ID,
				},
				StreamOffset: e.StreamOffset,
			},
			GasUsed:        e.GasUsed.Uint64(),
			Success:        e.Success,
			RelayerAddress: e.Relayer,
			TxHash:         e.Raw.TxHash,
		})
	}
	if err := iter.Error(); err != nil {
		return nil, errors.Wrap(err, "iterate receipts logs")
	}

	return receipts, nil
}

func (p *Provider) getXMsgLogs(ctx context.Context, chainID uint64, height uint64) ([]xchain.Msg, error) {
	chain, rpcClient, err := p.getChain(chainID)
	if err != nil {
		return nil, err
	}

	filterer, err := bindings.NewOmniPortalFilterer(common.HexToAddress(chain.PortalAddress), rpcClient)
	if err != nil {
		return nil, err
	}

	filterOpts := bind.FilterOpts{
		Start:   height,
		End:     &height,
		Context: ctx,
	}

	iter, err := filterer.FilterXMsg(&filterOpts, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "filter xmsg logs")
	}

	var xmsgs []xchain.Msg
	for iter.Next() {
		e := iter.Event
		xmsgs = append(xmsgs, xchain.Msg{
			MsgID: xchain.MsgID{
				StreamID: xchain.StreamID{
					SourceChainID: chain.ID,
					DestChainID:   e.DestChainId,
				},
				StreamOffset: e.StreamOffset,
			},
			SourceMsgSender: e.Sender,
			DestAddress:     e.To,
			Data:            e.Data,
			DestGasLimit:    e.GasLimit,
			TxHash:          e.Raw.TxHash,
		})
	}
	if err := iter.Error(); err != nil {
		return nil, errors.Wrap(err, "iterate xmsg logs")
	}

	return xmsgs, nil
}
