// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { XTypes } from "../libraries/XTypes.sol";
import { IOmniPortal } from "../interfaces/IOmniPortal.sol";
import { IOmniPortalAdmin } from "../interfaces/IOmniPortalAdmin.sol";

/**
 * @title OmniPortalStorage
 * @notice Storage layout for OmniPortal
 */
abstract contract OmniPortalStorage is IOmniPortal, IOmniPortalAdmin {
    /**
     * @notice Track latest seen validator set id. Validator set ids increment monotonically
     */
    uint64 public latestValidatorSetId;

    /**
     * @notice The address of the fee oracle contract
     */
    address public feeOracle;

    /**
     * @notice Offset of the last outbound XMsg that was sent to destChainId
     */
    mapping(uint64 => uint64) public outXStreamOffset;

    /**
     * @notice Offset of the last outbound XMsg that was sent to destChainId
     */
    mapping(uint64 => uint64) public inXStreamOffset;

    /**
     * @notice Offset of the last inbound XMsg that was received from sourceChainId
     */
    mapping(uint64 => uint64) public inXStreamBlockHeight;

    /**
     * @notice Maps validator set id -> total power
     */
    mapping(uint64 => uint64) public validatorSetTotalPower;

    /**
     * @notice Maps validator set id -> validator address -> power
     */
    mapping(uint64 => mapping(address => uint64)) public validatorSet;

    /**
     * @notice The current XMsg being executed, exposed via xmsg() getter
     * @dev Internal state + public getter preferred over public state with default getter
     *      so that we can use the XMsg struct type in the interface.
     */
    XTypes.MsgShort internal _currentXmsg;
}
