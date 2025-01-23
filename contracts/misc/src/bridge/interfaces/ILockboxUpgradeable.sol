// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

interface ILockboxUpgradeable {
    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                          STORAGE                           */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev The address of the ERC20 token stored in the lockbox.
     */
    function token() external view returns (address);

    /**
     * @dev The address of the bridgeable wrapper contract for `token`.
     */
    function wrapper() external view returns (address);

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                     LOCKBOX FUNCTIONS                      */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Deposit tokens into the lockbox.
     * @param value The amount of tokens to deposit.
     */
    function deposit(uint256 value) external;

    /**
     * @dev Deposit tokens into the lockbox for a specific address.
     * @param to    The address to deposit the tokens for.
     * @param value The amount of tokens to deposit.
     */
    function depositTo(address to, uint256 value) external;

    /**
     * @dev Withdraw tokens from the lockbox.
     * @param value The amount of tokens to withdraw.
     */
    function withdraw(uint256 value) external;

    /**
     * @dev Withdraws tokens from the lockbox to a specific address.
     * @param to    The address to withdraw the tokens to.
     * @param value The amount of tokens to withdraw.
     */
    function withdrawTo(address to, uint256 value) external;
}
