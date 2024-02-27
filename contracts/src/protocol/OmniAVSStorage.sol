// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { IAVSDirectory } from "eigenlayer-contracts/src/contracts/interfaces/IAVSDirectory.sol";

import { IDelegationManager } from "../interfaces/IDelegationManager.sol";
import { IOmniPortal } from "../interfaces/IOmniPortal.sol";
import { IOmniAVS } from "../interfaces/IOmniAVS.sol";

abstract contract OmniAVSStorage {
    /// @notice Constant used as a divisor in calculating weights
    uint256 internal constant WEIGHTING_DIVISOR = 1e18;

    /// @notice EigenLayer core DelegationManager
    IDelegationManager internal immutable _delegationManager;

    /// @notice EigenLayer core AVSDirectory
    IAVSDirectory internal immutable _avsDirectory;

    /// @notice Strategy parameters for restaking
    IOmniAVS.StrategyParams[] internal _strategyParams;

    /// @notice List of currently register operators, used to sync EigenCore
    address[] internal _operators;

    /// @notice Set of operators that are allowed to register
    mapping(address => bool) internal _allowlist;

    /// @notice Maximum number of operators that can be registered
    uint32 public maxOperatorCount;

    /// @notice Omni chain id, used to make xcalls to the Omni chain
    uint64 public omniChainId;

    /// @notice OmniPortal.xcall gas limit per each validator in syncWithOmni
    uint64 public xcallGasLimitPerValidator = 10_000;

    /// @notice OmniPortal.xcall base gas limit in syncWithOmnj
    uint64 public xcallBaseGasLimit = 75_000;

    /// @notice Minimum stake required for an operator to register, not including delegations
    uint96 public minOperatorStake;

    /// @notice Omni portal contract, used to make xcalls to the Omni chain
    IOmniPortal public omni;

    constructor(IDelegationManager delegationManager_, IAVSDirectory avsDirectory_) {
        _delegationManager = delegationManager_;
        _avsDirectory = avsDirectory_;
    }
}
