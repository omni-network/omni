// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

/**
 * @title IOmniPortalAdmin
 * @notice Defines the OmniPortal admin interface
 */
interface IOmniPortalAdmin {
    /**
     * @notice Emitted when the fee oracle is changed
     * @param oldFeeOracle  The old fee oracle
     * @param newFeeOracle  The new fee oracle
     */
    event FeeOracleChanged(address indexed oldFeeOracle, address indexed newFeeOracle);

    /**
     * @notice Emited when fees are collected
     * @param to        The address the fees are collected to
     * @param amount    The amount of fees collected
     */
    event FeesCollected(address indexed to, uint256 amount);

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
}
