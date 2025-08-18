// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import { IOmniPortal } from "src/interfaces/IOmniPortal.sol";
import { ConfLevel } from "src/libraries/ConfLevel.sol";
import { NominaBridgeNative } from "src/token/nomina/NominaBridgeNative.sol";
import { NominaBridgeL1 } from "src/token/nomina/NominaBridgeL1.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";
import { MockPortal } from "test/utils/MockPortal.sol";
import { Test } from "forge-std/Test.sol";
import { VmSafe } from "forge-std/Vm.sol";

// solhint-disable state-visibility

/**
 * @title BridgeL1PostUpgradeTest
 * @dev Test NominaBridgeL1 post-upgrade functionality
 */
contract BridgeL1PostUpgradeTest is Test {
    NominaBridgeL1 b;
    MockPortal portal;
    IERC20 omni;
    IERC20 nomina;
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
        b = NominaBridgeL1(addr);
        omni = b.omni();
        nomina = b.nomina();
        owner = b.owner();

        // NominaBridgeL1 portal at slot 0, no admin setters
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
        uint256 bridgeBalance = omni.balanceOf(address(b));

        vm.expectCall(
            address(portal),
            fee,
            abi.encodeCall(
                IOmniPortal.xcall,
                (
                    portal.omniChainId(),
                    ConfLevel.Finalized,
                    Predeploys.OmniBridgeNative,
                    abi.encodeCall(NominaBridgeNative.withdraw, (payor, to, amount)),
                    b.XCALL_WITHDRAW_GAS_LIMIT()
                )
            )
        );

        deal(address(omni), payor, amount);
        vm.deal(payor, fee);
        vm.startPrank(payor);
        omni.approve(address(b), amount);
        b.bridge{ value: fee }(to, amount);
        vm.stopPrank();

        assertEq(omni.balanceOf(address(b)), bridgeBalance + amount);
        assertEq(omni.balanceOf(payor), 0);
    }

    function _testWithdraw() internal {
        address to = makeAddr("to");
        uint256 amount = 1e18;
        uint256 bridgeBalance = omni.balanceOf(address(b));

        vm.expectCall(address(omni), abi.encodeCall(omni.transfer, (to, amount)));
        portal.mockXCall({
            sourceChainId: portal.omniChainId(),
            sender: Predeploys.OmniBridgeNative,
            to: address(b),
            data: abi.encodeCall(NominaBridgeL1.withdraw, (to, amount)),
            gasLimit: 100_000
        });

        assertEq(omni.balanceOf(to), amount);
        assertEq(omni.balanceOf(address(b)), bridgeBalance - amount);
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
