// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import { IFeeOracle } from "./IFeeOracle.sol";
import { IConversionRateOracle } from "./IConversionRateOracle.sol";

/**
 * @title IFeeOracleV2
 * @notice Extends IFeeOracle with FeeOracleV2 methods
 */
interface IFeeOracleV2 is IFeeOracle, IConversionRateOracle {
    /// @notice Thrown when the caller is not the manager.
    error NotManager();

    /// @notice Thrown when there are no fee parameters for a chain or data cost ID.
    error NoFeeParams();

    /// @notice Thrown when the address is zero.
    error ZeroAddress();

    /// @notice Thrown when the chain ID is zero.
    error ZeroChainId();

    /// @notice Thrown when the gas price is zero.
    error ZeroGasPrice();

    /// @notice Thrown when the gas token is zero.
    error ZeroGasToken();

    /// @notice Thrown when the data cost ID is zero.
    error ZeroDataCostId();

    /// @notice Thrown when the gas per byte is zero.
    error ZeroGasPerByte();

    /// @notice Thrown when the native rate is zero.
    error ZeroNativeRate();

    /// @notice Emitted when fee parameters for a chain are set.
    event FeeParamsSet(uint8 gasToken, uint64 chainId, uint64 gasPrice, uint64 dataCostId);

    /// @notice Emitted when data cost parameters for a data cost ID are set.
    event DataCostParamsSet(uint8 gasToken, uint64 dataCostId, uint64 gasPrice, uint64 gasPerByte);

    /// @notice Emitted when the gas price for a destination chain is set.
    event ExecGasPriceSet(uint64 chainId, uint64 gasPrice);

    /// @notice Emitted when the data gas price for a data cost ID is set.
    event DataGasPriceSet(uint64 dataCostId, uint64 gasPrice);

    /// @notice Emitted when the base gas limit for a destination chain is set.
    event BaseGasLimitSet(uint64 chainId, uint32 baseGasLimit);

    /// @notice Emitted when the base data buffer for a data cost ID is set.
    event BaseDataBufferSet(uint64 dataCostId, uint32 baseDataBuffer);

    /// @notice Emitted when the data cost ID for a destination chain is set.
    event DataCostIdSet(uint64 chainId, uint64 dataCostId);

    /// @notice Emitted when the gas per byte for a data cost ID is set.
    event GasPerByteSet(uint64 dataCostId, uint64 gasPerByte);

    /// @notice Emitted when the to-native conversion rate for a gas token is set.
    event ToNativeRateSet(uint8 gasToken, uint256 nativeRate);

    /// @notice Emitted when the base protocol fee is set.
    event ProtocolFeeSet(uint128 protocolFee);

    /// @notice Emitted when the manager is changed.
    event ManagerSet(address manager);

    /**
     * @notice Fee parameters for a specific chain.
     * @custom:field gasToken       The gas token ID.
     * @custom:field baseGasLimit   The base gas limit for that chain.
     * @custom:field chainId        The chain ID.
     * @custom:field gasPrice       The execution gas price on that chain (denominated in chains native token).
     * @custom:field dataCostId     The data cost ID for that chain.
     */
    struct FeeParams {
        uint8 gasToken;
        uint32 baseGasLimit;
        uint64 chainId;
        uint64 gasPrice;
        uint64 dataCostId;
    }

    /**
     * @notice Data cost parameters for a data cost ID.
     * @custom:field gasToken       The gas token ID.
     * @custom:field baseDataBuffer The base data buffer for that data cost ID.
     * @custom:field dataCostId     The data cost ID.
     * @custom:field gasPrice       The data gas price for that data cost ID (denominated in chains native token).
     * @custom:field gasPerByte     The gas per byte for that data cost ID.
     */
    struct DataCostParams {
        uint8 gasToken;
        uint32 baseDataBuffer;
        uint64 dataCostId;
        uint64 gasPrice;
        uint64 gasPerByte;
    }

    /**
     * @notice Parameters for a gas token's to-native conversion rate.
     * @custom:field gasToken       The gas token ID.
     * @custom:field nativeRate     The to-native conversion rate for that gas token.
     */
    struct NativeRateParams {
        uint8 gasToken;
        uint256 nativeRate;
    }

    /// @notice Returns the protocol fee.
    function protocolFee() external view returns (uint128);

    /// @notice Returns the manager's address.
    function manager() external view returns (address);

    /// @notice Returns the conversion rate from this chain's gas token to another native token, by gas token ID.
    function tokenToNativeRate(uint8 gasToken) external view returns (uint256);

    /// @notice Returns the fee parameters for a destination chain.
    function feeParams(uint64 chainId) external view returns (FeeParams memory);

    /// @notice Returns the data cost parameters for a data cost ID.
    function dataCostParams(uint64 dataCostId) external view returns (DataCostParams memory);

    /// @notice Returns the execution gas price for a destination chain.
    function execGasPrice(uint64 chainId) external view returns (uint64);

    /// @notice Returns the data gas price for a data cost ID.
    function dataGasPrice(uint64 dataCostId) external view returns (uint64);

    /// @notice Returns the base gas limit for a destination chain.
    function baseGasLimit(uint64 chainId) external view returns (uint32);

    /// @notice Returns the base data buffer for a data cost ID.
    function baseDataBuffer(uint64 dataCostId) external view returns (uint32);

    /// @notice Returns the gas token for a destination chain.
    function execGasToken(uint64 chainId) external view returns (uint8);

    /// @notice Returns the gas token for a data cost ID.
    function dataGasToken(uint64 dataCostId) external view returns (uint8);

    /// @notice Returns the data cost ID for a destination chain.
    function execDataCostId(uint64 chainId) external view returns (uint64);

    /// @notice Returns the gas per byte for a data cost ID.
    function dataGasPerByte(uint64 dataCostId) external view returns (uint64);

    /// @notice Returns the to-native conversion rate for a destination chain.
    function toNativeRate(uint64 chainId) external view returns (uint256);

    /// @notice Returns the to-native conversion rate for a data cost ID.
    function toNativeRateData(uint64 dataCostId) external view returns (uint256);

    /// @notice Set the fee parameters for a list of destination chains.
    function bulkSetFeeParams(FeeParams[] calldata params) external;

    /// @notice Set the data cost parameters for a list of data cost IDs.
    function bulkSetDataCostParams(DataCostParams[] calldata params) external;

    /// @notice Set the to-native conversion rate for a list of gas tokens.
    function bulkSetToNativeRate(NativeRateParams[] calldata params) external;

    /// @notice Set the execution gas price for a destination chain.
    function setExecGasPrice(uint64 chainId, uint64 gasPrice) external;

    /// @notice Set the data gas price for a data cost ID.
    function setDataGasPrice(uint64 dataCostId, uint64 gasPrice) external;

    /// @notice Set the base gas limit for a destination chain.
    function setBaseGasLimit(uint64 chainId, uint32 newBaseGasLimit) external;

    /// @notice Set the base data buffer for a data cost ID.
    function setBaseDataBuffer(uint64 dataCostId, uint32 newBaseDataBuffer) external;

    /// @notice Set the data cost ID for a destination chain.
    function setDataCostId(uint64 chainId, uint64 dataCostId) external;

    /// @notice Set the gas per byte for a data cost ID.
    function setGasPerByte(uint64 dataCostId, uint64 gasPerByte) external;

    /// @notice Set the to native conversion rate for a gas token.
    function setToNativeRate(uint8 gasToken, uint256 nativeRate) external;

    /// @notice Set the base protocol fee for each xmsg.
    function setProtocolFee(uint128 fee) external;

    /// @notice Set the manager admin account.
    function setManager(address manager) external;
}
