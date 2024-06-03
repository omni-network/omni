package app

import (
	"context"
	"time"

	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	rpctypes "github.com/cometbft/cometbft/rpc/core/types"
	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
	"github.com/cometbft/cometbft/test/e2e/pkg/infra/docker"
)

// perturb the running testnet.
func perturb(ctx context.Context, testnet types.Testnet) error {
	for _, node := range testnet.Nodes {
		for _, perturbation := range node.Perturbations {
			_, err := perturbNode(ctx, node, perturbation)
			if err != nil {
				return err
			}
			time.Sleep(3 * time.Second) // Give network some time to recover between each
		}
	}

	for service, purturbs := range testnet.Perturb {
		for _, p := range purturbs {
			if err := perturbService(ctx, service, testnet.Dir, p); err != nil {
				return errors.Wrap(err, "purturb service", "service", service)
			}
			time.Sleep(3 * time.Second) // Give network some time to recover between each
		}
	}

	return nil
}

// perturbService perturbs a docker service with a given perturbation.
func perturbService(ctx context.Context, service string, testnetDir string, perturb types.Perturb) error {
	ctx = log.WithCtx(ctx, "service", service)

	log.Info(ctx, "Perturbing service", "perturb", perturb)
	switch perturb {
	case types.PerturbRestart:
		if err := docker.ExecCompose(ctx, testnetDir, "restart", service); err != nil {
			return errors.Wrap(err, "restart service")
		}
	case types.PerturbStopStart:
		if err := docker.ExecCompose(ctx, testnetDir, "stop", service); err != nil {
			return errors.Wrap(err, "stop service")
		}
		time.Sleep(5 * time.Second)
		if err := docker.ExecCompose(ctx, testnetDir, "start", service); err != nil {
			return errors.Wrap(err, "start service")
		}
	case types.PerturbRollback:
		if err := docker.ExecCompose(ctx, testnetDir, "stop", service); err != nil {
			return errors.Wrap(err, "stop service")
		}
		if err := docker.ExecCompose(ctx, testnetDir, "run", service, "rollback"); err != nil {
			return errors.Wrap(err, "rollback service")
		}
		if err := docker.ExecCompose(ctx, testnetDir, "start", service); err != nil {
			return errors.Wrap(err, "start service")
		}
	case types.PerturbFuzzyHeadAttRoot, types.PerturbFuzzyHeadDropBlocks, types.PerturbFuzzyHeadDropMsgs, types.PerturbFuzzyHeadMoreMsgs:
		if err := docker.ExecCompose(ctx, testnetDir, "exec", service, "wget", "-O-", "localhost:8545/fuzzy_enable?perturb="+string(perturb)); err != nil {
			return errors.Wrap(err, "enable fuzzy head")
		}
		time.Sleep(6 * time.Second)
		if err := docker.ExecCompose(ctx, testnetDir, "exec", service, "wget", "-O-", "localhost:8545/fuzzy_disable"); err != nil {
			return errors.Wrap(err, "disable fuzzy head")
		}
	default:
		return errors.New("unknown service perturbation")
	}

	log.Info(ctx, "Perturbed service", "perturb", perturb)

	return nil
}

// perturbNode perturbs a node with a given perturbation, returning its status
// after recovering.
func perturbNode(ctx context.Context, node *e2e.Node, perturbation e2e.Perturbation) (*rpctypes.ResultStatus, error) {
	testnet := node.Testnet
	name := node.Name
	ctx = log.WithCtx(ctx, "name", name)

	switch perturbation {
	case e2e.PerturbationDisconnect:
		networkName := testnet.Name + "_" + testnet.Name
		log.Info(ctx, "Perturb node: disconnect")
		if err := docker.Exec(ctx, "network", "disconnect", networkName, name); err != nil {
			return nil, errors.Wrap(err, "disconnect node from network")
		}
		time.Sleep(10 * time.Second)
		if err := docker.Exec(ctx, "network", "connect", networkName, name); err != nil {
			return nil, errors.Wrap(err, "connect node tp network")
		}

	case e2e.PerturbationKill:
		log.Info(ctx, "Perturb node: kill")
		if err := docker.ExecCompose(ctx, testnet.Dir, "kill", "-s", "SIGKILL", name); err != nil {
			return nil, errors.Wrap(err, "kill node")
		}
		if err := docker.ExecCompose(ctx, testnet.Dir, "start", name); err != nil {
			return nil, errors.Wrap(err, "start node")
		}

	case e2e.PerturbationPause:
		log.Info(ctx, "Perturb node: pause")
		if err := docker.ExecCompose(ctx, testnet.Dir, "pause", name); err != nil {
			return nil, errors.Wrap(err, "pause node")
		}
		time.Sleep(10 * time.Second)
		if err := docker.ExecCompose(ctx, testnet.Dir, "unpause", name); err != nil {
			return nil, errors.Wrap(err, "unpause node")
		}

	case e2e.PerturbationRestart:
		log.Info(ctx, "Perturb node: restart")
		if err := docker.ExecCompose(ctx, testnet.Dir, "restart", name); err != nil {
			return nil, errors.Wrap(err, "restart node")
		}

	case e2e.PerturbationUpgrade:
		return nil, errors.New("upgrade perturbation not supported")

	default:
		return nil, errors.New("unexpected perturbation type", "type", perturbation)
	}

	status, err := waitForNode(ctx, node, 0, 20*time.Second)
	if err != nil {
		return nil, err
	}

	log.Info(ctx, "Node recovered from perturbation", "height", status.SyncInfo.LatestBlockHeight)

	return status, nil
}
