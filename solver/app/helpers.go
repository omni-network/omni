package app

import (
	"context"
	"encoding/hex"

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

// fmtReqID returns the least-significant 7 hex chars of the provided request ID.
// ReqIDs are monotonically incrementing numbers, not hashes.
func fmtReqID(reqID [32]byte) string {
	return hex.EncodeToString(reqID[:])[:7]
}
