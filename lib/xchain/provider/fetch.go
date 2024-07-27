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

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"golang.org/x/sync/errgroup"
)

// ChainVersionHeight returns the latest height for the provided chain version.
func (p *Provider) ChainVersionHeight(ctx context.Context, chainVer xchain.ChainVersion) (uint64, error) {
	if chainVer.ID == p.cChainID {
		// Consensus chain versions all reduce to `latest`.
		xblock, ok, err := p.cProvider.XBlock(ctx, 0, true)
		if err != nil {
			return 0, errors.Wrap(err, "fetch consensus xblock")
		} else if !ok {
			return 0, errors.Wrap(err, "unexpected missing latest block [BUG]")
		}

		return xblock.BlockHeight, nil
	}

	_, ethCl, err := p.getEVMChain(chainVer.ID)
	if err != nil {
		return 0, err
	}

	headType, ok := headTypeFromConfLevel(chainVer.ConfLevel)
	if !ok {
		return 0, errors.New("unsupported conf level")
	}

	header, err := ethCl.HeaderByType(ctx, headType)
	if err != nil {
		return 0, err
	}

	return header.Number.Uint64(), nil
}

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
		// Once we add multiple streams, we need to query portal module offset table using latest or historical blocks.
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

	offset, err := caller.OutXMsgOffset(opts, stream.DestChainID, uint64(stream.ShardID))
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

	msgOffset, err := caller.InXMsgOffset(callOpts, stream.SourceChainID, uint64(stream.ShardID))
	if err != nil {
		return xchain.SubmitCursor{}, false, errors.Wrap(err, "call InXMsgOffset")
	}

	if msgOffset == 0 {
		return xchain.SubmitCursor{}, false, nil
	}

	blockOffset, err := caller.InXBlockOffset(callOpts, stream.SourceChainID, uint64(stream.ShardID))
	if err != nil {
		return xchain.SubmitCursor{}, false, errors.Wrap(err, "call InXBlockOffset")
	}

	return xchain.SubmitCursor{
		StreamID:    stream,
		MsgOffset:   msgOffset,
		BlockOffset: blockOffset,
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
		} else if b.BlockHeight != req.Height {
			return xchain.Block{}, false, errors.New("unexpected block height [BUG]")
		}

		return b, true, nil
	}

	_, ethCl, err := p.getEVMChain(req.ChainID)
	if err != nil {
		return xchain.Block{}, false, err
	}

	// An xblock is constructed from an eth header, and xmsg logs, and xreceipt logs.
	var (
		header   *types.Header
		msgs     []xchain.Msg
		receipts []xchain.Receipt
	)

	// First check if height is confirmed.
	if !p.confirmedCache(req.ChainVersion(), req.Height) {
		// No higher cached header available, so fetch the latest head
		latest, err := p.headerByChainVersion(ctx, req.ChainVersion())
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

	// Fetch the header if we didn't find it in the cache
	if header == nil {
		header, err = ethCl.HeaderByNumber(ctx, big.NewInt(int64(req.Height)))
		if err != nil {
			return xchain.Block{}, false, errors.Wrap(err, "header by number")
		}
	}

	// Fetch the msgs and receipts in parallel.
	var eg errgroup.Group
	eg.Go(func() error {
		var err error
		msgs, err = p.getXMsgLogs(ctx, req.ChainID, header.Hash())

		return err
	})
	eg.Go(func() error {
		var err error
		receipts, err = p.getXReceiptLogs(ctx, req.ChainID, header.Hash())

		return err
	})

	if err := eg.Wait(); err != nil {
		return xchain.Block{}, false, errors.Wrap(err, "wait")
	}

	return xchain.Block{
		BlockHeader: xchain.BlockHeader{
			ChainID:     req.ChainID,
			BlockHeight: req.Height,
			BlockHash:   header.Hash(),
		},
		Msgs:       msgs,
		Receipts:   receipts,
		ParentHash: header.ParentHash,
		Timestamp:  time.Unix(int64(header.Time), 0),
	}, true, nil
}

func (p *Provider) getXReceiptLogs(ctx context.Context, chainID uint64, blockHash common.Hash) ([]xchain.Receipt, error) {
	ctx, span := tracer.Start(ctx, spanName("get_receipt_logs"))
	defer span.End()

	chain, rpcClient, err := p.getEVMChain(chainID)
	if err != nil {
		return nil, errors.Wrap(err, "get evm chain")
	}

	logs, err := getLogs(ctx, rpcClient, chain.PortalAddress, blockHash, "XReceipt")
	if err != nil {
		return nil, errors.Wrap(err, "get xreceipt logs")
	}

	expectedShards := make(map[uint64]bool)
	for _, stream := range p.network.StreamsTo(chainID) {
		expectedShards[uint64(stream.ShardID)] = true
	}

	filterer, err := bindings.NewOmniPortalFilterer(chain.PortalAddress, rpcClient)
	if err != nil {
		return nil, errors.Wrap(err, "new filterer")
	}

	var receipts []xchain.Receipt
	for _, xreceiptLog := range logs {
		e, err := filterer.ParseXReceipt(xreceiptLog)
		if err != nil {
			return nil, errors.Wrap(err, "parse xreceipt log")
		}

		if !expectedShards[e.ShardId] {
			return nil, errors.New("unexpected receipt shard",
				"shard", e.ShardId,
				"src_chain", e.SourceChainId,
				"expected", p.network.StreamsBetween(e.SourceChainId, chainID),
			)
		}

		receipts = append(receipts, xchain.Receipt{
			MsgID: xchain.MsgID{
				StreamID: xchain.StreamID{
					SourceChainID: e.SourceChainId,
					DestChainID:   chain.ID,
					ShardID:       xchain.ShardID(e.ShardId),
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

	return receipts, nil
}

func (p *Provider) getXMsgLogs(ctx context.Context, chainID uint64, blockHash common.Hash) ([]xchain.Msg, error) {
	ctx, span := tracer.Start(ctx, spanName("get_msg_logs"))
	defer span.End()

	chain, rpcClient, err := p.getEVMChain(chainID)
	if err != nil {
		return nil, errors.Wrap(err, "get evm chain")
	}

	logs, err := getLogs(ctx, rpcClient, chain.PortalAddress, blockHash, "XMsg")
	if err != nil {
		return nil, errors.Wrap(err, "get xmsg logs")
	}

	expectedShards := make(map[uint64]bool)
	for _, shard := range chain.Shards {
		expectedShards[uint64(shard)] = true
	}

	filterer, err := bindings.NewOmniPortalFilterer(chain.PortalAddress, rpcClient)
	if err != nil {
		return nil, err
	}

	var xmsgs []xchain.Msg
	for _, xmsgLog := range logs {
		e, err := filterer.ParseXMsg(xmsgLog)
		if err != nil {
			return nil, errors.Wrap(err, "parse xmsg log")
		}

		if !expectedShards[e.ShardId] {
			return nil, errors.New("unexpected xmsg shard", "shard", e.ShardId)
		}

		xmsgs = append(xmsgs, xchain.Msg{
			MsgID: xchain.MsgID{
				StreamID: xchain.StreamID{
					SourceChainID: chain.ID,
					DestChainID:   e.DestChainId,
					ShardID:       xchain.ShardID(e.ShardId),
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

	return xmsgs, nil
}

// confirmedCache returns true if the height is confirmedCache based on the chain version
// on the cached strategy head.
func (p *Provider) confirmedCache(chainVer xchain.ChainVersion, height uint64) bool {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.confHeads[chainVer] >= height
}

// headerByChainVersion returns the chain's header by confirmation level (finalization/latest)
// by querying via ethclient. It caches the result.
func (p *Provider) headerByChainVersion(ctx context.Context, chainVer xchain.ChainVersion) (*types.Header, error) {
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
	case xchain.ConfFinalized:
		return ethclient.HeadFinalized, true
	default:
		return "", false
	}
}

func getLogs(ctx context.Context, rpcClient ethclient.Client, contractAddr common.Address, blockHash common.Hash, topicName string) ([]types.Log, error) {
	portalAbi, err := bindings.OmniPortalMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "get abi")
	}

	logs, err := rpcClient.FilterLogs(ctx, ethereum.FilterQuery{
		BlockHash: &blockHash,
		Addresses: []common.Address{contractAddr},
		Topics:    [][]common.Hash{{portalAbi.Events[topicName].ID}},
	})
	if err != nil {
		return nil, errors.Wrap(err, "filter xreceipt logs")
	}

	return logs, nil
}
