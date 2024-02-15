// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import { DelegationManager } from "eigenlayer-contracts/src/contracts/core/DelegationManager.sol";
import { ISignatureUtils } from "eigenlayer-contracts/src/contracts/interfaces/ISignatureUtils.sol";
import { IStrategy } from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";
import { IDelegationManager } from "eigenlayer-contracts/src/contracts/interfaces/IDelegationManager.sol";

import { OperatorStateRetriever } from "eigenlayer-middleware/src/OperatorStateRetriever.sol";
import { ServiceManagerBase } from "eigenlayer-middleware/src/ServiceManagerBase.sol";
import { RegistryCoordinator } from "eigenlayer-middleware/src/RegistryCoordinator.sol";
import { IRegistryCoordinator } from "eigenlayer-middleware/src/interfaces/IRegistryCoordinator.sol";
import { IStakeRegistry } from "eigenlayer-middleware/src/interfaces/IStakeRegistry.sol";

import { OmniPredeploys } from "../libraries/OmniPredeploys.sol";
import { IOmniPortal } from "../interfaces/IOmniPortal.sol";
import { IOmniEthRestaking } from "../interfaces/IOmniEthRestaking.sol";
import { IOmniAVS } from "../interfaces/IOmniAVS.sol";
import { IOmniAVSAdmin } from "../interfaces/IOmniAVSAdmin.sol";

/**
 * @title OmniAVS
 * @notice Omni AVS contract. It is responsible for syncing Omni AVS operator
 *         stake and delegations with the Omni chain.
 */
contract OmniAVS is IOmniAVS, IOmniAVSAdmin, ServiceManagerBase, OperatorStateRetriever {
    /// @dev AVS Quorum numbers, Omni only has one quorum
    bytes public constant QUORUM_NUMBERS = hex"00";

    /// @dev AVS Quorum number
    uint8 public constant QUORUM_NUMBER = 0;

    /// @dev Omni chain id, used to make xcalls to the Omni chain
    uint64 public omniChainId;

    /// @dev Omni portal contract, used to make xcalls to the Omni chain
    IOmniPortal public omni;

    /// @dev List of currently register operators, used to sync EigenCore
    address[] public operators;

    constructor(
        IDelegationManager delegationManager_,
        IRegistryCoordinator registryCoordinator_,
        IStakeRegistry stakeRegistry_
    ) ServiceManagerBase(delegationManager_, registryCoordinator_, stakeRegistry_) { }

    /// @inheritdoc IOmniAVSAdmin
    function initialize(address owner_, IOmniPortal omni_, uint64 omniChainId_) external initializer {
        _transferOwnership(owner_);
        omni = omni_;
        omniChainId = omniChainId_;
    }

    /**
     * Omni sync
     */

    /// @inheritdoc IOmniAVS
    function feeForSync() external returns (uint256) {
        _syncWithEigenLayer();
        Validator[] memory vals = getValidators();
        return omni.feeFor(omniChainId, abi.encodeWithSelector(IOmniEthRestaking.sync.selector, vals));
    }

    /// @inheritdoc IOmniAVS
    function syncWithOmni() external payable {
        _syncWithEigenLayer();
        Validator[] memory vals = getValidators();
        omni.xcall{ value: msg.value }(
            omniChainId,
            OmniPredeploys.OMNI_ETH_RESTAKING,
            abi.encodeWithSelector(IOmniEthRestaking.sync.selector, vals)
        );
    }

    /**
     * ServiceManagerBase overrides
     */

    /// @dev Override ServiceManagerBase.registerOperatorToAVS, to track list of operators
    ///      We need to track list of operators so we can call registryCoordinator.updateOperators(operators);
    ///      before syncing with omni.
    function registerOperatorToAVS(
        address operator,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) public override {
        ServiceManagerBase.registerOperatorToAVS(operator, operatorSignature);
        _addOperator(operator);
    }

    /// @dev Override ServiceManagerBase.deregisterOperatorFromAVS, to track list of operators
    ///      We need to track list of operators so we can call registryCoordinator.updateOperators(operators);
    ///      before syncing with omni.
    function deregisterOperatorFromAVS(address operator) public override {
        ServiceManagerBase.deregisterOperatorFromAVS(operator);
        _removeOperator(operator);
    }

    /**
     * Admin controls
     */

    /// @inheritdoc IOmniAVSAdmin
    function setOmniPortal(IOmniPortal omni_) external onlyOwner {
        omni = omni_;
    }

    /// @inheritdoc IOmniAVSAdmin
    function setOmniChainId(uint64 omniChainId_) external onlyOwner {
        omniChainId = omniChainId_;
    }

    /**
     * View functions
     */

    /// @inheritdoc IOmniAVS
    function getValidators() public view returns (Validator[] memory) {
        return getValidators(block.number);
    }

    /// @inheritdoc IOmniAVS
    function getValidators(uint256 blockNumber) public view returns (Validator[] memory validators) {
        /// we only provide on quorum, so we only need the first element
        Operator[] memory ops = _getOperatorState(blockNumber)[0];

        // We translate OperatorStateRetriever.Operator[] to Validator[], splitting Operator.stake into:
        //  - Validator.delegated (total amount delegated to the operator) and
        //  - Validator.staked (total amount staked by the operator)

        validators = new Validator[](ops.length);

        for (uint256 i = 0; i < ops.length; i++) {
            address addr = ops[i].operator;
            uint96 staked = _getStaked(addr);
            uint96 delegations = ops[i].stake - staked;
            validators[i] = Validator(addr, delegations, staked);
        }

        return validators;
    }

    /**
     * Internal view functions
     */

    /// @dev Read operator state from avs registries, at a specific block number
    function _getOperatorState(uint256 blockNumber) public view returns (Operator[][] memory) {
        return OperatorStateRetriever.getOperatorState(_registryCoordinator, QUORUM_NUMBERS, uint32(blockNumber));
    }

    /// @dev Returns the total amount staked by the operator, not including deletations
    ///      This requires us translate the operators active delegatable shares into strategyParam
    ///      weigted steake amount, for QUORUM_NUMBER.
    function _getStaked(address operator) internal view returns (uint96) {
        IStakeRegistry.StrategyParams[] memory strategyParams = _strategyParams();
        (IStrategy[] memory strategies, uint256[] memory shares) =
            DelegationManager(address(_delegationManager)).getDelegatableShares(operator);

        uint96 staked;

        for (uint256 i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];
            uint256 sharesAmt = shares[i];

            // find the strategy params for the strategy
            IStakeRegistry.StrategyParams memory params;
            for (uint256 j = 0; j < strategyParams.length; j++) {
                if (address(strategyParams[j].strategy) == address(strat)) {
                    params = strategyParams[j];
                    break;
                }
            }

            // if strategy is not found, do not consider it in stake
            if (address(params.strategy) == address(0)) continue;

            // same calculation StakeRegistry.weightOfOperatorForQuorum
            staked += uint96(sharesAmt * params.multiplier / _stakeRegistry.WEIGHTING_DIVISOR());
        }

        return staked;
    }

    /// @dev Returns the strategy params for OmniAVS's single quorum
    function _strategyParams() internal view returns (IStakeRegistry.StrategyParams[] memory params) {
        uint256 paramsLen = _stakeRegistry.strategyParamsLength(QUORUM_NUMBER);
        params = new IStakeRegistry.StrategyParams[](paramsLen);
        for (uint256 i = 0; i < paramsLen; i++) {
            params[i] = _stakeRegistry.strategyParamsByIndex(QUORUM_NUMBER, i);
        }
        return params;
    }

    /**
     * Internal functions.
     */

    /// @dev Sync with OmniAVS StakeRegistry with EigenLayer core. To be called before any
    //       operation that requires an up-to-date view of operator stake.
    function _syncWithEigenLayer() internal {
        RegistryCoordinator(address(_registryCoordinator)).updateOperators(operators);
    }

    /// @dev Adds an operator to the list of operators
    function _addOperator(address operator) internal {
        for (uint256 i = 0; i < operators.length; i++) {
            // If operator already exists, do not add it again.
            // We do not revert. Instead, we allow ServiceManagerBase.registerOperatorToAVS to
            // determine when an "operator already exists", at which point it will revert.
            if (operators[i] == operator) return;
        }

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
}
