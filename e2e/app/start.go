package app

import (
	"context"
	"path/filepath"
	"slices"
	"sort"
	"sync"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
	"github.com/cometbft/cometbft/test/e2e/pkg/infra"
)

// StartInitial starts the initial nodes (start_at==0).
func StartInitial(ctx context.Context, testnet *e2e.Testnet, p infra.Provider) error {
	allNodes, err := getSortedNodes(testnet)
	if err != nil {
		return err
	}

	// Start initial nodes (StartAt: 0)
	initialNodes := make([]*e2e.Node, 0)
	for _, node := range allNodes {
		if node.StartAt > 0 {
			continue
		}
		initialNodes = append(initialNodes, node)
	}
	if len(initialNodes) == 0 {
		return errors.New("no initial nodes in testnet")
	}

	log.Info(ctx, "Starting initial network nodes...", "count", len(initialNodes))

	if err = p.StartNodes(ctx, initialNodes...); err != nil {
		return errors.Wrap(err, "starting initial nodes")
	}

	for _, node := range initialNodes {
		log.Info(ctx, "Starting node",
			"name", node.Name,
			"external_ip", node.ExternalIP,
			"proxy_port", node.ProxyPort,
			"prom", node.PrometheusProxyPort,
		)
		if _, err := waitForNode(ctx, node, 0, 15*time.Second); err != nil {
			return err
		}
	}

	networkHeight := testnet.InitialHeight

	// Wait for initial height
	log.Info(ctx, "Waiting for initial height",
		"height", networkHeight,
		"initial", len(initialNodes),
		"pending", len(allNodes)-len(initialNodes))

	_, _, err = waitForHeight(ctx, testnet, networkHeight)
	if err != nil {
		return err
	}

	return nil
}

// StartRemaining starts the remaining nodes (start_at>0).
func StartRemaining(ctx context.Context, testnet *e2e.Testnet, p infra.Provider) error {
	nodeQueue, err := getSortedNodes(testnet)
	if err != nil {
		return err
	}

	var remaining []*e2e.Node
	for _, node := range nodeQueue {
		if node.StartAt > 0 {
			remaining = append(remaining, node)
		}
	}

	if len(remaining) == 0 {
		return nil
	}

	block, blockID, err := waitForHeight(ctx, testnet, testnet.InitialHeight)
	if err != nil {
		return err
	}
	networkHeight := block.Height

	log.Debug(ctx, "Setting catchup node state sync", "height", block.Height, "catchup_nodes", len(remaining))

	// Update any state sync nodes with a trusted height and hash
	for _, node := range remaining {
		if node.StateSync || node.Mode == e2e.ModeLight {
			nodeDir := filepath.Join(testnet.Dir, node.Name)
			err = updateConfigStateSync(nodeDir, block.Height, blockID.Hash.Bytes())
			if err != nil {
				return err
			}
		}
	}

	for _, node := range remaining {
		if node.StartAt > networkHeight {
			// if we're starting a node that's ahead of
			// the last known height of the network, then
			// we should make sure that the rest of the
			// network has reached at least the height
			// that this node will start at before we
			// start the node.

			log.Info(ctx, "Waiting for network to advance before starting catchup node",
				"node", node.Name,
				"current_height", networkHeight,
				"wait_for_height", node.StartAt)

			networkHeight = node.StartAt

			if _, _, err := waitForHeight(ctx, testnet, networkHeight); err != nil {
				return err
			}
		}

		log.Info(ctx, "Starting catchup node", "node", node.Name, "height", node.StartAt)

		err := p.StartNodes(ctx, node)
		if err != nil {
			return errors.Wrap(err, "starting catchup node")
		}
		status, err := waitForNode(ctx, node, node.StartAt, 3*time.Minute)
		if err != nil {
			return err
		}
		log.Info(ctx, "Started catchup node", "name", node.Name, "height", status.SyncInfo.LatestBlockHeight)
	}

	return nil
}

// getSortedNodes returns a copy of the testnet nodes by startAt, then mode, then name.
func getSortedNodes(testnet *e2e.Testnet) ([]*e2e.Node, error) {
	if len(testnet.Nodes) == 0 {
		return nil, errors.New("no nodes in testnet")
	}

	nodeQueue := slices.Clone(testnet.Nodes)
	sort.SliceStable(nodeQueue, func(i, j int) bool {
		a, b := nodeQueue[i], nodeQueue[j]
		switch {
		case a.Mode == b.Mode:
			return false
		case a.Mode == e2e.ModeSeed:
			return true
		case a.Mode == e2e.ModeValidator && b.Mode == e2e.ModeFull:
			return true
		}

		return false
	})

	sort.SliceStable(nodeQueue, func(i, j int) bool {
		return nodeQueue[i].StartAt < nodeQueue[j].StartAt
	})

	return nodeQueue, nil
}

// waitForEVMs waits for EVMs to be available.
// This mitigates any issues if starting anvils or omni evms is slow.
// It also ensures the public RPCs are accessible.
func waitForEVMs(ctx context.Context, network netconf.Network, backends ethbackend.Backends) error {
	for _, chain := range network.EVMChains() {
		backend, err := backends.Backend(chain.ID)
		if err != nil {
			return errors.Wrap(err, "backend")
		}

		innerCtx := log.WithCtx(ctx, "chain", chain.Name)
		if err := waitForEVM(innerCtx, backend); err != nil {
			return errors.Wrap(err, "waiting for EVM", "chain", chain.Name)
		}
	}

	return nil
}

func waitForEVM(ctx context.Context, backend *ethbackend.Backend) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	var once sync.Once

	backoff := expbackoff.New(ctx)
	for ctx.Err() == nil {
		_, err := backend.BlockNumber(ctx)
		if err == nil {
			return nil
		}
		once.Do(func() {
			log.Warn(ctx, "Waiting for EVM chain to be available", err, "address", backend.Address())
		})
		backoff()
	}

	return errors.Wrap(ctx.Err(), "timeout")
}
