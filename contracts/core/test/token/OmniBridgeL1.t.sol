// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { MockPortal } from "test/utils/MockPortal.sol";
import { IOmniPortal } from "src/interfaces/IOmniPortal.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";
import { Omni } from "src/token/Omni.sol";
import { OmniBridgeNative } from "src/token/OmniBridgeNative.sol";
import { OmniBridgeL1 } from "src/token/OmniBridgeL1.sol";
import { ConfLevel } from "src/libraries/ConfLevel.sol";
import { Test } from "forge-std/Test.sol";
import { console } from "forge-std/console.sol";

/**
 * @title OmniBridgeL1_Test
 * @notice Test suite for OmniBridgeNative contract.
 */
contract OmniBridgeL1_Test is Test {
    // events copied from OmniBridgeL1.sol
    event Bridge(address indexed payor, address indexed to, uint256 amount);
    event Withdraw(address indexed to, uint256 amount);

    MockPortal portal;
    Omni omni;
    OmniBridgeL1Harness b;

    address owner;
    address proxyAdmin;
    address initialSupplyRecipient;
    uint256 totalSupply = 100_000_000 * 10 ** 18;

    function setUp() public {
        initialSupplyRecipient = makeAddr("initialSupplyRecipient");
        owner = makeAddr("owner");
        proxyAdmin = makeAddr("proxyAdmin");

        portal = new MockPortal();
        omni = new Omni(totalSupply, initialSupplyRecipient);

        address impl = address(new OmniBridgeL1Harness(address(omni)));
        b = OmniBridgeL1Harness(
            address(
                new TransparentUpgradeableProxy(
                    impl, proxyAdmin, abi.encodeCall(OmniBridgeL1.initialize, (owner, address(portal)))
                )
            )
        );
    }

    function test_initialize() public {
        address impl = address(new OmniBridgeL1(address(omni)));
        address proxy = address(new TransparentUpgradeableProxy(impl, proxyAdmin, ""));

        vm.expectRevert("OmniBridge: no zero addr");
        OmniBridgeL1(proxy).initialize(owner, address(0));
    }

    function test_bridge() public {
        address to = makeAddr("to");
        uint256 amount = 1e18;
        address payor = address(this);
        uint256 fee = b.bridgeFee(payor, to, amount);

        // requires amount > 0
        vm.expectRevert("OmniBridge: amount must be > 0");
        b.bridge(to, 0);

        // to must not be zero
        vm.expectRevert("OmniBridge: no bridge to zero");
        b.bridge(address(0), amount);

        // value must be greater than or equal fee
        vm.expectRevert("OmniBridge: insufficient fee");
        b.bridge{ value: fee - 1 }(to, amount);

        // requires allowance
        vm.expectRevert("ERC20: insufficient allowance");
        b.bridge{ value: fee }(to, amount);

        omni.approve(address(b), amount);

        // requires balance
        vm.expectRevert("ERC20: transfer amount exceeds balance");
        b.bridge{ value: fee }(to, amount);

        // succeeds
        //
        // fund payor
        vm.prank(initialSupplyRecipient);
        omni.transfer(payor, amount);

        // emits event
        vm.expectEmit();
        emit Bridge(payor, to, amount);

        // emits xcall
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
        b.bridge{ value: fee }(to, amount);

        // assert balance change
        assertEq(omni.balanceOf(address(b)), amount);
        assertEq(omni.balanceOf(payor), 0);
    }

    function test_withdraw() public {
        address to = makeAddr("to");
        uint256 amount = 1e18;
        uint64 omniChainId = portal.omniChainId();
        uint64 gasLimit = new OmniBridgeNative().XCALL_WITHDRAW_GAS_LIMIT();

        // sender must be portal
        vm.expectRevert("OmniBridge: not xcall");
        b.withdraw(to, amount);

        // xmsg must be from native bridge
        vm.expectRevert("OmniBridge: not bridge");
        portal.mockXCall({
            sourceChainId: omniChainId,
            sender: address(1234), // wrong
            to: address(b),
            data: abi.encodeCall(OmniBridgeL1.withdraw, (to, amount)),
            gasLimit: gasLimit
        });

        // xmsg must be from omni evm
        vm.expectRevert("OmniBridge: not omni portal");
        portal.mockXCall({
            sourceChainId: omniChainId + 1, // wrong
            sender: Predeploys.OmniBridgeNative,
            to: address(b),
            data: abi.encodeCall(OmniBridgeL1.withdraw, (to, amount)),
            gasLimit: gasLimit
        });

        // succeeds
        //
        // need to fund bridge first
        vm.prank(initialSupplyRecipient);
        omni.transfer(address(b), amount);

        // emit event
        vm.expectEmit();
        emit Withdraw(to, amount);

        // tranfers amount to to
        vm.expectCall(address(omni), abi.encodeCall(omni.transfer, (to, amount)));
        uint256 gasUsed = portal.mockXCall({
            sourceChainId: portal.omniChainId(),
            sender: Predeploys.OmniBridgeNative,
            to: address(b),
            data: abi.encodeCall(OmniBridgeL1.withdraw, (to, amount)),
            gasLimit: gasLimit
        });

        // assert balance change
        assertEq(omni.balanceOf(to), amount);
        assertEq(omni.balanceOf(address(b)), 0);

        // log gas, to inform xcall gas limit
        console.log("OmniBridgeL1.withdraw gas used: ", gasUsed);
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

        // bridge not paused (reverts, but not due to pause)
        vm.expectRevert("OmniBridge: insufficient fee");
        b.bridge(to, amount);
    }

    function test_pauseWithdraws() public {
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
        b.withdraw(to, amount);

        // unpause
        vm.prank(owner);
        b.unpause(action);

        // assert unpaued
        assertFalse(b.isPaused(action));

        // no longer paused
        vm.expectRevert("OmniBridge: not xcall");
        b.withdraw(to, amount);
    }

    function test_pauseAll() public {
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
        b.withdraw(to, amount);

        // unpause all
        vm.prank(owner);
        b.unpause();

        assertFalse(b.isPaused(b.ACTION_BRIDGE()));
        assertFalse(b.isPaused(b.ACTION_WITHDRAW()));
    }
}

contract OmniBridgeL1Harness is OmniBridgeL1 {
    constructor(address omni) OmniBridgeL1(omni) { }
}
