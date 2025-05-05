// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "../TestBase.sol";

contract SolverNet_Inbox_General_Test is TestBase {
    function test_setOutboxes_reverts() public {
        uint64[] memory chainIds = new uint64[](1);
        chainIds[0] = srcChainId;
        ISolverNetOutbox.InboxConfig[] memory configs = new ISolverNetOutbox.InboxConfig[](0);

        vm.expectRevert(ISolverNetOutbox.InvalidArrayLength.selector);
        outbox.setInboxes(chainIds, configs);
    }

    function test_fillFee_reverts() public {
        setRoutes(ISolverNetOutbox.Provider.None);

        (, IERC7683.OnchainCrossChainOrder memory order) = getNativeForNativeVaultOrder(defaultAmount, defaultAmount);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = inbox.resolve(order);

        vm.chainId(destChainId);
        vm.expectRevert(ISolverNetOutbox.InvalidConfig.selector);
        outbox.fillFee(resolvedOrder.fillInstructions[0].originData);
    }
}
