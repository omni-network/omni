// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { OwnableUpgradeable } from "@openzeppelin-v4/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { PausableUpgradeable } from "@openzeppelin-v4/contracts-upgradeable/security/PausableUpgradeable.sol";
import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import { IGenesisStake } from "../interfaces/IGenesisStake.sol";

contract GenesisClaim is OwnableUpgradeable, PausableUpgradeable {
    /**
     * @notice Emitted when rewards are set for an account.
     * @param account   The account that received rewards.
     * @param amount    The amount of rewards received.
     */
    event RewardsSet(address indexed account, uint256 amount);

    /**
     * @notice Emitted when an account claims rewards.
     * @param account   The account that claimed rewards.
     * @param amount    The amount of rewards claimed.
     */
    event Claimed(address indexed account, uint256 amount);

    /**
     * @notice Emitted when the GenesisStake contract is set.
     * @param newGenesisStake   The new GenesisStake contract address.
     * @param prevGenesisStake  The previous GenesisStake contract address.
     */
    event GenesisStakeChanged(address indexed newGenesisStake, address indexed prevGenesisStake);

    /**
     * @notice Emitted when the Clique signer address is changed.
     * @param newClique The new Clique signer address.
     * @param prevClique The previous Clique signer address.
     */
    event CliqueChanged(address indexed newClique, address indexed prevClique);

    /**
     * @notice Emitted when the Clique fee is changed.
     * @param newCliqueFee The new Clique fee.
     * @param prevCliqueFee The previous Clique fee.
     */
    event CliqueFeeChanged(uint256 newCliqueFee, uint256 prevCliqueFee);

    /**
     * @notice Emitted when claims are opened.
     */
    event ClaimsOpened();

    /**
     * @notice Emitted when claims are closed.
     */
    event ClaimsClosed();

    /**
     * @notice The token contract address.
     */
    IERC20 public immutable token;

    /**
     * @notice The Clique signer address.
     */
    address public clique;

    /**
     * @notice The GenesisStake contract address.
     */
    IGenesisStake public genesisStake;

    /**
     * @notice True if claims are open, false otherwise.
     */
    bool public isOpen;

    /**
     * @notice Timestamp at which claims were opened.
     */
    uint256 public openedAt;

    /**
     * @notice Clique fee
     */
    uint256 public cliqueFee;

    /**
     * @notice Rewards for each account.
     */
    mapping(address => uint256) public rewards;

    /**
     * @notice True if rewards have been set for an account, false otherwise.
     */
    mapping(address => bool) public rewardsSet;

    /**
     * @notice Restrict calls to the Clique address.
     */
    modifier onlyClique() {
        require(msg.sender == clique, "GenesisClaim: only clique");
        _;
    }

    /**
     * @notice Restrict calls to when claims are open.
     */
    modifier whenOpen() {
        require(isOpen, "GenesisClaim: not open");
        _;
    }

    constructor(address token_) {
        token = IERC20(token_);
        _disableInitializers();
    }

    /**
     * @notice Initialize the contract.
     * @param owner_        The owner of the contract.
     * @param genesisStake_ The GenesisStake contract address.
     */
    function initialize(address owner_, address genesisStake_, address clique_, uint256 cliqueFee_, bool isOpen_)
        external
        initializer
    {
        __Ownable_init();
        __Pausable_init();
        _transferOwnership(owner_);

        _setGenesisStake(genesisStake_);
        _setClique(clique_);
        _setCliqueFee(cliqueFee_);

        if (isOpen_) _openClaims();
    }

    /**
     * @notice Set the rewards for `account`.
     * @param account   The account to set rewards for.
     * @param amount    The amount of rewards to set.
     */
    function setRewards(address account, uint256 amount) external onlyClique whenNotPaused {
        _setRewards(account, amount);
    }

    /**
     * @notice Set rewards for multiple accounts.
     * @param accounts  The accounts to set rewards for.
     * @param amounts   The amounts of rewards to set.
     */
    function batchSetRewards(address[] calldata accounts, uint256[] calldata amounts)
        external
        onlyClique
        whenNotPaused
    {
        require(accounts.length == amounts.length, "GenesisClaim: length mismatch");

        for (uint256 i = 0; i < accounts.length; i++) {
            _setRewards(accounts[i], amounts[i]);
        }
    }

    function _setRewards(address account, uint256 amount) internal {
        require(amount > 0, "GenesisClaim: amount must be > 0");
        require(!rewardsSet[account], "GenesisClaim: already set");
        require(account != address(0), "GenesisClaim: no zero address");

        rewards[account] = amount;
        rewardsSet[account] = true;

        emit RewardsSet(account, amount);
    }

    /**
     * @notice Reset rewards for `accounts`.
     * @param accounts  The accounts to reset rewards for.
     */
    function resetRewards(address[] calldata accounts) external onlyOwner {
        for (uint256 i = 0; i < accounts.length; i++) {
            // this check is included to avoid resetting rewards for accounts that have already claimed
            require(rewards[accounts[i]] > 0, "GenesisClaim: no rewards");

            rewards[accounts[i]] = 0;
            rewardsSet[accounts[i]] = false;
        }
    }

    /**
     * @notice Return true if `account` has rewards, false otherwise.
     * @param account   The account to check.
     */
    function hasRewards(address account) external view returns (bool) {
        return rewards[account] > 0;
    }

    /**
     * @notice Claim all rewards for the caller.
     */
    function claim() external payable whenNotPaused whenOpen {
        require(msg.value >= cliqueFee, "GenesisClaim: insufficient fee");

        uint256 amount = _markClaimed(msg.sender);

        IERC20(token).transfer(msg.sender, amount);

        emit Claimed(msg.sender, amount);
    }

    /**
     * @notice Claim all rewards for the caller, and stake them on behalf of the caller.
     */
    function claimAndStake() external payable whenNotPaused whenOpen {
        require(msg.value >= cliqueFee, "GenesisClaim: insufficient fee");

        uint256 amount = _markClaimed(msg.sender);

        genesisStake.stakeFor(msg.sender, amount);

        emit Claimed(msg.sender, amount);
    }

    /**
     * @notice Mark rewards as claimed for `account`. Returns the amount claimed.
     * @param account   The account claiming rewards.
     */
    function _markClaimed(address account) internal returns (uint256) {
        uint256 amount = rewards[account];
        require(amount > 0, "GenesisClaim: no rewards");

        rewards[account] = 0;

        return amount;
    }

    /**
     * @notice Open claims.
     * @dev Just sets openedAt timestamp.
     */
    function openClaims() external onlyOwner {
        _openClaims();
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
     * @notice Withdraw all unclaimed tokens to the `withdrawTo` address, and close claims.
     */
    function withdrawAndClose(address withdrawTo) external onlyOwner whenOpen {
        require(block.timestamp >= openedAt + 45 days, "GenesisClaim: not 45 days");
        isOpen = false;
        token.transfer(withdrawTo, token.balanceOf(address(this)));
        emit ClaimsClosed();
    }

    /**
     * @notice Withdraw clique fees to the `withdrawTo` address.
     */
    function withdrawFees(address payable withdrawTo) external onlyClique whenNotPaused {
        (bool success,) = withdrawTo.call{ value: address(this).balance }("");
        require(success, "GenesisClaim: withdrawal failed");
    }

    /**
     * @notice Set the GenesisStake contract address.
     * @param genesisStake_ The new GenesisStake contract address.
     */
    function setGenesisStake(address genesisStake_) external onlyOwner {
        _setGenesisStake(genesisStake_);
    }

    /**
     * @notice Set the Clique address.
     * @param clique_ The new Clique address.
     */
    function setClique(address clique_) external onlyOwner {
        _setClique(clique_);
    }

    /**
     * @notice Set the GenesisStake contract address.
     * @param genesisStake_ The new GenesisStake contract address.
     */
    function _setGenesisStake(address genesisStake_) internal {
        require(genesisStake_ != address(0), "GenesisClaim: no zero address");

        address prevGenesisStake = address(genesisStake);
        genesisStake = IGenesisStake(genesisStake_);

        // Approve new GenesisStake contract to transfer tokens.
        token.approve(genesisStake_, type(uint256).max);

        // Revoke approval from previous GenesisStake contract.
        if (prevGenesisStake != address(0)) token.approve(prevGenesisStake, 0);

        emit GenesisStakeChanged(genesisStake_, prevGenesisStake);
    }

    /**
     * @notice Set the Clique address.
     * @param clique_ The new Clique address.
     */
    function _setClique(address clique_) internal {
        require(clique_ != address(0), "GenesisClaim: no zero address");
        emit CliqueChanged(clique_, clique);
        clique = clique_;
    }

    /**
     * @notice Set the Clique fee.
     * @param cliqueFee_ The new Clique fee.
     */
    function _setCliqueFee(uint256 cliqueFee_) internal {
        emit CliqueFeeChanged(cliqueFee, cliqueFee_);
        cliqueFee = cliqueFee_;
    }

    /**
     * @notice Open claims.
     */
    function _openClaims() internal {
        require(!isOpen, "GenesisClaim: already open");
        isOpen = true;
        openedAt = block.timestamp;
        emit ClaimsOpened();
    }
}
