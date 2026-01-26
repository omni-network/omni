// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.26;

import "../TestBase.sol";
import { IERC20Errors } from "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";

contract SendTokenTest is TestBase {
    function test_sendToken_reverts() public {
        uint256 fee = bridgeWithLockbox.bridgeFee(DEST_CHAIN_ID);
        bytes32 clawbackerRole = token.CLAWBACKER_ROLE();

        mockBridge({
            bridge: bridgeWithLockbox,
            srcChainId: SRC_CHAIN_ID,
            destChainId: DEST_CHAIN_ID,
            wrap: true,
            refundTo: user,
            from: user,
            to: user,
            value: 1
        });

        vm.startPrank(user);
        // `destChainId` must have a configured route.
        vm.expectRevert(abi.encodeWithSelector(IBridge.InvalidRoute.selector, 0));
        bridgeWithLockbox.sendToken{ value: fee }(0, address(0), 0, true, address(0));
        vm.expectRevert(abi.encodeWithSelector(IBridge.InvalidRoute.selector, 0));
        bridgeWithLockbox.sendToken{ value: fee }(0, address(0), 0, false, address(0));

        // `to` cannot be zero address.
        vm.expectRevert(abi.encodeWithSelector(IBridge.ZeroAddress.selector));
        bridgeWithLockbox.sendToken{ value: fee }(DEST_CHAIN_ID, address(0), 0, true, address(0));
        vm.expectRevert(abi.encodeWithSelector(IBridge.ZeroAddress.selector));
        bridgeWithLockbox.sendToken{ value: fee }(DEST_CHAIN_ID, address(0), 0, false, address(0));

        // `value` cannot be zero.
        vm.expectRevert(abi.encodeWithSelector(IBridge.ZeroAmount.selector));
        bridgeWithLockbox.sendToken{ value: fee }(DEST_CHAIN_ID, user, 0, true, address(0));
        vm.expectRevert(abi.encodeWithSelector(IBridge.ZeroAmount.selector));
        bridgeWithLockbox.sendToken{ value: fee }(DEST_CHAIN_ID, user, 0, false, address(0));

        // `wrap` cannot be true if `lockbox` is not set.
        vm.expectRevert(abi.encodeWithSelector(IBridge.CannotWrap.selector));
        bridgeNoLockbox.sendToken{ value: fee }(SRC_CHAIN_ID, user, 1, true, address(0));

        // `refundTo` cannot be zero address.
        vm.expectRevert(abi.encodeWithSelector(IBridge.ZeroAddress.selector));
        bridgeWithLockbox.sendToken{ value: fee }(DEST_CHAIN_ID, user, 1, true, address(0));

        // `amount` cannot exceed the user's balance.
        // wrap = true reverts with SafeTransferLib.TransferFromFailed (via token.transferFrom(...))
        vm.expectRevert(abi.encodeWithSelector(SafeTransferLib.TransferFromFailed.selector));
        bridgeWithLockbox.sendToken{ value: fee }(DEST_CHAIN_ID, user, INITIAL_USER_BALANCE + 1, true, user);

        // wrap = false reverts with ERC20InsufficientBalance (via xtoken.clawback())
        vm.expectRevert(
            abi.encodeWithSelector(IERC20Errors.ERC20InsufficientBalance.selector, user, 1, INITIAL_USER_BALANCE + 1)
        );
        bridgeWithLockbox.sendToken{ value: fee }(DEST_CHAIN_ID, user, INITIAL_USER_BALANCE + 1, false, user);

        // `bridgeFee` must be paid.
        vm.expectRevert("XApp: insufficient funds");
        bridgeWithLockbox.sendToken{ value: fee - 1 }(DEST_CHAIN_ID, user, 1, true, user);

        lockbox.deposit(1);
        vm.expectRevert("XApp: insufficient funds");
        bridgeWithLockbox.sendToken{ value: fee - 1 }(DEST_CHAIN_ID, user, 1, false, user);
        vm.stopPrank();

        // Reverts if bridge is paused
        vm.startPrank(pauser);
        bridgeWithLockbox.pause();
        bridgeNoLockbox.pause();
        vm.stopPrank();

        vm.startPrank(user);
        vm.expectRevert(PausableUpgradeable.EnforcedPause.selector);
        bridgeWithLockbox.sendToken{ value: fee }(DEST_CHAIN_ID, user, 1, true, user);
        vm.expectRevert(PausableUpgradeable.EnforcedPause.selector);
        bridgeWithLockbox.sendToken{ value: fee }(DEST_CHAIN_ID, user, 1, false, user);

        vm.expectRevert(PausableUpgradeable.EnforcedPause.selector);
        bridgeNoLockbox.sendToken{ value: fee }(SRC_CHAIN_ID, user, 1, true, user);
        vm.expectRevert(PausableUpgradeable.EnforcedPause.selector);
        bridgeNoLockbox.sendToken{ value: fee }(SRC_CHAIN_ID, user, 1, false, user);
        vm.stopPrank();

        vm.startPrank(unpauser);
        bridgeWithLockbox.unpause();
        bridgeNoLockbox.unpause();
        vm.stopPrank();

        // Reverts if `CLAWBACKER_ROLE` is revoked
        vm.startPrank(admin);
        wrapper.revokeRole(clawbackerRole, address(bridgeWithLockbox));
        wrapper.revokeRole(clawbackerRole, address(bridgeNoLockbox));
        vm.stopPrank();

        vm.startPrank(user);
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector, address(bridgeWithLockbox), clawbackerRole
            )
        );
        bridgeWithLockbox.sendToken{ value: fee }(DEST_CHAIN_ID, user, 1, true, user);

        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector, address(bridgeNoLockbox), clawbackerRole
            )
        );
        bridgeNoLockbox.sendToken{ value: fee }(SRC_CHAIN_ID, user, 1, false, user);
        vm.stopPrank();
    }

    function test_sendToken_withLockbox_token_succeeds() public {
        mockBridge({
            bridge: bridgeWithLockbox,
            srcChainId: SRC_CHAIN_ID,
            destChainId: DEST_CHAIN_ID,
            wrap: true,
            refundTo: user,
            from: user,
            to: user,
            value: INITIAL_USER_BALANCE
        });
        _assertBalances({
            addr: user, tokenUserBal: 0, tokenLockboxBal: INITIAL_USER_BALANCE, wrapperUserBal: INITIAL_USER_BALANCE
        });
    }

    function test_sendToken_withLockbox_wrapper_succeeds() public {
        vm.prank(user);
        lockbox.deposit(INITIAL_USER_BALANCE);
        _assertBalances({
            addr: user, tokenUserBal: 0, tokenLockboxBal: INITIAL_USER_BALANCE, wrapperUserBal: INITIAL_USER_BALANCE
        });

        mockBridge({
            bridge: bridgeWithLockbox,
            srcChainId: SRC_CHAIN_ID,
            destChainId: DEST_CHAIN_ID,
            wrap: false,
            refundTo: user,
            from: user,
            to: user,
            value: INITIAL_USER_BALANCE
        });
        _assertBalances({
            addr: user, tokenUserBal: 0, tokenLockboxBal: INITIAL_USER_BALANCE, wrapperUserBal: INITIAL_USER_BALANCE
        });
    }

    function test_sendToken_withoutLockbox_wrapper_succeeds() public {
        vm.prank(user);
        lockbox.deposit(INITIAL_USER_BALANCE);
        _assertBalances({
            addr: user, tokenUserBal: 0, tokenLockboxBal: INITIAL_USER_BALANCE, wrapperUserBal: INITIAL_USER_BALANCE
        });

        mockBridge({
            bridge: bridgeNoLockbox,
            srcChainId: DEST_CHAIN_ID,
            destChainId: SRC_CHAIN_ID,
            wrap: false,
            refundTo: user,
            from: user,
            to: user,
            value: INITIAL_USER_BALANCE
        });
        _assertBalances({ addr: user, tokenUserBal: INITIAL_USER_BALANCE, tokenLockboxBal: 0, wrapperUserBal: 0 });
    }

    function test_sendToken_empty_lockbox_succeeds() public {
        mockBridge({
            bridge: bridgeWithLockbox,
            srcChainId: SRC_CHAIN_ID,
            destChainId: DEST_CHAIN_ID,
            wrap: true,
            refundTo: user,
            from: user,
            to: user,
            value: INITIAL_USER_BALANCE
        });
        _assertBalances({
            addr: user, tokenUserBal: 0, tokenLockboxBal: INITIAL_USER_BALANCE, wrapperUserBal: INITIAL_USER_BALANCE
        });

        vm.startPrank(admin);
        wrapper.grantRole(wrapper.MINTER_ROLE(), admin);
        wrapper.mint(user, INITIAL_USER_BALANCE);
        _assertBalances({
            addr: user, tokenUserBal: 0, tokenLockboxBal: INITIAL_USER_BALANCE, wrapperUserBal: INITIAL_USER_BALANCE * 2
        });
        vm.stopPrank();

        mockBridge({
            bridge: bridgeNoLockbox,
            srcChainId: DEST_CHAIN_ID,
            destChainId: SRC_CHAIN_ID,
            wrap: false,
            refundTo: user,
            from: user,
            to: user,
            value: INITIAL_USER_BALANCE
        });
        _assertBalances({
            addr: user, tokenUserBal: INITIAL_USER_BALANCE, tokenLockboxBal: 0, wrapperUserBal: INITIAL_USER_BALANCE
        });

        mockBridge({
            bridge: bridgeNoLockbox,
            srcChainId: DEST_CHAIN_ID,
            destChainId: SRC_CHAIN_ID,
            wrap: false,
            refundTo: user,
            from: user,
            to: user,
            value: INITIAL_USER_BALANCE
        });
        _assertBalances({
            addr: user, tokenUserBal: INITIAL_USER_BALANCE, tokenLockboxBal: 0, wrapperUserBal: INITIAL_USER_BALANCE
        });
    }

    function test_sendToken_wrapper_overdrafted_lockbox_succeeds() public {
        mockBridge({
            bridge: bridgeWithLockbox,
            srcChainId: SRC_CHAIN_ID,
            destChainId: DEST_CHAIN_ID,
            wrap: true,
            refundTo: user,
            from: user,
            to: user,
            value: INITIAL_USER_BALANCE
        });
        _assertBalances({
            addr: user, tokenUserBal: 0, tokenLockboxBal: INITIAL_USER_BALANCE, wrapperUserBal: INITIAL_USER_BALANCE
        });

        vm.startPrank(admin);
        wrapper.grantRole(wrapper.MINTER_ROLE(), admin);
        wrapper.mint(user, INITIAL_USER_BALANCE);
        _assertBalances({
            addr: user, tokenUserBal: 0, tokenLockboxBal: INITIAL_USER_BALANCE, wrapperUserBal: INITIAL_USER_BALANCE * 2
        });
        vm.stopPrank();

        mockBridge({
            bridge: bridgeNoLockbox,
            srcChainId: DEST_CHAIN_ID,
            destChainId: SRC_CHAIN_ID,
            wrap: false,
            refundTo: user,
            from: user,
            to: user,
            value: INITIAL_USER_BALANCE * 2
        });
        _assertBalances({
            addr: user, tokenUserBal: 0, tokenLockboxBal: INITIAL_USER_BALANCE, wrapperUserBal: INITIAL_USER_BALANCE * 2
        });
    }

    function test_sendToken_fee_overpayment_refunded() public prank(user) {
        uint256 fee = bridgeWithLockbox.bridgeFee(DEST_CHAIN_ID);
        uint256 balance = user.balance;
        bytes memory data = abi.encodeCall(Bridge.receiveToken, (user, 1));

        bridgeWithLockbox.sendToken{ value: fee + 1 }(DEST_CHAIN_ID, user, 1, true, user);
        assertEq(user.balance, balance - fee, "Fee overpayment should be refunded");

        omni.mockXCall({
            sourceChainId: SRC_CHAIN_ID,
            sender: address(bridgeWithLockbox),
            to: address(bridgeNoLockbox),
            data: data,
            gasLimit: _getGasLimit(Bridge(bridgeNoLockbox))
        });

        balance = user.balance;
        bridgeNoLockbox.sendToken{ value: fee + 1 }(SRC_CHAIN_ID, user, 1, false, user);
        assertEq(user.balance, balance - fee, "Fee overpayment should be refunded");
    }

    function test_sendToken_fee_overpayment_refunded_refundTo() public prank(user) {
        uint256 fee = bridgeWithLockbox.bridgeFee(DEST_CHAIN_ID);
        bytes memory data = abi.encodeCall(Bridge.receiveToken, (user, 1));

        bridgeWithLockbox.sendToken{ value: fee + 1 }(DEST_CHAIN_ID, user, 1, true, other);
        assertEq(other.balance, 1, "Fee overpayment should be refunded");

        omni.mockXCall({
            sourceChainId: SRC_CHAIN_ID,
            sender: address(bridgeWithLockbox),
            to: address(bridgeNoLockbox),
            data: data,
            gasLimit: _getGasLimit(Bridge(bridgeNoLockbox))
        });

        bridgeNoLockbox.sendToken{ value: fee + 1 }(SRC_CHAIN_ID, user, 1, false, other);
        assertEq(other.balance, 2, "Fee overpayment should be refunded");
    }
}
