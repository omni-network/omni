// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "../TestBase.sol";

contract SolverNet_Outbox_General_Test is TestBase {
    function test_v2_setOutboxes_reverts() public {
        uint64[] memory chainIds = new uint64[](1);
        chainIds[0] = srcChainId;
        ISolverNetOutbox.InboxConfig[] memory configs = new ISolverNetOutbox.InboxConfig[](0);

        vm.expectRevert(ISolverNetOutbox.InvalidArrayLength.selector);
        outbox.setInboxes(chainIds, configs);
    }

    function test_v2_fillFee_reverts() public {
        setRoutes(ISolverNetOutbox.Provider.None);

        (, IERC7683.OnchainCrossChainOrder memory order) = getNativeForNativeVaultOrder(defaultAmount, defaultAmount);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);

        vm.chainId(destChainId);
        vm.expectRevert(ISolverNetOutbox.InvalidConfig.selector);
        outbox.fillFee(resolvedOrder.fillInstructions[0].originData);
    }

    function test_v2_didFill_oldFillHash_succeeds() public {
        (, IERC7683.OnchainCrossChainOrder memory order) = getNativeForNativeVaultOrder(defaultAmount, defaultAmount);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);
        bytes32 oldFillHash = keccak256(abi.encode(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData));
        bytes32 newFillHash = fillHash(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData);
        bytes32 value = bytes32(type(uint256).max);
        assertTrue(oldFillHash != newFillHash, "fill hashes should not match");

        // Calculate the specific storage slot for _filled[oldFillHash]
        // The base slot of the _filled mapping is 4, according to `forge inspect`.
        bytes32 slot = keccak256(abi.encode(oldFillHash, uint256(4)));
        vm.store(address(outbox), slot, value);
        assertTrue(
            outbox.didFill(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData),
            "should return true with old fill hash"
        );

        // Wipe and set value for new fill hash
        vm.store(address(outbox), slot, bytes32(0));
        slot = keccak256(abi.encode(newFillHash, uint256(4)));
        vm.store(address(outbox), slot, value);
        assertTrue(
            outbox.didFill(resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData),
            "should return true with new fill hash"
        );
    }
}
