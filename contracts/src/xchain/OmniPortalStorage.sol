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
     * @notice Maxium number of bytes allowed in xreceipt result
     */
    uint16 public xreceiptMaxErrorBytes;

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
     * @notice Maps shard id to true, if the shard is supported.
     */
    mapping(uint64 => bool) public isSupportedShard;

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
    XTypes.MsgShort internal _xmsg;

    /**
     * @notice List of supported shards.
     */
    uint64[] internal _shards;
}
