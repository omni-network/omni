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
func (p *Provider) GetEmittedCursor(ctx context.Context, ref xchain.EmitRef, stream xchain.StreamID,
) (xchain.EmitCursor, bool, error) {
	if !ref.Valid() {
		return xchain.EmitCursor{}, false, errors.New("invalid emit ref")
	}

	if stream.SourceChainID == p.cChainID {
		// Consensus xblocks only has a single stream/shard for now, so just query the latest block.
		block, err := getConsXBlock(ctx, ref, p.cProvider)
		if err != nil {
			return xchain.EmitCursor{}, false, err
		} else if len(block.Msgs) == 0 {
			return xchain.EmitCursor{}, false, errors.New("no consensus xmsgs [BUG]")
		}

		return xchain.EmitCursor{
			StreamID:  stream,
			MsgOffset: block.Msgs[len(block.Msgs)-1].StreamOffset,
		}, true, nil
	}

	chain, rpcClient, err := p.getEVMChain(stream.SourceChainID)
	if err != nil {
		return xchain.EmitCursor{}, false, err
	}

	caller, err := bindings.NewOmniPortalCaller(chain.PortalAddress, rpcClient)
	if err != nil {
		return xchain.EmitCursor{}, false, errors.Wrap(err, "new caller")
	}

	opts := &bind.CallOpts{Context: ctx}
	if ref.Height != nil {
		opts.BlockNumber = big.NewInt(int64(*ref.Height))
	} else if head, ok := headTypeFromConfLevel(*ref.ConfLevel); !ok {
		return xchain.EmitCursor{}, false, errors.New("invalid conf level")
	} else {
		// Populate an explicit block number if not querying latest head.
		header, err := rpcClient.HeaderByType(ctx, head)
		if err != nil {
			return xchain.EmitCursor{}, false, err
		}

		opts.BlockNumber = header.Number
	}

	offset, err := caller.OutXMsgOffset(opts, stream.DestChainID, stream.ShardID)
	if err != nil {
		return xchain.EmitCursor{}, false, errors.Wrap(err, "call OutXMgsOffset")
	}

	if offset == 0 {
		return xchain.EmitCursor{}, false, nil
	}

	return xchain.EmitCursor{
		StreamID:  stream,
		MsgOffset: offset,
	}, true, nil
}

// GetSubmittedCursor returns the submitted cursor for the source chain on the destination chain,
// or false if not available, or an error. Calls the destination chain portal InXStreamOffset method.
func (p *Provider) GetSubmittedCursor(ctx context.Context, stream xchain.StreamID,
) (xchain.SubmitCursor, bool, error) {
	chain, rpcClient, err := p.getEVMChain(stream.DestChainID)
	if err != nil {
		return xchain.SubmitCursor{}, false, err
	}

	caller, err := bindings.NewOmniPortalCaller(chain.PortalAddress, rpcClient)
	if err != nil {
		return xchain.SubmitCursor{}, false, errors.Wrap(err, "new caller")
	}

	height, err := rpcClient.BlockNumber(ctx)
	if err != nil {
		return xchain.SubmitCursor{}, false, err
	}

	callOpts := &bind.CallOpts{Context: ctx, BlockNumber: big.NewInt(int64(height))}

	msgOffset, err := caller.InXMsgOffset(callOpts, stream.SourceChainID, stream.ShardID)
	if err != nil {
		return xchain.SubmitCursor{}, false, errors.Wrap(err, "call InXMsgOffset")
	}

	if msgOffset == 0 {
		return xchain.SubmitCursor{}, false, nil
	}

	blockOffset, err := caller.InXBlockOffset(callOpts, stream.SourceChainID, stream.ShardID)
	if err != nil {
		return xchain.SubmitCursor{}, false, errors.Wrap(err, "call InXBlockOffset")
	}

	valSetID, err := caller.InXStreamValidatorSetId(callOpts, stream.SourceChainID, stream.ShardID)
	if err != nil {
		return xchain.SubmitCursor{}, false, errors.Wrap(err, "call InXStreamValidatorSetId")
	}

	return xchain.SubmitCursor{
		StreamID:       stream,
		MsgOffset:      msgOffset,
		BlockOffset:    blockOffset,
		ValidatorSetID: valSetID,
	}, true, nil
}

// GetBlock returns the XBlock for the provided chain and height, or false if not available yet (not finalized),
// or an error.
func (p *Provider) GetBlock(ctx context.Context, req xchain.ProviderRequest) (xchain.Block, bool, error) {
	ctx, span := tracer.Start(ctx, spanName("get_block"))
	defer span.End()

	if req.ChainID == p.cChainID {
		b, ok, err := p.cProvider.XBlock(ctx, req.Height, false)
		if err != nil {
			return xchain.Block{}, false, errors.Wrap(err, "fetch consensus xblock")
		} else if !ok {
			return xchain.Block{}, false, nil
		} else if len(b.Msgs) == 0 {
			return xchain.Block{}, false, errors.New("empty consensus block [BUG]")
		} else if b.BlockHeight != req.Height && b.BlockOffset != req.Offset {
			return xchain.Block{}, false, errors.New("unexpected block height and offset [BUG]")
		}

		return b, true, nil
	}

	_, ethCl, err := p.getEVMChain(req.ChainID)
	if err != nil {
		return xchain.Block{}, false, err
	}

	chainVer := reqToChainVersion(req)

	// An xblock is constructed from an eth header, and xmsg logs, and xreceipt logs.
	var (
		header   *types.Header
		msgs     []xchain.Msg
		receipts []xchain.Receipt
	)

	// First check if height is confirmed.
	if !p.confirmedCache(chainVer, req.Height) {
		// No higher cached header available, so fetch the latest head
		latest, err := p.headerByChainVersion(ctx, chainVer)
		if err != nil {
			return xchain.Block{}, false, errors.Wrap(err, "header by strategy")
		}

		// If still lower, we reached the head of the chain, return false
		if latest.Number.Uint64() < req.Height {
			return xchain.Block{}, false, nil
		}

		// Use this header if it matches height
		if latest.Number.Uint64() == req.Height {
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
		header, err = ethCl.HeaderByNumber(ctx, big.NewInt(int64(req.Height)))

		return err
	})
	eg.Go(func() error {
		var err error
		msgs, err = p.getXMsgLogs(ctx, req.ChainID, req.Height)

		return err
	})
	eg.Go(func() error {
		var err error
		receipts, err = p.getXReceiptLogs(ctx, req.ChainID, req.Height)

		return err
	})

	if err := eg.Wait(); err != nil {
		return xchain.Block{}, false, errors.Wrap(err, "wait")
	}

	resp := xchain.Block{
		BlockHeader: xchain.BlockHeader{
			SourceChainID: req.ChainID,
			ConfLevel:     req.ConfLevel,
			BlockHeight:   req.Height,
			BlockHash:     header.Hash(),
		},
		Msgs:       msgs,
		Receipts:   receipts,
		ParentHash: header.ParentHash,
		Timestamp:  time.Unix(int64(header.Time), 0),
	}
	if resp.ShouldAttest() {
		resp.BlockOffset = req.Offset
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

	expectedShards := make(map[uint64]bool)
	for _, shard := range chain.Shards {
		expectedShards[shard] = true
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

	iter, err := filterer.FilterXReceipt(&filterOpts, nil, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "filter receipts logs")
	}

	var receipts []xchain.Receipt
	for iter.Next() {
		e := iter.Event

		if !expectedShards[e.ShardId] {
			return nil, errors.New("unexpected receipt shard", "shard", e.ShardId)
		}

		receipts = append(receipts, xchain.Receipt{
			MsgID: xchain.MsgID{
				StreamID: xchain.StreamID{
					SourceChainID: e.SourceChainId,
					DestChainID:   chain.ID,
					ShardID:       e.ShardId,
				},
				StreamOffset: e.Offset,
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

	expectedShards := make(map[uint64]bool)
	for _, shard := range chain.Shards {
		expectedShards[shard] = true
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

	iter, err := filterer.FilterXMsg(&filterOpts, nil, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "filter xmsg logs")
	}

	var xmsgs []xchain.Msg
	for iter.Next() {
		e := iter.Event

		if !expectedShards[e.ShardId] {
			return nil, errors.New("unexpected xmsg shard", "shard", e.ShardId)
		}

		xmsgs = append(xmsgs, xchain.Msg{
			MsgID: xchain.MsgID{
				StreamID: xchain.StreamID{
					SourceChainID: chain.ID,
					DestChainID:   e.DestChainId,
					ShardID:       e.ShardId,
				},
				StreamOffset: e.Offset,
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

// confirmedCache returns true if the height is confirmedCache based on the chain version
// on the cached strategy head.
func (p *Provider) confirmedCache(chain chainVersion, height uint64) bool {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.confHeads[chain] >= height
}

// headerByChainVersion returns the chain's header by confirmation level (finalization/latest)
// by querying via ethclient. It caches the result.
func (p *Provider) headerByChainVersion(ctx context.Context, chainVer chainVersion) (*types.Header, error) {
	_, rpcClient, err := p.getEVMChain(chainVer.ID)
	if err != nil {
		return nil, err
	}

	headType, ok := headTypeFromConfLevel(chainVer.ConfLevel)
	if !ok {
		return nil, errors.New("unsupported conf level")
	}

	// Fetch the header from the ethclient
	header, err := rpcClient.HeaderByType(ctx, headType)
	if err != nil {
		return nil, err
	}

	// Update the strategy cache
	p.mu.Lock()
	defer p.mu.Unlock()
	p.confHeads[chainVer] = header.Number.Uint64()

	return header, nil
}

func getConsXBlock(ctx context.Context, ref xchain.EmitRef, cprov cchain.Provider) (xchain.Block, error) {
	var height uint64
	var latest bool
	if ref.Height != nil {
		height = *ref.Height
	} else if ref.ConfLevel != nil {
		// For consensus chain (instant finality), we can query the latest consensus xblock.
		latest = true
	}

	xblock, ok, err := cprov.XBlock(ctx, height, latest)
	if err != nil {
		return xchain.Block{}, errors.Wrap(err, "fetch consensus xblock")
	} else if !ok {
		return xchain.Block{}, errors.New("no consensus xblocks [BUG]")
	} else if xblock.Msgs[0].DestChainID != 0 {
		return xchain.Block{}, errors.New("non-broadcast consensus chain xmsg [BUG]")
	}

	return xblock, nil
}

func spanName(method string) string {
	return "xprovider/" + method
}

func headTypeFromConfLevel(conf xchain.ConfLevel) (ethclient.HeadType, bool) {
	switch conf {
	case xchain.ConfLatest:
		return ethclient.HeadLatest, true
	case xchain.ConfSafe:
		return ethclient.HeadSafe, true
	case xchain.ConfFinalized:
		return ethclient.HeadFinalized, true
	default:
		return "", false
	}
}

func reqToChainVersion(req xchain.ProviderRequest) chainVersion {
	return chainVersion{ID: req.ChainID, ConfLevel: req.ConfLevel}
}
