package pingpong

import (
	"context"
	"math/big"
	"time"

	"github.com/omni-network/omni/contracts/bindings/examples"
	"github.com/omni-network/omni/e2e/netman"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/forkjoin"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"

	"golang.org/x/sync/errgroup"
)

// XDapp defines the deployed pingpong contract xdapp.
// The XDapp is graph of ping pong pairs that connects all chains to all chains.
// So given a network of N chains (vertexes), there will be N! pairs (edges).
type XDapp struct {
	contracts map[uint64]contract
	edges     []Edge
	backends  ethbackend.Backends
	operator  common.Address
}

func Deploy(ctx context.Context, netMgr netman.Manager, backends ethbackend.Backends) (XDapp, error) {
	log.Info(ctx, "Deploying ping pong contracts")

	operator := netMgr.Operator()

	// Define a deploy function that deploys a ping pong contract to a chain.
	deployFunc := func(ctx context.Context, portal netman.Portal) (contract, error) {
		log.Debug(ctx, "Deploying ping pong contract", "chain", portal.Chain.Name, "chainID", portal.Chain.ID, "portal", portal.DeployInfo.PortalAddress)

		portalAddr := portal.DeployInfo.PortalAddress

		txOpts, backend, err := backends.BindOpts(ctx, portal.Chain.ID, operator)
		if err != nil {
			return contract{}, errors.Wrap(err, "deploy opts")
		}

		height, err := backend.BlockNumber(ctx)
		if err != nil {
			return contract{}, errors.Wrap(err, "block number")
		}

		addr, _, pingPong, err := examples.DeployPingPong(txOpts, backend, portalAddr)
		if err != nil {
			return contract{}, errors.Wrap(err, "deploy ping pong contract")
		}

		log.Debug(ctx, "Deployed ping pong contract", "addr", addr, "chain", portal.Chain.Name, "height", height)

		return contract{
			Chain:        portal.Chain,
			Address:      addr,
			PingPong:     pingPong,
			DeployHeight: height,
		}, nil
	}

	// Start forkjoin for all portals
	portals := flatten[uint64, netman.Portal](netMgr.Portals())
	results, cancel := forkjoin.NewWithInputs(ctx, deployFunc, portals)

	defer cancel()

	// Collect the resulting contracts
	contracts := make(map[uint64]contract)
	for res := range results {
		if res.Err != nil {
			return XDapp{}, errors.Wrap(res.Err, "deploy")
		}

		contracts[res.Input.Chain.ID] = res.Output
	}

	dapp := XDapp{
		contracts: contracts,
		backends:  backends,
		edges:     edges(contracts),
		operator:  operator,
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

func (d *XDapp) LogBalances(ctx context.Context) error {
	for _, contract := range d.contracts {
		backend, err := d.backends.Backend(contract.Chain.ID)
		if err != nil {
			return err
		}

		balance, err := backend.EtherBalanceAt(ctx, contract.Address)
		if err != nil {
			return errors.Wrap(err, "balance at", "chain", contract.Chain.Name)
		}

		log.Debug(ctx, "Ping pong balance", "chain", contract.Chain.Name, "balance", balance)
	}

	return nil
}

func (d *XDapp) fund(ctx context.Context) error {
	for _, contract := range d.contracts {
		txOpts, backend, err := d.backends.BindOpts(ctx, contract.Chain.ID, d.operator)
		if err != nil {
			return err
		}

		fund := new(big.Int).Div(big.NewInt(params.Ether), big.NewInt(10)) // Fund it with 0.1 ETH
		txOpts.Value = fund

		tx, err := contract.PingPong.Receive(txOpts)
		if err != nil {
			return errors.Wrap(err, "fund ping pong", "chain", contract.Chain.Name)
		}

		if rec, err := bind.WaitMined(ctx, backend, tx); err != nil {
			return errors.Wrap(err, "wait mined", "chain", contract.Chain.Name, "tx", tx.Hash())
		} else if rec.Status != ethtypes.ReceiptStatusSuccessful {
			return errors.New("fund unsuccessful", "chain", contract.Chain.Name)
		}
	}

	return nil
}

// StartAllEdges starts <parallel> ping pongs for all edges, each doing <count> hops.
func (d *XDapp) StartAllEdges(ctx context.Context, parallel, count uint64) error {
	log.Info(ctx, "Starting ping pong contracts")
	eg, ctx := errgroup.WithContext(ctx)
	for _, edge := range d.edges {
		from := d.contracts[edge.From]
		to := d.contracts[edge.To]

		log.Debug(ctx, "Starting ping pong contract",
			"from", from.Chain.Name,
			"to", to.Chain.Name,
			"parallel", parallel,
			"count", count,
		)

		for i := uint64(0); i < parallel; i++ {
			eg.Go(func() error {
				txOpts, backend, err := d.backends.BindOpts(ctx, from.Chain.ID, d.operator)
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

				return nil
			})
		}
	}

	if err := eg.Wait(); err != nil {
		return errors.Wrap(err, "wait parallel start")
	}

	return nil
}

// WaitDone waits for all edges to complete the hops of a single ping pong.
// Note this doesn't wait for all parallel ping pongs to complete, it only waits for one of P.
func (d *XDapp) WaitDone(ctx context.Context) error {
	log.Info(ctx, "Waiting for ping pongs to complete")
	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()
	for _, edge := range d.edges {
		// Retry fetching done log until found or context is done
		backoff := expbackoff.New(ctx)
		for {
			if ctx.Err() != nil {
				return errors.Wrap(ctx.Err(), "timeout")
			}

			from := d.contracts[edge.From]
			iter, err := from.PingPong.FilterDone(&bind.FilterOpts{
				Start:   from.DeployHeight,
				Context: ctx,
			})
			if err != nil {
				return errors.Wrap(err, "filter", "from", from.Chain.Name)
			}

			var found bool
			for iter.Next() {
				if iter.Event.DestChainID == edge.To {
					log.Debug(ctx, "Ping pong done", "from", from.Chain.Name,
						"to", d.contracts[edge.To].Chain.Name, "times", iter.Event.Times)
					found = true

					break
				}
			}
			if err := iter.Error(); err != nil {
				return errors.Wrap(err, "iter error", "from", from.Chain.Name)
			} else if err := iter.Close(); err != nil {
				return errors.Wrap(err, "iter close", "from", from.Chain.Name)
			}

			if found {
				break
			}

			backoff()
		}
	}

	return nil
}

// Edge defines a unique edge between two ping pong contracts.
type Edge struct {
	From uint64 // From chain ID
	To   uint64 // To chain ID
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

func flatten[K comparable, V any](m map[K]V) []V {
	var resp []V
	for _, v := range m {
		resp = append(resp, v)
	}

	return resp
}
