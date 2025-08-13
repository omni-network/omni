// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.24;

interface INominaPortalPausable {
    /**
     * @notice Return true if actionId for is paused for the given chain
     * @dev Also returns if all actions are paused or if the action is globally paused
     */
    function isPaused(bytes32 actionId, uint64 chainId) external view returns (bool);

    /**
     * @notice Return true if actionId is paused for all chains
     * @dev Also returns if all actions are paused
     */
    function isPaused(bytes32 actionId) external view returns (bool);

    /**
     * @notice Return true if all actions are paused
     */
    function isPaused() external view returns (bool);
}
