// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

interface ILockboxUpgradeable {
    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                           ERRORS                           */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Error thrown when a withdraw with an insufficient balance is attempted.
     */
    error InsufficientBalance(address token, address depositor, uint256 balance, uint256 value);

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                           EVENTS                           */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Event emitted when a deposit is made.
     */
    event Deposit(address indexed token, address indexed caller, address indexed beneficiary, uint256 value);

    /**
     * @dev Event emitted when a withdraw is made.
     */
    event Withdraw(address indexed token, address indexed caller, address indexed recipient, uint256 value);

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                       VIEW FUNCTIONS                       */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Mapping of depositor to token to balance.
     * @param depositor The address of the depositor.
     * @param token The address of the token.
     * @return balance The balance of the depositor for the token.
     */
    function balances(address depositor, address token) external view returns (uint256 balance);

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                     LOCKBOX FUNCTIONS                      */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Deposit tokens into the lockbox.
     * @param token The address of the token to deposit.
     * @param value The amount of tokens to deposit.
     */
    function deposit(address token, uint256 value) external;

    /**
     * @dev Deposit tokens into the lockbox for a specific address.
     * @param token The address of the token to deposit.
     * @param to The address to deposit the tokens for.
     * @param value The amount of tokens to deposit.
     */
    function depositTo(address token, address to, uint256 value) external;

    /**
     * @dev Withdraw tokens from the lockbox.
     * @param token The address of the token to withdraw.
     * @param value The amount of tokens to withdraw.
     */
    function withdraw(address token, uint256 value) external;

    /**
     * @dev Withdraw tokens from the lockbox to a specific address.
     * @param token The address of the token to withdraw.
     * @param to The address to withdraw the tokens to.
     * @param value The amount of tokens to withdraw.
     */
    function withdrawTo(address token, address to, uint256 value) external;
}
