// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import { IStrategy } from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";

/**
 * @title OmniAVS
 * @notice Interface for the Omni AVS contract. It is responsible for syncing Omni AVS operator
 *         stake and delegations with the Omni chain.
 */
interface IOmniAVS {
    /**
     * @notice Emitted when an operator is added to the OmniAVS.
     * @param operator The address of the operator
     */
    event OperatorAdded(address indexed operator);

    /**
     * @notice Emitted when an operator is removed from the OmniAVS.
     * @param operator The address of the operator
     */
    event OperatorRemoved(address indexed operator);

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
     * @return The fee in wei
     */
    function feeForSync() external view returns (uint256);

    /**
     * @notice Sync OmniAVS validator stake & delegations with Omni chain.
     */
    function syncWithOmni() external payable;

    /**
     * @notice Returns the currrent list of validators registered as OmniAVS
     *         operators, with their stake / delegations.
     */
    function getValidators() external view returns (Validator[] memory);

    /**
     * @notice Returns the current strategy parameters.
     */
    function strategyParams() external view returns (StrategyParams[] memory);
}
