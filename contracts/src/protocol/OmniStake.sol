// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

/**
 * @title OmniStake
 * @notice The deposit contract for OMNI-staked validators.
 */
contract OmniStake {
    /**
     * @notice Emitted when a user deposits funds into the contract.
     * @param depositor Address of the depositor
     * @param amount    Funds deposited
     */
    event Deposit(address indexed depositor, uint256 amount);

    /**
     * @notice Deposit OMNI. This is the entry point for validator staking. The consensus chain is
     *         notified of the deposit, and manages stake accounting / validator onboarding.
     */
    function deposit() external payable {
        require(msg.value > 1 ether, "OmniStake: deposit amt too low");
        require(msg.value < type(uint64).max, "OmniStake: deposit amt too high");
        emit Deposit(msg.sender, msg.value);
    }
}
