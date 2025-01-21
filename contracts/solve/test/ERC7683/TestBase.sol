// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { SolverNetInbox } from "src/ERC7683/SolverNetInbox.sol";
import { SolverNetOutbox } from "src/ERC7683/SolverNetOutbox.sol";

import { MockToken } from "test/utils/MockToken.sol";
import { MockVault } from "test/utils/MockVault.sol";
import { MockMultiTokenVault } from "test/utils/MockMultiTokenVault.sol";
import { MockPortal } from "core/test/utils/MockPortal.sol";

import { IERC7683 } from "src/ERC7683/interfaces/IERC7683.sol";
import { ISolverNet } from "src/ERC7683/interfaces/ISolverNet.sol";
import { ISolverNetInbox } from "src/ERC7683/interfaces/ISolverNetInbox.sol";
import { ISolverNetOutbox } from "src/ERC7683/interfaces/ISolverNetOutbox.sol";

import { Test, console2 } from "forge-std/Test.sol";
import { AddrUtils } from "src/ERC7683/lib/AddrUtils.sol";

/**
 * @title TestBase
 * @dev Shared test utils / fixtures.
 */
contract TestBase is Test {
    using AddrUtils for address;
    using AddrUtils for bytes32;

    SolverNetInbox inbox;
    SolverNetOutbox outbox;

    MockToken token1;
    MockToken token2;
    MockToken token3;
    MockToken token4;

    MockVault vault;
    MockMultiTokenVault multiTokenVault;

    MockPortal portal;

    uint64 srcChainId = 1;
    uint64 destChainId = 2;

    address user = makeAddr("user");
    address solver = makeAddr("solver");
    address proxyAdmin = makeAddr("proxy-admin-owner");

    bytes32 internal constant ORDER_DATA_TYPEHASH = keccak256(
        "OrderData(Call call,Deposit[] deposits)Call(uint64 chainId,bytes32 target,uint256 value,bytes data,TokenExpense[] expenses)TokenExpense(bytes32 token,bytes32 spender,uint256 amount)Deposit(bytes32 token,uint256 amount)"
    );

    modifier prankUser() {
        vm.startPrank(user);
        _;
        vm.stopPrank();
    }

    function setUp() public virtual {
        token1 = new MockToken();
        token2 = new MockToken();
        token3 = new MockToken();
        token4 = new MockToken();

        vault = new MockVault(address(token2));
        multiTokenVault = new MockMultiTokenVault();
        portal = new MockPortal();

        inbox = deploySolverNetInbox();
        outbox = deploySolverNetOutbox();
        initializeInbox();
        initializeOutbox();
    }

    function getVaultCalldata(address addr, uint256 amount) internal pure returns (bytes memory) {
        return abi.encodeCall(MockVault.deposit, (addr, amount));
    }

    function getMultiTokenVaultCalldata(address addr, ISolverNet.TokenExpense[] memory expenses, uint256 nativeAmount)
        internal
        pure
        returns (bytes memory)
    {
        address[] memory tokens = new address[](nativeAmount > 0 ? expenses.length + 1 : expenses.length);
        uint256[] memory amounts = new uint256[](nativeAmount > 0 ? expenses.length + 1 : expenses.length);

        for (uint256 i; i < expenses.length; ++i) {
            tokens[i] = expenses[i].token.toAddress();
            amounts[i] = expenses[i].amount;
        }

        if (nativeAmount > 0) {
            tokens[tokens.length - 1] = address(0);
            amounts[amounts.length - 1] = nativeAmount;
        }

        return abi.encodeCall(MockMultiTokenVault.deposit, (addr, tokens, amounts));
    }

    function getOrderDataBytes(
        uint64 chainId,
        bytes32 target,
        uint256 value,
        bytes memory data,
        ISolverNet.TokenExpense[] memory expenses,
        ISolverNet.Deposit[] memory deposits
    ) internal pure returns (bytes memory) {
        return abi.encode(
            ISolverNet.OrderData({ call: ISolverNet.Call(chainId, target, value, data, expenses), deposits: deposits })
        );
    }

    /**
     * @dev Generate a random order for a vault deposit.
     *      srcChainId = 1, destChainId = 2, amount = 1-1000
     *      token1 deposited into inbox on srcChain, token2 deposited into vault on destChain
     */
    function randOrder() internal returns (IERC7683.OnchainCrossChainOrder memory) {
        uint256 rand = vm.randomUint(1, 1000);

        ISolverNet.Deposit[] memory deposits = new ISolverNet.Deposit[](1);
        deposits[0] = ISolverNet.Deposit({ token: address(token1).toBytes32(), amount: rand * 1 ether });

        ISolverNet.TokenExpense[] memory expenses = new ISolverNet.TokenExpense[](1);
        expenses[0] = ISolverNet.TokenExpense({
            token: address(token2).toBytes32(),
            spender: address(vault).toBytes32(),
            amount: rand * 1 ether
        });

        ISolverNet.Call memory call = ISolverNet.Call({
            chainId: destChainId,
            target: address(vault).toBytes32(),
            value: 0,
            data: getVaultCalldata(user, rand * 1 ether),
            expenses: expenses
        });

        ISolverNet.OrderData memory orderData = ISolverNet.OrderData({ call: call, deposits: deposits });

        return IERC7683.OnchainCrossChainOrder({
            fillDeadline: uint32(block.timestamp + 1 minutes),
            orderDataType: ORDER_DATA_TYPEHASH,
            orderData: abi.encode(orderData)
        });
    }

    /**
     * @dev Generate a random order with native token deposits.
     *      srcChainId = 1, destChainId = 2, amount = 1-1000
     *      native token deposited into inbox on srcChain, token2 deposited into vault on destChain
     */
    function randNativeOrder() internal returns (IERC7683.OnchainCrossChainOrder memory) {
        uint256 rand = vm.randomUint(1, 1000);

        ISolverNet.Deposit[] memory deposits = new ISolverNet.Deposit[](1);
        deposits[0] = ISolverNet.Deposit({ token: bytes32(0), amount: rand * 1 ether });

        ISolverNet.TokenExpense[] memory expenses = new ISolverNet.TokenExpense[](1);
        expenses[0] = ISolverNet.TokenExpense({
            token: address(token2).toBytes32(),
            spender: address(vault).toBytes32(),
            amount: rand * 1 ether
        });

        ISolverNet.Call memory call = ISolverNet.Call({
            chainId: destChainId,
            target: address(vault).toBytes32(),
            value: 0,
            data: getVaultCalldata(user, rand * 1 ether),
            expenses: expenses
        });

        ISolverNet.OrderData memory orderData = ISolverNet.OrderData({ call: call, deposits: deposits });

        return IERC7683.OnchainCrossChainOrder({
            fillDeadline: uint32(block.timestamp + 1 minutes),
            orderDataType: ORDER_DATA_TYPEHASH,
            orderData: abi.encode(orderData)
        });
    }

    function randMultiTokenOrder(address[] memory srcDeposits, address[] memory destDeposits)
        internal
        returns (IERC7683.OnchainCrossChainOrder memory)
    {
        uint256 nativeExpense;
        for (uint256 i; i < destDeposits.length; ++i) {
            if (destDeposits[i] == address(0)) {
                nativeExpense = vm.randomUint(1, 1000) * 1 ether;
                break;
            }
        }

        ISolverNet.Deposit[] memory deposits = new ISolverNet.Deposit[](srcDeposits.length);
        for (uint256 i; i < srcDeposits.length; ++i) {
            uint256 rand = vm.randomUint(1, 1000);
            deposits[i] = ISolverNet.Deposit({ token: address(srcDeposits[i]).toBytes32(), amount: rand * 1 ether });
        }

        ISolverNet.TokenExpense[] memory expenses =
            new ISolverNet.TokenExpense[](nativeExpense > 0 ? destDeposits.length - 1 : destDeposits.length);
        bool nativeProcessed;
        for (uint256 i; i < destDeposits.length; ++i) {
            if (destDeposits[i] == address(0)) {
                nativeProcessed = true;
                continue;
            }
            uint256 rand = vm.randomUint(1, 1000);
            expenses[nativeProcessed ? i - 1 : i] = ISolverNet.TokenExpense({
                token: address(destDeposits[i]).toBytes32(),
                spender: address(multiTokenVault).toBytes32(),
                amount: rand * 1 ether
            });
        }

        ISolverNet.Call memory call = ISolverNet.Call({
            chainId: destChainId,
            target: address(multiTokenVault).toBytes32(),
            value: nativeExpense,
            data: getMultiTokenVaultCalldata(user, expenses, nativeExpense),
            expenses: expenses
        });

        ISolverNet.OrderData memory orderData = ISolverNet.OrderData({ call: call, deposits: deposits });

        return IERC7683.OnchainCrossChainOrder({
            fillDeadline: uint32(block.timestamp + 1 minutes),
            orderDataType: ORDER_DATA_TYPEHASH,
            orderData: abi.encode(orderData)
        });
    }

    function fundUser(IERC7683.Output[] memory deposits) internal {
        vm.chainId(srcChainId);
        for (uint256 i; i < deposits.length; ++i) {
            address token = deposits[i].token.toAddress();
            if (token == address(0)) {
                vm.deal(user, deposits[i].amount);
                continue;
            }

            vm.startPrank(user);
            MockToken(token).approve(address(inbox), deposits[i].amount);
            MockToken(token).mint(user, deposits[i].amount);
            vm.stopPrank();
        }
    }

    function fundSolver(IERC7683.Output[] memory expenses) internal {
        vm.chainId(destChainId);
        for (uint256 i; i < expenses.length; ++i) {
            address token = expenses[i].token.toAddress();
            if (token == address(0)) {
                vm.deal(solver, expenses[i].amount);
                continue;
            }

            vm.startPrank(solver);
            MockToken(token).approve(address(outbox), expenses[i].amount);
            MockToken(token).mint(solver, expenses[i].amount);
            vm.stopPrank();
        }

        uint256 fillFee = outbox.fillFee(srcChainId);
        vm.deal(solver, address(solver).balance + fillFee);
    }

    function deploySolverNetInbox() internal returns (SolverNetInbox) {
        address impl = address(new SolverNetInbox());
        return SolverNetInbox(address(new TransparentUpgradeableProxy(impl, proxyAdmin, bytes(""))));
    }

    function deploySolverNetOutbox() internal returns (SolverNetOutbox) {
        address impl = address(new SolverNetOutbox());
        return SolverNetOutbox(address(new TransparentUpgradeableProxy(impl, proxyAdmin, bytes(""))));
    }

    // Separate initialization functions are necessary as proxy addresses must be known prior.
    function initializeInbox() internal {
        inbox.initialize(address(this), solver, address(portal), address(outbox));
    }

    // Separate initialization functions are necessary as proxy addresses must be known prior.
    function initializeOutbox() internal {
        outbox.initialize(address(this), solver, address(portal), address(inbox));
    }

    function fillHash(bytes32 orderId, bytes memory originData) internal pure returns (bytes32) {
        return keccak256(abi.encode(orderId, originData));
    }

    function assertResolved(
        address userAddr,
        bytes32 orderId,
        IERC7683.OnchainCrossChainOrder memory order,
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder
    ) internal view {
        ISolverNet.OrderData memory orderData = abi.decode(order.orderData, (ISolverNet.OrderData));
        ISolverNet.Call memory orderCall = orderData.call;
        ISolverNet.TokenExpense[] memory orderExpenses = orderCall.expenses;
        ISolverNet.Deposit[] memory orderDeposits = orderData.deposits;
        ISolverNet.FillOriginData memory fillOriginData =
            abi.decode(resolvedOrder.fillInstructions[0].originData, (ISolverNet.FillOriginData));

        assertEq(userAddr, resolvedOrder.user, "assertResolved: user");
        assertEq(srcChainId, resolvedOrder.originChainId, "assertResolved: origin chain id");
        assertEq(uint32(block.timestamp), resolvedOrder.openDeadline, "assertResolved: open deadline");
        assertEq(order.fillDeadline, resolvedOrder.fillDeadline, "assertResolved: fill deadline");
        assertEq(orderId, resolvedOrder.orderId, "assertResolved: order id");

        assertEq(
            orderCall.value > 0 ? orderExpenses.length + 1 : orderExpenses.length,
            resolvedOrder.maxSpent.length,
            "assertResolved: max spent length"
        );
        assertEq(orderExpenses.length, fillOriginData.call.expenses.length, "assertResolved: call expense length");
        for (uint256 i; i < orderExpenses.length; ++i) {
            assertEq(orderExpenses[i].token, resolvedOrder.maxSpent[i].token, "assertResolved: max spent token");
            assertEq(
                orderExpenses[i].token, fillOriginData.call.expenses[i].token, "assertResolved: call expense token"
            );
            assertEq(
                orderExpenses[i].spender,
                fillOriginData.call.expenses[i].spender,
                "assertResolved: call expense spender"
            );
            assertEq(orderExpenses[i].amount, resolvedOrder.maxSpent[i].amount, "assertResolved: max spent amount");
            assertEq(
                orderExpenses[i].amount, fillOriginData.call.expenses[i].amount, "assertResolved: call expense amount"
            );
            assertEq(
                address(outbox).toBytes32(), resolvedOrder.maxSpent[i].recipient, "assertResolved: max spent recipient"
            );
            assertEq(orderCall.chainId, resolvedOrder.maxSpent[i].chainId, "assertResolved: max spent chain id");
        }
        if (orderCall.value > 0) {
            assertEq(bytes32(0), resolvedOrder.maxSpent[orderExpenses.length].token, "assertResolved: max spent token");
            assertEq(
                orderCall.value, resolvedOrder.maxSpent[orderExpenses.length].amount, "assertResolved: max spent amount"
            );
            assertEq(
                address(outbox).toBytes32(),
                resolvedOrder.maxSpent[orderExpenses.length].recipient,
                "assertResolved: max spent recipient"
            );
            assertEq(
                orderCall.chainId,
                resolvedOrder.maxSpent[orderExpenses.length].chainId,
                "assertResolved: max spent chain id"
            );
        }

        assertEq(orderDeposits.length, resolvedOrder.minReceived.length, "assertResolved: min received length");
        for (uint256 i; i < orderDeposits.length; ++i) {
            assertEq(orderDeposits[i].token, resolvedOrder.minReceived[i].token, "assertResolved: min received token");
            assertEq(
                orderDeposits[i].amount, resolvedOrder.minReceived[i].amount, "assertResolved: min received amount"
            );
            assertEq(bytes32(0), resolvedOrder.minReceived[i].recipient, "assertResolved: min received recipient");
            assertEq(srcChainId, resolvedOrder.minReceived[i].chainId, "assertResolved: min received chain id");
        }

        assertEq(1, resolvedOrder.fillInstructions.length, "assertResolved: fill instructions length");
        assertEq(
            orderCall.chainId,
            resolvedOrder.fillInstructions[0].destinationChainId,
            "assertResolved: fill instructions chain id"
        );
        assertEq(
            address(outbox).toBytes32(),
            resolvedOrder.fillInstructions[0].destinationSettler,
            "assertResolved: fill instructions destination"
        );
        assertEq(
            keccak256(abi.encode(ISolverNet.FillOriginData({ srcChainId: srcChainId, call: orderCall }))),
            keccak256(resolvedOrder.fillInstructions[0].originData),
            "assertResolved: fill instructions origin data"
        );
    }

    function assertNullOrder(bytes32 orderId) internal view {
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder;
        ISolverNetInbox.OrderState memory state;
        ISolverNetInbox.StatusUpdate[] memory history;
        (resolvedOrder, state, history) = inbox.getOrder(orderId);

        assertEq(resolvedOrder.user, address(0), "null order: user");
        assertEq(resolvedOrder.originChainId, 0, "null order: originChainId");
        assertEq(resolvedOrder.openDeadline, 0, "null order: openDeadline");
        assertEq(resolvedOrder.fillDeadline, 0, "null order: fillDeadline");
        assertEq(resolvedOrder.orderId, bytes32(0), "null order: orderId");
        assertEq(resolvedOrder.minReceived.length, 0, "null order: minReceived");
        assertEq(resolvedOrder.maxSpent.length, 0, "null order: maxSpent");
        assertEq(resolvedOrder.fillInstructions.length, 0, "null order: fillInstructions");
        assertEq(uint8(state.status), uint8(ISolverNetInbox.Status.Invalid), "null order: status");
        assertEq(state.acceptedBy, address(0), "null order: acceptedBy");
        assertEq(history.length, 0, "null order: history");
        assertEq(token1.balanceOf(address(inbox)), 0, "null order: inbox token1 balance");
    }

    function assertOpenedOrder(bytes32 orderId) internal view {
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder;
        ISolverNetInbox.OrderState memory state;
        ISolverNetInbox.StatusUpdate[] memory history;
        (resolvedOrder, state, history) = inbox.getOrder(orderId);

        assertEq(resolvedOrder.user, user, "opened order: user");
        assertEq(resolvedOrder.originChainId, srcChainId, "opened order: originChainId");
        assertEq(resolvedOrder.openDeadline, uint32(block.timestamp), "opened order: openDeadline");
        assertEq(resolvedOrder.fillDeadline, uint32(block.timestamp + 1 minutes), "opened order: fillDeadline");
        assertEq(resolvedOrder.orderId, orderId, "opened order: orderId");
        assertEq(uint8(state.status), uint8(ISolverNetInbox.Status.Pending), "opened order: status");
        assertEq(state.acceptedBy, address(0), "opened order: acceptedBy");
        assertEq(history.length, 1, "opened order: history");
        assertEq(uint8(history[0].status), uint8(ISolverNetInbox.Status.Pending), "opened order: history[0].status");
        assertEq(history[0].timestamp, uint40(block.timestamp), "opened order: history[0].timestamp");
        assertEq(inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Pending), orderId, "opened order: latestOrderId");
        assertEq(
            token1.balanceOf(address(inbox)), resolvedOrder.minReceived[0].amount, "opened order: inbox token1 balance"
        );
    }

    function assertAcceptedOrder(bytes32 orderId) internal view {
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder;
        ISolverNetInbox.OrderState memory state;
        ISolverNetInbox.StatusUpdate[] memory history;
        (resolvedOrder, state, history) = inbox.getOrder(orderId);

        assertEq(resolvedOrder.user, user, "accepted order: user");
        assertEq(resolvedOrder.originChainId, srcChainId, "accepted order: originChainId");
        assertEq(resolvedOrder.openDeadline, uint32(block.timestamp), "accepted order: openDeadline");
        assertEq(resolvedOrder.fillDeadline, uint32(block.timestamp + 1 minutes), "accepted order: fillDeadline");
        assertEq(resolvedOrder.orderId, orderId, "accepted order: orderId");
        assertEq(uint8(state.status), uint8(ISolverNetInbox.Status.Accepted), "accepted order: status");
        assertEq(state.acceptedBy, solver, "accepted order: acceptedBy");
        assertEq(history.length, 2, "accepted order: history");
        assertEq(uint8(history[0].status), uint8(ISolverNetInbox.Status.Pending), "accepted order: history[0].status");
        assertEq(history[0].timestamp, uint40(block.timestamp), "accepted order: history[0].timestamp");
        assertEq(uint8(history[1].status), uint8(ISolverNetInbox.Status.Accepted), "accepted order: history[1].status");
        assertEq(history[1].timestamp, uint40(block.timestamp), "accepted order: history[1].timestamp");
        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Accepted), orderId, "accepted order: latestOrderId"
        );
        assertEq(
            token1.balanceOf(address(inbox)),
            resolvedOrder.minReceived[0].amount,
            "accepted order: inbox token1 balance"
        );
    }

    function assertVaultDeposit(bytes32 orderId) internal view {
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder;
        (resolvedOrder,,) = inbox.getOrder(orderId);

        uint256 amount = resolvedOrder.maxSpent[0].amount;
        assertEq(vault.balances(user), amount, "vault deposit: amount");
        assertEq(token2.balanceOf(address(vault)), amount, "vault deposit: vault token2 balance");
        assertEq(token2.balanceOf(address(outbox)), 0, "vault deposit: outbox token2 balance");
        assertEq(token2.balanceOf(solver), 0, "vault deposit: solver token2 balance");
    }

    function assertFulfilledOrder(bytes32 orderId) internal view {
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder;
        ISolverNetInbox.OrderState memory state;
        ISolverNetInbox.StatusUpdate[] memory history;
        (resolvedOrder, state, history) = inbox.getOrder(orderId);

        assertEq(resolvedOrder.user, user, "fulfilled order: user");
        assertEq(resolvedOrder.originChainId, srcChainId, "fulfilled order: originChainId");
        assertEq(resolvedOrder.openDeadline, uint32(block.timestamp), "fulfilled order: openDeadline");
        assertEq(resolvedOrder.fillDeadline, uint32(block.timestamp + 1 minutes), "fulfilled order: fillDeadline");
        assertEq(resolvedOrder.orderId, orderId, "fulfilled order: orderId");
        assertEq(uint8(state.status), uint8(ISolverNetInbox.Status.Filled), "fulfilled order: status");
        assertEq(state.acceptedBy, solver, "fulfilled order: acceptedBy");
        assertEq(history.length, 3, "fulfilled order: history");
        assertEq(uint8(history[0].status), uint8(ISolverNetInbox.Status.Pending), "fulfilled order: history[0].status");
        assertEq(history[0].timestamp, uint40(block.timestamp), "fulfilled order: history[0].timestamp");
        assertEq(uint8(history[1].status), uint8(ISolverNetInbox.Status.Accepted), "fulfilled order: history[1].status");
        assertEq(history[1].timestamp, uint40(block.timestamp), "fulfilled order: history[1].timestamp");
        assertEq(uint8(history[2].status), uint8(ISolverNetInbox.Status.Filled), "fulfilled order: history[2].status");
        assertEq(history[2].timestamp, uint40(block.timestamp), "fulfilled order: history[2].timestamp");
        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Filled), orderId, "fulfilled order: latestOrderId"
        );
        assertEq(
            token1.balanceOf(address(inbox)),
            resolvedOrder.minReceived[0].amount,
            "fulfilled order: inbox token1 balance"
        );
        assertEq(token1.balanceOf(solver), 0, "fulfilled order: solver token1 balance");
    }

    function assertClaimedOrder(bytes32 orderId) internal view {
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder;
        ISolverNetInbox.OrderState memory state;
        ISolverNetInbox.StatusUpdate[] memory history;
        (resolvedOrder, state, history) = inbox.getOrder(orderId);

        assertEq(resolvedOrder.user, user, "accepted order: user");
        assertEq(resolvedOrder.originChainId, srcChainId, "accepted order: originChainId");
        assertEq(resolvedOrder.openDeadline, uint32(block.timestamp), "accepted order: openDeadline");
        assertEq(resolvedOrder.fillDeadline, uint32(block.timestamp + 1 minutes), "accepted order: fillDeadline");
        assertEq(resolvedOrder.orderId, orderId, "accepted order: orderId");
        assertEq(uint8(state.status), uint8(ISolverNetInbox.Status.Claimed), "accepted order: status");
        assertEq(state.acceptedBy, solver, "accepted order: acceptedBy");
        assertEq(history.length, 4, "accepted order: history");
        assertEq(uint8(history[0].status), uint8(ISolverNetInbox.Status.Pending), "accepted order: history[0].status");
        assertEq(history[0].timestamp, uint40(block.timestamp), "accepted order: history[0].timestamp");
        assertEq(uint8(history[1].status), uint8(ISolverNetInbox.Status.Accepted), "accepted order: history[1].status");
        assertEq(history[1].timestamp, uint40(block.timestamp), "accepted order: history[1].timestamp");
        assertEq(uint8(history[2].status), uint8(ISolverNetInbox.Status.Filled), "accepted order: history[2].status");
        assertEq(history[2].timestamp, uint40(block.timestamp), "accepted order: history[2].timestamp");
        assertEq(uint8(history[3].status), uint8(ISolverNetInbox.Status.Claimed), "accepted order: history[3].status");
        assertEq(history[3].timestamp, uint40(block.timestamp), "accepted order: history[3].timestamp");
        assertEq(
            inbox.getLatestOrderIdByStatus(ISolverNetInbox.Status.Claimed), orderId, "accepted order: latestOrderId"
        );
        assertEq(token1.balanceOf(solver), resolvedOrder.minReceived[0].amount, "claimed order: solver token1 balance");
        assertEq(token1.balanceOf(address(inbox)), 0, "claimed order: inbox token1 balance");
    }
}
