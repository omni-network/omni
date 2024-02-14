// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.0;

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
     * @notice Default xmsg execution gas limit, enforced on destination chain
     * @return Gas limit
     */
    function XMSG_DEFAULT_GAS_LIMIT() external view returns (uint64);

    /**
     * @notice Maximum allowed xmsg gas limit
     * @return Maximum gas limit
     */
    function XMSG_MAX_GAS_LIMIT() external view returns (uint64);

    /**
     * @notice Minimum allowed xmsg gas limit
     * @return Minimum gas limit
     */
    function XMSG_MIN_GAS_LIMIT() external view returns (uint64);

    /**
     * @notice Numerator of the fraction of total validator power required to
     *         accept an XSubmission. Ex 2/3 -> 2
     * @return Quorum threshold numerator
     */
    function XSUB_QUORUM_NUMERATOR() external view returns (uint8);

    /**
     * @notice Denominator of the fraction of total validator power required to
     *         accept an XSubmission. Ex 2/3 -> 3
     * @return Quorum threshold denominator
     */
    function XSUB_QUORUM_DENOMINATOR() external view returns (uint8);

    /**
     * @notice Chain ID of the chain to which this portal is deployed
     * @dev Used as sourceChainId for all outbound XMsgs
     * @return Chain ID
     */
    function chainId() external view returns (uint64);

    /**
     * @notice Offset of the last outbound XMsg that was sent to destChainId
     * @param destChainId Destination chain ID
     * @return Offset
     */
    function outXStreamOffset(uint64 destChainId) external view returns (uint64);

    /**
     * @notice Offset of the last inbound XMsg that was received from sourceChainId
     * @param sourceChainId Source chain ID
     * @return Offset
     */
    function inXStreamOffset(uint64 sourceChainId) external view returns (uint64);

    /**
     * @notice Source block height of the last inbound XMsg that was received from sourceChainId
     * @param sourceChainId Source chain ID
     * @return Block height
     */
    function inXStreamBlockHeight(uint64 sourceChainId) external view returns (uint64);

    /**
     * @notice The current XMsg being executed via this portal
     * @dev If no XMsg is being executed, all fields will be zero
     * @return XMsg
     */
    function xmsg() external view returns (XTypes.Msg memory);

    /**
     * @notice Whether the current transaction is an xcall
     * @return True if current transaction is an xcall, false otherwise
     */
    function isXCall() external view returns (bool);

    /**
     * @notice Calculate the fee for calling a contract on another chain
     * @dev Uses OmniPortal.XMSG_DEFAULT_GAS_LIMIT
     * @dev Fees denominated in wei
     * @param destChainId Destination chain ID
     * @param data Encoded function calldata
     */
    function feeFor(uint64 destChainId, bytes calldata data) external view returns (uint256);

    /**
     * @notice Calculate the fee for calling a contract on another chain
     * @dev Fees denominated in wei
     * @param destChainId Destination chain ID
     * @param data Encoded function calldata
     * @param gasLimit Execution gas limit, enforced on destination chain
     */
    function feeFor(uint64 destChainId, bytes calldata data, uint64 gasLimit) external view returns (uint256);

    /**
     * @notice Call a contract on another chain
     * @dev Uses OmniPortal.XMSG_DEFAULT_GAS_LIMIT as execution gas limit on destination chain
     * @param destChainId Destination chain ID
     * @param to Address of contract to call on destination chain
     * @param data Encoded function calldata (use abi.encodeWithSignature
     * 	or abi.encodeWithSelector)
     */
    function xcall(uint64 destChainId, address to, bytes calldata data) external payable;

    /**
     * @notice Call a contract on another chain
     * @dev Uses provide gasLimit as execution gas limit on destination chain.
     *      Reverts if gasLimit < XMSG_MAX_GAS_LIMIT or gasLimit > XMSG_MAX_GAS_LIMIT
     * @param destChainId Destination chain ID
     * @param to Address of contract to call on destination chain
     * @param data Encoded function calldata (use abi.encodeWithSignature
     * 	or abi.encodeWithSelector)
     */
    function xcall(uint64 destChainId, address to, bytes calldata data, uint64 gasLimit) external payable;

    /**
     * @notice Submit a batch of XMsgs to be executed on this chain
     * @param xsub An xchain submisison, including an attestation root w/ validator signatures,
     *        and a block header and message batch, proven against the attestation root.
     */
    function xsubmit(XTypes.Submission calldata xsub) external;
}
