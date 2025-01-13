// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { ISolverNet, ISolverNetExecutor } from "./interfaces/ISolverNetExecutor.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { AddrUtils } from "./lib/AddrUtils.sol";

contract SolverNetExecutor is ISolverNetExecutor {
    using SafeTransferLib for address;

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
     * @notice Approves a spender (usually call target) to spend a token.
     * @dev Called prior to `executeCall` in order to ensure tokens can be spent.
     */
    function tokenApproval(address token, address spender, uint256 amount) external onlyOutbox {
        token.safeApprove(spender, amount);
    }

    /**
     * @notice Executes a call.
     */
    function executeCall(ISolverNet.Call memory call) external payable onlyOutbox {
        address target = AddrUtils.bytes32ToAddress(call.target);
        (bool success,) = payable(target).call{ value: call.value }(call.data);
        if (!success) revert CallFailed();
    }

    /**
     * @notice Refunds excess tokens.
     * @dev Called after `executeCall` in order to refund any excess or returned tokens.
     */
    function refundExcess(address token, address spender, address to, uint256 amount) external onlyOutbox {
        token.safeApprove(spender, 0);
        token.safeTransfer(to, amount);
    }

    /**
     * @notice Refunds native currency.
     * @dev Called after `executeCall` in order to refund any native currency sent back to the executor.
     */
    function refundNative(address to) external onlyOutbox {
        to.safeTransferETH(address(this).balance);
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
