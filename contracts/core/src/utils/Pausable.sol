// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

/**
 * @title Pausable
 * @notice Contract module which provides a way to pause certain functions by key.
 * @dev We use a map of bytes32 key to bools, rather than uint256 bitmap, to allow keys to be generated dynamically.
 *      This allows for flexible pausing, but at higher gas cost.
 */
contract PausableUpgradeable {
    /// @custom:storage-location erc7201:omni.storage.Pauseable
    struct PauseableStorage {
        mapping(bytes32 => bool) _paused;
    }

    // keccak256(abi.encode(uint256(keccak256("omni.storage.Pauseable")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant PausableStorageSlot = 0xff37105740f03695c8f3597f3aff2b92fbe1c80abea3c28731ecff2efd693400;

    function _getPauseableStorage() internal pure returns (PauseableStorage storage $) {
        assembly {
            $.slot := PausableStorageSlot
        }
    }

    /**
     * @dev Special key for pausing all keys.
     */
    bytes32 public constant KeyPauseAll = keccak256("PAUSE_ALL");

    /**
     * @notice Pause by key.
     */
    function _pause(bytes32 key) internal {
        PauseableStorage storage $ = _getPauseableStorage();
        require(!$._paused[key], "Pausable: paused");
        $._paused[key] = true;
    }

    /**
     * @notice Unpause by key.
     */
    function _unpause(bytes32 key) internal {
        PauseableStorage storage $ = _getPauseableStorage();
        require($._paused[key], "Pausable: not paused");
        $._paused[key] = false;
    }

    /**
     * @notice Returns true if `key` is paused, or all keys are paused.
     */
    function _isPaused(bytes32 key) internal view returns (bool) {
        PauseableStorage storage $ = _getPauseableStorage();
        return $._paused[KeyPauseAll] || $._paused[key];
    }

    /**
     * @notice Returns true if either `key1` or `key2` is paused, or all keys are paused.
     */
    function _isPaused(bytes32 key1, bytes32 key2) internal view returns (bool) {
        PauseableStorage storage $ = _getPauseableStorage();
        return $._paused[KeyPauseAll] || $._paused[key1] || $._paused[key2];
    }

    /**
     * @notice Returns true if all keys are paused.
     */
    function _isAllPaused() internal view returns (bool) {
        PauseableStorage storage $ = _getPauseableStorage();
        return $._paused[KeyPauseAll];
    }

    /**
     * @notice Pause all keys.
     */
    function _pauseAll() internal {
        _pause(KeyPauseAll);
    }

    /**
     * @notice Unpause all keys.
     */
    function _unpauseAll() internal {
        _unpause(KeyPauseAll);
    }
}
