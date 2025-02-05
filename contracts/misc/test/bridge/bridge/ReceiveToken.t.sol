// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.26;

import "../TestBase.sol";

/**
 * @dev When triggering revert cases, gas needs to be 60_000 higher than the minimum necessary for successful transfers.
 */
contract ReceiveTokenTest is TestBase {
    function test_receiveToken_reverts() public {
        bytes memory data = abi.encodeCall(Bridge.receiveToken, (user, 1));
        uint64 unknownChainId = DEST_CHAIN_ID + 1;
        address unknownSender = makeAddr("unknownSender");
        bytes32 minterRole = wrapper.MINTER_ROLE();

        // Unknown source chain ID
        vm.expectRevert(abi.encodeWithSelector(IBridge.Unauthorized.selector, unknownChainId, address(bridgeNoLockbox)));
        omni.mockXCall({
            sourceChainId: unknownChainId,
            sender: address(bridgeNoLockbox),
            to: address(bridgeWithLockbox),
            data: data,
            gasLimit: DEFAULT_GAS_LIMIT
        });

        // Unknown source chain sender
        vm.expectRevert(abi.encodeWithSelector(IBridge.Unauthorized.selector, DEST_CHAIN_ID, unknownSender));
        omni.mockXCall({
            sourceChainId: DEST_CHAIN_ID,
            sender: unknownSender,
            to: address(bridgeWithLockbox),
            data: data,
            gasLimit: DEFAULT_GAS_LIMIT
        });

        // Unauthorized direct call
        vm.expectRevert(abi.encodeWithSelector(IBridge.Unauthorized.selector, SRC_CHAIN_ID, user));
        vm.prank(user);
        bridgeWithLockbox.receiveToken(user, 1);

        // Reverts if `MINTER_ROLE` is revoked
        vm.startPrank(admin);
        wrapper.revokeRole(minterRole, address(bridgeWithLockbox));
        wrapper.revokeRole(minterRole, address(bridgeNoLockbox));
        vm.stopPrank();

        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector, address(bridgeWithLockbox), minterRole
            )
        );
        omni.mockXCall({
            sourceChainId: DEST_CHAIN_ID,
            sender: address(bridgeNoLockbox),
            to: address(bridgeWithLockbox),
            data: data,
            gasLimit: DEFAULT_GAS_LIMIT
        });

        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector, address(bridgeNoLockbox), minterRole
            )
        );
        omni.mockXCall({
            sourceChainId: SRC_CHAIN_ID,
            sender: address(bridgeWithLockbox),
            to: address(bridgeNoLockbox),
            data: data,
            gasLimit: DEFAULT_GAS_LIMIT
        });
    }

    function test_receiveToken_succeeds_solvent_lockbox() public {
        bytes memory data = abi.encodeCall(Bridge.receiveToken, (user, INITIAL_USER_BALANCE));

        vm.prank(user);
        lockbox.deposit(INITIAL_USER_BALANCE);

        omni.mockXCall({
            sourceChainId: DEST_CHAIN_ID,
            sender: address(bridgeNoLockbox),
            to: address(bridgeWithLockbox),
            data: data,
            gasLimit: DEFAULT_GAS_LIMIT
        });

        _assertBalances({
            addr: user,
            tokenUserBal: INITIAL_USER_BALANCE,
            tokenLockboxBal: 0,
            wrapperUserBal: INITIAL_USER_BALANCE
        });
    }

    function test_receiveToken_succeeds_insolvent_lockbox() public {
        bytes memory data = abi.encodeCall(Bridge.receiveToken, (user, INITIAL_USER_BALANCE));

        omni.mockXCall({
            sourceChainId: DEST_CHAIN_ID,
            sender: address(bridgeNoLockbox),
            to: address(bridgeWithLockbox),
            data: data,
            gasLimit: DEFAULT_GAS_LIMIT
        });

        _assertBalances({
            addr: user,
            tokenUserBal: INITIAL_USER_BALANCE,
            tokenLockboxBal: 0,
            wrapperUserBal: INITIAL_USER_BALANCE
        });
    }

    function test_receiveToken_succeeds_no_lockbox() public {
        bytes memory data = abi.encodeCall(Bridge.receiveToken, (user, INITIAL_USER_BALANCE));

        omni.mockXCall({
            sourceChainId: SRC_CHAIN_ID,
            sender: address(bridgeWithLockbox),
            to: address(bridgeNoLockbox),
            data: data,
            gasLimit: DEFAULT_GAS_LIMIT
        });

        _assertBalances({
            addr: user,
            tokenUserBal: INITIAL_USER_BALANCE,
            tokenLockboxBal: 0,
            wrapperUserBal: INITIAL_USER_BALANCE
        });
    }

    function test_receiveToken_succeeds_paused() public prank(pauser) {
        bytes memory data = abi.encodeCall(Bridge.receiveToken, (user, INITIAL_USER_BALANCE));

        bridgeWithLockbox.pause();
        bridgeNoLockbox.pause();

        omni.mockXCall({
            sourceChainId: DEST_CHAIN_ID,
            sender: address(bridgeNoLockbox),
            to: address(bridgeWithLockbox),
            data: data,
            gasLimit: DEFAULT_GAS_LIMIT
        });
        omni.mockXCall({
            sourceChainId: SRC_CHAIN_ID,
            sender: address(bridgeWithLockbox),
            to: address(bridgeNoLockbox),
            data: data,
            gasLimit: DEFAULT_GAS_LIMIT
        });

        _assertBalances({
            addr: user,
            tokenUserBal: INITIAL_USER_BALANCE,
            tokenLockboxBal: 0,
            wrapperUserBal: INITIAL_USER_BALANCE * 2
        });
    }
}
