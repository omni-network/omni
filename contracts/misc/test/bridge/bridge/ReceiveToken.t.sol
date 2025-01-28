// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.26;

import "../TestBase.sol";

/**
 * @dev When triggering revert cases, gas needs to be 42_500 higher than the minimum necessary for successful transfers.
 */
contract ReceiveTokenTest is TestBase {
    function test_receiveToken_reverts() public {
        bytes memory data = abi.encodeCall(Bridge.receiveToken, (user, 1));
        uint64 unknownChainId = DEST_CHAIN_ID + 1;
        address unknownSender = makeAddr("unknownSender");
        bytes32 minterRole = srcWrapper.MINTER_ROLE();

        // Unknown source chain ID
        vm.expectRevert(abi.encodeWithSelector(IBridge.Unauthorized.selector, unknownChainId, address(destBridge)));
        omni.mockXCall({
            sourceChainId: unknownChainId,
            sender: address(destBridge),
            to: address(srcBridge),
            data: data,
            gasLimit: DEFAULT_GAS_LIMIT
        });

        // Unknown source chain sender
        vm.expectRevert(abi.encodeWithSelector(IBridge.Unauthorized.selector, DEST_CHAIN_ID, unknownSender));
        omni.mockXCall({
            sourceChainId: DEST_CHAIN_ID,
            sender: unknownSender,
            to: address(srcBridge),
            data: data,
            gasLimit: DEFAULT_GAS_LIMIT
        });

        // Unauthorized direct call
        vm.expectRevert(abi.encodeWithSelector(IBridge.Unauthorized.selector, SRC_CHAIN_ID, user));
        vm.prank(user);
        srcBridge.receiveToken(user, 1);

        // Reverts if `MINTER_ROLE` is revoked
        vm.startPrank(admin);
        srcWrapper.revokeRole(minterRole, address(srcBridge));
        destWrapper.revokeRole(minterRole, address(destBridge));
        vm.stopPrank();

        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector, address(srcBridge), minterRole
            )
        );
        omni.mockXCall({
            sourceChainId: DEST_CHAIN_ID,
            sender: address(destBridge),
            to: address(srcBridge),
            data: data,
            gasLimit: DEFAULT_GAS_LIMIT
        });

        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector, address(destBridge), minterRole
            )
        );
        omni.mockXCall({
            sourceChainId: SRC_CHAIN_ID,
            sender: address(srcBridge),
            to: address(destBridge),
            data: data,
            gasLimit: DEFAULT_GAS_LIMIT
        });
    }

    function test_receiveToken_succeeds_solvent_lockbox() public {
        bytes memory data = abi.encodeCall(Bridge.receiveToken, (user, INITIAL_USER_BALANCE));

        vm.prank(user);
        srcLockbox.deposit(INITIAL_USER_BALANCE);

        omni.mockXCall({
            sourceChainId: DEST_CHAIN_ID,
            sender: address(destBridge),
            to: address(srcBridge),
            data: data,
            gasLimit: DEFAULT_GAS_LIMIT
        });

        _assertBalances({
            addr: user,
            tokenUserBal: INITIAL_USER_BALANCE,
            tokenLockboxBal: 0,
            srcWrapperUserBal: INITIAL_USER_BALANCE,
            destWrapperUserBal: 0
        });
    }

    function test_receiveToken_succeeds_insolvent_lockbox() public {
        bytes memory data = abi.encodeCall(Bridge.receiveToken, (user, INITIAL_USER_BALANCE));

        omni.mockXCall({
            sourceChainId: DEST_CHAIN_ID,
            sender: address(destBridge),
            to: address(srcBridge),
            data: data,
            gasLimit: DEFAULT_GAS_LIMIT
        });

        _assertBalances({
            addr: user,
            tokenUserBal: INITIAL_USER_BALANCE,
            tokenLockboxBal: 0,
            srcWrapperUserBal: INITIAL_USER_BALANCE,
            destWrapperUserBal: 0
        });
    }

    function test_receiveToken_succeeds_no_lockbox() public {
        bytes memory data = abi.encodeCall(Bridge.receiveToken, (user, INITIAL_USER_BALANCE));

        omni.mockXCall({
            sourceChainId: SRC_CHAIN_ID,
            sender: address(srcBridge),
            to: address(destBridge),
            data: data,
            gasLimit: DEFAULT_GAS_LIMIT
        });

        _assertBalances({
            addr: user,
            tokenUserBal: INITIAL_USER_BALANCE,
            tokenLockboxBal: 0,
            srcWrapperUserBal: 0,
            destWrapperUserBal: INITIAL_USER_BALANCE
        });
    }

    function test_receiveToken_succeeds_paused() public prank(pauser) {
        bytes memory data = abi.encodeCall(Bridge.receiveToken, (user, INITIAL_USER_BALANCE));

        srcBridge.pause();
        destBridge.pause();

        omni.mockXCall({
            sourceChainId: DEST_CHAIN_ID,
            sender: address(destBridge),
            to: address(srcBridge),
            data: data,
            gasLimit: DEFAULT_GAS_LIMIT
        });
        omni.mockXCall({
            sourceChainId: SRC_CHAIN_ID,
            sender: address(srcBridge),
            to: address(destBridge),
            data: data,
            gasLimit: DEFAULT_GAS_LIMIT
        });

        _assertBalances({
            addr: user,
            tokenUserBal: INITIAL_USER_BALANCE,
            tokenLockboxBal: 0,
            srcWrapperUserBal: INITIAL_USER_BALANCE,
            destWrapperUserBal: INITIAL_USER_BALANCE
        });
    }
}
