package provider

import (
	"context"
	"math/big"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const xMsgSigString = "XMsg(uint64,uint64,address,address,bytes,uint64)"

//nolint:gochecknoglobals // This is a static hash.
var xMsgSigHash = crypto.Keccak256Hash([]byte(xMsgSigString))

func getCurrentFinalisedBlockHeader(ctx context.Context, rpcClient *ethclient.Client) (*types.Header, error) {
	// skip ethCLient and call the function directly as the "finalized" tag is not supported
	// by ethClient. This call will return the last finalized block.
	// var finalisedHeader types.Header
	// params := []string{"latest", "false"}
	// err := rpcClient.Client().CallContext(ctx, &finalisedHeader, "eth_getBlockByNumber", params)
	// if err != nil {
	//	 return nil, errors.Wrap(err, "could not get finalized block")
	// }

	// TODO(corver): Support different finalized methods (to be added to netconf).
	//  The only chain we support at this point is anvil, it doesn't support "finalized", so just use "latest" for now.
	height, err := rpcClient.BlockNumber(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get block number")
	}

	header, err := rpcClient.HeaderByNumber(ctx, big.NewInt(int64(height)))
	if err != nil {
		return nil, errors.Wrap(err, "get header by number")
	}

	return header, nil
}

func (*Provider) GetSubmittedCursor(context.Context, uint64, uint64) (xchain.StreamCursor, error) {
	return xchain.StreamCursor{}, errors.New("not implemented")
}

// GetBlock returns the XBlock for the provided chain and height, or false if not available yet (not finalized),
// or an error.
func (p *Provider) GetBlock(ctx context.Context, chainID uint64, height uint64) (xchain.Block, bool, error) {
	chain, rpcClient, err := p.getChain(chainID)
	if err != nil {
		return xchain.Block{}, false, err
	}

	// get the current finalized header
	finalisedHeader, err := getCurrentFinalisedBlockHeader(ctx, rpcClient)
	if err != nil {
		return xchain.Block{}, false, err
	}

	// ignore if our height is greater than the finalized height
	if height > finalisedHeader.Number.Uint64() {
		return xchain.Block{}, false, nil
	}

	// construct the query to fetch all the event logs in the given height
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(height)),
		ToBlock:   big.NewInt(int64(height)),
		Addresses: []common.Address{
			common.HexToAddress(chain.PortalAddress),
		},
	}

	// call the rpc to get the logs from the chain
	logs, err := rpcClient.FilterLogs(ctx, query)
	if err != nil {
		return xchain.Block{}, false, errors.Wrap(err, "could not filter logs")
	}

	// select the logs based on the required event signature
	// TODO(jmozah): extract receipts logs too
	selectedMsgLogs := make([]types.Log, 0)
	for _, vLog := range logs {
		switch vLog.Topics[0].Hex() {
		case xMsgSigHash.Hex():
			selectedMsgLogs = append(selectedMsgLogs, vLog)
		default:
			return xchain.Block{}, false, errors.New("log not expected")
		}
	}

	// check if we can reuse the header
	if height != finalisedHeader.Number.Uint64() {
		// fetch the block header for the given height
		hdr, err := rpcClient.HeaderByNumber(ctx, big.NewInt(int64(height)))
		if err != nil {
			return xchain.Block{}, false, errors.Wrap(err, "could not get header by number")
		}
		finalisedHeader = hdr
	}

	return constructXBlock(chain.ID, selectedMsgLogs, finalisedHeader), true, nil
}

// constructXBlock assembles the xBlock using the XMsgs and XReceipts found in the given block height.
func constructXBlock(chainID uint64, selectedMsgLogs []types.Log, header *types.Header) xchain.Block {
	// assemble the block header and skeleton
	xBlock := xchain.Block{
		BlockHeader: xchain.BlockHeader{
			SourceChainID: chainID,
			BlockHeight:   header.Number.Uint64(),
			BlockHash:     header.Hash(),
		},
		Timestamp: time.Unix(int64(header.Time), 0),
	}

	// add cross chain message and receipts to the block
	for _, vLog := range selectedMsgLogs {
		// construct the messages that go in this block
		streamID := xchain.StreamID{
			SourceChainID: chainID,
			DestChainID:   vLog.Topics[1].Big().Uint64(),
		}
		msgID := xchain.MsgID{
			StreamID:     streamID,
			StreamOffset: vLog.Topics[2].Big().Uint64(),
		}
		msg := xchain.Msg{
			MsgID:           msgID,
			SourceMsgSender: [20]byte(vLog.Topics[2].Bytes()),
			DestAddress:     [20]byte(vLog.Topics[3].Bytes()),
			Data:            vLog.Topics[4].Bytes(),
			DestGasLimit:    vLog.Topics[5].Big().Uint64(),
			TxHash:          vLog.TxHash,
		}
		xBlock.Msgs = append(xBlock.Msgs, msg)
	}

	return xBlock
}
