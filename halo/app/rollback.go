package app

import (
	"context"
	"fmt"
	"math/big"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	cmtcmd "github.com/cometbft/cometbft/cmd/cometbft/commands"

	"github.com/ethereum/go-ethereum/common"

	"cosmossdk.io/store"
	storemetrics "cosmossdk.io/store/metrics"
	dbm "github.com/cosmos/cosmos-db"
)

func Rollback(ctx context.Context, cfg Config, removeBlock bool) error {
	engineCl, err := newEngineClient(ctx, cfg, cfg.Network, nil)
	if err != nil {
		return err
	}

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

	db, err := dbm.NewDB("application", cfg.BackendType(), cfg.DataDir())
	if err != nil {
		return errors.Wrap(err, "create db")
	}

	// Rollback CometBFT state
	height, hash, err := cmtcmd.RollbackState(&cfg.Comet, removeBlock)
	if err != nil {
		return errors.Wrap(err, "rollback comet state")
	}

	// Rollback the multistore
	cms := store.NewCommitMultiStore(db, newSDKLogger(ctx), storemetrics.NewNoOpMetrics())
	if err := cms.RollbackToVersion(height); err != nil {
		return errors.Wrap(err, "rollback to height")
	}

	log.Info(ctx, "Rolled back consensus state", "height", height, "hash", fmt.Sprintf("%X", hash))

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
