// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.24;

interface ISolverNetExecutor {
    /**
     * @notice Error thrown when the sender is not the executor.
     */
    error NotSelf();

    /**
     * @notice Error thrown when the sender is not the outbox.
     */
    error NotOutbox();

    /**
     * @notice Error thrown when the call fails.
     */
    error CallFailed();

    /**
     * @notice Error thrown when the token is invalid.
     */
    error InvalidToken();

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
     * @dev Any tokens sent back to this contract after a call will remain in the contract for the next call
     *         Receiving tokens need to be followed up by an `approve` call to use them in following calls (unless calling `transfer`)
     * @param target Address of the contract to call.
     * @param value  Value to send with the call.
     * @param data   Data to send with the call.
     */
    function execute(address target, uint256 value, bytes calldata data) external payable;

    /**
     * @notice Execute a call and transfer any received ERC20 tokens back to the recipient
     * @dev Intended to be used when interacting with contracts that don't allow us to specify a recipient
     *         This should be triggered by `execute` by executing an external call against this Executor contract
     * @param token  Token to transfer
     * @param to     Recipient address
     * @param target Call target address
     * @param data   Calldata for the call
     */
    function executeAndTransfer(address token, address to, address target, bytes calldata data) external payable;

    /**
     * @notice Execute a call and transfer a received ERC721 token back to the recipient
     * @dev Intended to be used when interacting with contracts that don't allow us to specify a recipient
     *         This should be triggered by `execute` by executing an external call against this Executor contract
     * @param token     Token to transfer
     * @param tokenId   Token ID to transfer
     * @param to        Recipient address
     * @param target    Call target address
     * @param data      Calldata for the call
     */
    function executeAndTransfer721(address token, uint256 tokenId, address to, address target, bytes calldata data)
        external
        payable;

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
