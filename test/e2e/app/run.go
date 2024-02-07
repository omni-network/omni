package app

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/log"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
)

// Deploy a new e2e network. It also starts all services in order to deploy private portals.
func Deploy(ctx context.Context, def Definition, promSecrets PromSecrets) error {
	if err := Cleanup(ctx, def.Testnet.Testnet); err != nil {
		return err
	}

	genesisValSetID := uint64(1) // validator set IDs start at 1
	genesisVals, err := toPortalValidators(def.Testnet.Validators)
	if err != nil {
		return err
	}

	// Deploy public portals first so their addresses are available for setup.
	if err := def.Netman.DeployPublicPortals(ctx, genesisValSetID, genesisVals); err != nil {
		return err
	}

	if err := Setup(ctx, def, promSecrets); err != nil {
		return err
	}

	if err := Start(ctx, def.Testnet.Testnet, def.Infra); err != nil {
		return err
	}

	if err := def.Netman.DeployPrivatePortals(ctx, genesisValSetID, genesisVals); err != nil {
		return err
	}

	return nil
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
func E2ETest(ctx context.Context, def Definition, cfg E2ETestConfig) error {
	if err := Deploy(ctx, def, PromSecrets{}); err != nil { // No prometheus in tests for now.
		return err
	}

	msgBatches := []int{4, 3, 2, 1} // Send 10 msgs from each chain to each other chain
	msgsErr := StartSendingXMsgs(ctx, def.Netman.Portals(), msgBatches...)

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
	if err := <-msgsErr; err != nil {
		return err
	}

	if err := WaitAllSubmissions(ctx, def.Netman.Portals(), sum(msgBatches)); err != nil {
		return err
	}

	if err := Test(ctx, def); err != nil {
		return err
	}

	if err := LogMetrics(ctx, def); err != nil {
		return err
	}

	if cfg.Preserve {
		log.Warn(ctx, "Docker containers not stopped, --preserve=true", nil)
	} else if err := Cleanup(ctx, def.Testnet.Testnet); err != nil {
		return err
	}

	return nil
}

func sum(batches []int) int {
	var resp int
	for _, b := range batches {
		resp += b
	}

	return resp
}

// Convert cometbft testnet validators to solidity bindings.Validator, expected by portal constructor.
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
