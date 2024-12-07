// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { MockPortal } from "test/utils/MockPortal.sol";
import { NoReceive } from "test/utils/NoReceive.sol";
import { IOmniPortal } from "src/interfaces/IOmniPortal.sol";
import { OmniBridgeNative } from "src/token/OmniBridgeNative.sol";
import { OmniBridgeL1 } from "src/token/OmniBridgeL1.sol";
import { AddressUtils } from "src/libraries/AddressUtils.sol";
import { ConfLevel } from "src/libraries/ConfLevel.sol";
import { Test } from "forge-std/Test.sol";
import { console } from "forge-std/console.sol";

/**
 * @title OmniBridgeNative_Test
 * @notice Test suite for OmniBridgeNative contract.
 */
contract OmniBridgeNative_Test is Test {
    using AddressUtils for address;

    // Events copied from OmniBridgeNative.sol
    event Bridge(address indexed payor, address indexed to, uint256 amount);
    event Withdraw(address indexed payor, address indexed to, uint256 amount, bool success);
    event Claimed(address indexed claimant, address indexed to, uint256 amount);

    MockPortal portal;
    OmniBridgeNativeHarness b;
    OmniBridgeL1 l1Bridge;
    address owner;

    uint64 l1ChainId;
    uint256 totalSupply = 100_000_000 * 10 ** 18;

    function setUp() public {
        portal = new MockPortal();
        l1ChainId = 1;
        l1Bridge = new OmniBridgeL1(makeAddr("token"));
        owner = makeAddr("owner");

        address impl = address(new OmniBridgeNativeHarness());
        b = OmniBridgeNativeHarness(
            address(
                new TransparentUpgradeableProxy(
                    impl, owner, abi.encodeWithSelector(OmniBridgeNative.initialize.selector, (owner))
                )
            )
        );

        vm.prank(owner);
        b.setup(l1ChainId, address(portal), address(l1Bridge), 0);
        vm.deal(address(b), totalSupply);
    }

    function test_bridge() public {
        address to = makeAddr("to");
        uint256 amount = 1e18;
        uint256 fee = b.bridgeFee(to, amount);

        // to must not be zero
        vm.expectRevert("OmniBridge: no bridge to zero");
        b.bridge(address(0), amount);

        // requires amount > 0
        vm.expectRevert("OmniBridge: amount must be > 0");
        b.bridge(to, 0);

        // requires l1Deposits >= amount
        vm.expectRevert("OmniBridge: no liquidity");
        b.bridge(to, amount);

        b.setL1Deposits(amount - 1);

        // still too low
        vm.expectRevert("OmniBridge: no liquidity");
        b.bridge(to, amount);

        b.setL1Deposits(amount);

        // requires msg.value >= fee + amount
        vm.expectRevert("OmniBridge: insufficient funds");
        b.bridge{ value: amount + fee - 1 }(to, amount);

        // succeeds
        //

        // emits event
        vm.expectEmit();
        emit Bridge(address(this), to, amount);

        // emits xcall
        uint256 feeWithExcess = fee + 1; // test that bridge forwards excess fee to portal
        vm.expectCall(
            address(portal),
            feeWithExcess,
            abi.encodeCall(
                IOmniPortal.xcall,
                (
                    l1ChainId,
                    ConfLevel.Finalized,
                    address(l1Bridge).toBytes32(),
                    abi.encodeCall(OmniBridgeL1.withdraw, (to, amount)),
                    b.XCALL_WITHDRAW_GAS_LIMIT()
                )
            )
        );
        b.bridge{ value: amount + feeWithExcess }(to, amount);

        // decrements l1Deposits
        assertEq(b.l1Deposits(), 0);
        vm.expectRevert("OmniBridge: no liquidity");
        b.bridge(to, amount);
    }

    function test_withdraw() public {
        address payor = makeAddr("payor");
        address to = makeAddr("to");
        uint256 amount = 1e18;
        uint64 gasLimit = l1Bridge.XCALL_WITHDRAW_GAS_LIMIT();

        // sender must be portal
        vm.expectRevert("OmniBridge: not xcall");
        b.withdraw(payor, to, amount);

        // xmsg must be from l1Bridge
        vm.expectRevert("OmniBridge: not bridge");
        portal.mockXCall({
            sourceChainId: l1ChainId,
            sender: address(1234).toBytes32(), // wrong
            to: address(b).toBytes32(),
            data: abi.encodeCall(OmniBridgeNative.withdraw, (payor, to, amount)),
            gasLimit: gasLimit
        });

        // xmsg must be from l1ChainId
        vm.expectRevert("OmniBridge: not L1");
        portal.mockXCall({
            sourceChainId: l1ChainId + 1, // wrong
            sender: address(l1Bridge).toBytes32(),
            to: address(b).toBytes32(),
            data: abi.encodeCall(OmniBridgeNative.withdraw, (payor, to, amount)),
            gasLimit: gasLimit
        });

        // succeeds
        //
        // emits event
        vm.expectEmit();
        emit Withdraw(payor, to, amount, true);

        // transfers amount to to
        vm.expectCall(to, amount, "");
        uint256 gasUsed = portal.mockXCall({
            sourceChainId: l1ChainId,
            sender: address(l1Bridge).toBytes32(),
            to: address(b).toBytes32(),
            data: abi.encodeCall(OmniBridgeNative.withdraw, (payor, to, amount)),
            gasLimit: gasLimit
        });

        // log gas, to inform xcall gas limit
        console.log("OmniBridgeNative.withdraw(success=true) gas used: ", gasUsed);

        assertEq(to.balance, amount);

        // nothing claimable
        assertEq(b.claimable(payor), 0);

        // adds amount to l1Deposits
        assertEq(b.l1Deposits(), amount);

        // adds claimable if to.call fails
        //
        address noReceiver = address(new NoReceive());

        vm.expectEmit();
        emit Withdraw(payor, noReceiver, amount, false);

        vm.expectCall(noReceiver, amount, "");
        gasUsed = portal.mockXCall({
            sourceChainId: l1ChainId,
            sender: address(l1Bridge).toBytes32(),
            to: address(b).toBytes32(),
            data: abi.encodeCall(OmniBridgeNative.withdraw, (payor, noReceiver, amount)),
            gasLimit: gasLimit
        });

        assertEq(b.claimable(payor), amount);

        // log gas, to inform xcall gas limit
        console.log("OmniBridgeNative.withdraw(success=false) gas used: ", gasUsed);
    }

    function test_claim() public {
        address claimant = makeAddr("claimant");
        address to = makeAddr("to");

        // must be xcall
        vm.expectRevert("OmniBridge: not xcall");
        b.claim(address(0));

        // must be from l1
        vm.expectRevert("OmniBridge: not L1");
        portal.mockXCall({
            sourceChainId: l1ChainId + 1, // wrong
            sender: claimant.toBytes32(),
            to: address(b).toBytes32(),
            data: abi.encodeCall(OmniBridgeNative.claim, to),
            gasLimit: 100_000
        });

        // to must not be zero
        vm.expectRevert("OmniBridge: no claim to zero");
        portal.mockXCall({
            sourceChainId: l1ChainId,
            sender: claimant.toBytes32(),
            to: address(b).toBytes32(),
            data: abi.encodeCall(OmniBridgeNative.claim, address(0)),
            gasLimit: 100_000
        });

        // claimant must have claimable
        vm.expectRevert("OmniBridge: nothing to claim");
        portal.mockXCall({
            sourceChainId: l1ChainId,
            sender: claimant.toBytes32(),
            to: address(b).toBytes32(),
            data: abi.encodeCall(OmniBridgeNative.claim, to),
            gasLimit: 100_000
        });

        // reverts on to.call failure
        //
        uint256 amount = 1e18;
        address noReceiver = address(new NoReceive());

        b.setClaimable(claimant, amount);

        vm.expectRevert("OmniBridge: transfer failed");
        portal.mockXCall({
            sourceChainId: l1ChainId,
            sender: claimant.toBytes32(),
            to: address(b).toBytes32(),
            data: abi.encodeCall(OmniBridgeNative.claim, noReceiver),
            gasLimit: 100_000
        });

        // succeeds
        //

        // emits event
        vm.expectEmit();
        emit Claimed(claimant, to, amount);

        // transfers claimable to to
        vm.expectCall(to, amount, "");
        portal.mockXCall({
            sourceChainId: l1ChainId,
            sender: claimant.toBytes32(),
            to: address(b).toBytes32(),
            data: abi.encodeCall(OmniBridgeNative.claim, to),
            gasLimit: 100_000
        });

        // claimable is zero
        assertEq(b.claimable(claimant), 0);

        // to has amount
        assertEq(to.balance, amount);
    }

    function test_pauseBridging() public {
        address to = makeAddr("to");
        uint256 amount = 1e18;
        bytes32 action = b.ACTION_BRIDGE();

        // pause bridging
        vm.prank(owner);
        b.pause(action);

        // assert paused
        assertTrue(b.isPaused(action));

        // bridge reverts
        vm.expectRevert("OmniBridge: paused");
        b.bridge(to, amount);

        // unpause bridging
        vm.prank(owner);
        b.unpause(action);

        // assert unpaused
        assertFalse(b.isPaused(action));

        // no longer paused
        vm.expectRevert("OmniBridge: no liquidity");
        b.bridge(to, amount);
    }

    function test_pauseWithdraws() public {
        address payor = makeAddr("payor");
        address to = makeAddr("to");
        uint256 amount = 1e18;
        bytes32 action = b.ACTION_WITHDRAW();

        // pause withdraws
        vm.prank(owner);
        b.pause(action);

        // assert paused
        assertTrue(b.isPaused(action));

        // withdraw reverts
        vm.expectRevert("OmniBridge: paused");
        b.withdraw(payor, to, amount);

        // claim reverts
        vm.expectRevert("OmniBridge: paused");
        b.claim(to);

        // unpause
        vm.prank(owner);
        b.unpause(action);

        // assert unpaued
        assertFalse(b.isPaused(action));

        // no longer paused
        vm.expectRevert("OmniBridge: not xcall");
        b.withdraw(payor, to, amount);

        vm.expectRevert("OmniBridge: not xcall");
        b.claim(to);
    }

    function test_pauseAll() public {
        address payor = makeAddr("payor");
        address to = makeAddr("to");
        uint256 amount = 1e18;

        // pause all
        vm.prank(owner);
        b.pause();

        // assert actions paus
        assertTrue(b.isPaused(b.ACTION_BRIDGE()));
        assertTrue(b.isPaused(b.ACTION_WITHDRAW()));

        // bridge reverts
        vm.expectRevert("OmniBridge: paused");
        b.bridge(to, amount);

        // withdraw reverts
        vm.expectRevert("OmniBridge: paused");
        b.withdraw(payor, to, amount);

        // claim reverts
        vm.expectRevert("OmniBridge: paused");
        b.claim(to);

        // unpause all
        vm.prank(owner);
        b.unpause();

        assertFalse(b.isPaused(b.ACTION_BRIDGE()));
        assertFalse(b.isPaused(b.ACTION_WITHDRAW()));
    }
}

/**
 * @title OmniBridgeNativeHarness
 * @notice A harness for testing OmniBridgeNative that exposes setup and state modifiers.
 */
contract OmniBridgeNativeHarness is OmniBridgeNative {
    function setL1Deposits(uint256 deposits) public {
        l1Deposits = deposits;
    }

    function setClaimable(address claimant, uint256 amount) public {
        claimable[claimant] = amount;
    }
}
