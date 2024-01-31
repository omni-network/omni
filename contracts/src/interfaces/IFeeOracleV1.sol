// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

import { IFeeOracle } from "./IFeeOracle.sol";

/**
 * @title IFeeOracleV1
 * @notice Extends IFeeOracle with FeeOracleV1 methods
 */
interface IFeeOracleV1 is IFeeOracle {
    /**
     * @notice Emitted when the admin account is changed
     * @param oldAdmin The old admin account
     * @param newAdmin The new admin account
     */
    event AdminChanged(address indexed oldAdmin, address indexed newAdmin);

    /**
     * @notice Emitted when the fee is changed
     * @param oldFee The old fee
     * @param newFee The new fee
     */
    event FeeChanged(uint256 oldFee, uint256 newFee);

    /**
     * @notice The admin account, who can change the fee.
     * @return The admin.
     */
    function admin() external view returns (address);

    /**
     * @notice Set the admin account.
     * @dev Only callable by the current admin.
     */
    function setAdmin(address _admin) external;

    /**
     * @notice The current fee per transaction, in wei.
     * @return The fee.
     */
    function fee() external view returns (uint256);

    /**
     * @notice Set the fee per transaction, in wei.
     * @param _fee The fee to set.
     */
    function setFee(uint256 _fee) external;
}
