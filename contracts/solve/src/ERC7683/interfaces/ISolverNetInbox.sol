// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { IOriginSettler } from "./IOriginSettler.sol";
import { ISolverNet } from "./ISolverNet.sol";

interface ISolverNetInbox is IOriginSettler, ISolverNet {
    error NoCalls();
    error NotOutbox();
    error NoSpender();
    error NoCalldata();
    error NoDeposits();
    error NotPending();
    error ZeroAmount();
    error ZeroAddress();
    error NotAccepted();
    error NotFulfilled();
    error WrongCallHash();
    error NoTokenPrereqs();
    error InvalidSrcChain();
    error InvalidDestChain();
    error WrongSourceChain();
    error InvalidRecipient();
    error InvalidOrderData();
    error InvalidFillDeadline();
    error InvalidNativeDeposit();
    error NotPendingOrRejected();
    error InvalidOrderDataTypehash();

    /**
     * @notice Emitted when an order is accepted.
     * @param id  ID of the order.
     * @param by  Address of the solver who accepted the order.
     */
    event Accepted(bytes32 indexed id, address indexed by);

    /**
     * @notice Emitted when an order is rejected.
     * @param id      ID of the order.
     * @param by      Address of the solver who rejected the order.
     * @param reason  Reason code for rejecting the order.
     */
    event Rejected(bytes32 indexed id, address indexed by, uint8 indexed reason);

    /**
     * @notice Emitted when an order is cancelled.
     * @param id  ID of the order.
     */
    event Reverted(bytes32 indexed id);

    /**
     * @notice Emitted when an order is fulfilled.
     * @param id          ID of the order.
     * @param callHash    Hash of the call executed on another chain.
     * @param creditedTo  Address of the recipient credited the funds by the solver.
     */
    event Fulfilled(bytes32 indexed id, bytes32 indexed callHash, address indexed creditedTo);

    /**
     * @notice Emitted when an order is claimed.
     * @param id        ID of the order.
     * @param by        The solver address that claimed the order.
     * @param to        The recipient of claimed deposits.
     * @param deposits  Array of deposits claimed
     */
    event Claimed(bytes32 indexed id, address indexed by, address indexed to, Output[] deposits);

    /**
     * @notice Status of an order.
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
     * @notice Details of a token deposit backing an order.
     * @dev Not stored, only used in opening an order.
     * @param token  Address of the token.
     * @param amount Deposit amount.
     */
    struct TokenDeposit {
        address token;
        uint256 amount;
    }

    /**
     * @notice Data for a cross-chain order.
     * @param intent    Intent for the order, contains chain IDs, token pre-requisites, and the call to be executed.
     * @param deposits  Array of deposits backing the order.
     */
    struct SolverNetOrderData {
        SolverNetIntent intent;
        TokenDeposit[] deposits;
    }

    /**
     * @notice Status update for an order.
     * @param status    Order status.
     * @param timestamp Timestamp of the status update.
     */
    struct StatusUpdate {
        Status status;
        uint40 timestamp;
    }

    /**
     * @notice State of an order.
     * @param status      Order status.
     * @param acceptedBy  Address of the solver that accepted the order.
     */
    struct OrderState {
        Status status;
        address acceptedBy;
    }

    /**
     * @notice Returns resolved cross-chain order with current state and history.
     * @param id  Order ID.j:w
     */
    function getOrder(bytes32 id)
        external
        view
        returns (ResolvedCrossChainOrder memory order, OrderState memory state, StatusUpdate[] memory history);

    /**
     * @notice Returns the latest order ID with the given status.
     */
    function getLatestOrderIdByStatus(Status status) external view returns (bytes32);

    /**
     * @dev Validate the onchain order.
     */
    function validateOrder(OnchainCrossChainOrder calldata order) external view returns (bool);

    /**
     * @notice Accept an open order.
     * @dev Only a whitelisted solver can accept.
     * @param id  ID of the order.
     */
    function accept(bytes32 id) external;

    /**
     * @notice Reject an open order.
     * @dev Only a whitelisted solver can reject.
     * @param id      ID of the order.
     * @param reason  Reason code for rejecting the order.
     */
    function reject(bytes32 id, uint8 reason) external;

    /**
     * @notice Cancel an open or rejected order and refund deposits.
     * @dev Only order initiator can cancel.
     * @param id  ID of the order.
     */
    function cancel(bytes32 id) external;

    /**
     * @notice Fulfill an order.
     * @dev Only callable by the outbox.
     * @param id        ID of the order.
     * @param callHash  Hash of the calls for this order executed on another chain.
     */
    function markFulfilled(bytes32 id, bytes32 callHash) external;

    /**
     * @notice Claim a fulfilled order.
     * @param id  ID of the order.
     * @param to  Address to send deposits to.
     */
    function claim(bytes32 id, address to) external;
}
