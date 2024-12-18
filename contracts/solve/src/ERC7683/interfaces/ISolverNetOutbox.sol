// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { IDestinationSettler } from "./IDestinationSettler.sol";
import { ISolverNet } from "./ISolverNet.sol";

interface ISolverNetOutbox is IDestinationSettler, ISolverNet {
    error CallFailed();
    error InvalidPrereq();
    error WrongDestChain();
    error CallNotAllowed();
    error InsufficientFee();
    error AlreadyFulfilled();

    /**
     * @notice Emitted when a call is allowed.
     * @param target    Address of the target contract.
     * @param selector  4-byte selector of the function to allow.
     * @param allowed   Whether the call is allowed.
     */
    event AllowedCallSet(address indexed target, bytes4 indexed selector, bool allowed);

    /**
     * @notice Emitted when a cross-chain request is fulfilled on the destination chain
     * @param orderId     ID of the order on the source chain
     * @param callHash    Hash of the executed call and its parameters
     * @param solvedBy    Address of the solver that executed the fulfillment
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
     * @notice Returns the message passing fee required to mark a request as fulfilled on the source chain
     * @param srcChainId  ID of the source chain.
     * @return            Fee amount in native currency.
     */
    function fulfillFee(uint64 srcChainId) external view returns (uint256);

    /**
     * @notice Returns whether a call has been fulfilled.
     * @param srcReqId    ID of the on the source inbox.
     * @param originData  Data emitted on the origin to parameterize the fill
     * @return            Whether the call has been fulfilled.
     */
    function didFulfill(bytes32 srcReqId, bytes calldata originData) external view returns (bool);

    /**
     * @notice Sets an allowed call.
     * @param target    Address of the target contract.
     * @param selector  4-byte selector of the function to allow.
     * @param allowed   Whether the call is allowed.
     */
    function setAllowedCall(address target, bytes4 selector, bool allowed) external;
}
