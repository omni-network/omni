// SPDX-License-Identifier: Unlicense
pragma solidity 0.8.24;

/**
 * @title XGasLimits
 * @notice Constant gas limits used in xcalls.
 *         Values determined via unit tests, with buffer for safety.
 */
library GasLimits {
    /// @notice XStakeController.recordStake xcall gas limit.
    uint64 internal constant RecordStake = 100_000;

    /// @notice XStakeController.unstakeFor xcall gas limit.
    uint64 internal constant UnstakeFor = 100_000;

    /// @notice XStaker.withdraw xcall gas limit.
    uint64 internal constant Withdraw = 100_000;
}
