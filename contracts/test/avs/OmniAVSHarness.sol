// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { OmniAVS } from "src/protocol/OmniAVS.sol";
import { IDelegationManager } from "eigenlayer-contracts/src/contracts/interfaces/IDelegationManager.sol";
import { IAVSDirectory } from "eigenlayer-contracts/src/contracts/interfaces/IAVSDirectory.sol";

/**
 * @title OmniAVSHarness
 * @dev Wrapper around OmniAVS that exposes internal functions.
 */
contract OmniAVSHarness is OmniAVS {
    constructor(IDelegationManager delegationManager, IAVSDirectory avsDirectory)
        OmniAVS(delegationManager, avsDirectory)
    { }

    function xcallGasLimitFor(uint256 numValidators) external view returns (uint64) {
        return _xcallGasLimitFor(numValidators);
    }
}
