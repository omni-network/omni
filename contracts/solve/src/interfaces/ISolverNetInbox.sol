// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { IOriginSettler } from "../erc7683/IOriginSettler.sol";

interface ISolverNetInbox is IOriginSettler {
    // Validation errors
    error InvalidOrderTypehash();
    error InvalidOrderData();
    error InvalidChainId();
    error InvalidFillDeadline();
    error InvalidMissingCalls();
    error InvalidCallTarget();
    error InvalidExpenseToken();
    error InvalidExpenseAmount();

    // Open order errors
    error InvalidNativeDeposit();

    // Order accept/reject/cancel errors
    error OrderNotPending();
    error FillDeadlinePassed();

    // Order fill errors
    error OrderNotPendingOrAccepted();
    error WrongSourceChain();
    error WrongFillHash();

    // Order claim errors
    error OrderNotFilled();

    /**
     * @notice Emitted when an outbox is set.
     * @param chainId ID of the chain.
     * @param outbox  Address of the outbox.
     */
    event OutboxSet(uint64 indexed chainId, address indexed outbox);

    /**
     * @notice Emitted when an order is accepted.
     * @param id ID of the order.
     * @param by Address of the solver who accepted the order.
     */
    event Accepted(bytes32 indexed id, address indexed by);

    /**
     * @notice Emitted when an order is rejected.
     * @param id     ID of the order.
     * @param by     Address of the solver who rejected the order.
     * @param reason Reason code for rejecting the order.
     */
    event Rejected(bytes32 indexed id, address indexed by, uint8 indexed reason);

    /**
     * @notice Emitted when an order is cancelled.
     * @param id ID of the order.
     */
    event Reverted(bytes32 indexed id);

    /**
     * @notice Emitted when an order is filled.
     * @param id         ID of the order.
     * @param fillHash   Hash of the fill instructions origin data.
     * @param creditedTo Address of the recipient credited the funds by the solver.
     */
    event Filled(bytes32 indexed id, bytes32 indexed fillHash, address indexed creditedTo);

    /**
     * @notice Emitted when an order is claimed.
     * @param id ID of the order.
     * @param by The solver address that claimed the order.
     * @param to The recipient of claimed deposits.
     */
    event Claimed(bytes32 indexed id, address indexed by, address indexed to);

    /**
     * @notice Status of an order.
     */
    enum Status {
        Invalid,
        Pending,
        Accepted,
        Rejected,
        Reverted,
        Filled,
        Claimed
    }

    /**
     * @notice State of an order.
     * @param status    Latest order status.
     * @param timestamp Timestamp of the status update.
     * @param claimant  Address of the claimant, defined at fill.
     */
    struct OrderState {
        Status status;
        uint32 timestamp;
        address claimant;
    }

    /**
     * @notice Set the outbox addresses for the given chain IDs.
     * @param chainIds IDs of the chains.
     * @param outboxes Addresses of the outboxes.
     */
    function setOutboxes(uint64[] calldata chainIds, address[] calldata outboxes) external;

    /**
     * @notice Returns the order and its state with the given ID.
     * @param id ID of the order.
     */
    function getOrder(bytes32 id)
        external
        view
        returns (ResolvedCrossChainOrder memory order, OrderState memory state);

    /**
     * @notice Returns the next order ID.
     */
    function getNextId() external view returns (bytes32);

    /**
     * @notice Returns the latest order with the given status.
     * @param status Order status to query.
     */
    function getLatestOrderIdByStatus(Status status) external view returns (bytes32);

    /**
     * @dev Validate the onchain order.
     * @param order OnchainCrossChainOrder to validate.
     */
    function validate(OnchainCrossChainOrder calldata order) external view returns (bool);

    /**
     * @notice Accept an open order.
     * @dev Only a whitelisted solver can accept.
     * @param id ID of the order.
     */
    function accept(bytes32 id) external;

    /**
     * @notice Reject an open order and refund deposits.
     * @dev Only a whitelisted solver can reject.
     * @param id     ID of the order.
     * @param reason Reason code for rejection.
     */
    function reject(bytes32 id, uint8 reason) external;

    /**
     * @notice Cancel an open and refund deposits.
     * @dev Only order initiator can cancel.
     * @param id ID of the order.
     */
    function cancel(bytes32 id) external;

    /**
     * @notice Fill an order.
     * @dev Only callable by the outbox.
     * @param id        ID of the order.
     * @param fillHash  Hash of fill instructions origin data.
     * @param claimant  Address to claim the order, provided by the filler.
     */
    function markFilled(bytes32 id, bytes32 fillHash, address claimant) external;

    /**
     * @notice Claim a filled order.
     * @param id ID of the order.
     * @param to Address to send deposits to.
     */
    function claim(bytes32 id, address to) external;
}
