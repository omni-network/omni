// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "../TestBase.sol";
import { SolverNetExecutor, ISolverNetExecutor } from "src/SolverNetExecutor.sol";
import { Reverter } from "test/utils/Reverter.sol";

contract SolverNet_Outbox_Executor_Test is TestBase {
    using AddrUtils for address;

    SolverNetExecutor internal executor;
    Reverter internal reverter;

    function setUp() public override {
        super.setUp();
        executor = SolverNetExecutor(payable(outbox.executor()));
        reverter = new Reverter();
    }

    function test_all_reverts_not_outbox() public {
        vm.expectRevert(ISolverNetExecutor.NotOutbox.selector);
        executor.approve(address(0x123), address(0x456), 1 ether);

        ISolverNet.Call memory newCall;
        vm.expectRevert(ISolverNetExecutor.NotOutbox.selector);
        executor.execute(newCall);

        vm.expectRevert(ISolverNetExecutor.NotOutbox.selector);
        executor.transfer(address(0x123), address(0x456), 1 ether);

        vm.expectRevert(ISolverNetExecutor.NotOutbox.selector);
        executor.transferNative(address(0x123), 1 ether);
    }

    function test_execute_reverts_call_failed() public {
        ISolverNet.Call memory newCall;
        newCall.target = address(reverter).toBytes32();
        newCall.data = abi.encodeCall(MockVault.deposit, (user, 1 ether));

        vm.expectRevert(ISolverNetExecutor.CallFailed.selector);
        vm.prank(address(outbox));
        executor.execute(newCall);
    }
}
