package loadgen

import (
	"context"
	"crypto/ecdsa"
	"path/filepath"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain/connect"

	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

// Config is the configuration for the load generator.
type Config struct {
	ValidatorKeysGlob string
}

// Start starts the validator self delegation load generator.
// It does:
// - Validator self-delegation on periodic basis.
func Start(ctx context.Context, network netconf.Network, ethClients map[uint64]ethclient.Client, cfg Config, xCallerCfg XCallerConfig) error {
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

	var period = time.Hour * 6
	if network.ID == netconf.Devnet {
		period = time.Second * 5
	}

	for _, key := range keys {
		val := ethcrypto.PubkeyToAddress(key.PublicKey)
		go selfDelegateForever(ctx, contract, backend, val, period)
	}

	if xCallerCfg.Enabled {
		if len(xCallerCfg.ChainIDs) == 0 {
			return errors.New("xcaller enabled but no chain id pairs specified")
		}

		connector, err := connect.New(ctx, network.ID)
		if err != nil {
			return err
		}
		xcaller := NewXCaller(network.ID, connector, 2*time.Hour, xCallerCfg.ChainIDs)
		go xcaller.XCallForever(ctx)
	}

	return nil
}
