// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import { Initializable } from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import { AccessControlUpgradeable } from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import { PausableUpgradeable } from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import { ILockbox } from "./interfaces/ILockbox.sol";

import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { ITokenOps } from "./interfaces/ITokenOps.sol";

contract Lockbox is Initializable, AccessControlUpgradeable, PausableUpgradeable, ILockbox {
    using SafeTransferLib for address;

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                         CONSTANTS                          */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    // keccak256("PAUSER")
    bytes32 public constant PAUSER_ROLE = 0x539440820030c4994db4e31b6b800deafd503688728f932addfe7a410515c14c;

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                          STORAGE                           */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Address of the ERC20 token stored in the lockbox.
     */
    address public token;

    /**
     * @dev Address of the bridgeable wrapper contract for `token`.
     */
    address public wrapped;

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                        CONSTRUCTOR                         */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    constructor() {
        _disableInitializers();
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                        INITIALIZER                         */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    function initialize(address admin_, address pauser_, address token_, address wrapped_) external initializer {
        // Validate required inputs are not zero addresses.
        if (admin_ == address(0)) revert ZeroAddress();
        if (pauser_ == address(0)) revert ZeroAddress();
        if (token_ == address(0)) revert ZeroAddress();
        if (wrapped_ == address(0)) revert ZeroAddress();

        // Initialize everything and grant roles.
        __AccessControl_init();
        __Pausable_init();
        _grantRole(DEFAULT_ADMIN_ROLE, admin_);
        _grantRole(PAUSER_ROLE, pauser_);

        // Set configured values.
        token = token_;
        wrapped = wrapped_;
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                     LOCKBOX FUNCTIONS                      */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Deposit tokens into the lockbox.
     * @param value The amount of tokens to deposit.
     */
    function deposit(uint256 value) external whenNotPaused {
        _deposit(msg.sender, msg.sender, value);
    }

    /**
     * @dev Deposit tokens into the lockbox for a specific address.
     * @param to    The address to deposit the tokens for.
     * @param value The amount of tokens to deposit.
     */
    function depositTo(address to, uint256 value) external whenNotPaused {
        _deposit(msg.sender, to, value);
    }

    /**
     * @dev Withdraws tokens from the lockbox.
     * @param value The amount of tokens to withdraw.
     */
    function withdraw(uint256 value) external whenNotPaused {
        _withdraw(msg.sender, msg.sender, value);
    }

    /**
     * @dev Withdraws tokens from the lockbox to a specific address.
     * @param to    The address to withdraw the tokens to.
     * @param value The amount of tokens to withdraw.
     */
    function withdrawTo(address to, uint256 value) external whenNotPaused {
        _withdraw(msg.sender, to, value);
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                      ADMIN FUNCTIONS                       */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Pauses depositing into and withdrawing from the lockbox.
     */
    function pause() external onlyRole(PAUSER_ROLE) {
        _pause();
    }

    /**
     * @dev Unpauses depositing into and withdrawing from the lockbox.
     */
    function unpause() external onlyRole(PAUSER_ROLE) {
        _unpause();
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                     INTERNAL FUNCTIONS                     */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Internal function to deposit tokens into the lockbox.
     * @param from  The address of the depositor.
     * @param to    The address to deposit the tokens for.
     * @param value The amount of tokens to deposit.
     */
    function _deposit(address from, address to, uint256 value) internal {
        token.safeTransferFrom(from, address(this), value);
        ITokenOps(wrapped).mint(to, value);
    }

    /**
     * @dev Internal function to withdraw tokens from the lockbox.
     * @param from  The address of the withdrawer.
     * @param to    The address to withdraw the tokens to.
     * @param value The amount of tokens to withdraw.
     */
    function _withdraw(address from, address to, uint256 value) internal {
        ITokenOps(wrapped).clawback(from, value);
        token.safeTransfer(to, value);
    }
}
