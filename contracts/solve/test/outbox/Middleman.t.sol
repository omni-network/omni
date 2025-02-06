// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "../TestBase.sol";
import { MockLST } from "test/utils/MockLST.sol";

contract SolverNet_Outbox_Middleman_Test is TestBase {
    MockLST lst;

    function setUp() public override {
        super.setUp();
        lst = new MockLST();
    }

    function test_executeAndTransfer_succeeds() public {
        vm.deal(solver, 1 ether);
        vm.prank(solver);
        middleman.executeAndTransfer{ value: 1 ether }(
            address(lst.token()), user, address(lst), abi.encodeWithSelector(MockLST.deposit.selector)
        );
        assertEq(lst.token().balanceOf(user), 1 ether);
    }
}
