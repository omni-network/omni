// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.24;

import { ReentrancyGuard } from "solady/src/utils/ReentrancyGuard.sol";
import { IOriginSettler, IERC7683 } from "../../src/erc7683/IOriginSettler.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { SolverNet } from "../../src/lib/SolverNet.sol";
import { AddrUtils } from "../../src/lib/AddrUtils.sol";

contract MockSolverNetInbox is ReentrancyGuard, IOriginSettler {
    using SafeTransferLib for address;
    using AddrUtils for address;

    error InvalidOrderTypehash();
    error InvalidOrderData();
    error InvalidChainId();
    error InvalidFillDeadline();
    error InvalidCallTarget();
    error InvalidExpenseToken();
    error InvalidExpenseAmount();
    error InvalidNativeDeposit();

    enum Status {
        Invalid,
        Pending,
        Accepted,
        Rejected,
        Reverted,
        Filled,
        Claimed
    }

    struct OrderState {
        Status status;
        uint32 timestamp;
        address claimant;
    }

    bytes32 internal constant ORDERDATA_TYPEHASH = keccak256(
        "OrderData(address owner,uint64 destChainId,Deposit deposit,Call[] calls,Expense[] expenses)Deposit(address token,uint96 amount)Call(address target,bytes4 selector,uint256 value,bytes params)Expense(address spender,address token,uint96 amount)"
    );

    address public immutable outbox;

    uint256 internal _lastId;

    mapping(bytes32 id => SolverNet.Header) internal _orderHeader;
    mapping(bytes32 id => SolverNet.Deposit) internal _orderDeposit;
    mapping(bytes32 id => SolverNet.Call[]) internal _orderCalls;
    mapping(bytes32 id => SolverNet.Expense[]) internal _orderExpenses;
    mapping(bytes32 id => OrderState) internal _orderState;
    mapping(Status => bytes32 id) internal _latestOrderIdByStatus;

    constructor(address outbox_) {
        outbox = outbox_;
    }

    function getNextId() external view returns (bytes32) {
        return _nextId();
    }

    function resolve(OnchainCrossChainOrder calldata order) external view returns (ResolvedCrossChainOrder memory) {
        SolverNet.Order memory orderData = _validate(order);
        return _resolve(orderData, _nextId());
    }

    function open(OnchainCrossChainOrder calldata order) external payable nonReentrant {
        SolverNet.Order memory orderData = _validate(order);
        _processDeposit(orderData.deposit);
        ResolvedCrossChainOrder memory resolved = _openOrder(orderData);

        emit Open(resolved.orderId, resolved);
    }

    function _nextId() internal view returns (bytes32) {
        return bytes32(_lastId + 1);
    }

    function _incrementId() internal returns (bytes32) {
        return bytes32(++_lastId);
    }

    function _validate(OnchainCrossChainOrder calldata order) internal view returns (SolverNet.Order memory) {
        if (order.fillDeadline < block.timestamp && order.fillDeadline != 0) revert InvalidFillDeadline();
        if (order.orderDataType != ORDERDATA_TYPEHASH) revert InvalidOrderTypehash();
        if (order.orderData.length == 0) revert InvalidOrderData();

        SolverNet.OrderData memory orderData = abi.decode(order.orderData, (SolverNet.OrderData));

        if (orderData.owner == address(0)) orderData.owner = msg.sender;
        if (orderData.destChainId == 0 || orderData.destChainId == block.chainid) revert InvalidChainId();

        SolverNet.Header memory header = SolverNet.Header({
            owner: orderData.owner,
            destChainId: orderData.destChainId,
            fillDeadline: order.fillDeadline
        });

        SolverNet.Call[] memory calls = orderData.calls;
        for (uint256 i; i < calls.length; ++i) {
            SolverNet.Call memory call = calls[i];
            if (call.target == address(0)) revert InvalidCallTarget();
        }

        SolverNet.Expense[] memory expenses = orderData.expenses;
        for (uint256 i; i < expenses.length; ++i) {
            if (expenses[i].token == address(0)) revert InvalidExpenseToken();
            if (expenses[i].amount == 0) revert InvalidExpenseAmount();
        }

        return SolverNet.Order({ header: header, calls: calls, deposit: orderData.deposit, expenses: expenses });
    }

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
                recipient: outbox.toBytes32(),
                chainId: header.destChainId
            });
        }
        if (totalNativeValue > 0) {
            maxSpent[expenses.length] = IERC7683.Output({
                token: bytes32(0),
                amount: totalNativeValue,
                recipient: outbox.toBytes32(),
                chainId: header.destChainId
            });
        }

        return maxSpent;
    }

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
            destinationSettler: outbox.toBytes32(),
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

    function _processDeposit(SolverNet.Deposit memory deposit) internal {
        if (deposit.token == address(0)) {
            if (msg.value != deposit.amount) revert InvalidNativeDeposit();
        } else {
            deposit.token.safeTransferFrom(msg.sender, address(this), deposit.amount);
        }
    }

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

        _upsertOrder(id, Status.Pending, msg.sender);

        return resolved;
    }

    function _upsertOrder(bytes32 id, Status status, address updatedBy) internal {
        OrderState memory state = _orderState[id];

        state.status = status;
        state.timestamp = uint32(block.timestamp);
        if (status == Status.Accepted) state.claimant = updatedBy;
        if (status == Status.Filled && state.claimant == address(0)) state.claimant = updatedBy;

        _orderState[id] = state;
        _latestOrderIdByStatus[status] = id;
    }
}
