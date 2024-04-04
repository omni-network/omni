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
     * @notice  Chain ID of Omni's EVM execution chain
     */
    uint64 public omniEChainId;

    /**
     * @notice Virtual chain ID used in xmsgs from Omni's consensus chain
     */
    uint64 public omniCChainID;

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
     * @notice Soure block height of the last XSubmission that was received from sourceChainId
     */
    mapping(uint64 => uint64) public inXStreamBlockHeight;

    /**
     * @notice Validator set id of the last XSubmission that was received from sourceChainId
     */
    mapping(uint64 => uint64) public inXStreamValidatorSetId;

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
    XTypes.MsgShort internal _xmsg;
}
