package balancesnap

import (
	"context"
	"encoding/hex"
	"encoding/json"
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
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p/enode"
)

// Start starts a goroutine that snapshots balances.
// Returns immediately if not enabled.
// If haltHeight is 0, snapshots at latest height immediately.
// If haltHeight is set, waits for halt height before snapshotting.
func Start(
	ctx context.Context,
	network netconf.ID,
	haltHeight uint64,
	evmRedenomCfg evmredenomsubmit.Config,
	homeDir string,
	cprov cchain.Provider,
	rpcEndpoints xchain.RPCEndpoints,
	_ chan<- error,
) {
	// Only enabled if evmredenom is enabled
	if !evmRedenomCfg.Enabled {
		return
	}

	go func() {
		err := expbackoff.Retry(ctx,
			func() error {
				return run(ctx, network, haltHeight, evmRedenomCfg, homeDir, cprov, rpcEndpoints)
			},
			expbackoff.WithRetryLabel("BalanceSnap"),
			expbackoff.WithRetryCount(10),
			expbackoff.WithPeriodicConfig(time.Minute), // Mitigate geth 30s inboundThrottleTime
		)
		if err != nil {
			log.Error(ctx, "BalanceSnap: balance snapshot failed", err)
			// Don't abort the app, just log the error
		}
	}()
}

// run executes the balance snapshot workflow.
func run(
	ctx context.Context,
	network netconf.ID,
	haltHeight uint64,
	evmRedenomCfg evmredenomsubmit.Config,
	homeDir string,
	cprov cchain.Provider,
	rpcEndpoints xchain.RPCEndpoints,
) error {
	var consensusHeight uint64
	var outputSuffix string

	if haltHeight == 0 {
		// No halt height configured, snapshot at latest height
		log.Info(ctx, "BalanceSnap: no halt height configured, snapshotting at latest height")

		status, err := cprov.NodeStatus(ctx)
		if err != nil {
			return errors.Wrap(err, "get consensus status")
		}

		consensusHeight = status.Height
		outputSuffix = "latest.json"

		log.Info(ctx, "BalanceSnap: snapshotting at latest consensus height", "height", consensusHeight)
	} else {
		// Halt height configured, wait for it
		log.Info(ctx, "BalanceSnap: waiting for halt height", "halt_height", haltHeight)

		if err := waitForHeight(ctx, cprov, haltHeight); err != nil {
			return errors.Wrap(err, "wait for halt height")
		}

		consensusHeight = haltHeight
		outputSuffix = "halt.json"

		log.Info(ctx, "BalanceSnap: halt height reached, querying EVM head", "height", haltHeight)
	}

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
	evmHeader, err := archive.HeaderByHash(ctx, execHead.BlockHash)
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

	// Output paths based on whether we're at halt or latest
	evmOutputPath := filepath.Join(homeDir, "data", "evm_balances_"+outputSuffix)
	stakeOutputPath := filepath.Join(homeDir, "data", "staking_balances_"+outputSuffix)

	// Snapshot EVM balances
	log.Info(ctx, "BalanceSnap: snapshotting EVM balances", "output", evmOutputPath)
	if err := snapshotEVMBalances(ctx, evmRedenomCfg, evmHeader.Root, evmHeader.Number.Uint64(), consensusHeight, evmOutputPath); err != nil {
		return errors.Wrap(err, "snapshot EVM balances")
	}

	// Snapshot staking balances
	log.Info(ctx, "BalanceSnap: snapshotting staking balances", "output", stakeOutputPath)
	if err := snapshotStakingBalances(ctx, cprov, consensusHeight, stakeOutputPath); err != nil {
		return errors.Wrap(err, "snapshot staking balances")
	}

	// Only verify height isn't increasing if we're expecting halt
	if haltHeight != 0 {
		if err := waitForHeight(ctx, cprov, haltHeight); err != nil {
			return errors.Wrap(err, "verify halt height")
		}
	}

	log.Info(ctx, "BalanceSnap: balance snapshot completed successfully",
		"consensus_height", consensusHeight,
		"evm_block", evmHeader.Number.Uint64(),
	)

	// Only consolidate balances on mainnet
	if network != netconf.Mainnet {
		log.Warn(ctx, "BalanceSnap: skipping consolidation (only runs on mainnet)", nil, "network", network)

		return nil
	}

	// Get L1 RPC from endpoints
	l1ChainID, ok := netconf.EthereumChainID(network)
	if !ok {
		log.Warn(ctx, "BalanceSnap: skipping consolidation (no L1 chain for network)", nil, "network", network)

		return nil
	}

	l1RPC, err := rpcEndpoints.ByNameOrID("ethereum", l1ChainID)
	if err != nil {
		log.Warn(ctx, "BalanceSnap: skipping consolidation (no L1 RPC endpoint configured)", err)

		return nil
	}

	consolidatedOutputPath := filepath.Join(homeDir, "data", "combined_balances_"+outputSuffix)

	log.Info(ctx, "BalanceSnap: consolidating balances", "output", consolidatedOutputPath, "l1_rpc", l1RPC)

	_, summary, err := ConsolidateBalances(
		ctx,
		network,
		evmOutputPath,
		stakeOutputPath,
		l1RPC,
		consolidatedOutputPath,
	)
	if err != nil {
		// Don't fail the snapshot if consolidation fails - just log the error
		log.Warn(ctx, "BalanceSnap: consolidation failed (snapshot still succeeded)", err)

		return nil
	}

	log.Info(ctx, "BalanceSnap: consolidation completed",
		"total_payable", FormatBalance(summary.TotalPayable),
		"l1_bridge_balance", FormatBalance(summary.L1BridgeBalance),
		"foundation_shortfall", FormatBalance(summary.FoundationShortfall),
		"output", consolidatedOutputPath,
	)

	return nil
}

// waitForHeight polls consensus provider until target height is reached.
func waitForHeight(ctx context.Context, cprov cchain.Provider, targetHeight uint64) error {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return errors.Wrap(ctx.Err(), "context done")
		case <-ticker.C:
			status, err := cprov.NodeStatus(ctx)
			if err != nil {
				log.Warn(ctx, "BalanceSnap: failed to get consensus status (will retry)", err)
				continue
			}

			currentHeight := status.Height

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
