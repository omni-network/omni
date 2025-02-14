// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.26;

import "../TestBase.sol";

contract RetryTransferTest is TestBase {
    bytes data = abi.encodeCall(Bridge.receiveToken, (user, 1));
    bytes32 minterRole;

    function setUp() public override {
        super.setUp();
        minterRole = wrapper.MINTER_ROLE();

        vm.startPrank(admin);
        wrapper.revokeRole(minterRole, address(bridgeWithLockbox));
        wrapper.revokeRole(minterRole, address(bridgeNoLockbox));
        vm.stopPrank();
    }

    function test_retryTransfer_reverts() public {
        // Reverts if no claimable tokens
        vm.expectRevert(abi.encodeWithSelector(IBridge.NoClaimable.selector));
        bridgeWithLockbox.retryTransfer(user);

        omni.mockXCall({
            sourceChainId: DEST_CHAIN_ID,
            sender: address(bridgeNoLockbox),
            to: address(bridgeWithLockbox),
            data: data,
            gasLimit: _getGasLimit(Bridge(bridgeWithLockbox))
        });

        // Reverts if no minter role
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector, address(bridgeWithLockbox), minterRole
            )
        );
        bridgeWithLockbox.retryTransfer(user);
    }

    function test_receiveToken_caches_tokens_when_mint_reverts() public {
        bytes32 minterRole = wrapper.MINTER_ROLE();
        bytes memory data = abi.encodeCall(Bridge.receiveToken, (user, 1));

        vm.startPrank(admin);
        wrapper.revokeRole(minterRole, address(bridgeWithLockbox));
        wrapper.revokeRole(minterRole, address(bridgeNoLockbox));
        vm.stopPrank();

        omni.mockXCall({
            sourceChainId: DEST_CHAIN_ID,
            sender: address(bridgeNoLockbox),
            to: address(bridgeWithLockbox),
            data: data,
            gasLimit: _getGasLimit(Bridge(bridgeWithLockbox))
        });

        omni.mockXCall({
            sourceChainId: SRC_CHAIN_ID,
            sender: address(bridgeWithLockbox),
            to: address(bridgeNoLockbox),
            data: data,
            gasLimit: _getGasLimit(Bridge(bridgeNoLockbox))
        });

        _assertBalances({ addr: user, tokenUserBal: INITIAL_USER_BALANCE, tokenLockboxBal: 0, wrapperUserBal: 0 });

        assertEq(bridgeWithLockbox.claimable(user), 1, "bridgeWithLockbox claimable");
        assertEq(wrapper.balanceOf(address(bridgeWithLockbox)), 0, "bridgeWithLockbox balance");

        assertEq(bridgeNoLockbox.claimable(user), 1, "bridgeNoLockbox claimable");
        assertEq(wrapper.balanceOf(address(bridgeNoLockbox)), 0, "bridgeNoLockbox balance");
    }

    function test_retryTransfer_succeeds() public {
        // Send transactions that will get cached due to no mint permissions
        omni.mockXCall({
            sourceChainId: DEST_CHAIN_ID,
            sender: address(bridgeNoLockbox),
            to: address(bridgeWithLockbox),
            data: data,
            gasLimit: _getGasLimit(Bridge(bridgeWithLockbox))
        });
        omni.mockXCall({
            sourceChainId: SRC_CHAIN_ID,
            sender: address(bridgeWithLockbox),
            to: address(bridgeNoLockbox),
            data: data,
            gasLimit: _getGasLimit(Bridge(bridgeNoLockbox))
        });

        // Grant mint permissions
        vm.startPrank(admin);
        wrapper.grantRole(minterRole, address(bridgeWithLockbox));
        wrapper.grantRole(minterRole, address(bridgeNoLockbox));
        vm.stopPrank();

        // Deal a token to the lockbox so it doesn't revert as bridgeWithLockbox will try to unwrap
        vm.prank(minter);
        token.mint(address(lockbox), 1);

        bridgeWithLockbox.retryTransfer(user);
        bridgeNoLockbox.retryTransfer(user);

        assertEq(bridgeWithLockbox.claimable(user), 0, "bridgeWithLockbox claimable");
        assertEq(wrapper.balanceOf(address(bridgeWithLockbox)), 0, "bridgeWithLockbox balance");

        assertEq(bridgeNoLockbox.claimable(user), 0, "bridgeNoLockbox claimable");
        assertEq(wrapper.balanceOf(address(bridgeNoLockbox)), 0, "bridgeNoLockbox balance");

        _assertBalances({ addr: user, tokenUserBal: INITIAL_USER_BALANCE + 1, tokenLockboxBal: 0, wrapperUserBal: 1 });
    }
}
