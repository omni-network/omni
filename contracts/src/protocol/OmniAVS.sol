// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import { ISignatureUtils } from "eigenlayer-contracts/src/contracts/interfaces/ISignatureUtils.sol";
import { IStrategy } from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";
import { IDelegationManager } from "eigenlayer-contracts/src/contracts/interfaces/IDelegationManager.sol";
import { DelegationManager } from "eigenlayer-contracts/src/contracts/core/DelegationManager.sol";

import { OperatorStateRetriever } from "eigenlayer-middleware/src/OperatorStateRetriever.sol";
import { ServiceManagerBase } from "eigenlayer-middleware/src/ServiceManagerBase.sol";
import { IRegistryCoordinator } from "eigenlayer-middleware/src/interfaces/IRegistryCoordinator.sol";
import { IStakeRegistry } from "eigenlayer-middleware/src/interfaces/IStakeRegistry.sol";

import { IOmniPortal } from "../interfaces/IOmniPortal.sol";
import { OmniPredeploys } from "../libraries/OmniPredeploys.sol";

/**
 * @title IOmniEthRestaking
 * @dev Interface for OmniEthRestaking predeployed contract, that receives operator stake updates from OmniAVS
 *       TOOD: implement OmniEthRestaking (name TBD), move to separate file
 */
interface IOmniEthRestaking {
    /// @dev Syncs operator state with OmniAVS. Only callable by XMsg from OmniAVS
    function sync(OmniAVS.Validator[] calldata validators) external;
}

/**
 * @title OmniAVS
 */
contract OmniAVS is ServiceManagerBase, OperatorStateRetriever {
    // using BytesLib for bytes;

    struct Validator {
        // ethereum address of the operator
        address addr;
        // amount of delegated, not including operator stake
        uint96 delegated;
        // total amount staked by the operator
        uint96 staked;
    }

    /// @dev AVS Quorum numbers, Omni only has one quorum
    bytes public constant QUORUM_NUMBERS = hex"00";

    /// @dev AVS Quorum number
    uint8 public constant QUORUM_NUMBER = 0;

    /// TODO: what are ramifications of hardcoding single quorum as constants?

    uint64 public immutable omniChainId;
    IOmniPortal public immutable omni;

    address[] public operators;

    constructor(
        IDelegationManager delegationManager_,
        IRegistryCoordinator registryCoordinator_,
        IStakeRegistry stakeRegistry_,
        IOmniPortal omniPortal_,
        uint64 omniChainId_
    ) ServiceManagerBase(delegationManager_, registryCoordinator_, stakeRegistry_) {
        omni = omniPortal_;
        omniChainId = omniChainId_;
    }

    /// @dev Allow relayer to calculate xcall fee for syncWithOmni
    function feeForSync() external view returns (uint256) {
        Validator[] memory vals = getValidators();
        return omni.feeFor(omniChainId, abi.encodeWithSelector(IOmniEthRestaking.sync.selector, vals));
    }

    /// @dev Syncs operator state with OmniEthRestaking predeploy
    function syncWithOmni() external payable {
        Validator[] memory vals = getValidators();
        omni.xcall{ value: msg.value }(
            omniChainId,
            OmniPredeploys.OMNI_ETH_RESTAKING,
            abi.encodeWithSelector(IOmniEthRestaking.sync.selector, vals)
        );
    }

    /// @dev Override ServiceManagerBase.registerOperatorToAVS, to track list of operators
    ///      We need to track list of operators so we can call registryCoordinator.updateOperators(operators);
    ///      before syncing with omni.
    ///      TODO: call before sync - RegistryCoordinator(address(_registryCoordinator)).updateOperators(operators);
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
    ///      TODO: call before sync - RegistryCoordinator(address(_registryCoordinator)).updateOperators(operators);
    function deregisterOperatorFromAVS(address operator) public override {
        ServiceManagerBase.deregisterOperatorFromAVS(operator);

        // TODO: can we remove operator here? do we need to wait for withdrawals?
        _removeOperator(operator);
    }

    /// @dev Returns the list of validators for the current block
    function getValidators() public view returns (Validator[] memory) {
        return getValidators(block.number);
    }

    /// @dev Returns the list of validators for the given block
    function getValidators(uint256 blockNumber) public view returns (Validator[] memory validators) {
        /// we only provide on quorum, so we only need the first element of getOperatorState
        Operator[] memory ops = getOperatorState(blockNumber)[0];

        // we translate Operator[] to Validator[], splitting Operator.stake
        // into Validator.delegated (total amount delegated to the operator)
        // and Validator.staked (total amount staked by the operator)
        validators = new Validator[](ops.length);

        for (uint256 i = 0; i < ops.length; i++) {
            address addr = ops[i].operator;
            uint96 staked = _getStaked(addr);
            uint96 delegations = ops[i].stake - staked;
            validators[i] = Validator(addr, delegations, staked);
        }

        return validators;
    }

    /// @dev exposed for now, for testing
    function getOperatorState() public view returns (Operator[][] memory) {
        return getOperatorState(block.number);
    }

    /// @dev exposed for now, for testing
    function getOperatorState(uint256 blockNumber) public view returns (Operator[][] memory) {
        return OperatorStateRetriever.getOperatorState(_registryCoordinator, QUORUM_NUMBERS, uint32(blockNumber));
    }

    /// @dev Returns the total amount staked by the operator, not including deletations
    ///      This requires us translate the operators active delegatable shares into strategyParam
    ///      weigted steake amount, for QUORUM_NUMBER.
    ///      TODO: this is a big question mark - we are not using eigen the way they intended
    function _getStaked(address operator) internal view returns (uint96) {
        DelegationManager delegation = DelegationManager(address(_delegationManager));
        (IStrategy[] memory strategies, uint256[] memory shares) = delegation.getDelegatableShares(operator);

        IStakeRegistry.StrategyParams[] memory strategyParams = _strategyParams();

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

            // TODO: what should we do if the strategy is not found? this should not happen
            if (address(params.strategy) == address(0)) {
                continue;
            }

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

    /// @dev Adds an operator to the list of operators
    function _addOperator(address operator) internal {
        for (uint256 i = 0; i < operators.length; i++) {
            // TODO: we may not want to revert here, and instead just return
            // allow ServiceManagerBase.registerOperatorToAVS to determine when an "operator already exists"
            require(operators[i] != operator, "Operator already exists");
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
