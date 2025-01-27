// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OwnableRoles } from "solady/src/auth/OwnableRoles.sol";
import { ReentrancyGuard } from "solady/src/utils/ReentrancyGuard.sol";
import { Initializable } from "solady/src/utils/Initializable.sol";
import { XAppBase } from "core/src/pkg/XAppBase.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { DeployedAt } from "src/util/DeployedAt.sol";
import { AddrUtils } from "src/ERC7683/lib/AddrUtils.sol";
import { ISolverNetInbox } from "./interfaces/ISolverNetInbox.sol";

/**
 * @title SolverNetInbox
 * @notice Entrypoint and alt-mempool for user solve orders.
 */
contract SolverNetInbox is OwnableRoles, ReentrancyGuard, Initializable, DeployedAt, XAppBase, ISolverNetInbox {
    using SafeTransferLib for address;
    using AddrUtils for address;
    using AddrUtils for bytes32;

    /**
     * @notice Role for solvers.
     * @dev _ROLE_0 evaluates to '1'.
     */
    uint256 internal constant SOLVER = _ROLE_0;

    /**
     * @notice Typehash for the order data.
     */
    bytes32 internal constant ORDER_DATA_TYPEHASH = keccak256(
        "OrderData(address user,Call call,Deposit[] deposits,TokenExpense[] expenses)Call(uint64 chainId,bytes32 target,uint256 value,bytes data)Deposit(bytes32 token,uint256 amount)TokenExpense(bytes32 token,bytes32 spender,uint256 amount)"
    ); // Not really needed until we support more than one order type or gasless orders

    /**
     * @dev Counter for generating unique order IDs. Incremented each time a new order is created.
     */
    uint256 internal _lastId;

    /**
     * @notice Outbox contract addresses per destChainId.
     */
    mapping(uint64 destChainId => bytes32) internal _outboxes;

    /**
     * @notice Map order ID to order parameters.
     */
    mapping(bytes32 orderHash => OrderState) internal _orderState;

    /**
     * @notice Map status to latest order ID.
     */
    mapping(Status => bytes32 id) internal _latestOrderIdByStatus;

    constructor() {
        _disableInitializers();
    }

    /**
     * @notice Initialize the contract's owner and solver.
     * @dev Used instead of constructor as we want to use the transparent upgradeable proxy pattern.
     * @param owner_  Address of the owner.
     * @param solver_ Address of the solver.
     * @param omni_   Address of the OmniPortal contract.
     */
    function initialize(address owner_, address solver_, address omni_) external initializer {
        _initializeOwner(owner_);
        _grantRoles(solver_, SOLVER);
        _setOmniPortal(omni_);
    }

    /**
     * @notice Returns the next order ID.
     */
    function getNextId() external view returns (bytes32) {
        return _nextId();
    }

    /**
     * @notice Returns the outbox address for a given destination chain ID.
     * @param destChainId Destination chain ID.
     */
    function getOutbox(uint64 destChainId) external view returns (bytes32) {
        return _outboxes[destChainId];
    }

    /**
     * @notice Returns the order state for a given resolved order.
     * @param resolvedOrder Resolved order to query.
     */
    function getOrderState(ResolvedCrossChainOrder calldata resolvedOrder) external view returns (OrderState memory) {
        return (_orderState[resolvedOrder.orderId]);
    }

    /**
     * @notice Returns the latest order with the given status.
     * @param status Order status to query.
     */
    function getLatestOrderIdByStatus(Status status) external view returns (bytes32) {
        return _latestOrderIdByStatus[status];
    }

    /**
     * @dev Validate the onchain order.
     * @param order OnchainCrossChainOrder to validate.
     */
    function validate(OnchainCrossChainOrder calldata order) external view returns (bool) {
        _validate(order);
        return true;
    }

    /**
     * @notice Resolve the onchain order.
     * @param order  OnchainCrossChainOrder to resolve.
     */
    function resolve(OnchainCrossChainOrder calldata order) public view returns (ResolvedCrossChainOrder memory) {
        OrderData memory orderData = _validate(order);
        Call[] memory calls = orderData.calls;
        Deposit[] memory deposits = orderData.deposits;
        TokenExpense[] memory expenses = orderData.expenses;
        bytes32 orderId = _nextId();

        // Process `calls` into `fillInstructions`.
        FillInstruction[] memory fillInstructions = new FillInstruction[](calls.length);
        for (uint256 i; i < calls.length; ++i) {
            Call memory call = calls[i];
            fillInstructions[i] = FillInstruction({
                destinationChainId: call.chainId,
                destinationSettler: _outboxes[call.chainId],
                originData: abi.encode(OriginData({ srcChainId: block.chainid, orderId: orderId, call: call }))
            });
        }

        // Process `deposits` into `minReceived`.
        Output[] memory minReceived = new Output[](deposits.length);
        for (uint256 i; i < deposits.length; ++i) {
            Deposit memory deposit = deposits[i];
            minReceived[i] = Output({
                token: deposit.token,
                amount: deposit.amount,
                recipient: bytes32(0), // recipient is solver, which is only known on acceptance
                chainId: block.chainid
            });
        }

        // Process `expenses` into `maxSpent`.
        Output[] memory maxSpent = new Output[](expenses.length);
        for (uint256 i; i < expenses.length; ++i) {
            TokenExpense memory expense = expenses[i];
            maxSpent[i] = Output({
                token: expense.token,
                amount: expense.amount,
                recipient: expense.spender, // We utilize this field to specify what address needs a token approval
                chainId: expense.chainId
            });
        }

        return ResolvedCrossChainOrder({
            user: orderData.user != address(0) ? orderData.user : msg.sender,
            originChainId: block.chainid,
            openDeadline: uint32(block.timestamp),
            fillDeadline: order.fillDeadline,
            orderId: orderId, // use value from `nextId()` for view function here, may not be id when opened
            maxSpent: maxSpent,
            minReceived: minReceived,
            fillInstructions: fillInstructions
        });
    }

    /**
     * @notice Open an order to execute a call on another chain, backed by deposits.
     * @dev Token deposits are transferred from msg.sender to this inbox.
     * @param order OnchainCrossChainOrder to open.
     */
    function open(OnchainCrossChainOrder calldata order) external payable nonReentrant {
        OrderData memory orderData = abi.decode(order.orderData, (OrderData));
        ResolvedCrossChainOrder memory resolvedOrder = resolve(order);
        _processDeposits(orderData.deposits);
        _incrementId();
        emit Open(resolvedOrder.orderId, resolvedOrder);
    }

    /**
     * @notice Accept an open order.
     * @dev Only a whitelisted solver can accept.
     * @param resolvedOrder Resolved order to accept.
     */
    function accept(ResolvedCrossChainOrder calldata resolvedOrder) external onlyRoles(SOLVER) nonReentrant {
        bytes32 orderHash = _orderHash(resolvedOrder);
        OrderState memory state = _orderState[orderHash];

        if (state.latest.status != Status.Pending) revert OrderNotPending();

        _updateOrderState(resolvedOrder.orderId, orderHash, Status.Accepted, msg.sender);

        emit Accepted(resolvedOrder.orderId, msg.sender);
    }

    /**
     * @notice Reject an open order and refund deposits.
     * @dev Only a whitelisted solver can reject.
     * @param resolvedOrder Resolved order to reject.
     * @param reason        Reason code for rejection.
     */
    function reject(ResolvedCrossChainOrder calldata resolvedOrder, uint8 reason)
        external
        onlyRoles(SOLVER)
        nonReentrant
    {
        bytes32 orderHash = _orderHash(resolvedOrder);
        OrderState memory state = _orderState[orderHash];

        if (state.latest.status != Status.Pending) revert OrderNotPending();

        _updateOrderState(resolvedOrder.orderId, orderHash, Status.Rejected, address(0));
        _transferDeposits(resolvedOrder.user, resolvedOrder.minReceived);

        emit Rejected(resolvedOrder.orderId, msg.sender, reason);
    }

    /**
     * @notice Cancel an open and refund deposits.
     * @dev Only the resolved order's user can cancel.
     * @param resolvedOrder Resolved order to cancel.
     */
    function cancel(ResolvedCrossChainOrder calldata resolvedOrder) external nonReentrant {
        bytes32 orderHash = _orderHash(resolvedOrder);
        OrderState memory state = _orderState[orderHash];

        if (state.latest.status != Status.Pending) revert OrderNotPending();
        if (resolvedOrder.user != msg.sender) revert Unauthorized();

        _updateOrderState(resolvedOrder.orderId, orderHash, Status.Reverted, address(0));
        _transferDeposits(resolvedOrder.user, resolvedOrder.minReceived);

        emit Reverted(resolvedOrder.orderId);
    }

    /**
     * @notice Mark an order as filled.
     * @dev Only callable by the crosschain outbox. `acceptedBy` is only utilized if a solver filled but didn't accept.
     * @param orderId    ID of the order.
     * @param orderHash  Hash of the resolved order.
     * @param timestamp  Timestamp of the fill.
     * @param acceptedBy Address of a solver that filled but didn't accept.
     */
    function markFilled(bytes32 orderId, bytes32 orderHash, uint40 timestamp, address acceptedBy)
        external
        xrecv
        nonReentrant
    {
        OrderState memory state = _orderState[orderHash];

        if (state.latest.status != Status.Accepted) revert OrderNotAccepted();
        if (xmsg.sender != _outboxes[xmsg.sourceChainId].toAddress()) revert NotOutbox();

        // If the order was fulfilled without being accepted, apply the filled status and set the `acceptedBy` address.
        if (state.accepted.status != Status.Accepted) {
            _updateOrderState(orderId, orderHash, Status.Filled, acceptedBy);
            emit Filled(orderId, orderHash, acceptedBy);
            return;
        }

        // If the order was fulfilled before being accepted, override the `acceptedBy` address. Otherwise apply the filled status.
        address raceWinner = state.accepted.timestamp > timestamp ? acceptedBy : state.acceptedBy;
        if (raceWinner != state.acceptedBy) _updateOrderState(orderId, orderHash, Status.Filled, raceWinner);
        else _updateOrderState(orderId, orderHash, Status.Filled, address(0));

        emit Filled(orderId, orderHash, raceWinner);
    }

    /**
     * @notice Claim a filled order.
     * @dev Only the solver address set to `acceptedBy` can claim.
     * @param resolvedOrder Fulfilled resolved order to claim deposits for.
     * @param to            Address to send deposits to.
     */
    function claim(ResolvedCrossChainOrder memory resolvedOrder, address to) external nonReentrant {
        bytes32 orderHash = _orderHash(resolvedOrder);
        OrderState memory state = _orderState[orderHash];

        if (state.latest.status != Status.Filled) revert OrderNotFilled();
        if (state.acceptedBy != msg.sender) revert Unauthorized();

        _updateOrderState(resolvedOrder.orderId, orderHash, Status.Claimed, address(0));
        _transferDeposits(to, resolvedOrder.minReceived);

        emit Claimed(resolvedOrder.orderId, msg.sender, to, resolvedOrder.minReceived);
    }

    /**
     * @notice Set the outboxes for a given chain.
     * @dev Only callable by the owner.
     * @param chainIds Chain IDs to set outboxes for.
     * @param outboxes Outbox addresses to set.
     */
    function setOutboxes(uint64[] memory chainIds, bytes32[] memory outboxes) external onlyOwner {
        if (chainIds.length != outboxes.length) revert ArrayLengthMismatch();
        for (uint256 i; i < chainIds.length; ++i) {
            if (chainIds[i] == block.chainid) revert InvalidChainId();
            if (outboxes[i] == bytes32(0)) revert InvalidOutbox();
            _outboxes[chainIds[i]] = outboxes[i];
        }
    }

    /**
     * @dev Parse and return order data, validate correctness.
     * @param order OnchainCrossChainOrder to parse
     */
    function _validate(OnchainCrossChainOrder calldata order) internal view returns (OrderData memory) {
        if (order.fillDeadline < block.timestamp) revert InvalidFillDeadline();
        if (order.orderDataType != ORDER_DATA_TYPEHASH) revert InvalidOrderDataTypehash();
        if (order.orderData.length == 0) revert InvalidOrderData();

        OrderData memory orderData = abi.decode(order.orderData, (OrderData));
        Call[] memory calls = orderData.calls;
        Deposit[] memory deposits = orderData.deposits;
        TokenExpense[] memory expenses = orderData.expenses;

        uint256 callNativeTokens;
        if (calls.length == 0) revert NoCalls();
        for (uint256 i; i < calls.length; ++i) {
            Call memory call = calls[i];
            if (call.chainId == 0) revert NoCallChainId();
            if (call.target == bytes32(0)) revert NoCallTarget();
            if (call.value == 0 && call.data.length == 0) revert NoCallData(); // Seems frivolous, will block fallback calls

            unchecked {
                if (call.value > 0) callNativeTokens += call.value;
            }
        }

        bool hasNative;
        if (deposits.length == 0) revert NoDeposits();
        for (uint256 i; i < deposits.length; ++i) {
            Deposit memory deposit = deposits[i];
            if (deposit.amount == 0) revert NoDepositAmount();
            if (deposit.token == bytes32(0)) {
                if (hasNative) revert DuplicateNativeDeposit();
                hasNative = true;
            }
        }

        uint256 expenseNativeTokens;
        for (uint256 i; i < expenses.length; ++i) {
            TokenExpense memory expense = expenses[i];
            if (expense.amount == 0) revert NoExpenseAmount();
            if (expense.chainId == 0) revert NoExpenseChainId();

            unchecked {
                if (expense.token == bytes32(0)) expenseNativeTokens += expense.amount;
            }
        }

        if (callNativeTokens != expenseNativeTokens) {
            revert InvalidExpenseNativeAmount(callNativeTokens, expenseNativeTokens);
        }

        return orderData;
    }

    /**
     * @notice Validate and intake all deposits.
     * @dev If msg.value > 0, exactly one corresponding native deposit (address == 0x0) is expected.
     * @param deposits Deposits to intake.
     */
    function _processDeposits(Deposit[] memory deposits) internal {
        bool hasNative = msg.value > 0;
        bool processedNative = false;

        for (uint256 i; i < deposits.length; ++i) {
            Deposit memory deposit = deposits[i];

            if (deposit.token == bytes32(0)) {
                // Handle native deposit
                if (!hasNative) revert InvalidNativeDeposit();
                if (deposit.amount != msg.value) revert InvalidNativeDeposit();
                if (processedNative) revert DuplicateNativeDeposit();
                processedNative = true;
            } else {
                // Handle ERC20 deposit
                address token = deposit.token.toAddress();
                token.safeTransferFrom(msg.sender, address(this), deposit.amount);
            }
        }

        if (hasNative && !processedNative) revert InvalidNativeDeposit();
    }

    /**
     * @dev Transfer deposits to recipient. Used for both refunds and claims.
     * @param to       Address to send deposits to.
     * @param deposits Array of deposit Outputs to transfer.
     */
    function _transferDeposits(address to, Output[] memory deposits) internal {
        if (to == address(0)) revert InvalidRecipient();

        for (uint256 i; i < deposits.length; ++i) {
            Output memory deposit = deposits[i];
            if (deposit.token == bytes32(0)) {
                to.safeTransferETH(deposit.amount);
            } else {
                address token = deposit.token.toAddress();
                token.safeTransfer(to, deposit.amount);
            }
        }
    }

    /**
     * @dev Return the next order ID.
     */
    function _nextId() internal view returns (bytes32) {
        return bytes32(_lastId + 1);
    }

    /**
     * @dev Increment and return the next order ID.
     */
    function _incrementId() internal returns (bytes32) {
        return bytes32(++_lastId);
    }

    /**
     * @dev Returns call hash. Used to discern fulfillment.
     * @param resolvedOrder Resolved order to hash.
     */
    function _orderHash(ResolvedCrossChainOrder memory resolvedOrder) internal pure returns (bytes32) {
        return keccak256(abi.encode(resolvedOrder));
    }

    /**
     * @dev Update order state by order hash.
     * @param orderId   ID of the order.
     * @param orderHash Hash of the order.
     * @param status    Updated status.
     */
    function _updateOrderState(bytes32 orderId, bytes32 orderHash, Status status, address acceptedBy) internal {
        OrderState memory state = _orderState[orderHash];
        StatusUpdate memory update = StatusUpdate({ status: status, timestamp: uint40(block.timestamp) });

        if (status == Status.Accepted) {
            state.accepted = update;
            if (acceptedBy != address(0)) state.acceptedBy = acceptedBy;
        }
        state.latest = update;

        _orderState[orderHash] = state;
        _latestOrderIdByStatus[status] = orderId;
    }
}
