// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { PausableUpgradeable } from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { XAppUpgradeable } from "src/pkg/XAppUpgradeable.sol";
import { ConfLevel } from "src/libraries/ConfLevel.sol";

/**
 * @title OmniXFund
 * @notice A cross-chain fund, allows authorized contracts on other chains to withdraw funds.
 */
contract OmniXFund is XAppUpgradeable, OwnableUpgradeable, PausableUpgradeable {
    event SettleUp(
        address indexed recipient,
        uint64 indexed chainID,
        address indexed from,
        uint256 owed,
        uint256 funded,
        bool success
    );

    //// @notice Map chainID to addr to true, if authorized to withdraw
    mapping(uint64 => mapping(address => bool)) public authed;

    /// @notice Map recipient to chainID (xmsg.sourceChainId) to authed creditor (xmsg.sender) to total funded
    mapping(address => mapping(uint64 => mapping(address => uint256))) public funded;

    constructor() {
        _disableInitializers();
    }

    function initialize(address portal, address owner) external initializer {
        __XApp_init(portal, ConfLevel.Finalized);
        __Ownable_init(owner);
    }

    /**
     * @notice Settle up with a recipient. If `total` is more than what they've been
     *         funded, transfer the difference.
     * @param recipient     Address to receive the funds
     * @param owed         Total amount owed to `recipient` by way of `xmsg.sender`
     */
    function settleUp(address recipient, uint256 owed) external xrecv whenNotPaused {
        require(isXCall() && authed[xmsg.sourceChainId][xmsg.sender], "OmniFund: unauthorized");
        require(owed >= funded[recipient][xmsg.sourceChainId][xmsg.sender], "OmniFund: already funded");

        uint256 amt = owed - funded[recipient][xmsg.sourceChainId][xmsg.sender];
        (bool success,) = recipient.call{ value: amt }("");

        if (success) funded[recipient][xmsg.sourceChainId][xmsg.sender] += amt;

        emit SettleUp(
            recipient,
            xmsg.sourceChainId,
            xmsg.sender,
            owed,
            funded[recipient][xmsg.sourceChainId][xmsg.sender],
            success
        );
    }

    /// @notice Authorize `addr` on `chainId` to withdraw funds via xcall
    function authorize(uint64 chainID, address addr) external onlyOwner {
        authed[chainID][addr] = true;
    }

    /// @notice Pause withdrawals
    function pause() external onlyOwner {
        _pause();
    }

    /// @notice Unpause withdrawals
    function unpause() external onlyOwner {
        _unpause();
    }

    receive() external payable { }
}
