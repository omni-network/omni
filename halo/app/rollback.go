package app

import (
	"context"
	"fmt"
	"math/big"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	cmtcmd "github.com/cometbft/cometbft/cmd/cometbft/commands"

	"github.com/ethereum/go-ethereum/common"

	dbm "github.com/cosmos/cosmos-db"
)

type RollbackConfig struct {
	Config
	RollbackEVM      bool
	RemoveCometBlock bool
}

func Rollback(ctx context.Context, cfg RollbackConfig) error {
	db, err := dbm.NewDB("application", cfg.BackendType(), cfg.DataDir())
	if err != nil {
		return errors.Wrap(err, "create db")
	}

	baseAppOpts, err := makeBaseAppOpts(cfg.Config)
	if err != nil {
		return errors.Wrap(err, "make base app opts")
	}

	engineCl, err := newEngineClient(ctx, cfg.Config, cfg.Network, nil)
	if err != nil {
		return err
	}

	privVal, err := loadPrivVal(cfg.Config)
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
		netconf.ChainNamer(cfg.Network),
		baseAppOpts...,
	)
	if err != nil {
		return errors.Wrap(err, "new app")
	}

	// Rollback CometBFT state
	height, hash, err := cmtcmd.RollbackState(&cfg.Comet, cfg.RemoveCometBlock)
	if err != nil {
		return errors.Wrap(err, "rollback comet state")
	}

	// Rollback the multistore
	if err := app.CommitMultiStore().RollbackToVersion(height); err != nil {
		return errors.Wrap(err, "rollback to height")
	}

	log.Info(ctx, "Rolled back consensus state", "height", height, "hash", fmt.Sprintf("%X", hash))

	if !cfg.RollbackEVM {
		log.Debug(ctx, "Not rolling back Omni EVM since --rollback-evm=false")

		return nil
	}

	// TODO(corver): Rolling back the EVM fails with `debug_setHead` not active/enabled.
	// Rolling back the EVM might not be required, since EngineAPI hard-sets the head in any-case.
	// If it is required, figure out how to enable debug_setHead.

	latestHeigth, err := engineCl.BlockNumber(ctx)
	if err != nil {
		return err
	}

	latestBlock, err := engineCl.BlockByNumber(ctx, big.NewInt(int64(latestHeigth)))
	if err != nil {
		return err
	} else if latestBlock.BeaconRoot() == nil {
		return errors.New("cannot rollback EVM with nil beacon root", "height", latestHeigth)
	}

	// Rollback EVM if latest EVM block built on-top of new rolled-back consensus head.
	if *latestBlock.BeaconRoot() != common.BytesToHash(hash) {
		return errors.New("cannot rollback EVM, latest EVM block not built on new rolled-back state",
			"evm_height", latestHeigth,
			"evm_beacon_root", *latestBlock.BeaconRoot(),
		)
	}

	if err := engineCl.SetHead(ctx, latestHeigth-1); err != nil {
		return errors.Wrap(err, "set head")
	}

	rolledBackBlock, err := engineCl.BlockByNumber(ctx, big.NewInt(int64(latestHeigth-1)))
	if err != nil {
		return err
	}

	log.Info(ctx, "Rolled back execution state", "height", rolledBackBlock.Number(), "hash", rolledBackBlock.Hash())

	return nil
}
