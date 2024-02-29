// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { IOmniPortal } from "../interfaces/IOmniPortal.sol";
import { IOmniAVS } from "../interfaces/IOmniAVS.sol";

abstract contract OmniAVSStorage {
    /// @notice Strategy parameters for restaking
    IOmniAVS.StrategyParam[] internal _strategyParams;

    /// @notice Ethereum addresses of currently register operators
    address[] internal _operators;

    /// @notice Set of operators that are allowed to register
    mapping(address => bool) internal _allowlist;

    /// @notice Omni chain id, used to make xcalls to the Omni chain
    uint64 public omniChainId;

    /// @notice OmniPortal.xcall gas limit per each operator in syncWithOmni()
    uint64 public xcallGasLimitPerOperator = 10_000;

    /// @notice OmniPortal.xcall base gas limit in syncWithOmni()
    uint64 public xcallBaseGasLimit = 75_000;

    /// @notice Omni portal contract, used to make xcalls to the Omni chain
    IOmniPortal public omni;
}
