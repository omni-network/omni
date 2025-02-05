// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.26;

import "../TestBase.sol";

contract WithdrawTest is TestBase {
    function test_withdraws_reverts() public {
        bytes32 clawbackerRole = wrapper.CLAWBACKER_ROLE();

        // Cannot withdraw if paused
        vm.prank(user);
        lockbox.deposit(INITIAL_USER_BALANCE);

        vm.prank(pauser);
        lockbox.pause();

        vm.startPrank(user);
        vm.expectRevert(PausableUpgradeable.EnforcedPause.selector);
        lockbox.withdraw(INITIAL_USER_BALANCE);

        vm.expectRevert(PausableUpgradeable.EnforcedPause.selector);
        lockbox.withdrawTo(other, INITIAL_USER_BALANCE);
        vm.stopPrank();

        vm.prank(pauser);
        lockbox.unpause();

        // Cannot withdraw if wrapper does not have clawbacker role
        vm.prank(admin);
        wrapper.revokeRole(clawbackerRole, address(lockbox));

        vm.startPrank(user);
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector, address(lockbox), clawbackerRole
            )
        );
        lockbox.withdraw(INITIAL_USER_BALANCE);

        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector, address(lockbox), clawbackerRole
            )
        );
        lockbox.withdrawTo(other, INITIAL_USER_BALANCE);
        vm.stopPrank();
    }

    function test_withdraw_succeeds() public prank(user) {
        lockbox.deposit(INITIAL_USER_BALANCE);
        lockbox.withdraw(INITIAL_USER_BALANCE);

        assertEq(token.balanceOf(user), INITIAL_USER_BALANCE, "User token balance mismatch");
        assertEq(token.balanceOf(address(lockbox)), 0, "Lockbox token balance mismatch");
        assertEq(wrapper.balanceOf(user), 0, "User wrapper balance mismatch");
    }

    function test_withdrawTo_self_succeeds() public prank(user) {
        lockbox.deposit(INITIAL_USER_BALANCE);
        lockbox.withdrawTo(user, INITIAL_USER_BALANCE);

        assertEq(token.balanceOf(user), INITIAL_USER_BALANCE, "User token balance mismatch");
        assertEq(token.balanceOf(address(lockbox)), 0, "Lockbox token balance mismatch");
        assertEq(wrapper.balanceOf(user), 0, "User wrapper balance mismatch");
    }

    function test_withdrawTo_other_succeeds() public prank(user) {
        lockbox.deposit(INITIAL_USER_BALANCE);
        lockbox.withdrawTo(other, INITIAL_USER_BALANCE);

        assertEq(token.balanceOf(user), 0, "User token balance mismatch");
        assertEq(token.balanceOf(other), INITIAL_USER_BALANCE, "Other token balance mismatch");
        assertEq(token.balanceOf(address(lockbox)), 0, "Lockbox token balance mismatch");
        assertEq(wrapper.balanceOf(user), 0, "User wrapper balance mismatch");
        assertEq(wrapper.balanceOf(other), 0, "Other wrapper balance mismatch");
    }
}
