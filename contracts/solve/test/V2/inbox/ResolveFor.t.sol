// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "../TestBase.sol";

contract SolverNet_Inbox_ResolveFor_Test is TestBase {
    using AddrUtils for address;

    function test_v2_resolveFor_erc20Deposit_erc20Expense_succeeds() public {
        SolverNet.Call[] memory calls = new SolverNet.Call[](1);
        calls[0] = getVaultCall(address(erc20Vault), 0, user, defaultAmount);
        SolverNet.TokenExpense[] memory expenses = new SolverNet.TokenExpense[](1);
        expenses[0] = getExpense(address(erc20Vault), address(token2), defaultAmount);

        (SolverNet.OmniOrderData memory orderData, IERC7683.GaslessCrossChainOrder memory order) = getGaslessOrder(
            user,
            user,
            1,
            destChainId,
            defaultOpenDeadline,
            defaultFillDeadline,
            address(token1),
            defaultAmount,
            calls,
            expenses
        );

        bytes32 id = inbox.getOrderId(user, 1, true);
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolved = inbox.resolveFor(order, "");

        assertEq(resolved.user, user, "user");
        assertEq(resolved.originChainId, srcChainId, "originChainId");
        assertEq(resolved.openDeadline, order.openDeadline, "openDeadline");
        assertEq(resolved.fillDeadline, order.fillDeadline, "fillDeadline");
        assertEq(resolved.orderId, id, "orderId");
        assertEq(resolved.maxSpent.length, 1, "maxSpent.length");
        assertEq(resolved.maxSpent[0].token, address(token2).toBytes32(), "maxSpent[0].token");
        assertEq(resolved.maxSpent[0].amount, defaultAmount, "maxSpent[0].amount");
        assertEq(resolved.maxSpent[0].recipient, bytes32(0), "maxSpent[0].recipient");
        assertEq(resolved.maxSpent[0].chainId, destChainId, "maxSpent[0].chainId");
        assertEq(resolved.minReceived.length, 1, "minReceived.length");
        assertEq(resolved.minReceived[0].token, address(token1).toBytes32(), "minReceived[0].token");
        assertEq(resolved.minReceived[0].amount, defaultAmount, "minReceived[0].amount");
        assertEq(resolved.minReceived[0].recipient, bytes32(0), "minReceived[0].recipient");
        assertEq(resolved.minReceived[0].chainId, srcChainId, "minReceived[0].chainId");
        assertEq(resolved.fillInstructions.length, 1, "fillInstructions.length");
        assertEq(resolved.fillInstructions[0].destinationChainId, destChainId, "fillInstructions[0].destinationChainId");
        assertEq(
            resolved.fillInstructions[0].destinationSettler,
            address(outbox).toBytes32(),
            "fillInstructions[0].destinationSettler"
        );
        assertEq(
            resolved.fillInstructions[0].originData,
            abi.encode(
                SolverNet.FillOriginData({
                    srcChainId: srcChainId,
                    destChainId: orderData.destChainId,
                    fillDeadline: order.fillDeadline,
                    calls: orderData.calls,
                    expenses: orderData.expenses
                })
            ),
            "fillInstructions[0].originData"
        );
    }

    function test_v2_resolveFor_erc20Deposit_mixedExpenses_multicall_succeeds() public {
        SolverNet.Call[] memory calls = new SolverNet.Call[](2);
        calls[0] = getVaultCall(address(nativeVault), defaultAmount, user, defaultAmount);
        calls[1] = getVaultCall(address(erc20Vault), 0, user, defaultAmount);
        SolverNet.TokenExpense[] memory expenses = new SolverNet.TokenExpense[](1);
        expenses[0] = getExpense(address(erc20Vault), address(token2), defaultAmount);

        (SolverNet.OmniOrderData memory orderData, IERC7683.GaslessCrossChainOrder memory order) = getGaslessOrder(
            user,
            user,
            1,
            destChainId,
            defaultOpenDeadline,
            defaultFillDeadline,
            address(token1),
            defaultAmount * 2,
            calls,
            expenses
        );

        bytes32 id = inbox.getOrderId(user, 1, true);
        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolved = inbox.resolveFor(order, "");

        assertEq(resolved.user, user, "user");
        assertEq(resolved.originChainId, srcChainId, "originChainId");
        assertEq(resolved.openDeadline, order.openDeadline, "openDeadline");
        assertEq(resolved.fillDeadline, order.fillDeadline, "fillDeadline");
        assertEq(resolved.orderId, id, "orderId");
        assertEq(resolved.maxSpent.length, 2, "maxSpent.length");
        assertEq(resolved.maxSpent[0].token, address(token2).toBytes32(), "maxSpent[0].token");
        assertEq(resolved.maxSpent[0].amount, defaultAmount, "maxSpent[0].amount");
        assertEq(resolved.maxSpent[0].recipient, bytes32(0), "maxSpent[0].recipient");
        assertEq(resolved.maxSpent[0].chainId, destChainId, "maxSpent[0].chainId");
        assertEq(resolved.maxSpent[1].token, bytes32(0), "maxSpent[1].token");
        assertEq(resolved.maxSpent[1].amount, defaultAmount, "maxSpent[1].amount");
        assertEq(resolved.maxSpent[1].recipient, bytes32(0), "maxSpent[1].recipient");
        assertEq(resolved.maxSpent[1].chainId, destChainId, "maxSpent[1].chainId");
        assertEq(resolved.minReceived.length, 1, "minReceived.length");
        assertEq(resolved.minReceived[0].token, address(token1).toBytes32(), "minReceived[0].token");
        assertEq(resolved.minReceived[0].amount, defaultAmount * 2, "minReceived[0].amount");
        assertEq(resolved.minReceived[0].recipient, bytes32(0), "minReceived[0].recipient");
        assertEq(resolved.minReceived[0].chainId, srcChainId, "minReceived[0].chainId");
        assertEq(resolved.fillInstructions.length, 1, "fillInstructions.length");
        assertEq(resolved.fillInstructions[0].destinationChainId, destChainId, "fillInstructions[0].destinationChainId");
        assertEq(
            resolved.fillInstructions[0].destinationSettler,
            address(outbox).toBytes32(),
            "fillInstructions[0].destinationSettler"
        );
        assertEq(
            resolved.fillInstructions[0].originData,
            abi.encode(
                SolverNet.FillOriginData({
                    srcChainId: srcChainId,
                    destChainId: orderData.destChainId,
                    fillDeadline: order.fillDeadline,
                    calls: orderData.calls,
                    expenses: orderData.expenses
                })
            ),
            "fillInstructions[0].originData"
        );
    }

    function test_v2_resolveFor_hyperlane() public {
        address impl = address(new SolverNetInboxV2(address(0), address(mailboxes[uint32(srcChainId)])));
        inbox = SolverNetInboxV2(address(new TransparentUpgradeableProxy(impl, proxyAdmin, bytes(""))));
        inbox.initialize(address(this), solver);
        setRoutes(ISolverNetOutbox.Provider.Hyperlane);

        uint256 snapshot = vm.snapshotState();
        test_v2_resolveFor_erc20Deposit_erc20Expense_succeeds();
        vm.revertToState(snapshot);

        test_v2_resolveFor_erc20Deposit_mixedExpenses_multicall_succeeds();
    }
}
