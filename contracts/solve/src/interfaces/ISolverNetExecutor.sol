// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.24;

interface ISolverNetExecutor {
    /**
     * @notice Error thrown when the sender is not the outbox.
     */
    error NotOutbox();

    /**
     * @notice Error thrown when the call fails.
     */
    error CallFailed();

    /**
     * @notice Address of the outbox.
     */
    function outbox() external view returns (address);

    /**
     * @notice Approves a spender (usually call target) to spend a token held by the executor.
     * @dev Called prior to `execute` in order to ensure tokens can be spent and after to purge excess approvals.
     */
    function approve(address token, address spender, uint256 amount) external;

    /**
     * @notice Attempts to revoke an approval for a spender.
     * @dev If the token reverts when setting approval to zero, this will not revert.
     */
    function tryRevokeApproval(address token, address spender) external;

    /**
     * @notice Executes a call.
     * @param target Address of the contract to call.
     * @param value  Value to send with the call.
     * @param data   Data to send with the call.
     */
    function execute(address target, uint256 value, bytes calldata data) external payable;

    /**
     * @notice Transfers a token to a recipient.
     * @dev Called after `execute` in order to refund any excess or returned tokens.
     */
    function transfer(address token, address to, uint256 amount) external;

    /**
     * @notice Transfers native currency to a recipient.
     * @dev Called after `execute` in order to refund any native currency sent back to the executor.
     */
    function transferNative(address to, uint256 amount) external;
}
