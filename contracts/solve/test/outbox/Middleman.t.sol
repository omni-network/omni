// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "../TestBase.sol";
import { MockLST } from "test/utils/MockLST.sol";
import { Refunder } from "test/utils/Refunder.sol";
import { Reverter } from "test/utils/Reverter.sol";

contract SolverNet_Outbox_Middleman_Test is TestBase {
    MockLST lst;
    Refunder refunder;
    Reverter reverter;

    function setUp() public override {
        super.setUp();
        lst = new MockLST();
        refunder = new Refunder();
        reverter = new Reverter();
    }

    function test_executeAndTransfer_reverts() public {
        vm.deal(solver, 1 ether);
        vm.prank(solver);
        vm.expectRevert(SolverNetMiddleman.CallFailed.selector);
        middleman.executeAndTransfer{ value: 1 ether }(address(0), user, address(reverter), "");
    }

    function test_executeAndTransfer_erc20_succeeds() public {
        vm.deal(solver, 1 ether);
        vm.prank(solver);
        middleman.executeAndTransfer{ value: 1 ether }(
            address(lst.token()), user, address(lst), abi.encodeWithSelector(MockLST.deposit.selector)
        );
        assertEq(lst.token().balanceOf(user), 1 ether, "user should have received 1 ether of the LST from middleman");
    }

    function test_executeAndTransfer_native_succeeds() public {
        vm.deal(solver, 1 ether);
        vm.prank(solver);
        middleman.executeAndTransfer{ value: 1 ether }(address(0), user, address(refunder), "");
        assertEq(user.balance, 1 ether, "user should have received 1 ether of native tokens from middleman");
    }
}
