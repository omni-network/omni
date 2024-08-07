// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { XApp } from "src/pkg/XApp.sol";
import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";
import { ConfLevel } from "src/libraries/ConfLevel.sol";

contract OmniFund is XApp, Ownable {
    /// @notice Map chainID to addr to true, if authorized to withdraw
    mapping(uint64 => mapping(address => bool)) public authed;

    /// @notice Map address to chainID to total funded
    mapping(address => mapping(uint64 => uint256)) public funded;

    constructor(address portal, address owner) XApp(portal, ConfLevel.Finalized) Ownable(owner) { }

    /**
     * @notice Try to withdraw remaining funds owned to `to`.
     *         The amount owed is `total - funded[to][xmsg.sourceChainId]`.
     */
    function tryWithdrawRemaining(address to, uint256 total) external xrecv {
        require(isXCall() && authed[xmsg.sourceChainId][xmsg.sender], "OmniFund: unauthorized");

        // we've already funded total requested
        require(total >= funded[to][xmsg.sourceChainId], "OmniFund: already funded");

        uint256 amt = total - funded[to][xmsg.sourceChainId];
        (bool success,) = to.call{ value: amt }("");

        // Only update funded if the transfer was successful
        // This allows the user to retry if the transfer fails
        // A transer may fail if this fund runs out of funds
        if (success) funded[to][xmsg.sourceChainId] += amt;
    }

    function authorize(uint64 chainID, address addr) external onlyOwner {
        authed[chainID][addr] = true;
    }

    receive() external payable { }
}
