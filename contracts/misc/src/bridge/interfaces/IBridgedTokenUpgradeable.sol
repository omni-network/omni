// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

interface IBridgedTokenUpgradeable {
    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                   MANAGEMENT FUNCTIONS                     */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Mints tokens to an address.
     * @param to The address to mint tokens to.
     * @param amount The amount of tokens to mint.
     */
    function mint(address to, uint256 amount) external;

    /**
     * @dev Burns tokens from the caller.
     * @param amount The amount of tokens to burn.
     */
    function burn(uint256 amount) external;

    /**
     * @dev Burns tokens from an address.
     * @param from The address to burn tokens from.
     * @param amount The amount of tokens to burn.
     */
    function clawback(address from, uint256 amount) external;

    /**
     * @dev Pauses the contract.
     */
    function pause() external;

    /**
     * @dev Unpauses the contract.
     */
    function unpause() external;

    /**
     * @dev Pauses accounts.
     * @param accounts The accounts to pause.
     */
    function pauseAccounts(address[] calldata accounts) external;

    /**
     * @dev Unpauses an account.
     * @param account The account to unpause.
     */
    function unpauseAccount(address account) external;
}
