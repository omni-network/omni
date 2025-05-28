// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "../TestBase.sol";
import { ISolverNetOutbox } from "src/interfaces/ISolverNetOutbox.sol";
import { SolverNet } from "src/lib/SolverNet.sol";

contract SolverNet_Outbox_RetryMarkFilled_Test is TestBase {
    function test_v2_retryMarkFilled_reverts(uint8 provider) public {
        provider = uint8(bound(provider, uint8(1), uint8(2)));
        setRoutes(ISolverNetOutbox.Provider(provider));

        vm.chainId(srcChainId);
        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getNativeForNativeVaultOrder(defaultAmount, defaultAmount);

        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        bytes32 orderId = resolvedOrder.orderId;
        bytes memory originData = resolvedOrder.fillInstructions[0].originData;
        bytes memory fillerDataSolver = abi.encode(solver);

        vm.chainId(destChainId);

        // fillerData must be empty or 32 bytes
        bytes memory badFillerData = hex"0420";
        vm.expectRevert(ISolverNetOutbox.BadFillerData.selector);
        outbox.retryMarkFilled(orderId, originData, badFillerData);

        // order must already be filled
        vm.expectRevert(ISolverNetOutbox.NotFilled.selector);
        outbox.retryMarkFilled(orderId, originData, fillerDataSolver);

        // fill order for future test cases
        uint256 fillFee = outbox.fillFee(originData);
        fundSolver(orderData, fillFee); // Fund solver for the actual fill
        vm.prank(solver);
        outbox.fill{ value: defaultAmount + fillFee }(orderId, originData, fillerDataSolver);

        // settlement hash must match params used by filler
        bytes memory fillerDataUser = abi.encode(user);
        vm.deal(address(this), fillFee); // Deal current contract for retry
        vm.expectRevert(ISolverNetOutbox.InvalidSettlement.selector);
        outbox.retryMarkFilled{ value: fillFee }(orderId, originData, fillerDataUser);

        // `msg.value` must be at least `fillFee`
        vm.deal(address(outbox), fillFee);
        vm.expectRevert(ISolverNetOutbox.InsufficientFee.selector);
        outbox.retryMarkFilled{ value: fillFee - 1 }(orderId, originData, fillerDataSolver);
    }

    function test_v2_retryMarkFilled_succeeds(uint8 provider) public {
        provider = uint8(bound(provider, uint8(1), uint8(2)));
        setRoutes(ISolverNetOutbox.Provider(provider));

        vm.chainId(srcChainId);
        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getNativeForNativeVaultOrder(defaultAmount, defaultAmount);

        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        bytes32 orderId = resolvedOrder.orderId;
        bytes memory originData = resolvedOrder.fillInstructions[0].originData;
        bytes memory fillerDataSolver = abi.encode(solver);

        vm.chainId(destChainId);
        uint256 fillFee = outbox.fillFee(originData);
        fundSolver(orderData, defaultAmount + fillFee);

        vm.prank(solver);
        outbox.fill{ value: defaultAmount + fillFee }(orderId, originData, fillerDataSolver);

        vm.deal(solver, fillFee);
        vm.expectEmit(true, true, true, true, address(outbox));
        emit ISolverNetOutbox.MarkFilledRetry(orderId, fillHash(orderId, originData), solver);
        vm.prank(solver);
        outbox.retryMarkFilled{ value: fillFee }(orderId, originData, fillerDataSolver);

        assertEq(address(outbox).balance, 0, "Outbox native balance should be zero after successful retry (exact fee)");
    }

    function test_v2_retryMarkFilled_succeeds_feeRefund(uint8 provider) public {
        provider = uint8(bound(provider, uint8(1), uint8(2)));
        setRoutes(ISolverNetOutbox.Provider(provider));

        vm.chainId(srcChainId);
        (SolverNet.OrderData memory orderData, IERC7683.OnchainCrossChainOrder memory order) =
            getNativeForNativeVaultOrder(defaultAmount, defaultAmount);

        vm.prank(user);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        bytes32 orderId = resolvedOrder.orderId;
        bytes memory originData = resolvedOrder.fillInstructions[0].originData;
        bytes memory fillerDataSolver = abi.encode(solver);

        vm.chainId(destChainId);
        uint256 fillFee = outbox.fillFee(originData);
        fundSolver(orderData, defaultAmount + fillFee);

        vm.prank(solver);
        outbox.fill{ value: defaultAmount + fillFee }(orderId, originData, fillerDataSolver);

        vm.deal(solver, fillFee + 1 ether);
        vm.expectEmit(true, true, true, true, address(outbox));
        emit ISolverNetOutbox.MarkFilledRetry(orderId, fillHash(orderId, originData), solver);
        vm.prank(solver);
        outbox.retryMarkFilled{ value: fillFee + 1 ether }(orderId, originData, fillerDataSolver);

        assertEq(address(outbox).balance, 0, "Outbox native balance should be zero after successful retry");
        assertEq(solver.balance, 1 ether, "Solver balance should be 1 ether after successful retry");
    }
}
