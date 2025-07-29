// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";

/**
 * @title Distribution
 * @notice The EVM interface to the consensus chain's x/distribution module.
 *         Calls are proxied, and not executed synchronously. Their execution is left to
 *         the consensus chain, and they may fail.
 */
contract Distribution {
    using SafeTransferLib for address;

    /**
     * @notice Error thrown when the contract is temporarily disabled
     */
    error TemporarilyDisabled();

    /**
     * @notice Emitted when a rewards withdrawal request is made by a delegator
     * @param delegator     (MsgWithdraw.delegator_addr) The address of the delegator
     * @param validator     (MsgWithdraw.validator_addr) The address of the validator with a delegation
     */
    event Withdraw(address indexed delegator, address indexed validator);

    /**
     * @notice The address to burn fees to
     */
    address private constant BurnAddr = 0x000000000000000000000000000000000000dEaD;

    /**
     * @notice Static fee to withdraw. Used to prevent spamming of events, which require consensus
     *         chain work that is not metered by execution chain gas.
     */
    uint256 public constant Fee = 0.1 ether;

    /**
     * @notice Withdraw delegation rewards from a validator
     * @dev Proxies x/distribution.MsgWithdrawDelegatorReward
     * @param validator The address of the validator with a stake
     */
    function withdraw(address validator) external payable {
        revert TemporarilyDisabled(); // Remove this and the error, and fix test and admin script to reenable
        _burnFee();
        emit Withdraw(msg.sender, validator);
    }

    /**
     * @notice Burn the fee, requiring it be sent with the call
     */
    function _burnFee() internal {
        require(msg.value >= Fee, "Distribution: insufficient fee");

        BurnAddr.safeTransferETH(Fee);

        // Send any overpayment back to the caller
        uint256 remaining = msg.value - Fee;
        if (remaining > 0) {
            msg.sender.safeTransferETH(remaining);
        }
    }
}
