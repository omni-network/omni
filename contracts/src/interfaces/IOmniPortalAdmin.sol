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
     * @notice Emitted when the xmsgDefaultGasLimit is changed.
     * @param oldDefault The old xmsgDefaultGasLimit
     * @param newDefault The new xmsgDefaultGasLimit
     */
    event XMsgDefaultGasLimitChanged(uint64 indexed oldDefault, uint64 indexed newDefault);

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
     * @notice Emitted when xreceiptMaxErrorBytes is changed.
     * @param oldMax The old xreceiptMaxErrorBytes
     * @param newMax The new xreceiptMaxErrorBytes
     */
    event XReceiptMaxErrorBytesChanged(uint64 indexed oldMax, uint64 indexed newMax);

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
     * @notice Set the default gas limit for xmsg
     */
    function setXMsgDefaultGasLimit(uint64 gasLimit) external;

    /**
     * @notice Set the minimum gas limit for xmsg
     */
    function setXMsgMinGasLimit(uint64 gasLimit) external;

    /**
     * @notice Set the maximum gas limit for xmsg
     */
    function setXMsgMaxGasLimit(uint64 gasLimit) external;

    /**
     * @notice Set the maximum error bytes for xreceipt
     */
    function setXReceiptMaxErrorBytes(uint64 maxErrorBytes) external;

    /**
     * @notice Pause xcalls
     */
    function pause() external;

    /**
     * @notice Unpause xcalls
     */
    function unpause() external;
}
