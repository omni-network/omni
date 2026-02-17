// SPDX-License-Identifier: GPL-3.0-only
// solhint-disable no-console
pragma solidity 0.8.24;

import { Script } from "forge-std/Script.sol";
import { console } from "forge-std/console.sol";
import { MerkleGen } from "multiproof/src/MerkleGen.sol";
import { NominaBridgeL1 } from "src/token/nomina/NominaBridgeL1.sol";
import { stdJson } from "forge-std/StdJson.sol";
import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * @title PostHaltNominaL1BridgeWithdrawals
 * @notice Script to execute post-halt withdrawals from NominaBridgeL1 using merkle proofs.
 */
contract PostHaltNominaL1BridgeWithdrawals is Script {
    using stdJson for string;

    /// @notice Withdrawal data structure
    struct Withdrawal {
        address account;
        uint256 balance;
    }

    /// @notice Path to the withdrawals JSON file
    string public constant WITHDRAWALS_FILE = "script/admin/post-halt-withdrawals.json";

    /// @notice Maximum number of withdrawals to process in a single batch
    uint256 public constant BATCH_SIZE = 100;

    /// @notice Total number of withdrawals in the JSON file
    uint256 public constant TOTAL_WITHDRAWALS = 7526;

    /// @notice Cached withdrawals array to avoid multiple file reads
    Withdrawal[] private _withdrawals;

    /**
     * @notice Returns the list of withdrawals from JSON file.
     * @dev Reads from post-halt-withdrawals.json in the script/admin directory.
     *      Reads each withdrawal individually to handle large datasets efficiently.
     * @return withdrawals Array of (address, balance) pairs.
     */
    function getWithdrawals() public returns (Withdrawal[] memory withdrawals) {
        if (_withdrawals.length > 0) {
            return _withdrawals;
        }

        string memory root = vm.projectRoot();
        string memory path = string.concat(root, "/", WITHDRAWALS_FILE);
        string memory json = vm.readFile(path);

        // Read each withdrawal individually to avoid memory issues with large arrays
        for (uint256 i = 0; i < TOTAL_WITHDRAWALS; i++) {
            string memory basePath = string.concat(".[", vm.toString(i), "]");

            // Parse individual fields using type-specific parsers to avoid ABI decoding issues
            address account = vm.parseJsonAddress(json, string.concat(basePath, ".account"));
            uint256 balance = vm.parseJsonUint(json, string.concat(basePath, ".balance"));

            _withdrawals.push(Withdrawal({ account: account, balance: balance }));
        }

        return _withdrawals;
    }

    /**
     * @notice Returns the merkle root for all withdrawals.
     * @dev This should match the root set in initializeV3.
     * @return root The merkle root of all withdrawals.
     */
    function getWithdrawalRoot() public returns (bytes32 root) {
        Withdrawal[] memory withdrawals = getWithdrawals();

        // Create leaves for all withdrawals
        bytes32[] memory leaves = new bytes32[](withdrawals.length);
        for (uint256 i = 0; i < withdrawals.length; i++) {
            leaves[i] = keccak256(bytes.concat(keccak256(abi.encode(withdrawals[i].account, withdrawals[i].balance))));
        }

        // Generate proof for all to get root
        uint256[] memory allIndices = new uint256[](withdrawals.length);
        for (uint256 i = 0; i < withdrawals.length; i++) {
            allIndices[i] = i;
        }

        (,, root) = MerkleGen.generateMultiproof(leaves, allIndices);
    }

    /**
     * @notice Internal function to execute a single batch of withdrawals.
     * @param bridge The address of the NominaBridgeL1 contract.
     * @param allWithdrawals All withdrawals from the JSON file.
     * @param leaves All merkle leaves.
     * @param startIndex The index to start from (inclusive).
     * @param count The number of withdrawals to process in this batch.
     * @param broadcast Whether to broadcast the transaction or just simulate.
     */
    function _executeBatch(
        address bridge,
        Withdrawal[] memory allWithdrawals,
        bytes32[] memory leaves,
        uint256 startIndex,
        uint256 count,
        bool broadcast
    ) internal {
        console.log("\n=== Processing Batch ===");
        console.log("Start index:", startIndex);
        console.log("Count:", count);

        // Create indices for selected withdrawals
        uint256[] memory selectedIndices = new uint256[](count);
        for (uint256 i = 0; i < count; i++) {
            selectedIndices[i] = startIndex + i;
        }

        // Generate multiproof for selected withdrawals
        (bytes32[] memory multiProof, bool[] memory multiProofFlags,) =
            MerkleGen.generateMultiproof(leaves, selectedIndices);

        // Prepare accounts and amounts arrays
        address[] memory accounts = new address[](count);
        uint256[] memory amounts = new uint256[](count);
        for (uint256 i = 0; i < count; i++) {
            accounts[i] = allWithdrawals[startIndex + i].account;
            amounts[i] = allWithdrawals[startIndex + i].balance;
        }

        // Get token contract for balance checking
        IERC20 nomina = NominaBridgeL1(bridge).NOMINA();

        // Record balances before withdrawal
        uint256[] memory balancesBefore = new uint256[](count);
        for (uint256 i = 0; i < count; i++) {
            balancesBefore[i] = nomina.balanceOf(accounts[i]);
        }

        // Execute withdrawal (with or without broadcast)
        if (broadcast) {
            vm.startBroadcast();
            NominaBridgeL1(bridge).postHaltWithdraw(accounts, amounts, multiProof, multiProofFlags);
            vm.stopBroadcast();
        } else {
            NominaBridgeL1(bridge).postHaltWithdraw(accounts, amounts, multiProof, multiProofFlags);
        }

        // Verify balances increased by expected amounts
        for (uint256 i = 0; i < count; i++) {
            uint256 balanceAfter = nomina.balanceOf(accounts[i]);
            uint256 expectedIncrease = amounts[i];
            uint256 actualIncrease = balanceAfter - balancesBefore[i];

            require(
                actualIncrease == expectedIncrease,
                string(
                    abi.encodePacked(
                        "Balance mismatch for account ",
                        vm.toString(accounts[i]),
                        ": expected ",
                        vm.toString(expectedIncrease),
                        ", got ",
                        vm.toString(actualIncrease)
                    )
                )
            );
        }

        console.log("Batch complete - verified", count, "withdrawals");
    }

    /**
     * @notice Internal function containing shared logic for run and runNoBroadcast.
     * @dev Processes all withdrawals in batches of up to BATCH_SIZE (100).
     * @param bridge The address of the NominaBridgeL1 contract.
     * @param broadcast Whether to broadcast the transaction or just simulate.
     */
    function _executeWithdrawals(address bridge, bool broadcast) internal {
        require(bridge != address(0), "Invalid bridge address");

        Withdrawal[] memory allWithdrawals = getWithdrawals();

        console.log("=== Post-Halt Withdrawal Execution ===");
        console.log("Bridge address:", bridge);
        console.log("Total withdrawals:", allWithdrawals.length);
        console.log("Batch size:", BATCH_SIZE);
        console.log("Broadcast:", broadcast);

        // Verify the merkle root is set correctly
        bytes32 expectedRoot = getWithdrawalRoot();
        bytes32 actualRoot = NominaBridgeL1(bridge).postHaltRoot();
        require(actualRoot == expectedRoot, "Post halt root mismatch");
        console.log("Root verified:");
        console.logBytes32(expectedRoot);

        // Create leaves for all withdrawals (used for all batches)
        bytes32[] memory leaves = new bytes32[](allWithdrawals.length);
        for (uint256 i = 0; i < allWithdrawals.length; i++) {
            leaves[i] =
                keccak256(bytes.concat(keccak256(abi.encode(allWithdrawals[i].account, allWithdrawals[i].balance))));
        }

        // Process withdrawals in batches
        uint256 totalProcessed = 0;
        while (totalProcessed < allWithdrawals.length) {
            uint256 remaining = allWithdrawals.length - totalProcessed;
            uint256 batchSize = remaining < BATCH_SIZE ? remaining : BATCH_SIZE;

            _executeBatch(bridge, allWithdrawals, leaves, totalProcessed, batchSize, broadcast);

            totalProcessed += batchSize;
        }

        console.log("\n=== All Withdrawals Complete ===");
        console.log("Total processed:", totalProcessed);
    }

    /**
     * @notice Execute post-halt withdrawals for all accounts.
     * @dev Processes all withdrawals from the JSON file in batches of up to 100.
     *      This function broadcasts transactions to the network.
     * @param bridge The address of the NominaBridgeL1 contract.
     */
    function run(address bridge) external {
        _executeWithdrawals(bridge, true);
    }

    /**
     * @notice Execute post-halt withdrawals without broadcasting.
     * @dev This function performs the same operations as run() but WITHOUT broadcasting.
     *      It's useful for:
     *      - Testing in post-upgrade scenarios to verify withdrawals work correctly
     *      - Validating merkle proofs before executing real transactions
     *      - Checking that balances are updated correctly
     *      The simulation will revert if the proof is invalid or if any check fails.
     * @param bridge The address of the NominaBridgeL1 contract.
     */
    function runNoBroadcast(address bridge) external {
        _executeWithdrawals(bridge, false);
    }

    /**
     * @notice Execute a specific range of post-halt withdrawals without broadcasting.
     * @dev Useful for testing a subset of withdrawals (e.g. first N and last N).
     * @param bridge The address of the NominaBridgeL1 contract.
     * @param start The starting index (inclusive).
     * @param count The number of withdrawals to process.
     */
    function runNoBroadcastRange(address bridge, uint256 start, uint256 count) external {
        require(bridge != address(0), "Invalid bridge address");

        Withdrawal[] memory allWithdrawals = getWithdrawals();
        require(start + count <= allWithdrawals.length, "Range out of bounds");

        // Verify the merkle root is set correctly
        bytes32 expectedRoot = getWithdrawalRoot();
        bytes32 actualRoot = NominaBridgeL1(bridge).postHaltRoot();
        require(actualRoot == expectedRoot, "Post halt root mismatch");

        // Create leaves for all withdrawals (needed for multiproof generation)
        bytes32[] memory leaves = new bytes32[](allWithdrawals.length);
        for (uint256 i = 0; i < allWithdrawals.length; i++) {
            leaves[i] =
                keccak256(bytes.concat(keccak256(abi.encode(allWithdrawals[i].account, allWithdrawals[i].balance))));
        }

        _executeBatch(bridge, allWithdrawals, leaves, start, count, false);
    }

    /**
     * @notice Print all withdrawals for verification.
     */
    function printWithdrawals() external {
        Withdrawal[] memory withdrawals = getWithdrawals();

        console.log("=== All Withdrawals ===");
        console.log("Total count:", withdrawals.length);
        console.log("");

        uint256 totalBalance = 0;
        for (uint256 i = 0; i < withdrawals.length; i++) {
            console.log("Index:", i);
            console.log("Account:", withdrawals[i].account);
            console.log("Balance:", withdrawals[i].balance);
            console.log("---");
            totalBalance += withdrawals[i].balance;
        }

        console.log("\nTotal balance:", totalBalance);
        console.log("\nMerkle Root:");
        console.logBytes32(getWithdrawalRoot());
    }
}
