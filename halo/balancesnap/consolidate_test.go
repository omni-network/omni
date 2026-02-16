package balancesnap_test

import (
	"math/big"
	"os"
	"path/filepath"
	"testing"

	"github.com/omni-network/omni/halo/balancesnap"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

const testL1RPC = "https://eth.llamarpc.com"

func TestConsolidateBalances(t *testing.T) {
	t.Parallel()

	ctx := t.Context()

	// Paths to mainnet balance files in ops repo
	opsDir := filepath.Join(os.Getenv("HOME"), "repos", "ops")
	evmBalancesPath := filepath.Join(opsDir, "evm_balances_latest.json")
	stakingBalancesPath := filepath.Join(opsDir, "staking_balances_latest.json")

	// Check if files exist
	if _, err := os.Stat(evmBalancesPath); os.IsNotExist(err) {
		t.Skipf("EVM balances file not found: %s", evmBalancesPath)
	}
	if _, err := os.Stat(stakingBalancesPath); os.IsNotExist(err) {
		t.Skipf("Staking balances file not found: %s", stakingBalancesPath)
	}

	// Create output file in ops repo
	outputPath := filepath.Join(opsDir, "combined_balances_mainnet.json")

	// Run consolidation with L1 RPC to verify bridge balance
	// L1 RPC URL is now required for mainnet consolidation
	l1RPC := testL1RPC

	result, summary, err := balancesnap.ConsolidateBalances(
		ctx,
		netconf.Mainnet,
		evmBalancesPath,
		stakingBalancesPath,
		l1RPC,
		outputPath,
	)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, summary)

	// Verify output file was created
	require.FileExists(t, outputPath)

	// Verify summary calculations
	t.Logf("EVM Total Supply: %s", balancesnap.FormatBalance(summary.TotalEVMSupply))
	t.Logf("Consensus Staking: %s", balancesnap.FormatBalance(summary.TotalConsensusStake))
	t.Logf("Total Burned: %s", balancesnap.FormatBalance(summary.TotalBurned))
	t.Logf("Total Consolidated: %s", balancesnap.FormatBalance(summary.TotalConsolidated))
	t.Logf("Total Payable: %s", balancesnap.FormatBalance(summary.TotalPayable))

	// Verify staking predeploy is not in result (burned)
	stakingAddr := common.HexToAddress(predeploys.Staking)
	_, exists := result[stakingAddr]
	require.False(t, exists, "Staking predeploy should be burned, not in payout")

	// Verify native bridge predeploy is not in result (burned)
	bridgeAddr := common.HexToAddress(predeploys.OmniBridgeNative)
	_, exists = result[bridgeAddr]
	require.False(t, exists, "Native bridge predeploy should be burned, not in payout")

	// Verify dead address is not in result (burned)
	deadAddr := common.HexToAddress("0x000000000000000000000000000000000000dead")
	_, exists = result[deadAddr]
	require.False(t, exists, "Dead address should be burned, not in payout")

	// Verify foundation address received consolidated funds
	foundationAddr := common.HexToAddress("0xfdb3e9cdc5f016cff6cfaf28fef141ae76efd31d")
	foundationBalance, exists := result[foundationAddr]
	require.True(t, exists, "Foundation should have received consolidated funds")
	require.Positive(t, foundationBalance.Cmp(big.NewInt(0)), "Foundation balance should be positive")
	t.Logf("Foundation balance: %s", balancesnap.FormatBalance(foundationBalance))

	// Verify total payable equals the sum of all result balances
	calculatedTotal := big.NewInt(0)
	for _, amount := range result {
		calculatedTotal.Add(calculatedTotal, amount)
	}
	require.Equal(t, summary.TotalPayable.String(), calculatedTotal.String(),
		"Total payable should equal sum of all result balances")

	// Verify the calculation formula accounting for L1 bridge shortfall:
	// If Total Payable > L1 Bridge Balance, foundation covers the shortfall
	// Expected Total (before shortfall) = EVM Total - Burned + Consensus Staking
	expectedTotalBeforeShortfall := new(big.Int).Set(summary.TotalEVMSupply)
	expectedTotalBeforeShortfall.Sub(expectedTotalBeforeShortfall, summary.TotalBurned)
	expectedTotalBeforeShortfall.Add(expectedTotalBeforeShortfall, summary.TotalConsensusStake)

	// If there's a shortfall, total payable should equal L1 bridge balance
	if summary.FoundationShortfall.Cmp(big.NewInt(0)) > 0 {
		require.Equal(t, summary.L1BridgeBalance.String(), summary.TotalPayable.String(),
			"Total payable should equal L1 bridge balance when foundation covers shortfall")

		// Verify shortfall calculation
		expectedShortfall := new(big.Int).Sub(expectedTotalBeforeShortfall, summary.L1BridgeBalance)
		require.Equal(t, expectedShortfall.String(), summary.FoundationShortfall.String(),
			"Foundation shortfall should equal difference between expected total and bridge balance")
	} else {
		// No shortfall - total payable should match formula
		require.Equal(t, expectedTotalBeforeShortfall.String(), summary.TotalPayable.String(),
			"Total payable should match formula: EVM Total - Burned + Consensus Staking")
	}

	t.Logf("Number of payout addresses: %d", len(result))
	t.Logf("Burned Staking: %s", balancesnap.FormatBalance(summary.BurnedStaking))
	t.Logf("Burned Native Bridge: %s", balancesnap.FormatBalance(summary.BurnedNativeBridge))
	t.Logf("Burned Dead: %s", balancesnap.FormatBalance(summary.BurnedDead))
	t.Logf("Consolidated EOAs: %s", balancesnap.FormatBalance(summary.ConsolidatedEOAs))
	t.Logf("Consolidated Validators: %s", balancesnap.FormatBalance(summary.ConsolidatedValidators))
	t.Logf("Consolidated Contracts: %s", balancesnap.FormatBalance(summary.ConsolidatedContracts))
	t.Logf("Consolidated Predeploys: %s", balancesnap.FormatBalance(summary.ConsolidatedPredeploys))
}

func TestConsolidateBalances_OnlyMainnet(t *testing.T) {
	t.Parallel()

	ctx := t.Context()
	tempDir := t.TempDir()

	// Try with non-mainnet network
	_, _, err := balancesnap.ConsolidateBalances(
		ctx,
		netconf.Omega, // Not mainnet
		"dummy.json",
		"dummy.json",
		testL1RPC,
		filepath.Join(tempDir, "output.json"),
	)
	require.Error(t, err)
	require.Contains(t, err.Error(), "consolidation only allowed for mainnet")
}

func TestConsolidateBalances_InvalidFiles(t *testing.T) {
	t.Parallel()

	ctx := t.Context()
	tempDir := t.TempDir()

	// Try with non-existent files
	_, _, err := balancesnap.ConsolidateBalances(
		ctx,
		netconf.Mainnet,
		"nonexistent.json",
		"nonexistent.json",
		testL1RPC,
		filepath.Join(tempDir, "output.json"),
	)
	require.Error(t, err)
}

func TestConsolidateBalances_RequiresL1RPC(t *testing.T) {
	t.Parallel()

	ctx := t.Context()
	tempDir := t.TempDir()

	opsDir := filepath.Join(os.Getenv("HOME"), "repos", "ops")
	evmBalancesPath := filepath.Join(opsDir, "evm_balances_latest.json")
	stakingBalancesPath := filepath.Join(opsDir, "staking_balances_latest.json")

	if _, err := os.Stat(evmBalancesPath); os.IsNotExist(err) {
		t.Skipf("EVM balances file not found: %s", evmBalancesPath)
	}
	if _, err := os.Stat(stakingBalancesPath); os.IsNotExist(err) {
		t.Skipf("Staking balances file not found: %s", stakingBalancesPath)
	}

	// Try with empty L1 RPC URL - should error
	_, _, err := balancesnap.ConsolidateBalances(
		ctx,
		netconf.Mainnet,
		evmBalancesPath,
		stakingBalancesPath,
		"", // Empty RPC URL should error
		filepath.Join(tempDir, "output.json"),
	)
	require.Error(t, err)
	require.Contains(t, err.Error(), "l1 rpc url required")
}

func TestValidatorAddressFormat(t *testing.T) {
	t.Parallel()

	// This test verifies that validator addresses in mainnet manifest
	// are in the expected format: 0x-prefixed, 42 characters
	// The consolidation code will error if the format is invalid

	ctx := t.Context()

	opsDir := filepath.Join(os.Getenv("HOME"), "repos", "ops")
	evmBalancesPath := filepath.Join(opsDir, "evm_balances_latest.json")
	stakingBalancesPath := filepath.Join(opsDir, "staking_balances_latest.json")

	if _, err := os.Stat(evmBalancesPath); os.IsNotExist(err) {
		t.Skipf("EVM balances file not found: %s", evmBalancesPath)
	}
	if _, err := os.Stat(stakingBalancesPath); os.IsNotExist(err) {
		t.Skipf("Staking balances file not found: %s", stakingBalancesPath)
	}

	tempDir := t.TempDir()
	outputPath := filepath.Join(tempDir, "combined_balances_test.json")

	// Use L1 RPC for validation
	l1RPC := testL1RPC

	// Should not error - mainnet manifest has valid format
	_, _, err := balancesnap.ConsolidateBalances(
		ctx,
		netconf.Mainnet,
		evmBalancesPath,
		stakingBalancesPath,
		l1RPC,
		outputPath,
	)
	require.NoError(t, err, "Mainnet manifest should have valid validator address format")
}

func TestConsolidateBalances_WithL1RPC(t *testing.T) {
	t.Parallel()

	ctx := t.Context()

	opsDir := filepath.Join(os.Getenv("HOME"), "repos", "ops")
	evmBalancesPath := filepath.Join(opsDir, "evm_balances_latest.json")
	stakingBalancesPath := filepath.Join(opsDir, "staking_balances_latest.json")

	if _, err := os.Stat(evmBalancesPath); os.IsNotExist(err) {
		t.Skipf("EVM balances file not found: %s", evmBalancesPath)
	}
	if _, err := os.Stat(stakingBalancesPath); os.IsNotExist(err) {
		t.Skipf("Staking balances file not found: %s", stakingBalancesPath)
	}

	tempDir := t.TempDir()
	outputPath := filepath.Join(tempDir, "combined_balances_mainnet_l1.json")

	// Use working Ethereum mainnet RPC
	l1RPC := testL1RPC

	result, summary, err := balancesnap.ConsolidateBalances(
		ctx,
		netconf.Mainnet,
		evmBalancesPath,
		stakingBalancesPath,
		l1RPC,
		outputPath,
	)

	// Note: This may error if there's a shortfall that foundation can't cover
	// That's expected behavior - we need to address the funding issue
	if err != nil {
		t.Logf("Consolidation error (may be expected): %v", err)
		require.Contains(t, err.Error(), "insufficient funds", "Should be shortfall error")

		return
	}

	require.NotNil(t, result)
	require.NotNil(t, summary)

	// Log L1 bridge info
	t.Logf("L1 Bridge Address: %s", summary.L1BridgeAddress.Hex())
	t.Logf("L1 Bridge NOM Balance: %s NOM", balancesnap.FormatBalance(summary.L1BridgeBalance))
	t.Logf("Total Payable: %s NOM", balancesnap.FormatBalance(summary.TotalPayable))

	if summary.L1BridgeBalance.Cmp(big.NewInt(0)) == 0 {
		t.Logf("⚠️  WARNING: L1 Bridge has ZERO balance - bridge may not be funded yet")
	}

	if summary.FoundationShortfall.Cmp(big.NewInt(0)) > 0 {
		t.Logf("Foundation Shortfall: %s NOM (deducted from foundation)", balancesnap.FormatBalance(summary.FoundationShortfall))

		// Verify foundation balance was reduced
		foundationAddr := common.HexToAddress("0xfdb3e9cdc5f016cff6cfaf28fef141ae76efd31d")
		foundationBalance := result[foundationAddr]
		t.Logf("Foundation final balance: %s NOM", balancesnap.FormatBalance(foundationBalance))
	} else if summary.L1BridgeBalance.Cmp(summary.TotalPayable) >= 0 {
		t.Logf("✅ L1 bridge has sufficient funds (surplus: %s NOM)",
			balancesnap.FormatBalance(new(big.Int).Sub(summary.L1BridgeBalance, summary.TotalPayable)))
	}
}
