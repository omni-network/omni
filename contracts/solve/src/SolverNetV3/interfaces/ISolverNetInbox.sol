// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { IOriginSettler } from "./IOriginSettler.sol";
import { ISolverNet } from "./ISolverNet.sol";

interface ISolverNetInbox is IOriginSettler, ISolverNet {
    // markFilled authorization
    error NotOutbox();

    // OnchainCrossChainOrder validation errors
    error InvalidOrderData();
    error InvalidOrderDataTypehash();
    error InvalidFillDeadline();
    error InvalidExpenseNativeAmount(uint256 callNativeTokens, uint256 expenseNativeTokens);

    // deposit validation errors
    error NoDeposits();
    error NoDepositAmount();
    error DuplicateNativeDeposit();
    error InvalidNativeDeposit();

    // call validation errors
    error NoCalls();
    error NoCallData();
    error NoCallTarget();
    error NoCallChainId();

    // expense validation errors
    error NoExpenseAmount();
    error NoExpenseChainId();

    // order state transition errors
    error OrderNotPending();
    error OrderNotAccepted();
    error OrderNotFilled();

    // transfer errors
    error InvalidRecipient();

    // outbox config errors
    error InvalidOutbox();
    error InvalidChainId();
    error ArrayLengthMismatch();

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
     * @notice Emitted when an order is filled.
     * @param id          ID of the order.
     * @param callHash    Hash of the call executed on another chain.
     * @param creditedTo  Address of the recipient credited the funds by the solver.
     */
    event Filled(bytes32 indexed id, bytes32 indexed callHash, address indexed creditedTo);

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
        Filled,
        Claimed
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
     * @param latest      Latest order status.
     * @param accepted    Accepted order status.
     * @param acceptedBy  Address of the solver that accepted the order.
     */
    struct OrderState {
        StatusUpdate latest;
        StatusUpdate accepted;
        address acceptedBy;
    }

    /**
     * @notice Returns the next order ID.
     */
    function getNextId() external view returns (bytes32);

    /**
     * @notice Returns the outbox address for a given destination chain ID.
     * @param destChainId Destination chain ID.
     */
    function getOutbox(uint64 destChainId) external view returns (bytes32);

    /**
     * @notice Returns the order state for a given resolved order.
     * @param resolvedOrder Resolved order to query.
     */
    function getOrderState(ResolvedCrossChainOrder calldata resolvedOrder) external view returns (OrderState memory);

    /**
     * @notice Returns the latest order ID with the given status.
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
     * @param resolvedOrder Resolved order to accept.
     */
    function accept(ResolvedCrossChainOrder calldata resolvedOrder) external;

    /**
     * @notice Reject an open order and refund deposits.
     * @dev Only a whitelisted solver can reject.
     * @param resolvedOrder Resolved order to reject.
     * @param reason        Reason code for rejection.
     */
    function reject(ResolvedCrossChainOrder calldata resolvedOrder, uint8 reason) external;

    /**
     * @notice Cancel an open and refund deposits.
     * @dev Only the resolved order's user can cancel.
     * @param resolvedOrder Resolved order to cancel.
     */
    function cancel(ResolvedCrossChainOrder calldata resolvedOrder) external;

    /**
     * @notice Mark an order as filled.
     * @dev Only callable by the crosschain outbox. `acceptedBy` is only utilized if a solver filled but didn't accept.
     * @param orderId    ID of the order.
     * @param orderHash  Hash of the resolved order.
     * @param timestamp  Timestamp of the fill.
     * @param acceptedBy Address of a solver that filled but didn't accept.
     */
    function markFilled(bytes32 orderId, bytes32 orderHash, uint40 timestamp, address acceptedBy) external;

    /**
     * @notice Claim a filled order.
     * @dev Only the solver address set to `acceptedBy` can claim.
     * @param resolvedOrder Fulfilled resolved order to claim deposits for.
     * @param to            Address to send deposits to.
     */
    function claim(ResolvedCrossChainOrder memory resolvedOrder, address to) external;

    /**
     * @notice Set the outboxes for a given chain.
     * @dev Only callable by the owner.
     * @param chainIds Chain IDs to set outboxes for.
     * @param outboxes Outbox addresses to set.
     */
    function setOutboxes(uint64[] memory chainIds, bytes32[] memory outboxes) external;
}
