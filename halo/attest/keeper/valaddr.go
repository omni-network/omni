package keeper

import (
	"context"
	"sync"

	vtypes "github.com/omni-network/omni/halo/valsync/types"
	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/errors"

	"github.com/cometbft/cometbft/crypto"

	"github.com/ethereum/go-ethereum/common"
)

// valAddrCache is a simple read-through cache for validator comet-to-eth address lookups.
type valAddrCache struct {
	sync.RWMutex
	ethAddrs map[[crypto.AddressSize]byte]common.Address
}

func (c *valAddrCache) GetEthAddress(cmtAddr [crypto.AddressSize]byte) (common.Address, bool) {
	c.RLock()
	defer c.RUnlock()

	ethAddr, ok := c.ethAddrs[cmtAddr]

	return ethAddr, ok
}

func (c *valAddrCache) SetAll(vals []vtypes.Validator) error {
	var ethAddrs = make(map[[crypto.AddressSize]byte]common.Address, len(vals))
	for _, val := range vals {
		cmtAddr, err := val.CometAddress()
		if err != nil {
			return err
		} else if len(cmtAddr) != crypto.AddressSize {
			return errors.New("invalid comet address length [BUG]", "len", len(cmtAddr))
		}

		cmtAddrArr, err := cast.Array20(cmtAddr)
		if err != nil {
			return err
		}

		ethAddr, err := val.EthereumAddress()
		if err != nil {
			return err
		}

		ethAddrs[cmtAddrArr] = ethAddr
	}

	c.Lock()
	defer c.Unlock()

	c.ethAddrs = ethAddrs

	return nil
}

// getValEthAddr returns the validator's ethereum address via reverse lookup using the provided validator cometBFT address.
// It uses the validator read-through-cache to avoid querying the underlying validator set provider.
// Note it assumes the provided validator is inside the provided set. It doesn't ensure this.
func (k *Keeper) getValEthAddr(ctx context.Context, cmtAddr []byte, setHeight uint64) (common.Address, error) {
	addr, err := cast.Array20(cmtAddr)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "invalid comet address length [BUG]")
	}

	// Check cache
	if ethAddr, ok := k.valAddrCache.GetEthAddress(addr); ok {
		return ethAddr, nil
	}

	// Cache is stale, rehydrate it.
	valset, err := k.valProvider.ActiveSetByHeight(ctx, setHeight)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "get active set")
	}
	if err := k.valAddrCache.SetAll(valset.Validators); err != nil {
		return common.Address{}, err
	}

	// Check the rehydrated cache
	ethAddr, ok := k.valAddrCache.GetEthAddress(addr)
	if !ok {
		return common.Address{}, errors.New("validator not in current set [BUG]")
	}

	return ethAddr, nil
}
