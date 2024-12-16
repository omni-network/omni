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

    /**
     * @notice Returns whether a call is allowed.
     * @param target    Address of the target contract.
     * @param selector  4-byte selector of the function to allow.
     * @return          Whether the call is allowed.
     */
    function allowedCalls(address target, bytes4 selector) external view returns (bool);

    /**
     * @notice Returns whether a call is fulfilled.
     * @param callHash  Hash of the call executed.
     * @return          Whether the call is fulfilled.
     */
    function fulfilledCalls(bytes32 callHash) external view returns (bool);

    /**
     * @notice Returns the fee for a fulfill call.
     * @param srcChainId  ID of the source chain.
     * @return            Fee for the fulfill call.
     */
    function fulfillFee(uint64 srcChainId) external view returns (uint256);

    /**
     * @notice Returns whether a call has been fulfilled.
     * @param srcReqId          ID of the on the source inbox.
     * @param srcChainId        ID of the source chain.
     * @param fillOriginData    Data emitted on the origin to parameterize the fill
     * @return                  Whether the call has been fulfilled.
     */
    function didFulfill(bytes32 srcReqId, uint64 srcChainId, bytes calldata fillOriginData)
        external
        view
        returns (bool);

    /**
     * @notice Sets an allowed call.
     * @param target    Address of the target contract.
     * @param selector  4-byte selector of the function to allow.
     * @param allowed   Whether the call is allowed.
     */
    function setAllowedCall(address target, bytes4 selector, bool allowed) external;
}
