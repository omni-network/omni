// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.24;

import { IERC7683 } from "./IERC7683.sol";

/// @title IOriginSettler
/// @notice Standard interface for settlement contracts on the origin chain
/// @dev See https://github.com/ethereum/ERCs/blob/master/ERCS/erc-7683.md
interface IOriginSettler is IERC7683 {
    /// @notice Opens a cross-chain order
    /// @dev To be called by the user
    /// @dev This method must emit the Open event
    /// @param order The OnchainCrossChainOrder definition
    function open(OnchainCrossChainOrder calldata order) external payable;

    /// @notice Resolves a specific OnchainCrossChainOrder into a generic ResolvedCrossChainOrder
    /// @dev Intended to improve standardized integration of various order types and settlement contracts
    /// @param order The OnchainCrossChainOrder definition
    /// @return ResolvedCrossChainOrder hydrated order data including the inputs and outputs of the order
    function resolve(OnchainCrossChainOrder calldata order) external view returns (ResolvedCrossChainOrder memory);
}
