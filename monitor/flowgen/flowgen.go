package flowgen

import (
	"context"
	"fmt"
	"time"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
	"github.com/omni-network/omni/monitor/flowgen/bridging"
	"github.com/omni-network/omni/monitor/flowgen/types"
)

func Start(ctx context.Context, network netconf.Network, rpcEndpoints xchain.RPCEndpoints, keyPath string) error {
	if keyPath == "" {
		return errors.New("private key is required")
	}

	privKey, err := ethcrypto.LoadECDSA(keyPath)
	if err != nil {
		return errors.Wrap(err, "load xcaller key", "path", keyPath)
	}

	backends, err := ethbackend.BackendsFromNetwork(network, rpcEndpoints, privKey)
	if err != nil {
		return err
	}

	jobs := []types.Job{
		bridging.NewJob(),
		// {
		// 	Name:    "Symbiotic",
		// 	Run:     symbiotic.NewRunner(),
		// 	Cadence: 1 * time.Hour,
		// 	Spend:   symbiotic.Deposit,
		// },
	}

	for _, job := range jobs {
		go func() {
			ticker := time.NewTicker(job.Cadence())
			defer ticker.Stop()

			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					if err := job.Run(ctx, backends); err != nil {
						log.Error(ctx, fmt.Sprintf("job %s failed", job.Name()), err)
					}
				}
			}
		}()
	}

	return nil
}
