// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { OwnableUpgradeable } from "@openzeppelin-v4/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { PausableUpgradeable } from "@openzeppelin-v4/contracts-upgradeable/security/PausableUpgradeable.sol";
import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import { IGenesisStakeV2 } from "../interfaces/IGenesisStake.sol";

/**
 * @title GenesisStake
 * @notice Omni's genesis staking contract. It allows
 */
contract GenesisStakeV2 is IGenesisStakeV2, OwnableUpgradeable, PausableUpgradeable {
    /**
     * @notice Emitted when an account stakes.
     * @param recipient The recipient of the stake.
     * @param amount    The amount of tokens staked.
     */
    event Staked(address indexed recipient, uint256 amount);

    /**
     * @notice Emitted when a user unstakes tokens.
     * @param account   The account that unstaked tokens.
     * @param amount    The amount of tokens unstaked.
     */
    event Unstaked(address indexed account, uint256 amount);

    /**
     * @notice Emitted when a user migrates their stake.
     * @param account   The account that migrated their stake.
     * @param amount    The amount of tokens migrated.
     */
    event Migrated(address indexed account, uint256 amount);

    /**
     * @notice Emitted when staking is opened.
     */
    event Opened();

    /**
     * @notice Emitted when staking is closed.
     */
    event Closed();

    /**
     * @notice Omni erc20 token.
     */
    IERC20 public immutable token;

    /**
     * @notice The rewards distributor for the staking contract.
     * @dev This contract is allowed to withdraw user staking deposits for migration to Omni.
     */
    address public immutable rewardsDistributor;
    /**
     * @notice The staked balance of each user.
     */
    mapping(address => uint256) public balanceOf;

    /**
     * @notice The timestamp at which each user unstaked.
     */
    mapping(address => uint256) public unstakedAt;

    /**
     * @notice True is staking is open, false otherwise.
     */
    bool public isOpen;

    /**
     * @notice Restrict function to when staking is open.
     */
    modifier whenOpen() {
        require(isOpen, "GenesisStake: not open");
        _;
    }

    constructor(address token_, address rewardsDistributor_) {
        token = IERC20(token_);
        rewardsDistributor = rewardsDistributor_;
        _disableInitializers();
    }

    /**
     * @notice Initialize the contract.
     * @param owner_            The owner of the contract.
     */
    function initialize(address owner_) external initializer {
        __Ownable_init();
        __Pausable_init();
        _transferOwnership(owner_);
        _open();
    }

    /**
     * @notice Stake `amount` tokens for `user`, paid by the caller.
     * @param recipient The recipient of the stake.
     * @param amount    The amount of tokens to stake.
     */
    function stakeFor(address recipient, uint256 amount) external whenNotPaused whenOpen {
        _stake(recipient, msg.sender, amount);
    }

    /**
     * @notice Stake `amount` tokens for the caller.
     * @param amount    The amount of tokens to stake.
     */
    function stake(uint256 amount) external whenNotPaused whenOpen {
        _stake(msg.sender, msg.sender, amount);
    }

    /**
     * @notice Internal function to stake `amount` tokens for `recipient`, paid by `patron`.
     * @param recipient The recipient of the stake.
     * @param patron    The account paying for the stake.
     * @param amount    The amount of tokens to stake.
     */
    function _stake(address recipient, address patron, uint256 amount) internal {
        require(amount > 0, "GenesisStake: amount must be > 0");
        require(unstakedAt[recipient] == 0, "GenesisStake: unstaked");

        balanceOf[recipient] += amount;

        require(token.transferFrom(patron, address(this), amount), "GenesisStake: transfer failed");

        emit Staked(recipient, amount);
    }

    /**
     * @notice Migrate a user's stake to the rewards distributor.
     * @param addr The address of the user to migrate.
     * @return The amount of tokens migrated.
     */
    function migrateStake(address addr) external whenNotPaused returns (uint256) {
        require(msg.sender == rewardsDistributor, "GenesisStake: unauthorized");

        uint256 amount = balanceOf[addr];
        if (amount == 0) return amount;

        // reset balance & timestamps
        _resetValues(addr);

        require(token.transfer(rewardsDistributor, amount), "GenesisStake: transfer failed");

        emit Migrated(addr, amount);
        return amount;
    }

    /**
     * @notice Pause the contract.
     */
    function pause() external onlyOwner {
        _pause();
    }

    /**
     * @notice Unpause the contract.
     */
    function unpause() external onlyOwner {
        _unpause();
    }

    /**
     * @notice Open staking.
     */
    function open() external onlyOwner {
        _open();
    }

    /**
     * @notice Close staking.
     */
    function close() external onlyOwner {
        _close();
    }

    /**
     * @notice Reset the balance and unstaked timestamp for an address.
     * @param addr The address to reset.
     */
    function _resetValues(address addr) internal {
        balanceOf[addr] = 0;
        unstakedAt[addr] = 0;
    }

    /**
     * @notice Open staking.
     */
    function _open() internal {
        require(!isOpen, "GenesisStake: already open");
        isOpen = true;
        emit Opened();
    }

    /**
     * @notice Close staking.
     */
    function _close() internal {
        require(isOpen, "GenesisStake: already closed");
        isOpen = false;
        emit Closed();
    }

    fallback() external { }
}
