package ethclient

import (
	"context"
	"net/http"
	"time"

	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

const (
	defaultRPCHTTPTimeout = time.Second * 30

	newPayloadV2 = "engine_newPayloadV2"
	newPayloadV3 = "engine_newPayloadV3"

	forkchoiceUpdatedV2 = "engine_forkchoiceUpdatedV2"
	forkchoiceUpdatedV3 = "engine_forkchoiceUpdatedV3"

	getPayloadV2 = "engine_getPayloadV2"
	getPayloadV3 = "engine_getPayloadV3"
)

// EngineClient defines the Engine API authenticated JSON-RPC endpoints.
// It extends the normal Client interface with the Engine API.
type EngineClient interface {
	Client

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

// engineClient implements EngineClient using JSON-RPC.
type engineClient struct {
	Wrapper
}

// NewAuthClient returns a new authenticated JSON-RPc engineClient.
func NewAuthClient(ctx context.Context, urlAddr string, jwtSecret []byte) (EngineClient, error) {
	transport := http.DefaultTransport
	if len(jwtSecret) > 0 {
		transport = newJWTRoundTripper(http.DefaultTransport, jwtSecret)
	}

	client := &http.Client{Timeout: defaultRPCHTTPTimeout, Transport: transport}

	rpcClient, err := rpc.DialOptions(ctx, urlAddr, rpc.WithHTTPClient(client))
	if err != nil {
		return engineClient{}, errors.Wrap(err, "rpc dial")
	}

	return engineClient{
		Wrapper: NewClient(rpcClient, "engine", urlAddr),
	}, nil
}

func (c engineClient) NewPayloadV2(ctx context.Context, params engine.ExecutableData) (engine.PayloadStatusV1, error) {
	const endpoint = "new_payload_v2"
	defer latency(c.chain, endpoint)()

	var resp engine.PayloadStatusV1
	err := c.cl.Client().CallContext(ctx, &resp, newPayloadV2, params)
	if err != nil {
		incError(c.chain, endpoint)
		return engine.PayloadStatusV1{}, errors.Wrap(err, "rpc new payload v2")
	}

	return resp, nil
}

func (c engineClient) NewPayloadV3(ctx context.Context, params engine.ExecutableData, versionedHashes []common.Hash,
	beaconRoot *common.Hash,
) (engine.PayloadStatusV1, error) {
	const endpoint = "new_payload_v3"
	defer latency(c.chain, endpoint)()

	var resp engine.PayloadStatusV1
	err := c.cl.Client().CallContext(ctx, &resp, newPayloadV3, params, versionedHashes, beaconRoot)
	if err != nil {
		incError(c.chain, endpoint)
		return engine.PayloadStatusV1{}, errors.Wrap(err, "rpc new payload v3")
	}

	return resp, nil
}

func (c engineClient) ForkchoiceUpdatedV2(ctx context.Context, update engine.ForkchoiceStateV1,
	payloadAttributes *engine.PayloadAttributes,
) (engine.ForkChoiceResponse, error) {
	const endpoint = "forkchoice_updated_v2"
	defer latency(c.chain, endpoint)()

	var resp engine.ForkChoiceResponse
	err := c.cl.Client().CallContext(ctx, &resp, forkchoiceUpdatedV2, update, payloadAttributes)
	if err != nil {
		incError(c.chain, endpoint)
		return engine.ForkChoiceResponse{}, errors.Wrap(err, "rpc forkchoice updated v2")
	}

	return resp, nil
}

func (c engineClient) ForkchoiceUpdatedV3(ctx context.Context, update engine.ForkchoiceStateV1,
	payloadAttributes *engine.PayloadAttributes,
) (engine.ForkChoiceResponse, error) {
	const endpoint = "forkchoice_updated_v3"
	defer latency(c.chain, endpoint)()

	var resp engine.ForkChoiceResponse
	err := c.cl.Client().CallContext(ctx, &resp, forkchoiceUpdatedV3, update, payloadAttributes)
	if err != nil {
		incError(c.chain, endpoint)
		return engine.ForkChoiceResponse{}, errors.Wrap(err, "rpc forkchoice updated v3")
	}

	return resp, nil
}

func (c engineClient) GetPayloadV2(ctx context.Context, payloadID engine.PayloadID) (
	*engine.ExecutionPayloadEnvelope, error,
) {
	const endpoint = "get_payload_v2"
	defer latency(c.chain, endpoint)()

	var resp engine.ExecutionPayloadEnvelope
	err := c.cl.Client().CallContext(ctx, &resp, getPayloadV2, payloadID)
	if err != nil {
		incError(c.chain, endpoint)
		return nil, errors.Wrap(err, "rpc get payload v2")
	}

	return &resp, nil
}

func (c engineClient) GetPayloadV3(ctx context.Context, payloadID engine.PayloadID) (
	*engine.ExecutionPayloadEnvelope, error,
) {
	const endpoint = "get_payload_v3"
	defer latency(c.chain, endpoint)()

	var resp engine.ExecutionPayloadEnvelope
	err := c.cl.Client().CallContext(ctx, &resp, getPayloadV3, payloadID)
	if err != nil {
		incError(c.chain, endpoint)
		return nil, errors.Wrap(err, "rpc get payload v3")
	}

	return &resp, nil
}
