// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.26;

import "./TestBase.sol";

contract BridgeTest is TestBase {
    using SafeTransferLib for address;

    function test_bridge_success() public prankUser(user) {
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
}
