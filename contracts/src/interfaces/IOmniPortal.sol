// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import { XTypes } from "../libraries/XTypes.sol";

/**
 * @title IOmniPortal
 * @notice The OmniPortal is the on-chain interface to Omni's cross-chain
 *         messaging protocol. It is used to initiate and execute cross-chain calls.
 */
interface IOmniPortal {
    /**
     * @notice Emitted when an xcall is made to a contract on another chain
     * @param destChainId Destination chain ID
     * @param streamOffset Offset this XMsg in the source -> dest XStream
     * @param sender msg.sender of the source xcall
     * @param to Address of the contract to call on the destination chain
     * @param gasLimit Gas limit for execution on destination chain
     * @param data Encoded function calldata
     */
    event XMsg(
        uint64 indexed destChainId, uint64 indexed streamOffset, address sender, address to, bytes data, uint64 gasLimit
    );

    /**
     * @notice Emitted when an XMsg is executed on its destination chain
     * @param sourceChainId Source chain ID
     * @param streamOffset Offset the XMsg in the source -> dest XStream
     * @param gasUsed Gas used in execution of the XMsg
     * @param success Whether the execution succeeded
     * @param relayer Address of the relayer who submitted the XMsg
     */
    event XReceipt(
        uint64 indexed sourceChainId, uint64 indexed streamOffset, uint256 gasUsed, address relayer, bool success
    );

    /**
     * @notice Emitted when a new validator set is added
     * @param setId Validator set ID
     */
    event ValidatorSetAdded(uint64 indexed setId);

    /**
     * @notice Returns Chain ID of the chain to which this portal is deployed
     */
    function chainId() external view returns (uint64);

    /**
     * @notice Returns the offset of the last outbound XMsg that was sent to destChainId
     * @param destChainId Destination chain ID
     */
    function outXStreamOffset(uint64 destChainId) external view returns (uint64);

    /**
     * @notice Returns the offset of the last inbound XMsg that was received from sourceChainId
     * @param sourceChainId Source chain ID
     */
    function inXStreamOffset(uint64 sourceChainId) external view returns (uint64);

    /**
     * @notice Returns the source block height of the last inbound XMsg that was received from
     *         sourceChainId
     * @param sourceChainId Source chain ID
     */
    function inXStreamBlockHeight(uint64 sourceChainId) external view returns (uint64);

    /**
     * @notice Returns the current XMsg being executed via this portal.
     *          - xmsg().sourceChainId  Chain ID of the source xcall
     *          - xmsg().sender         msg.sender of the source xcall
     *         If no XMsg is being executed, all fields will be zero.
     *          - xmsg().sourceChainId  == 0
     *          - xmsg().sender         == address(0)
     */
    function xmsg() external view returns (XTypes.MsgShort memory);

    /**
     * @notice Returns true the current transaction is an xcall, false otherwise
     */
    function isXCall() external view returns (bool);

    /**
     * @notice Calculate the fee for calling a contract on another chain. Uses
     *         OmniPortal.XMSG_DEFAULT_GAS_LIMIT. Fees denominated in wei.
     * @param destChainId   Destination chain ID
     * @param data          Encoded function calldata
     */
    function feeFor(uint64 destChainId, bytes calldata data) external view returns (uint256);

    /**
     * @notice Calculate the fee for calling a contract on another chain
     *         Fees denominated in wei.
     * @param destChainId   Destination chain ID
     * @param data          Encoded function calldata
     * @param gasLimit      Execution gas limit, enforced on destination chain
     */
    function feeFor(uint64 destChainId, bytes calldata data, uint64 gasLimit) external view returns (uint256);

    /**
     * @notice Call a contract on another chain Uses OmniPortal.XMSG_DEFAULT_GAS_LIMIT as execution
     *         gas limit on destination chain
     * @param destChainId   Destination chain ID
     * @param to            Address of contract to call on destination chain
     * @param data          ABI Encoded function calldata
     */
    function xcall(uint64 destChainId, address to, bytes calldata data) external payable;

    /**
     * @notice Call a contract on another chain Uses provide gasLimit as execution gas limit on
     *          destination chain. Reverts if gasLimit < XMSG_MAX_GAS_LIMIT or gasLimit >
     *          XMSG_MAX_GAS_LIMIT
     * @param destChainId   Destination chain ID
     * @param to            Address of contract to call on destination chain
     * @param data          ABI Encoded function calldata
     * @param gasLimit      Execution gas limit, enforced on destination chain
     */
    function xcall(uint64 destChainId, address to, bytes calldata data, uint64 gasLimit) external payable;

    /**
     * @notice Submit a batch of XMsgs to be executed on this chain
     * @param xsub  An xchain submisison, including an attestation root w/ validator signatures,
     *              and a block header and message batch, proven against the attestation root.
     */
    function xsubmit(XTypes.Submission calldata xsub) external;
}
