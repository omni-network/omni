// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

/**
 * @title OmniPortalConstants
 * @notice Constants used by the OmniPortal contract
 */
contract OmniPortalConstants {
    /**
     * @notice Numerator of the fraction of total validator power required to accept an XSubmission. Ex 2/3 -> 2
     */
    uint8 public constant XSubQuorumNumerator = 2;

    /**
     * @notice Denominator of the fraction of total validator power required to accept an XSubmission. Ex 2/3 -> 3
     */
    uint8 public constant XSubQuorumDenominator = 3;

    /**
     * @notice Action ID for xsubmissions, used as Pauseable key
     */
    bytes32 public constant ActionXSubmit = keccak256("xsubmit");

    /**
     * @notice Action ID for xcalls, used as Pauseable key
     */
    bytes32 public constant ActionXCall = keccak256("xcall");

    /**
     * @dev xmsg.destChainId for "broadcast" xcalls, intended for all portals
     */
    uint64 internal constant BroadcastChainId = 0;

    /**
     * @dev xmg.sender for xmsgs from Omni's consensus chain
     */
    address internal constant CChainSender = address(0);

    /**
     * @dev xmsg.to for xcalls to be executed on the portal itself
     */
    address internal constant VirtualPortalAddress = address(0);
}
