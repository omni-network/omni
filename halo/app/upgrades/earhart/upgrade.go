// Package drake defines the third Omni consensus chain upgrade, named after
// Sir Francis Drake, an English explorer best known for making the second
// circumnavigation of the world in a single expedition between 1577 and 1580.
//
// It includes:
// - Unstaking of funds,
// - Processing of rewards withdrawals.
package drake

import (
	"context"
	"encoding/json"
	"time"

	evmredenomkeeper "github.com/omni-network/omni/halo/evmredenom/keeper"
	evmredenomsubmit "github.com/omni-network/omni/halo/evmredenom/submit"
	evmredenomtypes "github.com/omni-network/omni/halo/evmredenom/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/ethp2p"
	"github.com/omni-network/omni/lib/log"
	evmenginekeeper "github.com/omni-network/omni/octane/evmengine/keeper"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p/enode"

	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
)

const UpgradeName = "4_earhart"

func StoreUpgrades(_ context.Context) *storetypes.StoreUpgrades {
	return &storetypes.StoreUpgrades{
		Added: []string{
			evmredenomtypes.ModuleName, // Add the evmredenom module
		},
	}
}

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	evmEngine *evmenginekeeper.Keeper,
	redenom *evmredenomkeeper.Keeper,
	submitCfg evmredenomsubmit.Config,
) upgradetypes.UpgradeHandler {
	return func(ctx context.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		log.Info(ctx, "Running 4_earhart upgrade handler")

		// Initialize redenomination status to current execution head state root.
		header, err := evmEngine.GetExecutionHeader(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "get execution head")
		}

		if err := redenom.InitStatus(ctx, header.Root); err != nil {
			return nil, errors.Wrap(err, "initialize redenomination status")
		}

		if err := maybeSubmitRedenomination(ctx, submitCfg, header.Root); err != nil {
			return nil, errors.Wrap(err, "maybe submit redenomination")
		}

		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}

// maybeSubmitRedenomination submits redenomination account batches if configured.
func maybeSubmitRedenomination(ctx context.Context, cfg evmredenomsubmit.Config, root common.Hash) error {
	if !cfg.Enabled() {
		log.Debug(ctx, "Redenomination submission not enabled")
		return nil
	}

	log.Info(ctx, "Submitting redenomination account batches")

	privkey, err := crypto.LoadECDSA(cfg.PrivKey)
	if err != nil {
		return errors.Wrap(err, "load ECDSA private key")
	}
	from := crypto.PubkeyToAddress(privkey.PublicKey)

	ethCl, err := ethclient.DialContext(ctx, "redenom", cfg.EVMAddr)
	if err != nil {
		return errors.Wrap(err, "dial EVM client")
	}

	chainID, err := ethCl.ChainID(ctx)
	if err != nil {
		return errors.Wrap(err, "get chain ID")
	}
	backend, err := ethbackend.NewBackend("redenom", chainID.Uint64(), time.Second, ethCl, privkey)
	if err != nil {
		return errors.Wrap(err, "create backend")
	}

	enr, err := enode.ParseV4(cfg.EVMENR)
	if err != nil {
		return errors.Wrap(err, "parse ENR")
	}
	enr, err = ethp2p.DNSResolveHostname(ctx, enr)
	if err != nil {
		return errors.Wrap(err, "resolve ENR hostname")
	}

	return evmredenomsubmit.Do(ctx, from, backend, enr, root, cfg.BatchSize)
}

func GenesisState(codec.JSONCodec) (map[string]json.RawMessage, error) {
	return nil, nil //nolint:nilnil // map is for reading only
}
