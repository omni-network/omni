// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.26;

import "../TestBase.sol";

contract GeneralLockboxest is TestBase {
    function test_initialize_reverts() public {
        address impl = address(new Lockbox());
        lockbox = Lockbox(address(new TransparentUpgradeableProxy(impl, admin, "")));

        // `admin` cannot be zero address.
        vm.expectRevert(ILockbox.ZeroAddress.selector);
        lockbox.initialize(address(0), address(0), address(0), address(0));

        // `pauser` cannot be zero address.
        vm.expectRevert(ILockbox.ZeroAddress.selector);
        lockbox.initialize(admin, address(0), address(0), address(0));

        // `token` cannot be zero address.
        vm.expectRevert(ILockbox.ZeroAddress.selector);
        lockbox.initialize(admin, pauser, address(0), address(0));

        // `wrapped` cannot be zero address.
        vm.expectRevert(ILockbox.ZeroAddress.selector);
        lockbox.initialize(admin, pauser, address(token), address(0));
    }
}
