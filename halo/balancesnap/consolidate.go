package balancesnap

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"

	"github.com/omni-network/omni/contracts/allocs"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/e2e/app/key"
	"github.com/omni-network/omni/e2e/manifests"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tokens/tokenutil"

	"github.com/ethereum/go-ethereum/common"
)

// ConsolidationSummary contains statistics about the consolidation process.
type ConsolidationSummary struct {
	TotalEVMSupply         *big.Int
	TotalConsensusStake    *big.Int
	BurnedStaking          *big.Int
	BurnedNativeBridge     *big.Int
	BurnedDead             *big.Int
	ConsolidatedEOAs       *big.Int
	ConsolidatedContracts  *big.Int
	ConsolidatedPredeploys *big.Int
	ConsolidatedValidators *big.Int
	TotalBurned            *big.Int
	TotalConsolidated      *big.Int
	TotalPayable           *big.Int
	L1BridgeBalance        *big.Int
	L1BridgeAddress        common.Address
	FoundationShortfall    *big.Int
	FoundationAddress      common.Address
}

// ConsolidateBalances consolidates EVM and consensus balances for mainnet L1 payout.
//
// It performs the following:
// - Burns staking and native bridge predeploy balances
// - Burns 0xdead address
// - Sweeps EOA role accounts to foundation address
// - Sweeps validator addresses to foundation address
// - Sweeps contract addresses to foundation address
// - Consolidates remaining predeploys to foundation address
// - Adds consensus staking balances
// - Queries L1 bridge balance and ensures sufficient funds
// - Deducts shortfall from foundation if total exceeds L1 bridge balance
//
// Returns a map of address -> wei amount representing L1 payouts.
func ConsolidateBalances(
	ctx context.Context,
	network netconf.ID,
	evmBalancesPath string,
	stakingBalancesPath string,
	foundationAddr common.Address,
	l1RPCURL string,
	outputPath string,
) (map[common.Address]*big.Int, *ConsolidationSummary, error) {
	// Only allow mainnet
	if network != netconf.Mainnet {
		return nil, nil, errors.New("consolidation only allowed for mainnet")
	}

	log.Info(ctx, "Starting balance consolidation", "network", network, "foundation", foundationAddr.Hex())

	// Load balance files
	evmBalances, err := loadEVMBalances(evmBalancesPath)
	if err != nil {
		return nil, nil, errors.Wrap(err, "load evm balances")
	}

	stakingBalances, err := loadStakingBalances(stakingBalancesPath)
	if err != nil {
		return nil, nil, errors.Wrap(err, "load staking balances")
	}

	// Initialize summary
	summary := &ConsolidationSummary{
		TotalEVMSupply:         bi.Zero(),
		TotalConsensusStake:    stakingBalances.TotalStake,
		BurnedStaking:          bi.Zero(),
		BurnedNativeBridge:     bi.Zero(),
		BurnedDead:             bi.Zero(),
		ConsolidatedEOAs:       bi.Zero(),
		ConsolidatedContracts:  bi.Zero(),
		ConsolidatedPredeploys: bi.Zero(),
		ConsolidatedValidators: bi.Zero(),
		TotalBurned:            bi.Zero(),
		TotalConsolidated:      bi.Zero(),
		TotalPayable:           bi.Zero(),
		L1BridgeBalance:        bi.Zero(),
		FoundationShortfall:    bi.Zero(),
		FoundationAddress:      foundationAddr,
	}

	// Calculate total EVM supply
	for _, bal := range evmBalances.Balances {
		summary.TotalEVMSupply.Add(summary.TotalEVMSupply, bal.Balance)
	}

	// Get all internal addresses to consolidate
	internalAddrs, err := getInternalAddresses(network)
	if err != nil {
		return nil, nil, errors.Wrap(err, "get internal addresses")
	}

	// Build result map
	result := make(map[common.Address]*big.Int)

	// Process each EVM balance
	for _, bal := range evmBalances.Balances {
		addr := common.HexToAddress(bal.Address)
		amount := new(big.Int).Set(bal.Balance)

		// Check if this is an address to consolidate or burn
		action := categorizeAddress(addr, internalAddrs)

		switch action {
		case actionBurnStaking:
			summary.BurnedStaking.Add(summary.BurnedStaking, amount)
			summary.TotalBurned.Add(summary.TotalBurned, amount)
			log.Debug(ctx, "Burning staking predeploy", "address", addr.Hex(), "amount", FormatBalance(amount))

		case actionBurnNativeBridge:
			summary.BurnedNativeBridge.Add(summary.BurnedNativeBridge, amount)
			summary.TotalBurned.Add(summary.TotalBurned, amount)
			log.Debug(ctx, "Burning native bridge predeploy", "address", addr.Hex(), "amount", FormatBalance(amount))

		case actionBurnDead:
			summary.BurnedDead.Add(summary.BurnedDead, amount)
			summary.TotalBurned.Add(summary.TotalBurned, amount)
			log.Debug(ctx, "Burning dead address", "address", addr.Hex(), "amount", FormatBalance(amount))

		case actionConsolidateEOA:
			summary.ConsolidatedEOAs.Add(summary.ConsolidatedEOAs, amount)
			summary.TotalConsolidated.Add(summary.TotalConsolidated, amount)
			addToResult(result, foundationAddr, amount)
			log.Debug(ctx, "Consolidating EOA", "address", addr.Hex(), "amount", FormatBalance(amount))

		case actionConsolidateContract:
			summary.ConsolidatedContracts.Add(summary.ConsolidatedContracts, amount)
			summary.TotalConsolidated.Add(summary.TotalConsolidated, amount)
			addToResult(result, foundationAddr, amount)
			log.Debug(ctx, "Consolidating contract", "address", addr.Hex(), "amount", FormatBalance(amount))

		case actionConsolidatePredeploy:
			summary.ConsolidatedPredeploys.Add(summary.ConsolidatedPredeploys, amount)
			summary.TotalConsolidated.Add(summary.TotalConsolidated, amount)
			addToResult(result, foundationAddr, amount)
			log.Debug(ctx, "Consolidating predeploy", "address", addr.Hex(), "amount", FormatBalance(amount))

		case actionConsolidateValidator:
			summary.ConsolidatedValidators.Add(summary.ConsolidatedValidators, amount)
			summary.TotalConsolidated.Add(summary.TotalConsolidated, amount)
			addToResult(result, foundationAddr, amount)
			log.Debug(ctx, "Consolidating validator", "address", addr.Hex(), "amount", FormatBalance(amount))

		case actionKeep:
			// Regular user balance, keep it
			addToResult(result, addr, amount)
		}
	}

	// Add consensus staking balances
	// Check if they belong to validators and consolidate to foundation if so
	for _, stake := range stakingBalances.Delegators {
		addr := common.HexToAddress(stake.Address)

		// Check if this is a validator address
		if internalAddrs.validatorAddrs[addr] {
			// Consolidate validator staking to foundation
			summary.ConsolidatedValidators.Add(summary.ConsolidatedValidators, stake.Total)
			summary.TotalConsolidated.Add(summary.TotalConsolidated, stake.Total)
			addToResult(result, foundationAddr, stake.Total)
			log.Debug(ctx, "Consolidating validator staking", "address", addr.Hex(), "amount", FormatBalance(stake.Total))
		} else {
			// Regular delegator, keep their staking
			addToResult(result, addr, stake.Total)
		}
	}

	// Calculate total payable
	for _, amount := range result {
		summary.TotalPayable.Add(summary.TotalPayable, amount)
	}

	// Get L1 bridge address from contracts package
	contractAddrs, err := contracts.GetAddresses(ctx, network)
	if err != nil {
		return nil, nil, errors.Wrap(err, "get contract addresses")
	}

	l1BridgeAddr := contractAddrs.L1Bridge
	summary.L1BridgeAddress = l1BridgeAddr

	if l1RPCURL != "" {
		l1BridgeBalance, err := queryL1BridgeBalance(ctx, l1RPCURL, l1BridgeAddr)
		if err != nil {
			return nil, nil, errors.Wrap(err, "query L1 bridge balance")
		}
		summary.L1BridgeBalance = l1BridgeBalance

		log.Info(ctx, "L1 bridge NOM token balance queried",
			"bridge_address", l1BridgeAddr.Hex(),
			"balance", FormatBalance(l1BridgeBalance),
		)
	} else {
		log.Warn(ctx, "L1 RPC URL not provided, skipping L1 bridge balance check", nil)
	}

	// Check if we have enough funds on L1 (only if L1 balance was queried)
	if l1RPCURL != "" && summary.L1BridgeBalance.Cmp(bi.Zero()) > 0 && summary.TotalPayable.Cmp(summary.L1BridgeBalance) > 0 {
		// Calculate shortfall
		shortfall := new(big.Int).Sub(summary.TotalPayable, summary.L1BridgeBalance)
		summary.FoundationShortfall = shortfall

		log.Warn(ctx, "Total payable exceeds L1 bridge balance", nil,
			"total_payable", FormatBalance(summary.TotalPayable),
			"l1_bridge", FormatBalance(summary.L1BridgeBalance),
			"shortfall", FormatBalance(shortfall),
		)

		// Check if foundation has enough to cover shortfall
		foundationBalance, exists := result[foundationAddr]
		if !exists || foundationBalance.Cmp(shortfall) < 0 {
			foundationBal := bi.Zero()
			if exists {
				foundationBal = foundationBalance
			}
			return nil, nil, errors.New("insufficient funds: foundation cannot cover shortfall",
				"shortfall", FormatBalance(shortfall),
				"foundation_balance", FormatBalance(foundationBal),
			)
		}

		// Deduct shortfall from foundation
		log.Info(ctx, "Deducting shortfall from foundation",
			"foundation", foundationAddr.Hex(),
			"before", FormatBalance(foundationBalance),
			"shortfall", FormatBalance(shortfall),
		)

		foundationBalance.Sub(foundationBalance, shortfall)

		log.Info(ctx, "Foundation balance after shortfall deduction",
			"after", FormatBalance(foundationBalance),
		)

		// Recalculate total payable
		summary.TotalPayable = bi.Zero()
		for _, amount := range result {
			summary.TotalPayable.Add(summary.TotalPayable, amount)
		}
	} else {
		summary.FoundationShortfall = bi.Zero()
	}

	// Print summary
	printConsolidationSummary(ctx, summary)

	// Write output file
	if err := writeConsolidatedBalances(outputPath, result); err != nil {
		return nil, nil, errors.Wrap(err, "write consolidated balances")
	}

	log.Info(ctx, "Balance consolidation completed", "output", outputPath, "addresses", len(result))

	return result, summary, nil
}

type action int

const (
	actionKeep action = iota
	actionBurnStaking
	actionBurnNativeBridge
	actionBurnDead
	actionConsolidateEOA
	actionConsolidateContract
	actionConsolidatePredeploy
	actionConsolidateValidator
)

type internalAddresses struct {
	eoaAddrs       map[common.Address]bool
	contractAddrs  map[common.Address]bool
	validatorAddrs map[common.Address]bool
	predeployAddrs map[common.Address]bool
}

// getInternalAddresses returns all internal addresses that should be consolidated or burned.
func getInternalAddresses(network netconf.ID) (*internalAddresses, error) {
	addrs := &internalAddresses{
		eoaAddrs:       make(map[common.Address]bool),
		contractAddrs:  make(map[common.Address]bool),
		validatorAddrs: make(map[common.Address]bool),
		predeployAddrs: make(map[common.Address]bool),
	}

	// Get EOA addresses
	for _, account := range eoa.AllAccounts(network) {
		addrs.eoaAddrs[account.Address] = true
	}

	// Get contract addresses from genesis allocs
	genesisAllocs, err := allocs.Alloc(network)
	if err == nil {
		for addr := range genesisAllocs {
			addrs.contractAddrs[addr] = true
		}
	}

	// Get validator addresses from manifest
	manifest, err := manifests.Manifest(network)
	if err != nil {
		return nil, errors.Wrap(err, "load manifest")
	}

	for nodeName, keys := range manifest.Keys {
		if valAddr, ok := keys[key.Validator]; ok {
			// Validator keys in manifest are stored as addresses (0x-prefixed, 42 chars)
			if len(valAddr) != 42 || valAddr[:2] != "0x" {
				return nil, errors.New("invalid validator address format in manifest",
					"node", nodeName,
					"value", valAddr,
					"expected", "0x-prefixed address (42 chars)",
				)
			}
			addrs.validatorAddrs[common.HexToAddress(valAddr)] = true
		}
	}

	// Add predeploy addresses
	predeployList := []string{
		predeploys.PortalRegistry,
		predeploys.OmniBridgeNative,
		predeploys.WOmni,
		predeploys.WNomina,
		predeploys.Staking,
		predeploys.Slashing,
		predeploys.Upgrade,
		predeploys.Distribution,
		predeploys.Redenom,
	}

	for _, addr := range predeployList {
		addrs.predeployAddrs[common.HexToAddress(addr)] = true
	}

	return addrs, nil
}

// categorizeAddress determines what action to take for an address.
func categorizeAddress(addr common.Address, internal *internalAddresses) action {
	// Burn specific predeploys
	if addr == common.HexToAddress(predeploys.Staking) {
		return actionBurnStaking
	}
	if addr == common.HexToAddress(predeploys.OmniBridgeNative) {
		return actionBurnNativeBridge
	}

	// Burn dead address
	if addr == common.HexToAddress("0x000000000000000000000000000000000000dead") {
		return actionBurnDead
	}

	// Consolidate other predeploys
	if internal.predeployAddrs[addr] {
		return actionConsolidatePredeploy
	}

	// Consolidate EOAs
	if internal.eoaAddrs[addr] {
		return actionConsolidateEOA
	}

	// Consolidate validators
	if internal.validatorAddrs[addr] {
		return actionConsolidateValidator
	}

	// Consolidate contracts
	if internal.contractAddrs[addr] {
		return actionConsolidateContract
	}

	// Keep regular user balances
	return actionKeep
}

// queryL1BridgeBalance queries the NOM token balance held by the L1 bridge contract.
func queryL1BridgeBalance(ctx context.Context, l1RPCURL string, bridgeAddr common.Address) (*big.Int, error) {
	client, err := ethclient.DialContext(ctx, "l1", l1RPCURL)
	if err != nil {
		return nil, errors.Wrap(err, "dial L1 RPC")
	}
	defer client.Close()

	// Query NOM token balance held by the bridge contract on Ethereum L1
	balance, err := tokenutil.BalanceOfAsset(ctx, client, tokens.NOM, bridgeAddr)
	if err != nil {
		return nil, errors.Wrap(err, "query NOM token balance")
	}

	return balance, nil
}

// addToResult adds an amount to the result map, summing if the address already exists.
func addToResult(result map[common.Address]*big.Int, addr common.Address, amount *big.Int) {
	if existing, ok := result[addr]; ok {
		existing.Add(existing, amount)
	} else {
		result[addr] = new(big.Int).Set(amount)
	}
}

// printConsolidationSummary prints a summary of the consolidation process.
func printConsolidationSummary(ctx context.Context, s *ConsolidationSummary) {
	log.Info(ctx, "═══════════════════════════════════════════════════════════")
	log.Info(ctx, "BALANCE CONSOLIDATION SUMMARY")
	log.Info(ctx, "═══════════════════════════════════════════════════════════")
	log.Info(ctx, "")
	log.Info(ctx, "INPUT:")
	log.Info(ctx, fmt.Sprintf("  EVM Total Supply:        %s NOM", FormatBalance(s.TotalEVMSupply)))
	log.Info(ctx, fmt.Sprintf("  Consensus Staking Total: %s NOM", FormatBalance(s.TotalConsensusStake)))
	log.Info(ctx, "")
	log.Info(ctx, "BURNED (not paid on L1):")
	log.Info(ctx, fmt.Sprintf("  Staking Predeploy:       %s NOM", FormatBalance(s.BurnedStaking)))
	log.Info(ctx, fmt.Sprintf("  Native Bridge Predeploy: %s NOM", FormatBalance(s.BurnedNativeBridge)))
	log.Info(ctx, fmt.Sprintf("  Dead Address:            %s NOM", FormatBalance(s.BurnedDead)))
	log.Info(ctx, fmt.Sprintf("  Total Burned:            %s NOM", FormatBalance(s.TotalBurned)))
	log.Info(ctx, "")
	log.Info(ctx, fmt.Sprintf("CONSOLIDATED (swept to %s):", s.FoundationAddress.Hex()))
	log.Info(ctx, fmt.Sprintf("  EOA Accounts:            %s NOM", FormatBalance(s.ConsolidatedEOAs)))
	log.Info(ctx, fmt.Sprintf("  Validator Addresses:     %s NOM", FormatBalance(s.ConsolidatedValidators)))
	log.Info(ctx, fmt.Sprintf("  Contract Addresses:      %s NOM", FormatBalance(s.ConsolidatedContracts)))
	log.Info(ctx, fmt.Sprintf("  Other Predeploys:        %s NOM", FormatBalance(s.ConsolidatedPredeploys)))
	log.Info(ctx, fmt.Sprintf("  Total Consolidated:      %s NOM", FormatBalance(s.TotalConsolidated)))
	log.Info(ctx, "")
	log.Info(ctx, "L1 BRIDGE:")
	log.Info(ctx, fmt.Sprintf("  Address:                 %s", s.L1BridgeAddress.Hex()))
	log.Info(ctx, fmt.Sprintf("  Available Balance:       %s NOM", FormatBalance(s.L1BridgeBalance)))
	log.Info(ctx, "")
	log.Info(ctx, "═══════════════════════════════════════════════════════════")
	log.Info(ctx, "RESULT:")
	log.Info(ctx, "")
	log.Info(ctx, fmt.Sprintf("  💰 Total to Pay on L1:   %s NOM", FormatBalance(s.TotalPayable)))
	if s.FoundationShortfall.Cmp(bi.Zero()) > 0 {
		log.Info(ctx, "")
		log.Info(ctx, fmt.Sprintf("  ⚠️  Foundation Shortfall: %s NOM", FormatBalance(s.FoundationShortfall)))
		log.Info(ctx, "      (deducted from foundation payout)")
	}
	log.Info(ctx, "")
	log.Info(ctx, "═══════════════════════════════════════════════════════════")
}

// writeConsolidatedBalances writes the consolidated balances to a JSON file.
func writeConsolidatedBalances(outputPath string, balances map[common.Address]*big.Int) error {
	// Convert to output format (just address and wei amount)
	type output struct {
		Address string `json:"address"`
		Wei     string `json:"wei"`
	}

	outputs := make([]output, 0, len(balances))
	for addr, amount := range balances {
		outputs = append(outputs, output{
			Address: addr.Hex(),
			Wei:     amount.String(),
		})
	}

	data, err := json.MarshalIndent(outputs, "", "  ")
	if err != nil {
		return errors.Wrap(err, "marshal json")
	}

	if err := os.WriteFile(outputPath, data, 0o644); err != nil {
		return errors.Wrap(err, "write file")
	}

	return nil
}

type evmBalancesFile struct {
	Balances []struct {
		Address string   `json:"address"`
		Balance *big.Int `json:"balance"`
	} `json:"balances"`
}

type stakingBalancesFile struct {
	TotalStake *big.Int `json:"-"`
	Delegators []struct {
		Address string   `json:"address"`
		Total   *big.Int `json:"total"`
	} `json:"delegators"`
}

// loadEVMBalances loads EVM balances from a JSON file.
func loadEVMBalances(path string) (*evmBalancesFile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "read file")
	}

	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, errors.Wrap(err, "unmarshal json")
	}

	var result evmBalancesFile
	if err := json.Unmarshal(raw["balances"], &result.Balances); err != nil {
		return nil, errors.Wrap(err, "unmarshal balances")
	}

	return &result, nil
}

// loadStakingBalances loads staking balances from a JSON file.
func loadStakingBalances(path string) (*stakingBalancesFile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "read file")
	}

	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, errors.Wrap(err, "unmarshal json")
	}

	var result stakingBalancesFile
	if err := json.Unmarshal(raw["delegators"], &result.Delegators); err != nil {
		return nil, errors.Wrap(err, "unmarshal delegators")
	}

	var totalStake string
	if err := json.Unmarshal(raw["total_stake"], &totalStake); err != nil {
		return nil, errors.Wrap(err, "unmarshal total_stake")
	}

	result.TotalStake = new(big.Int)
	if _, ok := result.TotalStake.SetString(totalStake, 10); !ok {
		return nil, errors.New("invalid total_stake format")
	}

	return &result, nil
}
