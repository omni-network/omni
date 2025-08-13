// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { PausableUpgradeable } from "src/utils/PausableUpgradeable.sol";

/**
 * @title NominaBridgeCommon
 * @notice Common constants and functions for NominaBridge contracts
 */
abstract contract NominaBridgeCommon is OwnableUpgradeable, PausableUpgradeable {
    /// @notice Pausable key for withdraws
    bytes32 public constant ACTION_WITHDRAW = keccak256("withdraw");

    /// @notice Pausable key for bridges
    bytes32 public constant ACTION_BRIDGE = keccak256("bridge");

    /// @notice Revert if `action` is paused
    modifier whenNotPaused(bytes32 action) {
        require(!_isPaused(action), "NominaBridge: paused");
        _;
    }

    /// @notice Pause `action`
    function pause(bytes32 action) external onlyOwner {
        _pause(action);
    }

    /// @notice Unpause `action`
    function unpause(bytes32 action) external onlyOwner {
        _unpause(action);
    }

    /// @notice Pause all actions
    function pause() external onlyOwner {
        _pauseAll();
    }

    /// @notice Unpause all actions
    function unpause() external onlyOwner {
        _unpauseAll();
    }

    /// @notice Returns true if `action` is paused
    function isPaused(bytes32 action) external view returns (bool) {
        return _isPaused(action);
    }
}
