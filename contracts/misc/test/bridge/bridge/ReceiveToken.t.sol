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

        // Unknown source chain ID
        vm.expectRevert(abi.encodeWithSelector(IBridge.Unauthorized.selector, unknownChainId, address(destBridge)));
        omni.mockXCall(unknownChainId, address(destBridge), address(srcBridge), data, DEFAULT_GAS_LIMIT);

        // Unknown source chain sender
        vm.expectRevert(abi.encodeWithSelector(IBridge.Unauthorized.selector, DEST_CHAIN_ID, unknownSender));
        omni.mockXCall(DEST_CHAIN_ID, unknownSender, address(srcBridge), data, DEFAULT_GAS_LIMIT);

        // Unauthorized direct call
        vm.expectRevert(abi.encodeWithSelector(IBridge.Unauthorized.selector, SRC_CHAIN_ID, user));
        vm.prank(user);
        srcBridge.receiveToken(user, 1);

        // Message delivers if sender is `OmniPortal`
        omni.mockXCall(DEST_CHAIN_ID, address(destBridge), address(srcBridge), data, DEFAULT_GAS_LIMIT);
    }

    function test_receiveToken_reverts_without_wrapper_minter_role() public {
        bytes memory data = abi.encodeCall(Bridge.receiveToken, (user, INITIAL_USER_BALANCE));
        bytes32 minterRole = srcWrapper.MINTER_ROLE();

        vm.startPrank(admin);
        srcWrapper.revokeRole(minterRole, address(srcBridge));
        destWrapper.revokeRole(minterRole, address(destBridge));
        vm.stopPrank();

        vm.startPrank(user);
        // `srcBridge` must have `BURNER_ROLE` for `srcWrapper`.
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector, address(srcBridge), minterRole
            )
        );
        omni.mockXCall(DEST_CHAIN_ID, address(destBridge), address(srcBridge), data, DEFAULT_GAS_LIMIT);

        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector, address(destBridge), minterRole
            )
        );
        omni.mockXCall(SRC_CHAIN_ID, address(srcBridge), address(destBridge), data, DEFAULT_GAS_LIMIT);
        vm.stopPrank();
    }

    function test_receiveToken_succeeds_solvent_lockbox() public {
        bytes memory data = abi.encodeCall(Bridge.receiveToken, (user, INITIAL_USER_BALANCE));

        vm.prank(user);
        srcLockbox.deposit(INITIAL_USER_BALANCE);

        omni.mockXCall(DEST_CHAIN_ID, address(destBridge), address(srcBridge), data, DEFAULT_GAS_LIMIT);

        _assertBalances(user, INITIAL_USER_BALANCE, 0, INITIAL_USER_BALANCE, 0);
    }

    function test_receiveToken_succeeds_insolvent_lockbox() public {
        bytes memory data = abi.encodeCall(Bridge.receiveToken, (user, INITIAL_USER_BALANCE));

        omni.mockXCall(DEST_CHAIN_ID, address(destBridge), address(srcBridge), data, DEFAULT_GAS_LIMIT);

        _assertBalances(user, INITIAL_USER_BALANCE, 0, INITIAL_USER_BALANCE, 0);
    }

    function test_receiveToken_succeeds_no_lockbox() public {
        bytes memory data = abi.encodeCall(Bridge.receiveToken, (user, INITIAL_USER_BALANCE));

        omni.mockXCall(SRC_CHAIN_ID, address(srcBridge), address(destBridge), data, DEFAULT_GAS_LIMIT);

        _assertBalances(user, INITIAL_USER_BALANCE, 0, 0, INITIAL_USER_BALANCE);
    }

    function test_receiveToken_succeeds_paused() public prank(pauser) {
        bytes memory data = abi.encodeCall(Bridge.receiveToken, (user, INITIAL_USER_BALANCE));

        srcBridge.pause();
        destBridge.pause();

        omni.mockXCall(DEST_CHAIN_ID, address(destBridge), address(srcBridge), data, DEFAULT_GAS_LIMIT);
        omni.mockXCall(SRC_CHAIN_ID, address(srcBridge), address(destBridge), data, DEFAULT_GAS_LIMIT);

        _assertBalances(user, INITIAL_USER_BALANCE, 0, INITIAL_USER_BALANCE, INITIAL_USER_BALANCE);
    }
}
