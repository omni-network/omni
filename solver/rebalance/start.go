package rebalance

import (
	"context"

	"github.com/omni-network/omni/lib/cctp"
	cctpdb "github.com/omni-network/omni/lib/cctp/db"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokenpricer"

	"github.com/ethereum/go-ethereum/common"

	cosmosdb "github.com/cosmos/cosmos-db"
)

// Start starts rebalancing the solver's balance on the given network.
func Start(
	ctx context.Context,
	network netconf.Network,
	cctpClient cctp.Client,
	pricer tokenpricer.Pricer,
	backends ethbackend.Backends,
	solver common.Address,
	dbDir string,
	opts ...Options,
) error {
	if network.ID != netconf.Mainnet {
		// Rebalancing is only supported on mainnet.
		return nil
	}

	ctx = log.WithCtx(ctx, "process", "rebalance")

	if err := monitorForever(ctx, network, backends.Clients(), solver); err != nil {
		return errors.Wrap(err, "monitor forever")
	}

	network = newCCTPNetwork(network)

	db, err := newCCTPDB(dbDir)
	if err != nil {
		return errors.Wrap(err, "new cctp db")
	}

	if err := cctp.MintAuditForever(ctx, db, cctpClient, network, backends, solver, solver); err != nil {
		return errors.Wrap(err, "mint forever")
	}

	o := defaultOps()
	for _, opt := range opts {
		opt(&o)
	}

	go rebalanceCCTPForever(ctx, o.interval, db, network, pricer, backends, solver)
	go rebalanceMantleForever(ctx, o.interval, backends, solver)

	return nil
}

// newCCTPDB returns a new CCTP DB instance based on the given directory.
func newCCTPDB(dbDir string) (*cctpdb.DB, error) {
	if dbDir == "" {
		memDB := cosmosdb.NewMemDB()
		return cctpdb.New(memDB)
	}

	var err error
	lvlDB, err := cosmosdb.NewGoLevelDB("rebalance.cctp", dbDir, nil)
	if err != nil {
		return nil, errors.Wrap(err, "new golevel db")
	}

	return cctpdb.New(lvlDB)
}

// newCCTPNetwork returns the subset of `network` that can be rebalanced via CCTP.
func newCCTPNetwork(network netconf.Network) netconf.Network {
	out := netconf.Network{ID: network.ID}

	for _, chain := range network.EVMChains() {
		if !cctp.IsSupportedChain(chain.ID) {
			continue
		}

		out.Chains = append(out.Chains, chain)
	}

	return out
}
