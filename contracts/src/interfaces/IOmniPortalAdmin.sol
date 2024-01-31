// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

/**
 * @title IOmniPortalAdmin
 * @notice Defines the OmniPortal admin interface
 */
interface IOmniPortalAdmin {
    /**
     * @notice Emitted when the admin account is changed
     * @param oldAdmin The old admin account
     * @param newAdmin The new admin account
     */
    event AdminChanged(address indexed oldAdmin, address indexed newAdmin);

    /**
     * @notice Emitted when the fee oracle is changed
     * @param oldFeeOracle The old fee oracle
     * @param newFeeOracle The new fee oracle
     */
    event FeeOracleChanged(address indexed oldFeeOracle, address indexed newFeeOracle);

    /**
     * @notice The current admin account.
     * @return The admin.
     */
    function admin() external view returns (address);

    /**
     * @notice Set the admin account.
     * @dev Only callable by the current admin.
     */
    function setAdmin(address admin) external;

    /**
     * @notice The current fee oracle.
     * @return The fee oracle.
     */
    function feeOracle() external view returns (address);

    /**
     * @notice Set the fee oracle.
     * @dev Only callable by the current admin.
     */
    function setFeeOracle(address feeOracle) external;
}
