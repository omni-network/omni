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
    uint72 public protocolFee;

    /**
     * @notice Base gas limit for each xmsg.
     */
    uint24 public baseGasLimit;

    /**
     * @notice Address allowed to set gas prices and to-native conversion rates.
     */
    address public manager;

    /**
     * @notice Execution fee parameters for a specific chain, by chain ID.
     */
    mapping(uint64 => IFeeOracleV2.ExecFeeParams) internal _execFeeParams;

    /**
     * @notice Data inclusion fee parameters for a specific chain, by chain ID.
     */
    mapping(uint64 => IFeeOracleV2.DataFeeParams) internal _dataFeeParams;

    /**
     * @notice Denominator for conversion rate calculations.
     */
    uint256 public constant CONVERSION_RATE_DENOM = 1e6;

    modifier onlyManager() {
        require(msg.sender == manager, "FeeOracleV2: not manager");
        _;
    }

    constructor() {
        _disableInitializers();
    }

    function initialize(
        address owner_,
        address manager_,
        uint24 baseGasLimit_,
        uint72 protocolFee_,
        ExecFeeParams[] calldata execParams,
        DataFeeParams[] calldata dataParams
    ) public initializer {
        __Ownable_init(owner_);

        _setManager(manager_);
        _setBaseGasLimit(baseGasLimit_);
        _setProtocolFee(protocolFee_);
        _bulkSetExecFeeParams(execParams);
        _bulkSetDataFeeParams(dataParams);
    }

    /// @inheritdoc IFeeOracle
    function version() external pure override returns (uint64) {
        return 2;
    }

    /// @inheritdoc IFeeOracle
    function feeFor(uint64 destChainId, bytes calldata data, uint64 gasLimit) external view returns (uint256) {
        IFeeOracleV2.ExecFeeParams storage execChain = _execFeeParams[destChainId];
        IFeeOracleV2.DataFeeParams storage dataChain = _dataFeeParams[execChain.postsTo];

        uint256 _execGasPrice = execChain.execGasPrice * execChain.toNativeRate / CONVERSION_RATE_DENOM;
        uint256 _dataGasPrice = dataChain.dataGasPrice * dataChain.toNativeRate / CONVERSION_RATE_DENOM;

        require(_execGasPrice > 0, "FeeOracleV2: no exec fee params");
        require(_dataGasPrice > 0, "FeeOracleV2: no data fee params");

        // 16 gas per non-zero byte, assume non-zero bytes
        uint256 dataGas = data.length * 16;

        return
            protocolFee + (baseGasLimit + gasLimit) * _execGasPrice + ((dataChain.sizeBuffer + dataGas) * _dataGasPrice);
    }

    /**
     * @notice Returns all of the fee parameters for a destination chain.
     */
    function feeParams(uint64 chainId) external view returns (FeeParams memory) {
        IFeeOracleV2.ExecFeeParams memory execParams = _execFeeParams[chainId];
        IFeeOracleV2.DataFeeParams memory dataParams = _dataFeeParams[execParams.postsTo];
        IFeeOracleV2.FeeParams memory params = IFeeOracleV2.FeeParams({
            execChainId: chainId,
            execPostsTo: execParams.postsTo,
            execGasPrice: execParams.execGasPrice,
            execToNativeRate: execParams.toNativeRate,
            dataChainId: execParams.postsTo,
            dataSizeBuffer: dataParams.sizeBuffer,
            dataGasPrice: dataParams.dataGasPrice,
            dataToNativeRate: dataParams.toNativeRate
        });

        return params;
    }

    /**
     * @notice Returns the chain ID to which an execution chain posts tx calldata.
     */
    function execPostsTo(uint64 chainId) external view returns (uint64) {
        return _execFeeParams[chainId].postsTo;
    }

    /**
     * @notice Returns the size buffer for a chain performing data inclusion.
     */
    function dataSizeBuffer(uint64 chainId) external view returns (uint64) {
        return _dataFeeParams[chainId].sizeBuffer;
    }

    /**
     * @notice Returns the execution gas price for a destination chain.
     */
    function execGasPrice(uint64 chainId) external view returns (uint64) {
        return _execFeeParams[chainId].execGasPrice;
    }

    /**
     * @notice Returns the data inclusion gas price for a destination chain.
     */
    function dataGasPrice(uint64 chainId) external view returns (uint64) {
        return _dataFeeParams[chainId].dataGasPrice;
    }

    /**
     * @notice Returns the to-native conversion rate for a chain performing execution.
     */
    function execToNativeRate(uint64 chainId) external view returns (uint64) {
        return _execFeeParams[chainId].toNativeRate;
    }

    /**
     * @notice Returns the to-native conversion rate for a chain performing data inclusion.
     */
    function dataToNativeRate(uint64 chainId) external view returns (uint64) {
        return _dataFeeParams[chainId].toNativeRate;
    }

    /**
     * @notice Set the execution fee parameters for a list of destination chains.
     */
    function bulkSetExecFeeParams(ExecFeeParams[] calldata params) external onlyManager {
        _bulkSetExecFeeParams(params);
    }

    /**
     * @notice Set the data inclusion fee parameters for a list of chains.
     */
    function bulkSetDataFeeParams(DataFeeParams[] calldata params) external onlyManager {
        _bulkSetDataFeeParams(params);
    }

    /**
     * @notice Set the chain ID to which an execution chain posts tx calldata.
     */
    function setExecPostsTo(uint64 chainId, uint64 postsTo) external onlyManager {
        _setExecPostsTo(chainId, postsTo);
    }

    /**
     * @notice Set the size buffer for a chain performing data inclusion.
     */
    function setDataSizeBuffer(uint64 chainId, uint64 sizeBuffer) external onlyManager {
        _setDataSizeBuffer(chainId, sizeBuffer);
    }

    /**
     * @notice Set the execution gas price for a destination chain.
     */
    function setExecGasPrice(uint64 chainId, uint64 gasPrice) external onlyManager {
        _setExecGasPrice(chainId, gasPrice);
    }

    /**
     * @notice Set the data inclusion gas price for a chain.
     */
    function setDataGasPrice(uint64 chainId, uint64 gasPrice) external onlyManager {
        _setDataGasPrice(chainId, gasPrice);
    }

    /**
     * @notice Set the to-native conversion rate for an execution chain.
     */
    function setExecToNativeRate(uint64 chainId, uint64 rate) external onlyManager {
        _setExecToNativeRate(chainId, rate);
    }

    /**
     * @notice Set the to-native conversion rate for a data inclusion chain.
     */
    function setDataToNativeRate(uint64 chainId, uint64 rate) external onlyManager {
        _setDataToNativeRate(chainId, rate);
    }

    /**
     * @notice Set the base protocol fee for each xmsg.
     */
    function setProtocolFee(uint72 fee) external onlyOwner {
        _setProtocolFee(fee);
    }

    /**
     * @notice Set the base gas limit for each xmsg.
     */
    function setBaseGasLimit(uint24 gasLimit) external onlyOwner {
        _setBaseGasLimit(gasLimit);
    }

    /**
     * @notice Set the manager admin account.
     */
    function setManager(address manager_) external onlyOwner {
        require(manager_ != address(0), "FeeOracleV2: no zero manager");
        _setManager(manager_);
    }

    /**
     * @notice Set the execution fee parameters for a list of destination chains.
     */
    function _bulkSetExecFeeParams(ExecFeeParams[] calldata params) internal {
        for (uint256 i = 0; i < params.length; i++) {
            ExecFeeParams memory p = params[i];

            require(p.execGasPrice > 0, "FeeOracleV2: no zero gas price");
            require(p.postsTo > 0, "FeeOracleV2: no zero chain id");
            require(p.toNativeRate > 0, "FeeOracleV2: no zero rate");
            require(p.chainId != 0, "FeeOracleV2: no zero chain id");

            _execFeeParams[p.chainId] = p;

            emit ExecFeeParamsSet(p.chainId, p.postsTo, p.execGasPrice, p.toNativeRate);
        }
    }

    /**
     * @notice Set the data inclusion fee parameters for a list of chains.
     */
    function _bulkSetDataFeeParams(DataFeeParams[] calldata params) internal {
        for (uint256 i = 0; i < params.length; i++) {
            DataFeeParams memory p = params[i];

            require(p.dataGasPrice > 0, "FeeOracleV2: no zero gas price");
            require(p.toNativeRate > 0, "FeeOracleV2: no zero rate");
            require(p.chainId != 0, "FeeOracleV2: no zero chain id");

            _dataFeeParams[p.chainId] = p;

            emit DataFeeParamsSet(p.chainId, p.sizeBuffer, p.dataGasPrice, p.toNativeRate);
        }
    }

    /**
     * @notice Set the chain ID to which this execution chain posts tx calldata.
     */
    function _setExecPostsTo(uint64 chainId, uint64 postsTo) internal {
        require(postsTo != 0, "FeeOracleV2: no zero chain id");
        require(chainId != 0, "FeeOracleV2: no zero chain id");

        _execFeeParams[chainId].postsTo = postsTo;
        emit ExecPostsToSet(chainId, postsTo);
    }

    /**
     * @notice Set the size buffer for a data inclusion chain.
     */
    function _setDataSizeBuffer(uint64 chainId, uint64 sizeBuffer) internal {
        require(chainId != 0, "FeeOracleV2: no zero chain id");

        _dataFeeParams[chainId].sizeBuffer = sizeBuffer;
        emit DataSizeBufferSet(chainId, sizeBuffer);
    }

    /**
     * @notice Set the execution gas price for a destination chain.
     */
    function _setExecGasPrice(uint64 chainId, uint64 gasPrice) internal {
        require(gasPrice > 0, "FeeOracleV2: no zero gas price");
        require(chainId != 0, "FeeOracleV2: no zero chain id");

        _execFeeParams[chainId].execGasPrice = gasPrice;
        emit ExecGasPriceSet(chainId, gasPrice);
    }

    /**
     * @notice Set the data inclusion gas price for a destination chain.
     */
    function _setDataGasPrice(uint64 chainId, uint64 gasPrice) internal {
        require(gasPrice > 0, "FeeOracleV2: no zero gas price");
        require(chainId != 0, "FeeOracleV2: no zero chain id");

        _dataFeeParams[chainId].dataGasPrice = gasPrice;
        emit DataGasPriceSet(chainId, gasPrice);
    }

    /**
     * @notice Set the to-native conversion rate for an execution chain.
     */
    function _setExecToNativeRate(uint64 chainId, uint64 rate) internal {
        require(rate > 0, "FeeOracleV2: no zero rate");
        require(chainId != 0, "FeeOracleV2: no zero chain id");

        _execFeeParams[chainId].toNativeRate = rate;
        emit ExecToNativeRateSet(chainId, rate);
    }

    /**
     * @notice Set the to-native conversion rate for a data inclusion chain.
     */
    function _setDataToNativeRate(uint64 chainId, uint64 rate) internal {
        require(rate > 0, "FeeOracleV2: no zero rate");
        require(chainId != 0, "FeeOracleV2: no zero chain id");

        _dataFeeParams[chainId].toNativeRate = rate;
        emit DataToNativeRateSet(chainId, rate);
    }

    /**
     * @notice Set the base protocol fee for each xmsg.
     */
    function _setProtocolFee(uint72 fee) internal {
        protocolFee = fee;
        emit ProtocolFeeSet(fee);
    }

    /**
     * @notice Set the base gas limit for each xmsg.
     */
    function _setBaseGasLimit(uint24 gasLimit) internal {
        baseGasLimit = gasLimit;
        emit BaseGasLimitSet(gasLimit);
    }

    /**
     * @notice Set the manager admin account.
     */
    function _setManager(address manager_) internal {
        manager = manager_;
        emit ManagerSet(manager_);
    }
}
