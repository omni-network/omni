// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.26;

import "../TestBase.sol";

contract WithdrawTest is TestBase {
    function test_withdraws_revert_paused() public {
        vm.prank(user);
        srcLockbox.deposit(INITIAL_USER_BALANCE);

        vm.prank(pauser);
        srcLockbox.pause();

        vm.startPrank(user);
        vm.expectRevert(PausableUpgradeable.EnforcedPause.selector);
        srcLockbox.withdraw(INITIAL_USER_BALANCE);

        vm.expectRevert(PausableUpgradeable.EnforcedPause.selector);
        srcLockbox.withdrawTo(other, INITIAL_USER_BALANCE);
        vm.stopPrank();

        // Works again after unpausing
        vm.prank(pauser);
        srcLockbox.unpause();

        vm.startPrank(user);
        srcLockbox.withdraw(1);
        srcLockbox.withdrawTo(other, 1);
        vm.stopPrank();
    }

    function test_withdraws_revert_without_wrapper_clawbacker_role() public {
        bytes32 clawbackerRole = srcWrapper.CLAWBACKER_ROLE();

        vm.prank(user);
        srcLockbox.deposit(INITIAL_USER_BALANCE);

        vm.prank(admin);
        srcWrapper.revokeRole(clawbackerRole, address(srcLockbox));

        vm.startPrank(user);
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector, address(srcLockbox), clawbackerRole
            )
        );
        srcLockbox.withdraw(INITIAL_USER_BALANCE);

        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector, address(srcLockbox), clawbackerRole
            )
        );
        srcLockbox.withdrawTo(other, INITIAL_USER_BALANCE);
        vm.stopPrank();
    }

    function test_withdraw_succeeds() public prank(user) {
        srcLockbox.deposit(INITIAL_USER_BALANCE);
        srcLockbox.withdraw(INITIAL_USER_BALANCE);

        assertEq(originalToken.balanceOf(user), INITIAL_USER_BALANCE, "Original token balance mismatch");
        assertEq(originalToken.balanceOf(address(srcLockbox)), 0, "Source lockbox balance mismatch");
        assertEq(srcWrapper.balanceOf(user), 0, "Source wrapped token balance mismatch");
    }

    function test_withdrawTo_self_succeeds() public prank(user) {
        srcLockbox.deposit(INITIAL_USER_BALANCE);
        srcLockbox.withdrawTo(user, INITIAL_USER_BALANCE);

        assertEq(originalToken.balanceOf(user), INITIAL_USER_BALANCE, "Original token balance mismatch");
        assertEq(originalToken.balanceOf(address(srcLockbox)), 0, "Source lockbox balance mismatch");
        assertEq(srcWrapper.balanceOf(user), 0, "Source wrapped token balance mismatch");
    }

    function test_withdrawTo_other_succeeds() public prank(user) {
        srcLockbox.deposit(INITIAL_USER_BALANCE);
        srcLockbox.withdrawTo(other, INITIAL_USER_BALANCE);

        assertEq(originalToken.balanceOf(user), 0, "Original token sender balance mismatch");
        assertEq(originalToken.balanceOf(other), INITIAL_USER_BALANCE, "Original token recipient balance mismatch");
        assertEq(originalToken.balanceOf(address(srcLockbox)), 0, "Source lockbox balance mismatch");
        assertEq(srcWrapper.balanceOf(user), 0, "Source wrapped token sender balance mismatch");
        assertEq(srcWrapper.balanceOf(other), 0, "Source wrapped token recipient balance mismatch");
    }
}
