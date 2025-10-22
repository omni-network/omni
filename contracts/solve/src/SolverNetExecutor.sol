// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { Receiver } from "solady/src/accounts/Receiver.sol";
import { ISolverNetExecutor } from "./interfaces/ISolverNetExecutor.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { AddrUtils } from "./lib/AddrUtils.sol";
import { IERC20 } from "@openzeppelin/contracts/interfaces/IERC20.sol";
import { IERC721 } from "@openzeppelin/contracts/interfaces/IERC721.sol";

contract SolverNetExecutor is Receiver, ISolverNetExecutor {
    using SafeTransferLib for address;
    using AddrUtils for bytes32;

    /**
     * @notice Address of the outbox.
     */
    address public immutable outbox;

    /**
     * @notice Modifier to provide access control to the outbox.
     * @dev This was used as it is more efficient than using Ownable. Only the outbox will call these functions.
     */
    modifier onlyOutbox() {
        if (msg.sender != outbox) revert NotOutbox();
        _;
    }

    /**
     * @notice Modifier to provide access control to the executor.
     * @dev This was used as it is more efficient than using Ownable. Only the executor will call these functions.
     */
    modifier onlySelf() {
        if (msg.sender != address(this)) revert NotSelf();
        _;
    }

    constructor(address _outbox) {
        outbox = _outbox;
    }

    /**
     * @notice Approves a spender (usually call target) to spend a token held by the executor.
     * @dev Called prior to `execute` in order to ensure tokens can be spent and after to purge excess approvals.
     */
    function approve(address token, address spender, uint256 amount) external onlyOutbox {
        token.safeApproveWithRetry(spender, amount);
    }

    /**
     * @notice Attempts to revoke an approval for a spender.
     * @dev If the token reverts when setting approval to zero, this will not revert.
     */
    function tryRevokeApproval(address token, address spender) external onlyOutbox {
        try IERC20(token).approve(spender, 0) { }

            // If the token reverts when setting approval to zero, try setting it to 1
            catch {
            try IERC20(token).approve(spender, 1) { } catch { }
        }
    }

    /**
     * @notice Executes a call.
     * @dev Any tokens sent back to this contract after a call will remain in the contract for the next call
     *      Receiving tokens need to be followed up by an `approve` call to use them in following calls (unless calling `transfer`)
     * @param target Address of the contract to call.
     * @param value  Value to send with the call.
     * @param data   Data to send with the call.
     */
    function execute(address target, uint256 value, bytes calldata data) external payable onlyOutbox {
        if (target == address(0)) target = address(this);
        // If fallback is later enabled, we need to validate that self-calls are for existing function selectors
        (bool success,) = payable(target).call{ value: value }(data);
        if (!success) revert CallFailed();
    }

    /**
     * @notice Execute a call and transfer any received ERC20 tokens back to the recipient
     * @dev Intended to be used when interacting with contracts that don't allow us to specify a recipient
     *      This should be triggered by `execute` by executing an external call against this Executor contract
     * @param token  Token to transfer
     * @param to     Recipient address
     * @param target Call target address
     * @param data   Calldata for the call
     */
    function executeAndTransfer(address token, address to, address target, bytes calldata data)
        external
        payable
        onlySelf
    {
        (bool success,) = target.call{ value: msg.value }(data);
        if (!success) revert CallFailed();

        if (token == address(0)) SafeTransferLib.safeTransferAllETH(to);
        else token.safeTransferAll(to);
    }

    /**
     * @notice Execute a call and transfer a received ERC721 token back to the recipient
     * @dev Intended to be used when interacting with contracts that don't allow us to specify a recipient
     *      This should be triggered by `execute` by executing an external call against this Executor contract
     * @param token     Token to transfer
     * @param tokenId   Token ID to transfer
     * @param to        Recipient address
     * @param target    Call target address
     * @param data      Calldata for the call
     */
    function executeAndTransfer721(address token, uint256 tokenId, address to, address target, bytes calldata data)
        external
        payable
        onlySelf
    {
        (bool success,) = target.call{ value: msg.value }(data);
        if (!success) revert CallFailed();

        if (token == address(0)) revert InvalidToken();
        IERC721(token).transferFrom(address(this), to, tokenId);
    }

    /**
     * @notice Transfers a token to a recipient.
     * @dev Called after `execute` in order to refund any excess or returned tokens.
     */
    function transfer(address token, address to, uint256 amount) external onlyOutbox {
        token.safeTransfer(to, amount);
    }

    /**
     * @notice Transfers native currency to a recipient.
     * @dev Called after `execute` in order to refund any native currency sent back to the executor.
     */
    function transferNative(address to, uint256 amount) external onlyOutbox {
        to.safeTransferETH(amount);
    }

    /**
     * @dev Do not allow self-transfers, as this indicates user error.
     */
    receive() external payable override {
        if (msg.sender == address(this)) revert CallFailed();
    }
}
