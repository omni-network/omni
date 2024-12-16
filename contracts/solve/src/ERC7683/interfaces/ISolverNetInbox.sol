// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { IOriginSettler } from "./IOriginSettler.sol";
import { ISolverNet } from "./ISolverNet.sol";

interface ISolverNetInbox is IOriginSettler, ISolverNet {
    error NoCalls();
    error NotOutbox();
    error NoSpender();
    error NoDeposits();
    error NotPending();
    error ZeroAmount();
    error ZeroAddress();
    error NotAccepted();
    error NotFulfilled();
    error InvalidChain();
    error WrongCallHash();
    error NoTokenPrereqs();
    error WrongSourceChain();
    error InvalidRecipient();
    error InvalidOrderData();
    error InvalidFillDeadline();
    error InvalidNativeDeposit();
    error NotPendingOrRejected();

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
    event Rejected(bytes32 indexed id, address indexed by, RejectReason indexed reason);

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
     * @param id        ID of the request.
     * @param by        The solver address that claimed the request.
     * @param to        The recipient of claimed deposits.
     * @param deposits  Array of deposits claimed
     */
    event Claimed(bytes32 indexed id, address indexed by, address indexed to, Output[] deposits);

    /**
     * @notice Status of a request.
     */
    enum Status {
        Invalid,
        Pending,
        Accepted,
        Rejected,
        Reverted,
        Fulfilled,
        Claimed
    }

    /**
     * @notice Reason for rejecting a request.
     */
    enum RejectReason {
        None,
        DestCallReverts,
        InsufficientFee,
        InsufficientInventory
    }

    /**
     * @notice Details of a token deposit backing a request.
     * @dev Not stored, only used in opening a request.
     * @param token  Address of the token.
     * @param amount Deposit amount.
     */
    struct TokenDeposit {
        address token;
        uint256 amount;
    }

    /**
     * @notice Order data for a request.
     * @param destChainId  ID of the destination chain.
     * @param deposits     Array of deposits backing the request.
     * @param prereqs      Array of token pre-requisites for the destination calls.
     * @param calls        Array of calls to be executed on the destination chain.
     */
    struct OrderData {
        uint64 destChainId;
        TokenDeposit[] deposits;
        TokenPrereq[] prereqs;
        Call[] calls;
    }

    /**
     * @notice Status update for a request.
     * @param status    Request status.
     * @param timestamp Timestamp of the status update.
     */
    struct StatusUpdate {
        Status status;
        uint40 timestamp;
    }

    /**
     * @notice A request to execute a call on another chain, backed by a deposit.
     * @param order         The order to be executed.
     * @param status        Request status (open, accepted, cancelled, rejected, fulfilled, paid).
     * @param updatedAt     Timestamp request status was last updated.
     * @param acceptedBy    Address of the solver that accepted the request.
     * @param history       Array of status updates including timestamps.
     */
    struct Request {
        ResolvedCrossChainOrder order;
        Status status;
        uint40 updatedAt;
        address acceptedBy;
        StatusUpdate[] history;
    }

    /**
     * @notice Returns the request with the given ID.
     */
    function getRequest(bytes32 id) external view returns (Request memory);

    /**
     * @notice Returns the latest request with the given status.
     */
    function getLatestRequestByStatus(Status status) external view returns (Request memory);

    /**
     * @notice Returns the update history for a request.
     */
    function getUpdateHistory(bytes32 id) external view returns (StatusUpdate[] memory);

    /**
     * @dev Validate the onchain order.
     */
    function validateOnchainOrder(OnchainCrossChainOrder calldata order) external view returns (bool);

    /**
     * @dev Resolve the onchain order.
     */
    function resolve(OnchainCrossChainOrder calldata order) external view returns (ResolvedCrossChainOrder memory);

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
    function reject(bytes32 id, RejectReason reason) external;

    /**
     * @notice Cancel an open or rejected request and refund deposits.
     * @dev Only request initiator can cancel.
     * @param id  ID of the request.
     */
    function cancel(bytes32 id) external;

    /**
     * @notice Fulfill a request.
     * @dev Only callable by the outbox.
     * @param id        ID of the request.
     * @param callHash  Hash of the calls for this request executed on another chain.
     */
    function markFulfilled(bytes32 id, bytes32 callHash) external;

    /**
     * @notice Claim a fulfilled request.
     * @param id  ID of the request.
     * @param to  Address to send deposits to.
     */
    function claim(bytes32 id, address to) external;
}
