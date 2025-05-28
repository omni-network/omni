// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.24;

import { IOriginSettler } from "../erc7683/IOriginSettler.sol";
import { IMessageRecipient } from "@hyperlane-xyz/core/contracts/interfaces/IMessageRecipient.sol";
import { SolverNet } from "../lib/SolverNet.sol";

interface ISolverNetInboxV2 is IOriginSettler, IMessageRecipient {
    // Validation errors
    error InvalidOrderTypehash();
    error InvalidOrderData();
    error InvalidOriginSettler();
    error InvalidOriginChainId();
    error InvalidDestinationChainId();
    error InvalidOpenDeadline();
    error InvalidFillDeadline();
    error InvalidMissingCalls();
    error InvalidExpenseToken();
    error InvalidExpenseAmount();
    error InvalidArrayLength();
    error InvalidUser();

    // Open order errors
    error InvalidERC20Deposit();
    error InvalidNativeDeposit();

    // Reject order errors
    error InvalidReason();

    // Order status errors
    error OrderNotPending();
    error OrderStillValid();

    // Order fill errors
    error WrongSourceChain();
    error WrongFillHash();

    // Order claim errors
    error OrderNotFilled();

    // Pause errors
    error IsPaused();
    error AllPaused();

    /**
     * @notice Emitted when an outbox is set.
     * @param chainId ID of the chain.
     * @param outbox  Address of the outbox.
     */
    event OutboxSet(uint64 indexed chainId, address indexed outbox);

    /**
     * @notice Emitted when a pause state is set.
     * @param key Pause key.
     * @param pause True if paused, false if unpaused.
     * @param pauseState Current pause state.
     */
    event Paused(bytes32 indexed key, bool indexed pause, uint8 indexed pauseState);

    /**
     * @notice Emitted when an order is opened.
     * @dev This event emits the FillOriginData typed `originData`, rather than ABI-encoded as seen in `IERC7683.Open`.
     * @param id ID of the order.
     * @param fillOriginData Order fill originData.
     */
    event FillOriginData(bytes32 indexed id, SolverNet.FillOriginData fillOriginData);

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
    event Closed(bytes32 indexed id);

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
        Rejected,
        Closed,
        Filled,
        Claimed
    }

    /**
     * @notice State of an order.
     * @param status       Latest order status.
     * @param rejectReason Reason code for rejecting the order, if rejected.
     * @param timestamp    Timestamp of the status update.
     * @param updatedBy    Address for who last updated the order.
     */
    struct OrderState {
        Status status;
        uint8 rejectReason;
        uint32 timestamp;
        address updatedBy;
    }

    /**
     * @notice Set the outbox addresses for the given chain IDs.
     * @param chainIds IDs of the chains.
     * @param outboxes Addresses of the outboxes.
     */
    function setOutboxes(uint64[] calldata chainIds, address[] calldata outboxes) external;

    /**
     * @notice Returns the outbox address for the given chain ID.
     * @param chainId ID of the chain.
     * @return outbox Outbox address.
     */
    function getOutbox(uint64 chainId) external view returns (address);

    /**
     * @notice Returns the order, its state, and offset with the given ID.
     * @param id ID of the order.
     */
    function getOrder(bytes32 id)
        external
        view
        returns (ResolvedCrossChainOrder memory order, OrderState memory state, uint248 offset);

    /**
     * @notice Returns the order ID for the given user and nonce.
     * @param user  Address of the user.
     * @param nonce Nonce of the order.
     * @param gasless Whether the order is gasless.
     */
    function getOrderId(address user, uint256 nonce, bool gasless) external view returns (bytes32);

    /**
     * @notice Returns the next onchain order ID for the given user.
     * @param user Address of the user the order is opened for.
     */
    function getNextOnchainOrderId(address user) external view returns (bytes32);

    /**
     * @notice Returns the onchain nonce for the given user.
     * @param user Address of the user the order is opened for.
     */
    function getOnchainUserNonce(address user) external view returns (uint256);

    /**
     * @notice Returns the order offset of the latest order opened at this inbox.
     */
    function getLatestOrderOffset() external view returns (uint248);

    /**
     * @dev Validate the onchain order.
     * @param order OnchainCrossChainOrder to validate.
     */
    function validate(OnchainCrossChainOrder calldata order) external view returns (bool);

    /**
     * @dev Validate the gasless order.
     * @param order GaslessCrossChainOrder to validate.
     */
    function validateFor(GaslessCrossChainOrder calldata order) external view returns (bool);

    /**
     * @notice Reject an open order and refund deposits.
     * @dev Only a whitelisted solver can reject.
     * @param id     ID of the order.
     * @param reason Reason code for rejection.
     */
    function reject(bytes32 id, uint8 reason) external;

    /**
     * @notice Close order and refund deposits after fill deadline has elapsed.
     * @dev Only order initiator can close.
     * @param id ID of the order.
     */
    function close(bytes32 id) external;

    /**
     * @notice Fill an order via Omni Core.
     * @dev Only callable by the outbox.
     * @param id         ID of the order.
     * @param fillHash   Hash of fill instructions origin data.
     * @param creditedTo Address deposits are credited to, provided by the filler.
     */
    function markFilled(bytes32 id, bytes32 fillHash, address creditedTo) external;

    /**
     * @notice Claim deposits for a filled order.
     * @param id ID of the order.
     * @param to Address to send deposits to.
     */
    function claim(bytes32 id, address to) external;
}
