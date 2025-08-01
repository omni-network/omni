// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.30;

/**
 * @title IFeeOracle
 * @notice Defines the interface expected of a fee oracle by the NominaPortal
 */
interface IFeeOracle {
    /**
     * @notice Calculate the fee for calling a contract on another chain
     * @dev Fees denominated in wei
     * @param destChainId Destination chain ID
     * @param data Encoded function calldata
     * @param gasLimit Execution gas limit, enforced on destination chain
     */
    function feeFor(uint64 destChainId, bytes calldata data, uint64 gasLimit) external view returns (uint256);

    /**
     * @notice Returns the version of the fee oracle
     */
    function version() external view returns (uint64);
}
