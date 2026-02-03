package balancesnap

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	halocfg "github.com/omni-network/omni/halo/config"
	evmredenomsubmit "github.com/omni-network/omni/halo/evmredenom/submit"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethp2p"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/cometbft/cometbft/rpc/client"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p/enode"
)

// Start starts a goroutine that waits for halt height and snapshots balances.
// Returns immediately if not enabled or network is not staging/devnet.
func Start(
	ctx context.Context,
	network netconf.ID,
	evmRedenomCfg evmredenomsubmit.Config,
	homeDir string,
	consensusClient client.Client,
	cprov cchain.Provider,
	asyncAbort chan<- error,
) {
	// Only enabled if evmredenom is enabled
	if !evmRedenomCfg.Enabled {
		return
	}

	// Get halt height for this network (0 means disabled)
	haltHeight := halocfg.HaltHeight(network)
	if haltHeight == 0 {
		return
	}

	go func() {
		if err := run(ctx, haltHeight, evmRedenomCfg, homeDir, consensusClient, cprov); err != nil {
			log.Error(ctx, "Balance snapshot failed", err)
			asyncAbort <- errors.Wrap(err, "balance snapshot")
		}
	}()
}

// run executes the balance snapshot workflow.
func run(
	ctx context.Context,
	haltHeight uint64,
	evmRedenomCfg evmredenomsubmit.Config,
	homeDir string,
	consensusClient client.Client,
	cprov cchain.Provider,
) error {
	log.Info(ctx, "Balance snapshot enabled", "halt_height", haltHeight)

	// Wait for consensus to reach halt height
	if err := waitForHeight(ctx, consensusClient, haltHeight); err != nil {
		return errors.Wrap(err, "wait for halt height")
	}

	log.Info(ctx, "Halt height reached, querying EVM head", "height", haltHeight)

	// Connect to archive RPC to get latest EVM header
	archive, err := ethclient.DialContext(ctx, "omni_evm", evmRedenomCfg.RPCArchive)
	if err != nil {
		return errors.Wrap(err, "dial archive RPC")
	}
	defer archive.Close()

	// Query latest EVM block header
	evmHeader, err := archive.HeaderByType(ctx, ethclient.HeadLatest)
	if err != nil {
		return errors.Wrap(err, "get latest EVM header")
	}

	log.Info(ctx, "EVM head retrieved",
		"block_number", evmHeader.Number.Uint64(),
		"state_root", evmHeader.Root.Hex(),
	)

	// Hardcode output paths
	evmOutputPath := filepath.Join(homeDir, "data", "evm_balances_halt.json")
	stakeOutputPath := filepath.Join(homeDir, "data", "staking_balances_halt.json")

	// Snapshot EVM balances
	log.Info(ctx, "Snapshotting EVM balances", "output", evmOutputPath)
	if err := snapshotEVMBalances(ctx, evmRedenomCfg, evmHeader.Root, evmHeader.Number.Uint64(), evmOutputPath); err != nil {
		return errors.Wrap(err, "snapshot EVM balances")
	}

	// Snapshot staking balances
	log.Info(ctx, "Snapshotting staking balances", "output", stakeOutputPath)
	if err := snapshotStakingBalances(ctx, cprov, stakeOutputPath); err != nil {
		return errors.Wrap(err, "snapshot staking balances")
	}

	log.Info(ctx, "Balance snapshot completed successfully")

	return nil
}

// waitForHeight polls consensus client until target height is reached.
func waitForHeight(ctx context.Context, cl client.Client, targetHeight uint64) error {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return errors.Wrap(ctx.Err(), "context done")
		case <-ticker.C:
			status, err := cl.Status(ctx)
			if err != nil {
				log.Warn(ctx, "Failed to get consensus status (will retry)", err)
				continue
			}

			if status.SyncInfo.LatestBlockHeight < 0 {
				log.Warn(ctx, "Invalid block height (will retry)", nil, "height", status.SyncInfo.LatestBlockHeight)
				continue
			}

			currentHeight := uint64(status.SyncInfo.LatestBlockHeight) //nolint:gosec // Validated >= 0 above
			if currentHeight >= targetHeight {
				return nil
			}

			if currentHeight%100 == 0 {
				log.Debug(ctx, "Waiting for halt height",
					"current", currentHeight,
					"target", targetHeight,
				)
			}
		}
	}
}

// snapshotEVMBalances fetches EVM balances using snap protocol and writes to JSON.
func snapshotEVMBalances(
	ctx context.Context,
	evmRedenomCfg evmredenomsubmit.Config,
	stateRoot common.Hash,
	blockNumber uint64,
	outputPath string,
) error {
	// Parse ENR
	peer, err := enode.ParseV4(evmRedenomCfg.EVMENR)
	if err != nil {
		return errors.Wrap(err, "parse ENR")
	}

	// Resolve hostname
	peer, err = ethp2p.DNSResolveHostname(ctx, peer)
	if err != nil {
		return errors.Wrap(err, "resolve ENR hostname")
	}

	// Connect to archive RPC
	archive, err := ethclient.DialContext(ctx, "omni_evm", evmRedenomCfg.RPCArchive)
	if err != nil {
		return errors.Wrap(err, "dial archive RPC")
	}
	defer archive.Close()

	// Load preimages from genesis file
	preimages, err := loadGenesisPreimages(evmRedenomCfg.Genesis)
	if err != nil {
		return errors.Wrap(err, "load genesis preimages")
	}

	log.Info(ctx, "Fetching EVM balances via snap protocol",
		"peer", peer.String(),
		"state_root", stateRoot.Hex(),
		"preimages", len(preimages),
	)

	// Call GetEVMBalances
	balances, err := GetEVMBalances(ctx, peer, stateRoot, archive, preimages, evmRedenomCfg.BatchSize)
	if err != nil {
		return errors.Wrap(err, "get EVM balances")
	}

	// Calculate total supply
	totalSupply := bi.Zero()
	for _, bal := range balances {
		totalSupply.Add(totalSupply, bal.Balance)
	}

	// Prepare output structure
	output := map[string]any{
		"block_number":        blockNumber,
		"state_root":          stateRoot.Hex(),
		"total_supply":        totalSupply.String(),
		"total_supply_pretty": FormatBalance(totalSupply),
		"accounts":            len(balances),
		"balances":            balances,
	}

	// Write to file
	data, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		return errors.Wrap(err, "marshal JSON")
	}

	if err := os.WriteFile(outputPath, data, 0o644); err != nil {
		return errors.Wrap(err, "write file")
	}

	log.Info(ctx, "EVM balances written",
		"path", outputPath,
		"accounts", len(balances),
		"total_supply", FormatBalance(totalSupply),
	)

	return nil
}

// snapshotStakingBalances fetches staking balances from consensus and writes to JSON.
func snapshotStakingBalances(ctx context.Context, cprov cchain.Provider, outputPath string) error {
	log.Info(ctx, "Fetching staking balances from consensus chain")

	// Call GetStakingBalances
	stakes, err := GetStakingBalances(ctx, cprov)
	if err != nil {
		return errors.Wrap(err, "get staking balances")
	}

	// Calculate totals
	totalStake := bi.Zero()
	totalDelegation := bi.Zero()
	totalUnbonding := bi.Zero()
	totalRewards := bi.Zero()

	for _, stake := range stakes {
		totalStake.Add(totalStake, stake.Total)
		totalDelegation.Add(totalDelegation, stake.Delegation)
		totalUnbonding.Add(totalUnbonding, stake.Unbonding)
		totalRewards.Add(totalRewards, stake.Rewards)
	}

	// Prepare output structure
	output := map[string]any{
		"total_stake":        totalStake.String(),
		"total_stake_pretty": FormatBalance(totalStake),
		"total_delegation":   FormatBalance(totalDelegation),
		"total_unbonding":    FormatBalance(totalUnbonding),
		"total_rewards":      FormatBalance(totalRewards),
		"total_delegators":   len(stakes),
		"delegators":         stakes,
	}

	// Write to file
	data, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		return errors.Wrap(err, "marshal JSON")
	}

	if err := os.WriteFile(outputPath, data, 0o644); err != nil {
		return errors.Wrap(err, "write file")
	}

	log.Info(ctx, "Staking balances written",
		"path", outputPath,
		"delegators", len(stakes),
		"total_stake", FormatBalance(totalStake),
	)

	return nil
}

// loadGenesisPreimages loads address preimages from EVM genesis file.
func loadGenesisPreimages(genesisPath string) (map[common.Hash]common.Address, error) {
	// Read genesis file
	bz, err := os.ReadFile(genesisPath)
	if err != nil {
		return nil, errors.Wrap(err, "read genesis file")
	}

	// Parse genesis allocs
	var allocs struct {
		Alloc map[string]json.RawMessage `json:"alloc"`
	}
	if err := json.Unmarshal(bz, &allocs); err != nil {
		return nil, errors.Wrap(err, "unmarshal genesis")
	}

	// Build preimages map: keccak256(address) -> address
	preimages := make(map[common.Hash]common.Address, len(allocs.Alloc))
	for addrHex := range allocs.Alloc {
		addrBz, err := hex.DecodeString(addrHex)
		if err != nil {
			return nil, errors.Wrap(err, "decode address", "hex", addrHex)
		}
		addr, err := cast.EthAddress(addrBz)
		if err != nil {
			return nil, err
		}
		preimages[crypto.Keccak256Hash(addr[:])] = addr
	}

	return preimages, nil
}
