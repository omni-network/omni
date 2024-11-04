// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { IOmniPortal } from "src/interfaces/IOmniPortal.sol";
import { ConfLevel } from "src/libraries/ConfLevel.sol";
import { OmniBridgeNative } from "src/token/OmniBridgeNative.sol";
import { OmniBridgeL1 } from "src/token/OmniBridgeL1.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";
import { MockPortal } from "test/utils/MockPortal.sol";
import { NoReceive } from "test/utils/NoReceive.sol";
import { Test } from "forge-std/Test.sol";
import { VmSafe } from "forge-std/Vm.sol";

// solhint-disable state-visibility

/**
 * @title BridgeNativePostUpgradeTest
 * @dev Test OmniBridgeNative post-upgrade functionality
 */
contract BridgeNativePostUpgradeTest is Test {
    OmniBridgeNative b;
    MockPortal portal;
    address l1Bridge;
    address owner;
    uint64 l1ChainId;

    function run() public {
        (VmSafe.CallerMode mode,,) = vm.readCallers();
        require(mode == VmSafe.CallerMode.None, "no broadcast");

        _setup();
        _testWithdraw(); // test withdraw() before bridge(), to update l1BridgeBalance
        _testBridge();
        _testClaim();
        _testPauseUnpause();
    }

    function _setup() internal {
        b = OmniBridgeNative(Predeploys.OmniBridgeNative);
        l1Bridge = b.l1Bridge();
        l1ChainId = b.l1ChainId();
        owner = b.owner();
        portal = new MockPortal();
        uint256 l1Deposits = b.l1Deposits();

        // change portal to mock portal
        vm.prank(owner);
        b.setup(l1ChainId, address(portal), l1Bridge, l1Deposits);
    }

    function _testWithdraw() internal {
        address to = makeAddr("to");
        uint256 amount = 1e18;
        address payor = makeAddr("payor");
        uint256 l1Deposits = b.l1Deposits();

        vm.expectCall(to, amount, "");

        portal.mockXCall({
            sourceChainId: l1ChainId,
            sender: address(l1Bridge),
            to: address(b),
            data: abi.encodeCall(OmniBridgeNative.withdraw, (payor, to, amount)),
            gasLimit: 100_000
        });

        assertEq(b.l1Deposits(), l1Deposits + amount);
        assertEq(b.claimable(payor), 0);
    }

    function _testBridge() internal {
        address to = makeAddr("to");
        uint256 amount = 1e18;
        uint256 fee = b.bridgeFee(to, amount);
        vm.expectCall(
            address(portal),
            fee,
            abi.encodeCall(
                IOmniPortal.xcall,
                (
                    l1ChainId,
                    ConfLevel.Finalized,
                    address(l1Bridge),
                    abi.encodeCall(OmniBridgeL1.withdraw, (to, amount)),
                    b.XCALL_WITHDRAW_GAS_LIMIT()
                )
            )
        );
        vm.deal(to, amount + fee);
        vm.prank(to);
        b.bridge{ value: amount + fee }(to, amount);
    }

    function _testClaim() internal {
        address to = makeAddr("to");
        uint256 amount = 1e18;
        address payor = makeAddr("payor");

        // will revert on withdraw
        address noReceiver = address(new NoReceive());

        // make claimable with failed withdraw
        vm.expectCall(noReceiver, amount, "");
        portal.mockXCall({
            sourceChainId: l1ChainId,
            sender: address(l1Bridge),
            to: address(b),
            data: abi.encodeCall(OmniBridgeNative.withdraw, (payor, noReceiver, amount)),
            gasLimit: 100_000
        });

        assertEq(b.claimable(payor), amount);

        // claim
        vm.expectCall(to, amount, "");
        portal.mockXCall({
            sourceChainId: l1ChainId,
            sender: payor,
            to: address(b),
            data: abi.encodeCall(OmniBridgeNative.claim, to),
            gasLimit: 100_000
        });

        assertEq(b.claimable(payor), 0);
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
