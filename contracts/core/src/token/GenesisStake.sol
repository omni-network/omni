// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { OwnableUpgradeable } from "@openzeppelin-v4/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { PausableUpgradeable } from "@openzeppelin-v4/contracts-upgradeable/security/PausableUpgradeable.sol";
import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import { IGenesisStake } from "../interfaces/IGenesisStake.sol";

/**
 * @title GenesisStake
 * @notice Omni's genesis staking contract. It allows
 */
contract GenesisStake is IGenesisStake, OwnableUpgradeable, PausableUpgradeable {
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
     * @notice Emitted when a user withdraws tokens.
     * @param account   The account that withdrew tokens.
     * @param amount    The amount of tokens withdrawn.
     */
    event Withdrawn(address indexed account, uint256 amount);

    /**
     * @notice Emitted when the unboding period is changed.
     * @param newDuration   The new unboding period.
     * @param prevDuration  The previous unboding period.
     */
    event UnbondingPeriodChanged(uint256 newDuration, uint256 prevDuration);

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
     * @notice Duration (in seconds) that a user must wait to withdraw after unstaking.
     */
    uint256 public unbondingPeriod;

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

    constructor(address token_) {
        token = IERC20(token_);
        _disableInitializers();
    }

    /**
     * @notice Initialize the contract.
     * @param owner_            The owner of the contract.
     * @param unbondingPeriod_  The unboding period.
     */
    function initialize(address owner_, uint256 unbondingPeriod_) external initializer {
        __Ownable_init();
        __Pausable_init();
        _transferOwnership(owner_);
        _setUnbondingPeriod(unbondingPeriod_);
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
     * @notice Unstake your entire balance.
     */
    function unstake() external whenNotPaused {
        require(balanceOf[msg.sender] > 0, "GenesisStake: not staked");
        require(unstakedAt[msg.sender] == 0, "GenesisStake: already unstaked");

        unstakedAt[msg.sender] = block.timestamp;

        emit Unstaked(msg.sender, balanceOf[msg.sender]);
    }

    /**
     * @notice Withdraw your entire balance after the unbonding period.
     */
    function withdraw() external whenNotPaused {
        require(balanceOf[msg.sender] > 0, "GenesisStake: not staked");
        require(unstakedAt[msg.sender] > 0, "GenesisStake: not unstaked");
        require(block.timestamp >= unstakedAt[msg.sender] + unbondingPeriod, "GenesisStake: not unbonded");

        uint256 amount = balanceOf[msg.sender];

        // reset balance & timestamps
        balanceOf[msg.sender] = 0;
        unstakedAt[msg.sender] = 0;

        require(token.transfer(msg.sender, amount), "GenesisStake: transfer failed");

        emit Withdrawn(msg.sender, amount);
    }

    /**
     * @notice Returns timestamp at which `account` can withdraw.
     *         Reverts if the account has not staked & unstaked.
     */
    function canWithdrawAt(address account) external view returns (uint256) {
        require(balanceOf[account] > 0, "GenesisStake: not staked");
        require(unstakedAt[account] > 0, "GenesisStake: not unstaked");

        return unstakedAt[account] + unbondingPeriod;
    }

    /**
     * @notice Set the unboding period.
     * @param duration The unboding period.
     */
    function setUnbondingPeriod(uint256 duration) external onlyOwner {
        _setUnbondingPeriod(duration);
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
     * @notice Set the unboding period.
     * @param duration The unboding period.
     */
    function _setUnbondingPeriod(uint256 duration) internal {
        require(duration > 0, "GenesisStake: dur must be > 0");
        uint256 prev = unbondingPeriod;
        unbondingPeriod = duration;
        emit UnbondingPeriodChanged(duration, prev);
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
}
