package rebalance

import (
	"context"

	"github.com/omni-network/omni/lib/cctp"
	cctpdb "github.com/omni-network/omni/lib/cctp/db"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	xprovider "github.com/omni-network/omni/lib/xchain/provider"

	"github.com/ethereum/go-ethereum/common"

	cosmosdb "github.com/cosmos/cosmos-db"
)

// Start starts rebalancing the solver's balance on the given network.
func Start(
	ctx context.Context,
	cfg Config,
	network netconf.Network,
	cctpClient cctp.Client,
	backends ethbackend.Backends,
	solver common.Address,
	dbDir string,
) error {
	ctx = log.WithCtx(ctx, "process", "rebalance")

	network, chains, err := newRebalanceNetwork(network)
	if err != nil {
		return err
	}

	xprov := xprovider.New(network, backends.Clients(), nil)

	db, err := newCCTPDB(dbDir)
	if err != nil {
		return errors.Wrap(err, "new cctp db")
	}

	if err := cctp.MintForever(ctx, db, cctpClient, backends, chains, solver); err != nil {
		return errors.Wrap(err, "mint forever")
	}

	if err := cctp.AuditForever(ctx, db, network.ID, xprov, backends.Clients(), chains, solver); err != nil {
		return errors.Wrap(err, "rebalance forever")
	}

	cctp.MonitorForever(ctx, db)

	go rebalanceForever(ctx, cfg, db, network, backends, solver)

	return nil
}

// newCCTPDB returns a new CCTP DB instance based on the given directory.
func newCCTPDB(dbDir string) (*cctpdb.DB, error) {
	if dbDir == "" {
		memDB := cosmosdb.NewMemDB()
		return cctpdb.New(memDB)
	}

	var err error
	lvlDB, err := cosmosdb.NewGoLevelDB("solver.rebalance.cctp", dbDir, nil)
	if err != nil {
		return nil, errors.Wrap(err, "new golevel db")
	}

	return cctpdb.New(lvlDB)
}

// newRebalanceNetwork returns the subset of `network` that can be rebalanced, along with list of in-network chains.
func newRebalanceNetwork(network netconf.Network) (netconf.Network, []evmchain.Metadata, error) {
	out := netconf.Network{ID: network.ID}
	chains := []evmchain.Metadata{}

	for _, chain := range network.EVMChains() {
		meta, ok := evmchain.MetadataByID(chain.ID)
		if !ok {
			return netconf.Network{}, nil, errors.New("no chain metadata", "chain", chain.Name)
		}

		if !canRebalance(chain.ID) {
			continue
		}

		chains = append(chains, meta)
		out.Chains = append(out.Chains, chain)
	}

	return out, chains, nil
}

// canRebalance returns true if the chain can be rebalanced.
func canRebalance(chainID uint64) bool {
	return cctp.IsSupportedChain(chainID)
}
