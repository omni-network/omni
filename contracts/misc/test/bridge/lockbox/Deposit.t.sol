// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.26;

import "../TestBase.sol";

contract DepositTest is TestBase {
    function test_deposits_reverts() public {
        bytes32 minterRole = wrapper.MINTER_ROLE();

        // Cannot deposit if paused
        vm.prank(pauser);
        lockbox.pause();

        vm.startPrank(user);
        vm.expectRevert(PausableUpgradeable.EnforcedPause.selector);
        lockbox.deposit(INITIAL_USER_BALANCE);

        vm.expectRevert(PausableUpgradeable.EnforcedPause.selector);
        lockbox.depositTo(other, INITIAL_USER_BALANCE);
        vm.stopPrank();

        vm.prank(pauser);
        lockbox.unpause();

        // Cannot deposit if wrapper does not have minter role
        vm.prank(admin);
        wrapper.revokeRole(minterRole, address(lockbox));

        vm.startPrank(user);
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector, address(lockbox), minterRole
            )
        );
        lockbox.deposit(INITIAL_USER_BALANCE);

        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector, address(lockbox), minterRole
            )
        );
        lockbox.depositTo(other, INITIAL_USER_BALANCE);
        vm.stopPrank();
    }

    function test_deposit_succeeds() public prank(user) {
        lockbox.deposit(INITIAL_USER_BALANCE);

        assertEq(token.balanceOf(user), 0, "User token balance mismatch");
        assertEq(token.balanceOf(address(lockbox)), INITIAL_USER_BALANCE, "Lockbox token balance mismatch");
        assertEq(wrapper.balanceOf(user), INITIAL_USER_BALANCE, "User wrapper balance mismatch");
    }

    function test_depositTo_self_succeeds() public prank(user) {
        lockbox.depositTo(user, INITIAL_USER_BALANCE);

        assertEq(token.balanceOf(user), 0, "User token balance mismatch");
        assertEq(token.balanceOf(address(lockbox)), INITIAL_USER_BALANCE, "Lockbox token balance mismatch");
        assertEq(wrapper.balanceOf(user), INITIAL_USER_BALANCE, "User wrapper balance mismatch");
    }

    function test_depositTo_other_succeeds() public prank(user) {
        lockbox.depositTo(other, INITIAL_USER_BALANCE);

        assertEq(token.balanceOf(user), 0, "User token balance mismatch");
        assertEq(token.balanceOf(other), 0, "Other token balance mismatch");
        assertEq(token.balanceOf(address(lockbox)), INITIAL_USER_BALANCE, "Lockbox token balance mismatch");
        assertEq(wrapper.balanceOf(user), 0, "User wrapper balance mismatch");
        assertEq(wrapper.balanceOf(other), INITIAL_USER_BALANCE, "Other wrapper balance mismatch");
    }
}
