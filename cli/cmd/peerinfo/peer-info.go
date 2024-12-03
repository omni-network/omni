package peerinfo

import (
	_ "embed"
	"encoding/json"

	cmtjson "github.com/cometbft/cometbft/libs/json"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	"github.com/cometbft/cometbft/rpc/jsonrpc/types"

	"github.com/omni-network/omni/lib/errors"
)

//go:embed omega-net-info.json
var omegaNetInfo []byte

//go:embed mainnet-net-info.json
var mainnetNetInfo []byte

// OmegaNetInfo returns omega net info.
func OmegaNetInfo() (*coretypes.ResultNetInfo, error) {
	var omega coretypes.ResultNetInfo
	omegaResp := &types.RPCResponse{}
	if err := json.Unmarshal(omegaNetInfo, omegaResp); err != nil {
		return nil, errors.New("Failed to Unmarshal omega-net-info.json: %v", err)
	}

	if err := cmtjson.Unmarshal(omegaResp.Result, &omega); err != nil {
		return nil, errors.New("failed to unmarshal omega response: %v", err)
	}

	return &omega, nil
}

// MainnetNetInfo returns mainnet net info.
func MainnetNetInfo() (*coretypes.ResultNetInfo, error) {
	var mainnet coretypes.ResultNetInfo
	mainnetResp := &types.RPCResponse{}
	if err := json.Unmarshal(mainnetNetInfo, mainnetResp); err != nil {
		return nil, errors.New("Failed to Unmarshal mainnet-net-info.json: %v", err)
	}

	if err := cmtjson.Unmarshal(mainnetResp.Result, &mainnet); err != nil {
		return nil, errors.New("failed to unmarshal mainnet response: %v", err)
	}

	return &mainnet, nil
}
