package provider

import (
	"context"
	"math/big"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/tracer"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	"golang.org/x/sync/errgroup"
)

// GetEmittedCursor returns the emitted cursor for the destination chain on the source chain,
// or false if not available, or an error. Calls the source chain portal OutXStreamOffset method.
//
// Note that the BlockOffset field is not populated for emit cursors, since it isn't stored on-chain
// but tracked off-chain.
func (p *Provider) GetEmittedCursor(ctx context.Context, ref xchain.EmitRef, sourceChainID uint64, destinationChainID uint64,
) (xchain.StreamCursor, bool, error) {
	const unknownBlockOffset uint64 = 0
	if sourceChainID == p.cChainID {
		block, err := getConsXBlock(ctx, ref, p.cProvider)
		if err != nil {
			return xchain.StreamCursor{}, false, err
		}

		return xchain.StreamCursor{
			StreamID: xchain.StreamID{
				SourceChainID: sourceChainID,
				DestChainID:   destinationChainID, // Consensus xmsgs are broadcast, so use the provided destination chain ID.
			},
			MsgOffset:   block.Msgs[0].StreamOffset, // Consensus xblocks only have a single xmsg.
			BlockOffset: unknownBlockOffset,
		}, true, nil
	}

	chain, rpcClient, err := p.getEVMChain(sourceChainID)
	if err != nil {
		return xchain.StreamCursor{}, false, err
	}

	caller, err := bindings.NewOmniPortalCaller(chain.PortalAddress, rpcClient)
	if err != nil {
		return xchain.StreamCursor{}, false, errors.Wrap(err, "new caller")
	}

	opts := &bind.CallOpts{Context: ctx}
	if ref.HeadType != nil && *ref.HeadType != ethclient.HeadLatest {
		// Populate an explicit block number if not querying latest head.
		header, err := rpcClient.HeaderByType(ctx, *ref.HeadType)
		if err != nil {
			return xchain.StreamCursor{}, false, err
		}

		opts.BlockNumber = header.Number
	} else if ref.Height != nil {
		opts.BlockNumber = big.NewInt(int64(*ref.Height))
	}

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
		MsgOffset:   offset,
		BlockOffset: unknownBlockOffset,
	}, true, nil
}

// GetSubmittedCursor returns the submitted cursor for the source chain on the destination chain,
// or false if not available, or an error. Calls the destination chain portal InXStreamOffset method.
func (p *Provider) GetSubmittedCursor(ctx context.Context, destChainID uint64, sourceChainID uint64,
) (xchain.StreamCursor, bool, error) {
	chain, rpcClient, err := p.getEVMChain(destChainID)
	if err != nil {
		return xchain.StreamCursor{}, false, err
	}

	caller, err := bindings.NewOmniPortalCaller(chain.PortalAddress, rpcClient)
	if err != nil {
		return xchain.StreamCursor{}, false, errors.Wrap(err, "new caller")
	}

	height, err := rpcClient.BlockNumber(ctx)
	if err != nil {
		return xchain.StreamCursor{}, false, err
	}

	callOpts := &bind.CallOpts{Context: ctx, BlockNumber: big.NewInt(int64(height))}

	// TODO(corver): Rename portal variable to InXStreamMsgOffset
	msgOffset, err := caller.InXStreamOffset(callOpts, sourceChainID)
	if err != nil {
		return xchain.StreamCursor{}, false, errors.Wrap(err, "call inXStreamOffset")
	}

	if msgOffset == 0 {
		return xchain.StreamCursor{}, false, nil
	}

	// TODO(corver): Rename portal variable to InXStreamBlockOffset
	blockOffset, err := caller.InXStreamBlockHeight(callOpts, sourceChainID)
	if err != nil {
		return xchain.StreamCursor{}, false, errors.Wrap(err, "call inXStreamBlockHeight")
	}

	return xchain.StreamCursor{
		StreamID: xchain.StreamID{
			SourceChainID: sourceChainID,
			DestChainID:   destChainID,
		},
		MsgOffset:   msgOffset,
		BlockOffset: blockOffset,
	}, true, nil
}

// GetBlock returns the XBlock for the provided chain and height, or false if not available yet (not finalized),
// or an error.
func (p *Provider) GetBlock(ctx context.Context, chainID uint64, height uint64, xOffset uint64) (xchain.Block, bool, error) {
	ctx, span := tracer.Start(ctx, spanName("get_block"))
	defer span.End()

	if chainID == p.cChainID {
		b, ok, err := p.cProvider.XBlock(ctx, height, false)
		if err != nil {
			return xchain.Block{}, false, errors.Wrap(err, "fetch consensus xblock")
		} else if !ok {
			return xchain.Block{}, false, nil
		} else if b.BlockHeight != height && b.BlockOffset != xOffset {
			return xchain.Block{}, false, errors.New("unexpected block height and offset [BUG]")
		}

		return b, true, nil
	}

	_, ethCl, err := p.getEVMChain(chainID)
	if err != nil {
		return xchain.Block{}, false, err
	}

	// An xblock is constructed from an eth header, and xmsg logs, and xreceipt logs.
	var (
		header   *types.Header
		msgs     []xchain.Msg
		receipts []xchain.Receipt
	)

	// First check if height is finalized by the chain's finalization strategy.
	if !p.finalisedInCache(chainID, height) {
		// No higher cached header available, so fetch the latest head
		latest, err := p.headerByStrategy(ctx, chainID)
		if err != nil {
			return xchain.Block{}, false, errors.Wrap(err, "header by strategy")
		}

		// If still lower, we reached the head of the chain, return false
		if latest.Number.Uint64() < height {
			return xchain.Block{}, false, nil
		}

		// Use this header if it matches height
		if latest.Number.Uint64() == height {
			header = latest
		}
	}

	// Fetch the msgs and receipts (and header if required) in parallel.
	var eg errgroup.Group
	eg.Go(func() error {
		if header != nil {
			return nil // No need to fetch header again.
		}

		var err error
		header, err = ethCl.HeaderByNumber(ctx, big.NewInt(int64(height)))

		return err
	})
	eg.Go(func() error {
		var err error
		msgs, err = p.getXMsgLogs(ctx, chainID, height)

		return err
	})
	eg.Go(func() error {
		var err error
		receipts, err = p.getXReceiptLogs(ctx, chainID, height)

		return err
	})

	if err := eg.Wait(); err != nil {
		return xchain.Block{}, false, errors.Wrap(err, "wait")
	}

	resp := xchain.Block{
		BlockHeader: xchain.BlockHeader{
			SourceChainID: chainID,
			BlockHeight:   height,
			BlockHash:     header.Hash(),
		},
		Msgs:      msgs,
		Receipts:  receipts,
		Timestamp: time.Unix(int64(header.Time), 0),
	}
	if resp.ShouldAttest() {
		resp.BlockOffset = xOffset
	}

	return resp, true, nil
}

func (p *Provider) getXReceiptLogs(ctx context.Context, chainID uint64, height uint64) ([]xchain.Receipt, error) {
	ctx, span := tracer.Start(ctx, spanName("get_receipt_logs"))
	defer span.End()

	chain, rpcClient, err := p.getEVMChain(chainID)
	if err != nil {
		return nil, err
	}

	filterer, err := bindings.NewOmniPortalFilterer(chain.PortalAddress, rpcClient)
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
			Error:          e.Error,
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
	ctx, span := tracer.Start(ctx, spanName("get_msg_logs"))
	defer span.End()

	chain, rpcClient, err := p.getEVMChain(chainID)
	if err != nil {
		return nil, err
	}

	filterer, err := bindings.NewOmniPortalFilterer(chain.PortalAddress, rpcClient)
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

// finalisedInCache returns true if the chain height is finalized based
// on the cached strategy head.
func (p *Provider) finalisedInCache(chainID uint64, height uint64) bool {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.stratHeads[chainID] >= height
}

// headerByStrategy returns the chain's header by strategy (finalization/latest)
// by querying via ethclient. It caches the result.
func (p *Provider) headerByStrategy(ctx context.Context, chainID uint64) (*types.Header, error) {
	chain, rpcClient, err := p.getEVMChain(chainID)
	if err != nil {
		return nil, err
	}

	// Fetch the header from the ethclient
	header, err := rpcClient.HeaderByType(ctx, ethclient.HeadType(chain.FinalizationStrat))
	if err != nil {
		return nil, err
	}

	// Update the strategy cache
	p.mu.Lock()
	defer p.mu.Unlock()
	p.stratHeads[chainID] = header.Number.Uint64()

	return header, nil
}

func getConsXBlock(ctx context.Context, ref xchain.EmitRef, cprov cchain.Provider) (xchain.Block, error) {
	var height uint64
	var latest bool
	if ref.Height != nil {
		height = *ref.Height
	} else if ref.HeadType != nil {
		// For consensus chain (instant finality), we can query the latest consensus xblock.
		latest = true
	}

	xblock, ok, err := cprov.XBlock(ctx, height, latest)
	if err != nil {
		return xchain.Block{}, errors.Wrap(err, "fetch consensus xblock")
	} else if !ok {
		return xchain.Block{}, errors.New("no consensus xblocks [BUG]")
	} else if len(xblock.Msgs) != 1 {
		return xchain.Block{}, errors.New("unexpected xblock msg conut [BUG]")
	} else if xblock.Msgs[0].DestChainID != 0 {
		return xchain.Block{}, errors.New("non-broadcast consensus chain xmsg [BUG]")
	}

	return xblock, nil
}

func spanName(method string) string {
	return "xprovider/" + method
}
