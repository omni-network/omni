// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import { IDelegationManager } from "eigenlayer-contracts/src/contracts/interfaces/IDelegationManager.sol";

import { OperatorStateRetriever } from "eigenlayer-middleware/src/OperatorStateRetriever.sol";
import { IRegistryCoordinator } from "eigenlayer-middleware/src/interfaces/IRegistryCoordinator.sol";
import { IStakeRegistry } from "eigenlayer-middleware/src/interfaces/IStakeRegistry.sol";

import { OmniAVS } from "src/protocol/OmniAVS.sol";

/**
 * @title OmniAVSHarness
 * @dev A wrapper over OmniAVS that exposes internal functions for testing.
 */
contract OmniAVSHarness is OmniAVS {
    constructor(
        IDelegationManager delegationManager_,
        IRegistryCoordinator registryCoordinator_,
        IStakeRegistry stakeRegistry_
    ) OmniAVS(delegationManager_, registryCoordinator_, stakeRegistry_) { }

    function getOperatorState() external view returns (Operator[][] memory) {
        return _getOperatorState(block.number);
    }
}
