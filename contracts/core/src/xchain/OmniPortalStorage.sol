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
     * @notice Number of validator sets since the latest that can be used to validate an XSubmission
     */
    uint8 public xsubValsetCutoff;

    /**
     * @notice Maxium number of bytes allowed in xreceipt result
     */
    uint16 public xreceiptMaxErrorSize;

    /**
     * @notice Maximum number of bytes allowed in xmsg data
     */
    uint16 public xmsgMaxDataSize;

    /**
     * @notice Maximum allowed xmsg gas limit
     */
    uint64 public xmsgMaxGasLimit;

    /**
     * @notice Minimum allowed xmsg gas limit
     */
    uint64 public xmsgMinGasLimit;

    /**
     * @notice ID of the latest validator set relayed to this portal from the consensus chain.
     */
    uint64 public latestValSetId;

    /**
     * @notice Chain ID of Omni's EVM execution chain
     */
    uint64 public omniChainId;

    /**
     * @notice Virtual Chain ID of Omni's consensus chain
     */
    uint64 public omniCChainId;

    /**
     * @notice The address of the fee oracle contract
     */
    address public feeOracle;

    /**
     * @notice A list of supported chains & shards.
     */
    XTypes.Chain[] public network;

    /**
     * @notice Maps shard ID to true, if the shard is supported.
     */
    mapping(uint64 => bool) public isSupportedShard;

    /**
     * @notice Maps chain ID to true, if the chain is supported.
     */
    mapping(uint64 => bool) public isSupportedDest;

    /**
     * @notice Offset of the last outbound XMsg that was sent to destChainId in shardId
     *         Maps destChainId -> shardId -> offset.
     */
    mapping(uint64 => mapping(uint64 => uint64)) public outXMsgOffset;

    /**
     * @notice Offset of the last outbound XMsg that was sent to destChainId in shardId
     *         Maps destChainId -> shardId -> offset.
     */
    mapping(uint64 => mapping(uint64 => uint64)) public inXMsgOffset;

    /**
     * @notice The xblock offset of the last inbound XMsg that was received from sourceChainId in shardIdj
     *         Maps sourceChainId -> shardId -> xblockOffset.
     */
    mapping(uint64 => mapping(uint64 => uint64)) public inXBlockOffset;

    /**
     * @notice Maps validator set id -> total power
     */
    mapping(uint64 => uint64) public valSetTotalPower;

    /**
     * @notice Maps validator set id -> validator address -> power
     */
    mapping(uint64 => mapping(address => uint64)) public valSet;

    /**
     * @notice The current XMsg being executed, exposed via xmsg() getter
     * @dev Internal state + public getter preferred over public state with default getter
     *      so that we can use the XMsg struct type in the interface.
     */
    XTypes.MsgContext internal _xmsg;
}
