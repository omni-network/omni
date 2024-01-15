package ethclient

import (
	"context"
	"math/big"
	"strings"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
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

// GetBlock fetches the cross chain block, if present in a given rollup block height.
func (e *EthClient) GetBlock(ctx context.Context, height uint64) (xchain.Block, bool, error) {
	var xBlock xchain.Block

	// TODO(jmozah): check if block is finalized else return
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
		}
	}

	// construct an xblock only if some cross chain events are found
	if len(selectedMsgLogs) > 0 {
		xBlock = e.constructXBlock(selectedMsgLogs)
	}

	return xBlock, true, nil
}

// constructXBlocks assembles the xBlock using the XMsgs and XReceipts found in the block height.
func (e *EthClient) constructXBlock(selectedMsgLogs []types.Log) xchain.Block {
	var header xchain.BlockHeader
	messages := make([]xchain.Msg, 0)
	var block xchain.Block

	// construct the block based on cross chain message or receipts that are found
	for _, vLog := range selectedMsgLogs {
		// create the BlockHeader once
		if header.SourceChainID != e.chainID {
			header = xchain.BlockHeader{
				SourceChainID: e.chainID,
				BlockHeight:   vLog.BlockNumber,
				BlockHash:     vLog.BlockHash,
			}
		}

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
		messages = append(messages, msg)
	}

	// assemble the entire block
	block = xchain.Block{
		BlockHeader: header,
		Msgs:        messages,
	}

	return block
}
