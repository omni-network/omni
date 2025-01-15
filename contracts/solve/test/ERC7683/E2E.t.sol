// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "./TestBase.sol";

contract SolverNet_E2E_Test is TestBase {
    function test_e2e_complete_order() public {
        // Prep: Set chainId to srcChainId
        vm.chainId(srcChainId);

        // 0. Generate order, validate it, resolve it, and prepare deposit tokens
        IERC7683.OnchainCrossChainOrder memory order = randOrder();
        assertTrue(inbox.validate(order));
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        mintAndApprove(resolvedOrder.minReceived, resolvedOrder.maxSpent);

        assertNullOrder(resolvedOrder.orderId);

        // 1. Open order on srcChain
        vm.prank(user);
        inbox.open(order);

        assertOpenedOrder(resolvedOrder.orderId);

        // 2. Accept order on srcChain
        vm.prank(solver);
        inbox.accept(resolvedOrder.orderId);

        assertAcceptedOrder(resolvedOrder.orderId);

        // Prep: Set chainId to destChainId and give solver some funds
        vm.chainId(destChainId);
        uint256 fillFee = outbox.fillFee(srcChainId);
        vm.deal(address(solver), fillFee);

        // 3. Fill order on destChain
        bytes32 fillHash = fillHash(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData);
        vm.expectEmit(true, true, true, true);
        emit ISolverNetOutbox.Filled(resolvedOrder.orderId, fillHash, solver);
        // Solver token mint and approval is taken care of in step 0 `mintAndApprove` helper call
        vm.prank(solver);
        outbox.fill{ value: fillFee }(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData, bytes(""));

        assertVaultDeposit(resolvedOrder.orderId);
        assertTrue(outbox.didFill(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData));

        // Prep: Set chainId back to srcChainId
        vm.chainId(srcChainId);

        // 4. Mock markFulfilled call from destChain to srcChain
        portal.mockXCall(
            destChainId,
            address(outbox),
            address(inbox),
            abi.encodeCall(
                ISolverNetInbox.markFilled, (resolvedOrder.orderId, fillHash, uint40(block.timestamp), bytes32(0))
            ),
            100_000
        );

        assertFulfilledOrder(resolvedOrder.orderId);

        // 5. Claim order deposits on srcChain as solver
        vm.prank(solver);
        inbox.claim(resolvedOrder.orderId, solver);

        assertClaimedOrder(resolvedOrder.orderId);
    }
}
