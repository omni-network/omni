// SPDX-License-Identifier: No License (None)
pragma solidity ^0.8.23;

/**
 * @title IOmniPortal
 * @notice Portal in and out or Omni's xchain messaging protocal
 */
interface IOmniPortal {
    /**
     * @notice Emitted when a contract is called on another chain
     * @param destChainId Destination chain ID
     * @param streamOffset Offset of corresponding XMsg in source -> dest XStream
     * @param sender msg.sender of the source xcall
     * @param to Address of the contract to call on the destination chain
     * @param gasLimit Gas limit for execution on destination chain
     * @param data Encoded function calldata
     */
    event XMsg(
        uint64 indexed destChainId, uint64 indexed streamOffset, address sender, address to, bytes data, uint64 gasLimit
    );

    /**
     * @notice Default xmsg execution gas limit, enforced on destination chain
     * @return Gas limit
     */
    function XMSG_DEFAULT_GAS_LIMIT() external view returns (uint64);

    /**
     * @notice Chain ID of the chain to which this portal is deployed
     * @dev Used as sourceChainId for all outbound XMsgs
     * @return Chain ID
     */
    function chainId() external view returns (uint64);

    /**
     * @notice Offset of the next XMsg to be sent to the given chain
     * @param destChainId Destination chain ID
     * @return Offset
     */
    function outXStreamOffset(uint64 destChainId) external view returns (uint64);

    /**
     * @notice Call a contract on another chain
     * @dev Uses OmniPortal.XMSG_DEFAULT_GAS_LIMIT as execution gas limit
     * @param destChainId Destination chain ID
     * @param to Address of contract to call on destination chain
     * @param data Encoded function calldata (use abi.encodeWithSignature
     * 	or abi.encodeWithSelector)
     */
    function xcall(uint64 destChainId, address to, bytes calldata data) external payable;
}
