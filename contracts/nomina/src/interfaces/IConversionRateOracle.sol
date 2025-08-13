// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.30;

/**
 * @title IConversionRateOracle
 * @notice Provides conversion rate from other chain's native token to this chain's native token.
 */
interface IConversionRateOracle {
    /**
     * @notice Returns the conversion rate (as a numerator over CONVERSION_RATE_DENOM) from `chainId`'s
     *         native token to this chain's native token.
     */
    function toNativeRate(uint64 chainId) external view returns (uint256);

    /**
     * @notice Denominator used in to conversion rate calculations.
     */
    function CONVERSION_RATE_DENOM() external view returns (uint256);
}
