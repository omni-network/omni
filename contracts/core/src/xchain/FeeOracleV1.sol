// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { IFeeOracle } from "../interfaces/IFeeOracle.sol";
import { IFeeOracleV1 } from "../interfaces/IFeeOracleV1.sol";

/**
 * @title FeeOracleV1
 * @notice A simple fee oracle with a fixed fee, controlled by an admin account
 *         Used by OmniPortal to calculate xmsg fees
 */
contract FeeOracleV1 is IFeeOracle, IFeeOracleV1, OwnableUpgradeable {
    /**
     * @notice Base gas limit for each xmsg.
     */
    uint64 public baseGasLimit;

    /**
     * @notice Base protocol fee for each xmsg.
     */
    uint256 public protocolFee;

    /**
     * @notice Address allowed to set gas prices and to-native conversion rates.
     */
    address public manager;

    /**
     * @notice Fee parameters for a specific chain, by chain ID.
     */
    mapping(uint64 => IFeeOracleV1.ChainFeeParams) internal _feeParams;

    /**
     * @notice Denominator for conversion rate calculations.
     */
    uint256 public constant CONVERSION_RATE_DENOM = 1e6;

    modifier onlyManager() {
        require(msg.sender == manager, "FeeOracleV1: not manager");
        _;
    }

    constructor() {
        _disableInitializers();
    }

    function initialize(
        address owner_,
        address manager_,
        uint64 baseGasLimit_,
        uint256 protocolFee_,
        ChainFeeParams[] calldata params
    ) public initializer {
        __Ownable_init(owner_);

        _setManager(manager_);
        _setBaseGasLimit(baseGasLimit_);
        _setProtocolFee(protocolFee_);
        _bulkSetFeeParams(params);
    }

    /// @inheritdoc IFeeOracle
    function feeFor(uint64 destChainId, bytes calldata data, uint64 gasLimit) external view returns (uint256) {
        IFeeOracleV1.ChainFeeParams storage execP = _feeParams[destChainId];
        IFeeOracleV1.ChainFeeParams storage dataP = _feeParams[execP.postsTo];

        require(execP.gasPrice > 0 && execP.toNativeRate > 0, "FeeOracleV1: no fee params");
        require(dataP.gasPrice > 0 && dataP.toNativeRate > 0, "FeeOracleV1: no fee params");

        uint256 execGasPrice = execP.gasPrice * execP.toNativeRate / CONVERSION_RATE_DENOM;
        uint256 dataGasPrice = dataP.gasPrice * dataP.toNativeRate / CONVERSION_RATE_DENOM;

        // 16 gas per non-zero byte, assume non-zero bytes
        // TODO: given we mostly support rollups that post data to L1, it may be cheaper for users to count
        //       non-zero bytes (consuming L2 execution gas) to reduce their L1 data fee
        uint256 dataGas = data.length * 16;

        return protocolFee + (baseGasLimit * execGasPrice) + (gasLimit * execGasPrice) + (dataGas * dataGasPrice);
    }

    /**
     * @notice Returns the fee parameters for a destination chain.
     */
    function feeParams(uint64 chainId) external view returns (ChainFeeParams memory) {
        return _feeParams[chainId];
    }

    /**
     * @notice Returns the gas price for a destination chain.
     */
    function gasPriceOn(uint64 chainId) external view returns (uint256) {
        return _feeParams[chainId].gasPrice;
    }

    /**
     * @notice Returns the to-native conversion rate for a destination chain.
     */
    function toNativeRate(uint64 chainId) external view returns (uint256) {
        return _feeParams[chainId].toNativeRate;
    }

    /**
     * @notice Returns the chainId of the chain that the given destination chain posts tx data to.
     *         For rollups, this is L1.
     */
    function postsTo(uint64 chainId) external view returns (uint64) {
        return _feeParams[chainId].postsTo;
    }

    /**
     * @notice Set the fee parameters for a list of destination chains.
     */
    function bulkSetFeeParams(ChainFeeParams[] calldata params) external onlyManager {
        _bulkSetFeeParams(params);
    }

    /**
     * @notice Set the gas price for a destination chain.
     */
    function setGasPrice(uint64 chainId, uint256 gasPrice) external onlyManager {
        _setGasPrice(chainId, gasPrice);
    }

    /**
     * @notice Set the to native conversion rate for a destination chain.
     */
    function setToNativeRate(uint64 chainId, uint256 rate) external onlyManager {
        _setToNativeRate(chainId, rate);
    }

    /**
     * @notice Set the base gas limit for each xmsg.
     */
    function setBaseGasLimit(uint64 gasLimit) external onlyOwner {
        _setBaseGasLimit(gasLimit);
    }

    /**
     * @notice Set the base protocol fee for each xmsg.
     */
    function setProtocolFee(uint256 fee) external onlyOwner {
        _setProtocolFee(fee);
    }

    /**
     * @notice Set the manager admin account.
     */
    function setManager(address manager_) external onlyOwner {
        require(manager_ != address(0), "FeeOracleV1: no zero manager");
        _setManager(manager_);
    }

    /**
     * @notice Set the fee parameters for a list of destination chains.
     */
    function _bulkSetFeeParams(ChainFeeParams[] calldata params) internal {
        for (uint256 i = 0; i < params.length; i++) {
            ChainFeeParams memory p = params[i];

            require(p.gasPrice > 0, "FeeOracleV1: no zero gas price");
            require(p.toNativeRate > 0, "FeeOracleV1: no zero rate");
            require(p.chainId != 0, "FeeOracleV1: no zero chain id");
            require(p.postsTo != 0, "FeeOracleV1: no zero postsTo");

            _feeParams[p.chainId] = p;

            emit FeeParamsSet(p.chainId, p.postsTo, p.gasPrice, p.toNativeRate);
        }
    }

    /**
     * @notice Set the gas price for a destination chain.
     */
    function _setGasPrice(uint64 chainId, uint256 gasPrice) internal {
        require(gasPrice > 0, "FeeOracleV1: no zero gas price");
        require(chainId != 0, "FeeOracleV1: no zero chain id");

        _feeParams[chainId].gasPrice = gasPrice;
        emit GasPriceSet(chainId, gasPrice);
    }

    /**
     * @notice Set the to-native conversion rate for a destination chain.
     */
    function _setToNativeRate(uint64 chainId, uint256 rate) internal {
        require(rate > 0, "FeeOracleV1: no zero rate");
        require(chainId != 0, "FeeOracleV1: no zero chain id");

        _feeParams[chainId].toNativeRate = rate;
        emit ToNativeRateSet(chainId, rate);
    }

    /**
     * @notice Set the base gas limit for each xmsg.
     */
    function _setBaseGasLimit(uint64 gasLimit) internal {
        baseGasLimit = gasLimit;
        emit BaseGasLimitSet(gasLimit);
    }

    /**
     * @notice Set the base protocol fee for each xmsg.
     */
    function _setProtocolFee(uint256 fee) internal {
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
