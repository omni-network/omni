package comet

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/tracer"

	lightprovider "github.com/cometbft/cometbft/light/provider"
	lighthttp "github.com/cometbft/cometbft/light/provider/http"
	rpcclient "github.com/cometbft/cometbft/rpc/client"
	cmttypes "github.com/cometbft/cometbft/types"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

var _ API = adapter{}

type API interface {
	// Validators returns the cometBFT validators at the given height.
	Validators(ctx context.Context, height int64) (*cmttypes.ValidatorSet, error)
}

func NewAPI(cl rpcclient.Client, chainID string) API {
	return adapter{
		cl: lighthttp.NewWithClient(chainID, remoteCl{cl}),
	}
}

type adapter struct {
	cl lightprovider.Provider
}

// Validators returns the cometBFT validators at the given height.
func (a adapter) Validators(ctx context.Context, height int64) (*cmttypes.ValidatorSet, error) {
	ctx, span := tracer.Start(ctx, "comet/validators", trace.WithAttributes(attribute.Int64("height", height)))
	defer span.End()

	block, err := a.cl.LightBlock(ctx, height) // LightBlock does all the heavy lifting to query the validator set.
	if err != nil {
		return nil, errors.Wrap(err, "fetch light block")
	}

	return block.ValidatorSet, nil
}

// remoteCl is a wrapper around rpcclient.Client to implement rpcclient.RemoteClient.
type remoteCl struct {
	rpcclient.Client
}

func (remoteCl) Remote() string {
	return ""
}
