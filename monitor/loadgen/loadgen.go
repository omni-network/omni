package loadgen

import (
	"context"
	"crypto/ecdsa"
	"path/filepath"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

const (
	selfDelegationPeriod       = time.Minute * 2
	selfDelegationPeriodDevnet = time.Second * 5
	xCallerPeriod              = time.Hour * 2
	xCallerPeriodDevnet        = time.Second * 30
)

// Config is the configuration for the load generator.
type Config struct {
	// ValidatorKeysGlob defines the paths to the validator keys used for self-delegation.
	ValidatorKeysGlob string
	// XCallerKey path to the xcaller private key.
	XCallerKey string
}

// Start starts the validator self delegation load generator.
// It does:
// - Validator self- and normal-delegation on periodic basis.
// - Makes XCalls from -> to random EVM portals on periodic basis.
func Start(ctx context.Context, network netconf.Network, ethClients map[uint64]ethclient.Client, cfg Config) error {
	err := startDelegation(ctx, network, ethClients, cfg)
	if err != nil {
		return errors.Wrap(err, "start self delegation")
	}

	err = startXCaller(ctx, network, ethClients, cfg.XCallerKey)
	if err != nil {
		return errors.Wrap(err, "start xcaller")
	}

	return nil
}

func startDelegation(ctx context.Context, network netconf.Network, ethClients map[uint64]ethclient.Client, cfg Config) error {
	// Only generate load in ephemeral networks, devnet and staging.
	if !network.ID.IsEphemeral() {
		return nil
	} else if cfg.ValidatorKeysGlob == "" {
		// Skip if no validator keys are provided.
		return nil
	}

	var keys []*ecdsa.PrivateKey
	keysPaths, err := filepath.Glob(cfg.ValidatorKeysGlob)
	if err != nil {
		return errors.Wrap(err, "glob validator keys", "glob", cfg.ValidatorKeysGlob)
	}
	for _, keyPath := range keysPaths {
		key, err := ethcrypto.LoadECDSA(keyPath)
		if err != nil {
			return errors.Wrap(err, "load validator key", "path", keyPath)
		}

		keys = append(keys, key)
	}

	omniEVM, ok := network.OmniEVMChain()
	if !ok {
		return errors.New("omniEVM chain not found")
	}

	ethCl, ok := ethClients[omniEVM.ID]
	if !ok {
		return errors.New("eth client not found")
	}

	backend, err := ethbackend.NewBackend(omniEVM.Name, omniEVM.ID, omniEVM.BlockPeriod, ethCl, keys...)
	if err != nil {
		return err
	}

	contract, err := bindings.NewStaking(common.HexToAddress(predeploys.Staking), backend)
	if err != nil {
		return errors.Wrap(err, "new omni stake")
	}

	period := selfDelegationPeriod
	if network.ID == netconf.Devnet {
		period = selfDelegationPeriodDevnet
	}

	for i, key := range keys {
		// Use each validator key as delegator
		delegator := ethcrypto.PubkeyToAddress(key.PublicKey)
		val := delegator // For even i, delegate to self.
		if i%2 == 1 {    // For odd i, delegate to previous validator (normal non-self delegation).
			val = ethcrypto.PubkeyToAddress(keys[i-1].PublicKey)
		}

		if network.ID == netconf.Staging {
			go maybeDosForever(ctx, backend, delegator, val, period)
		} else if network.ID == netconf.Devnet {
			go delegateForever(ctx, contract, backend, delegator, val, period)
		}
	}

	return nil
}

func startXCaller(ctx context.Context, network netconf.Network, ethClients map[uint64]ethclient.Client, keyPath string) error {
	if keyPath == "" {
		// Skip if no key is provided.
		return nil
	}

	privKey, err := ethcrypto.LoadECDSA(keyPath)
	if err != nil {
		return errors.Wrap(err, "load xcaller key", "path", keyPath)
	}

	backends, err := ethbackend.BackendsFromClients(ethClients, privKey)
	if err != nil {
		return err
	}

	xCallerAddr := eoa.MustAddress(network.ID, eoa.RoleXCaller)

	period := xCallerPeriod
	if network.ID == netconf.Devnet {
		period = xCallerPeriodDevnet
	}
	xCallCfg := xCallConfig{
		NetworkID:   network.ID,
		XCallerAddr: xCallerAddr,
		Period:      period,
		Backends:    backends,
		Chains:      network.EVMChains(),
	}
	go xCallForever(ctx, xCallCfg)

	return nil
}
