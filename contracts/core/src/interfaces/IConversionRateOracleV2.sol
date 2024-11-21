// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

/**
 * @title IConversionRateOracle
 * @notice Provides conversion rate from other chain's native token to this chain's native token.
 */
interface IConversionRateOracleV2 {
    /**
     * @notice Returns the conversion rate (as a numerator over CONVERSION_RATE_DENOM) from `chainId`'s
     *         native token to this chain's native token.
     */
    function execToNativeRate(uint64 chainId) external view returns (uint64);

    /**
     * @notice Returns the conversion rate (as a numerator over CONVERSION_RATE_DENOM) from `chainId`'s
     *         native token to this chain's native token.
     */
    function dataToNativeRate(uint64 chainId) external view returns (uint64);

    /**
     * @notice Denominator used in to conversion rate calculations.
     */
    function CONVERSION_RATE_DENOM() external view returns (uint256);
}
