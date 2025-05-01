// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

interface IGenesisStakeV2 {
    /**
     * @notice Migrate a user's stake to the rewards distributor.
     * @param addr The address of the user to migrate.
     * @return The amount of tokens migrated.
     */
    function migrateStake(address addr) external returns (uint256);
}
