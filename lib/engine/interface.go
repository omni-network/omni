// Package engine defines the execution chain' Engine API interface and RPC client implementation.
// See reference https://github.com/ethereum/execution-apis/tree/main/src/engine.
package engine

import (
	"context"

	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
)

// API defines the execution chain' Engine API.
type API interface {
	// NewPayloadV2 creates an Eth1 block, inserts it in the chain, and returns the status of the chain.
	NewPayloadV2(ctx context.Context, params engine.ExecutableData) (engine.PayloadStatusV1, error)
	// NewPayloadV3 creates an Eth1 block, inserts it in the chain, and returns the status of the chain.
	NewPayloadV3(ctx context.Context, params engine.ExecutableData, versionedHashes []common.Hash,
		beaconRoot *common.Hash) (engine.PayloadStatusV1, error)

	// ForkchoiceUpdatedV2 has several responsibilities:
	//  - It sets the chain the head.
	//  - And/or it sets the chain's finalized block hash.
	//  - And/or it starts assembling (async) a block with the payload attributes.
	ForkchoiceUpdatedV2(ctx context.Context, update engine.ForkchoiceStateV1,
		payloadAttributes *engine.PayloadAttributes) (engine.ForkChoiceResponse, error)

	// ForkchoiceUpdatedV3 is equivalent to V2 with the addition of parent beacon block root in the payload attributes.
	ForkchoiceUpdatedV3(ctx context.Context, update engine.ForkchoiceStateV1,
		payloadAttributes *engine.PayloadAttributes) (engine.ForkChoiceResponse, error)

	// GetPayloadV2 returns a cached payload by id.
	GetPayloadV2(ctx context.Context, payloadID engine.PayloadID) (*engine.ExecutionPayloadEnvelope, error)
	// GetPayloadV3 returns a cached payload by id.
	GetPayloadV3(ctx context.Context, payloadID engine.PayloadID) (*engine.ExecutionPayloadEnvelope, error)
}
