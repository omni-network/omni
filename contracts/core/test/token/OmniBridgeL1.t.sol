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
    Omni token;
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
        token = new Omni(totalSupply, initialSupplyRecipient);

        address impl = address(new OmniBridgeL1Harness(address(token)));
        b = OmniBridgeL1Harness(
            address(
                new TransparentUpgradeableProxy(
                    impl, proxyAdmin, abi.encodeWithSelector(OmniBridgeL1.initialize.selector, owner, address(portal))
                )
            )
        );
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

        // value must equal fee
        vm.expectRevert("OmniBridge: incorrect fee");
        b.bridge{ value: fee + 1 }(to, amount);
        vm.expectRevert("OmniBridge: incorrect fee");
        b.bridge{ value: fee - 1 }(to, amount);

        // requires allowance
        vm.expectRevert("ERC20: insufficient allowance");
        b.bridge{ value: fee }(to, amount);

        token.approve(address(b), amount);

        // requires balance
        vm.expectRevert("ERC20: transfer amount exceeds balance");
        b.bridge{ value: fee }(to, amount);

        // succeeds
        //
        // fund payor
        vm.prank(initialSupplyRecipient);
        token.transfer(payor, amount);

        // emits event
        vm.expectEmit();
        emit Bridge(payor, to, amount);

        // emits xcall
        vm.expectCall(
            address(portal),
            fee,
            abi.encodeWithSelector(
                IOmniPortal.xcall.selector,
                portal.omniChainId(),
                ConfLevel.Finalized,
                Predeploys.OmniBridgeNative,
                abi.encodeWithSelector(OmniBridgeNative.withdraw.selector, payor, to, amount),
                b.XCALL_WITHDRAW_GAS_LIMIT()
            )
        );
        b.bridge{ value: fee }(to, amount);

        // assert balance change
        assertEq(token.balanceOf(address(b)), amount);
        assertEq(token.balanceOf(payor), 0);
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
            data: abi.encodeWithSelector(OmniBridgeL1.withdraw.selector, to, amount),
            gasLimit: gasLimit
        });

        // xmsg must be from omni evm
        vm.expectRevert("OmniBridge: not omni");
        portal.mockXCall({
            sourceChainId: omniChainId + 1, // wrong
            sender: Predeploys.OmniBridgeNative,
            to: address(b),
            data: abi.encodeWithSelector(OmniBridgeL1.withdraw.selector, to, amount),
            gasLimit: gasLimit
        });

        // succeeds
        //
        // need to fund bridge first
        vm.prank(initialSupplyRecipient);
        token.transfer(address(b), amount);

        // emit event
        vm.expectEmit();
        emit Withdraw(to, amount);

        // tranfers amount to to
        vm.expectCall(address(token), abi.encodeWithSelector(token.transfer.selector, to, amount));
        uint256 gasUsed = portal.mockXCall({
            sourceChainId: portal.omniChainId(),
            sender: Predeploys.OmniBridgeNative,
            to: address(b),
            data: abi.encodeWithSelector(OmniBridgeL1.withdraw.selector, to, amount),
            gasLimit: gasLimit
        });

        // assert balance change
        assertEq(token.balanceOf(to), amount);
        assertEq(token.balanceOf(address(b)), 0);

        // log gas, to inform xcall gas limit
        console.log("OmniBridgeL1.withdraw gas used: ", gasUsed);
    }
}

contract OmniBridgeL1Harness is OmniBridgeL1 {
    constructor(address token) OmniBridgeL1(token) { }
}
