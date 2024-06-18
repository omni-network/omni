package keeper

import (
	"context"
	"sync"

	vtypes "github.com/omni-network/omni/halo/valsync/types"
	"github.com/omni-network/omni/lib/errors"

	"github.com/cometbft/cometbft/crypto"

	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// valCache is a simple read-through cache for validator eth-to-comet address lookups.
type valCache struct {
	sync.RWMutex
	valsByAddr map[common.Address]*vtypes.Validator
}

func (c *valCache) GetCometAddress(ethAddr common.Address) (crypto.Address, bool, error) {
	c.RLock()
	defer c.RUnlock()

	val, ok := c.valsByAddr[ethAddr]
	if !ok {
		return nil, false, nil
	}

	cmtAddr, err := val.CometConsensusAddress()
	if err != nil {
		return nil, false, err
	}

	return cmtAddr, true, nil
}

func (c *valCache) SetAll(vals []*vtypes.Validator) error {
	c.Lock()
	defer c.Unlock()

	var valsByAddr = make(map[common.Address]*vtypes.Validator, len(vals))
	for _, val := range vals {
		addr, err := val.EthConsensusAddress()
		if err != nil {
			return err
		}

		valsByAddr[addr] = val
	}

	c.valsByAddr = valsByAddr

	return nil
}

// getValCometAddr returns the validator's cometBFT address via reverse lookup using the provided validator ethereum address.
// It uses the validator read-through-cache to avoid querying the underlying validator set provider.
func (k *Keeper) getValCometAddr(ctx context.Context, ethAddr common.Address) (crypto.Address, error) {
	cmtAddr, ok, err := k.valCache.GetCometAddress(ethAddr)
	if err != nil {
		return nil, err
	} else if ok {
		return cmtAddr, nil
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	valset, err := k.valProvider.ActiveSetByHeight(ctx, uint64(sdkCtx.BlockHeight()))
	if err != nil {
		return nil, errors.Wrap(err, "get active set")
	}

	if err := k.valCache.SetAll(valset.Validators); err != nil {
		return nil, err
	}

	cmtAddr, ok, err = k.valCache.GetCometAddress(ethAddr)
	if err != nil {
		return nil, err
	} else if !ok {
		return nil, errors.New("validator not in current set")
	}

	return cmtAddr, nil
}
