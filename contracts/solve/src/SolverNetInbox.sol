// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OwnableRoles } from "solady/src/auth/OwnableRoles.sol";
import { ReentrancyGuard } from "solady/src/utils/ReentrancyGuard.sol";
import { Initializable } from "solady/src/utils/Initializable.sol";
import { DeployedAt } from "./util/DeployedAt.sol";
import { XAppBase } from "core/src/pkg/XAppBase.sol";
import { IERC7683 } from "./erc7683/IERC7683.sol";
import { ISolverNetInbox } from "./interfaces/ISolverNetInbox.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { SolverNet } from "./lib/SolverNet.sol";
import { AddrUtils } from "./lib/AddrUtils.sol";

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
     * @notice Typehash for the OrderData struct.
     */
    bytes32 internal constant ORDERDATA_TYPEHASH = keccak256(
        "OrderData(address owner,uint64 destChainId,Deposit deposit,Call[] calls,Expense[] expenses)Deposit(address token,uint96 amount)Call(address target,bytes4 selector,uint256 value,bytes params)Expense(address spender,address token,uint96 amount)"
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
     * @notice Map order ID to header parameters.
     * @dev (owner, destChainId, fillDeadline)
     */
    mapping(bytes32 id => SolverNet.Header) internal _orderHeader;

    /**
     * @notice Map order ID to deposit parameters.
     * @dev (token, amount)
     */
    mapping(bytes32 id => SolverNet.Deposit) internal _orderDeposit;

    /**
     * @notice Map order ID to call parameters.
     * @dev (target, selector, value, params)
     */
    mapping(bytes32 id => SolverNet.Call[]) internal _orderCalls;

    /**
     * @notice Map order ID to expense parameters.
     * @dev (spender, token, amount)
     */
    mapping(bytes32 id => SolverNet.Expense[]) internal _orderExpenses;

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
        SolverNet.Order memory orderData = _getOrder(id);
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
        SolverNet.Order memory orderData = _validate(order);
        return _resolve(orderData);
    }

    /**
     * @notice Open an order to execute a call on another chain, backed by deposits.
     * @dev Token deposits are transferred from msg.sender to this inbox.
     * @param order OnchainCrossChainOrder to open.
     */
    function open(OnchainCrossChainOrder calldata order) external payable nonReentrant {
        SolverNet.Order memory orderData = _validate(order);
        _processDeposit(orderData.deposit);
        ResolvedCrossChainOrder memory resolved = _openOrder(orderData);

        emit Open(resolved.orderId, resolved);
    }

    /**
     * @notice Accept an open order.
     * @dev Only a whitelisted solver can accept.
     * @param id ID of the order.
     */
    function accept(bytes32 id) external onlyRoles(SOLVER) nonReentrant {
        SolverNet.Header memory header = _orderHeader[id];
        OrderState memory state = _orderState[id];

        if (state.status != Status.Pending) revert OrderNotPending();
        if (header.fillDeadline < block.timestamp && header.fillDeadline != 0) revert FillDeadlinePassed();

        _upsertOrder(id, Status.Accepted, msg.sender);

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

        if (state.status != Status.Pending) {
            if (state.status != Status.Accepted) revert Unauthorized();
            if (state.claimant != msg.sender) revert Unauthorized();
        }

        _upsertOrder(id, Status.Reverted, msg.sender);
        _transferDeposit(id, _orderHeader[id].owner);

        emit Rejected(id, msg.sender, reason);
    }

    /**
     * @notice Cancel an open and refund deposits.
     * @dev Only order initiator can cancel.
     * @param id ID of the order.
     */
    function cancel(bytes32 id) external nonReentrant {
        OrderState memory state = _orderState[id];
        address user = _orderHeader[id].owner;

        if (state.status != Status.Pending) revert OrderNotPending();
        if (user != msg.sender) revert Unauthorized();

        _upsertOrder(id, Status.Reverted, msg.sender);
        _transferDeposit(id, user);

        emit Reverted(id);
    }

    /**
     * @notice Fill an order.
     * @dev Only callable by the outbox.
     * @param id        ID of the order.
     * @param fillHash  Hash of fill instructions origin data.
     * @param claimant  Address to claim the order, provided by the filler.
     */
    function markFilled(bytes32 id, bytes32 fillHash, address claimant) external xrecv nonReentrant {
        SolverNet.Header memory header = _orderHeader[id];
        OrderState memory state = _orderState[id];

        if (state.status != Status.Pending && state.status != Status.Accepted) {
            revert OrderNotPendingOrAccepted();
        }
        if (xmsg.sender != _outboxes[xmsg.sourceChainId]) revert Unauthorized();
        if (xmsg.sourceChainId != header.destChainId) revert WrongSourceChain();

        // Ensure reported fill hash matches origin data
        if (fillHash != _fillHash(id)) {
            revert WrongFillHash();
        }

        _upsertOrder(id, Status.Filled, claimant);
        emit Filled(id, fillHash, claimant);
    }

    /**
     * @notice Claim a filled order.
     * @param id ID of the order.
     * @param to Address to send deposits to.
     */
    function claim(bytes32 id, address to) external nonReentrant {
        OrderState memory state = _orderState[id];

        if (state.status != Status.Filled) revert OrderNotFilled();
        if (state.claimant != msg.sender) revert Unauthorized();

        _upsertOrder(id, Status.Claimed, msg.sender);
        _transferDeposit(id, to);

        emit Claimed(id, msg.sender, to);
    }

    /**
     * @dev Return the order for the given ID.
     * @param id ID of the order.
     */
    function _getOrder(bytes32 id) internal view returns (SolverNet.Order memory) {
        return SolverNet.Order({
            header: _orderHeader[id],
            calls: _orderCalls[id],
            deposit: _orderDeposit[id],
            expenses: _orderExpenses[id]
        });
    }

    /**
     * @dev Parse and return order data, validate correctness.
     * @param order OnchainCrossChainOrder to parse
     */
    function _validate(OnchainCrossChainOrder calldata order) internal view returns (SolverNet.Order memory) {
        // Validate OnchainCrossChainOrder
        if (order.fillDeadline < block.timestamp && order.fillDeadline != 0) revert InvalidFillDeadline();
        if (order.orderDataType != ORDERDATA_TYPEHASH) revert InvalidOrderTypehash();
        if (order.orderData.length == 0) revert InvalidOrderData();

        SolverNet.OrderData memory orderData = abi.decode(order.orderData, (SolverNet.OrderData));

        // Validate SolverNet.OrderData.Header fields
        if (orderData.owner == address(0)) orderData.owner = msg.sender;
        if (orderData.destChainId == 0 || orderData.destChainId == block.chainid) revert InvalidChainId();

        SolverNet.Header memory header = SolverNet.Header({
            owner: orderData.owner,
            destChainId: orderData.destChainId,
            fillDeadline: order.fillDeadline
        });

        // Validate SolverNet.OrderData.Call
        SolverNet.Call[] memory calls = orderData.calls;
        for (uint256 i; i < calls.length; ++i) {
            SolverNet.Call memory call = calls[i];
            if (call.target == address(0)) revert InvalidCallTarget();
        }

        // Validate SolverNet.OrderData.Expenses
        SolverNet.Expense[] memory expenses = orderData.expenses;
        for (uint256 i; i < expenses.length; ++i) {
            if (expenses[i].token == address(0)) revert InvalidExpenseToken();
            if (expenses[i].amount == 0) revert InvalidExpenseAmount();
        }

        return SolverNet.Order({ header: header, calls: calls, deposit: orderData.deposit, expenses: expenses });
    }

    /**
     * @dev Derive the maxSpent Output for the order.
     * @param orderData Order data to derive from.
     */
    function _deriveMaxSpent(SolverNet.Order memory orderData) internal view returns (IERC7683.Output[] memory) {
        SolverNet.Header memory header = orderData.header;
        SolverNet.Call[] memory calls = orderData.calls;
        SolverNet.Expense[] memory expenses = orderData.expenses;

        uint256 totalNativeValue;
        for (uint256 i; i < calls.length; ++i) {
            if (calls[i].value > 0) totalNativeValue += calls[i].value;
        }

        IERC7683.Output[] memory maxSpent =
            new IERC7683.Output[](totalNativeValue > 0 ? expenses.length + 1 : expenses.length);
        for (uint256 i; i < expenses.length; ++i) {
            maxSpent[i] = IERC7683.Output({
                token: expenses[i].token.toBytes32(),
                amount: expenses[i].amount,
                recipient: _outboxes[header.destChainId].toBytes32(),
                chainId: header.destChainId
            });
        }
        if (totalNativeValue > 0) {
            maxSpent[expenses.length] = IERC7683.Output({
                token: bytes32(0),
                amount: totalNativeValue,
                recipient: _outboxes[header.destChainId].toBytes32(),
                chainId: header.destChainId
            });
        }

        return maxSpent;
    }

    /**
     * @dev Derive the minReceived Output for the order.
     * @param orderData Order data to derive from.
     */
    function _deriveMinReceived(SolverNet.Order memory orderData) internal view returns (IERC7683.Output[] memory) {
        SolverNet.Deposit memory deposit = orderData.deposit;

        IERC7683.Output[] memory minReceived = new IERC7683.Output[](deposit.amount > 0 ? 1 : 0);
        if (deposit.amount > 0) {
            minReceived[0] = IERC7683.Output({
                token: deposit.token.toBytes32(),
                amount: deposit.amount,
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
    function _deriveFillInstructions(SolverNet.Order memory orderData)
        internal
        view
        returns (IERC7683.FillInstruction[] memory)
    {
        SolverNet.Header memory header = orderData.header;
        SolverNet.Call[] memory calls = orderData.calls;
        SolverNet.Expense[] memory expenses = orderData.expenses;

        IERC7683.FillInstruction[] memory fillInstructions = new IERC7683.FillInstruction[](1);
        fillInstructions[0] = IERC7683.FillInstruction({
            destinationChainId: header.destChainId,
            destinationSettler: _outboxes[header.destChainId].toBytes32(),
            originData: abi.encode(
                SolverNet.FillOriginData({
                    srcChainId: uint64(block.chainid),
                    destChainId: header.destChainId,
                    fillDeadline: header.fillDeadline,
                    calls: calls,
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
    function _resolve(SolverNet.Order memory orderData) internal view returns (ResolvedCrossChainOrder memory) {
        SolverNet.Header memory header = orderData.header;

        IERC7683.Output[] memory maxSpent = _deriveMaxSpent(orderData);
        IERC7683.Output[] memory minReceived = _deriveMinReceived(orderData);
        IERC7683.FillInstruction[] memory fillInstructions = _deriveFillInstructions(orderData);

        return ResolvedCrossChainOrder({
            user: header.owner,
            originChainId: block.chainid,
            openDeadline: 0,
            fillDeadline: header.fillDeadline,
            orderId: _nextId(),
            maxSpent: maxSpent,
            minReceived: minReceived,
            fillInstructions: fillInstructions
        });
    }

    /**
     * @notice Validate and intake an ERC20 or native deposit.
     * @param deposit Deposit to process.
     */
    function _processDeposit(SolverNet.Deposit memory deposit) internal {
        if (deposit.token == address(0)) {
            if (msg.value != deposit.amount) revert InvalidNativeDeposit();
        } else {
            deposit.token.safeTransferFrom(msg.sender, address(this), deposit.amount);
        }
    }

    /**
     * @dev Opens a new order by initializing its state.
     * @param orderData Order data to open.
     */
    function _openOrder(SolverNet.Order memory orderData) internal returns (ResolvedCrossChainOrder memory resolved) {
        resolved = _resolve(orderData);
        bytes32 id = _incrementId();

        _orderHeader[id] = orderData.header;
        _orderDeposit[id] = orderData.deposit;
        for (uint256 i; i < orderData.calls.length; ++i) {
            _orderCalls[id].push(orderData.calls[i]);
        }
        for (uint256 i; i < orderData.expenses.length; ++i) {
            _orderExpenses[id].push(orderData.expenses[i]);
        }

        _upsertOrder(id, Status.Pending, msg.sender);

        return resolved;
    }

    /**
     * @dev Transfer deposit to recipient. Used for both refunds and claims.
     * @param id ID of the order.
     * @param to Address to send deposits to.
     */
    function _transferDeposit(bytes32 id, address to) internal {
        SolverNet.Deposit memory deposit = _orderDeposit[id];

        if (deposit.amount > 0) {
            if (deposit.token == address(0)) to.safeTransferETH(deposit.amount);
            else deposit.token.safeTransfer(to, deposit.amount);
        }
    }

    /**
     * @dev Update or insert order state by id.
     * @param id        ID of the order.
     * @param status    Status to upsert.
     * @param updatedBy Address updating the order, only written to state if status is Accepted or Filled.
     */
    function _upsertOrder(bytes32 id, Status status, address updatedBy) internal {
        OrderState memory state = _orderState[id];

        state.status = status;
        state.timestamp = uint32(block.timestamp);
        if (status == Status.Accepted) state.claimant = updatedBy;
        if (status == Status.Filled && state.claimant == address(0)) state.claimant = updatedBy;

        _orderState[id] = state;
        _latestOrderIdByStatus[status] = id;
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
        SolverNet.Header memory header = _orderHeader[orderId];
        SolverNet.Call[] memory calls = _orderCalls[orderId];
        SolverNet.Expense[] memory expenses = _orderExpenses[orderId];

        SolverNet.FillOriginData memory fillOriginData = SolverNet.FillOriginData({
            srcChainId: uint64(block.chainid),
            destChainId: header.destChainId,
            fillDeadline: header.fillDeadline,
            calls: calls,
            expenses: expenses
        });

        return keccak256(abi.encode(orderId, abi.encode(fillOriginData)));
    }
}
