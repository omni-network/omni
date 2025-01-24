// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.26;

import "../TestBase.sol";

contract SendTokenTest is TestBase {
    function test_sendToken_reverts() public prank(user) {
        uint256 fee = srcBridge.bridgeFee(DEST_CHAIN_ID);
        bytes32 burnerRole = originalToken.BURNER_ROLE();

        // `destChainId` must have a configured route.
        vm.expectRevert(abi.encodeWithSelector(IBridge.InvalidRoute.selector, 0));
        srcBridge.sendToken{ value: fee }(0, address(0), 0, true);
        vm.expectRevert(abi.encodeWithSelector(IBridge.InvalidRoute.selector, 0));
        srcBridge.sendToken{ value: fee }(0, address(0), 0, false);

        // `to` cannot be zero address.
        vm.expectRevert(abi.encodeWithSelector(IBridge.ZeroAddress.selector));
        srcBridge.sendToken{ value: fee }(DEST_CHAIN_ID, address(0), 0, true);
        vm.expectRevert(abi.encodeWithSelector(IBridge.ZeroAddress.selector));
        srcBridge.sendToken{ value: fee }(DEST_CHAIN_ID, address(0), 0, false);

        // `value` cannot be zero.
        vm.expectRevert(abi.encodeWithSelector(IBridge.ZeroAmount.selector));
        srcBridge.sendToken{ value: fee }(DEST_CHAIN_ID, user, 0, true);
        vm.expectRevert(abi.encodeWithSelector(IBridge.ZeroAmount.selector));
        srcBridge.sendToken{ value: fee }(DEST_CHAIN_ID, user, 0, false);

        // `wrap` cannot be true if `lockbox` is not set.
        vm.expectRevert(abi.encodeWithSelector(IBridge.CannotWrap.selector));
        destBridge.sendToken{ value: fee }(SRC_CHAIN_ID, user, 1, true);

        // `amount` cannot exceed the user's balance.
        vm.expectRevert(abi.encodeWithSelector(SafeTransferLib.TransferFromFailed.selector));
        srcBridge.sendToken{ value: fee }(DEST_CHAIN_ID, user, INITIAL_USER_BALANCE + 1, true);
        vm.expectRevert(abi.encodeWithSelector(SafeTransferLib.TransferFromFailed.selector));
        srcBridge.sendToken{ value: fee }(DEST_CHAIN_ID, user, INITIAL_USER_BALANCE + 1, false);

        // `bridgeFee` must be paid.
        vm.expectRevert("XApp: insufficient funds");
        srcBridge.sendToken{ value: fee - 1 }(DEST_CHAIN_ID, user, 1, true);

        srcLockbox.deposit(1);
        vm.expectRevert("XApp: insufficient funds");
        srcBridge.sendToken{ value: fee - 1 }(DEST_CHAIN_ID, user, 1, false);

        // Proper usage succeeds.
        srcBridge.sendToken{ value: fee }(DEST_CHAIN_ID, user, 1, true);
    }

    function test_sendToken_reverts_paused() public {
        uint256 fee = srcBridge.bridgeFee(DEST_CHAIN_ID);

        mockBridge(srcBridge, SRC_CHAIN_ID, DEST_CHAIN_ID, true, user, user, 1);

        vm.startPrank(pauser);
        srcBridge.pause();
        destBridge.pause();
        vm.stopPrank();

        vm.startPrank(user);
        vm.expectRevert(PausableUpgradeable.EnforcedPause.selector);
        srcBridge.sendToken{ value: fee }(DEST_CHAIN_ID, user, 1, true);
        vm.expectRevert(PausableUpgradeable.EnforcedPause.selector);
        srcBridge.sendToken{ value: fee }(DEST_CHAIN_ID, user, 1, false);

        fee = destBridge.bridgeFee(SRC_CHAIN_ID);

        vm.expectRevert(PausableUpgradeable.EnforcedPause.selector);
        destBridge.sendToken{ value: fee }(SRC_CHAIN_ID, user, 1, true);
        vm.expectRevert(PausableUpgradeable.EnforcedPause.selector);
        destBridge.sendToken{ value: fee }(SRC_CHAIN_ID, user, 1, false);
        vm.stopPrank();

        // Works again after unpausing
        vm.startPrank(pauser);
        srcBridge.unpause();
        destBridge.unpause();
        vm.stopPrank();

        vm.startPrank(user);
        srcBridge.sendToken{ value: fee }(DEST_CHAIN_ID, user, 1, true);
        destBridge.sendToken{ value: fee }(SRC_CHAIN_ID, user, 1, false);
        vm.stopPrank();
    }

    function test_sendToken_reverts_without_wrapper_burner_role() public {
        uint256 fee = srcBridge.bridgeFee(DEST_CHAIN_ID);
        bytes32 burnerRole = srcWrapper.BURNER_ROLE();

        mockBridge(srcBridge, SRC_CHAIN_ID, DEST_CHAIN_ID, true, user, user, 1);

        vm.startPrank(admin);
        srcWrapper.revokeRole(burnerRole, address(srcBridge));
        destWrapper.revokeRole(burnerRole, address(destBridge));
        vm.stopPrank();

        vm.startPrank(user);
        // `srcBridge` must have `BURNER_ROLE` for `srcWrapper`.
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector, address(srcBridge), burnerRole
            )
        );
        srcBridge.sendToken{ value: fee }(DEST_CHAIN_ID, user, 1, true);

        fee = destBridge.bridgeFee(SRC_CHAIN_ID);

        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector, address(destBridge), burnerRole
            )
        );
        destBridge.sendToken{ value: fee }(SRC_CHAIN_ID, user, 1, false);
        vm.stopPrank();
    }

    function test_sendToken_originalToken_succeeds() public {
        mockBridge(srcBridge, SRC_CHAIN_ID, DEST_CHAIN_ID, true, user, user, INITIAL_USER_BALANCE);
        _assertBalances(user, 0, INITIAL_USER_BALANCE, 0, INITIAL_USER_BALANCE);
    }

    function test_sendToken_srcWrapper_succeeds() public {
        vm.prank(user);
        srcLockbox.deposit(INITIAL_USER_BALANCE);
        _assertBalances(user, 0, INITIAL_USER_BALANCE, INITIAL_USER_BALANCE, 0);

        mockBridge(srcBridge, SRC_CHAIN_ID, DEST_CHAIN_ID, false, user, user, INITIAL_USER_BALANCE);
        _assertBalances(user, 0, INITIAL_USER_BALANCE, 0, INITIAL_USER_BALANCE);
    }

    function test_sendToken_destWrapper_succeeds() public {
        mockBridge(srcBridge, SRC_CHAIN_ID, DEST_CHAIN_ID, true, user, user, INITIAL_USER_BALANCE);
        _assertBalances(user, 0, INITIAL_USER_BALANCE, 0, INITIAL_USER_BALANCE);

        mockBridge(destBridge, DEST_CHAIN_ID, SRC_CHAIN_ID, false, user, user, INITIAL_USER_BALANCE);
        _assertBalances(user, INITIAL_USER_BALANCE, 0, 0, 0);
    }

    function test_sendToken_destWrapper_empty_lockbox_succeeds() public {
        mockBridge(srcBridge, SRC_CHAIN_ID, DEST_CHAIN_ID, true, user, user, INITIAL_USER_BALANCE);
        _assertBalances(user, 0, INITIAL_USER_BALANCE, 0, INITIAL_USER_BALANCE);

        vm.startPrank(admin);
        destWrapper.grantRole(destWrapper.MINTER_ROLE(), admin);
        destWrapper.mint(user, INITIAL_USER_BALANCE);
        _assertBalances(user, 0, INITIAL_USER_BALANCE, 0, INITIAL_USER_BALANCE * 2);
        vm.stopPrank();

        mockBridge(destBridge, DEST_CHAIN_ID, SRC_CHAIN_ID, false, user, user, INITIAL_USER_BALANCE);
        _assertBalances(user, INITIAL_USER_BALANCE, 0, 0, INITIAL_USER_BALANCE);

        mockBridge(destBridge, DEST_CHAIN_ID, SRC_CHAIN_ID, false, user, user, INITIAL_USER_BALANCE);
        _assertBalances(user, INITIAL_USER_BALANCE, 0, INITIAL_USER_BALANCE, 0);
    }

    function test_sendToken_destWrapper_overdrafted_lockbox_succeeds() public {
        mockBridge(srcBridge, SRC_CHAIN_ID, DEST_CHAIN_ID, true, user, user, INITIAL_USER_BALANCE);
        _assertBalances(user, 0, INITIAL_USER_BALANCE, 0, INITIAL_USER_BALANCE);

        vm.startPrank(admin);
        destWrapper.grantRole(destWrapper.MINTER_ROLE(), admin);
        destWrapper.mint(user, INITIAL_USER_BALANCE);
        _assertBalances(user, 0, INITIAL_USER_BALANCE, 0, INITIAL_USER_BALANCE * 2);
        vm.stopPrank();

        mockBridge(destBridge, DEST_CHAIN_ID, SRC_CHAIN_ID, false, user, user, INITIAL_USER_BALANCE * 2);
        _assertBalances(user, 0, INITIAL_USER_BALANCE, INITIAL_USER_BALANCE * 2, 0);
    }

    function test_sendToken_fee_overpayment_refunded() public prank(user) {
        uint256 fee = srcBridge.bridgeFee(DEST_CHAIN_ID);
        uint256 balance = user.balance;
        bytes memory data = abi.encodeCall(Bridge.receiveToken, (user, 1));

        srcBridge.sendToken{ value: fee + 1 }(DEST_CHAIN_ID, user, 1, true);
        assertEq(user.balance, balance - fee, "Fee overpayment should be refunded");

        omni.mockXCall(SRC_CHAIN_ID, address(srcBridge), address(destBridge), data, DEFAULT_GAS_LIMIT);

        balance = user.balance;
        destBridge.sendToken{ value: fee + 1 }(SRC_CHAIN_ID, user, 1, false);
        assertEq(user.balance, balance - fee, "Fee overpayment should be refunded");
    }
}
