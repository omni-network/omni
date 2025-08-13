// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { MockPortal } from "test/utils/nomina/MockPortal.sol";
import { INominaPortal } from "src/interfaces/nomina/INominaPortal.sol";
import { Predeploys } from "src/libraries/nomina/Predeploys.sol";
import { MockOmni } from "nomina/test/utils/MockOmni.sol";
import { Nomina } from "nomina/src/token/Nomina.sol";
import { NominaBridgeNative } from "src/token/nomina/NominaBridgeNative.sol";
import { NominaBridgeL1 } from "src/token/nomina/NominaBridgeL1.sol";
import { ConfLevel } from "src/libraries/ConfLevel.sol";
import { Test } from "forge-std/Test.sol";
import { console } from "forge-std/console.sol";
import { ERC20 } from "solady/src/tokens/ERC20.sol";

/**
 * @title NominaBridgeL1_Test
 * @notice Test suite for NominaBridgeNative contract.
 */
contract NominaBridgeL1_Test is Test {
    // events copied from NominaBridgeL1.sol
    event Bridge(address indexed payor, address indexed to, uint256 amount);
    event Withdraw(address indexed to, uint256 amount);

    MockPortal portal;
    MockOmni omni;
    Nomina nomina;
    NominaBridgeL1Harness b;

    address owner;
    address minter;
    address mintAuthority;
    address proxyAdmin;
    address initialSupplyRecipient;
    uint8 conversionRate = 75;
    uint256 amount = 1 ether;
    uint256 totalSupply = 100_000_000 ether;

    function setUp() public {
        initialSupplyRecipient = makeAddr("initialSupplyRecipient");
        owner = makeAddr("owner");
        proxyAdmin = makeAddr("proxyAdmin");
        minter = makeAddr("minter");
        mintAuthority = makeAddr("mintAuthority");

        portal = new MockPortal();
        omni = new MockOmni(totalSupply, initialSupplyRecipient);

        nomina = new Nomina(address(omni), mintAuthority);
        vm.prank(mintAuthority);
        nomina.setMinter(minter);

        address impl = address(new NominaBridgeL1Harness(address(omni), address(nomina)));
        b = NominaBridgeL1Harness(
            address(
                new TransparentUpgradeableProxy(
                    impl, proxyAdmin, abi.encodeCall(NominaBridgeL1.initialize, (owner, address(portal)))
                )
            )
        );
        b.initializeV2();
    }

    function test_initialize() public {
        address impl = address(new NominaBridgeL1(address(omni), address(nomina)));
        address proxy = address(new TransparentUpgradeableProxy(impl, proxyAdmin, ""));

        // reverts
        vm.expectRevert("NominaBridge: no zero addr");
        NominaBridgeL1(proxy).initialize(owner, address(0));

        // succeeds
        NominaBridgeL1(proxy).initialize(owner, address(portal));

        // initializeV2 converts omni balance to nomina
        vm.prank(initialSupplyRecipient);
        omni.transfer(proxy, amount);
        NominaBridgeL1(proxy).initializeV2();

        // assert balance
        assertEq(nomina.balanceOf(proxy), amount * conversionRate);
        assertEq(omni.balanceOf(proxy), 0);
    }

    function test_bridge() public {
        address to = makeAddr("to");
        address payor = address(this);
        uint256 fee = b.bridgeFee(payor, to, amount);

        // requires amount > 0
        vm.expectRevert("NominaBridge: amount must be > 0");
        b.bridge(to, 0);

        // to must not be zero
        vm.expectRevert("NominaBridge: no bridge to zero");
        b.bridge(address(0), amount);

        // value must be greater than or equal fee
        vm.expectRevert("NominaBridge: insufficient fee");
        b.bridge{ value: fee - 1 }(to, amount);

        // requires allowance
        vm.expectRevert(ERC20.InsufficientAllowance.selector);
        b.bridge{ value: fee }(to, amount);

        nomina.approve(address(b), amount);

        // requires balance
        vm.expectRevert(ERC20.InsufficientBalance.selector);
        b.bridge{ value: fee }(to, amount);

        // succeeds
        //
        // fund payor
        vm.startPrank(initialSupplyRecipient);
        omni.approve(address(nomina), amount);
        nomina.convert(payor, amount);
        vm.stopPrank();

        // emits event
        vm.expectEmit();
        emit Bridge(payor, to, amount);

        // emits xcall
        vm.expectCall(
            address(portal),
            fee,
            abi.encodeCall(
                INominaPortal.xcall,
                (
                    portal.nominaChainId(),
                    ConfLevel.Finalized,
                    Predeploys.NominaBridgeNative,
                    abi.encodeCall(NominaBridgeNative.withdraw, (payor, to, amount)),
                    b.XCALL_WITHDRAW_GAS_LIMIT()
                )
            )
        );
        b.bridge{ value: fee }(to, amount);

        // assert balance change
        assertEq(nomina.balanceOf(address(b)), amount);
        assertEq(nomina.balanceOf(payor), (amount * conversionRate) - amount);
    }

    function test_withdraw() public {
        address to = makeAddr("to");
        uint64 nominaChainId = portal.nominaChainId();
        uint64 gasLimit = new NominaBridgeNative().XCALL_WITHDRAW_GAS_LIMIT();

        // sender must be portal
        vm.expectRevert("NominaBridge: not xcall");
        b.withdraw(to, amount);

        // xmsg must be from native bridge
        vm.expectRevert("NominaBridge: not bridge");
        portal.mockXCall({
            sourceChainId: nominaChainId,
            sender: address(1234), // wrong
            to: address(b),
            data: abi.encodeCall(NominaBridgeL1.withdraw, (to, amount)),
            gasLimit: gasLimit
        });

        // xmsg must be from nomina evm
        vm.expectRevert("NominaBridge: not nomina portal");
        portal.mockXCall({
            sourceChainId: nominaChainId + 1, // wrong
            sender: Predeploys.NominaBridgeNative,
            to: address(b),
            data: abi.encodeCall(NominaBridgeL1.withdraw, (to, amount)),
            gasLimit: gasLimit
        });

        // succeeds
        //
        // need to fund bridge first
        vm.startPrank(initialSupplyRecipient);
        omni.approve(address(nomina), amount);
        nomina.convert(address(b), amount);
        vm.stopPrank();

        // emit event
        vm.expectEmit();
        emit Withdraw(to, amount);

        // tranfers amount to to
        vm.expectCall(address(nomina), abi.encodeCall(nomina.transfer, (to, amount)));
        uint256 gasUsed = portal.mockXCall({
            sourceChainId: portal.nominaChainId(),
            sender: Predeploys.NominaBridgeNative,
            to: address(b),
            data: abi.encodeCall(NominaBridgeL1.withdraw, (to, amount)),
            gasLimit: gasLimit
        });

        // assert balance change
        assertEq(nomina.balanceOf(to), amount);
        assertEq(nomina.balanceOf(address(b)), (amount * conversionRate) - amount);

        // log gas, to inform xcall gas limit
        console.log("NominaBridgeL1.withdraw gas used: ", gasUsed);
    }

    function test_pauseBridging() public {
        address to = makeAddr("to");
        bytes32 action = b.ACTION_BRIDGE();

        // pause bridging
        vm.prank(owner);
        b.pause(action);

        // assert paused
        assertTrue(b.isPaused(action));

        // bridge reverts
        vm.expectRevert("NominaBridge: paused");
        b.bridge(to, amount);

        // unpause bridging
        vm.prank(owner);
        b.unpause(action);

        // assert unpaused
        assertFalse(b.isPaused(action));

        // bridge not paused (reverts, but not due to pause)
        vm.expectRevert("NominaBridge: insufficient fee");
        b.bridge(to, amount);
    }

    function test_pauseWithdraws() public {
        address to = makeAddr("to");
        bytes32 action = b.ACTION_WITHDRAW();

        // pause withdraws
        vm.prank(owner);
        b.pause(action);

        // assert paused
        assertTrue(b.isPaused(action));

        // withdraw reverts
        vm.expectRevert("NominaBridge: paused");
        b.withdraw(to, amount);

        // unpause
        vm.prank(owner);
        b.unpause(action);

        // assert unpaued
        assertFalse(b.isPaused(action));

        // no longer paused
        vm.expectRevert("NominaBridge: not xcall");
        b.withdraw(to, amount);
    }

    function test_pauseAll() public {
        address to = makeAddr("to");

        // pause all
        vm.prank(owner);
        b.pause();

        // assert actions paus
        assertTrue(b.isPaused(b.ACTION_BRIDGE()));
        assertTrue(b.isPaused(b.ACTION_WITHDRAW()));

        // bridge reverts
        vm.expectRevert("NominaBridge: paused");
        b.bridge(to, amount);

        // withdraw reverts
        vm.expectRevert("NominaBridge: paused");
        b.withdraw(to, amount);

        // unpause all
        vm.prank(owner);
        b.unpause();

        assertFalse(b.isPaused(b.ACTION_BRIDGE()));
        assertFalse(b.isPaused(b.ACTION_WITHDRAW()));
    }
}

contract NominaBridgeL1Harness is NominaBridgeL1 {
    constructor(address omni, address nomina) NominaBridgeL1(omni, nomina) { }
}
