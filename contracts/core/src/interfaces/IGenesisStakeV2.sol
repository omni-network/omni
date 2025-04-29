// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

interface IGenesisStakeV2 {
    /**
     * @notice Stake `amount` tokens.
     * @param amount    The amount of tokens to stake.
     */
    function stake(uint256 amount) external;

    /**
     * @notice Stake `amount` tokens for `recipient`, paid by the caller.
     * @param recipient The recipient to stake tokens for.
     * @param amount    The amount of tokens to stake.
     */
    function stakeFor(address recipient, uint256 amount) external;

    /**
     * @notice Migrate a user's stake to the rewards distributor.
     * @param addr The address of the user to migrate.
     * @return The amount of tokens migrated.
     */
    function migrateStake(address addr) external returns (uint256);
}
