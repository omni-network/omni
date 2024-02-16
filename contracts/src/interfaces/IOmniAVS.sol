// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import { IStrategy } from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";
import { ISignatureUtils } from "eigenlayer-contracts/src/contracts/interfaces/ISignatureUtils.sol";

/**
 * @title OmniAVS
 * @notice Interface for the Omni AVS contract. It is responsible for syncing Omni AVS operator
 *         stake and delegations with the Omni chain.
 */
interface IOmniAVS {
    struct Validator {
        // ethereum address of the operator
        address addr;
        // total amount delegated, not including operator stake
        uint96 delegated;
        // total amount staked by the operator, not including delegations
        uint96 staked;
    }

    struct StrategyParams {
        // strategy contract
        IStrategy strategy;
        // stake multiplier, to weight strategy against others
        uint96 multiplier;
    }

    /**
     * @notice Calculate the omni xcall fee for a syncWithOmni call.
     * @dev This is not a view function because it updates state in the OmniAVS
     *      StakeRegistry, by syncing state with EigenLayer core. It is meant to
     *      be called offchain with an eth_call.
     * @return The fee in wei
     */
    function feeForSync() external returns (uint256);

    /**
     * @notice Sync OmniAVS validator stake & delegations with Omni chain.
     */
    function syncWithOmni() external payable;

    /**
     * @notice Get the list of validators registered as OmniAVS operators, with
     *         their stake / delegations.
     * @return The list of validators
     */
    function getValidators() external view returns (Validator[] memory);

    /**
     * @notice Forwards a call to EigenLayer's DelegationManager contract to confirm operator registration with the AVS
     * @param operatorSignature The signature, salt, and expiry of the operator's signature.
     */
    function registerOperator(ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature) external;

    /**
     * @notice Forwards a call to EigenLayer's DelegationManager contract to confirm operator deregistration from the AVS
     */
    function deregisterOperator() external;

    /**
     * @notice Returns the list of strategies that the AVS supports for restaking
     * @dev Matches eigenlayer-middleware ServiceManagerBase.getRestakeableStrategies
     * @dev This function is intended to be called off-chain
     * @dev No guarantee is made on uniqueness of each element in the returned array.
     *      The off-chain service should do that validation separately
     */
    function getRestakeableStrategies() external view returns (address[] memory);

    /**
     * @notice Returns the list of strategies that the operator has potentially restaked on the AVS
     * @param operator The address of the operator to get restaked strategies for
     * @dev Matches eigenlayer-middleware ServiceManagerBase.getOperatorRestakedStrategies
     * @dev This function is intended to be called off-chain
     * @dev No guarantee is made on whether the operator has shares for a strategy in a quorum or uniqueness
     *      of each element in the returned array. The off-chain service should do that validation separately
     */
    function getOperatorRestakedStrategies(address operator) external view returns (address[] memory);
}
