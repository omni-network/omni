package app

import (
	"context"
	"fmt"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	cmtcmd "github.com/cometbft/cometbft/cmd/cometbft/commands"

	dbm "github.com/cosmos/cosmos-db"
)

func DefaultRollbackConfig() RollbackConfig {
	return RollbackConfig{
		// RemoveCometBlock (--hard=true) rollback doesn't add value in most use-cases.
		// Blocks cannot be re-proposed/re-consensused since that would result in validator slashing.
		RemoveCometBlock: false,
	}
}

type RollbackConfig struct {
	RemoveCometBlock bool
}

func Rollback(ctx context.Context, cfg Config, rCfg RollbackConfig) error {
	db, err := dbm.NewDB("application", cfg.BackendType(), cfg.DataDir())
	if err != nil {
		return errors.Wrap(err, "create db")
	}

	baseAppOpts, err := makeBaseAppOpts(cfg)
	if err != nil {
		return errors.Wrap(err, "make base app opts")
	}

	engineCl, err := newEngineClient(ctx, cfg, cfg.Network, nil)
	if err != nil {
		return err
	}

	privVal, err := loadPrivVal(cfg)
	if err != nil {
		return errors.Wrap(err, "load validator key")
	}

	voter, err := newVoterLoader(privVal.Key.PrivKey)
	if err != nil {
		return errors.Wrap(err, "new voter loader")
	}

	//nolint:contextcheck // False positive.
	app, err := newApp(
		newSDKLogger(ctx),
		db,
		engineCl,
		voter,
		netconf.ChainVersionNamer(cfg.Network),
		netconf.ChainNamer(cfg.Network),
		burnEVMFees{},
		baseAppOpts...,
	)
	if err != nil {
		return errors.Wrap(err, "new app")
	}

	// Rollback CometBFT state
	height, hash, err := cmtcmd.RollbackState(&cfg.Comet, rCfg.RemoveCometBlock)
	if err != nil {
		return errors.Wrap(err, "rollback comet state")
	}

	// Rollback the multistore
	if err := app.CommitMultiStore().RollbackToVersion(height); err != nil {
		return errors.Wrap(err, "rollback to height")
	}

	log.Info(ctx, "Rolled back consensus state", "height", height, "hash", fmt.Sprintf("%X", hash))

	return nil
}
