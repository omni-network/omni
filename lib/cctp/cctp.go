package cctp

import (
	"context"

	"github.com/omni-network/omni/lib/cctp/db"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
	xprovider "github.com/omni-network/omni/lib/xchain/provider"

	"github.com/ethereum/go-ethereum/common"
)

// MintAuditForever forever mints, audits, and monitors all CCTP messages for
// the given network / recipient, forever.
func MintAuditForever(
	ctx context.Context,
	db *db.DB,
	client Client,
	network netconf.Network,
	backends ethbackend.Backends,
	recipient common.Address,
	minter common.Address,
	opts ...Option,
) error {
	clients := backends.Clients()
	xprov := xprovider.New(network, clients, nil)

	chains, err := getChains(network)
	if err != nil {
		return errors.Wrap(err, "get chains")
	}

	if err := auditForever(ctx, db, network.ID, xprov, clients, chains, recipient); err != nil {
		return errors.Wrap(err, "audit forever")
	}

	if err := mintForever(ctx, db, client, backends, chains, minter, opts...); err != nil {
		return errors.Wrap(err, "mint forever")
	}

	go monitorForever(ctx, chains, db)

	return nil
}

func getChains(network netconf.Network) ([]evmchain.Metadata, error) {
	var chains []evmchain.Metadata
	for _, c := range network.EVMChains() {
		chain, ok := evmchain.MetadataByID(c.ID)
		if !ok {
			return nil, errors.New("unknown chain", "chain_id", c.ID)
		}

		chains = append(chains, chain)
	}

	return chains, nil
}
