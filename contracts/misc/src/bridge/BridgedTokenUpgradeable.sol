// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import { Initializable } from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import { ERC20Upgradeable } from "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";
import { UUPSUpgradeable } from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import { AccessControlUpgradeable } from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import { ERC20PausableUpgradeable } from
    "@openzeppelin/contracts-upgradeable/token/ERC20/extensions/ERC20PausableUpgradeable.sol";
import { AccountPausableUpgradeable } from "./libraries/AccountPausableUpgradeable.sol";
import { IBridgedTokenUpgradeable } from "./interfaces/IBridgedTokenUpgradeable.sol";

contract BridgedTokenUpgradeable is
    Initializable,
    ERC20Upgradeable,
    UUPSUpgradeable,
    AccessControlUpgradeable,
    ERC20PausableUpgradeable,
    AccountPausableUpgradeable,
    IBridgedTokenUpgradeable
{
    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                         CONSTANTS                          */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    //keccak256("MINTER")
    bytes32 public constant MINTER_ROLE = 0xf0887ba65ee2024ea881d91b74c2450ef19e1557f03bed3ea9f16b037cbe2dc9;
    //keccak256("UPGRADER")
    bytes32 public constant UPGRADER_ROLE = 0xa615a8afb6fffcb8c6809ac0997b5c9c12b8cc97651150f14c8f6203168cff4c;
    //keccak256("PAUSER")
    bytes32 public constant PAUSER_ROLE = 0x539440820030c4994db4e31b6b800deafd503688728f932addfe7a410515c14c;
    //keccak256("BURNER")
    bytes32 public constant BURNER_ROLE = 0x9667e80708b6eeeb0053fa0cca44e028ff548e2a9f029edfeac87c118b08b7c8;
    //keccak256("CLAWBACKER")
    bytes32 public constant CLAWBACKER_ROLE = 0x715bacafb7a853b9b91b59ae724920a9eb0c006c5b318ac393fa1bc8974edd98;

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
    function initialize(
        string memory name_,
        string memory symbol_,
        address bridge_,
        address admin_,
        address upgrader_,
        address pauser_,
        address clawbacker_
    ) external initializer {
        __ERC20_init(name_, symbol_);
        __UUPSUpgradeable_init();
        __AccessControl_init();
        _grantRole(DEFAULT_ADMIN_ROLE, admin_);
        _grantRole(MINTER_ROLE, bridge_);
        _grantRole(BURNER_ROLE, bridge_);
        _grantRole(UPGRADER_ROLE, upgrader_);
        __ERC20Pausable_init();
        __AccountPausable_init();
        _grantRole(PAUSER_ROLE, pauser_);
        _grantRole(CLAWBACKER_ROLE, clawbacker_);
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                   MANAGEMENT FUNCTIONS                     */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Creates a `value` amount of tokens and assigns them to `to`, by transferring it from address(0).
     * Relies on the `_update` mechanism
     *
     * @param to    The address that will be receiving the minted amount.
     * @param value The amount of tokens that are being minted to the account.
     *
     * Emits a {Transfer} event with `source` set to the zero address.
     */
    function mint(address to, uint256 value) public virtual onlyRole(MINTER_ROLE) {
        _mint(to, value);
    }

    /**
     * @dev Destroys a `value` amount of tokens from `msg.sender`, lowering the total supply.
     * Relies on the `_update` mechanism.
     *
     * @param value The amount of tokens that are being burned from message sender's holdings.
     *
     * Emits a {Transfer} event with `destination` set to the zero address.
     */
    function burn(uint256 value) public virtual onlyRole(BURNER_ROLE) {
        _burn(msg.sender, value);
    }

    /**
     * @dev Destroys a `value` amount of tokens from `from` account, lowering the total supply.
     * Relies on the `_update` mechanism.
     *
     * @param from  The address from which the tokens will be burned.
     * @param value The amount of tokens that will be burned.
     *
     * Emits a {Transfer} event with `destination` set to the zero address.
     */
    function clawback(address from, uint256 value) public virtual onlyRole(CLAWBACKER_ROLE) {
        _burn(from, value);
    }

    /**
     * Allow an authorized minter to pause the circulation of ERC20 tokens from all accounts.
     *
     * Requirements:
     * - The contract must not be paused.
     */
    function pause() public virtual onlyRole(PAUSER_ROLE) {
        _pause();
    }

    /**
     * Allow an authorized minter to resume/unpause the circulation of ERC20 tokens from all account.
     *
     * Requirements:
     * - The contract must be paused.
     */
    function unpause() public virtual onlyRole(PAUSER_ROLE) {
        _unpause();
    }

    /**
     * Allow an authorized minter to pause the circulation of ERC20 tokens from a specified account.
     *
     * @param accounts An array of addresses that will be paused, restricting them from taking any value moving actions.
     *
     * Requirements:
     * - accounts in the {accounts} list should be unpaused
     */
    function pauseAccounts(address[] calldata accounts) public virtual onlyRole(PAUSER_ROLE) {
        address lastAdd = address(0);
        uint256 accountsLength = accounts.length;
        for (uint256 i = 0; i < accountsLength; ++i) {
            require(accounts[i] > lastAdd, "Addresses should be sorted");
            _pauseAccount(accounts[i]);
            lastAdd = accounts[i];
        }
    }

    /**
     * Allow an authorized minter to resume/unpause the circulation of ERC20 tokens from a specified account.
     *
     * @param account The address that is being unpaused.
     *
     * Requirements:
     * - the {account} should be paused
     */
    function unpauseAccount(address account) public virtual onlyRole(PAUSER_ROLE) {
        _unpauseAccount(account);
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                    OVERRIDDEN FUNCTIONS                    */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * An overridden method from {UUPSUpgradeable} which defines the permissions for authorizing an upgrade to a
     * new implementation.
     */
    function _authorizeUpgrade(address newImplementation) internal virtual override onlyRole(UPGRADER_ROLE) { }

    /**
     * An overridden method to add modifiers to check if the accounts being used to transfer are not frozen.
     *
     * Requirements:
     * - the {from} account should not be paused/frozen
     * - the {to} account should not be paused/frozen
     * - the {msg.sender} account should not be paused/frozen
     */
    function _update(address from, address to, uint256 value)
        internal
        override(ERC20Upgradeable, ERC20PausableUpgradeable)
        whenAccountNotPaused(from)
        whenAccountNotPaused(to)
        whenAccountNotPaused(_msgSender())
    {
        ERC20PausableUpgradeable._update(from, to, value);
    }

    /**
     * An overridden method to block approvals when contract or accounts are paused.
     *
     * @param spender The account being approved by message sender to spend from their holdings.
     * @param value   The number of tokens the spender is being allowed to spend.
     *
     * @return A boolean indicating if the approval was successful or not.
     */
    function approve(address spender, uint256 value)
        public
        virtual
        override(ERC20Upgradeable)
        whenAccountNotPaused(spender)
        whenAccountNotPaused(_msgSender())
        whenNotPaused
        returns (bool)
    {
        return super.approve(spender, value);
    }
}
