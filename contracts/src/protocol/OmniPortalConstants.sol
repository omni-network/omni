// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

/**
 * @title OmniPortalConstants
 * @notice Constants used by the OmniPortal contract
 */
contract OmniPortalConstants {
    /// @notice Default xmsg execution gas limit, enforced on destination chain
    uint64 public constant XMSG_DEFAULT_GAS_LIMIT = 200_000;

    /// @notice Maximum allowed xmsg gas limit
    uint64 public constant XMSG_MAX_GAS_LIMIT = 5_000_000;

    /// @notice Minimum allowed xmsg gas limit
    uint64 public constant XMSG_MIN_GAS_LIMIT = 21_000;

    /// @notice Numerator of the fraction of total validator power required to accept an XSubmission. Ex 2/3 -> 2
    uint8 public constant XSUB_QUORUM_NUMERATOR = 2;

    /// @notice Denominator of the fraction of total validator power required to accept an XSubmission. Ex 2/3 -> 3
    uint8 public constant XSUB_QUORUM_DENOMINATOR = 3;

    // TODO: make gas limits admin-configurable
    // TODO: make quorum admin-configurable
}
