// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import { IOmniPortal } from "src/interfaces/IOmniPortal.sol";
import { ConfLevel } from "src/libraries/ConfLevel.sol";
import { OmniBridgeNative } from "src/token/OmniBridgeNative.sol";
import { OmniBridgeL1 } from "src/token/OmniBridgeL1.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";
import { MockPortal } from "test/utils/MockPortal.sol";
import { Test } from "forge-std/Test.sol";
import { VmSafe } from "forge-std/Vm.sol";

// solhint-disable state-visibility

/**
 * @title BridgeL1PostUpgradeTest
 * @dev Test OmniBridgeL1 post-upgrade functionality
 */
contract BridgeL1PostUpgradeTest is Test {
    OmniBridgeL1 b;
    MockPortal portal;
    IERC20 token;
    address owner;

    function run(address addr) public {
        (VmSafe.CallerMode mode,,) = vm.readCallers();
        require(mode == VmSafe.CallerMode.None, "no broadcast");

        _setup(addr);
        _testBridge();
        _testWithdraw();
        _testPauseUnpause();
    }

    function _setup(address addr) internal {
        b = OmniBridgeL1(addr);
        token = b.token();
        owner = b.owner();

        // OmniBridgeL1 portal at slot 0, no admin setters
        portal = new MockPortal();
        vm.store(addr, bytes32(0), bytes32(uint256(uint160(address(portal)))));

        // ensure bridge is fully unpaused prior to tests
        vm.startPrank(owner);
        if (b.isPaused(b.ACTION_BRIDGE())) b.unpause(b.ACTION_BRIDGE());
        if (b.isPaused(b.ACTION_WITHDRAW())) b.unpause(b.ACTION_WITHDRAW());
        vm.stopPrank();
    }

    function _testBridge() internal {
        address to = makeAddr("to");
        uint256 amount = 1e18;
        address payor = address(this);
        uint256 fee = b.bridgeFee(payor, to, amount);
        uint256 bridgeBalance = token.balanceOf(address(b));

        vm.expectCall(
            address(portal),
            fee,
            abi.encodeCall(
                IOmniPortal.xcall,
                (
                    portal.omniChainId(),
                    ConfLevel.Finalized,
                    Predeploys.OmniBridgeNative,
                    abi.encodeCall(OmniBridgeNative.withdraw, (payor, to, amount)),
                    b.XCALL_WITHDRAW_GAS_LIMIT()
                )
            )
        );

        deal(address(token), payor, amount);
        vm.deal(payor, fee);
        vm.startPrank(payor);
        token.approve(address(b), amount);
        b.bridge{ value: fee }(to, amount);
        vm.stopPrank();

        assertEq(token.balanceOf(address(b)), bridgeBalance + amount);
        assertEq(token.balanceOf(payor), 0);
    }

    function _testWithdraw() internal {
        address to = makeAddr("to");
        uint256 amount = 1e18;
        uint256 bridgeBalance = token.balanceOf(address(b));

        vm.expectCall(address(token), abi.encodeCall(token.transfer, (to, amount)));
        portal.mockXCall({
            sourceChainId: portal.omniChainId(),
            sender: Predeploys.OmniBridgeNative,
            to: address(b),
            data: abi.encodeCall(OmniBridgeL1.withdraw, (to, amount)),
            gasLimit: 100_000
        });

        assertEq(token.balanceOf(to), amount);
        assertEq(token.balanceOf(address(b)), bridgeBalance - amount);
    }

    function _testPauseUnpause() internal {
        vm.prank(owner);
        b.pause();

        assertTrue(b.isPaused(b.ACTION_BRIDGE()));
        assertTrue(b.isPaused(b.ACTION_WITHDRAW()));

        vm.prank(owner);
        b.unpause();

        assertFalse(b.isPaused(b.ACTION_BRIDGE()));
        assertFalse(b.isPaused(b.ACTION_WITHDRAW()));
    }
}
