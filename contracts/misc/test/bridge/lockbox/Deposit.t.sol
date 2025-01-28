// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.26;

import "../TestBase.sol";

contract DepositTest is TestBase {
    function test_deposits_reverts() public {
        bytes32 minterRole = srcWrapper.MINTER_ROLE();

        // Cannot deposit if paused
        vm.prank(pauser);
        srcLockbox.pause();

        vm.startPrank(user);
        vm.expectRevert(PausableUpgradeable.EnforcedPause.selector);
        srcLockbox.deposit(INITIAL_USER_BALANCE);

        vm.expectRevert(PausableUpgradeable.EnforcedPause.selector);
        srcLockbox.depositTo(other, INITIAL_USER_BALANCE);
        vm.stopPrank();

        vm.prank(pauser);
        srcLockbox.unpause();

        // Cannot deposit if wrapper does not have minter role
        vm.prank(admin);
        srcWrapper.revokeRole(minterRole, address(srcLockbox));

        vm.startPrank(user);
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector, address(srcLockbox), minterRole
            )
        );
        srcLockbox.deposit(INITIAL_USER_BALANCE);

        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector, address(srcLockbox), minterRole
            )
        );
        srcLockbox.depositTo(other, INITIAL_USER_BALANCE);
        vm.stopPrank();
    }

    function test_deposit_succeeds() public prank(user) {
        srcLockbox.deposit(INITIAL_USER_BALANCE);

        assertEq(originalToken.balanceOf(user), 0, "Original token balance mismatch");
        assertEq(originalToken.balanceOf(address(srcLockbox)), INITIAL_USER_BALANCE, "Source lockbox balance mismatch");
        assertEq(srcWrapper.balanceOf(user), INITIAL_USER_BALANCE, "Source wrapped token balance mismatch");
    }

    function test_depositTo_self_succeeds() public prank(user) {
        srcLockbox.depositTo(user, INITIAL_USER_BALANCE);

        assertEq(originalToken.balanceOf(user), 0, "Original token balance mismatch");
        assertEq(originalToken.balanceOf(address(srcLockbox)), INITIAL_USER_BALANCE, "Source lockbox balance mismatch");
        assertEq(srcWrapper.balanceOf(user), INITIAL_USER_BALANCE, "Source wrapped token balance mismatch");
    }

    function test_depositTo_other_succeeds() public prank(user) {
        srcLockbox.depositTo(other, INITIAL_USER_BALANCE);

        assertEq(originalToken.balanceOf(user), 0, "Original token sender balance mismatch");
        assertEq(originalToken.balanceOf(other), 0, "Original token recipient balance mismatch");
        assertEq(originalToken.balanceOf(address(srcLockbox)), INITIAL_USER_BALANCE, "Source lockbox balance mismatch");
        assertEq(srcWrapper.balanceOf(user), 0, "Source wrapped token sender balance mismatch");
        assertEq(srcWrapper.balanceOf(other), INITIAL_USER_BALANCE, "Source wrapped token recipient balance mismatch");
    }
}
