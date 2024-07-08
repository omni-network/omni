// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { MockPortal } from "test/utils/MockPortal.sol";
import { IOmniPortal } from "src/interfaces/IOmniPortal.sol";
import { OmniBridgeNative } from "src/token/OmniBridgeNative.sol";
import { OmniBridgeL1 } from "src/token/OmniBridgeL1.sol";
import { ConfLevel } from "src/libraries/ConfLevel.sol";
import { Test } from "forge-std/Test.sol";
import { console } from "forge-std/console.sol";

/**
 * @title OmniBridgeNative_Test
 * @notice Test suite for OmniBridgeNative contract.
 */
contract OmniBridgeNative_Test is Test {
    // Events copied from OmniBridgeNative.sol
    event Bridge(address indexed payor, address indexed to, uint256 amount);
    event Withdraw(address indexed payor, address indexed to, uint256 amount, bool success);
    event Claimed(address indexed claimant, address indexed to, uint256 amount);

    MockPortal portal;
    OmniBridgeNativeHarness b;
    OmniBridgeL1 l1Bridge;

    uint64 l1ChainId;

    uint256 totalSupply = 100_000_000 * 10 ** 18;

    function setUp() public {
        portal = new MockPortal();
        b = new OmniBridgeNativeHarness();
        l1ChainId = 1;
        l1Bridge = new OmniBridgeL1(makeAddr("token"));
        b.setupNoAuth(l1ChainId, address(portal), address(l1Bridge));
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

        // requires l1BridgeBalance >= amount
        vm.expectRevert("OmniBridge: no liquidity");
        b.bridge(to, amount);

        b.setL1BridgeBalance(amount - 1);

        // still too low
        vm.expectRevert("OmniBridge: no liquidity");
        b.bridge(to, amount);

        b.setL1BridgeBalance(amount);

        // requires msg.value == fee + amount
        vm.expectRevert("OmniBridge: incorrect funds");
        b.bridge{ value: amount }(to, amount);
        vm.expectRevert("OmniBridge: incorrect funds");
        b.bridge{ value: amount + 1 }(to, amount);

        // must equal amount + fee
        vm.expectRevert("OmniBridge: incorrect funds");
        b.bridge{ value: amount + fee + 1 }(to, amount);

        // succeeds
        //

        // emits event
        vm.expectEmit();
        emit Bridge(address(this), to, amount);

        // emits xcall
        vm.expectCall(
            address(portal),
            fee,
            abi.encodeWithSelector(
                IOmniPortal.xcall.selector,
                l1ChainId,
                ConfLevel.Finalized,
                l1Bridge,
                abi.encodeWithSelector(OmniBridgeL1.withdraw.selector, to, amount),
                b.XCALL_WITHDRAW_GAS_LIMIT()
            )
        );
        b.bridge{ value: amount + fee }(to, amount);

        // decrements l1BridgeBalance
        assertEq(b.l1BridgeBalance(), 0);
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
            sender: address(1234), // wrong
            to: address(b),
            data: abi.encodeWithSelector(OmniBridgeNative.withdraw.selector, payor, to, amount),
            gasLimit: gasLimit
        });

        // xmsg must be from l1ChainId
        vm.expectRevert("OmniBridge: not L1");
        portal.mockXCall({
            sourceChainId: l1ChainId + 1, // wrong
            sender: address(l1Bridge),
            to: address(b),
            data: abi.encodeWithSelector(OmniBridgeNative.withdraw.selector, payor, to, amount),
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
            sender: address(l1Bridge),
            to: address(b),
            data: abi.encodeWithSelector(OmniBridgeNative.withdraw.selector, payor, to, amount),
            gasLimit: gasLimit
        });

        // log gas, to inform xcall gas limit
        console.log("OmniBridgeNative.withdraw(success=true) gas used: ", gasUsed);

        assertEq(to.balance, amount);

        // nothing claimable
        assertEq(b.claimable(payor), 0);

        // increments l1BridgeBalance
        assertEq(b.l1BridgeBalance(), amount);

        // adds claimable if to.call fails
        //
        address noReceiver = address(new NoReceive());

        vm.expectEmit();
        emit Withdraw(payor, noReceiver, amount, false);

        vm.expectCall(noReceiver, amount, "");
        gasUsed = portal.mockXCall({
            sourceChainId: l1ChainId,
            sender: address(l1Bridge),
            to: address(b),
            data: abi.encodeWithSelector(OmniBridgeNative.withdraw.selector, payor, noReceiver, amount),
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
            sender: claimant,
            to: address(b),
            data: abi.encodeWithSelector(OmniBridgeNative.claim.selector, to),
            gasLimit: 100_000
        });

        // to must not be zero
        vm.expectRevert("OmniBridge: no claim to zero");
        portal.mockXCall({
            sourceChainId: l1ChainId,
            sender: claimant,
            to: address(b),
            data: abi.encodeWithSelector(OmniBridgeNative.claim.selector, address(0)),
            gasLimit: 100_000
        });

        // claimant must have claimable
        vm.expectRevert("OmniBridge: nothing to claim");
        portal.mockXCall({
            sourceChainId: l1ChainId,
            sender: claimant,
            to: address(b),
            data: abi.encodeWithSelector(OmniBridgeNative.claim.selector, to),
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
            sender: claimant,
            to: address(b),
            data: abi.encodeWithSelector(OmniBridgeNative.claim.selector, noReceiver),
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
            sender: claimant,
            to: address(b),
            data: abi.encodeWithSelector(OmniBridgeNative.claim.selector, to),
            gasLimit: 100_000
        });

        // claimable is zero
        assertEq(b.claimable(claimant), 0);

        // to has amount
        assertEq(to.balance, amount);
    }
}

/**
 * @title OmniBridgeNativeHarness
 * @notice A harness for testing OmniBridgeNative that exposes setup and state modifiers.
 */
contract OmniBridgeNativeHarness is OmniBridgeNative {
    function setupNoAuth(uint64 l1ChainId_, address omni_, address l1Bridge_) public {
        l1ChainId = l1ChainId_;
        omni = IOmniPortal(omni_);
        l1Bridge = l1Bridge_;
    }

    function setL1BridgeBalance(uint256 balance) public {
        l1BridgeBalance = balance;
    }

    function setClaimable(address claimant, uint256 amount) public {
        claimable[claimant] = amount;
    }
}

/**
 * @title NoReceive
 * @notice An contract that does not implement the receive function.
 */
contract NoReceive { }
