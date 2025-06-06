package provider

import (
	"context"
	"math/big"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/tracer"
	"github.com/omni-network/omni/lib/umath"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ChainVersionHeight returns the latest height for the provided chain version.
func (p *Provider) ChainVersionHeight(ctx context.Context, chainVer xchain.ChainVersion) (uint64, error) {
	if chainVer.ID == p.cChainID {
		if p.cProvider == nil {
			return 0, errors.New("consensus provider not set")
		}

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

	return heightByConfLevel(ctx, ethCl, chainVer.ConfLevel)
}

// heightByConfLevel returns the current height for the provided conf level.
// This is slightly more efficient for MinX conf levels than headerByConfLevel.
func heightByConfLevel(ctx context.Context, ethCl ethclient.Client, confLevel xchain.ConfLevel) (uint64, error) {
	// Some conf levels are directly queryable
	headers := map[xchain.ConfLevel]ethclient.HeadType{
		xchain.ConfFinalized: ethclient.HeadFinalized,
	}
	if headType, ok := headers[confLevel]; ok {
		header, err := ethCl.HeaderByType(ctx, headType)
		if err != nil {
			return 0, err
		}

		return header.Number.Uint64(), nil
	}

	// Other conf levels are related to `latest` height
	delay := confLevel.MinX()
	if delay == 0 && confLevel != xchain.ConfLatest {
		return 0, errors.New("unsupported 0-delay conf level")
	}

	latest, err := ethCl.BlockNumber(ctx)
	if err != nil {
		return 0, err
	}

	return umath.SubtractOrZero(latest, delay), nil // Use SubtractOrZero since genesis always confirmed for all levels
}

// headerByConfLevel returns the current header for the provided conf level.
func headerByConfLevel(ctx context.Context, ethCl ethclient.Client, confLevel xchain.ConfLevel) (*types.Header, error) {
	var headType ethclient.HeadType
	if confLevel == xchain.ConfLatest || confLevel.MinX() > 0 {
		headType = ethclient.HeadLatest
	} else if confLevel == xchain.ConfFinalized {
		headType = ethclient.HeadFinalized
	} else {
		return nil, errors.New("unsupported conf level")
	}

	header, err := ethCl.HeaderByType(ctx, headType)
	if err != nil {
		return nil, errors.Wrap(err, "header by type")
	}

	// Walk back MinX blocks using parent hash (if applicable)
	for range confLevel.MinX() {
		// Return if we reached the genesis block
		if bi.IsZero(header.Number) && header.ParentHash == (common.Hash{}) {
			return header, nil
		}

		header, err = ethCl.HeaderByHash(ctx, header.ParentHash)
		if err != nil {
			return nil, errors.Wrap(err, "header cache by hash")
		}
	}

	return header, nil
}

// GetEmittedCursor returns the emitted cursor for the destination chain on the source chain,
// or false if not available, or an error. Calls the source chain portal OutXStreamOffset method.
//
// Note that the AttestOffset field is not populated for emit cursors, since it isn't stored on-chain
// but tracked off-chain.
func (p *Provider) GetEmittedCursor(ctx context.Context, ref xchain.Ref, stream xchain.StreamID,
) (xchain.EmitCursor, bool, error) {
	if !ref.Valid() {
		return xchain.EmitCursor{}, false, errors.New("invalid emit ref")
	}

	if stream.SourceChainID == p.cChainID {
		if p.cProvider == nil {
			return xchain.EmitCursor{}, false, errors.New("consensus provider not set")
		}

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
	} else if !chain.HasEmitPortal {
		return xchain.EmitCursor{}, false, errors.New("emit portal not available on source chain")
	}

	caller, err := bindings.NewOmniPortalCaller(chain.PortalAddress, rpcClient)
	if err != nil {
		return xchain.EmitCursor{}, false, errors.Wrap(err, "new caller")
	}

	height, err := heightForRef(ctx, rpcClient, ref)
	if err != nil {
		return xchain.EmitCursor{}, false, err
	}

	opts := &bind.CallOpts{Context: ctx, BlockNumber: height}

	offset, err := caller.OutXMsgOffset(opts, stream.DestChainID, uint64(stream.ShardID))
	if err != nil {
		return xchain.EmitCursor{}, false, errors.Wrap(err, "call OutXMsgOffset")
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
func (p *Provider) GetSubmittedCursor(
	ctx context.Context,
	ref xchain.Ref,
	stream xchain.StreamID,
) (xchain.SubmitCursor, bool, error) {
	chain, rpcClient, err := p.getEVMChain(stream.DestChainID)
	if err != nil {
		return xchain.SubmitCursor{}, false, err
	} else if !chain.HasSubmitPortal {
		return xchain.SubmitCursor{}, false, errors.New("submit portal not available on destination chain")
	}

	caller, err := bindings.NewOmniPortalCaller(chain.PortalAddress, rpcClient)
	if err != nil {
		return xchain.SubmitCursor{}, false, errors.Wrap(err, "new caller")
	}

	height, err := heightForRef(ctx, rpcClient, ref)
	if err != nil {
		return xchain.SubmitCursor{}, false, err
	}

	callOpts := &bind.CallOpts{Context: ctx, BlockNumber: height}

	msgOffset, err := caller.InXMsgOffset(callOpts, stream.SourceChainID, uint64(stream.ShardID))
	if err != nil {
		return xchain.SubmitCursor{}, false, errors.Wrap(err, "call InXMsgOffset")
	}

	if msgOffset == 0 {
		return xchain.SubmitCursor{}, false, nil
	}

	attestOffset, err := caller.InXBlockOffset(callOpts, stream.SourceChainID, uint64(stream.ShardID))
	if err != nil {
		return xchain.SubmitCursor{}, false, errors.Wrap(err, "call InXBlockOffset")
	}

	return xchain.SubmitCursor{
		StreamID:     stream,
		MsgOffset:    msgOffset,
		AttestOffset: attestOffset,
	}, true, nil
}

// GetBlock returns the XBlock for the provided chain and height, or false if not available yet (not finalized),
// or an error.
func (p *Provider) GetBlock(ctx context.Context, req xchain.ProviderRequest) (xchain.Block, bool, error) {
	ctx, span := tracer.Start(ctx, spanName("get_block"))
	defer span.End()

	if evmchain.IsSVM(req.ChainID) {
		return xchain.Block{}, false, errors.New("svm chains are not supported")
	}

	//nolint:nestif // Not so bad
	if req.ChainID == p.cChainID {
		if p.cProvider == nil {
			return xchain.Block{}, false, errors.New("consensus provider not set")
		}

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
		header, err = ethCl.HeaderByNumber(ctx, bi.N(req.Height))
		if err != nil {
			return xchain.Block{}, false, errors.Wrap(err, "header by number")
		}
	}

	msgs, receipts, err = p.getMsgsAndReceipts(ctx, req.ChainID, header.Hash())
	if err != nil {
		return xchain.Block{}, false, errors.Wrap(err, "get msgs and receipts")
	}

	timeSecs, err := umath.ToInt64(header.Time)
	if err != nil {
		return xchain.Block{}, false, err
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
		Timestamp:  time.Unix(timeSecs, 0),
	}, true, nil
}

// getMsgsAndReceipts returns the xmsgs and xreceipts for the chain and block hash.
//
//nolint:nestif // Not worth refactoring
func (p *Provider) getMsgsAndReceipts(ctx context.Context, chainID uint64, blockHash common.Hash) ([]xchain.Msg, []xchain.Receipt, error) {
	ctx, span := tracer.Start(ctx, spanName("get_mgs_and_receipts"))
	defer span.End()

	chain, rpcClient, err := p.getEVMChain(chainID)
	if err != nil {
		return nil, nil, errors.Wrap(err, "get evm chain")
	}

	portalAbi, err := bindings.OmniPortalMetaData.GetAbi()
	if err != nil {
		return nil, nil, errors.Wrap(err, "get abi")
	}

	msgEvent, ok := portalAbi.Events["XMsg"]
	if !ok {
		return nil, nil, errors.New("missing XMsg event [BUG]")
	}

	receiptEvent, ok := portalAbi.Events["XReceipt"]
	if !ok {
		return nil, nil, errors.New("missing XReceipt event [BUG]")
	}

	topics := []common.Hash{msgEvent.ID, receiptEvent.ID}
	addrs := []common.Address{chain.PortalAddress}
	events, err := getEventLogs(ctx, rpcClient, addrs, blockHash, topics)
	if err != nil {
		return nil, nil, errors.Wrap(err, "get logs")
	}

	expectedShards := make(map[xchain.ShardID]bool)
	for _, stream := range p.network.StreamsTo(chainID) {
		expectedShards[stream.ShardID] = true
	}
	verifyShard := func(shardID xchain.ShardID) error {
		if !expectedShards[shardID] {
			return errors.New("unexpected shard",
				"shard", shardID,
				"chain", chainID,
				"expected", p.network.StreamsTo(chainID),
			)
		}

		return nil
	}

	filterer, err := bindings.NewOmniPortalFilterer(chain.PortalAddress, rpcClient)
	if err != nil {
		return nil, nil, errors.Wrap(err, "new filterer")
	}

	var msgs []xchain.Msg
	var receipts []xchain.Receipt
	for _, event := range events {
		if event.Topics[0] == msgEvent.ID {
			msg, err := parseXMsg(filterer, event, chainID)
			if err != nil {
				return nil, nil, err
			} else if err := verifyShard(msg.ShardID); err != nil {
				return nil, nil, err
			}

			msgs = append(msgs, msg)
		} else if event.Topics[0] == receiptEvent.ID {
			receipt, err := parseXReceipt(filterer, event, chainID)
			if err != nil {
				return nil, nil, err
			} else if err := verifyShard(receipt.ShardID); err != nil {
				return nil, nil, err
			}

			receipts = append(receipts, receipt)
		} else {
			return nil, nil, errors.New("unexpected event topic")
		}
	}

	return msgs, receipts, nil
}

// GetSubmission returns the submission associated with the transaction hash or an error.
func (p *Provider) GetSubmission(ctx context.Context, chainID uint64, txHash common.Hash) (xchain.Submission, error) {
	chain, rpcClient, err := p.getEVMChain(chainID)
	if err != nil {
		return xchain.Submission{}, errors.Wrap(err, "get evm chain")
	}

	tx, _, err := rpcClient.TransactionByHash(ctx, txHash)
	if err != nil {
		return xchain.Submission{}, errors.Wrap(err, "tx by hash")
	}

	sub, err := xchain.DecodeXSubmit(tx.Data())
	if err != nil {
		return xchain.Submission{}, errors.Wrap(err, "decode xsubmit")
	}

	return xchain.SubmissionFromBinding(sub, chain.ID)
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

	// Fetch the header from the ethclient
	header, err := headerByConfLevel(ctx, rpcClient, chainVer.ConfLevel)
	if err != nil {
		return nil, errors.Wrap(err, "header by conf level")
	}

	// Update the head cache
	p.mu.Lock()
	defer p.mu.Unlock()
	p.confHeads[chainVer] = header.Number.Uint64()

	return header, nil
}

func parseXMsg(filterer *bindings.OmniPortalFilterer, event types.Log, chainID uint64) (xchain.Msg, error) {
	e, err := filterer.ParseXMsg(event)
	if err != nil {
		return xchain.Msg{}, errors.Wrap(err, "parse xmsg log")
	}

	return xchain.Msg{
		MsgID: xchain.MsgID{
			StreamID: xchain.StreamID{
				SourceChainID: chainID,
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
		Fees:            e.Fees,
		LogIndex:        uint64(e.Raw.Index),
	}, nil
}

func parseXReceipt(filterer *bindings.OmniPortalFilterer, event types.Log, chainID uint64) (xchain.Receipt, error) {
	e, err := filterer.ParseXReceipt(event)
	if err != nil {
		return xchain.Receipt{}, errors.Wrap(err, "parse xreceipt log")
	}

	return xchain.Receipt{
		MsgID: xchain.MsgID{
			StreamID: xchain.StreamID{
				SourceChainID: e.SourceChainId,
				DestChainID:   chainID,
				ShardID:       xchain.ShardID(e.ShardId),
			},
			StreamOffset: e.Offset,
		},
		GasUsed:        e.GasUsed.Uint64(),
		Success:        e.Success,
		Error:          e.Err,
		RelayerAddress: e.Relayer,
		TxHash:         e.Raw.TxHash,
		LogIndex:       uint64(e.Raw.Index),
	}, nil
}

func getConsXBlock(ctx context.Context, ref xchain.Ref, cprov cchain.Provider) (xchain.Block, error) {
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
	} else if len(xblock.Msgs) == 0 {
		return xchain.Block{}, errors.New("empty consensus xblock [BUG]")
	} else if xblock.Msgs[0].DestChainID != 0 {
		return xchain.Block{}, errors.New("non-broadcast consensus chain xmsg [BUG]")
	}

	return xblock, nil
}

func spanName(method string) string {
	return "xprovider/" + method
}

// getEventLogs returns the logs for the contract address and block hash with any of the provided topics in the first position.
func getEventLogs(ctx context.Context, rpcClient ethclient.Client, contractAddrs []common.Address, blockHash common.Hash, topics []common.Hash) ([]types.Log, error) {
	logs, err := rpcClient.FilterLogs(ctx, ethereum.FilterQuery{
		BlockHash: &blockHash,
		Addresses: contractAddrs,
		Topics:    [][]common.Hash{topics}, // Match any of the topics in the first position.
	})
	if err != nil {
		return nil, errors.Wrap(err, "filter logs", "block_hash", blockHash)
	}

	// Ensure events are valid and sorted by index.
	for i, log := range logs {
		if i > 0 && log.Index <= logs[i-1].Index {
			return nil, errors.New("unordered log index", "index", i)
		} else if log.BlockHash != blockHash {
			return nil, errors.New("unexpected log block hash", "index", i)
		} else if len(log.Topics) == 0 {
			return nil, errors.New("missing log topics", "index", i)
		}
	}

	return logs, nil
}

// heightForRef returns block height from a valid ref.
// It uses the provided client to fetch on-chain header.
func heightForRef(ctx context.Context, client ethclient.Client, ref xchain.Ref) (*big.Int, error) {
	if !ref.Valid() {
		return nil, errors.New("invalid ref")
	}

	if ref.Height != nil {
		return bi.N(*ref.Height), nil
	}

	height, err := heightByConfLevel(ctx, client, *ref.ConfLevel)
	if err != nil {
		return nil, err
	}

	return bi.N(height), nil
}
