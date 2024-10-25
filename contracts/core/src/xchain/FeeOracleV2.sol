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
    uint64 public protocolFee;

    /**
     * @notice Base gas limit for each xmsg.
     */
    uint32 public baseGasLimit;

    /**
     * @notice Address allowed to set gas prices and to-native conversion rates.
     */
    address public manager;

    /**
     * @notice Fee parameters for a specific chain, by chain ID.
     */
    mapping(uint64 => IFeeOracleV2.FeeParams) internal _feeParams;

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
        uint32 baseGasLimit_,
        uint64 protocolFee_,
        FeeParams[] calldata params
    ) public initializer {
        __Ownable_init(owner_);

        _setManager(manager_);
        _setBaseGasLimit(baseGasLimit_);
        _setProtocolFee(protocolFee_);
        _bulkSetFeeParams(params);
    }

    /// @inheritdoc IFeeOracle
    function version() external pure override returns (uint64) {
        return 2;
    }

    /// @inheritdoc IFeeOracle
    function feeFor(uint64 destChainId, bytes calldata data, uint64 gasLimit) external view returns (uint256) {
        IFeeOracleV2.FeeParams storage p = _feeParams[destChainId];

        uint256 _execGasPrice = p.execGasPrice * p.toNativeRate / CONVERSION_RATE_DENOM;
        uint256 _dataGasPrice = p.dataGasPrice * p.toNativeRate / CONVERSION_RATE_DENOM;

        require(_execGasPrice > 0, "FeeOracleV2: no fee params");
        require(_dataGasPrice > 0, "FeeOracleV2: no fee params");

        // 16 gas per non-zero byte, assume non-zero bytes
        // TODO: given we mostly support rollups that post data to L1, it may be cheaper for users to count
        //       non-zero bytes (consuming L2 execution gas) to reduce their L1 data fee
        uint256 dataGas = data.length * 16;

        return protocolFee + (baseGasLimit + gasLimit) * _execGasPrice + (dataGas * _dataGasPrice);
    }

    /**
     * @notice Returns the fee parameters for a destination chain.
     */
    function feeParams(uint64 chainId) external view returns (FeeParams memory) {
        return _feeParams[chainId];
    }

    /**
     * @notice Returns the gas price for a destination chain.
     */
    function execGasPrice(uint64 chainId) external view returns (uint64) {
        return _feeParams[chainId].execGasPrice;
    }

    /**
     * @notice Returns the data gas price for a destination chain.
     */
    function dataGasPrice(uint64 chainId) external view returns (uint64) {
        return _feeParams[chainId].dataGasPrice;
    }

    /**
     * @notice Returns the to-native conversion rate for a destination chain.
     */
    function toNativeRate(uint64 chainId) external view returns (uint256) {
        return _feeParams[chainId].toNativeRate;
    }

    /**
     * @notice Set the fee parameters for a list of destination chains.
     */
    function bulkSetFeeParams(FeeParams[] calldata params) external onlyManager {
        _bulkSetFeeParams(params);
    }

    /**
     * @notice Set the execution gas price for a destination chain.
     */
    function setExecGasPrice(uint64 chainId, uint64 gasPrice) external onlyManager {
        _setExecGasPrice(chainId, gasPrice);
    }

    /**
     * @notice Set the data gas price for a destination chain.
     */
    function setDataGasPrice(uint64 chainId, uint64 gasPrice) external onlyManager {
        _setDataGasPrice(chainId, gasPrice);
    }

    /**
     * @notice Set the to native conversion rate for a destination chain.
     */
    function setToNativeRate(uint64 chainId, uint64 rate) external onlyManager {
        _setToNativeRate(chainId, rate);
    }

    /**
     * @notice Set the base gas limit for each xmsg.
     */
    function setBaseGasLimit(uint32 gasLimit) external onlyOwner {
        _setBaseGasLimit(gasLimit);
    }

    /**
     * @notice Set the base protocol fee for each xmsg.
     */
    function setProtocolFee(uint64 fee) external onlyOwner {
        _setProtocolFee(fee);
    }

    /**
     * @notice Set the manager admin account.
     */
    function setManager(address manager_) external onlyOwner {
        require(manager_ != address(0), "FeeOracleV2: no zero manager");
        _setManager(manager_);
    }

    /**
     * @notice Set the fee parameters for a list of destination chains.
     */
    function _bulkSetFeeParams(FeeParams[] calldata params) internal {
        for (uint256 i = 0; i < params.length; i++) {
            FeeParams memory p = params[i];

            require(p.execGasPrice > 0, "FeeOracleV2: no zero gas price");
            require(p.dataGasPrice > 0, "FeeOracleV2: no zero gas price");
            require(p.toNativeRate > 0, "FeeOracleV2: no zero rate");
            require(p.chainId != 0, "FeeOracleV2: no zero chain id");

            _feeParams[p.chainId] = p;

            emit FeeParamsSet(p.chainId, p.execGasPrice, p.dataGasPrice, p.toNativeRate);
        }
    }

    /**
     * @notice Set the execution gas price for a destination chain.
     */
    function _setExecGasPrice(uint64 chainId, uint64 gasPrice) internal {
        require(gasPrice > 0, "FeeOracleV2: no zero gas price");
        require(chainId != 0, "FeeOracleV2: no zero chain id");

        _feeParams[chainId].execGasPrice = gasPrice;
        emit ExecGasPriceSet(chainId, gasPrice);
    }

    /**
     * @notice Set the data gas price for a destination chain.
     */
    function _setDataGasPrice(uint64 chainId, uint64 gasPrice) internal {
        require(gasPrice > 0, "FeeOracleV2: no zero gas price");
        require(chainId != 0, "FeeOracleV2: no zero chain id");

        _feeParams[chainId].dataGasPrice = gasPrice;
        emit DataGasPriceSet(chainId, gasPrice);
    }

    /**
     * @notice Set the to-native conversion rate for a destination chain.
     */
    function _setToNativeRate(uint64 chainId, uint64 rate) internal {
        require(rate > 0, "FeeOracleV2: no zero rate");
        require(chainId != 0, "FeeOracleV2: no zero chain id");

        _feeParams[chainId].toNativeRate = rate;
        emit ToNativeRateSet(chainId, rate);
    }

    /**
     * @notice Set the base gas limit for each xmsg.
     */
    function _setBaseGasLimit(uint32 gasLimit) internal {
        baseGasLimit = gasLimit;
        emit BaseGasLimitSet(gasLimit);
    }

    /**
     * @notice Set the base protocol fee for each xmsg.
     */
    function _setProtocolFee(uint64 fee) internal {
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
