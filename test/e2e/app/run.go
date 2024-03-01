package app

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/test/e2e/app/agent"
	"github.com/omni-network/omni/test/e2e/netman/pingpong"
	"github.com/omni-network/omni/test/e2e/types"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
)

// defaultPingPongN defines a few hours of ping pong messages after each deploy.
const defaultPingPongN = 1000

func DefaultDeployConfig() DeployConfig {
	return DeployConfig{
		AgentSecrets: agent.Secrets{}, // Empty secrets
		PingPongN:    defaultPingPongN,
	}
}

type DeployConfig struct {
	AgentSecrets agent.Secrets
	EigenFile    string
	PingPongN    uint64
	testConfig   bool // Internal use only (no command line flag).
}

// Deploy a new e2e network. It also starts all services in order to deploy private portals.
func Deploy(ctx context.Context, def Definition, cfg DeployConfig) (types.DeployInfos, error) {
	if err := Cleanup(ctx, def); err != nil {
		return nil, err
	}

	genesisValSetID := uint64(1) // validator set IDs start at 1
	genesisVals, err := toPortalValidators(def.Testnet.Validators)
	if err != nil {
		return nil, err
	}

	// Deploy public portals first so their addresses are available for setup.
	if err := def.Netman.DeployPublicPortals(ctx, genesisValSetID, genesisVals); err != nil {
		return nil, err
	}

	if err := Setup(ctx, def, cfg.AgentSecrets, cfg.testConfig); err != nil {
		return nil, err
	}

	if err := StartInitial(ctx, def.Testnet.Testnet, def.Infra); err != nil {
		return nil, err
	}

	if err := def.Netman.DeployPrivatePortals(ctx, genesisValSetID, genesisVals); err != nil {
		return nil, err
	}

	deployInfo := make(types.DeployInfos)

	if err := deployAVS(ctx, def, cfg, deployInfo); err != nil {
		return nil, err
	}

	for chain, info := range def.Netman.DeployInfo() {
		deployInfo.Set(chain.ID, types.ContractPortal, info.PortalAddress, info.DeployHeight)
	}

	if cfg.PingPongN == 0 {
		return deployInfo, nil
	}

	pp, err := pingpong.Deploy(ctx, def.Netman, def.Backends)
	if err != nil {
		return nil, errors.Wrap(err, "deploy pingpong")
	}

	err = pp.StartAllEdges(ctx, cfg.PingPongN)
	if err != nil {
		return nil, errors.Wrap(err, "start all edges")
	}

	pp.ExportDeployInfo(deployInfo)

	return deployInfo, nil
}

// E2ETestConfig is the configuration required to run a full e2e test.
type E2ETestConfig struct {
	Preserve bool
}

// DefaultE2ETestConfig returns a default configuration for a e2e test.
func DefaultE2ETestConfig() E2ETestConfig {
	return E2ETestConfig{}
}

// E2ETest runs a full e2e test.
func E2ETest(ctx context.Context, def Definition, cfg E2ETestConfig, secrets agent.Secrets) error {
	const pingpongN = 4 // Deploy and start ping pong
	depCfg := DeployConfig{
		AgentSecrets: secrets,
		PingPongN:    pingpongN,
		testConfig:   true,
	}

	deployInfo, err := Deploy(ctx, def, depCfg)
	if err != nil {
		return err
	}

	msgBatches := []int{3, 2, 1} // Send 6 msgs from each chain to each other chain
	msgsErr := StartSendingXMsgs(ctx, def.Netman, def.Backends, msgBatches...)

	if err := StartRemaining(ctx, def.Testnet.Testnet, def.Infra); err != nil {
		return err
	}

	if err := Wait(ctx, def.Testnet.Testnet, 5); err != nil { // allow some txs to go through
		return err
	}

	if def.Testnet.HasPerturbations() {
		return errors.New("perturbations not supported yet")
	}

	if def.Testnet.Evidence > 0 {
		return errors.New("evidence injection not supported yet")
	}

	// Wait for all messages to be sent
	log.Info(ctx, "Waiting for all cross chain messages to be sent")
	select {
	case <-ctx.Done():
		return errors.Wrap(ctx.Err(), "cancel")
	case err := <-msgsErr:
		if err != nil {
			return err
		}
	}

	if err := WaitAllSubmissions(ctx, def.Netman.Portals(), sum(msgBatches)); err != nil {
		return err
	}

	// Anvil doens't support subscriptions, we need to poll.
	// if err := pp.WaitDone(ctx); err != nil {
	//	return errors.Wrap(err, "wait pingpong")
	//}

	if err := Test(ctx, def, deployInfo, false); err != nil {
		return err
	}

	if err := LogMetrics(ctx, def); err != nil {
		return err
	}

	if cfg.Preserve {
		log.Warn(ctx, "Docker containers not stopped, --preserve=true", nil)
	} else if err := Cleanup(ctx, def); err != nil {
		return err
	}

	return nil
}

// Upgrade generates all local artifacts, but only copies the docker-compose file to the VMs.
// It them calls docker-compose up.
func Upgrade(ctx context.Context, def Definition) error {
	if err := Setup(ctx, def, agent.Secrets{}, false); err != nil {
		return err
	}

	return def.Infra.Upgrade(ctx)
}

// toPortalValidators returns the provided validator set as a lice of portal validators.
func toPortalValidators(validators map[*e2e.Node]int64) ([]bindings.Validator, error) {
	vals := make([]bindings.Validator, 0, len(validators))
	for val, power := range validators {
		addr, err := k1util.PubKeyToAddress(val.PrivvalKey.PubKey())
		if err != nil {
			return nil, errors.Wrap(err, "convert validator pubkey to address")
		}

		vals = append(vals, bindings.Validator{
			Addr:  addr,
			Power: uint64(power),
		})
	}

	return vals, nil
}
