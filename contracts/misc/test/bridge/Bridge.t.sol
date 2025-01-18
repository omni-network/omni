// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "./TestBase.sol";

contract BridgeTest is TestBase {
    using SafeTransferLib for address;

    function test_bridge_direct() public prankUser(user) {
        // Bridge from src to A and B
        address srcToken = address(tokenSrc);
        address destTokenA = bridgeSrc.tokenRoutes(srcToken, destChainIdA);
        address destTokenB = bridgeSrc.tokenRoutes(srcToken, destChainIdB);
        uint256 amount1 = 500_000 ether;

        srcToken.safeApprove(address(bridgeSrc), type(uint256).max);
        mockBridge(bridgeSrc, srcChainId, destChainIdA, srcToken, user, amount1);
        mockBridge(bridgeSrc, srcChainId, destChainIdB, srcToken, user, amount1);

        assertEq(srcToken.balanceOf(user), 0, "src: srcToken.balanceOf(user)");
        assertEq(destTokenA.balanceOf(user), amount1, "src: destTokenA.balanceOf(user)");
        assertEq(destTokenB.balanceOf(user), amount1, "src: destTokenB.balanceOf(user)");
        assertEq(
            srcToken.balanceOf(address(lockboxSrc)),
            INITIAL_USER_BALANCE,
            "src: srcToken.balanceOf(address(lockboxSrc))"
        );

        // Bridge from A to B and src
        uint256 amount2 = amount1 / 2;

        destTokenA.safeApprove(address(bridgeA), type(uint256).max);
        mockBridge(bridgeA, destChainIdA, destChainIdB, destTokenA, user, amount2);
        mockBridge(bridgeA, destChainIdA, srcChainId, destTokenA, user, amount2);

        assertEq(destTokenA.balanceOf(user), 0, "A: destTokenA.balanceOf(user)");
        assertEq(destTokenB.balanceOf(user), amount1 + amount2, "A: destTokenB.balanceOf(user)");
        assertEq(srcToken.balanceOf(user), amount2, "A: srcToken.balanceOf(user)");
        assertEq(
            srcToken.balanceOf(address(lockboxSrc)), amount1 + amount2, "A: srcToken.balanceOf(address(lockboxSrc))"
        );

        // Bridge from B to src and A
        uint256 amount3 = amount2 + (amount2 / 2);

        destTokenB.safeApprove(address(bridgeB), type(uint256).max);
        mockBridge(bridgeB, destChainIdB, srcChainId, destTokenB, user, amount3);
        mockBridge(bridgeB, destChainIdB, destChainIdA, destTokenB, user, amount3);

        assertEq(destTokenB.balanceOf(user), 0, "B: destTokenB.balanceOf(user)");
        assertEq(srcToken.balanceOf(user), amount2 + amount3, "B: srcToken.balanceOf(user)");
        assertEq(destTokenA.balanceOf(user), amount3, "B: destTokenA.balanceOf(user)");
        assertEq(srcToken.balanceOf(address(lockboxSrc)), amount3, "B: srcToken.balanceOf(address(lockboxSrc))");
    }

    function test_bridge_intent() public {
        vm.prank(admin);
        tokenSrc.mint(solver, INITIAL_SOLVER_BALANCE);

        // Bridge from src to A and B for solver
        address srcToken = address(tokenSrc);
        address destTokenA = bridgeSrc.tokenRoutes(srcToken, destChainIdA);
        address destTokenB = bridgeSrc.tokenRoutes(srcToken, destChainIdB);
        uint256 amount = INITIAL_USER_BALANCE;

        vm.startPrank(solver);
        srcToken.safeApprove(address(bridgeSrc), type(uint256).max);
        mockBridge(bridgeSrc, srcChainId, destChainIdA, srcToken, solver, amount);
        mockBridge(bridgeSrc, srcChainId, destChainIdB, srcToken, solver, amount);
        vm.stopPrank();

        assertEq(srcToken.balanceOf(solver), amount, "src (solver): srcToken.balanceOf(solver)");
        assertEq(destTokenA.balanceOf(solver), amount, "src (solver): destTokenA.balanceOf(solver)");
        assertEq(destTokenB.balanceOf(solver), amount, "src (solver): destTokenB.balanceOf(solver)");
        assertEq(
            srcToken.balanceOf(address(lockboxSrc)), amount * 2, "src (solver): srcToken.balanceOf(address(lockboxSrc))"
        );

        // Bridge from src to A and B for user
        uint256 amount1 = amount / 2;
        uint256 solverFee1 = bridgeSrc.bridgeIntentFee(amount1);

        vm.prank(user);
        srcToken.safeApprove(address(bridgeSrc), type(uint256).max);
        mockBridgeIntent(bridgeSrc, srcChainId, destChainIdA, srcToken, user, amount1);
        mockBridgeIntent(bridgeSrc, srcChainId, destChainIdB, srcToken, user, amount1);

        assertEq(srcToken.balanceOf(user), 0, "src (user): srcToken.balanceOf(user)");
        assertEq(srcToken.balanceOf(solver), amount * 2, "src (user): srcToken.balanceOf(solver)");
        assertEq(destTokenA.balanceOf(user), amount1 - solverFee1, "src (user): destTokenA.balanceOf(user)");
        assertEq(destTokenA.balanceOf(solver), amount1 + solverFee1, "src (user): destTokenA.balanceOf(solver)");
        assertEq(destTokenB.balanceOf(user), amount1 - solverFee1, "src (user): destTokenB.balanceOf(user)");
        assertEq(destTokenB.balanceOf(solver), amount1 + solverFee1, "src (user): destTokenB.balanceOf(solver)");
        assertEq(
            srcToken.balanceOf(address(lockboxSrc)), amount * 2, "src (user): srcToken.balanceOf(address(lockboxSrc))"
        );

        // Bridge from A to B and src for user
        uint256 amount2 = (amount1 - solverFee1) / 2;
        uint256 solverFee2 = bridgeA.bridgeIntentFee(amount2);

        vm.prank(user);
        destTokenA.safeApprove(address(bridgeA), type(uint256).max);
        mockBridgeIntent(bridgeA, destChainIdA, destChainIdB, destTokenA, user, amount2);
        mockBridgeIntent(bridgeA, destChainIdA, srcChainId, destTokenA, user, amount2);

        assertEq(destTokenA.balanceOf(user), 0, "A (user): destTokenA.balanceOf(user)");
        assertEq(
            destTokenA.balanceOf(solver), amount1 + solverFee1 + (amount2 * 2), "A (user): destTokenA.balanceOf(solver)"
        );
        assertEq(
            destTokenB.balanceOf(user),
            amount1 - solverFee1 + amount2 - solverFee2,
            "A (user): destTokenB.balanceOf(user)"
        );
        assertEq(
            destTokenB.balanceOf(solver),
            amount1 + solverFee1 - amount2 + solverFee2,
            "A (user): destTokenB.balanceOf(solver)"
        );
        assertEq(srcToken.balanceOf(user), amount2 - solverFee2, "A (user): srcToken.balanceOf(user)");
        assertEq(
            srcToken.balanceOf(solver), (amount * 2) - amount2 + solverFee2, "A (user): srcToken.balanceOf(solver)"
        );
        assertEq(
            srcToken.balanceOf(address(lockboxSrc)), amount * 2, "A (user): srcToken.balanceOf(address(lockboxSrc))"
        );

        // Bridge from B to src and A for user
        uint256 amount3 = ((amount1 - solverFee1) + amount2 - solverFee2) / 2;
        uint256 solverFee3 = bridgeB.bridgeIntentFee(amount3);

        vm.prank(user);
        destTokenB.safeApprove(address(bridgeB), type(uint256).max);
        mockBridgeIntent(bridgeB, destChainIdB, srcChainId, destTokenB, user, amount3);
        mockBridgeIntent(bridgeB, destChainIdB, destChainIdA, destTokenB, user, amount3);

        assertEq(destTokenB.balanceOf(user), 0, "B (user): destTokenB.balanceOf(user)");
        assertEq(
            destTokenB.balanceOf(solver),
            amount1 + solverFee1 - amount2 + solverFee2 + (amount3 * 2),
            "B (user): destTokenB.balanceOf(solver)"
        );
        assertEq(
            srcToken.balanceOf(user), amount2 - solverFee2 + amount3 - solverFee3, "B (user): srcToken.balanceOf(user)"
        );
        assertEq(
            srcToken.balanceOf(solver),
            (amount * 2) - amount2 + solverFee2 - amount3 + solverFee3,
            "B (user): srcToken.balanceOf(solver)"
        );
        assertEq(destTokenA.balanceOf(user), amount3 - solverFee3, "B (user): destTokenA.balanceOf(user)");
        assertEq(
            destTokenA.balanceOf(solver),
            amount1 + solverFee1 + (amount2 * 2) - amount3 + solverFee3,
            "B (user): destTokenA.balanceOf(solver)"
        );
        assertEq(
            srcToken.balanceOf(address(lockboxSrc)), amount * 2, "B (user): srcToken.balanceOf(address(lockboxSrc))"
        );
    }
}
