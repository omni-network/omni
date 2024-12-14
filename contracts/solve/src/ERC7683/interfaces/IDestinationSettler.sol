// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { IERC7683 } from "./IERC7683.sol";

/// @title IDestinationSettler
/// @notice Standard interface for settlement contracts on the destination chain
/// @dev See https://github.com/ethereum/ERCs/blob/master/ERCS/erc-7683.md
interface IDestinationSettler is IERC7683 {
    /// @notice Fills a single leg of a particular order on the destination chain
    /// @param orderId Unique order identifier for this order
    /// @param originData Data emitted on the origin to parameterize the fill
    /// @param fillerData Data provided by the filler to inform the fill or express their preferences
    function fill(bytes32 orderId, bytes calldata originData, bytes calldata fillerData) external;
}
