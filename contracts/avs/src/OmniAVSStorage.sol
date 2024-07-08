// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { IOmniPortal } from "core/interfaces/IOmniPortal.sol";
import { IOmniAVS } from "./interfaces/IOmniAVS.sol";

abstract contract OmniAVSStorage {
    /// @notice Strategy parameters for restaking
    IOmniAVS.StrategyParam[] internal _strategyParams;

    /// @notice Ethereum addresses of currently register operators
    address[] internal _operators;

    /// @notice Map operator address to secp256k1 public key
    mapping(address => bytes) internal _operatorPubkeys;

    /// @notice Set of operators that are allowed to register
    mapping(address => bool) internal _allowlist;

    /// @notice Maximum number of operators that can be registered
    uint32 public maxOperatorCount;

    /// @notice Omni chain id, used to make xcalls to Omni
    uint64 public omniChainId;

    /// @notice OmniPortal.xcall gas limit per each operator in syncWithOmni()
    uint64 public xcallGasLimitPerOperator;

    /// @notice OmniPortal.xcall base gas limit in syncWithOmni()
    uint64 public xcallBaseGasLimit;

    /// @notice Minimum stake required for an operator to register
    uint96 public minOperatorStake;

    /// @notice Whether or not the allowlist is enabled
    bool public allowlistEnabled;

    /// @notice Address of EthStakeInbox contract, predeployed on Omni
    address public ethStakeInbox;

    /// @notice Omni portal contract, used to make xcalls to Omni
    IOmniPortal public omni;
}
