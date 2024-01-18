package ethclient

import (
	"context"
	"math/big"
	"strings"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	xMsgSigString = "XMsg(uint64,uint64,address,address,bytes,uint64)"
)

// EthClient is the configuration for the rpc client to connect and get information.
type EthClient struct {
	chainID       uint64
	portalAddress common.Address
	contractABI   abi.ABI
	xMsgSigHash   common.Hash
	rpcClient     *ethclient.Client
}

// NewEthClient is the client implementation of the json rpc interface to the rollup chain.
func NewEthClient(
	chainID uint64,
	portalAddress common.Address,
	rpcClient *ethclient.Client,
) (*EthClient, error) {
	// TODO(jmozah): validate chainID , portalAddress etc

	// construct the omni portal contract ABI from the bindings
	contractAbi, err := abi.JSON(strings.NewReader(bindings.OmniPortalMetaData.ABI))
	if err != nil {
		return nil, errors.Wrap(err, "could not create contract abi")
	}

	// create the method signatures
	xMsgSig := []byte(xMsgSigString)

	return &EthClient{
		chainID:       chainID,
		portalAddress: portalAddress,
		contractABI:   contractAbi,
		xMsgSigHash:   crypto.Keccak256Hash(xMsgSig),
		rpcClient:     rpcClient,
	}, nil
}

func (e *EthClient) getCurrentFinalisedBlockHeader(ctx context.Context) (*types.Header, error) {
	// skip ethCLient and call the function directly as the "finalized" tag is not supported
	// by ethClient. This call will return the last finalized block.
	var finalisedHeader types.Header

	params := []string{"finalized", "false"}
	err := e.rpcClient.Client().CallContext(ctx, &finalisedHeader, "eth_getBlockByNumber", params)
	if err != nil {
		return nil, errors.Wrap(err, "could not get finalized block")
	}

	return &finalisedHeader, nil
}

// GetBlock fetches the cross chain block, if present in a given rollup block height.
func (e *EthClient) GetBlock(ctx context.Context, height uint64) (xchain.Block, bool, error) {
	var xBlock xchain.Block

	// get the current finalized header
	finalisedHeader, err := e.getCurrentFinalisedBlockHeader(ctx)
	if err != nil {
		return xBlock, false, err
	}

	// ignore if our height is greater than the finalized height
	if height > finalisedHeader.Number.Uint64() {
		log.Debug(ctx, "Block not finalized yet",
			"height", height,
			"finalized_height", finalisedHeader.Number.Uint64())

		return xBlock, false, nil
	}

	// construct the query to fetch all the event logs in the given height
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(height)),
		ToBlock:   big.NewInt(int64(height)),
		Addresses: []common.Address{
			e.portalAddress,
		},
	}

	// call the rpc to get the logs from the chain
	logs, err := e.rpcClient.FilterLogs(ctx, query)
	if err != nil {
		return xBlock, false, errors.Wrap(err, "could not filter logs")
	}

	// select the logs based on the required event signature
	// TODO(jmozah): extract receipts logs too
	selectedMsgLogs := make([]types.Log, 0)
	for _, vLog := range logs {
		switch vLog.Topics[0].Hex() {
		case e.xMsgSigHash.Hex():
			selectedMsgLogs = append(selectedMsgLogs, vLog)
		default:
			return xBlock, false, errors.New("log not expected")
		}
	}

	// check if we can reuse the header
	if height != finalisedHeader.Number.Uint64() {
		// fetch the block header for the given height
		hdr, err := e.rpcClient.HeaderByNumber(ctx, big.NewInt(int64(height)))
		if err != nil {
			return xBlock, false, errors.Wrap(err, "could not get header by number")
		}
		finalisedHeader = hdr
	}
	xBlock = e.constructXBlock(selectedMsgLogs, finalisedHeader)

	return xBlock, true, nil
}

// constructXBlock assembles the xBlock using the XMsgs and XReceipts found in the given block height.
func (e *EthClient) constructXBlock(selectedMsgLogs []types.Log, header *types.Header) xchain.Block {
	// assemble the block header and skeleton
	xBlock := xchain.Block{
		BlockHeader: xchain.BlockHeader{
			SourceChainID: e.chainID,
			BlockHeight:   header.Number.Uint64(),
			BlockHash:     header.Hash(),
		},
		Timestamp: time.Unix(int64(header.Time), 0),
	}

	// add cross chain message and receipts to the block
	for _, vLog := range selectedMsgLogs {
		// construct the messages that go in this block
		streamID := xchain.StreamID{
			SourceChainID: e.chainID,
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
