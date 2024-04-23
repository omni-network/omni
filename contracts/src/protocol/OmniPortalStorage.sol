// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { IOmniPortal } from "../interfaces/IOmniPortal.sol";
import { IOmniPortalAdmin } from "../interfaces/IOmniPortalAdmin.sol";
import { XTypes } from "../libraries/XTypes.sol";

/**
 * @title OmniPortalStorage
 * @notice Storage layout for OmniPortal
 */
abstract contract OmniPortalStorage is IOmniPortal, IOmniPortalAdmin {
    /**
     * @notice Default xmsg execution gas limit, enforced on destination chain
     */
    uint64 public xmsgDefaultGasLimit;

    /**
     * @notice Maximum allowed xmsg gas limit
     */
    uint64 public xmsgMaxGasLimit;

    /**
     * @notice Minimum allowed xmsg gas limit
     */
    uint64 public xmsgMinGasLimit;

    /**
     * @notice Maxium number of bytes allowed in xreceipt result
     */
    uint64 public xreceiptMaxErrorBytes;

    /**
     * @notice  Chain ID of Omni's EVM execution chain
     */
    uint64 public omniChainId;

    /**
     * @notice Virtual chain ID used in xmsgs from Omni's consensus chain
     */
    uint64 public omniCChainID;

    /**
     * @notice The address of the fee oracle contract
     */
    address public feeOracle;

    /**
     * @notice The address of the XRegistry replica contract on this chain
     */
    address public xregistry;

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
