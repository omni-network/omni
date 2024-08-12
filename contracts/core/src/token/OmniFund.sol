// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { XTypes } from "src/libraries/XTypes.sol";
import { IOmniPortal } from "src/interfaces/IOmniPortal.sol";
import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";

contract OmniFund is Ownable {
    /// @notice Map chainID to addr to true, if authorized to withdraw
    mapping(uint64 => mapping(address => bool)) public authed;

    /// @notice Map address to chainID to total funded
    mapping(address => mapping(uint64 => uint256)) public funded;

    IOmniPortal public omni;

    constructor(address portal, address owner) Ownable(owner) {
        omni = IOmniPortal(portal);
    }

    /**
     * @notice Try to withdraw remaining funds owned to `to`.
     *         The amount owed is `total - funded[to][xmsg.sourceChainId]`.
     * @param recipient     Address to receive the funds
     * @param total         Total (historical) amount requested for `recipient`
     */
    function tryWithdrawRemaining(address recipient, uint256 total) external {
        XTypes.MsgContext memory xmsg = omni.xmsg();
        require(msg.sender == address(omni) && authed[xmsg.sourceChainId][xmsg.sender], "OmniFund: unauthorized");

        // we've already funded total requested
        require(total >= funded[recipient][xmsg.sourceChainId], "OmniFund: already funded");

        uint256 amt = total - funded[recipient][xmsg.sourceChainId];
        (bool success,) = recipient.call{ value: amt }("");

        // Only update funded if the transfer was successful
        // This allows the user to retry if the transfer fails
        // A transer may fail if this fund runs out of funds
        if (success) funded[recipient][xmsg.sourceChainId] += amt;
    }

    function authorize(uint64 chainID, address addr) external onlyOwner {
        authed[chainID][addr] = true;
    }

    receive() external payable { }
}
