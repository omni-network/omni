// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { IDestinationSettler } from "./IDestinationSettler.sol";
import { ISolverNet } from "./ISolverNet.sol";

interface ISolverNetOutbox is IDestinationSettler, ISolverNet {
    error CallFailed();
    error WrongDestChain();
    error CallNotAllowed();
    error InsufficientFee();
    error AlreadyFulfilled();
    error IncorrectPrereqs();

    /**
     * @notice Emitted when a call is allowed.
     * @param target    Address of the target contract.
     * @param selector  4-byte selector of the function to allow.
     * @param allowed   Whether the call is allowed.
     */
    event AllowedCallSet(address indexed target, bytes4 indexed selector, bool allowed);

    /**
     * @notice Emitted when a request is fulfilled.
     * @param orderId     ID of the order.
     * @param callHash    Hash of the call executed.
     * @param solvedBy    Address of the solver.
     */
    event Fulfilled(bytes32 indexed orderId, bytes32 indexed callHash, address indexed solvedBy);
}
