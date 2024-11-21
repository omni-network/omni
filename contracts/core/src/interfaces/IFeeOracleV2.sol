// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import { IFeeOracle } from "./IFeeOracle.sol";
import { IConversionRateOracleV2 } from "./IConversionRateOracleV2.sol";

/**
 * @title IFeeOracleV2
 * @notice Extends IFeeOracle with FeeOracleV2 methods
 */
interface IFeeOracleV2 is IFeeOracle, IConversionRateOracleV2 {
    /// @notice Emitted when fee parameters for a chain are set.
    event ExecFeeParamsSet(uint64 chainId, uint64 postsTo, uint64 execGasPrice, uint64 toNativeRate);

    /// @notice Emitted when fee parameters for a chain are set.
    event DataFeeParamsSet(uint64 chainId, uint64 sizeBuffer, uint64 dataGasPrice, uint64 toNativeRate);

    /// @notice Emitted when the chain ID to which this chain posts tx calldata is set.
    event ExecPostsToSet(uint64 chainId, uint64 postsTo);

    /// @notice Emitted when the size buffer for a destination chain is set.
    event DataSizeBufferSet(uint64 chainId, uint64 sizeBuffer);

    /// @notice Emitted when the gas price for a destination chain is set.
    event ExecGasPriceSet(uint64 chainId, uint64 gasPrice);

    /// @notice Emitted when the data gas price for a destination chain is set.
    event DataGasPriceSet(uint64 chainId, uint64 gasPrice);

    /// @notice Emitted when the to-native conversion rate for an execution chain is set.
    event ExecToNativeRateSet(uint64 chainId, uint64 toNativeRate);

    /// @notice Emitted when the to-native conversion rate for a data inclusion chain is set.
    event DataToNativeRateSet(uint64 chainId, uint64 toNativeRate);

    /// @notice Emitted when the base protocol fee is set.
    event ProtocolFeeSet(uint72 protocolFee);

    /// @notice Emitted when the base gas limit is set.
    event BaseGasLimitSet(uint24 baseGasLimit);

    /// @notice Emitted when the manager is changed.
    event ManagerSet(address manager);

    /**
     * @notice Execution fee parameters for a specific chain.
     * @custom:field chainId        The chain ID.
     * @custom:field postsTo        The chain ID to which this chain posts tx calldata.
     * @custom:field execGasPrice   The execution gas price on that chain (denominated in chains native token).
     * @custom:field toNativeRate   The conversion rate from the chains native token to this chain's
     *                              native token. Rate is numerator over CONVERSION_RATE_DENOM.
     */
    struct ExecFeeParams {
        uint64 chainId;
        uint64 postsTo;
        uint64 execGasPrice;
        uint64 toNativeRate;
    }

    /**
     * @notice Data inclusion fee parameters for a specific chain.
     * @custom:field chainId        The chain ID.
     * @custom:field sizeBuffer     The size buffer for data inclusion on that chain.
     * @custom:field dataGasPrice   The data gas price on that chain (denominated in chains native token).
     *                              ex. for Optimism, dataGasPrice is Ethereum L1's blob gas price.
     * @custom:field toNativeRate   The conversion rate from the chains native token to this chain's
     *                              native token. Rate is numerator over CONVERSION_RATE_DENOM.
     */
    struct DataFeeParams {
        uint64 chainId;
        uint64 sizeBuffer;
        uint64 dataGasPrice;
        uint64 toNativeRate;
    }

    /**
     * @notice Fee parameters for a specific chain.
     * @custom:field execChainId       The execution chain ID.
     * @custom:field execPostsTo       The chain ID to which the execution chain posts tx calldata.
     * @custom:field execGasPrice      The execution gas price on the execution chain (denominated in chains native token).
     * @custom:field execToNativeRate  The conversion rate from the execution chain's native token to this chain's
     *                                 native token. Rate is numerator over CONVERSION_RATE_DENOM.
     * @custom:field dataChainId       The data inclusion chain ID.
     * @custom:field dataSizeBuffer    The size buffer for data inclusion on the data inclusion chain.
     * @custom:field dataGasPrice      The data gas price on the data inclusion chain (denominated in chains native token).
     *                                 ex. for Optimism, dataGasPrice is Ethereum L1's blob gas price.
     * @custom:field dataToNativeRate  The conversion rate from the data inclusion chain's native token to this chain's
     *                                 native token. Rate is numerator over CONVERSION_RATE_DENOM.
     */
    struct FeeParams {
        uint64 execChainId;
        uint64 execPostsTo;
        uint64 execGasPrice;
        uint64 execToNativeRate;
        uint64 dataChainId;
        uint64 dataSizeBuffer;
        uint64 dataGasPrice;
        uint64 dataToNativeRate;
    }

    /// @notice Returns the protocol fee.
    function protocolFee() external view returns (uint72);

    /// @notice Returns the base gas limit.
    function baseGasLimit() external view returns (uint24);

    /// @notice Returns the manager's address.
    function manager() external view returns (address);

    /// @notice Returns all of the fee parameters for a destination chain.
    function feeParams(uint64 chainId) external view returns (FeeParams memory);

    /// @notice Returns the chain ID to which an execution chain posts tx calldata.
    function execPostsTo(uint64 chainId) external view returns (uint64);

    /// @notice Returns the size buffer for a chain performing data inclusion.
    function dataSizeBuffer(uint64 chainId) external view returns (uint64);

    /// @notice Returns the execution gas price for a destination chain.
    function execGasPrice(uint64 chainId) external view returns (uint64);

    /// @notice Returns the data inclusion gas price for a destination chain.
    function dataGasPrice(uint64 chainId) external view returns (uint64);

    /// @notice Returns the to-native conversion rate for a chain performing execution.
    function execToNativeRate(uint64 chainId) external view returns (uint64);

    /// @notice Returns the to-native conversion rate for a chain performing data inclusion.
    function dataToNativeRate(uint64 chainId) external view returns (uint64);

    /// @notice Set the execution fee parameters for a list of destination chains.
    function bulkSetExecFeeParams(ExecFeeParams[] calldata params) external;

    /// @notice Set the data inclusion fee parameters for a list of chains.
    function bulkSetDataFeeParams(DataFeeParams[] calldata params) external;

    /// @notice Set the chain ID to which an execution chain posts tx calldata.
    function setExecPostsTo(uint64 chainId, uint64 postsTo) external;

    /// @notice Set the size buffer for a chain performing data inclusion.
    function setDataSizeBuffer(uint64 chainId, uint64 sizeBuffer) external;

    /// @notice Set the execution gas price for a destination chain.
    function setExecGasPrice(uint64 chainId, uint64 execGasPrice) external;

    /// @notice Set the data inclusion gas price for a chain.
    function setDataGasPrice(uint64 chainId, uint64 dataGasPrice) external;

    /// @notice Set the to-native conversion rate for an execution chain.
    function setExecToNativeRate(uint64 chainId, uint64 toNativeRate) external;

    /// @notice Set the to-native conversion rate for a data inclusion chain.
    function setDataToNativeRate(uint64 chainId, uint64 toNativeRate) external;

    /// @notice Set the base protocol fee for each xmsg.
    function setProtocolFee(uint72 fee) external;

    /// @notice Set the base gas limit for each xmsg.
    function setBaseGasLimit(uint24 gasLimit) external;

    /// @notice Set the manager admin account.
    function setManager(address manager) external;
}
