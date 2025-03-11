// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "../TestBase.sol";

contract SolverNet_Inbox_General_Test is TestBase {
    function test_setOutboxes_reverts() public {
        uint64[] memory chainIds = new uint64[](1);
        chainIds[0] = srcChainId;
        address[] memory inboxes = new address[](0);

        vm.expectRevert(ISolverNetOutbox.InvalidArrayLength.selector);
        outbox.setInboxes(chainIds, inboxes);
    }
}
