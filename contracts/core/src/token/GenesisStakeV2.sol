// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { OwnableUpgradeable } from "@openzeppelin-v4/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { PausableUpgradeable } from "@openzeppelin-v4/contracts-upgradeable/security/PausableUpgradeable.sol";
import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import { IGenesisStakeV2 } from "../interfaces/IGenesisStakeV2.sol";

/**
 * @title GenesisStakeV2
 * @notice Omni's genesis staking contract
 */
contract GenesisStakeV2 is IGenesisStakeV2, OwnableUpgradeable, PausableUpgradeable {
    /**
     * @notice Emitted when a user migrates their stake.
     * @param account   The account that migrated their stake.
     * @param amount    The amount of tokens migrated.
     */
    event Migrated(address indexed account, uint256 amount);

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
     * @notice The unbonding period (deprecated).
     * @dev This variable is kept for storage layout compatibility.
     */
    uint256 private _deprecated_unbondingPeriod;

    /**
     * @notice The staked balance of each user.
     */
    mapping(address => uint256) public balanceOf;

    /**
     * @notice The timestamp at which each user unstaked (deprecated).
     * @dev This variable is kept for storage layout compatibility.
     */
    mapping(address => uint256) private _deprecated_unstakedAt;

    /**
     * @notice True if staking is open, false otherwise (deprecated).
     * @dev This variable is kept for storage layout compatibility.
     */
    bool private _deprecated_isOpen;

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
    }

    /**
     * @notice Migrate a user's stake to the rewards distributor.
     * @param addr The address of the user to migrate.
     * @return The amount of tokens migrated.
     */
    function migrateStake(address addr) external whenNotPaused returns (uint256) {
        require(msg.sender == rewardsDistributor, "GenesisStakeV2: unauthorized");

        uint256 amount = balanceOf[addr];
        if (amount == 0) return amount;

        // reset balance
        balanceOf[addr] = 0;

        require(token.transfer(rewardsDistributor, amount), "GenesisStakeV2: transfer failed");

        emit Migrated(addr, amount);
        return amount;
    }

    /**
     * @notice Stake `amount` tokens for `recipient`, paid by the caller.
     * @param recipient The recipient to stake tokens for.
     * @param amount    The amount of tokens to stake.
     */
    function stakeFor(address recipient, uint256 amount) external whenNotPaused {
        require(recipient != address(0), "GenesisStakeV2: no zero address");
        require(amount > 0, "GenesisStakeV2: amount must be > 0");
        
        // Transfer tokens from caller to this contract
        require(token.transferFrom(msg.sender, address(this), amount), "GenesisStakeV2: transfer failed");
        
        // Update recipient's balance
        balanceOf[recipient] += amount;
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

    fallback() external { }
}
