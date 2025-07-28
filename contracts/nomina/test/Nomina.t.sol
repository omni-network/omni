// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.30;

import { Test } from "forge-std/Test.sol";
import { Nomina } from "../src/Nomina.sol";
import { MockOmni } from "./utils/MockOmni.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";

contract Nomina_Test is Test {
    Nomina public nomina;
    MockOmni public omni;

    address public minter = makeAddr("minter");
    address public user = makeAddr("user");

    function setUp() public {
        omni = new MockOmni(1_000_000 ether, minter);
        nomina = new Nomina(address(omni));
    }

    function testMetadata() public view {
        assertEq(nomina.name(), "Nomina");
        assertEq(nomina.symbol(), "NOM");
        assertEq(nomina.decimals(), 18);
        assertEq(nomina.totalSupply(), 0);
        assertEq(nomina.omni(), address(omni));
    }

    function testConvertReverts() public {
        // Token approval is required
        vm.expectRevert(SafeTransferLib.TransferFromFailed.selector);
        vm.prank(minter);
        nomina.convert(user, 1 ether);

        // Token balance is required
        vm.startPrank(user);
        omni.approve(address(nomina), type(uint256).max);
        vm.expectRevert(SafeTransferLib.TransferFromFailed.selector);
        nomina.convert(user, 1 ether);
        vm.stopPrank();

        // Burn on conversion is disallowed
        vm.expectRevert(Nomina.ZeroAddress.selector);
        vm.prank(minter);
        nomina.convert(address(0), 1 ether);

        // Zero amount is not allowed
        vm.expectRevert(Nomina.ZeroAmount.selector);
        nomina.convert(user, 0);

        // Conversion must not be disabled
        nomina = new Nomina(address(0));
        vm.expectRevert(Nomina.ConversionDisabled.selector);
        nomina.convert(user, 1 ether);
    }

    function testConvert() public {
        vm.startPrank(minter);
        omni.approve(address(nomina), 1 ether);
        nomina.convert(minter, 1 ether);
        vm.stopPrank();

        // Ensure 1:75 ratio is preserved
        assertEq(nomina.balanceOf(minter), 75 ether);
        assertEq(nomina.totalSupply(), 75 ether);
    }

    function testBurnReverts() public {
        // Zero amount is not allowed
        vm.expectRevert(Nomina.ZeroAmount.selector);
        nomina.burn(0);
    }

    function testBurn() public {
        vm.startPrank(minter);
        omni.approve(address(nomina), 1 ether);
        nomina.convert(minter, 1 ether);
        nomina.burn(1 ether);
        vm.stopPrank();

        // Ensure no conversion ratio is applied to burns
        assertEq(nomina.balanceOf(minter), 74 ether);
        assertEq(nomina.totalSupply(), 74 ether);
    }
}
