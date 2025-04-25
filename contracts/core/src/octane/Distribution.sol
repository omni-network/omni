// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

/**
 * @title Distribution
 * @notice The EVM interface to the consensus chain's x/distribution module.
 *         Calls are proxied, and not executed synchronously. Their execution is left to
 *         the consensus chain, and they may fail.
 */
contract Distribution is OwnableUpgradeable {
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
     * @notice Initialize the contract, used for fresh deployment
     */
    function initialize(address owner_) public initializer {
        __Ownable_init(owner_);
    }

    /**
     * @notice Withdraw delegation rewards from a validator
     * @dev Proxies x/distribution.MsgWithdrawDelegatorReward
     * @param validator The address of the validator with a stake
     */
    function withdraw(address validator) external payable {
        _burnFee();
        emit Withdraw(msg.sender, validator);
    }

    /**
     * @notice Burn the fee, requiring it be sent with the call
     */
    function _burnFee() internal {
        require(msg.value >= Fee, "Distribution: insufficient fee");
        payable(BurnAddr).transfer(msg.value);
    }
}
