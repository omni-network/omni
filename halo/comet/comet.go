package comet

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/tracer"

	rpcclient "github.com/cometbft/cometbft/rpc/client"
	cmttypes "github.com/cometbft/cometbft/types"

	"github.com/ethereum/go-ethereum/common"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

const perPageConst = 100

var _ API = adapter{}

type API interface {
	// Validators returns the cometBFT validators at the given height or false if not
	// available (probably due to snapshot sync after height).
	Validators(ctx context.Context, height int64) (*cmttypes.ValidatorSet, bool, error)

	// IsValidator returns true if the given address is a validator at the latest height.
	// It is best-effort, so returns false on any error.
	IsValidator(ctx context.Context, valAddress common.Address) bool
}

func NewAPI(cl rpcclient.Client) API {
	return adapter{cl: cl}
}

type adapter struct {
	cl rpcclient.Client
}

// IsValidator returns true if the given address is a validator at the latest height.
// It is best-effort, so returns false on any error.
func (a adapter) IsValidator(ctx context.Context, valAddress common.Address) bool {
	ctx, span := tracer.Start(ctx, "comet/is_validator")
	defer span.End()

	status, err := a.cl.Status(ctx)
	if err != nil || status.SyncInfo.CatchingUp {
		return false // Best effort
	}

	valset, ok, err := a.Validators(ctx, status.SyncInfo.LatestBlockHeight)
	if !ok || err != nil {
		return false // Best effort
	}

	for _, val := range valset.Validators {
		addr, err := k1util.PubKeyToAddress(val.PubKey)
		if err != nil {
			continue // Best effort
		}

		if addr == valAddress {
			return true
		}
	}

	return false
}

// Validators returns the cometBFT validators at the given height or false if not
// available (probably due to snapshot sync after height).
func (a adapter) Validators(ctx context.Context, height int64) (*cmttypes.ValidatorSet, bool, error) {
	ctx, span := tracer.Start(ctx, "comet/validators", trace.WithAttributes(attribute.Int64("height", height)))
	defer span.End()

	perPage := perPageConst // Can't take a pointer to a const directly.

	var vals []*cmttypes.Validator
	for page := 1; ; page++ { // Pages are 1-indexed.
		if page > 10 { // Sanity check.
			return nil, false, errors.New("too many validators [BUG]")
		}

		status, err := a.cl.Status(ctx)
		if err != nil {
			return nil, false, errors.Wrap(err, "fetch status")
		} else if height < status.SyncInfo.EarliestBlockHeight {
			// This can happen if height is before snapshot restore.
			return nil, false, nil
		}

		valResp, err := a.cl.Validators(ctx, &height, &page, &perPage)
		if err != nil {
			return nil, false, errors.Wrap(err, "fetch validators")
		}

		for _, v := range valResp.Validators {
			vals = append(vals, cmttypes.NewValidator(v.PubKey, v.VotingPower))
		}

		if len(vals) == valResp.Total {
			break
		}
	}

	// cmttypes.NewValidatorSet() panics on error, so manually construct it for proper error handling.
	valset := new(cmttypes.ValidatorSet)
	if err := valset.UpdateWithChangeSet(vals); err != nil {
		return nil, false, errors.Wrap(err, "update with change set")
	}
	if len(vals) > 0 {
		valset.IncrementProposerPriority(1) // See cmttypes.NewValidatorSet
	}

	if err := valset.ValidateBasic(); err != nil {
		return nil, false, errors.Wrap(err, "validate basic")
	}

	return valset, true, nil
}
