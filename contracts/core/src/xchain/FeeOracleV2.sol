// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { IFeeOracle } from "../interfaces/IFeeOracle.sol";
import { IFeeOracleV2 } from "../interfaces/IFeeOracleV2.sol";

/**
 * @title FeeOracleV2
 * @notice A simple fee oracle with a fixed fee, controlled by an admin account
 *         Used by OmniPortal to calculate xmsg fees
 */
contract FeeOracleV2 is IFeeOracle, IFeeOracleV2, OwnableUpgradeable {
    /**
     * @notice Base protocol fee for each xmsg.
     */
    uint128 public protocolFee;

    /**
     * @notice Address allowed to set gas prices and to-native conversion rates.
     */
    address public manager;

    /**
     * @notice Conversion rate from this chain's gas token to another native token, by gas token ID.
     */
    mapping(uint8 gasToken => uint256) public tokenToNativeRate;

    /**
     * @notice Fee parameters for a specific chain, by chain ID.
     */
    mapping(uint64 chainId => IFeeOracleV2.FeeParams) internal _feeParams;

    /**
     * @notice Data cost parameters for a specific chain, by data cost ID.
     */
    mapping(uint64 dataCostId => IFeeOracleV2.DataCostParams) internal _dataCostParams;

    /**
     * @notice Denominator for conversion rate calculations.
     */
    uint256 public constant CONVERSION_RATE_DENOM = 1e6;

    modifier onlyManager() {
        if (msg.sender != manager) revert IFeeOracleV2.NotManager();
        _;
    }

    constructor() {
        _disableInitializers();
    }

    function initialize(
        address owner_,
        address manager_,
        uint128 protocolFee_,
        FeeParams[] calldata feeParams_,
        DataCostParams[] calldata dataCostParams_,
        NativeRateParams[] calldata nativeRateParams_
    ) public initializer {
        __Ownable_init(owner_);

        _setManager(manager_);
        _setProtocolFee(protocolFee_);
        _bulkSetFeeParams(feeParams_);
        _bulkSetDataCostParams(dataCostParams_);
        _bulkSetToNativeRate(nativeRateParams_);
    }

    /// @inheritdoc IFeeOracle
    function version() external pure override returns (uint64) {
        return 2;
    }

    /// @inheritdoc IFeeOracle
    function feeFor(uint64 destChainId, bytes calldata data, uint64 gasLimit) external view returns (uint256) {
        IFeeOracleV2.FeeParams storage p = _feeParams[destChainId];
        IFeeOracleV2.DataCostParams storage d = _dataCostParams[p.dataCostId];

        uint256 _execGasPrice = p.gasPrice * tokenToNativeRate[p.gasToken] / CONVERSION_RATE_DENOM;
        uint256 _dataGasPrice = d.gasPrice * tokenToNativeRate[d.gasToken] / CONVERSION_RATE_DENOM;

        if (_execGasPrice == 0) revert IFeeOracleV2.NoFeeParams();
        if (_dataGasPrice == 0) revert IFeeOracleV2.NoFeeParams();

        // 16 gas per non-zero byte, assume non-zero bytes
        uint256 dataGas = (d.baseDataBuffer + data.length) * d.gasPerByte;

        return protocolFee + (p.baseGasLimit + gasLimit) * _execGasPrice + (dataGas * _dataGasPrice);
    }

    /**
     * @notice Returns the fee parameters for a destination chain.
     */
    function feeParams(uint64 chainId) external view returns (FeeParams memory) {
        return _feeParams[chainId];
    }

    /**
     * @notice Returns the data cost parameters for a data cost ID.
     */
    function dataCostParams(uint64 dataCostId) external view returns (DataCostParams memory) {
        return _dataCostParams[dataCostId];
    }

    /**
     * @notice Returns the gas price for a destination chain.
     */
    function execGasPrice(uint64 chainId) external view returns (uint64) {
        return _feeParams[chainId].gasPrice;
    }

    /**
     * @notice Returns the data gas price for a data cost ID.
     */
    function dataGasPrice(uint64 dataCostId) external view returns (uint64) {
        return _dataCostParams[dataCostId].gasPrice;
    }

    /**
     * @notice Returns the base gas limit for a destination chain.
     */
    function baseGasLimit(uint64 chainId) external view returns (uint32) {
        return _feeParams[chainId].baseGasLimit;
    }

    /**
     * @notice Returns the base data buffer for a data cost ID.
     */
    function baseDataBuffer(uint64 dataCostId) external view returns (uint32) {
        return _dataCostParams[dataCostId].baseDataBuffer;
    }

    /**
     * @notice Returns the gas token for a destination chain.
     */
    function execGasToken(uint64 chainId) external view returns (uint8) {
        return _feeParams[chainId].gasToken;
    }

    /**
     * @notice Returns the gas token for a data cost ID.
     */
    function dataGasToken(uint64 dataCostId) external view returns (uint8) {
        return _dataCostParams[dataCostId].gasToken;
    }

    /**
     * @notice Returns the data cost ID for a destination chain.
     */
    function execDataCostId(uint64 chainId) external view returns (uint64) {
        return _feeParams[chainId].dataCostId;
    }

    /**
     * @notice Returns the gas per byte for a data cost ID.
     */
    function dataGasPerByte(uint64 dataCostId) external view returns (uint64) {
        return _dataCostParams[dataCostId].gasPerByte;
    }

    /**
     * @notice Returns the to-native conversion rate for a destination chain.
     */
    function toNativeRate(uint64 chainId) external view returns (uint256) {
        return tokenToNativeRate[_feeParams[chainId].gasToken];
    }

    /**
     * @notice Returns the to-native conversion rate for a data cost ID.
     */
    function toNativeRateData(uint64 dataCostId) external view returns (uint256) {
        return tokenToNativeRate[_dataCostParams[dataCostId].gasToken];
    }

    /**
     * @notice Set the fee parameters for a list of destination chains.
     */
    function bulkSetFeeParams(FeeParams[] calldata params) external onlyManager {
        _bulkSetFeeParams(params);
    }

    /**
     * @notice Set the data cost parameters for a list of data cost IDs.
     */
    function bulkSetDataCostParams(DataCostParams[] calldata params) external onlyManager {
        _bulkSetDataCostParams(params);
    }

    /**
     * @notice Set the to-native conversion rate for a list of gas tokens.
     */
    function bulkSetToNativeRate(NativeRateParams[] calldata params) external onlyManager {
        _bulkSetToNativeRate(params);
    }

    /**
     * @notice Set the execution gas price for a destination chain.
     */
    function setExecGasPrice(uint64 chainId, uint64 gasPrice) external onlyManager {
        _setExecGasPrice(chainId, gasPrice);
    }

    /**
     * @notice Set the data gas price for a data cost ID.
     */
    function setDataGasPrice(uint64 dataCostId, uint64 gasPrice) external onlyManager {
        _setDataGasPrice(dataCostId, gasPrice);
    }

    /**
     * @notice Set the base gas limit for a destination chain.
     */
    function setBaseGasLimit(uint64 chainId, uint32 newBaseGasLimit) external onlyManager {
        _setBaseGasLimit(chainId, newBaseGasLimit);
    }

    /**
     * @notice Set the base data buffer for a data cost ID.
     */
    function setBaseDataBuffer(uint64 dataCostId, uint32 newBaseDataBuffer) external onlyManager {
        _setBaseDataBuffer(dataCostId, newBaseDataBuffer);
    }

    /**
     * @notice Set the data cost ID for a destination chain.
     */
    function setDataCostId(uint64 chainId, uint64 dataCostId) external onlyManager {
        _setDataCostId(chainId, dataCostId);
    }

    /**
     * @notice Set the gas per byte for a data cost ID.
     */
    function setGasPerByte(uint64 dataCostId, uint64 gasPerByte) external onlyManager {
        _setGasPerByte(dataCostId, gasPerByte);
    }

    /**
     * @notice Set the to native conversion rate for a gas token.
     */
    function setToNativeRate(uint8 gasToken, uint256 nativeRate) external onlyManager {
        _setToNativeRate(gasToken, nativeRate);
    }

    /**
     * @notice Set the base protocol fee for each xmsg.
     */
    function setProtocolFee(uint128 fee) external onlyOwner {
        _setProtocolFee(fee);
    }

    /**
     * @notice Set the manager admin account.
     */
    function setManager(address manager_) external onlyOwner {
        if (manager_ == address(0)) revert IFeeOracleV2.ZeroAddress();
        _setManager(manager_);
    }

    /**
     * @notice Set the fee parameters for a list of destination chains.
     */
    function _bulkSetFeeParams(FeeParams[] calldata params) internal {
        for (uint256 i = 0; i < params.length; i++) {
            FeeParams memory p = params[i];

            if (p.gasToken == 0) revert IFeeOracleV2.ZeroGasToken();
            if (p.chainId == 0) revert IFeeOracleV2.ZeroChainId();
            if (p.gasPrice == 0) revert IFeeOracleV2.ZeroGasPrice();
            if (p.dataCostId == 0) revert IFeeOracleV2.ZeroDataCostId();

            _feeParams[p.chainId] = p;

            emit FeeParamsSet(p.gasToken, p.chainId, p.gasPrice, p.dataCostId);
        }
    }

    /**
     * @notice Set the data cost parameters for a list of data cost IDs.
     */
    function _bulkSetDataCostParams(DataCostParams[] calldata params) internal {
        for (uint256 i = 0; i < params.length; i++) {
            DataCostParams memory d = params[i];

            if (d.gasToken == 0) revert IFeeOracleV2.ZeroGasToken();
            if (d.dataCostId == 0) revert IFeeOracleV2.ZeroDataCostId();
            if (d.gasPrice == 0) revert IFeeOracleV2.ZeroGasPrice();
            if (d.gasPerByte == 0) revert IFeeOracleV2.ZeroGasPerByte();

            _dataCostParams[d.dataCostId] = d;

            emit DataCostParamsSet(d.gasToken, d.dataCostId, d.gasPrice, d.gasPerByte);
        }
    }

    /**
     * @notice Set the to-native conversion rate for a list of gas tokens.
     */
    function _bulkSetToNativeRate(NativeRateParams[] calldata params) internal {
        for (uint256 i = 0; i < params.length; i++) {
            NativeRateParams memory n = params[i];

            if (n.gasToken == 0) revert IFeeOracleV2.ZeroGasToken();
            if (n.nativeRate == 0) revert IFeeOracleV2.ZeroNativeRate();

            tokenToNativeRate[n.gasToken] = n.nativeRate;

            emit ToNativeRateSet(n.gasToken, n.nativeRate);
        }
    }

    /**
     * @notice Set the execution gas price for a destination chain.
     */
    function _setExecGasPrice(uint64 chainId, uint64 gasPrice) internal {
        if (gasPrice == 0) revert IFeeOracleV2.ZeroGasPrice();
        if (chainId == 0) revert IFeeOracleV2.ZeroChainId();

        _feeParams[chainId].gasPrice = gasPrice;
        emit ExecGasPriceSet(chainId, gasPrice);
    }

    /**
     * @notice Set the data gas price for a destination chain.
     */
    function _setDataGasPrice(uint64 dataCostId, uint64 gasPrice) internal {
        if (gasPrice == 0) revert IFeeOracleV2.ZeroGasPrice();
        if (dataCostId == 0) revert IFeeOracleV2.ZeroDataCostId();

        _dataCostParams[dataCostId].gasPrice = gasPrice;
        emit DataGasPriceSet(dataCostId, gasPrice);
    }

    /**
     * @notice Set the base gas limit for a destination chain.
     */
    function _setBaseGasLimit(uint64 chainId, uint32 newBaseGasLimit) internal {
        if (chainId == 0) revert IFeeOracleV2.ZeroChainId();

        _feeParams[chainId].baseGasLimit = newBaseGasLimit;
        emit BaseGasLimitSet(chainId, newBaseGasLimit);
    }

    /**
     * @notice Set the base data buffer for a data cost ID.
     */
    function _setBaseDataBuffer(uint64 dataCostId, uint32 newBaseDataBuffer) internal {
        if (dataCostId == 0) revert IFeeOracleV2.ZeroDataCostId();

        _dataCostParams[dataCostId].baseDataBuffer = newBaseDataBuffer;
        emit BaseDataBufferSet(dataCostId, newBaseDataBuffer);
    }

    /**
     * @notice Set the data cost ID for a destination chain.
     */
    function _setDataCostId(uint64 chainId, uint64 dataCostId) internal {
        if (chainId == 0) revert IFeeOracleV2.ZeroChainId();
        if (dataCostId == 0) revert IFeeOracleV2.ZeroDataCostId();

        _feeParams[chainId].dataCostId = dataCostId;
        emit DataCostIdSet(chainId, dataCostId);
    }

    /**
     * @notice Set the gas per byte for a data cost ID.
     */
    function _setGasPerByte(uint64 dataCostId, uint64 gasPerByte) internal {
        if (dataCostId == 0) revert IFeeOracleV2.ZeroDataCostId();
        if (gasPerByte == 0) revert IFeeOracleV2.ZeroGasPerByte();

        _dataCostParams[dataCostId].gasPerByte = gasPerByte;
        emit GasPerByteSet(dataCostId, gasPerByte);
    }

    /**
     * @notice Set the to-native conversion rate for a gas token.
     */
    function _setToNativeRate(uint8 gasToken, uint256 nativeRate) internal {
        if (nativeRate == 0) revert IFeeOracleV2.ZeroNativeRate();
        if (gasToken == 0) revert IFeeOracleV2.ZeroGasToken();

        tokenToNativeRate[gasToken] = nativeRate;
        emit ToNativeRateSet(gasToken, nativeRate);
    }

    /**
     * @notice Set the base protocol fee for each xmsg.
     */
    function _setProtocolFee(uint128 fee) internal {
        protocolFee = fee;
        emit ProtocolFeeSet(fee);
    }

    /**
     * @notice Set the manager admin account.
     */
    function _setManager(address manager_) internal {
        manager = manager_;
        emit ManagerSet(manager_);
    }
}
