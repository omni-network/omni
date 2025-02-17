// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { SolverNetInbox } from "src/SolverNetInbox.sol";
import { SolverNetOutbox } from "src/SolverNetOutbox.sol";
import { SolverNetMiddleman } from "src/SolverNetMiddleman.sol";
import { ISolverNetInbox } from "src/interfaces/ISolverNetInbox.sol";
import { ISolverNetOutbox } from "src/interfaces/ISolverNetOutbox.sol";

import { IERC20 } from "@openzeppelin/contracts/interfaces/IERC20.sol";
import { MockERC20 } from "test/utils/MockERC20.sol";
import { MockVault } from "test/utils/MockVault.sol";
import { MockMultiTokenVault } from "test/utils/MockMultiTokenVault.sol";
import { MockPortal } from "core/test/utils/MockPortal.sol";

import { IERC7683 } from "src/erc7683/IERC7683.sol";
import { SolverNet } from "src/lib/SolverNet.sol";

import { Test, console2 } from "forge-std/Test.sol";
import { Ownable } from "solady/src/auth/Ownable.sol";
import { FixedPointMathLib } from "solady/src/utils/FixedPointMathLib.sol";
import { AddrUtils } from "src/lib/AddrUtils.sol";

/**
 * @title TestBase
 * @dev Shared test utils / fixtures.
 */
contract TestBase is Test {
    using AddrUtils for address;
    using AddrUtils for bytes32;

    SolverNetInbox inbox;
    SolverNetOutbox outbox;
    SolverNetMiddleman middleman;

    MockERC20 token1;
    MockERC20 token2;

    MockVault nativeVault;
    MockVault erc20Vault;
    MockMultiTokenVault multiTokenVault;

    MockPortal portal;

    uint64 srcChainId = 1;
    uint64 destChainId = 2;

    address user = makeAddr("user");
    address solver = makeAddr("solver");
    address proxyAdmin = makeAddr("proxy-admin-owner");

    uint96 defaultAmount = 100 ether;
    uint32 defaultFillDeadline = uint32(block.timestamp + 1 hours);

    bytes32 internal constant ORDER_DATA_TYPEHASH = keccak256(
        "OrderData(address owner,uint64 destChainId,Deposit deposit,Call[] calls,Expense[] expenses)Deposit(address token,uint96 amount)Call(address target,bytes4 selector,uint256 value,bytes params)Expense(address spender,address token,uint96 amount)"
    );

    modifier prankUser() {
        vm.startPrank(user);
        _;
        vm.stopPrank();
    }

    function setUp() public virtual {
        token1 = new MockERC20("Token 1", "TKN1");
        token2 = new MockERC20("Token 2", "TKN2");

        nativeVault = new MockVault(address(0));
        erc20Vault = new MockVault(address(token2));
        multiTokenVault = new MockMultiTokenVault();
        portal = new MockPortal();

        inbox = deploySolverNetInbox();
        outbox = deploySolverNetOutbox();
        middleman = new SolverNetMiddleman();
        initializeInbox();
        initializeOutbox();

        vm.chainId(srcChainId);
    }

    // Helper functions

    function assertResolvedEq(
        IERC7683.ResolvedCrossChainOrder memory resolved1,
        IERC7683.ResolvedCrossChainOrder memory resolved2
    ) internal pure {
        assertEq(keccak256(abi.encode(resolved1)), keccak256(abi.encode(resolved2)), "resolved orders are not equal");
    }

    function fillHash(bytes32 orderId, bytes memory originData) internal pure returns (bytes32) {
        return keccak256(abi.encode(orderId, originData));
    }

    function fundUser(SolverNet.OrderData memory orderData) internal {
        SolverNet.Deposit memory deposit = orderData.deposit;
        address token = deposit.token;
        uint96 amount = deposit.amount;

        if (amount > 0) {
            if (token == address(0)) {
                vm.deal(user, amount);
            } else {
                vm.prank(user);
                MockERC20(token).approve(address(inbox), type(uint256).max);
                MockERC20(token).mint(user, amount);
            }
        }
    }

    function fundSolver(SolverNet.OrderData memory orderData, uint256 fillFees) internal {
        SolverNet.Call[] memory calls = orderData.calls;
        SolverNet.Expense[] memory expenses = orderData.expenses;

        uint256 nativeValue;
        for (uint256 i; i < calls.length; ++i) {
            SolverNet.Call memory call = calls[i];
            if (call.value > 0) nativeValue += call.value;
        }
        if (nativeValue + fillFees > 0) vm.deal(solver, nativeValue + fillFees);

        for (uint256 i; i < expenses.length; ++i) {
            SolverNet.Expense memory expense = expenses[i];
            address token = expense.token;
            uint96 amount = expense.amount;

            if (amount > 0) {
                vm.prank(solver);
                MockERC20(token).approve(address(outbox), type(uint256).max);
                MockERC20(token).mint(solver, amount);
            }
        }
    }

    function getVaultCall(address vault, uint256 callValue, address depositRecipient, uint256 depositAmount)
        internal
        pure
        returns (SolverNet.Call memory)
    {
        return SolverNet.Call({
            target: vault,
            selector: MockVault.deposit.selector,
            value: callValue,
            params: abi.encode(depositRecipient, depositAmount)
        });
    }

    function getExpense(address spender, address token, uint96 amount)
        internal
        pure
        returns (SolverNet.Expense memory)
    {
        return SolverNet.Expense({ spender: spender, token: token, amount: amount });
    }

    function getOrder(uint256 fillDeadline, SolverNet.OrderData memory orderData)
        internal
        pure
        returns (IERC7683.OnchainCrossChainOrder memory)
    {
        return IERC7683.OnchainCrossChainOrder({
            fillDeadline: uint32(fillDeadline),
            orderDataType: ORDER_DATA_TYPEHASH,
            orderData: abi.encode(orderData)
        });
    }

    function getOrder(
        address owner,
        uint64 chainId,
        uint32 fillDeadline,
        address depositToken,
        uint96 depositAmount,
        SolverNet.Call[] memory calls,
        SolverNet.Expense[] memory expenses
    ) internal pure returns (SolverNet.OrderData memory, IERC7683.OnchainCrossChainOrder memory) {
        SolverNet.OrderData memory orderData = SolverNet.OrderData({
            owner: owner,
            destChainId: chainId,
            deposit: SolverNet.Deposit({ token: depositToken, amount: depositAmount }),
            calls: calls,
            expenses: expenses
        });

        IERC7683.OnchainCrossChainOrder memory order = IERC7683.OnchainCrossChainOrder({
            fillDeadline: fillDeadline,
            orderDataType: ORDER_DATA_TYPEHASH,
            orderData: abi.encode(orderData)
        });

        return (orderData, order);
    }

    function getNativeForNativeVaultOrder(uint256 depositAmount, uint256 expenseAmount)
        internal
        view
        returns (SolverNet.OrderData memory, IERC7683.OnchainCrossChainOrder memory)
    {
        SolverNet.Deposit memory deposit = SolverNet.Deposit({ token: address(0), amount: uint96(depositAmount) });

        SolverNet.Call[] memory calls = new SolverNet.Call[](1);
        calls[0] = SolverNet.Call({
            target: address(nativeVault),
            selector: MockVault.deposit.selector,
            value: depositAmount,
            params: abi.encode(user, expenseAmount)
        });

        SolverNet.Expense[] memory expenses = new SolverNet.Expense[](0);

        SolverNet.OrderData memory orderData = SolverNet.OrderData({
            owner: address(0),
            destChainId: destChainId,
            deposit: deposit,
            calls: calls,
            expenses: expenses
        });

        IERC7683.OnchainCrossChainOrder memory order = IERC7683.OnchainCrossChainOrder({
            fillDeadline: defaultFillDeadline,
            orderDataType: ORDER_DATA_TYPEHASH,
            orderData: abi.encode(orderData)
        });

        return (orderData, order);
    }

    function getErc20ForErc20VaultOrder(uint256 depositAmount, uint256 expenseAmount)
        internal
        view
        returns (SolverNet.OrderData memory, IERC7683.OnchainCrossChainOrder memory)
    {
        SolverNet.Deposit memory deposit = SolverNet.Deposit({ token: address(token1), amount: uint96(depositAmount) });

        SolverNet.Call[] memory calls = new SolverNet.Call[](1);
        calls[0] = SolverNet.Call({
            target: address(erc20Vault),
            selector: MockVault.deposit.selector,
            value: 0,
            params: abi.encode(user, expenseAmount)
        });

        SolverNet.Expense[] memory expenses = new SolverNet.Expense[](1);
        expenses[0] =
            SolverNet.Expense({ spender: address(erc20Vault), token: address(token2), amount: uint96(expenseAmount) });

        SolverNet.OrderData memory orderData = SolverNet.OrderData({
            owner: address(0),
            destChainId: destChainId,
            deposit: deposit,
            calls: calls,
            expenses: expenses
        });

        IERC7683.OnchainCrossChainOrder memory order = IERC7683.OnchainCrossChainOrder({
            fillDeadline: defaultFillDeadline,
            orderDataType: ORDER_DATA_TYPEHASH,
            orderData: abi.encode(orderData)
        });

        return (orderData, order);
    }

    function getArbitraryVaultOrder(
        address depositToken,
        uint96 depositAmount,
        address[] memory expenseTokens,
        uint96[] memory expenseAmounts
    ) internal view returns (SolverNet.OrderData memory, IERC7683.OnchainCrossChainOrder memory) {
        require(expenseTokens.length == expenseAmounts.length, "array length mismatch");

        SolverNet.Deposit memory deposit = SolverNet.Deposit({ token: depositToken, amount: depositAmount });

        SolverNet.Call[] memory calls = new SolverNet.Call[](expenseTokens.length);
        for (uint256 i; i < expenseTokens.length; ++i) {
            calls[i] = SolverNet.Call({
                target: expenseTokens[i] == address(0) ? address(nativeVault) : address(erc20Vault),
                selector: MockVault.deposit.selector,
                value: expenseTokens[i] == address(0) ? expenseAmounts[i] : 0,
                params: abi.encode(user, expenseAmounts[i])
            });
        }

        uint256 expenseLength;
        for (uint256 i; i < expenseTokens.length; ++i) {
            if (expenseTokens[i] != address(0)) ++expenseLength;
        }

        uint256 bias;
        SolverNet.Expense[] memory expenses = new SolverNet.Expense[](expenseLength);
        for (uint256 i; i < expenseTokens.length; ++i) {
            if (expenseTokens[i] == address(0)) {
                ++bias;
                continue;
            }

            expenses[i - bias] =
                SolverNet.Expense({ spender: address(erc20Vault), token: expenseTokens[i], amount: expenseAmounts[i] });
        }

        SolverNet.OrderData memory orderData = SolverNet.OrderData({
            owner: address(0),
            destChainId: destChainId,
            deposit: deposit,
            calls: calls,
            expenses: expenses
        });

        IERC7683.OnchainCrossChainOrder memory order = IERC7683.OnchainCrossChainOrder({
            fillDeadline: defaultFillDeadline,
            orderDataType: ORDER_DATA_TYPEHASH,
            orderData: abi.encode(orderData)
        });

        return (orderData, order);
    }

    function getGasLimit(bytes memory originData) internal pure returns (uint64) {
        SolverNet.FillOriginData memory fillData = abi.decode(originData, (SolverNet.FillOriginData));

        // 2500 gas for the Metadata struct SLOAD.
        uint256 metadataGas = 2500;

        // 2500 gas for Call array length SLOAD + dynamic cost of reading each call.
        uint256 callsGas = 2500;
        for (uint256 i; i < fillData.calls.length; ++i) {
            SolverNet.Call memory call = fillData.calls[i];
            unchecked {
                // 5000 gas for the two slots that hold target, selector, and value.
                // 2500 gas per params slot (1 per function argument) used (minimum of 1 slot).
                callsGas += 5000 + (FixedPointMathLib.divUp(call.params.length + 32, 32) * 2500);
            }
        }

        // 2500 gas for Expense array length SLOAD + cost of reading each expense.
        uint256 expensesGas = 2500;
        unchecked {
            expensesGas += fillData.expenses.length * 5000;
        }

        return uint64(metadataGas + callsGas + expensesGas + 100_000); // 100k base gas limit
    }

    // Setup functions

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
}
