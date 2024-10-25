// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import { IFeeOracle } from "./IFeeOracle.sol";
import { IConversionRateOracle } from "./IConversionRateOracle.sol";

/**
 * @title IFeeOracleV2
 * @notice Extends IFeeOracle with FeeOracleV2 methods
 */
interface IFeeOracleV2 is IFeeOracle, IConversionRateOracle {
    /// @notice Emitted when fee parameters for a chain are set.
    event FeeParamsSet(uint64 chainId, uint256 execGasPrice, uint256 dataGasPrice, uint256 toNativeRate);

    /// @notice Emitted when the base gas limit is set.
    event BaseGasLimitSet(uint64 baseGasLimit);

    /// @notice Emitted when the base protocol fee is set.
    event ProtocolFeeSet(uint256 protocolFee);

    /// @notice Emitted when the gas price for a destination chain is set.
    event ExecGasPriceSet(uint64 chainId, uint256 gasPrice);

    /// @notice Emitted when the data gas price for a destination chain is set.
    event DataGasPriceSet(uint64 chainId, uint256 gasPrice);

    /// @notice Emitted when the to-native conversion rate for a destination chain is set.
    event ToNativeRateSet(uint64 chainId, uint256 toNativeRate);

    /// @notice Emitted when the manager is changed.
    event ManagerSet(address manager);

    /**
     * @notice Fee parameters for a specific chain.
     * @custom:field chainId        The chain ID.
     * @custom:field execGasPrice   The execution gas price on that chain (denominated in chains native token).
     * @custom:field dataGasPrice   The data gas price on that chain (denominated in chains native token).
     *                              ex. for Optimism, dataGasPrice is Ethereum L1's blob gas price.
     * @custom:field toNativeRate   The conversion rate from the chains native token to this chain's
     *                              native token. Rate is numerator over CONVERSION_RATE_DENOM.
     */
    struct FeeParams {
        uint64 chainId;
        uint64 execGasPrice;
        uint64 dataGasPrice;
        uint64 toNativeRate;
    }

    /// @notice Returns the fee parameters for a destination chain.
    function feeParams(uint64 chainId) external view returns (FeeParams memory);

    /// @notice Returns the execution gas price for a destination chain.
    function execGasPrice(uint64 chainId) external view returns (uint64);

    /// @notice Returns the data gas price for a destination chain.
    function dataGasPrice(uint64 chainId) external view returns (uint64);

    /// @notice Returns the to-native conversion rate for a destination chain.
    function toNativeRate(uint64 chainId) external view returns (uint256);

    /// @notice Returns the manager's address.
    function manager() external view returns (address);

    /// @notice Returns the protocol fee.
    function protocolFee() external view returns (uint64);

    /// @notice Returns the base gas limit.
    function baseGasLimit() external view returns (uint32);

    /// @notice Set the fee parameters for a list of destination chains.
    function bulkSetFeeParams(FeeParams[] calldata params) external;

    /// @notice Set the execution gas price for a destination chain.
    function setExecGasPrice(uint64 chainId, uint64 execGasPrice) external;

    /// @notice Set the data gas price for a destination chain.
    function setDataGasPrice(uint64 chainId, uint64 dataGasPrice) external;

    /// @notice Set the to native conversion rate for a destination chain.
    function setToNativeRate(uint64 chainId, uint64 toNativeRate) external;

    /// @notice Set the base gas limit for each xmsg.
    function setBaseGasLimit(uint32 gasLimit) external;

    /// @notice Set the base protocol fee for each xmsg.
    function setProtocolFee(uint64 fee) external;

    /// @notice Set the manager admin account.
    function setManager(address manager) external;
}
