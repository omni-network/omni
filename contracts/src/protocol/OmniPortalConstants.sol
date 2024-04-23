// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

/**
 * @title OmniPortalConstants
 * @notice Constants used by the OmniPortal contract
 */
contract OmniPortalConstants {
    /// @notice Numerator of the fraction of total validator power required to accept an XSubmission. Ex 2/3 -> 2
    uint8 public constant XSUB_QUORUM_NUMERATOR = 2;

    /// @notice Denominator of the fraction of total validator power required to accept an XSubmission. Ex 2/3 -> 3
    uint8 public constant XSUB_QUORUM_DENOMINATOR = 3;

    /// @notice Default xmsg execution gas limit, enforced on destination chain
    uint64 public constant XMSG_DEFAULT_GAS_LIMIT = 200_000;

    /// @notice Maximum allowed xmsg gas limit
    uint64 public constant XMSG_MAX_GAS_LIMIT = 5_000_000;

    /// @notice Minimum allowed xmsg gas limit
    uint64 public constant XMSG_MIN_GAS_LIMIT = 21_000;

    /// @notice Maxium number of bytes allowed in xreceipt result
    uint64 public constant XRECEIPT_MAX_ERROR_BYTES = 256;

    /// @notice Error code that indicates xmsg error result exceeds XRECEIPT_MAX_ERROR_BYTES
    /// @dev Error code is abi.encodeWithSignature("OmniError(uint256)", 1)
    bytes public constant XRECEIPT_ERROR_EXCEEDS_MAX_BYTES =
        hex"0ea804990000000000000000000000000000000000000000000000000000000000000001";

    // TODO: make gas limits admin-configurable
    // TODO: make quorum admin-configurable

    // @dev xmsg.destChainId for "broadcast" xcalls, intended for all portals
    uint64 internal constant _BROADCAST_CHAIN_ID = 0;

    // @dev xmg.sender for xmsgs from Omni's consensus chain
    address internal constant _CCHAIN_SENDER = address(0);

    // @dev xmsg.to for xcalls to be executed on the portal itself
    address internal constant _VIRTUAL_PORTAL_ADDRESS = address(0);
}
