// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { SolverNetInbox } from "src/_optimized/SolverNetInbox.sol";
import { SolverNetOutbox } from "src/_optimized/SolverNetOutbox.sol";

import { ERC20 } from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import { MockERC20 } from "test/utils/MockERC20.sol";
import { MockVault } from "test/utils/MockVault.sol";
import { MockPortal } from "core/test/utils/MockPortal.sol";

import { IERC7683 } from "src/erc7683/IERC7683.sol";
import { SolverNet } from "src/_optimized/lib/SolverNet.sol";

import { Test, console2 } from "forge-std/Test.sol";
import { AddrUtils } from "src/lib/AddrUtils.sol";

contract SolverNet_Bench_Test is Test {
    using AddrUtils for address;
    using AddrUtils for bytes32;

    SolverNetInbox inbox;
    SolverNetOutbox outbox;

    MockERC20 token1;
    MockERC20 token2;

    MockVault nativeVault;
    MockVault erc20Vault;

    MockPortal portal;

    uint64 srcChainId = 1;
    uint64 destChainId = 2;

    address user = makeAddr("user");
    address solver = makeAddr("solver");
    address proxyAdmin = makeAddr("proxy-admin-owner");

    uint96 defaultAmount = 100 ether;

    bytes32 internal constant ORDER_DATA_TYPEHASH = keccak256(
        "OrderData(Header header,Deposit deposit,Call[] calls,Expense[] expenses)Header(address owner,uint64 destChainId,uint32 fillDeadline)Call(address target,bytes4 selector,uint256 value,bytes params)Deposit(address token,uint96 amount)Expense(address spender,address token,uint96 amount)"
    );

    function setUp() public {
        token1 = new MockERC20("Token 1", "TKN1");
        token2 = new MockERC20("Token 2", "TKN2");

        nativeVault = new MockVault(address(0));
        erc20Vault = new MockVault(address(token2));
        portal = new MockPortal();

        inbox = deploySolverNetInbox();
        outbox = deploySolverNetOutbox();
        initializeInbox();
        initializeOutbox();

        vm.chainId(srcChainId);
    }

    function deploySolverNetInbox() internal returns (SolverNetInbox) {
        address impl = address(new SolverNetInbox());
        return SolverNetInbox(address(new TransparentUpgradeableProxy(impl, proxyAdmin, bytes(""))));
    }

    function deploySolverNetOutbox() internal returns (SolverNetOutbox) {
        address impl = address(new SolverNetOutbox());
        return SolverNetOutbox(address(new TransparentUpgradeableProxy(impl, proxyAdmin, bytes(""))));
    }

    function initializeInbox() internal {
        inbox.initialize(address(this), solver, address(portal));

        uint64[] memory chainIds = new uint64[](1);
        chainIds[0] = destChainId;
        address[] memory outboxes = new address[](1);
        outboxes[0] = address(outbox);
        inbox.setOutboxes(chainIds, outboxes);
    }

    function initializeOutbox() internal {
        outbox.initialize(address(this), solver, address(portal));

        uint64[] memory chainIds = new uint64[](1);
        chainIds[0] = srcChainId;
        address[] memory inboxes = new address[](1);
        inboxes[0] = address(inbox);
        outbox.setInboxes(chainIds, inboxes);
    }

    function fillHash(bytes32 orderId, bytes memory originData) internal pure returns (bytes32) {
        return keccak256(abi.encode(orderId, originData));
    }

    function test_bench_nativeDeposit_nativeExpense() public {
        SolverNet.Header memory header =
            SolverNet.Header({ owner: user, destChainId: destChainId, fillDeadline: uint32(block.timestamp + 1 days) });

        SolverNet.Deposit memory deposit = SolverNet.Deposit({ token: address(0), amount: defaultAmount });

        SolverNet.Call[] memory calls = new SolverNet.Call[](1);
        calls[0] = SolverNet.Call({
            target: address(nativeVault),
            selector: MockVault.deposit.selector,
            value: defaultAmount,
            params: abi.encode(user, defaultAmount)
        });

        SolverNet.Expense[] memory expenses = new SolverNet.Expense[](0);

        SolverNet.OrderData memory orderData =
            SolverNet.OrderData({ header: header, deposit: deposit, calls: calls, expenses: expenses });

        IERC7683.OnchainCrossChainOrder memory order = IERC7683.OnchainCrossChainOrder({
            fillDeadline: uint32(block.timestamp + 1 days),
            orderDataType: ORDER_DATA_TYPEHASH,
            orderData: abi.encode(orderData)
        });

        assertTrue(inbox.validate(order), "order should be valid");

        vm.deal(user, defaultAmount);
        vm.startPrank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        inbox.open{ value: defaultAmount }(order);
        vm.stopPrank();

        uint256 fillFee = outbox.fillFee(resolvedOrder.fillInstructions[0].originData);
        bytes32 fillhash = fillHash(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData);
        vm.chainId(destChainId);
        vm.deal(solver, defaultAmount + fillFee);
        vm.startPrank(solver);
        inbox.accept(resolvedOrder.orderId);
        outbox.fill{ value: defaultAmount + fillFee }(
            resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData, abi.encode(solver)
        );
        vm.stopPrank();

        vm.chainId(srcChainId);
        portal.mockXCall(
            destChainId,
            address(outbox),
            address(inbox),
            abi.encodeCall(SolverNetInbox.markFilled, (resolvedOrder.orderId, fillhash, solver)),
            200_000
        );

        vm.prank(solver);
        inbox.claim(resolvedOrder.orderId, solver);
    }

    function test_bench_nativeDeposit_erc20Expense() public {
        SolverNet.Header memory header =
            SolverNet.Header({ owner: user, destChainId: destChainId, fillDeadline: uint32(block.timestamp + 1 days) });

        SolverNet.Deposit memory deposit = SolverNet.Deposit({ token: address(0), amount: defaultAmount });

        SolverNet.Call[] memory calls = new SolverNet.Call[](1);
        calls[0] = SolverNet.Call({
            target: address(erc20Vault),
            selector: MockVault.deposit.selector,
            value: 0,
            params: abi.encode(user, defaultAmount)
        });

        SolverNet.Expense[] memory expenses = new SolverNet.Expense[](1);
        expenses[0] = SolverNet.Expense({ spender: address(erc20Vault), token: address(token2), amount: defaultAmount });

        SolverNet.OrderData memory orderData =
            SolverNet.OrderData({ header: header, deposit: deposit, calls: calls, expenses: expenses });

        IERC7683.OnchainCrossChainOrder memory order = IERC7683.OnchainCrossChainOrder({
            fillDeadline: uint32(block.timestamp + 1 days),
            orderDataType: ORDER_DATA_TYPEHASH,
            orderData: abi.encode(orderData)
        });

        assertTrue(inbox.validate(order), "order should be valid");

        vm.deal(user, defaultAmount);
        vm.startPrank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        inbox.open{ value: defaultAmount }(order);
        vm.stopPrank();

        uint256 fillFee = outbox.fillFee(resolvedOrder.fillInstructions[0].originData);
        bytes32 fillhash = fillHash(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData);
        vm.chainId(destChainId);
        vm.deal(solver, fillFee);
        vm.startPrank(solver);
        token2.mint(solver, defaultAmount);
        token2.approve(address(outbox), defaultAmount);
        inbox.accept(resolvedOrder.orderId);
        outbox.fill{ value: fillFee }(
            resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData, abi.encode(solver)
        );
        vm.stopPrank();

        vm.chainId(srcChainId);
        portal.mockXCall(
            destChainId,
            address(outbox),
            address(inbox),
            abi.encodeCall(SolverNetInbox.markFilled, (resolvedOrder.orderId, fillhash, solver)),
            200_000
        );

        vm.prank(solver);
        inbox.claim(resolvedOrder.orderId, solver);
    }

    function test_bench_erc20Deposit_nativeExpense() public {
        SolverNet.Header memory header =
            SolverNet.Header({ owner: user, destChainId: destChainId, fillDeadline: uint32(block.timestamp + 1 days) });

        SolverNet.Deposit memory deposit = SolverNet.Deposit({ token: address(token1), amount: defaultAmount });

        SolverNet.Call[] memory calls = new SolverNet.Call[](1);
        calls[0] = SolverNet.Call({
            target: address(nativeVault),
            selector: MockVault.deposit.selector,
            value: defaultAmount,
            params: abi.encode(user, defaultAmount)
        });

        SolverNet.Expense[] memory expenses = new SolverNet.Expense[](0);

        SolverNet.OrderData memory orderData =
            SolverNet.OrderData({ header: header, deposit: deposit, calls: calls, expenses: expenses });

        IERC7683.OnchainCrossChainOrder memory order = IERC7683.OnchainCrossChainOrder({
            fillDeadline: uint32(block.timestamp + 1 days),
            orderDataType: ORDER_DATA_TYPEHASH,
            orderData: abi.encode(orderData)
        });

        assertTrue(inbox.validate(order), "order should be valid");

        vm.startPrank(user);
        token1.mint(user, defaultAmount);
        token1.approve(address(inbox), defaultAmount);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        inbox.open(order);
        vm.stopPrank();

        uint256 fillFee = outbox.fillFee(resolvedOrder.fillInstructions[0].originData);
        bytes32 fillhash = fillHash(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData);
        vm.chainId(destChainId);
        vm.deal(solver, defaultAmount + fillFee);
        vm.startPrank(solver);
        inbox.accept(resolvedOrder.orderId);
        outbox.fill{ value: defaultAmount + fillFee }(
            resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData, abi.encode(solver)
        );
        vm.stopPrank();

        vm.chainId(srcChainId);
        portal.mockXCall(
            destChainId,
            address(outbox),
            address(inbox),
            abi.encodeCall(SolverNetInbox.markFilled, (resolvedOrder.orderId, fillhash, solver)),
            200_000
        );

        vm.prank(solver);
        inbox.claim(resolvedOrder.orderId, solver);
    }

    function test_bench_erc20Deposit_erc20Expense() public {
        SolverNet.Header memory header =
            SolverNet.Header({ owner: user, destChainId: destChainId, fillDeadline: uint32(block.timestamp + 1 days) });

        SolverNet.Deposit memory deposit = SolverNet.Deposit({ token: address(token1), amount: defaultAmount });

        SolverNet.Call[] memory calls = new SolverNet.Call[](1);
        calls[0] = SolverNet.Call({
            target: address(erc20Vault),
            selector: MockVault.deposit.selector,
            value: 0,
            params: abi.encode(user, defaultAmount)
        });

        SolverNet.Expense[] memory expenses = new SolverNet.Expense[](1);
        expenses[0] = SolverNet.Expense({ spender: address(erc20Vault), token: address(token2), amount: defaultAmount });

        SolverNet.OrderData memory orderData =
            SolverNet.OrderData({ header: header, deposit: deposit, calls: calls, expenses: expenses });

        IERC7683.OnchainCrossChainOrder memory order = IERC7683.OnchainCrossChainOrder({
            fillDeadline: uint32(block.timestamp + 1 days),
            orderDataType: ORDER_DATA_TYPEHASH,
            orderData: abi.encode(orderData)
        });

        assertTrue(inbox.validate(order), "order should be valid");

        vm.startPrank(user);
        token1.mint(user, defaultAmount);
        token1.approve(address(inbox), defaultAmount);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        inbox.open(order);
        vm.stopPrank();

        uint256 fillFee = outbox.fillFee(resolvedOrder.fillInstructions[0].originData);
        bytes32 fillhash = fillHash(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData);
        vm.chainId(destChainId);
        vm.deal(solver, fillFee);
        vm.startPrank(solver);
        token2.mint(solver, defaultAmount);
        token2.approve(address(outbox), defaultAmount);
        inbox.accept(resolvedOrder.orderId);
        outbox.fill{ value: fillFee }(
            resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData, abi.encode(solver)
        );
        vm.stopPrank();

        vm.chainId(srcChainId);
        portal.mockXCall(
            destChainId,
            address(outbox),
            address(inbox),
            abi.encodeCall(SolverNetInbox.markFilled, (resolvedOrder.orderId, fillhash, solver)),
            200_000
        );

        vm.prank(solver);
        inbox.claim(resolvedOrder.orderId, solver);
    }
}
