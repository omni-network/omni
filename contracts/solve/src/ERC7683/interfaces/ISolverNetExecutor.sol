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
     * @notice Approves a spender (usually call target) to spend a token.
     * @dev Called prior to `executeCall` in order to ensure tokens can be spent.
     */
    function tokenApproval(address token, address spender, uint256 amount) external;

    /**
     * @notice Executes a call.
     */
    function executeCall(ISolverNet.Call memory call) external payable;

    /**
     * @notice Refunds excess tokens.
     * @dev Called after `executeCall` in order to refund any excess or returned tokens.
     */
    function refundExcess(address token, address spender, address to, uint256 amount) external;

    /**
     * @notice Refunds native currency.
     * @dev Called after `executeCall` in order to refund any native currency sent back to the executor.
     */
    function refundNative(address to) external;
}
