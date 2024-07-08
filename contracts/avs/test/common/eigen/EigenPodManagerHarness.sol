// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import { IBeacon } from "@openzeppelin/contracts/proxy/beacon/IBeacon.sol";

import { EigenPodManager } from "eigenlayer-contracts/src/contracts/pods/EigenPodManager.sol";
import { IStrategyManager } from "eigenlayer-contracts/src/contracts/interfaces/IStrategyManager.sol";
import { ISlasher } from "eigenlayer-contracts/src/contracts/interfaces/ISlasher.sol";
import { IDelegationManager } from "eigenlayer-contracts/src/contracts/interfaces/IDelegationManager.sol";
import { IETHPOSDeposit } from "eigenlayer-contracts/src/contracts/interfaces/IETHPOSDeposit.sol";

/**
 * @title EigenPodManagerHarness
 * @dev Test harness for EigenPodManager
 */
contract EigenPodManagerHarness is EigenPodManager {
    constructor(
        IETHPOSDeposit ethPOS,
        IBeacon eigenPodBeacon,
        IStrategyManager strategyManager,
        ISlasher slasher,
        IDelegationManager delegationManager
    ) EigenPodManager(ethPOS, eigenPodBeacon, strategyManager, slasher, delegationManager) { }

    function updatePodOwnerShares(address podOwner, int256 sharesDelta) public {
        int256 currentPodOwnerShares = podOwnerShares[podOwner];
        int256 updatedPodOwnerShares = currentPodOwnerShares + sharesDelta;
        podOwnerShares[podOwner] = updatedPodOwnerShares;

        int256 changeInDelegatableShares = _calculateChangeInDelegatableShares({
            sharesBefore: currentPodOwnerShares,
            sharesAfter: updatedPodOwnerShares
        });

        delegationManager.increaseDelegatedShares({
            staker: podOwner,
            strategy: beaconChainETHStrategy,
            shares: uint256(changeInDelegatableShares)
        });
    }
}
