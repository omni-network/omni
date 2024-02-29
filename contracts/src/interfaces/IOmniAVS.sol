// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import { IStrategy } from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";

/**
 * @title IOmniAVS
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

    /**
     * @notice Struct representing an OmniAVS operator
     * @custom:field addr       The operator's ethereum address
     * @custom:field delegated  The total amount delegated, not including operator stake
     * @custom:field staked     The total amount staked by the operator, not including delegations
     */
    struct Operator {
        address addr;
        uint96 delegated;
        uint96 staked;
    }

    /**
     * @notice Represents a single supported strategy.
     * @custom:field strategy   The strategy contract
     * @custom:field multiplier The stake multiplier, to weight strategy against others
     */
    struct StrategyParam {
        IStrategy strategy;
        uint96 multiplier;
    }

    /**
     * @notice Returns the fee required for syncWithOmni(), for the current operator set.
     */
    function feeForSync() external view returns (uint256);

    /**
     * @notice Sync OmniAVS operator stake & delegations with Omni chain.
     */
    function syncWithOmni() external payable;

    /**
     * @notice Returns the currrent list of operator registered as OmniAVS.
     */
    function operators() external view returns (Operator[] memory);

    /**
     * @notice Returns the current strategy parameters.
     */
    function strategyParams() external view returns (StrategyParam[] memory);
}
