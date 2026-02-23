// SPDX-License-Identifier: GPL-3.0-only
// solhint-disable no-console
pragma solidity 0.8.24;

import { Script } from "forge-std/Script.sol";
import { console } from "forge-std/console.sol";
import { MerkleGen } from "multiproof/src/MerkleGen.sol";
import { NominaBridgeL1 } from "src/token/nomina/NominaBridgeL1.sol";
import { stdJson } from "forge-std/StdJson.sol";

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

    /// @notice Maximum number of withdrawals to process in a single tx batch
    uint256 public constant BATCH_SIZE = 200;

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

        // Execute withdrawal (with or without broadcast)
        if (broadcast) {
            vm.startBroadcast();
            NominaBridgeL1(bridge).postHaltWithdraw(accounts, amounts, multiProof, multiProofFlags);
            vm.stopBroadcast();
        } else {
            NominaBridgeL1(bridge).postHaltWithdraw(accounts, amounts, multiProof, multiProofFlags);
        }

        console.log("Batch complete - verified", count, "withdrawals");
    }

    /**
     * @notice Internal function containing shared logic for run and runNoBroadcast.
     * @dev Processes all withdrawals in tx batches of up to BATCH_SIZE.
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

        // Verify the merkle root is set correctly
        require(NominaBridgeL1(bridge).postHaltRoot() == getWithdrawalRoot(), "Post halt root mismatch");

        // Create leaves for all withdrawals (needed for multiproof generation)
        bytes32[] memory leaves = new bytes32[](allWithdrawals.length);
        for (uint256 i = 0; i < allWithdrawals.length; i++) {
            leaves[i] =
                keccak256(bytes.concat(keccak256(abi.encode(allWithdrawals[i].account, allWithdrawals[i].balance))));
        }

        uint256 cursor = 0;
        while (cursor < allWithdrawals.length) {
            uint256 remaining = allWithdrawals.length - cursor;
            uint256 batchSize = remaining < BATCH_SIZE ? remaining : BATCH_SIZE;

            _executeBatch(bridge, allWithdrawals, leaves, cursor, batchSize, broadcast);

            cursor += batchSize;
        }

        console.log("\n=== All Withdrawals Complete ===");
        console.log("Total processed:", allWithdrawals.length);
    }

    /**
     * @notice Execute post-halt withdrawals for all accounts.
     * @dev Processes all withdrawals from the JSON file in tx batches of up to BATCH_SIZE.
     *      This function broadcasts transactions to the network.
     * @param bridge The address of the NominaBridgeL1 contract.
     */
    function run(address bridge) external {
        _executeWithdrawals(bridge, true);
    }

    /**
     * @notice Pre-compute all batch calldata and write to files.
     * @dev Writes abi-encoded postHaltWithdraw calldata for each batch to script/admin/batches/.
     *      Run this once offline, then use submitBatches to broadcast quickly.
     */
    function prepareBatches() external {
        Withdrawal[] memory allWithdrawals = getWithdrawals();

        bytes32[] memory leaves = _makeLeaves(allWithdrawals);

        string memory outDir = string.concat(vm.projectRoot(), "/script/admin/batches");
        vm.createDir(outDir, true);

        uint256 cursor = 0;
        uint256 batchNum = 0;
        while (cursor < allWithdrawals.length) {
            uint256 remaining = allWithdrawals.length - cursor;
            uint256 batchSize = remaining < BATCH_SIZE ? remaining : BATCH_SIZE;

            bytes memory callData = _encodeBatch(allWithdrawals, leaves, cursor, batchSize);

            string memory filePath = string.concat(outDir, "/batch_", vm.toString(batchNum), ".hex");
            vm.writeFile(filePath, vm.toString(callData));
            console.log("Batch %d written: %d withdrawals", batchNum, batchSize);

            cursor += batchSize;
            batchNum++;
        }

        console.log("Total batches prepared:", batchNum);
    }

    /**
     * @notice Submit pre-computed batches from files.
     * @dev Reads calldata from script/admin/batches/ and broadcasts each tx.
     * @param bridge The address of the NominaBridgeL1 contract.
     * @param numBatches The number of batch files to submit.
     * @param offset The number of batches already submitted (start from this index).
     */
    function submitBatches(address bridge, uint256 numBatches, uint256 offset) external {
        string memory batchDir = string.concat(vm.projectRoot(), "/script/admin/batches");

        vm.startBroadcast();
        for (uint256 i = offset; i < offset + numBatches; i++) {
            string memory filePath = string.concat(batchDir, "/batch_", vm.toString(i), ".hex");
            bytes memory callData = vm.parseBytes(vm.readFile(filePath));

            console.log("Submitting batch", i);
            (bool success,) = bridge.call(callData);
            require(success, string.concat("Batch ", vm.toString(i), " failed"));
        }
        vm.stopBroadcast();

        console.log("Total batches submitted:", numBatches);
    }

    /**
     * @notice Build merkle leaves from withdrawals.
     */
    function _makeLeaves(Withdrawal[] memory allWithdrawals) internal pure returns (bytes32[] memory leaves) {
        leaves = new bytes32[](allWithdrawals.length);
        for (uint256 i = 0; i < allWithdrawals.length; i++) {
            leaves[i] =
                keccak256(bytes.concat(keccak256(abi.encode(allWithdrawals[i].account, allWithdrawals[i].balance))));
        }
    }

    /**
     * @notice Encode a single batch as postHaltWithdraw calldata.
     */
    function _encodeBatch(
        Withdrawal[] memory allWithdrawals,
        bytes32[] memory leaves,
        uint256 startIndex,
        uint256 count
    ) internal pure returns (bytes memory) {
        uint256[] memory selectedIndices = new uint256[](count);
        address[] memory accounts = new address[](count);
        uint256[] memory amounts = new uint256[](count);
        for (uint256 i = 0; i < count; i++) {
            selectedIndices[i] = startIndex + i;
            accounts[i] = allWithdrawals[startIndex + i].account;
            amounts[i] = allWithdrawals[startIndex + i].balance;
        }

        (bytes32[] memory multiProof, bool[] memory multiProofFlags,) =
            MerkleGen.generateMultiproof(leaves, selectedIndices);

        return abi.encodeCall(NominaBridgeL1.postHaltWithdraw, (accounts, amounts, multiProof, multiProofFlags));
    }

    /**
     * @notice Execute post-halt withdrawals without broadcasting.
     * @dev Same as run() but without broadcasting. Useful for testing and validation.
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
