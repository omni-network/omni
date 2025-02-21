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
     * @notice Buffer for closing orders after fill deadline to give Omni Core relayer time to act.
     */
    uint256 internal constant CLOSE_BUFFER = 6 hours;

    /**
     * @notice Role for solvers.
     * @dev _ROLE_0 evaluates to '1'.
     */
    uint256 internal constant SOLVER = _ROLE_0;

    /**
     * @notice Typehash for the OrderData struct.
     */
    bytes32 internal constant ORDERDATA_TYPEHASH = keccak256(
        "OrderData(address owner,uint64 destChainId,Deposit deposit,Call[] calls,TokenExpense[] expenses)Deposit(address token,uint96 amount)Call(address target,bytes4 selector,uint256 value,bytes params)TokenExpense(address spender,address token,uint96 amount)"
    );

    /**
     * @notice Key for pausing the `open` function.
     */
    bytes32 internal constant OPEN = keccak256("OPEN");

    /**
     * @notice Key for pausing the `close` function.
     */
    bytes32 internal constant CLOSE = keccak256("CLOSE");

    uint8 internal constant NONE_PAUSED = 0;
    uint8 internal constant OPEN_PAUSED = 1;
    uint8 internal constant CLOSE_PAUSED = 2;
    uint8 internal constant ALL_PAUSED = 3;

    /**
     * @dev Counter for generating unique order IDs. Incremented each time a new order is created.
     */
    uint248 internal _lastId;

    /**
     * @notice Pause state.
     * @dev 0 = no pause, 1 = open paused, 2 = close paused, 3 = all paused.
     */
    uint8 public pauseState;

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
    mapping(bytes32 id => SolverNet.TokenExpense[]) internal _orderExpenses;

    /**
     * @notice Map order ID to order parameters.
     */
    mapping(bytes32 id => OrderState) internal _orderState;

    /**
     * @notice Map status to latest order ID.
     */
    mapping(Status => bytes32 id) internal _latestOrderIdByStatus;

    /**
     * @notice Modifier to ensure contract functions are not paused.
     */
    modifier whenNotPaused(bytes32 pauseKey) {
        uint8 _pauseState = pauseState;
        if (_pauseState != NONE_PAUSED) {
            if (_pauseState == OPEN_PAUSED && pauseKey == OPEN) revert IsPaused();
            if (_pauseState == CLOSE_PAUSED && pauseKey == CLOSE) revert IsPaused();
            if (_pauseState == ALL_PAUSED) revert AllPaused();
        }
        _;
    }

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
     * @notice Pause the `open` function, preventing new orders from being opened.
     * @dev Cannot override ALL_PAUSED state.
     * @param pause True to pause, false to unpause.
     */
    function pauseOpen(bool pause) external onlyOwnerOrRoles(SOLVER) {
        _setPauseState(OPEN, pause);
    }

    /**
     * @notice Pause the `close` function, preventing orders from being closed by users.
     * @dev `close` should only be paused if the Omni Core relayer is not available.
     * @dev Cannot override ALL_PAUSED state.
     * @param pause True to pause, false to unpause.
     */
    function pauseClose(bool pause) external onlyOwnerOrRoles(SOLVER) {
        _setPauseState(CLOSE, pause);
    }

    /**
     * @notice Pause open and close functions.
     * @dev Can override OPEN_PAUSED or CLOSE_PAUSED states.
     * @param pause True to pause, false to unpause.
     */
    function pauseAll(bool pause) external onlyOwnerOrRoles(SOLVER) {
        pause ? pauseState = ALL_PAUSED : pauseState = NONE_PAUSED;
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
        return (_resolve(orderData, id), _orderState[id]);
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
        return _resolve(orderData, _nextId());
    }

    /**
     * @notice Open an order to execute a call on another chain, backed by deposits.
     * @dev Token deposits are transferred from msg.sender to this inbox.
     * @param order OnchainCrossChainOrder to open.
     */
    function open(OnchainCrossChainOrder calldata order) external payable whenNotPaused(OPEN) nonReentrant {
        SolverNet.Order memory orderData = _validate(order);
        _processDeposit(orderData.deposit);
        ResolvedCrossChainOrder memory resolved = _openOrder(orderData);

        emit Open(resolved.orderId, resolved);
    }

    /**
     * @notice Reject an open order and refund deposits.
     * @dev Only a whitelisted solver can reject.
     * @param id     ID of the order.
     * @param reason Reason code for rejection.
     */
    function reject(bytes32 id, uint8 reason) external onlyRoles(SOLVER) nonReentrant {
        OrderState memory state = _orderState[id];

        if (reason == 0) revert InvalidReason();
        if (state.status != Status.Pending) revert OrderNotPending();

        _upsertOrder(id, Status.Rejected, reason, msg.sender);
        _transferDeposit(id, _orderHeader[id].owner);

        emit Rejected(id, msg.sender, reason);
    }

    /**
     * @notice Close order and refund deposits after fill deadline has elapsed.
     * @dev Only order initiator can close.
     * @param id ID of the order.
     */
    function close(bytes32 id) external whenNotPaused(CLOSE) nonReentrant {
        OrderState memory state = _orderState[id];
        SolverNet.Header memory header = _orderHeader[id];

        if (state.status != Status.Pending) revert OrderNotPending();
        if (header.owner != msg.sender) revert Unauthorized();
        if (header.fillDeadline + CLOSE_BUFFER >= block.timestamp) revert OrderStillValid();

        _upsertOrder(id, Status.Closed, 0, msg.sender);
        _transferDeposit(id, header.owner);

        emit Closed(id);
    }

    /**
     * @notice Fill an order.
     * @dev Only callable by the outbox.
     * @param id         ID of the order.
     * @param fillHash   Hash of fill instructions origin data.
     * @param creditedTo Address deposits are credited to, provided by the filler.
     */
    function markFilled(bytes32 id, bytes32 fillHash, address creditedTo) external xrecv nonReentrant {
        SolverNet.Header memory header = _orderHeader[id];
        OrderState memory state = _orderState[id];

        if (state.status != Status.Pending) revert OrderNotPending();
        if (xmsg.sourceChainId != header.destChainId) revert WrongSourceChain();
        if (xmsg.sender != _outboxes[xmsg.sourceChainId]) revert Unauthorized();

        // Ensure reported fill hash matches origin data
        if (fillHash != _fillHash(id)) {
            revert WrongFillHash();
        }

        _upsertOrder(id, Status.Filled, 0, creditedTo);
        emit Filled(id, fillHash, creditedTo);
    }

    /**
     * @notice Claim deposits for a filled order.
     * @param id ID of the order.
     * @param to Address to send deposits to.
     */
    function claim(bytes32 id, address to) external nonReentrant {
        OrderState memory state = _orderState[id];

        if (state.status != Status.Filled) revert OrderNotFilled();
        if (state.updatedBy != msg.sender) revert Unauthorized();

        _upsertOrder(id, Status.Claimed, 0, msg.sender);
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
        if (order.fillDeadline <= block.timestamp) revert InvalidFillDeadline();
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
        if (calls.length == 0) revert InvalidMissingCalls();
        for (uint256 i; i < calls.length; ++i) {
            SolverNet.Call memory call = calls[i];
            if (call.target == address(0)) revert InvalidCallTarget();
        }

        // Validate SolverNet.OrderData.Expenses
        SolverNet.TokenExpense[] memory expenses = orderData.expenses;
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
        SolverNet.TokenExpense[] memory expenses = orderData.expenses;

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
        SolverNet.TokenExpense[] memory expenses = orderData.expenses;

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
    function _resolve(SolverNet.Order memory orderData, bytes32 id)
        internal
        view
        returns (ResolvedCrossChainOrder memory)
    {
        SolverNet.Header memory header = orderData.header;

        IERC7683.Output[] memory maxSpent = _deriveMaxSpent(orderData);
        IERC7683.Output[] memory minReceived = _deriveMinReceived(orderData);
        IERC7683.FillInstruction[] memory fillInstructions = _deriveFillInstructions(orderData);

        return ResolvedCrossChainOrder({
            user: header.owner,
            originChainId: block.chainid,
            openDeadline: 0,
            fillDeadline: header.fillDeadline,
            orderId: id,
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
        bytes32 id = _incrementId();
        resolved = _resolve(orderData, id);

        _orderHeader[id] = orderData.header;
        _orderDeposit[id] = orderData.deposit;
        for (uint256 i; i < orderData.calls.length; ++i) {
            _orderCalls[id].push(orderData.calls[i]);
        }
        for (uint256 i; i < orderData.expenses.length; ++i) {
            _orderExpenses[id].push(orderData.expenses[i]);
        }

        _upsertOrder(id, Status.Pending, 0, msg.sender);

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
     * @param id           ID of the order.
     * @param status       Status to upsert.
     * @param rejectReason Reason code for rejecting the order, if rejected.
     * @param updatedBy    Address updating the order.
     */
    function _upsertOrder(bytes32 id, Status status, uint8 rejectReason, address updatedBy) internal {
        uint8 _rejectReason = _orderState[id].rejectReason;
        _orderState[id] = OrderState({
            status: status,
            rejectReason: rejectReason > 0 ? rejectReason : _rejectReason,
            timestamp: uint32(block.timestamp),
            updatedBy: updatedBy
        });
        _latestOrderIdByStatus[status] = id;
    }

    /**
     * @dev Return the next order ID.
     */
    function _nextId() internal view returns (bytes32) {
        return bytes32(uint256(_lastId + 1));
    }

    /**
     * @dev Increment and return the next order ID.
     */
    function _incrementId() internal returns (bytes32) {
        return bytes32(uint256(++_lastId));
    }

    /**
     * @dev Returns call hash. Used to discern fulfillment.
     * @param orderId ID of the order.
     */
    function _fillHash(bytes32 orderId) internal view returns (bytes32) {
        SolverNet.Header memory header = _orderHeader[orderId];
        SolverNet.Call[] memory calls = _orderCalls[orderId];
        SolverNet.TokenExpense[] memory expenses = _orderExpenses[orderId];

        SolverNet.FillOriginData memory fillOriginData = SolverNet.FillOriginData({
            srcChainId: uint64(block.chainid),
            destChainId: header.destChainId,
            fillDeadline: header.fillDeadline,
            calls: calls,
            expenses: expenses
        });

        return keccak256(abi.encode(orderId, abi.encode(fillOriginData)));
    }

    /**
     * @notice Pause the `open` or `close` function
     * @dev Cannot override ALL_PAUSED state
     * @param key OPEN or CLOSE pause key
     * @param pause True to pause, false to unpause
     */
    function _setPauseState(bytes32 key, bool pause) internal {
        uint8 _pauseState = pauseState;
        if (_pauseState == ALL_PAUSED) revert AllPaused();

        uint8 targetState = key == OPEN ? OPEN_PAUSED : CLOSE_PAUSED;
        if (pause ? _pauseState == targetState : _pauseState != targetState) revert IsPaused();

        pauseState = pause ? targetState : NONE_PAUSED;
    }
}
