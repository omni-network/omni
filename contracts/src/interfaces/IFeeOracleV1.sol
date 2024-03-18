// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import { IFeeOracle } from "./IFeeOracle.sol";

/**
 * @title IFeeOracleV1
 * @notice Extends IFeeOracle with FeeOracleV1 methods
 */
interface IFeeOracleV1 is IFeeOracle {
    /**
     * @notice Emitted when the fee is changed
     * @param oldFee The old fee
     * @param newFee The new fee
     */
    event FeeChanged(uint256 oldFee, uint256 newFee);

    /**
     * @notice Returns the current fee per transaction, in wei.
     */
    function fee() external view returns (uint256);

    /**
     * @notice Set the fee per transaction, in wei.
     * @param _fee The fee to set.
     */
    function setFee(uint256 _fee) external;
}
