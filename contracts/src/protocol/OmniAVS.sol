// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { OwnableUpgradeable } from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import { PausableUpgradeable } from "@openzeppelin-upgrades/contracts/security/PausableUpgradeable.sol";

import { IAVSDirectory } from "eigenlayer-contracts/src/contracts/interfaces/IAVSDirectory.sol";
import { IStrategy } from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";
import { ISignatureUtils } from "eigenlayer-contracts/src/contracts/interfaces/ISignatureUtils.sol";
import { IServiceManager } from "eigenlayer-middleware/src/interfaces/IServiceManager.sol";

import { OmniPredeploys } from "../libraries/OmniPredeploys.sol";
import { IDelegationManager } from "../interfaces/IDelegationManager.sol";
import { IOmniEthRestaking } from "../interfaces/IOmniEthRestaking.sol";
import { IOmniPortal } from "../interfaces/IOmniPortal.sol";
import { IOmniAVS } from "../interfaces/IOmniAVS.sol";
import { IOmniAVSAdmin } from "../interfaces/IOmniAVSAdmin.sol";

import { OmniAVSStorage } from "./OmniAVSStorage.sol";

/**
 * @title OmniAVS
 * @notice Omni's AVS contract. It is responsible for faciiltating registration / deregistration of
 *         EigenLayer opators, and for syncing operator delegations with the Omni chain.
 */
contract OmniAVS is
    IOmniAVS,
    IOmniAVSAdmin,
    IServiceManager,
    OwnableUpgradeable,
    PausableUpgradeable,
    OmniAVSStorage
{
    /// @notice Constant used as a divisor in calculating weights
    uint256 internal constant STRATEGY_WEIGHTING_DIVISOR = 1e18;

    /// @notice EigenLayer core DelegationManager
    IDelegationManager internal immutable _delegationManager;

    /// @notice EigenLayer core AVSDirectory
    IAVSDirectory internal immutable _avsDirectory;

    constructor(IDelegationManager delegationManager_, IAVSDirectory avsDirectory_) {
        _delegationManager = delegationManager_;
        _avsDirectory = avsDirectory_;
        _disableInitializers();
    }

    /**
     * @notice Initialize the Omni AVS admin contract.
     * @param owner_            Intiial contract owner
     * @param omni_             Omni portal contract
     * @param omniChainId_      Omni chain id
     * @param strategyParams_   List of accepted strategies and their multipliers
     */
    function initialize(
        address owner_,
        IOmniPortal omni_,
        uint64 omniChainId_,
        StrategyParam[] calldata strategyParams_
    ) external initializer {
        omni = omni_;
        omniChainId = omniChainId_;

        _transferOwnership(owner_);
        _setStrategyParams(strategyParams_);
    }

    //////////////////////////////////////////////////////////////////////////////
    //                          Operator Registration                           //
    //////////////////////////////////////////////////////////////////////////////

    /**
     * @notice Forwards a call to EigenLayer's DelegationManager contract to confirm operator
     *         registration with the AVS
     * @dev Adds operator address to internally tracked list of operators
     * @param operator          The address of the operator to register.
     * @param operatorSignature The signature, salt, and expiry of the operator's signature.
     */
    function registerOperatorToAVS(
        address operator,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) external whenNotPaused {
        require(msg.sender == operator, "OmniAVS: only operator");
        require(_allowlist[operator], "OmniAVS: not allowed");
        require(!_isOperator(operator), "OmniAVS: already an operator"); // we could let _avsDirectory.regsiterOperatorToAVS handle this, they do check

        _avsDirectory.registerOperatorToAVS(operator, operatorSignature);
        _addOperator(operator);

        emit OperatorAdded(operator);
    }

    /**
     * @notice Forwards a call to EigenLayer's DelegationManager contract to confirm operator deregistration from the AVS
     * @dev Removes operator address from internally tracked list of operators
     * @param operator The address of the operator to deregister.
     */
    function deregisterOperatorFromAVS(address operator) external whenNotPaused {
        require(msg.sender == operator || msg.sender == owner(), "OmniAVS: only operator or owner");
        require(_isOperator(operator), "OmniAVS: not an operator");

        _avsDirectory.deregisterOperatorFromAVS(operator);
        _removeOperator(operator);

        emit OperatorRemoved(operator);
    }

    /**
     * @notice Returns true if the operator is in the allowlist.
     * @param operator The operator to check
     */
    function isInAllowlist(address operator) external view returns (bool) {
        return _allowlist[operator];
    }

    /**
     * @inheritdoc IServiceManager
     * @dev Implemented to match IServiceManager interface - required for compatibility with
     *      eigenlayer frontend.
     */
    function avsDirectory() external view returns (address) {
        return address(_avsDirectory);
    }

    //////////////////////////////////////////////////////////////////////////////
    //                              Omni Sync                                   //
    //////////////////////////////////////////////////////////////////////////////

    /**
     * @notice Sync OmniAVS operator stake & delegations with Omni chain.
     */
    function syncWithOmni() external payable whenNotPaused {
        Operator[] memory ops = _getOperators();
        omni.xcall{ value: msg.value }(
            omniChainId,
            OmniPredeploys.OMNI_ETH_RESTAKING,
            abi.encodeWithSelector(IOmniEthRestaking.sync.selector, ops),
            _xcallGasLimitFor(ops.length)
        );
    }

    /**
     * @notice Returns the fee required for syncWithOmni(), for the current operator set.
     */
    function feeForSync() external view returns (uint256) {
        Operator[] memory ops = _getOperators();
        return omni.feeFor(
            omniChainId, abi.encodeWithSelector(IOmniEthRestaking.sync.selector, ops), _xcallGasLimitFor(ops.length)
        );
    }

    //////////////////////////////////////////////////////////////////////////////
    //                              AVS Views                                   //
    //////////////////////////////////////////////////////////////////////////////

    /**
     * @notice Returns the currrent list of operator registered as OmniAVS.
     *         Operator.addr        = The operator's ethereum address
     *         Operator.staked      = The total amount staked by the operator, not including delegations
     *         Operator.delegated   = The total amount delegated, not including operator stake
     */
    function operators() external view returns (Operator[] memory) {
        return _getOperators();
    }

    /**
     * @notice Returns the current strategy parameters. Strategy parameters determine which
     *         eigenlayer strateies the AVS considers when determining operator stake.
     */
    function strategyParams() external view returns (StrategyParam[] memory) {
        return _strategyParams;
    }

    /**
     * @inheritdoc IServiceManager
     * @dev Implemented to match IServiceManager interface - required for compatibility with
     *      eigenlayer frontend.
     */
    function getRestakeableStrategies() external view returns (address[] memory) {
        address[] memory strategies = new address[](_strategyParams.length);
        for (uint256 j = 0; j < _strategyParams.length; j++) {
            strategies[j] = address(_strategyParams[j].strategy);
        }
        return strategies;
    }

    /**
     * @inheritdoc IServiceManager
     * @dev Implemented to match IServiceManager interface - required for compatibility with
     *      eigenlayer frontend.
     */
    function getOperatorRestakedStrategies(address operator) external view returns (address[] memory) {
        address[] memory strategies = new address[](_strategyParams.length);
        for (uint256 j = 0; j < _strategyParams.length; j++) {
            address strat = address(_strategyParams[j].strategy);
            if (_delegationManager.operatorShares(operator, IStrategy(strat)) > 0) {
                strategies[j] = strat;
            }
        }
        return strategies;
    }

    //////////////////////////////////////////////////////////////////////////////
    //                              Admin functions                             //
    //////////////////////////////////////////////////////////////////////////////

    /**
     * @inheritdoc IServiceManager
     * @dev Sets AVS metadata URI with the AVSDirectory. Implemented to match IServiceManager interface
     */
    function setMetadataURI(string memory metadataURI) external onlyOwner {
        _avsDirectory.updateAVSMetadataURI(metadataURI);
    }

    /**
     * @notice Set the Omni portal contract.
     * @param portal The Omni portal contract
     */
    function setOmniPortal(IOmniPortal portal) external onlyOwner {
        omni = portal;
    }

    /**
     * @notice Set the Omni chain id.
     * @param chainId The Omni chain id
     */
    function setOmniChainId(uint64 chainId) external onlyOwner {
        omniChainId = chainId;
    }

    /**
     * @notice Set the strategy parameters.
     * @param params The strategy parameters
     */
    function setStrategyParams(StrategyParam[] calldata params) external onlyOwner {
        _setStrategyParams(params);
    }

    /**
     * @notice Set the xcall gas limits.
     * @param base          The base xcall gas limit
     * @param perOperator   The per-operator additional xcall gas limit
     */
    function setXcallGasLimits(uint64 base, uint64 perOperator) external onlyOwner {
        xcallBaseGasLimit = base;
        xcallGasLimitPerOperator = perOperator;
    }

    /**
     * @notice Add an operator to the allowlist.
     * @param operator The operator to add
     */
    function addToAllowlist(address operator) external onlyOwner {
        require(operator != address(0), "OmniAVS: zero address");
        require(!_allowlist[operator], "OmniAVS: already in allowlist");
        _allowlist[operator] = true;
        emit OperatorAllowed(operator);
    }

    /**
     * @notice Remove an operator from the allowlist.
     * @param operator The operator to remove
     */
    function removeFromAllowlist(address operator) external onlyOwner {
        require(_allowlist[operator], "OmniAVS: not in allowlist");
        _allowlist[operator] = false;
        emit OperatorDisallowed(operator);
    }

    /**
     * @notice Pause the contract.
     * @dev This pauses registerOperatorToAVS, deregisterOperatorFromAVS, and syncWithOmni.
     */
    function pause() external onlyOwner {
        _pause();
    }

    /**
     * @notice Unpause the contract.
     */
    function unpause() external onlyOwner {
        _unpause();
    }

    //////////////////////////////////////////////////////////////////////////////
    //                              Internal setters                            //
    //////////////////////////////////////////////////////////////////////////////

    /**
     * @dev Adds an operator to the list of operators, does not check if operator already exists
     */
    function _addOperator(address operator) private {
        _operators.push(operator);
    }

    /**
     * @notice Removes an operator from the list of operators
     */
    function _removeOperator(address operator) private {
        for (uint256 i = 0; i < _operators.length; i++) {
            if (_operators[i] == operator) {
                _operators[i] = _operators[_operators.length - 1];
                _operators.pop();
                break;
            }
        }
    }

    /**
     * @notice Returns true if the operator is in the list of operators
     */
    function _isOperator(address operator) private view returns (bool) {
        for (uint256 i = 0; i < _operators.length; i++) {
            if (_operators[i] == operator) {
                return true;
            }
        }
        return false;
    }

    /**
     * @notice Set the strategy parameters.
     * @param params The strategy parameters
     */
    function _setStrategyParams(StrategyParam[] calldata params) internal {
        delete _strategyParams;

        for (uint256 i = 0; i < params.length; i++) {
            // TODO: add zero addr and duplicate strat tests
            require(address(params[i].strategy) != address(0), "OmniAVS: zero strategy");

            // ensure no duplicates
            for (uint256 j = i + 1; j < params.length; j++) {
                require(address(params[i].strategy) != address(params[j].strategy), "OmniAVS: duplicate strategy");
            }

            _strategyParams.push(params[i]);
        }
    }

    //////////////////////////////////////////////////////////////////////////////
    //                              Internal views                              //
    //////////////////////////////////////////////////////////////////////////////

    /**
     * @notice Returns the gas limit for OmniEthRestaking.sync xcall for some number of operators
     */
    function _xcallGasLimitFor(uint256 numOperators) internal view returns (uint64) {
        return uint64(numOperators * xcallGasLimitPerOperator + xcallBaseGasLimit);
    }

    /**
     * @notice Return current list of Operators, including their personal stake and delegated stake
     */
    function _getOperators() internal view returns (Operator[] memory) {
        Operator[] memory ops = new Operator[](_operators.length);

        for (uint256 i = 0; i < ops.length; i++) {
            address operator = _operators[i];

            uint96 total = _getTotalDelegations(operator);
            uint96 staked = _getSelfDelegations(operator);

            // this should never happen, but just in case
            uint96 delegated = total > staked ? total - staked : 0;

            ops[i] = Operator(operator, delegated, staked);
        }

        return ops;
    }

    /**
     * @notice Returns the operator's self-delegations
     * @param operator The operator address
     */
    function _getSelfDelegations(address operator) internal view returns (uint96) {
        (IStrategy[] memory strategies, uint256[] memory shares) = _delegationManager.getDelegatableShares(operator);

        uint96 staked;
        for (uint256 i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];

            // find the strategy params for the strategy
            StrategyParam memory params;
            for (uint256 j = 0; j < _strategyParams.length; j++) {
                if (address(_strategyParams[j].strategy) == address(strat)) {
                    params = _strategyParams[j];
                    break;
                }
            }

            // if strategy is not found, do not consider it in stake
            if (address(params.strategy) == address(0)) continue;

            // TODO: should we convert shares to underlying?
            // uint256 amt = IStrategy(params.strategy).sharesToUnderlying(shares[i]);
            // This would convert "shares in the stETH strategy" to "stETH tokens"
            // Shares do not map 1:1 to underlying for rebalancing tokens

            staked += _weight(shares[i], params.multiplier);
        }

        return staked;
    }

    /**
     * @notice Returns total delegations to the operator, including self delegations
     * @param operator The operator address
     */
    function _getTotalDelegations(address operator) internal view returns (uint96) {
        uint96 total;
        StrategyParam memory params;

        for (uint256 j = 0; j < _strategyParams.length; j++) {
            params = _strategyParams[j];
            uint256 shares = _delegationManager.operatorShares(operator, params.strategy);

            // TODO: should we convert shares to underlying?
            // uint256 amt = IStrategy(params.strategy).sharesToUnderlying(shares);
            // This would convert "shares in the stETH strategy" to "stETH tokens"
            // Shares do not map 1:1 to underlying for rebalancing tokens

            total += _weight(shares, params.multiplier);
        }

        return total;
    }

    /**
     * @notice Returns the weighted stake for shares with specified multiplier
     */
    function _weight(uint256 shares, uint96 multiplier) internal pure returns (uint96) {
        return uint96(shares * multiplier / STRATEGY_WEIGHTING_DIVISOR);
    }
}
