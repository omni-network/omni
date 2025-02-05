// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

/// @dev Interface used by bridge and lockbox to mint and clawback tokens.
interface ITokenOps {
    /**
     * @dev Mint tokens to the given address.
     * @param to    The address to mint tokens to.
     * @param value The amount of tokens to mint.
     */
    function mint(address to, uint256 value) external;

    /**
     * @dev Burn tokens from the caller.
     * @param value The amount of tokens to burn.
     */
    function burn(uint256 value) external;

    /**
     * @dev Clawback tokens from the given address.
     * @param from  The address to clawback tokens from.
     * @param value The amount of tokens to clawback.
     */
    function clawback(address from, uint256 value) external;
}
