// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

/**
 * @title ConfLevel
 * @notice XMsg confirmation levels. Matches ConfLevels in lib/xchain/types.go
 * @dev We prefer explicit constants over Enums, because we want uint8 values to start at 1, not 0, as they do in
 *      lib/xchain/types.go, such that 0 can represent "unset".
 */
library ConfLevel {
    /**
     * @notice XMsg confirmation level "latest", last byte of xmsg.shardId.
     */
    uint8 internal constant Latest = 1;

    /**
     * @notice XMsg confirmation level "fast", last byte of xmsg.shardId.
     */
    uint8 internal constant Fast = 2;

    /**
     * @notice XMsg confirmation level "safe", last byte of xmsg.shardId.
     */
    uint8 internal constant Safe = 3;

    /**
     * @notice XMsg confirmation level "finalized", last byte of xmsg.shardId.
     */
    uint8 internal constant Finalized = 4;

    /**
     * @notice Returns true if the given level is valid.
     */
    function isValid(uint8 level) internal pure returns (bool) {
        return level >= Latest && level <= Finalized;
    }
}
