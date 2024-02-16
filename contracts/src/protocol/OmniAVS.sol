// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { OwnableUpgradeable } from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";

import { DelegationManager } from "eigenlayer-contracts/src/contracts/core/DelegationManager.sol";
import { IDelegationManager } from "eigenlayer-contracts/src/contracts/interfaces/IDelegationManager.sol";
import { IAVSDirectory } from "eigenlayer-contracts/src/contracts/interfaces/IAVSDirectory.sol";
import { IStrategy } from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";
import { ISignatureUtils } from "eigenlayer-contracts/src/contracts/interfaces/ISignatureUtils.sol";
import { IServiceManager } from "eigenlayer-middleware/src/interfaces/IServiceManager.sol";

import { OmniPredeploys } from "../libraries/OmniPredeploys.sol";
import { IOmniPortal } from "../interfaces/IOmniPortal.sol";
import { IOmniEthRestaking } from "../interfaces/IOmniEthRestaking.sol";
import { IOmniAVS } from "../interfaces/IOmniAVS.sol";
import { IOmniAVSAdmin } from "../interfaces/IOmniAVSAdmin.sol";

contract OmniAVS is IOmniAVS, IOmniAVSAdmin, IServiceManager, OwnableUpgradeable {
    /// @notice Constant used as a divisor in calculating weights
    uint256 public constant WEIGHTING_DIVISOR = 1e18;

    /// @notice EigenLayer core DelegationManager
    IDelegationManager public immutable _delegationManager;

    /// @notice EigenLayer core AVSDirectory
    IAVSDirectory public immutable _avsDirectory;

    /// @notice Maximum number of operators that can be registered
    uint32 public maxOperatorCount;

    /// @notice Omni chain id, used to make xcalls to the Omni chain
    uint64 public omniChainId;

    /// @notice Minimum stake required for an operator to register, not including delegations
    uint96 public minimumOperatorStake;

    /// @notice List of currently register operators, used to sync EigenCore
    address[] public operators;

    /// @notice Omni portal contract, used to make xcalls to the Omni chain
    IOmniPortal public omni;

    /// @notice Strategy parameters for restaking
    IOmniAVS.StrategyParams[] public strategyParams;

    /// @dev OmniPortal.xcall gas limit per each validator in syncWithOmni
    uint256 internal xcallGasLimitPerValidator = 10_000;

    /// @dev OmniPortal.xcall base gas limit in syncWithOmni
    uint256 internal xcallBaseGasLimit = 75_000;

    constructor(IDelegationManager delegationManager_, IAVSDirectory avsDirectory_) {
        _delegationManager = delegationManager_;
        _avsDirectory = avsDirectory_;
        _disableInitializers();
    }

    /// @inheritdoc IOmniAVSAdmin
    function initialize(
        address owner_,
        IOmniPortal omni_,
        uint64 omniChainId_,
        uint96 minimumOperatorStake_,
        uint32 maxOperatorCount_,
        StrategyParams[] calldata strategyParams_
    ) external initializer {
        _transferOwnership(owner_);
        omni = omni_;
        omniChainId = omniChainId_;
        minimumOperatorStake = minimumOperatorStake_;
        maxOperatorCount = maxOperatorCount_;
        _setStrategyParams(strategyParams_);
    }

    /**
     * Omni sync
     */

    /// @inheritdoc IOmniAVS
    function feeForSync() external view returns (uint256) {
        Validator[] memory vals = getValidators();
        return omni.feeFor(
            omniChainId, abi.encodeWithSelector(IOmniEthRestaking.sync.selector, vals), _xcallGasLimitFor(vals.length)
        );
    }

    /// @inheritdoc IOmniAVS
    function syncWithOmni() external payable {
        Validator[] memory vals = getValidators();
        omni.xcall{ value: msg.value }(
            omniChainId,
            OmniPredeploys.OMNI_ETH_RESTAKING,
            abi.encodeWithSelector(IOmniEthRestaking.sync.selector, vals),
            _xcallGasLimitFor(vals.length)
        );
    }

    /// @dev Returns the gas limit for OmniEthRestaking.sync xcall for some number of validators
    function _xcallGasLimitFor(uint256 numValidators) internal view returns (uint64) {
        return uint64(numValidators * xcallGasLimitPerValidator + xcallBaseGasLimit);
    }

    /**
     * Operator registration
     */

    /// @inheritdoc IServiceManager
    function registerOperatorToAVS(
        address operator,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) external {
        require(msg.sender == operator, "OmniAVS: only operator");
        require(operators.length < maxOperatorCount, "OmniAVS: max operators reached");
        require(_getStaked(operator) >= minimumOperatorStake, "OmniAVS: minimum stake not met");
        require(!_isOperator(operator), "OmniAVS: already an operator"); // we could let delegation.regsiterOperatorToAVS handle this, they do check

        _avsDirectory.registerOperatorToAVS(operator, operatorSignature);
        _addOperator(operator);

        emit OperatorAdded(operator);
    }

    /// @inheritdoc IServiceManager
    function deregisterOperatorFromAVS(address operator) external {
        require(msg.sender == operator, "OmniAVS: only operator");
        require(_isOperator(operator), "OmniAVS: not an operator");

        _avsDirectory.deregisterOperatorFromAVS(operator);
        _removeOperator(operator);

        emit OperatorRemoved(operator);
    }

    /// @dev Adds an operator to the list of operators, does not check if operator already exists
    function _addOperator(address operator) internal {
        operators.push(operator);
    }

    /// @dev Removes an operator from the list of operators
    function _removeOperator(address operator) internal {
        for (uint256 i = 0; i < operators.length; i++) {
            if (operators[i] == operator) {
                operators[i] = operators[operators.length - 1];
                operators.pop();
                break;
            }
        }
    }

    /// @dev Returns true if the operator is in the list of operators
    function _isOperator(address operator) internal view returns (bool) {
        for (uint256 i = 0; i < operators.length; i++) {
            if (operators[i] == operator) {
                return true;
            }
        }
        return false;
    }

    /**
     * Admin controls
     */

    /// @inheritdoc IServiceManager
    function setMetadataURI(string memory metadataURI) external onlyOwner {
        _avsDirectory.updateAVSMetadataURI(metadataURI);
    }

    /// @inheritdoc IOmniAVSAdmin
    function setOmniPortal(IOmniPortal portal) external onlyOwner {
        omni = portal;
    }

    /// @inheritdoc IOmniAVSAdmin
    function setOmniChainId(uint64 chainId) external onlyOwner {
        omniChainId = chainId;
    }

    /// @inheritdoc IOmniAVSAdmin
    function setStrategyParams(StrategyParams[] calldata params) external onlyOwner {
        _setStrategyParams(params);
    }

    /// @inheritdoc IOmniAVSAdmin
    function setMinimumOperatorStake(uint96 stake) external onlyOwner {
        minimumOperatorStake = stake;
    }

    /// @inheritdoc IOmniAVSAdmin
    function setMaxOperatorCount(uint32 count) external onlyOwner {
        maxOperatorCount = count;
    }

    /// @inheritdoc IOmniAVSAdmin
    function setXcallGasLimits(uint256 base, uint256 perValidator) external onlyOwner {
        xcallBaseGasLimit = base;
        xcallGasLimitPerValidator = perValidator;
    }

    function _setStrategyParams(StrategyParams[] calldata params) internal {
        delete strategyParams;
        for (uint256 i = 0; i < params.length; i++) {
            strategyParams.push(params[i]);
        }
    }

    /**
     * View functions
     */

    /// @inheritdoc IOmniAVS
    function getValidators() public view returns (Validator[] memory) {
        return _getValidators();
    }

    /// @inheritdoc IServiceManager
    function getRestakeableStrategies() external view returns (address[] memory) {
        address[] memory strategies = new address[](strategyParams.length);
        for (uint256 j = 0; j < strategyParams.length; j++) {
            strategies[j] = address(strategyParams[j].strategy);
        }
        return strategies;
    }

    /// @inheritdoc IServiceManager
    function getOperatorRestakedStrategies(address operator) external view returns (address[] memory) {
        address[] memory strategies = new address[](strategyParams.length);
        for (uint256 j = 0; j < strategyParams.length; j++) {
            address strat = address(strategyParams[j].strategy);
            if (_delegationManager.operatorShares(operator, IStrategy(strat)) > 0) strategies[j] = strat;
        }
        return strategies;
    }

    /// @inheritdoc IServiceManager
    function avsDirectory() external view returns (address) {
        return address(_avsDirectory);
    }

    /// @dev Return current list of Validators, including their personal stake and delegated stake
    function _getValidators() internal view returns (Validator[] memory) {
        Validator[] memory vals = new Validator[](operators.length);

        for (uint256 i = 0; i < vals.length; i++) {
            address addr = operators[i];
            uint96 totalStaked;
            StrategyParams memory params;

            // get total opearator stake (their own stake + delegations)
            for (uint256 j = 0; j < strategyParams.length; j++) {
                params = strategyParams[j];

                // shares of the operator in the strategy
                uint256 sharesAmount = _delegationManager.operatorShares(addr, params.strategy);

                // add the weight from the shares for this strategy to the total weight
                if (sharesAmount > 0) totalStaked += _weight(sharesAmount, params.multiplier);
            }

            uint96 staked = _getStaked(addr);
            uint96 delegated = totalStaked - staked;

            vals[i] = Validator(addr, delegated, staked);
        }

        return vals;
    }

    /// @dev Returns the total amount staked by the operator, not including deletations
    function _getStaked(address operator) internal view returns (uint96) {
        (IStrategy[] memory strategies, uint256[] memory shares) =
            DelegationManager(address(_delegationManager)).getDelegatableShares(operator);

        uint96 staked;

        for (uint256 i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];
            uint256 sharesAmt = shares[i];

            // find the strategy params for the strategy
            StrategyParams memory params;
            for (uint256 j = 0; j < strategyParams.length; j++) {
                if (address(strategyParams[j].strategy) == address(strat)) {
                    params = strategyParams[j];
                    break;
                }
            }

            // if strategy is not found, do not consider it in stake
            if (address(params.strategy) == address(0)) continue;

            staked += _weight(sharesAmt, params.multiplier);
        }

        return staked;
    }

    /// @dev Returns the weighted stake for shares with specified multiplier
    function _weight(uint256 shares, uint96 multiplier) internal pure returns (uint96) {
        return uint96(shares * multiplier / WEIGHTING_DIVISOR);
    }
}
