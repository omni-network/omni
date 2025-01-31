// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OwnableRoles } from "solady/src/auth/OwnableRoles.sol";
import { ReentrancyGuard } from "solady/src/utils/ReentrancyGuard.sol";
import { Initializable } from "solady/src/utils/Initializable.sol";
import { DeployedAt } from "../util/DeployedAt.sol";
import { XAppBase } from "core/src/pkg/XAppBase.sol";
import { IERC7683 } from "../erc7683/IERC7683.sol";
import { ISolverNetInbox } from "./interfaces/ISolverNetInbox.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { AddrUtils } from "../lib/AddrUtils.sol";

/**
 * @title SolverNetInbox
 * @notice Entrypoint and alt-mempool for user solve orders.
 */
contract SolverNetInbox is OwnableRoles, ReentrancyGuard, Initializable, DeployedAt, XAppBase, ISolverNetInbox {
    using SafeTransferLib for address;
    using AddrUtils for address;

    /**
     * @notice Role for solvers.
     * @dev _ROLE_0 evaluates to '1'.
     */
    uint256 internal constant SOLVER = _ROLE_0;

    /**
     * @notice Typehash for the Order struct.
     */
    bytes32 internal constant ORDER_TYPEHASH = keccak256(
        "Order(address owner,Call call,Values values,Deposit deposit,Expense[] expenses)Call(uint64 chainId,address target,bytes4 selector,bytes callParams)Values(uint96 nativeTip,uint96 callValue,uint32 openDeadline,uint32 fillDeadline)Deposit(address token,uint96 amount)Expense(address spender,address token,uint96 amount)"
    );

    /**
     * @dev Counter for generating unique order IDs. Incremented each time a new order is created.
     */
    uint256 internal _lastId;

    /**
     * @notice Addresses of the outbox contracts.
     */
    mapping(uint64 chainId => address outbox) internal _outboxes;

    /**
     * @notice Map order ID to owner.
     */
    mapping(bytes32 id => address owner) internal _orderOwners;

    /**
     * @notice Map order ID to call parameters.
     * @dev (chainId, target, selector, callParams)
     */
    mapping(bytes32 id => Call) internal _orderCalls;

    /**
     * @notice Map order ID to order values.
     * @dev (nativeTip, callValue, openDeadline, fillDeadline)
     */
    mapping(bytes32 id => Values) internal _orderValues;

    /**
     * @notice Map order ID to deposit parameters.
     * @dev (token, amount)
     */
    mapping(bytes32 id => Deposit) internal _orderDeposits;

    /**
     * @notice Map order ID to expense parameters.
     * @dev (spender, token, amount)
     */
    mapping(bytes32 id => Expense[]) internal _orderExpenses;

    /**
     * @notice Map order ID to order parameters.
     */
    mapping(bytes32 id => OrderState) internal _orderState;

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
     * @param omni_   Address of the OmniPortal.
     */
    function initialize(address owner_, address solver_, address omni_) external initializer {
        _initializeOwner(owner_);
        _grantRoles(solver_, SOLVER);
        _setOmniPortal(omni_);
    }

    /**
     * @notice Set the outbox addresses for the given chain IDs.
     * @param chainIds IDs of the chains.
     * @param outboxes Addresses of the outboxes.
     */
    function setOutboxes(uint64[] calldata chainIds, address[] calldata outboxes) external onlyOwner {
        for (uint256 i; i < chainIds.length; ++i) {
            _outboxes[chainIds[i]] = outboxes[i];
            emit OutboxSet(chainIds[i], outboxes[i]);
        }
    }

    /**
     * @notice Returns the order and its state with the given ID.
     * @param id ID of the order.
     */
    function getOrder(bytes32 id)
        external
        view
        returns (ResolvedCrossChainOrder memory resolved, OrderState memory state)
    {
        OrderData memory orderData = _getOrderData(id);
        return (_resolve(orderData), _orderState[id]);
    }

    /**
     * @notice Returns the next order ID.
     */
    function getNextId() external view returns (bytes32) {
        return _nextId();
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
     * @notice Resolve the onchain order with validation.
     * @param order OnchainCrossChainOrder to resolve.
     */
    function resolve(OnchainCrossChainOrder calldata order) public view returns (ResolvedCrossChainOrder memory) {
        OrderData memory orderData = _validate(order);
        return _resolve(orderData);
    }

    /**
     * @notice Open an order to execute a call on another chain, backed by deposits.
     * @dev Token deposits are transferred from msg.sender to this inbox.
     * @param order OnchainCrossChainOrder to open.
     */
    function open(OnchainCrossChainOrder calldata order) external payable nonReentrant {
        OrderData memory orderData = _validate(order);
        _processDeposits(orderData);
        ResolvedCrossChainOrder memory resolved = _openOrder(orderData);

        emit Open(resolved.orderId, resolved);
    }

    /**
     * @notice Accept an open order.
     * @dev Only a whitelisted solver can accept.
     * @param id ID of the order.
     */
    function accept(bytes32 id) external onlyRoles(SOLVER) nonReentrant {
        Values memory values = _orderValues[id];
        OrderState memory state = _orderState[id];

        if (state.latest.status != Status.Pending) revert OrderNotPending();
        if (values.fillDeadline < block.timestamp && values.fillDeadline != 0) revert FillDeadlinePassed();

        StatusUpdate memory statusUpdate = StatusUpdate({ status: Status.Accepted, timestamp: uint32(block.timestamp) });
        _upsertOrder(id, statusUpdate, msg.sender);

        emit Accepted(id, msg.sender);
    }

    /**
     * @notice Reject an open order and refund deposits.
     * @dev Only a whitelisted solver can reject.
     * @param id     ID of the order.
     * @param reason Reason code for rejection.
     */
    function reject(bytes32 id, uint8 reason) external onlyRoles(SOLVER) nonReentrant {
        OrderState memory state = _orderState[id];

        if (state.latest.status != Status.Pending) {
            if (state.latest.status == Status.Accepted) {
                if (state.acceptedBy != msg.sender) revert Unauthorized();
            } else {
                revert OrderNotPending();
            }
        }

        StatusUpdate memory statusUpdate = StatusUpdate({ status: Status.Rejected, timestamp: uint32(block.timestamp) });
        _upsertOrder(id, statusUpdate, address(0));
        _transferDeposits(id, _orderOwners[id]);

        emit Rejected(id, msg.sender, reason);
    }

    /**
     * @notice Cancel an open and refund deposits.
     * @dev Only order initiator can cancel.
     * @param id ID of the order.
     */
    function cancel(bytes32 id) external nonReentrant {
        OrderState memory state = _orderState[id];
        address user = _orderOwners[id];

        if (state.latest.status != Status.Pending) revert OrderNotPending();
        if (user != msg.sender) revert Unauthorized();

        StatusUpdate memory statusUpdate = StatusUpdate({ status: Status.Reverted, timestamp: uint32(block.timestamp) });
        _upsertOrder(id, statusUpdate, address(0));
        _transferDeposits(id, user);

        emit Reverted(id);
    }

    /**
     * @notice Fill an order.
     * @dev Only callable by the outbox.
     * @param id         ID of the order.
     * @param fillHash   Hash of fill instructions origin data.
     * @param timestamp  Timestamp of the fill.
     * @param acceptedBy Address of the solver that filled the order if they didnt accept it first.
     */
    function markFilled(bytes32 id, bytes32 fillHash, uint256 timestamp, address acceptedBy)
        external
        xrecv
        nonReentrant
    {
        Call memory call = _orderCalls[id];
        OrderState memory state = _orderState[id];

        if (state.latest.status != Status.Pending && state.latest.status != Status.Accepted) {
            revert OrderNotPendingOrAccepted();
        }
        if (xmsg.sender != _outboxes[xmsg.sourceChainId]) revert Unauthorized();
        if (xmsg.sourceChainId != call.chainId) revert WrongSourceChain();

        // Ensure reported fill hash matches origin data
        if (fillHash != _fillHash(id)) {
            revert WrongFillHash();
        }

        StatusUpdate memory statusUpdate = StatusUpdate({ status: Status.Filled, timestamp: uint32(timestamp) });
        _upsertOrder(id, statusUpdate, acceptedBy);

        emit Filled(id, fillHash, state.acceptedBy);
    }

    /**
     * @notice Claim a filled order.
     * @param id ID of the order.
     * @param to Address to send deposits to.
     */
    function claim(bytes32 id, address to) external nonReentrant {
        OrderState memory state = _orderState[id];
        if (state.latest.status != Status.Filled) revert OrderNotFilled();
        if (state.acceptedBy != msg.sender) revert Unauthorized();

        StatusUpdate memory statusUpdate = StatusUpdate({ status: Status.Claimed, timestamp: uint32(block.timestamp) });
        _upsertOrder(id, statusUpdate, address(0));
        _transferDeposits(id, to);

        emit Claimed(id, msg.sender, to);
    }

    /**
     * @dev Return the order data for the given ID.
     * @param id ID of the order.
     */
    function _getOrderData(bytes32 id) internal view returns (OrderData memory) {
        return OrderData({
            owner: _orderOwners[id],
            call: _orderCalls[id],
            values: _orderValues[id],
            deposit: _orderDeposits[id],
            expenses: _orderExpenses[id]
        });
    }

    /**
     * @dev Parse and return order data, validate correctness.
     * @param order OnchainCrossChainOrder to parse
     */
    function _validate(OnchainCrossChainOrder calldata order) internal view returns (OrderData memory) {
        if (order.fillDeadline < block.timestamp && order.fillDeadline != 0) revert InvalidFillDeadline();
        if (order.orderDataType != ORDER_TYPEHASH) revert InvalidOrderTypehash();
        if (order.orderData.length == 0) revert InvalidOrderData();

        OrderData memory orderData = abi.decode(order.orderData, (OrderData));
        Call memory call = orderData.call;
        Values memory values = orderData.values;
        Deposit memory deposit = orderData.deposit;
        Expense[] memory expenses = orderData.expenses;

        if (orderData.owner == address(0)) orderData.owner = msg.sender;
        if (call.chainId == 0 || call.chainId == block.chainid) revert InvalidCallChainId();
        if (call.target == address(0)) revert InvalidCallTarget();
        if (values.openDeadline < block.timestamp) revert InvalidOpenDeadline();
        if (values.fillDeadline <= values.openDeadline || order.fillDeadline != values.fillDeadline) {
            revert InvalidFillDeadline();
        }
        if (deposit.token == address(0) && deposit.amount != 0) revert InvalidDepositAmount();
        if (deposit.token != address(0) && deposit.amount == 0) revert InvalidDepositAmount();
        for (uint256 i; i < expenses.length; ++i) {
            if (expenses[i].token == address(0)) revert InvalidExpenseToken();
            if (expenses[i].amount == 0) revert InvalidExpenseAmount();
        }

        return orderData;
    }

    /**
     * @dev Derive the maxSpent Output for the order.
     * @param orderData Order data to derive from.
     */
    function _deriveMaxSpent(OrderData memory orderData) internal view returns (IERC7683.Output[] memory) {
        Call memory call = orderData.call;
        Values memory values = orderData.values;
        Expense[] memory expenses = orderData.expenses;

        IERC7683.Output[] memory maxSpent =
            new IERC7683.Output[](values.callValue > 0 ? expenses.length + 1 : expenses.length);
        for (uint256 i; i < expenses.length; ++i) {
            maxSpent[i] = IERC7683.Output({
                token: expenses[i].token.toBytes32(),
                amount: expenses[i].amount,
                recipient: _outboxes[call.chainId].toBytes32(),
                chainId: call.chainId
            });
        }
        if (values.callValue > 0) {
            maxSpent[expenses.length] = IERC7683.Output({
                token: bytes32(0),
                amount: values.callValue,
                recipient: _outboxes[call.chainId].toBytes32(),
                chainId: call.chainId
            });
        }

        return maxSpent;
    }

    /**
     * @dev Derive the minReceived Output for the order.
     * @param orderData Order data to derive from.
     */
    function _deriveMinReceived(OrderData memory orderData) internal view returns (IERC7683.Output[] memory) {
        Values memory values = orderData.values;
        Deposit memory deposit = orderData.deposit;

        uint8 deposits;
        bool erc20Deposit = deposit.token != address(0);
        bool nativeDeposit = values.nativeTip > 0;
        unchecked {
            if (erc20Deposit) ++deposits;
            if (nativeDeposit) ++deposits;
        }

        IERC7683.Output[] memory minReceived = new IERC7683.Output[](deposits);
        if (erc20Deposit) {
            minReceived[0] = IERC7683.Output({
                token: deposit.token.toBytes32(),
                amount: deposit.amount,
                recipient: bytes32(0),
                chainId: block.chainid
            });
        }
        if (nativeDeposit) {
            minReceived[deposits - 1] = IERC7683.Output({
                token: bytes32(0),
                amount: values.nativeTip,
                recipient: bytes32(0),
                chainId: block.chainid
            });
        }

        return minReceived;
    }

    /**
     * @dev Derive the fillInstructions for the order.
     * @param orderData Order data to derive from.
     */
    function _deriveFillInstructions(OrderData memory orderData)
        internal
        view
        returns (IERC7683.FillInstruction[] memory)
    {
        Call memory call = orderData.call;
        Values memory values = orderData.values;
        Expense[] memory expenses = orderData.expenses;

        IERC7683.FillInstruction[] memory fillInstructions = new IERC7683.FillInstruction[](1);
        fillInstructions[0] = IERC7683.FillInstruction({
            destinationChainId: call.chainId,
            destinationSettler: _outboxes[call.chainId].toBytes32(),
            originData: abi.encode(
                FillOriginData({
                    srcChainId: uint64(block.chainid),
                    destChainId: call.chainId,
                    fillDeadline: values.fillDeadline,
                    callValue: values.callValue,
                    target: call.target,
                    callData: abi.encodePacked(call.selector, call.callParams),
                    expenses: expenses
                })
            )
        });

        return fillInstructions;
    }

    /**
     * @dev Resolve the order without validation.
     * @param orderData Order data to resolve.
     */
    function _resolve(OrderData memory orderData) internal view returns (ResolvedCrossChainOrder memory) {
        Values memory values = orderData.values;

        IERC7683.Output[] memory maxSpent = _deriveMaxSpent(orderData);
        IERC7683.Output[] memory minReceived = _deriveMinReceived(orderData);
        IERC7683.FillInstruction[] memory fillInstructions = _deriveFillInstructions(orderData);

        return ResolvedCrossChainOrder({
            user: orderData.owner,
            originChainId: block.chainid,
            openDeadline: values.openDeadline,
            fillDeadline: values.fillDeadline,
            orderId: _nextId(),
            maxSpent: maxSpent,
            minReceived: minReceived,
            fillInstructions: fillInstructions
        });
    }

    /**
     * @notice Validate and intake ERC20 and/or native deposits.
     * @param orderData  Order data to process.
     */
    function _processDeposits(OrderData memory orderData) internal {
        Values memory values = orderData.values;
        Deposit memory deposit = orderData.deposit;

        if (msg.value != values.nativeTip) revert InvalidNativeTip();
        if (deposit.token != address(0)) {
            deposit.token.safeTransferFrom(msg.sender, address(this), deposit.amount);
        }
    }

    /**
     * @dev Opens a new order by initializing its state.
     * @param orderData Order data to open.
     */
    function _openOrder(OrderData memory orderData) internal returns (ResolvedCrossChainOrder memory resolved) {
        resolved = _resolve(orderData);
        bytes32 id = _incrementId();

        _orderOwners[id] = orderData.owner;
        _orderCalls[id] = orderData.call;
        _orderValues[id] = orderData.values;
        _orderDeposits[id] = orderData.deposit;
        for (uint256 i; i < orderData.expenses.length; ++i) {
            _orderExpenses[id].push(orderData.expenses[i]);
        }

        StatusUpdate memory statusUpdate = StatusUpdate({ status: Status.Pending, timestamp: uint32(block.timestamp) });
        _upsertOrder(id, statusUpdate, address(0));

        return resolved;
    }

    /**
     * @dev Transfer deposits to recipient. Used for both refunds and claims.
     * @param id ID of the order.
     * @param to Address to send deposits to.
     */
    function _transferDeposits(bytes32 id, address to) internal {
        Values memory values = _orderValues[id];
        Deposit memory deposit = _orderDeposits[id];

        if (values.nativeTip > 0) {
            to.safeTransferETH(values.nativeTip);
        }
        if (deposit.token != address(0)) {
            deposit.token.safeTransfer(to, deposit.amount);
        }
    }

    /**
     * @dev Update or insert order state by id.
     * @param id           ID of the order.
     * @param statusUpdate Status update to upsert.
     * @param acceptedBy   Address of the solver who accepted the order, if applicable.
     */
    function _upsertOrder(bytes32 id, StatusUpdate memory statusUpdate, address acceptedBy) internal {
        OrderState memory state = _orderState[id];

        // Apply most recent status update
        state.latest = statusUpdate;

        // If statusUpdate is accepted, update accepted status
        if (statusUpdate.status == Status.Accepted) {
            state.accepted = statusUpdate;
            state.acceptedBy = acceptedBy;
        }

        // If statusUpdate is filled, update acceptedBy if it was filled before being accepted
        if (statusUpdate.status == Status.Filled) {
            if (state.accepted.timestamp > statusUpdate.timestamp) {
                state.acceptedBy = acceptedBy;
            }
        }

        _orderState[id] = state;
        _latestOrderIdByStatus[statusUpdate.status] = id;
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
     * @param orderId ID of the order.
     */
    function _fillHash(bytes32 orderId) internal view returns (bytes32) {
        Call memory call = _orderCalls[orderId];
        Values memory values = _orderValues[orderId];
        Expense[] memory expenses = _orderExpenses[orderId];

        FillOriginData memory fillOriginData = FillOriginData({
            srcChainId: uint64(block.chainid),
            destChainId: call.chainId,
            fillDeadline: values.fillDeadline,
            callValue: values.callValue,
            target: call.target,
            callData: abi.encodePacked(call.selector, call.callParams),
            expenses: expenses
        });

        return keccak256(abi.encode(orderId, fillOriginData));
    }
}
