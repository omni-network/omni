// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OwnableRoles } from "solady/src/auth/OwnableRoles.sol";
import { ReentrancyGuard } from "solady/src/utils/ReentrancyGuard.sol";
import { Initializable } from "solady/src/utils/Initializable.sol";
import { XAppBase } from "core/src/pkg/XAppBase.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { DeployedAt } from "./util/DeployedAt.sol";
import { AddrUtils } from "./lib/AddrUtils.sol";
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
        "OrderData(address owner,Call call,Deposit[] deposits)Call(uint64 chainId,bytes32 target,uint256 value,bytes data,TokenExpense[] expenses)TokenExpense(bytes32 token,bytes32 spender,uint256 amount)Deposit(bytes32 token,uint256 amount)"
    ); // Not really needed until we support more than one order type or gasless orders

    /**
     * @dev Counter for generating unique order IDs. Incremented each time a new order is created.
     */
    uint256 internal _lastId;

    /**
     * @notice Address of the outbox contract.
     */
    address internal _outbox;

    /**
     * @notice Map order ID to resolved onchain order.
     */
    mapping(bytes32 id => OnchainCrossChainOrder) internal _orders;

    /**
     * @notice Map order ID to order parameters.
     */
    mapping(bytes32 id => OrderState) internal _orderState;

    /**
     * @notice Map order ID to order history.
     */
    mapping(bytes32 id => StatusUpdate[]) internal _orderHistory;

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
     */
    function initialize(address owner_, address solver_, address omni_, address outbox_) external initializer {
        _initializeOwner(owner_);
        _grantRoles(solver_, SOLVER);
        _setOmniPortal(omni_);
        _outbox = outbox_;
    }

    /**
     * @notice Returns the order with the given ID.
     * @param id  ID of the order.
     */
    function getOrder(bytes32 id)
        external
        view
        returns (ResolvedCrossChainOrder memory resolved, OrderState memory state, StatusUpdate[] memory history)
    {
        OnchainCrossChainOrder memory order = _orders[id];

        if (order.orderData.length > 0) {
            OrderData memory orderData = abi.decode(order.orderData, (OrderData));
            resolved = _resolve(orderData, order.fillDeadline);
        }

        return (resolved, _orderState[id], _orderHistory[id]);
    }

    /**
     * @notice Returns the next order ID.
     */
    function getNextId() external view returns (bytes32) {
        return _nextId();
    }

    /**
     * @notice Returns the latest order with the given status.
     * @param status  Order status to query.
     */
    function getLatestOrderIdByStatus(Status status) external view returns (bytes32) {
        return _latestOrderIdByStatus[status];
    }

    /**
     * @dev Validate the onchain order.
     * @param order  OnchainCrossChainOrder to validate.
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
        return _resolve(order);
    }

    /**
     * @notice Open an order to execute a call on another chain, backed by deposits.
     * @dev Token deposits are transferred from msg.sender to this inbox.
     * @param order OnchainCrossChainOrder to open.
     */
    function open(OnchainCrossChainOrder calldata order) external payable nonReentrant {
        OrderData memory orderData = abi.decode(order.orderData, (OrderData));

        _processDeposits(orderData.deposits);

        ResolvedCrossChainOrder memory resolved = _openOrder(order);

        emit Open(resolved.orderId, resolved);
    }

    /**
     * @notice Accept an open order.
     * @dev Only a whitelisted solver can accept.
     * @param id  ID of the order.
     */
    function accept(bytes32 id) external onlyRoles(SOLVER) nonReentrant {
        OnchainCrossChainOrder memory order = _orders[id];
        OrderState memory state = _orderState[id];

        if (state.status != Status.Pending) revert OrderNotPending();
        if (order.fillDeadline < block.timestamp && order.fillDeadline != 0) revert FillDeadlinePassed();

        state.status = Status.Accepted;
        state.acceptedBy = msg.sender;
        _upsertOrder(id, state);

        emit Accepted(id, msg.sender);
    }

    /**
     * @notice Reject an open order and refund deposits.
     * @dev Only a whitelisted solver can reject.
     * @param id      ID of the order.
     * @param reason  Reason code for rejection.
     */
    function reject(bytes32 id, uint8 reason) external onlyRoles(SOLVER) nonReentrant {
        OnchainCrossChainOrder memory order = _orders[id];
        OrderData memory orderData = abi.decode(order.orderData, (OrderData));
        ResolvedCrossChainOrder memory resolved = _resolve(orderData, order.fillDeadline);
        OrderState memory state = _orderState[id];

        if (state.status != Status.Pending) {
            if (state.status == Status.Accepted) {
                if (state.acceptedBy != msg.sender) revert Unauthorized();
            } else {
                revert OrderNotPending();
            }
        }

        state.status = Status.Rejected;
        _upsertOrder(id, state);
        _transferDeposits(resolved.user, resolved.minReceived);

        emit Rejected(id, msg.sender, reason);
    }

    /**
     * @notice Cancel an open and refund deposits.
     * @dev Only order initiator can cancel.
     * @param id  ID of the order.
     */
    function cancel(bytes32 id) external nonReentrant {
        OnchainCrossChainOrder memory order = _orders[id];
        OrderData memory orderData = abi.decode(order.orderData, (OrderData));
        ResolvedCrossChainOrder memory resolved = _resolve(orderData, order.fillDeadline);
        OrderState memory state = _orderState[id];

        if (state.status != Status.Pending) revert OrderNotPending();
        if (resolved.user != msg.sender) revert Unauthorized();

        state.status = Status.Reverted;
        _upsertOrder(id, state);
        _transferDeposits(resolved.user, resolved.minReceived);

        emit Reverted(id);
    }

    /**
     * @notice Fill an order.
     * @dev Only callable by the outbox.
     * @param id        ID of the order.
     * @param fillHash  Hash of fill instructions origin data.
     */
    function markFilled(bytes32 id, bytes32 fillHash) external xrecv nonReentrant {
        OnchainCrossChainOrder memory order = _orders[id];
        OrderData memory orderData = abi.decode(order.orderData, (OrderData));
        ResolvedCrossChainOrder memory resolved = _resolve(orderData, order.fillDeadline);
        OrderState memory state = _orderState[id];

        if (state.status != Status.Accepted) revert OrderNotAccepted();
        if (xmsg.sender != _outbox) revert NotOutbox();
        if (xmsg.sourceChainId != resolved.fillInstructions[0].destinationChainId) revert WrongSourceChain();

        // Ensure reported fill hash matches origin data
        if (fillHash != _fillHash(id, resolved.fillInstructions[0].originData)) {
            revert WrongFillHash();
        }

        state.status = Status.Filled;
        _upsertOrder(id, state);

        emit Filled(id, fillHash, state.acceptedBy);
    }

    /**
     * @notice Claim a filled order.
     * @param id  ID of the order.
     * @param to  Address to send deposits to.
     */
    function claim(bytes32 id, address to) external nonReentrant {
        OnchainCrossChainOrder memory order = _orders[id];
        OrderData memory orderData = abi.decode(order.orderData, (OrderData));
        ResolvedCrossChainOrder memory resolved = _resolve(orderData, order.fillDeadline);
        OrderState memory state = _orderState[id];

        if (state.status != Status.Filled) revert OrderNotFilled();
        if (state.acceptedBy != msg.sender) revert Unauthorized();

        state.status = Status.Claimed;

        _upsertOrder(id, state);
        _transferDeposits(to, resolved.minReceived);

        emit Claimed(id, msg.sender, to, resolved.minReceived);
    }

    /**
     * @dev Parse and return order data, validate correctness.
     * @param order  OnchainCrossChainOrder to parse
     */
    function _validate(OnchainCrossChainOrder calldata order) internal view returns (OrderData memory) {
        if (order.fillDeadline < block.timestamp && order.fillDeadline != 0) revert InvalidFillDeadline();
        if (order.orderDataType != ORDER_DATA_TYPEHASH) revert InvalidOrderDataTypehash();
        if (order.orderData.length == 0) revert InvalidOrderData();

        OrderData memory orderData = abi.decode(order.orderData, (OrderData));
        Call memory call = orderData.call;
        Deposit[] memory deposits = orderData.deposits;

        if (call.chainId == 0) revert NoCallChainId();
        if (call.target == bytes32(0)) revert NoCallTarget();
        if (deposits.length == 0) revert NoDeposits();

        bool hasNative;
        for (uint256 i; i < deposits.length; ++i) {
            Deposit memory deposit = deposits[i];
            if (deposit.amount == 0) revert NoDepositAmount();
            if (deposit.token == bytes32(0)) {
                if (hasNative) revert DuplicateNativeDeposit();
                hasNative = true;
            }
        }

        for (uint256 i; i < call.expenses.length; ++i) {
            TokenExpense memory expense = call.expenses[i];
            if (expense.token == bytes32(0)) revert NoExpenseToken();
            if (expense.amount == 0) revert NoExpenseAmount();
        }

        return orderData;
    }

    /**
     * @dev Resolve the order after validating it
     * @param order  OnchainCrossChainOrder to resolve.
     */
    function _resolve(OnchainCrossChainOrder calldata order) internal view returns (ResolvedCrossChainOrder memory) {
        OrderData memory orderData = _validate(order);
        Call memory call = orderData.call;
        Deposit[] memory deposits = orderData.deposits;

        bool hasNative = call.value > 0;
        Output[] memory maxSpent = new Output[](hasNative ? call.expenses.length + 1 : call.expenses.length);
        for (uint256 i; i < call.expenses.length; ++i) {
            maxSpent[i] = Output({
                token: call.expenses[i].token,
                amount: call.expenses[i].amount,
                recipient: _outbox.toBytes32(), // for solver, recipient is always outbox
                chainId: call.chainId
            });
        }
        if (hasNative) {
            maxSpent[call.expenses.length] =
                Output({ token: bytes32(0), amount: call.value, recipient: _outbox.toBytes32(), chainId: call.chainId });
        }

        Output[] memory minReceived = new Output[](deposits.length);
        for (uint256 i; i < deposits.length; ++i) {
            minReceived[i] = Output({
                token: deposits[i].token,
                amount: deposits[i].amount,
                recipient: bytes32(0), // recipient is solver, which is only known on acceptance
                chainId: block.chainid
            });
        }

        FillInstruction[] memory fillInstructions = new FillInstruction[](1);
        fillInstructions[0] = FillInstruction({
            destinationChainId: call.chainId,
            destinationSettler: _outbox.toBytes32(),
            originData: abi.encode(
                FillOriginData({ srcChainId: uint64(block.chainid), fillDeadline: order.fillDeadline, call: call })
            )
        });

        return ResolvedCrossChainOrder({
            user: orderData.owner,
            originChainId: block.chainid,
            openDeadline: uint32(block.timestamp),
            fillDeadline: order.fillDeadline,
            orderId: _nextId(), // use next id for view, may not be id when opened
            maxSpent: maxSpent,
            minReceived: minReceived,
            fillInstructions: fillInstructions
        });
    }

    /**
     * @dev Resolve the orderData stored onchain without validating it
     * @param orderData  OrderData to resolve.
     */
    function _resolve(OrderData memory orderData, uint40 fillDeadline)
        internal
        view
        returns (ResolvedCrossChainOrder memory)
    {
        Call memory call = orderData.call;
        Deposit[] memory deposits = orderData.deposits;

        bool hasNative = call.value > 0;
        Output[] memory maxSpent = new Output[](hasNative ? call.expenses.length + 1 : call.expenses.length);
        for (uint256 i; i < call.expenses.length; ++i) {
            maxSpent[i] = Output({
                token: call.expenses[i].token,
                amount: call.expenses[i].amount,
                recipient: _outbox.toBytes32(), // for solver, recipient is always outbox
                chainId: call.chainId
            });
        }
        if (hasNative) {
            maxSpent[call.expenses.length] =
                Output({ token: bytes32(0), amount: call.value, recipient: _outbox.toBytes32(), chainId: call.chainId });
        }

        Output[] memory minReceived = new Output[](deposits.length);
        for (uint256 i; i < deposits.length; ++i) {
            minReceived[i] = Output({
                token: deposits[i].token,
                amount: deposits[i].amount,
                recipient: bytes32(0), // recipient is solver, which is only known on acceptance
                chainId: block.chainid
            });
        }

        FillInstruction[] memory fillInstructions = new FillInstruction[](1);
        fillInstructions[0] = FillInstruction({
            destinationChainId: call.chainId,
            destinationSettler: _outbox.toBytes32(),
            originData: abi.encode(
                FillOriginData({ srcChainId: uint64(block.chainid), fillDeadline: fillDeadline, call: call })
            )
        });

        return ResolvedCrossChainOrder({
            user: orderData.owner,
            originChainId: block.chainid,
            openDeadline: uint32(block.timestamp),
            fillDeadline: uint32(fillDeadline),
            orderId: _nextId(), // use next id for view, may not be id when opened
            maxSpent: maxSpent,
            minReceived: minReceived,
            fillInstructions: fillInstructions
        });
    }

    /**
     * @notice Validate and intake all deposits.
     * @dev If msg.value > 0, exactly one corresponding native deposit (address == 0x0) is expected.
     * @param deposits  Deposits to intake.
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
     * @dev Opens a new order and initializes its state.
     * @param order The order to open.
     * @return resolved The storage reference to the newly created order.
     */
    function _openOrder(OnchainCrossChainOrder calldata order)
        internal
        returns (ResolvedCrossChainOrder memory resolved)
    {
        OnchainCrossChainOrder memory _order = order;
        OrderData memory _orderData = abi.decode(_order.orderData, (OrderData));
        if (_orderData.owner == address(0)) {
            _orderData.owner = msg.sender;
            _order.orderData = abi.encode(_orderData);
        }

        resolved = _resolve(order);
        _orders[resolved.orderId] = _order;
        _upsertOrder(resolved.orderId, OrderState({ status: Status.Pending, acceptedBy: address(0) }));

        return resolved;
    }

    /**
     * @dev Transfer deposits to recipient. Used for both refunds and claims.
     * @param to  Address to send deposits to.
     * @param deposits  Array of Output to transfer.
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
     * @param orderId      ID of the order.
     * @param originData   Encoded fill instruction origin data.
     */
    function _fillHash(bytes32 orderId, bytes memory originData) internal pure returns (bytes32) {
        return keccak256(abi.encode(orderId, originData));
    }

    /**
     * @dev Update or insert order state by id.
     */
    function _upsertOrder(bytes32 id, OrderState memory state) internal {
        _orderState[id] = state;
        _orderHistory[id].push(StatusUpdate({ status: state.status, timestamp: uint40(block.timestamp) }));
        _latestOrderIdByStatus[state.status] = id;
    }
}
