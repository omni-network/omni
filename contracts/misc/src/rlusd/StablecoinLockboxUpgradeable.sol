// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import { Initializable } from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import { UUPSUpgradeable } from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import { AccessControlUpgradeable } from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import { PausableUpgradeable } from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { IStablecoinLockboxUpgradeable } from "./interfaces/IStablecoinLockboxUpgradeable.sol";

contract StablecoinLockboxUpgradeable is
    Initializable,
    UUPSUpgradeable,
    AccessControlUpgradeable,
    PausableUpgradeable,
    IStablecoinLockboxUpgradeable
{
    using SafeTransferLib for address;

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                         CONSTANTS                          */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    //keccak256("UPGRADER")
    bytes32 public constant UPGRADER_ROLE = 0xa615a8afb6fffcb8c6809ac0997b5c9c12b8cc97651150f14c8f6203168cff4c;
    //keccak256("PAUSER")
    bytes32 public constant PAUSER_ROLE = 0x539440820030c4994db4e31b6b800deafd503688728f932addfe7a410515c14c;

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                          STORAGE                           */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Mapping of depositor to token to balance.
     * @return balance The balance of the depositor for the token.
     */
    mapping(address depositor => mapping(address token => uint256 balance)) public balances;

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                        CONSTRUCTOR                         */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                        INITIALIZER                         */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev This method is used to initialize the contract with values that we want to use to bootstrap and run things.
     * The modifier initializer here helps us block initialization in the constructor so that we initialize value only
     * when deploying the proxy and not the contract itself. The initializer also tracks how many times this method is
     * called and it can only be called once.
     */
    function initialize(address admin_, address upgrader_, address pauser_) external initializer {
        __UUPSUpgradeable_init();
        __AccessControl_init();
        __Pausable_init();
        _grantRole(DEFAULT_ADMIN_ROLE, admin_);
        _grantRole(UPGRADER_ROLE, upgrader_);
        _grantRole(PAUSER_ROLE, pauser_);
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                     LOCKBOX FUNCTIONS                      */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Deposit tokens into the lockbox.
     * @param token The address of the token to deposit.
     * @param value The amount of tokens to deposit.
     */
    function deposit(address token, uint256 value) external whenNotPaused {
        _deposit(token, msg.sender, msg.sender, value);
    }

    /**
     * @dev Deposit tokens into the lockbox for a specific address.
     * @param token The address of the token to deposit.
     * @param to The address to deposit the tokens for.
     * @param value The amount of tokens to deposit.
     */
    function depositTo(address token, address to, uint256 value) external whenNotPaused {
        _deposit(token, msg.sender, to, value);
    }

    /**
     * @dev Withdraws tokens from the lockbox.
     * @param token The address of the token to withdraw.
     * @param value The amount of tokens to withdraw.
     */
    function withdraw(address token, uint256 value) external whenNotPaused {
        _withdraw(token, msg.sender, msg.sender, value);
    }

    /**
     * @dev Withdraws tokens from the lockbox to a specific address.
     * @param token The address of the token to withdraw.
     * @param to The address to withdraw the tokens to.
     * @param value The amount of tokens to withdraw.
     */
    function withdrawTo(address token, address to, uint256 value) external whenNotPaused {
        _withdraw(token, msg.sender, to, value);
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                     INTERNAL FUNCTIONS                     */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Internal function to deposit tokens into the lockbox.
     * @param token The address of the token to deposit.
     * @param from The address of the depositor.
     * @param to The address to deposit the tokens for.
     * @param value The amount of tokens to deposit.
     */
    function _deposit(address token, address from, address to, uint256 value) internal {
        token.safeTransferFrom(from, address(this), value);
        unchecked {
            balances[to][token] += value;
        }
        emit Deposit(token, from, to, value);
    }

    /**
     * @dev Internal function to withdraw tokens from the lockbox.
     * @param token The address of the token to withdraw.
     * @param from The address of the withdrawer.
     * @param to The address to withdraw the tokens to.
     * @param value The amount of tokens to withdraw.
     */
    function _withdraw(address token, address from, address to, uint256 value) internal {
        uint256 balance = balances[from][token];
        if (balance < value) revert InsufficientBalance(token, from, balance, value);

        unchecked {
            balances[from][token] -= value;
        }
        token.safeTransfer(to, value);

        emit Withdraw(token, from, to, value);
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                    OVERRIDDEN FUNCTIONS                    */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * An overridden method from {UUPSUpgradeable} which defines the permissions for authorizing an upgrade to a
     * new implementation.
     */
    function _authorizeUpgrade(address newImplementation) internal virtual override onlyRole(UPGRADER_ROLE) { }
}
