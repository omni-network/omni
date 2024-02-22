package comet

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"

	rpcclient "github.com/cometbft/cometbft/rpc/client"
	cmttypes "github.com/cometbft/cometbft/types"
)

const perPageConst = 100

var _ API = adapter{}

type API interface {
	// Validators returns the cometBFT validators at the given height.
	Validators(ctx context.Context, height int64) (*cmttypes.ValidatorSet, error)
}

func NewAPI(cl rpcclient.Client) API {
	return adapter{cl: cl}
}

type adapter struct {
	cl rpcclient.Client
}

func (a adapter) Validators(ctx context.Context, height int64) (*cmttypes.ValidatorSet, error) {
	perPage := perPageConst // Can't take a pointer to a const directly.

	var vals []*cmttypes.Validator
	for page := 1; ; page++ { // Pages are 1-indexed.
		if page > 10 { // Sanity check.
			return nil, errors.New("too many validators [BUG]")
		}

		valResp, err := a.cl.Validators(ctx, &height, &page, &perPage)
		if err != nil {
			return nil, errors.Wrap(err, "fetch validators")
		}

		for _, v := range valResp.Validators {
			val := cmttypes.NewValidator(v.PubKey, v.VotingPower)

			// Set the address to our Ethereum-style address.
			addr, err := k1util.PubKeyToAddress(val.PubKey)
			if err != nil {
				return nil, errors.Wrap(err, "pubkey to address")
			}
			val.Address = addr[:]

			vals = append(vals, val)
		}

		if len(vals) == valResp.Total {
			break
		}
	}

	// cmttypes.NewValidatorSet() panics on error, so manually construct it for proper error handling.
	valset := new(cmttypes.ValidatorSet)
	if err := valset.UpdateWithChangeSet(vals); err != nil {
		return nil, errors.Wrap(err, "update with change set")
	}
	if len(vals) > 0 {
		valset.IncrementProposerPriority(1) // See cmttypes.NewValidatorSet, not sure about this actually...?
	}
	// Note we can not call valSet.ValidateBasic() since we changed the address format.

	return valset, nil
}
