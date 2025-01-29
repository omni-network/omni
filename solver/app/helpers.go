package app

import (
	"bytes"
	"context"

	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
)

// detectContractChains returns the chains on which the contract is deployed at the provided address.
func detectContractChains(ctx context.Context, network netconf.Network, backends ethbackend.Backends, address common.Address) ([]uint64, error) {
	var resp []uint64
	for _, chain := range network.EVMChains() {
		backend, err := backends.Backend(chain.ID)
		if err != nil {
			return nil, err
		}

		code, err := backend.CodeAt(ctx, address, nil)
		if err != nil {
			return nil, errors.Wrap(err, "get code", "chain", chain.Name)
		} else if len(code) == 0 {
			continue
		}

		resp = append(resp, chain.ID)
	}

	return resp, nil
}

// toEthAddr converts a 32-byte address to an Ethereum address.
func toEthAddr(bz [32]byte) common.Address {
	return cast.MustEthAddress(bz[12:])
}

// cmpAddrs returns true if the eth address is equal to the 32-byte address.
func cmpAddrs(addr common.Address, bz [32]byte) bool {
	addrBz := addr.Bytes()
	var addrBz32 [32]byte
	copy(addrBz32[12:], addrBz)

	return bytes.Equal(addrBz32[:], bz[:])
}

// toBz32 converts an Ethereum address to a 32-byte address.
func toBz32(addr common.Address) [32]byte {
	var bz [32]byte
	copy(bz[12:], addr.Bytes())

	return bz
}
