package app

import (
	"context"
	"encoding/json"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	"math/rand/v2"
)

var (
	portalABI = mustGetABI(bindings.OmniPortalMetaData)
	xMsgEvent = mustGetEvent(portalABI, "XMsg")
)

func parseAndFuzzXMsgs(ctx context.Context, height uint64, respBody []byte) ([]byte, bool, error) {
	var respMsg jsonRPCMessage
	if err := json.Unmarshal(respBody, &respMsg); err != nil {
		return nil, false, errors.Wrap(err, "unmarshal response")
	}

	var logs []ethtypes.Log
	if err := json.Unmarshal(respMsg.Result, &logs); err != nil {
		return nil, false, errors.Wrap(err, "unmarshal result")
	}

	if len(logs) == 0 {
		return respBody, false, nil
	}

	filterer, err := bindings.NewOmniPortalFilterer(common.Address{}, nil)
	if err != nil {
		return nil, false, err
	}

	var msgs []*bindings.OmniPortalXMsg
	for _, eventLog := range logs {
		msg, err := filterer.ParseXMsg(eventLog)
		if err != nil {
			return nil, false, errors.Wrap(err, "parse xmsg log")
		}

		msgs = append(msgs, msg)
	}

	var fuzzedLogs []ethtypes.Log
	for i, msg := range fuzzXMsgs(ctx, height, msgs) {
		data, err := portalABI.Events["XMsg"].Inputs.NonIndexed().Pack(msg.Sender, msg.To, msg.Data, msg.GasLimit, msg.Fees)
		if err != nil {
			return nil, false, errors.Wrap(err, "pack xmsg")
		}
		eventLog := logs[i]
		eventLog.Data = data
		fuzzedLogs = append(fuzzedLogs, eventLog)
	}

	result, err := json.Marshal(fuzzedLogs)
	if err != nil {
		return nil, false, errors.Wrap(err, "marshal result")
	}

	respMsg.Result = result

	bz, err := json.Marshal(respMsg)
	if err != nil {
		return nil, false, errors.Wrap(err, "marshal response")
	}

	return bz, true, nil
}

func fuzzXMsgs(ctx context.Context, height uint64, msgs []*bindings.OmniPortalXMsg) []*bindings.OmniPortalXMsg {
	switch height % 2 {
	case 0: // Every even block.
		// Change 50% of 1st message gas limits.
		// Consensus would not be possible.
		do := rand.Float64() < 0.5
		if do {
			msgs[0].GasLimit++
		}
		log.Info(ctx, "Fuzzed inconsistent attestation root", "height", height, "did", do)
	case 1: // Every odd block.
		// Remove all msgs.
		// Will reach consensus, but results in BlockOffset mismatch of subsequent xblock with Finalized.
		msgs = nil
		log.Info(ctx, "Fuzzed empty xblock", "height", height)
	}

	return msgs
}

func isFuzzyXMsgLogFilter(ctx context.Context, target string, reqMsg jsonRPCMessage) (bool, uint64, error) {
	if reqMsg.Method != "eth_getLogs" {
		return false, 0, nil
	}

	var args []struct {
		Topics    [][]common.Hash `json:"topics"`
		FromBlock string          `json:"fromBlock"`
		ToBlock   string          `json:"toBlock"`
	}
	if err := json.Unmarshal(reqMsg.Params, &args); err != nil {
		return false, 0, errors.Wrap(err, "unmarshal params")
	}
	if len(args) != 1 {
		return false, 0, nil
	}

	arg := args[0]

	if len(arg.Topics) == 0 || len(arg.Topics[0]) != 1 {
		return false, 0, nil
	} else if arg.Topics[0][0] != xMsgEvent.ID {
		return false, 0, nil
	}

	if arg.FromBlock != arg.ToBlock {
		return false, 0, errors.New("fromBlock and toBlock must be equal")
	}

	block, err := hexutil.DecodeUint64(arg.FromBlock)
	if err != nil {
		return false, 0, errors.Wrap(err, "decode block number")
	}

	ethCl, err := ethclient.Dial("proxy", target)
	if err != nil {
		return false, 0, errors.Wrap(err, "dial ethclient")
	}

	finalized, err := ethCl.HeaderByType(ctx, ethclient.HeadFinalized)
	if err != nil {
		return false, 0, errors.Wrap(err, "get finalized header")
	}

	const buffer = 2 // Avoid race conditions, require query to be more than 2 ahead of finalized.

	return block-buffer > finalized.Number.Uint64(), block, nil
}

// mustGetABI returns the metadata's ABI as an abi.ABI type.
// It panics on error.
func mustGetABI(metadata *bind.MetaData) *abi.ABI {
	abi, err := metadata.GetAbi()
	if err != nil {
		panic(err)
	}

	return abi
}

// mustGetEvent returns the event with the given name from the ABI.
// It panics if the event is not found.
func mustGetEvent(abi *abi.ABI, name string) abi.Event {
	event, ok := abi.Events[name]
	if !ok {
		panic("event not found")
	}

	return event
}
