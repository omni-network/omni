// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

/**
 * @title IOmniPortalAdmin
 * @notice Defines the OmniPortal admin interface
 */
interface IOmniPortalAdmin {
    /**
     * @notice Emitted when the fee oracle is changed
     * @param oldOracle  The old fee oracle
     * @param newOrcale  The new fee oracle
     */
    event FeeOracleChanged(address indexed oldOracle, address indexed newOrcale);

    /**
     * @notice Emited when fees are collected
     * @param to        The address the fees are collected to
     * @param amount    The amount of fees collected
     */
    event FeesCollected(address indexed to, uint256 amount);

    /**
     * @notice Emitted when xmsgMinGasLimit is changed.
     * @param oldMin The old xmsgMinGasLimit
     * @param newMin The new xmsgMinGasLimit
     */
    event XMsgMinGasLimitChanged(uint64 indexed oldMin, uint64 indexed newMin);

    /**
     * @notice Emitted when xmsgMaxGasLimit is changed.
     * @param oldMax The old xmsgMaxGasLimit
     * @param newMax The new xmsgMaxGasLimit
     */
    event XMsgMaxGasLimitChanged(uint64 indexed oldMax, uint64 indexed newMax);

    /**
     * @notice Emitted when xmsgMaxDataSize is changed.
     * @param oldMax The old max size
     * @param newMax The new max size
     */
    event XMsgMaxDataSizeChanged(uint16 indexed oldMax, uint16 indexed newMax);

    /**
     * @notice Emitted when xreceiptMaxErrorSize is changed.
     * @param oldMax The old max size
     * @param newMax The new max size
     */
    event XReceiptMaxErrorSizeChanged(uint16 indexed oldMax, uint16 indexed newMax);

    /**
     * @notice Emitted the portal is paused, all xcalls and xsubmissions.
     */
    event Paused();

    /**
     * @notice Emitted the portal is unpaused, all xcalls and xsubmissions.
     */
    event Unpaused();

    /**
     * @notice Emitted when all xcalls are paused.
     */
    event XCallPaused();

    /**
     * @notice Emitted when all xcalls are unpaused.
     */
    event XCallUnpaused();

    /**
     * @notice Emitted when all xsubmissions are paused.
     */
    event XSubmitPaused();

    /**
     * @notice Emitted when all xsubmissions are unpaused.
     */
    event XSubmitUnpaused();

    /**
     * @notice Emitted when xcalls to a specific chain are paused.
     * @param chainId   The destination chain
     */
    event XCallToPaused(uint64 indexed chainId);

    /**
     * @notice Emitted when xcalls to a specific chain are unpaused.
     * @param chainId   The destination chain
     */
    event XCallToUnpaused(uint64 indexed chainId);

    /**
     * @notice Emitted when xsubmissions from a specific chain are paused.
     * @param chainId    The source chain
     */
    event XSubmitFromPaused(uint64 indexed chainId);

    /**
     * @notice Emitted when xsubmissions from a specific chain are unpaused.
     * @param chainId    The source chain
     */
    event XSubmitFromUnpaused(uint64 indexed chainId);

    /**
     * @notice Returns the current fee oracle address
     */
    function feeOracle() external view returns (address);

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
