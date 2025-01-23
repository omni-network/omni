// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.26;

import "./TestBase.sol";

contract BridgeTest is TestBase {
    using SafeTransferLib for address;

    function test_bridge_success() public {
        uint256 amount = INITIAL_USER_BALANCE / 2;

        // Split user's balance 50/50 between original and wrapped tokens.
        vm.prank(user);
        srcLockbox.deposit(amount);

        assertEq(originalToken.balanceOf(user), amount, "INIT: Original token balance mismatch");
        assertEq(originalToken.balanceOf(address(srcLockbox)), amount, "INIT: Source lockbox balance mismatch");
        assertEq(srcWrapper.balanceOf(user), amount, "INIT: Source wrapped token balance mismatch");
        assertEq(destWrapper.balanceOf(user), 0, "INIT: Destination wrapped token balance mismatch");

        // Bridge original tokens to the destination chain.
        mockBridge(srcBridge, SRC_CHAIN_ID, DEST_CHAIN_ID, true, user, user, amount);

        assertEq(originalToken.balanceOf(user), 0, "TX1: Original token balance mismatch");
        assertEq(
            originalToken.balanceOf(address(srcLockbox)), INITIAL_USER_BALANCE, "TX1: Source lockbox balance mismatch"
        );
        assertEq(srcWrapper.balanceOf(user), amount, "TX1: Source wrapped token balance mismatch");
        assertEq(destWrapper.balanceOf(user), amount, "TX1: Destination wrapped token balance mismatch");

        // Bridge wrapped tokens to the destination chain.
        mockBridge(srcBridge, SRC_CHAIN_ID, DEST_CHAIN_ID, false, user, user, amount);

        assertEq(originalToken.balanceOf(user), 0, "TX2: Original token balance mismatch");
        assertEq(
            originalToken.balanceOf(address(srcLockbox)), INITIAL_USER_BALANCE, "TX2: Source lockbox balance mismatch"
        );
        assertEq(srcWrapper.balanceOf(user), 0, "TX2: Source wrapped token balance mismatch");
        assertEq(destWrapper.balanceOf(user), INITIAL_USER_BALANCE, "TX2: Destination wrapped token balance mismatch");

        // Bridge all tokens back to the source chain where they should automatically unwrap.
        mockBridge(destBridge, DEST_CHAIN_ID, SRC_CHAIN_ID, false, user, user, INITIAL_USER_BALANCE);

        assertEq(originalToken.balanceOf(user), INITIAL_USER_BALANCE, "END: Original token balance mismatch");
        assertEq(originalToken.balanceOf(address(srcLockbox)), 0, "END: Source lockbox balance mismatch");
        assertEq(srcWrapper.balanceOf(user), 0, "END: Source wrapped token balance mismatch");
        assertEq(destWrapper.balanceOf(user), 0, "END: Destination wrapped token balance mismatch");
    }

    function test_bridge_empty_lockbox() public {
        assertEq(originalToken.balanceOf(user), INITIAL_USER_BALANCE, "INIT: Original token balance mismatch");
        assertEq(originalToken.balanceOf(address(srcLockbox)), 0, "INIT: Source lockbox balance mismatch");
        assertEq(srcWrapper.balanceOf(user), 0, "INIT: Source wrapped token balance mismatch");
        assertEq(destWrapper.balanceOf(user), 0, "INIT: Destination wrapped token balance mismatch");

        // Bridge all of user's tokens to the destination chain.
        mockBridge(srcBridge, SRC_CHAIN_ID, DEST_CHAIN_ID, true, user, user, INITIAL_USER_BALANCE);

        assertEq(originalToken.balanceOf(user), 0, "TX1: Original token balance mismatch");
        assertEq(
            originalToken.balanceOf(address(srcLockbox)), INITIAL_USER_BALANCE, "TX1: Source lockbox balance mismatch"
        );
        assertEq(srcWrapper.balanceOf(user), 0, "TX1: Source wrapped token balance mismatch");
        assertEq(destWrapper.balanceOf(user), INITIAL_USER_BALANCE, "TX1: Destination wrapped token balance mismatch");

        // Incorrectly mint destination wrapped tokens to the user.
        vm.startPrank(admin);
        destWrapper.grantRole(destWrapper.MINTER_ROLE(), admin);
        destWrapper.mint(user, INITIAL_USER_BALANCE);
        vm.stopPrank();

        assertEq(originalToken.balanceOf(user), 0, "TX2: Original token balance mismatch");
        assertEq(
            originalToken.balanceOf(address(srcLockbox)), INITIAL_USER_BALANCE, "TX2: Source lockbox balance mismatch"
        );
        assertEq(srcWrapper.balanceOf(user), 0, "TX2: Source wrapped token balance mismatch");
        assertEq(
            destWrapper.balanceOf(user), INITIAL_USER_BALANCE * 2, "TX2: Destination wrapped token balance mismatch"
        );

        // Bridge user's original balance of wrapped tokens back to the source chain.
        mockBridge(destBridge, DEST_CHAIN_ID, SRC_CHAIN_ID, false, user, user, INITIAL_USER_BALANCE);

        assertEq(originalToken.balanceOf(user), INITIAL_USER_BALANCE, "TX3: Original token balance mismatch");
        assertEq(originalToken.balanceOf(address(srcLockbox)), 0, "TX3: Source lockbox balance mismatch");
        assertEq(srcWrapper.balanceOf(user), 0, "TX3: Source wrapped token balance mismatch");
        assertEq(destWrapper.balanceOf(user), INITIAL_USER_BALANCE, "TX3: Destination wrapped token balance mismatch");

        // Bridge user's duplicated wrapped tokens back to the source chain.
        mockBridge(destBridge, DEST_CHAIN_ID, SRC_CHAIN_ID, false, user, user, INITIAL_USER_BALANCE);

        // Confirm the user received wrapped tokens back on the source chain.
        assertEq(originalToken.balanceOf(user), INITIAL_USER_BALANCE, "END: Original token balance mismatch");
        assertEq(originalToken.balanceOf(address(srcLockbox)), 0, "END: Source lockbox balance mismatch");
        assertEq(srcWrapper.balanceOf(user), INITIAL_USER_BALANCE, "END: Source wrapped token balance mismatch");
        assertEq(destWrapper.balanceOf(user), 0, "END: Destination wrapped token balance mismatch");
    }
}
