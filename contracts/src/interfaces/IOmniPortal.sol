// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

import { XChain } from "../libraries/XChain.sol";

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
     * @notice Emitted when an XMsg is executed on another chain
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
     * @notice Chain ID of the chain to which this portal is deployed
     * @dev Used as sourceChainId for all outbound XMsgs
     * @return Chain ID
     */
    function chainId() external view returns (uint64);

    /**
     * @notice Offset of the next outbound XMsg to be sent in the corresponding source -> dest XStream
     * @param destChainId Destination chain ID
     * @return Offset
     */
    function outXStreamOffset(uint64 destChainId) external view returns (uint64);

    /**
     * @notice Offset of the next inbound XMsg to be received in the corresponding source -> dest XStream
     * @param sourceChainId Destination chain ID
     * @return Offset
     */
    function inXStreamOffset(uint64 sourceChainId) external view returns (uint64);

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
    function xsubmit(XChain.Submission calldata xsub) external;
}
