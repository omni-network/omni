// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { XAppBase } from "core/src/pkg/XAppBase.sol";

/**
 * @title Pledge
 * @notice Receives pledges to Ethereum from users on other chains.
 */
contract Pledge is XAppBase {
    /**
     * @notice Thrown when the caller is not the OmniPortal.
     */
    error NotXCall();

    /**
     * @notice Emitted when a user pledges to Ethereum.
     * @param from The address of the user who pledged.
     * @param chainId The chain ID of the source chain.
     * @param timestamp The timestamp the pledge was received.
     */
    event Pledged(bytes32 indexed from, uint256 indexed chainId, uint256 indexed timestamp);

    /**
     * @param omni_ The address of the OmniPortal.
     */
    constructor(address omni_) payable {
        _setOmniPortal(omni_);
    }

    /**
     * @notice Initiates a user's pledge to Ethereum.
     * @dev Only the OmniPortal can call this function, user cannot call it directly.
     * @dev Function selector is mined to be 0x00000000
     */
    function pledge_jwkilcxtschdbaaa() external xrecv {
        // only accept xcalls
        if (!isXCall()) revert NotXCall();
        // emit pledge event, frontend will contextualize
        emit Pledged(xmsg.sender, xmsg.sourceChainId, block.timestamp);
    }
}
