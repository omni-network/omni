// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { Solve } from "../Solve.sol";

interface ISolveInbox {
    /**
     * @notice Emitted when a request is created.
     * @param id        ID of the request.
     * @param from      Address of the user who created the request.
     * @param call      Details of the call to be executed on another chain.
     * @param deposits  Array of deposits backing the request.
     */
    event Requested(bytes32 indexed id, address indexed from, Solve.Call call, Solve.Deposit[] deposits);

    /**
     * @notice Emitted when a request is accepted.
     * @param id  ID of the request.
     * @param by  Address of the solver who accepted the request.
     */
    event Accepted(bytes32 indexed id, address indexed by);

    /**
     * @notice Emitted when a request is rejected.
     * @param id      ID of the request.
     * @param by      Address of the solver who rejected the request.
     * @param reason  Reason for rejecting the request.
     */
    event Rejected(bytes32 indexed id, address indexed by, Solve.RejectReason indexed reason);

    /**
     * @notice Emitted when a request is cancelled.
     * @param id  ID of the request.
     */
    event Reverted(bytes32 indexed id);

    /**
     * @notice Emitted when a request is fulfilled.
     * @param id          ID of the request.
     * @param callHash    Hash of the call executed on another chain.
     * @param creditedTo  Address of the recipient credited the funds by the solver.
     */
    event Fulfilled(bytes32 indexed id, bytes32 indexed callHash, address indexed creditedTo);

    /**
     * @notice Emitted when a request is claimed.
     * @param id  ID of the request.
     */
    event Claimed(bytes32 indexed id);

    /**
     * /**
     * @notice Returns the request with the given ID.
     */
    function getRequest(bytes32 id) external view returns (Solve.Request memory);

    /**
     * @notice Suggest the amount of native currency to send with a request.
     * @param call        Details of the call to be executed on another chain.
     * @param gasLimit    Maximum gas limit for the call.
     * @param gasPrice    Destination chain gas price in wei.
     * @param fulfillFee  Fee for the fulfill call, retrieved from the destination outbox.
     */
    function suggestNativePayment(Solve.Call calldata call, uint64 gasLimit, uint64 gasPrice, uint256 fulfillFee)
        external
        view
        returns (uint256);

    /**
     * @notice Open a request to execute a call on another chain, backed by deposits.
     *  Token deposits are transferred from msg.sender to this inbox.
     * @param call      Details of the call to be executed on another chain.
     * @param deposits  Array of deposits backing the request.
     */
    function request(Solve.Call calldata call, Solve.TokenDeposit[] calldata deposits)
        external
        payable
        returns (bytes32 id);

    /**
     * @notice Accept an open request.
     * @dev Only a whitelisted solver can accept.
     * @param id  ID of the request.
     */
    function accept(bytes32 id) external;

    /**
     * @notice Reject an open request.
     * @dev Only a whitelisted solver can reject.
     * @param id  ID of the request.
     */
    function reject(bytes32 id, Solve.RejectReason reason) external;

    /**
     * @notice Cancel an open or rejected request and refund deposits.
     * @dev Only request initiator can cancel.
     * @param id  ID of the request.
     */
    function cancel(bytes32 id) external;

    /**
     * @notice Fulfill a request.
     * @dev Only callable by the outbox.
     */
    function markFulfilled(bytes32 id, bytes32 callHash, address creditTo) external;

    /**
     * @notice Claim a fulfilled request.
     * @param id  ID of the request.
     */
    function claim(bytes32 id) external;
}
