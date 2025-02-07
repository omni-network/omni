// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

interface IStaking {
    /**
     * @notice Delegate tokens to a validator for another address
     * @param delegator The address of the delegator
     * @param validator The address of the validator to delegate to
     */
    function delegateFor(address delegator, address validator) external payable;
}
