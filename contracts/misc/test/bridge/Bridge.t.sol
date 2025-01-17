// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "./TestBase.sol";

contract BridgeTest is TestBase {
    using SafeTransferLib for address;

    function test_bridge() public prankUser {
        // Bridge from src to A and B
        address srcToken = address(tokenSrc);
        address destTokenA = bridgeSrc.tokenRoutes(srcToken, destChainIdA);
        address destTokenB = bridgeSrc.tokenRoutes(srcToken, destChainIdB);

        srcToken.safeApprove(address(bridgeSrc), type(uint256).max);
        mockBridge(bridgeSrc, srcChainId, destChainIdA, srcToken, user, 500_000 ether);
        mockBridge(bridgeSrc, srcChainId, destChainIdB, srcToken, user, 500_000 ether);

        assertEq(srcToken.balanceOf(user), 0);
        assertEq(destTokenA.balanceOf(user), 500_000 ether);
        assertEq(destTokenB.balanceOf(user), 500_000 ether);

        // Bridge from A to B and src
        srcToken = address(tokenA);
        destTokenA = bridgeA.tokenRoutes(srcToken, destChainIdB);
        destTokenB = bridgeA.tokenRoutes(srcToken, srcChainId);

        srcToken.safeApprove(address(bridgeA), type(uint256).max);
        mockBridge(bridgeA, destChainIdA, destChainIdB, srcToken, user, 250_000 ether);
        mockBridge(bridgeA, destChainIdA, srcChainId, srcToken, user, 250_000 ether);

        assertEq(srcToken.balanceOf(user), 0);
        assertEq(destTokenA.balanceOf(user), 750_000 ether);
        assertEq(destTokenB.balanceOf(user), 250_000 ether);

        // Bridge from B to src and A
        srcToken = address(tokenB);
        destTokenA = bridgeB.tokenRoutes(srcToken, srcChainId);
        destTokenB = bridgeB.tokenRoutes(srcToken, destChainIdA);

        srcToken.safeApprove(address(bridgeB), type(uint256).max);
        mockBridge(bridgeB, destChainIdB, srcChainId, srcToken, user, 375_000 ether);
        mockBridge(bridgeB, destChainIdB, destChainIdA, srcToken, user, 375_000 ether);

        assertEq(srcToken.balanceOf(user), 0);
        assertEq(destTokenA.balanceOf(user), 625_000 ether);
        assertEq(destTokenB.balanceOf(user), 375_000 ether);
    }
}
