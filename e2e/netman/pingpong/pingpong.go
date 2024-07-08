package pingpong

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"math/big"
	"slices"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/forkjoin"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"

	"golang.org/x/sync/errgroup"
)

const startPingPongEdgesTimeout = 15 * time.Minute

// XDapp defines the deployed pingpong contract xdapp.
// The XDapp is graph of ping pong pairs that connects all chains to all chains.
// So given a network of N chains (vertexes), there will be N! pairs (edges).
type XDapp struct {
	contracts map[uint64]contract
	edges     []Edge
	backends  ethbackend.Backends
	deployer  common.Address
	network   netconf.ID
}

func Deploy(ctx context.Context, network netconf.Network, backends ethbackend.Backends) (XDapp, error) {
	log.Info(ctx, "Deploying ping pong contracts")

	deployer := eoa.MustAddress(network.ID, eoa.RoleTester)

	// Define a deploy function that deploys a ping pong contract to a chain.
	deployFunc := func(ctx context.Context, chain netconf.Chain) (contract, error) {
		log.Debug(ctx, "Deploying ping pong contract", "chain", chain.Name, "portal", chain.PortalAddress)

		txOpts, backend, err := backends.BindOpts(ctx, chain.ID, deployer)
		if err != nil {
			return contract{}, errors.Wrap(err, "deploy opts")
		}

		height, err := backend.BlockNumber(ctx)
		if err != nil {
			return contract{}, errors.Wrap(err, "block number")
		}

		addr, _, pingPong, err := bindings.DeployPingPong(txOpts, backend, chain.PortalAddress)
		if err != nil {
			return contract{}, errors.Wrap(err, "deploy ping pong contract")
		}

		log.Debug(ctx, "Deployed ping pong contract", "addr", addr, "chain", chain.Name, "height", height)

		return contract{
			Chain:        chain,
			Address:      addr,
			PingPong:     pingPong,
			DeployHeight: height,
		}, nil
	}

	// Start forkjoin for all portals
	results, cancel := forkjoin.NewWithInputs(ctx, deployFunc, network.EVMChains())

	defer cancel()

	// Collect the resulting contracts
	contracts := make(map[uint64]contract)
	for res := range results {
		if res.Err != nil {
			return XDapp{}, errors.Wrap(res.Err, "deploy")
		}

		contracts[res.Input.ID] = res.Output
	}

	dapp := XDapp{
		contracts: contracts,
		backends:  backends,
		edges:     edges(contracts),
		deployer:  deployer,
		network:   network.ID,
	}

	if err := dapp.fund(ctx); err != nil {
		return XDapp{}, errors.Wrap(err, "fund")
	}

	return dapp, nil
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
		txOpts, backend, err := d.backends.BindOpts(ctx, contract.Chain.ID, d.deployer)
		if err != nil {
			return err
		}

		// For ETH chains, fund it with 0.5 ETH
		fund := new(big.Int).Div(big.NewInt(params.Ether), big.NewInt(2))

		// for OMNI chains, fund it with 100 OMNI
		if contract.Chain.ID == d.network.Static().OmniExecutionChainID {
			fund = new(big.Int).Mul(big.NewInt(params.Ether), big.NewInt(100))
		}

		txOpts.Value = fund

		tx, err := contract.PingPong.Receive(txOpts)
		if err != nil {
			return errors.Wrap(err, "fund ping pong", "chain", contract.Chain.Name)
		}

		if _, err := backend.WaitMined(ctx, tx); err != nil {
			return errors.Wrap(err, "wait mined", "chain", contract.Chain.Name, "tx", tx.Hash())
		}
	}

	return nil
}

// StartAllEdges starts <parallel> ping pongs for all edges, each doing <count> hops.
func (d *XDapp) StartAllEdges(ctx context.Context, latest, parallel, count uint64) error {
	log.Info(ctx, "Starting ping pong contracts")
	ctx, cancel := context.WithTimeout(ctx, startPingPongEdgesTimeout)
	defer cancel()
	eg, ctx := errgroup.WithContext(ctx)

	// Group edges by from
	edgesByFrom := make(map[uint64][]contract)
	for _, edge := range d.edges {
		edgesByFrom[edge.From] = append(edgesByFrom[edge.From], d.contracts[edge.To])
	}

	for fromID, toContracts := range edgesByFrom {
		from := d.contracts[fromID]
		log.Debug(ctx, "Starting pingpong deployments for sender", "from", from.Chain.Name)
		eg.Go(func() error {
			// For a particular sender (from), we deploy contracts in series to avoid nonce errors
			for _, to := range toContracts {
				log.Debug(ctx, "Starting ping pong contract",
					"from", from.Chain.Name,
					"to", to.Chain.Name,
					"from_chain_id", from.Chain.ID,
					"parallel", parallel,
					"count", count,
					"shards", from.Chain.Shards,
				)

				shards := from.Chain.Shards
				for i := uint64(0); i < parallel; i++ {
					// First are latest, rest is finalized
					conf := xchain.ConfFinalized
					// Only use latest shard if the chain has it and "latest" is enabled (i.e. not 0)
					if slices.Contains(shards, xchain.ShardLatest0) && i < latest {
						conf = xchain.ConfLatest
					}

					txOpts, backend, err := d.backends.BindOpts(ctx, from.Chain.ID, d.deployer)
					if err != nil {
						return err
					}

					id := randomHex7()
					tx, err := from.PingPong.Start(txOpts, id, to.Chain.ID, uint8(conf), to.Address, count)
					if err != nil {
						return errors.Wrap(err, "start ping pong", "id", id, "from", from.Chain.Name, "to", to.Chain.Name, "conf", conf)
					}

					if _, err := bind.WaitMined(ctx, backend, tx); err != nil {
						return errors.Wrap(err, "wait mined", "chain", from.Chain.Name, "tx", tx.Hash())
					}
				}
			}

			return nil
		})
	}

	log.Debug(ctx, "Waiting for all ping pong edges to start", "timeout", startPingPongEdgesTimeout)
	if err := eg.Wait(); err != nil {
		return errors.Wrap(err, "wait parallel start")
	}

	return nil
}

// Watch watches all PingPong contracts for Ping events and logs them.
func (d *XDapp) Watch(ctx context.Context) error {
	// watch an individual pingpong contract
	watch := func(ctx context.Context, contract contract, backend *ethbackend.Backend) {
		lastBlockHeight := contract.DeployHeight
		ticker := time.NewTicker(5 * time.Second)

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				blockNumber, err := backend.BlockNumber(ctx)
				if err != nil {
					log.Error(ctx, "Error getting block number", err, "chain", contract.Chain.Name)
					continue
				}

				if blockNumber <= lastBlockHeight {
					continue
				}

				iter, err := contract.PingPong.FilterPing(&bind.FilterOpts{
					Start:   lastBlockHeight,
					End:     &blockNumber,
					Context: ctx,
				})
				if err != nil {
					log.Error(ctx, "Error filtering Ping events", err, "chain", contract.Chain.Name)
					continue
				}

				for iter.Next() {
					log.Debug(ctx, "Ping",
						"id", iter.Event.Id,
						"n", iter.Event.N,
						"on", contract.Chain.Name,
						"from", d.contracts[iter.Event.SrcChainID].Chain.Name,
						"height", iter.Event.Raw.BlockNumber,
					)
				}

				lastBlockHeight = blockNumber
			}
		}
	}

	for _, contract := range d.contracts {
		backend, err := d.backends.Backend(contract.Chain.ID)
		if err != nil {
			return err
		}

		go watch(ctx, contract, backend)
	}

	return nil
}

// WaitDone waits for all edges to complete the hops of a single ping pong.
// Note this doesn't wait for all parallel ping pongs to complete, it only waits for one of P.
func (d *XDapp) WaitDone(ctx context.Context) error {
	log.Info(ctx, "Waiting for ping pongs to complete")
	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	err := d.Watch(ctx)
	if err != nil {
		return errors.Wrap(err, "watch")
	}

	for _, edge := range d.edges {
		// Retry fetching done log until found or context is done
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
					log.Debug(ctx, "Ping pong done", "id", iter.Event.Id,
						"from", from.Chain.Name, "to", d.contracts[edge.To].Chain.Name, "times", iter.Event.Times)
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

			time.Sleep(time.Second)
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
	Chain        netconf.Chain
	Address      common.Address
	PingPong     *bindings.PingPong
	DeployHeight uint64
}

// randomHex7 returns a random 7-character hex string.
func randomHex7() string {
	bytes := make([]byte, 4)
	_, _ = rand.Read(bytes)
	hexString := hex.EncodeToString(bytes)

	// Trim the string to 7 characters
	if len(hexString) > 7 {
		hexString = hexString[:7]
	}

	return hexString
}
