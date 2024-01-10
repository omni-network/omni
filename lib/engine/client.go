package engine

import (
	"context"
	"net/http"
	"time"

	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

const defaultRPCHTTPTimeout = time.Second * 30

var _ API = Client{}

// Client implements the Engine API using JSON-RPC.
type Client struct {
	ethclient.Client
	client *rpc.Client
}

// NewClient returns a new Engine API JSON-RPC client.
func NewClient(ctx context.Context, urlAddr string, jwtSecret []byte) (Client, error) {
	transport := http.DefaultTransport
	if len(jwtSecret) > 0 {
		transport = newJWTRoundTripper(http.DefaultTransport, jwtSecret)
	}

	client := &http.Client{Timeout: defaultRPCHTTPTimeout, Transport: transport}

	rpcClient, err := rpc.DialOptions(ctx, urlAddr, rpc.WithHTTPClient(client))
	if err != nil {
		return Client{}, errors.Wrap(err, "rpc dial")
	}

	return Client{
		Client: *ethclient.NewClient(rpcClient),
		client: rpcClient,
	}, nil
}

func (c Client) NewPayloadV2(ctx context.Context, params engine.ExecutableData) (engine.PayloadStatusV1, error) {
	var resp engine.PayloadStatusV1
	err := c.client.CallContext(ctx, &resp, newPayloadV2, params)
	if err != nil {
		return engine.PayloadStatusV1{}, errors.Wrap(err, "rpc new payload v2")
	}

	return resp, nil
}

func (c Client) NewPayloadV3(ctx context.Context, params engine.ExecutableData, versionedHashes []common.Hash,
	beaconRoot *common.Hash,
) (engine.PayloadStatusV1, error) {
	var resp engine.PayloadStatusV1
	err := c.client.CallContext(ctx, &resp, newPayloadV3, params, versionedHashes, beaconRoot)
	if err != nil {
		return engine.PayloadStatusV1{}, errors.Wrap(err, "rpc new payload v3")
	}

	return resp, nil
}

func (c Client) ForkchoiceUpdatedV2(ctx context.Context, update engine.ForkchoiceStateV1,
	payloadAttributes *engine.PayloadAttributes,
) (engine.ForkChoiceResponse, error) {
	var resp engine.ForkChoiceResponse
	err := c.client.CallContext(ctx, &resp, forkchoiceUpdatedV2, update, payloadAttributes)
	if err != nil {
		return engine.ForkChoiceResponse{}, errors.Wrap(err, "rpc forkchoice updated v2")
	}

	return resp, nil
}

func (c Client) ForkchoiceUpdatedV3(ctx context.Context, update engine.ForkchoiceStateV1,
	payloadAttributes *engine.PayloadAttributes,
) (engine.ForkChoiceResponse, error) {
	var resp engine.ForkChoiceResponse
	err := c.client.CallContext(ctx, &resp, forkchoiceUpdatedV3, update, payloadAttributes)
	if err != nil {
		return engine.ForkChoiceResponse{}, errors.Wrap(err, "rpc forkchoice updated v3")
	}

	return resp, nil
}

func (c Client) GetPayloadV2(ctx context.Context, payloadID engine.PayloadID) (
	*engine.ExecutionPayloadEnvelope, error,
) {
	var resp engine.ExecutionPayloadEnvelope
	err := c.client.CallContext(ctx, &resp, getPayloadV2, payloadID)
	if err != nil {
		return nil, errors.Wrap(err, "rpc get payload v2")
	}

	return &resp, nil
}

func (c Client) GetPayloadV3(ctx context.Context, payloadID engine.PayloadID) (
	*engine.ExecutionPayloadEnvelope, error,
) {
	var resp engine.ExecutionPayloadEnvelope
	err := c.client.CallContext(ctx, &resp, getPayloadV3, payloadID)
	if err != nil {
		return nil, errors.Wrap(err, "rpc get payload v3")
	}

	return &resp, nil
}
