// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { ISolverNetExecutor } from "./interfaces/ISolverNetExecutor.sol";
import { AddrUtils } from "../lib/AddrUtils.sol";

contract SolverNetExecutor is ISolverNetExecutor {
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

    constructor(address _outbox) {
        outbox = _outbox;
    }

    /**
     * @notice Approves a spender (usually call target) to spend a token held by the executor.
     * @dev Called prior to `execute` in order to ensure tokens can be spent and after to purge excess approvals.
     */
    function approve(address token, address spender, uint256 amount) external onlyOutbox {
        token.safeApprove(spender, amount);
    }

    /**
     * @notice Executes a call.
     * @param target Address of the contract to call.
     * @param value  Value to send with the call.
     * @param data   Data to send with the call.
     */
    function execute(address target, uint256 value, bytes calldata data) external payable onlyOutbox {
        (bool success,) = payable(target).call{ value: value }(data);
        if (!success) revert CallFailed();
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
     * @dev Allows target contracts to arbitrarily return native tokens to the executor.
     */
    receive() external payable { }

    /**
     * @dev Allows target contracts to arbitrarily return native tokens to the executor.
     */
    fallback() external payable { }
}
