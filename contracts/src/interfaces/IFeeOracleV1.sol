// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import { IFeeOracle } from "./IFeeOracle.sol";

/**
 * @title IFeeOracleV1
 * @notice Extends IFeeOracle with FeeOracleV1 methods
 */
interface IFeeOracleV1 is IFeeOracle {
    /**
     * @notice Emitted when fee parameters for a chain are set.
     */
    event FeeParamsSet(uint64 chainId, uint64 postsTo, uint256 gasPrice, uint256 toNativeRate);

    /**
     * @notice Emitted when the base gas limit is set.
     */
    event BaseGasLimitSet(uint64 baseGasLimit);

    /**
     * @notice Emitted when the base protocol fee is set.
     */
    event ProtocolFeeSet(uint256 protocolFee);

    /**
     * @notice Emitted when the gas price for a destination chain is set.
     */
    event GasPriceSet(uint64 chainId, uint256 gasPrice);

    /**
     * @notice Emitted when the to-native conversion rate for a destination chain is set.
     */
    event ToNativeRateSet(uint64 chainId, uint256 toNativeRate);

    /**
     * @notice Emitted when the manager is changed.
     */
    event ManagerSet(address manager);

    /**
     * @notice Fee parameters for a specific chain.
     * @custom:field chainId        The chain ID.
     * @custom:field postTo         The chain ID to which this chain posts tx calldata, used to calculate
     *                              calldata fees. For non-rollups, this should be the same as chainId.
     * @custom:field gasPrice       The gas price on that chain (denominated in chains native token).
     * @custom:field toNativeRate   The conversion rate from the chains native token to this chain's
     *                              native token. Rate is numerator over CONVERSION_RATE_DENOM.
     */
    struct ChainFeeParams {
        uint64 chainId;
        uint64 postsTo;
        uint256 gasPrice;
        uint256 toNativeRate;
    }

    /**
     * @notice Set the fee parameters for a list of destination chains.
     */
    function bulkSetFeeParams(ChainFeeParams[] calldata params) external;

    /**
     * @notice Set the gas price for a destination chain.
     */
    function setGasPrice(uint64 chainId, uint256 gasPrice) external;

    /**
     * @notice Set the to native conversion rate for a destination chain.
     */
    function setToNativeRate(uint64 chainId, uint256 toNativeRate) external;

    /**
     * @notice Set the base gas limit for each xmsg.
     */
    function setBaseGasLimit(uint64 gasLimit) external;

    /**
     * @notice Set the base protocol fee for each xmsg.
     */
    function setProtocolFee(uint256 fee) external;

    /**
     * @notice Set the manager admin account.
     */
    function setManager(address manager_) external;

    /**
     * @notice returns the conversion rate denominator, used in to
     */
    function CONVERSION_RATE_DENOM() external view returns (uint256);
}
