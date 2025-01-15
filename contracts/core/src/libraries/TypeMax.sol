// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

/**
 * @title TypeMax
 * @dev Util lib with type maximum values
 */
library TypeMax {
    uint40 internal constant Uint40 = type(uint40).max;
    uint256 internal constant Uint256 = type(uint256).max;
    bytes32 internal constant Bytes32 = bytes32(Uint256);
    address internal constant Address = address(uint160(Uint256));
}
