package pingpong

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings/examples"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/test/e2e/backend"
	"github.com/omni-network/omni/test/e2e/netman"
	"github.com/omni-network/omni/test/e2e/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
)

// XDapp defines the deployed pingpong contract xdapp.
// The XDapp is graph of ping pong pairs that connects all chains to all chains.
// So given a network of N chains (vertexes), there will be N! pairs (edges).
type XDapp struct {
	contracts map[uint64]contract
	edges     []Edge
	backends  backend.Backends
}

func Deploy(ctx context.Context, netman netman.Manager, backends backend.Backends) (XDapp, error) {
	log.Info(ctx, "Deploying ping pong contracts")

	contracts := make(map[uint64]contract)
	for chainID, portal := range netman.Portals() {
		portalAddr := portal.DeployInfo.PortalAddress

		txOpts, backend, err := backends.BindOpts(ctx, chainID)
		if err != nil {
			return XDapp{}, errors.Wrap(err, "deploy opts")
		}

		height, err := backend.BlockNumber(ctx)
		if err != nil {
			return XDapp{}, errors.Wrap(err, "block number")
		}

		addr, _, pingPong, err := examples.DeployPingPong(txOpts, backend, portalAddr)
		if err != nil {
			return XDapp{}, errors.Wrap(err, "deploy ping pong contract")
		}

		contracts[chainID] = contract{
			Chain:        portal.Chain,
			Address:      addr,
			PingPong:     pingPong,
			DeployHeight: height,
		}
	}

	dapp := XDapp{
		contracts: contracts,
		backends:  backends,
		edges:     edges(contracts),
	}

	if err := dapp.fund(ctx); err != nil {
		return XDapp{}, errors.Wrap(err, "fund")
	}

	return dapp, nil
}

func (d *XDapp) ExportDeployInfo(resp types.DeployInfos) {
	for chainID, contract := range d.contracts {
		resp.Set(chainID, types.ContractPingPong, contract.Address, contract.DeployHeight)
	}
}

func (d *XDapp) fund(ctx context.Context) error {
	for _, contract := range d.contracts {
		txOpts, backend, err := d.backends.BindOpts(ctx, contract.Chain.ID)
		if err != nil {
			return err
		}

		fund := new(big.Int).Mul(big.NewInt(1_000_000), big.NewInt(params.GWei)) // Fund it with 0.1 ETH
		txOpts.Value = fund

		tx, err := contract.PingPong.Receive(txOpts)
		if err != nil {
			return errors.Wrap(err, "fund ping pong", "chain", contract.Chain.Name)
		}

		if _, err := bind.WaitMined(ctx, backend, tx); err != nil {
			return errors.Wrap(err, "wait mined", "chain", contract.Chain.Name, "tx", tx.Hash())
		}
	}

	return nil
}

func (d *XDapp) StartAllEdges(ctx context.Context, count uint64) error {
	log.Info(ctx, "Starting ping pong contracts")
	for _, edge := range d.edges {
		from := d.contracts[edge.From]
		to := d.contracts[edge.To]

		log.Debug(ctx, "Starting ping pong contract",
			"from", from.Chain.Name,
			"to", to.Chain.Name,
			"count", count,
		)

		txOpts, backend, err := d.backends.BindOpts(ctx, from.Chain.ID)
		if err != nil {
			return err
		}

		tx, err := from.PingPong.Start(txOpts, to.Chain.ID, to.Address, count)
		if err != nil {
			return errors.Wrap(err, "start ping pong", "from", from.Chain.Name, "to", to.Chain.Name)
		}

		if _, err := bind.WaitMined(ctx, backend, tx); err != nil {
			return errors.Wrap(err, "wait mined", "chain", from.Chain.Name, "tx", tx.Hash())
		}
	}

	return nil
}

// WaitDone waits for all ping pongs to be done.
// Note this doesn't work on anvil since it doesn't support subscriptions.
func (d *XDapp) WaitDone(ctx context.Context) error {
	log.Info(ctx, "Waiting for ping pongs to complete")
	done := make(chan *examples.PingPongDone, len(d.edges))
	for _, edge := range d.edges {
		from := d.contracts[edge.From]
		_, err := from.PingPong.WatchDone(&bind.WatchOpts{
			Start:   &from.DeployHeight,
			Context: ctx,
		}, done)
		if err != nil {
			return errors.Wrap(err, "watch done", "from", from.Chain.Name)
		}
	}

	// Wait for all ping pongs to be done
	for range d.edges {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case resp := <-done:
			log.Debug(ctx, "Ping pong done", "to", d.contracts[resp.DestChainID].Chain.Name, "count", resp.Times)
		}
	}

	return nil
}

type Edge struct {
	From uint64
	To   uint64
}

// edges creates a deterministic map of unique edges between chains.
func edges(contracts map[uint64]contract) []Edge {
	var resp []Edge
	var arr []contract
	// flatten contracts
	for _, v := range contracts {
		arr = append(arr, v)
	}

	// get all unique edges
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			resp = append(resp, Edge{From: arr[i].Chain.ID, To: arr[j].Chain.ID})
		}
	}

	return resp
}

// contract defines a deployed contract.
type contract struct {
	Chain        types.EVMChain
	Address      common.Address
	PingPong     *examples.PingPong
	DeployHeight uint64
}
