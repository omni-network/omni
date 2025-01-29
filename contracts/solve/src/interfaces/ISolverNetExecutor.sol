// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { ISolverNet } from "./ISolverNet.sol";

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
     * @notice Executes a call.
     */
    function execute(ISolverNet.Call memory call) external payable;

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
