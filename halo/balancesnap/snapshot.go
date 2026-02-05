package balancesnap

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"math/big"
	"os"
	"path/filepath"
	"time"

	evmredenomsubmit "github.com/omni-network/omni/halo/evmredenom/submit"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethp2p"
	"github.com/omni-network/omni/lib/log"

	"github.com/cometbft/cometbft/rpc/client"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p/enode"
)

// Start starts a goroutine that waits for halt height and snapshots balances.
// Returns immediately if not enabled or halt height is 0.
func Start(
	ctx context.Context,
	haltHeight uint64,
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

	// Skip if halt height is disabled
	if haltHeight == 0 {
		return
	}

	go func() {
		if err := run(ctx, haltHeight, evmRedenomCfg, homeDir, consensusClient, cprov); err != nil {
			log.Error(ctx, "BalanceSnap: balance snapshot failed", err)
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
	log.Info(ctx, "BalanceSnap: waiting for halt height", "halt_height", haltHeight)

	// Wait for consensus to reach halt height
	if err := waitForHeight(ctx, consensusClient, haltHeight); err != nil {
		return errors.Wrap(err, "wait for halt height")
	}

	log.Info(ctx, "BalanceSnap: halt height reached, querying EVM head", "height", haltHeight)

	// Get ExecutionHead from cprovider to verify consensus state
	execHead, err := cprov.ExecutionHead(ctx)
	if err != nil {
		return errors.Wrap(err, "get execution head from consensus")
	}

	// Connect to archive RPC to get full header with state root
	archive, err := ethclient.DialContext(ctx, "omni_evm", evmRedenomCfg.RPCArchive)
	if err != nil {
		return errors.Wrap(err, "dial archive RPC")
	}
	defer archive.Close()

	// Query EVM header by block number to get state root
	blockNum := new(big.Int).SetUint64(execHead.BlockNumber)
	evmHeader, err := archive.HeaderByNumber(ctx, blockNum)
	if err != nil {
		return errors.Wrap(err, "get EVM header from archive")
	}

	// Verify archive header matches ExecutionHead
	if evmHeader.Number.Uint64() != execHead.BlockNumber || evmHeader.Hash() != execHead.BlockHash {
		return errors.New("archive header mismatch with execution head",
			"archive_number", evmHeader.Number.Uint64(),
			"exec_number", execHead.BlockNumber,
			"archive_hash", evmHeader.Hash().Hex(),
			"exec_hash", execHead.BlockHash.Hex(),
		)
	}

	log.Info(ctx, "BalanceSnap: evm head retrieved",
		"block_number", evmHeader.Number.Uint64(),
		"state_root", evmHeader.Root.Hex(),
	)

	// Hardcode output paths
	evmOutputPath := filepath.Join(homeDir, "data", "evm_balances_halt.json")
	stakeOutputPath := filepath.Join(homeDir, "data", "staking_balances_halt.json")

	// Snapshot EVM balances
	log.Info(ctx, "BalanceSnap: snapshotting EVM balances", "output", evmOutputPath)
	if err := snapshotEVMBalances(ctx, evmRedenomCfg, evmHeader.Root, evmHeader.Number.Uint64(), haltHeight, evmOutputPath); err != nil {
		return errors.Wrap(err, "snapshot EVM balances")
	}

	// Snapshot staking balances
	log.Info(ctx, "BalanceSnap: snapshotting staking balances", "output", stakeOutputPath)
	if err := snapshotStakingBalances(ctx, cprov, haltHeight, stakeOutputPath); err != nil {
		return errors.Wrap(err, "snapshot staking balances")
	}

	// Ensure height isn't increasing
	if err := waitForHeight(ctx, consensusClient, haltHeight); err != nil {
		return errors.Wrap(err, "wait for halt height")
	}

	log.Info(ctx, "BalanceSnap: balance snapshot completed successfully")

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
				log.Warn(ctx, "BalanceSnap: failed to get consensus status (will retry)", err)
				continue
			}

			if status.SyncInfo.LatestBlockHeight < 0 {
				log.Warn(ctx, "BalanceSnap: invalid block height (will retry)", nil, "height", status.SyncInfo.LatestBlockHeight)
				continue
			}

			currentHeight := uint64(status.SyncInfo.LatestBlockHeight) //nolint:gosec // Validated >= 0 above

			// Error if consensus height is above halt height
			if currentHeight > targetHeight {
				return errors.New("consensus height exceeded halt height",
					"current_height", currentHeight,
					"halt_height", targetHeight,
				)
			}

			if currentHeight >= targetHeight {
				return nil
			}

			if currentHeight%100 == 0 {
				log.Debug(ctx, "BalanceSnap: waiting for halt height",
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
	consensusHeight uint64,
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

	log.Info(ctx, "BalanceSnap: fetching EVM balances via snap protocol",
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
		"consensus_height":    consensusHeight,
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

	log.Info(ctx, "BalanceSnap: evm balances written",
		"path", outputPath,
		"accounts", len(balances),
		"total_supply", FormatBalance(totalSupply),
	)

	return nil
}

// snapshotStakingBalances fetches staking balances from consensus and writes to JSON.
func snapshotStakingBalances(ctx context.Context, cprov cchain.Provider, consensusHeight uint64, outputPath string) error {
	log.Info(ctx, "BalanceSnap: fetching staking balances from consensus chain")

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
		"consensus_height":   consensusHeight,
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

	log.Info(ctx, "BalanceSnap: staking balances written",
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
