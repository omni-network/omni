// This script opens Holesky -> Base Sepolia ETH bridge orders every 2 seconds
// It is used to loadgen our solver on omega / staging
package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/crypto"
)

var (
	flagPrivKey = flag.String("private-key", "", "Private key to open orders with")
	flagL1RPC   = flag.String("l1-rpc", "", "L1 rpc url")
	flagNetwork = flag.String("network", "staging", "Network to open orders on")
)

func main() {
	flag.Parse()

	ctx := context.Background()
	err := run(ctx)
	if err != nil {
		log.Error(ctx, "❌ Failed", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	network := netconf.ID(*flagNetwork)

	if network.Verify() != nil {
		return errors.New("invalid network")
	}

	if network == netconf.Mainnet {
		return errors.New("no mainnet")
	}

	chainID, ok := netconf.EthereumChainID(network)
	if !ok {
		return errors.New("no L1 chain for network")
	}

	chain, ok := evmchain.MetadataByID(chainID)
	if !ok {
		return errors.New("metadata by id", "chain_id", chainID)
	}

	var rpc string
	if *flagL1RPC != "" {
		rpc = *flagL1RPC
	} else {
		rpc = types.PublicRPCByName(chain.Name)
	}

	if *flagPrivKey == "" {
		return errors.New("private key required")
	}

	pk, err := crypto.HexToECDSA(strings.TrimPrefix(*flagPrivKey, "0x"))
	if err != nil {
		return errors.Wrap(err, "hex to ecdsa")
	}

	user := crypto.PubkeyToAddress(pk.PublicKey)

	log.Info(ctx, "Running loadgen", "network", network, "rpc", rpc, "user", user.Hex())

	client, err := ethclient.Dial(chain.Name, rpc)
	if err != nil {
		return errors.Wrap(err, "dial")
	}

	backend, err := ethbackend.NewBackend(
		chain.Name,
		chain.ChainID,
		chain.BlockPeriod,
		client,
		pk,
	)
	if err != nil {
		return errors.Wrap(err, "new backend")
	}

	backends := ethbackend.BackendsFrom(map[uint64]*ethbackend.Backend{chainID: backend})

	// open an order every 2 seconds

	count := 0
	go func() {
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				log.Info(ctx, "Exiting")
				return
			case <-ticker.C:
				if ctx.Err() != nil {
					log.Info(ctx, "Exiting")
					return
				}

				log.Info(ctx, "Opening order", "count", count)

				// open order async
				go func() {
					count++

					orderID, err := solvernet.OpenOrder(
						ctx,
						network,
						chainID,
						backends,
						user,
						bindings.SolverNetOrderData{
							DestChainId: evmchain.IDBaseSepolia, // (works for staing / omega)
							Deposit: bindings.SolverNetDeposit{
								Amount: bi.Ether(0.01003),
							},
							Expenses: []bindings.SolverNetTokenExpense{},
							Calls: []bindings.SolverNetCall{{
								Target: user,
								Value:  bi.Ether(0.01),
							}},
						},
					)

					if err != nil {
						log.Error(ctx, "Failed to open order", err)
					} else {
						log.Info(ctx, "Opened order", "order_id", orderID)
					}
				}()
			}
		}
	}()

	// wait for signal to exit
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	return nil
}
