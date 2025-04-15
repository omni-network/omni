// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "../TestBase.sol";
import { MockLST } from "test/utils/MockLST.sol";
import { Refunder } from "test/utils/Refunder.sol";
import { Reverter } from "test/utils/Reverter.sol";
import { MockERC721 } from "test/utils/MockERC721.sol";

contract SolverNet_Outbox_Middleman_Test is TestBase {
    MockLST lst;
    Refunder refunder;
    Reverter reverter;
    MockERC721 milady;

    function setUp() public override {
        super.setUp();
        lst = new MockLST();
        refunder = new Refunder();
        reverter = new Reverter();
        milady = new MockERC721("Milady Maker", "MILADY", "https://www.miladymaker.net/milady/json/");
    }

    function test_executeAndTransfer_reverts() public {
        vm.deal(solver, 1 ether);
        vm.prank(solver);
        vm.expectRevert(SolverNetMiddleman.CallFailed.selector);
        middleman.executeAndTransfer{ value: 1 ether }(address(0), user, address(reverter), "");
    }

    function test_executeAndTransfer721_reverts() public {
        vm.deal(solver, 1 ether);
        vm.prank(solver);
        vm.expectRevert(SolverNetMiddleman.CallFailed.selector);
        middleman.executeAndTransfer721{ value: 1 ether }(address(0), 0, user, address(reverter), "");
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

    function test_executeAndTransfer721_succeeds() public {
        vm.prank(solver);
        middleman.executeAndTransfer721(
            address(milady), 1, user, address(milady), abi.encodeWithSelector(MockERC721.mint.selector)
        );
        assertEq(milady.ownerOf(1), user, "user should have received the Milady NFT from middleman");
    }
}
