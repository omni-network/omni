package pingpong

import (
	"context"
	"math/big"
	"sort"

	examples "github.com/omni-network/omni/contracts/bindings/examples"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/test/e2e/netman"
	"github.com/omni-network/omni/test/e2e/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	etypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
)

// XDapp defines the deployed pingpong contract xdapp.
// The XDapp is graph of ping pong pairs that connects all chains to all chains.
// So given a network of N chains (vertexes), there will be N! pairs (edges).
type XDapp struct {
	contracts map[uint64]Contract
}

func Deploy(ctx context.Context, portals map[uint64]netman.Portal) (XDapp, error) {
	log.Info(ctx, "Deploying ping pong contracts")

	contracts := make(map[uint64]Contract)
	miner := make(miner)
	for chainID := range portals {
		portal := portals[chainID]
		portalAddr := portal.DeployInfo.PortalAddress

		height, err := portal.Client.BlockNumber(ctx)
		if err != nil {
			return XDapp{}, errors.Wrap(err, "block number")
		}

		txOpts := portal.TxOpts(ctx, nil)
		addr, tx, pingPong, err := examples.DeployPingPong(txOpts, portal.Client, portalAddr)
		if err != nil {
			return XDapp{}, errors.Wrap(err, "deploy ping pong contract")
		}

		contracts[chainID] = Contract{
			Client:       portal.Client,
			Chain:        portal.Chain,
			Address:      addr,
			PingPong:     pingPong,
			DeployHeight: height,
			txOpts:       txOpts,
		}

		miner[chainID] = tx
	}

	// TODO: replace with txmgr
	if err := miner.WaitMined(ctx, contracts); err != nil {
		return XDapp{}, errors.Wrap(err, "wait mined")
	}

	resp := XDapp{contracts: contracts}

	if err := resp.fund(ctx); err != nil {
		return XDapp{}, errors.Wrap(err, "fund")
	}

	return resp, nil
}

func (a *XDapp) ExportDeployInfo(resp types.DeployInfos) {
	for chainID, contract := range a.contracts {
		resp.Set(chainID, types.ContractPingPong, contract.Address, contract.DeployHeight)
	}
}

func (a *XDapp) fund(ctx context.Context) error {
	for _, contract := range a.contracts {
		fund := new(big.Int).Mul(big.NewInt(1_000_000), big.NewInt(params.GWei)) // Also fund it with 0.1 ETH
		opts := contract.TxOpts(ctx, fund)

		tx, err := contract.PingPong.Receive(opts)
		if err != nil {
			return errors.Wrap(err, "fund ping pong", "chain", contract.Chain.Name)
		}

		// TODO: replace with txmgr
		if _, err := bind.WaitMined(ctx, contract.Client, tx); err != nil {
			return errors.Wrap(err, "wait mined", "chain", contract.Chain.Name, "tx", tx.Hash())
		}
	}

	return nil
}

func (a *XDapp) StartAllEdges(ctx context.Context, count uint64) error {
	log.Info(ctx, "Starting ping pong contracts")
	for _, edge := range a.edges() {
		from := a.contracts[edge.From]
		to := a.contracts[edge.To]

		log.Info(ctx, "Starting ping pong contract",
			"from", from.Chain.Name,
			"to", to.Chain.Name,
			"count", count,
		)

		txOpts := from.TxOpts(ctx, nil)
		tx, err := from.PingPong.Start(txOpts, to.Chain.ID, to.Address, count)
		if err != nil {
			return errors.Wrap(err, "start ping pong", "from", from.Chain.Name, "to", to.Chain.Name)
		}

		// TODO: replace with txmgr
		if _, err := bind.WaitMined(ctx, from.Client, tx); err != nil {
			return errors.Wrap(err, "wait mined", "chain", from.Chain.Name, "tx", tx.Hash())
		}
	}

	return nil
}

// WaitDone waits for all ping pongs to be done.
// Note this doesn't work on anvil since it doesn't support subscriptions.
func (a *XDapp) WaitDone(ctx context.Context) error {
	log.Info(ctx, "Waiting for ping pongs to complete")
	done := make(chan *examples.PingPongDone, len(a.edges()))
	for _, edge := range a.edges() {
		from := a.contracts[edge.From]
		_, err := from.PingPong.WatchDone(&bind.WatchOpts{
			Start:   &from.DeployHeight,
			Context: ctx,
		}, done)
		if err != nil {
			return errors.Wrap(err, "watch done", "from", from.Chain.Name)
		}
	}

	// Wait for all ping pongs to be done
	for range a.edges() {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case d := <-done:
			log.Debug(ctx, "Ping pong done", "to", a.contracts[d.DestChainID].Chain.Name, "count", d.Times)
		}
	}

	return nil
}

type Edge struct {
	From uint64
	To   uint64
}

func (e Edge) Equals(other Edge) bool {
	return e.From == other.From && e.To == other.To || e.From == other.To && e.To == other.From
}

// edges returns a deterministic map of unique edges between chains.
func (a *XDapp) edges() []Edge {
	var all []Edge
	for fromChainID := range a.contracts {
		for toChainID := range a.contracts {
			if fromChainID == toChainID {
				continue
			}

			all = append(all, Edge{From: fromChainID, To: toChainID})
		}
	}

	// Order by fromChainID
	sort.Slice(all, func(i, j int) bool {
		return all[i].From < all[j].From
	})

	// Deduplicate
	var resp []Edge
	for _, candidate := range all {
		unique := true
		for _, existing := range resp {
			if candidate.Equals(existing) {
				unique = false
				break
			}
		}
		if unique {
			resp = append(resp, candidate)
		}
	}

	return resp
}

// Contract defines a deployed contract.
type Contract struct {
	Client       *ethclient.Client
	Chain        types.EVMChain
	Address      common.Address
	PingPong     *examples.PingPong
	txOpts       *bind.TransactOpts // TODO(corver): Replace this with a txmgr.
	DeployHeight uint64
}

// TxOpts returns transaction options using the deploy key.
func (p Contract) TxOpts(ctx context.Context, value *big.Int) *bind.TransactOpts {
	clone := *p.txOpts
	clone.Context = ctx
	clone.Value = value

	return &clone
}

type miner map[uint64]*etypes.Transaction

// TODO: remove after txmgr is implemented.
func (m miner) WaitMined(ctx context.Context, contracts map[uint64]Contract) error {
	for chainID, tx := range m {
		if _, err := bind.WaitMined(ctx, contracts[chainID].Client, tx); err != nil {
			return errors.Wrap(err, "wait mined", "chain", contracts[chainID].Chain.Name, "tx", tx.Hash())
		}
	}

	return nil
}
