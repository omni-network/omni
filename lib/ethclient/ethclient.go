package ethclient

import (
	"context"
	"math/big"
	"strings"

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

// EthClient is the configuration for the rpc client to connect and get information.
type EthClient struct {
	chainID       uint64
	rpcURL        string
	portalAddress common.Address
	contractABI   abi.ABI
	xMsgSigHash   common.Hash
	rpcClient     *ethclient.Client
}

// NewEthClient is the client implementation of the json rpc interface to the rollup chain.
func NewEthClient(
	chainID uint64,
	rpcURL string,
	portalAddress common.Address,
) (*EthClient, error) {
	// TODO(jmozah): validate chainID , portalAddress etc

	// construct the omni portal contract ABI from the bindings
	contractAbi, err := abi.JSON(strings.NewReader(bindings.OmniPortalMetaData.ABI))
	if err != nil {
		return nil, errors.Wrap(err, "could not create contract abi")
	}

	// create the method signatures
	xMsgSig := []byte("XMsg(uint64,uint64,address,address,bytes,uint64)")

	return &EthClient{
		chainID:       chainID,
		rpcURL:        rpcURL,
		portalAddress: portalAddress,
		contractABI:   contractAbi,
		xMsgSigHash:   crypto.Keccak256Hash(xMsgSig),
	}, nil
}

// GetBlocks fetches the blocks that are present in a given block height.
func (e *EthClient) GetBlocks(ctx context.Context, height uint64) ([]xchain.Block, error) {
	// connect to rpc if not connected
	err := e.connect(ctx)
	if err != nil {
		return nil, err
	}

	// get the xBlock
	xBlocks, err := e.constructXBlocks(ctx, height)
	if err != nil {
		return nil, err
	}

	return xBlocks, nil
}

// connect initiates a connection if it is not already dialed in.
func (e *EthClient) connect(ctx context.Context) error {
	if e.rpcClient != nil {
		return nil
	}
	eClient, err := ethclient.Dial(e.rpcURL)
	if err != nil {
		return errors.Wrap(err, "could not connect to chain")
	}
	e.rpcClient = eClient
	log.Info(ctx, "Connected to chain. ",
		"chainId", e.chainID,
		"rpcURL", e.rpcURL)

	return nil
}

// constructXBlocks assembles the blocks by parsing the XMsg event for a given height.
func (e *EthClient) constructXBlocks(ctx context.Context, height uint64) ([]xchain.Block, error) {
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
		return nil, errors.Wrap(err, "could not filter logs")
	}

	// select the logs based on the required event signature
	selectedLogs := make(map[uint64][]types.Log)
	for _, vLog := range logs {
		switch vLog.Topics[0].Hex() {
		case e.xMsgSigHash.Hex():
			destChainID := vLog.Topics[1].Big().Uint64()
			if blockLogs, ok := selectedLogs[destChainID]; ok {
				blockLogs = append(blockLogs, vLog)
				selectedLogs[destChainID] = blockLogs
			} else {
				selectedLogs[destChainID] = []types.Log{vLog}
			}
		default:
		}
	}

	// construct the blocks based on the destination chain
	// i.e. one block for each destination chain
	blocks := make([]xchain.Block, 0)
	for destID, vLogs := range selectedLogs {
		// create the BlockHeader
		header := xchain.BlockHeader{
			SourceChainID: e.chainID,
			BlockHeight:   vLogs[0].BlockNumber,
			BlockHash:     vLogs[0].BlockHash,
		}

		// construct the messages that go in this block
		var msgs []xchain.Msg
		for _, msg := range vLogs {
			streamID := xchain.StreamID{
				SourceChainID: e.chainID,
				DestChainID:   destID,
			}
			msgID := xchain.MsgID{
				StreamID:     streamID,
				StreamOffset: msg.Topics[2].Big().Uint64(),
			}
			m := xchain.Msg{
				MsgID:           msgID,
				SourceMsgSender: [20]byte(msg.Topics[2].Bytes()),
				DestAddress:     [20]byte(msg.Topics[3].Bytes()),
				Data:            msg.Topics[4].Bytes(),
				DestGasLimit:    msg.Topics[5].Big().Uint64(),
				TxHash:          msg.TxHash,
			}
			msgs = append(msgs, m)
		}

		// assemble the entire block
		block := xchain.Block{
			BlockHeader: header,
			Msgs:        msgs,
		}
		blocks = append(blocks, block)
	}

	return blocks, nil
}
