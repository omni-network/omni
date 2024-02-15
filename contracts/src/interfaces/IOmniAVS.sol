// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import { IStrategy } from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";

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
        IStrategy strategy;
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
}
