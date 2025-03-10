// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "../TestBase.sol";

contract SolverNet_Inbox_General_Test is TestBase {
    function test_setOutboxes_reverts() public {
        uint64[] memory chainIds = new uint64[](1);
        chainIds[0] = destChainId;
        address[] memory outboxes = new address[](0);

        vm.expectRevert(ISolverNetInbox.InvalidArrayLength.selector);
        inbox.setOutboxes(chainIds, outboxes);
    }

    function test_pause_reverts() public {
        // Should revert if overwriting an active OPEN pause state
        inbox.pauseOpen(true);
        vm.expectRevert(ISolverNetInbox.IsPaused.selector);
        inbox.pauseOpen(true);

        // Should revert if overwriting an active CLOSE pause state
        inbox.pauseClose(true);
        vm.expectRevert(ISolverNetInbox.IsPaused.selector);
        inbox.pauseClose(true);

        // Should revert if overriding an active ALL_PAUSED pause state
        inbox.pauseAll(true);
        vm.expectRevert(ISolverNetInbox.AllPaused.selector);
        inbox.pauseOpen(true);
        vm.expectRevert(ISolverNetInbox.AllPaused.selector);
        inbox.pauseClose(true);
    }
}
