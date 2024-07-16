// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

/**
 * @title IOmniPortalAdmin
 * @notice Defines the OmniPortal admin interface
 */
interface IOmniPortalAdmin {
    /**
     * @notice Emitted when the fee oracle is updated.
     * @param oracle  The new fee oracle address
     */
    event FeeOracleSet(address oracle);

    /**
     * @notice Emited when fees are collected
     * @param to        The address the fees are collected to
     * @param amount    The amount of fees collected
     */
    event FeesCollected(address indexed to, uint256 amount);

    /**
     * @notice Emitted when xmsgMinGasLimit is updated
     * @param gasLimit The new xmsgMinGasLimit
     */
    event XMsgMinGasLimitSet(uint64 gasLimit);

    /**
     * @notice Emitted when xmsgMaxGasLimit is updated
     * @param gasLimit The new xmsgMaxGasLimit
     */
    event XMsgMaxGasLimitSet(uint64 gasLimit);

    /**
     * @notice Emitted when xmsgMaxDataSize is updated
     * @param size The new max size
     */
    event XMsgMaxDataSizeSet(uint16 size);

    /**
     * @notice Emitted when xreceiptMaxErrorSize is updated
     * @param size The new max size
     */
    event XReceiptMaxErrorSizeSet(uint16 size);

    /**
     * @notice Emitted when the xsubValsetCutoff is updated
     * @param cutoff The new cutoff
     */
    event XSubValsetCutoffSet(uint8 cutoff);

    /**
     * @notice Emitted the portal is paused, all xcalls and xsubmissions
     */
    event Paused();

    /**
     * @notice Emitted the portal is unpaused, all xcalls and xsubmissions
     */
    event Unpaused();

    /**
     * @notice Emitted when all xcalls are paused
     */
    event XCallPaused();

    /**
     * @notice Emitted when inbound xmsg offset is updated
     */
    event InXMsgOffsetSet(uint64 indexed srcChainId, uint64 indexed shardId, uint64 offset);

    /**
     * @notice Emitted when all inbound xblock offset is updated
     */
    event InXBlockOffsetSet(uint64 indexed srcChainId, uint64 indexed shardId, uint64 offset);

    /**
     * @notice Emitted when all xcalls are unpaused
     */
    event XCallUnpaused();

    /**
     * @notice Emitted when all xsubmissions are paused
     */
    event XSubmitPaused();

    /**
     * @notice Emitted when all xsubmissions are unpaused
     */
    event XSubmitUnpaused();

    /**
     * @notice Emitted when xcalls to a specific chain are paused
     * @param chainId   The destination chain
     */
    event XCallToPaused(uint64 indexed chainId);

    /**
     * @notice Emitted when xcalls to a specific chain are unpaused
     * @param chainId   The destination chain
     */
    event XCallToUnpaused(uint64 indexed chainId);

    /**
     * @notice Emitted when xsubmissions from a specific chain are paused
     * @param chainId    The source chain
     */
    event XSubmitFromPaused(uint64 indexed chainId);

    /**
     * @notice Emitted when xsubmissions from a specific chain are unpaused
     * @param chainId    The source chain
     */
    event XSubmitFromUnpaused(uint64 indexed chainId);

    /**
     * @notice Set the inbound xmsg offset for a chain and shard
     * @param sourceChainId    Source chain ID
     * @param shardId          Shard ID
     * @param offset           New xmsg offset
     */
    function setInXMsgOffset(uint64 sourceChainId, uint64 shardId, uint64 offset) external;

    /**
     * @notice Set the inbound xblock offset for a chain and shard
     * @param sourceChainId    Source chain ID
     * @param shardId          Shard ID
     * @param offset           New xblock offset
     */
    function setInXBlockOffset(uint64 sourceChainId, uint64 shardId, uint64 offset) external;

    /**
     * @notice Set the fee oracle
     */
    function setFeeOracle(address feeOracle) external;

    /**
     * @notice Transfer all collected fees to the give address
     * @param to    The address to transfer the fees to
     */
    function collectFees(address to) external;

    /**
     * @notice Set the minimum gas limit for xmsg
     */
    function setXMsgMinGasLimit(uint64 gasLimit) external;

    /**
     * @notice Set the maximum gas limit for xmsg
     */
    function setXMsgMaxGasLimit(uint64 gasLimit) external;

    /**
     * @notice Set the maximum data bytes for xmsg
     */
    function setXMsgMaxDataSize(uint16 numBytes) external;

    /**
     * @notice Set the maximum error bytes for xreceipt
     */
    function setXReceiptMaxErrorSize(uint16 numBytes) external;

    /**
     * @notice Pause xcalls
     */
    function pause() external;

    /**
     * @notice Unpause xcalls
     */
    function unpause() external;
}
