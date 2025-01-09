// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { SolverNetInbox } from "src/ERC7683/SolverNetInbox.sol";
import { SolverNetOutbox } from "src/ERC7683/SolverNetOutbox.sol";

import { IERC7683 } from "src/ERC7683/interfaces/IERC7683.sol";
import { ISolverNet } from "src/ERC7683/interfaces/ISolverNet.sol";
import { ISolverNetInbox } from "src/ERC7683/interfaces/ISolverNetInbox.sol";

import { Test } from "forge-std/Test.sol";
import { MockToken } from "test/utils/MockToken.sol";
import { MockVault } from "test/utils/MockVault.sol";
import { MockPortal } from "core/test/utils/MockPortal.sol";

/**
 * @title TestBase
 * @dev Shared test utils / fixtures.
 */
contract TestBase is Test {
    SolverNetInbox inbox;
    SolverNetOutbox outbox;

    MockToken token1;
    MockToken token2;
    MockVault vault;
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

    function setUp() public {
        token1 = new MockToken();
        token2 = new MockToken();
        vault = new MockVault(address(token2));
        portal = new MockPortal();
        inbox = deploySolverNetInbox();
        outbox = deploySolverNetOutbox();
        initializeInbox();
        initializeOutbox();
        allowCall(address(vault), vault.deposit.selector);
    }

    /**
     * @dev Generate a random order for a vault deposit.
     *      srcChainId = 1, destChainId = 2, amount = 1-1000
     *      token1 deposited into inbox on srcChain, token2 deposited into vault on destChain
     */
    function randOrder() internal returns (IERC7683.OnchainCrossChainOrder memory) {
        uint256 rand = vm.randomUint(1, 1000);

        ISolverNet.TokenExpense[] memory expenses = new ISolverNet.TokenExpense[](1);
        expenses[0] = ISolverNet.TokenExpense({
            token: addressToBytes32(address(token2)),
            spender: addressToBytes32(address(vault)),
            amount: rand * 1 ether
        });

        ISolverNet.Call memory call = ISolverNet.Call({
            chainId: destChainId,
            target: addressToBytes32(address(vault)),
            value: 0,
            data: abi.encodeCall(MockVault.deposit, (user, rand * 1 ether)),
            expenses: expenses
        });

        ISolverNet.Deposit[] memory deposits = new ISolverNet.Deposit[](1);
        deposits[0] = ISolverNet.Deposit({ token: addressToBytes32(address(token1)), amount: rand * 1 ether });

        ISolverNet.OrderData memory orderData = ISolverNet.OrderData({ call: call, deposits: deposits });

        return IERC7683.OnchainCrossChainOrder({
            fillDeadline: uint32(block.timestamp + 1 minutes),
            orderDataType: ORDER_DATA_TYPEHASH,
            orderData: abi.encode(orderData)
        });
    }

    function mintAndApprove(IERC7683.Output[] memory deposits, IERC7683.Output[] memory expenses) internal {
        for (uint256 i; i < deposits.length; ++i) {
            vm.startPrank(user);
            MockToken(bytes32ToAddress(deposits[i].token)).approve(address(inbox), deposits[i].amount);
            MockToken(bytes32ToAddress(deposits[i].token)).mint(user, deposits[i].amount);
            vm.stopPrank();
        }

        for (uint256 i; i < expenses.length; ++i) {
            vm.startPrank(solver);
            MockToken(bytes32ToAddress(expenses[i].token)).approve(address(outbox), expenses[i].amount);
            MockToken(bytes32ToAddress(expenses[i].token)).mint(solver, expenses[i].amount);
            vm.stopPrank();
        }
    }

    function deploySolverNetInbox() internal returns (SolverNetInbox) {
        address impl = address(new SolverNetInbox());
        return SolverNetInbox(address(new TransparentUpgradeableProxy(impl, proxyAdmin, bytes(""))));
    }

    function deploySolverNetOutbox() internal returns (SolverNetOutbox) {
        address impl = address(new SolverNetOutbox());
        return SolverNetOutbox(address(new TransparentUpgradeableProxy(impl, proxyAdmin, bytes(""))));
    }

    // Seperate initialization functions are necessary as proxy addresses must be known prior.
    function initializeInbox() internal {
        inbox.initialize(address(this), solver, address(portal), address(outbox));
    }

    // Seperate initialization functions are necessary as proxy addresses must be known prior.
    function initializeOutbox() internal {
        outbox.initialize(address(this), solver, address(portal), address(inbox));
    }

    function allowCall(address target, bytes4 selector) internal {
        outbox.setAllowedCall(target, selector, true);
    }

    function fillHash(bytes32 orderId, bytes memory originData) internal pure returns (bytes32) {
        return keccak256(abi.encode(orderId, originData));
    }

    function addressToBytes32(address a) internal pure returns (bytes32) {
        return bytes32(uint256(uint160(a)));
    }

    function bytes32ToAddress(bytes32 b) internal pure returns (address) {
        return address(uint160(uint256(b)));
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
