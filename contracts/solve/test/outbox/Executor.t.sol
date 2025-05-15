// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import "../TestBase.sol";
import { SolverNetExecutor, ISolverNetExecutor } from "src/SolverNetExecutor.sol";
import { MockLST } from "test/utils/MockLST.sol";
import { Refunder } from "test/utils/Refunder.sol";
import { Reverter } from "test/utils/Reverter.sol";
import { ApprovalReverterERC20 } from "test/utils/ApprovalReverterERC20.sol";
import { MockERC721 } from "test/utils/MockERC721.sol";

contract SolverNet_Outbox_Executor_Test is TestBase {
    using AddrUtils for address;

    MockLST internal lst;
    Refunder internal refunder;
    Reverter internal reverter;
    ApprovalReverterERC20 internal approvalReverter;
    MockERC721 internal milady;

    function setUp() public override {
        super.setUp();
        lst = new MockLST();
        refunder = new Refunder();
        reverter = new Reverter();
        approvalReverter = new ApprovalReverterERC20();
        approvalReverter.mint(address(executor), 1 ether);
        milady = new MockERC721("Milady Maker", "MILADY", "https://www.miladymaker.net/milady/json/");
    }

    function test_executor_reverts() public {
        vm.expectRevert(ISolverNetExecutor.NotOutbox.selector);
        executor.approve(address(0), address(0), 0);

        vm.expectRevert(ISolverNetExecutor.NotOutbox.selector);
        executor.execute(address(0), 0, "");

        vm.expectRevert(ISolverNetExecutor.NotOutbox.selector);
        executor.transfer(address(0), address(0), 0);

        vm.expectRevert(ISolverNetExecutor.NotOutbox.selector);
        executor.transferNative(address(0), 0);

        vm.expectRevert(ISolverNetExecutor.CallFailed.selector);
        vm.prank(address(outbox));
        executor.execute(address(reverter), 0, "");

        vm.expectRevert(ISolverNetExecutor.NotSelf.selector);
        executor.executeAndTransfer(address(0), address(0), address(0), "");

        vm.expectRevert(ISolverNetExecutor.NotSelf.selector);
        executor.executeAndTransfer721(address(0), 0, address(0), address(0), "");

        vm.expectRevert(ISolverNetExecutor.InvalidToken.selector);
        vm.prank(address(executor));
        executor.executeAndTransfer721(address(0), 0, address(0), address(refunder), "");

        vm.deal(address(executor), 1 ether);
        vm.prank(address(executor));
        vm.expectRevert(ISolverNetExecutor.CallFailed.selector);
        executor.executeAndTransfer{ value: 1 ether }(address(0), user, address(reverter), "");

        vm.prank(address(executor));
        vm.expectRevert(ISolverNetExecutor.CallFailed.selector);
        executor.executeAndTransfer721{ value: 1 ether }(address(0), 0, user, address(reverter), "");
    }

    function test_fallback_reverts() public {
        vm.expectRevert(Receiver.FnSelectorNotRecognized.selector);
        SolverNetInbox(address(executor)).markFilled(bytes32(0), bytes32(0), address(0));
    }

    function test_onERC721Received_succeeds() public {
        milady.mint();
        milady.safeTransferFrom(address(this), address(executor), 1);
        assertEq(milady.ownerOf(1), address(executor), "executor should have received the Milady NFT");
    }

    function test_approve_succeeds() public {
        vm.prank(address(outbox));
        executor.approve(address(token1), user, 1 ether);

        assertEq(token1.allowance(address(executor), user), 1 ether, "allowance should be 1 ether");
    }

    function test_tryRevokeApproval_succeeds_approve_reverts() public {
        vm.startPrank(address(outbox));
        executor.approve(address(approvalReverter), user, 1 ether);
        executor.tryRevokeApproval(address(approvalReverter), user);
        vm.stopPrank();

        assertEq(approvalReverter.allowance(address(executor), user), 1 ether, "allowance should be 1 ether");
    }

    function test_execute_succeeds() public {
        token1.mint(address(executor), 1 ether);

        vm.prank(address(outbox));
        executor.execute(address(token1), 0, abi.encodeCall(IERC20.transfer, (user, 1 ether)));

        assertEq(token1.balanceOf(user), 1 ether, "balance should be 1 ether");
    }

    function test_execute_target_is_executor() public {
        vm.prank(address(outbox));
        executor.execute(
            address(0),
            0,
            abi.encodeCall(
                ISolverNetExecutor.executeAndTransfer721,
                (address(milady), 1, user, address(milady), abi.encodeWithSelector(MockERC721.mint.selector))
            )
        );
        assertEq(milady.ownerOf(1), user, "user should have received the Milady NFT from executor");
    }

    function test_executeAndTransfer_erc20_succeeds() public {
        address lstToken = address(lst.token());
        vm.deal(address(executor), 1 ether);
        vm.prank(address(executor));
        executor.executeAndTransfer{ value: 1 ether }(
            lstToken, user, address(lst), abi.encodeWithSelector(MockLST.deposit.selector)
        );
        assertEq(lst.token().balanceOf(user), 1 ether, "user should have received 1 ether of the LST from executor");
    }

    function test_executeAndTransfer_native_succeeds() public {
        vm.deal(address(executor), 1 ether);
        vm.prank(address(executor));
        executor.executeAndTransfer{ value: 1 ether }(address(0), user, address(refunder), "");
        assertEq(user.balance, 1 ether, "user should have received 1 ether of native tokens from executor");
    }

    function test_executeAndTransfer721_succeeds() public {
        vm.prank(address(executor));
        executor.executeAndTransfer721(
            address(milady), 1, user, address(milady), abi.encodeWithSelector(MockERC721.mint.selector)
        );
        assertEq(milady.ownerOf(1), user, "user should have received the Milady NFT from executor");
    }

    function test_transfer_succeeds() public {
        token1.mint(address(executor), 1 ether);

        vm.prank(address(outbox));
        executor.transfer(address(token1), user, 1 ether);

        assertEq(token1.balanceOf(user), 1 ether, "balance should be 1 ether");
    }

    function test_transferNative_succeeds() public {
        vm.deal(address(executor), 1 ether);

        vm.prank(address(outbox));
        executor.transferNative(user, 1 ether);

        assertEq(user.balance, 1 ether, "balance should be 1 ether");
    }
}
