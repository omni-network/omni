// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { Test } from "forge-std/Test.sol";
import { ERC20 } from "solady/src/tokens/ERC20.sol";
import { Nomina } from "src/token/Nomina.sol";
import { MockOmni } from "test/utils/MockOmni.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";

contract Nomina_Test is Test {
    Nomina public nomina;
    MockOmni public omni;

    address public mintAuthority = makeAddr("mintAuthority");
    address public minter = makeAddr("minter");
    address public user = makeAddr("user");

    function setUp() public {
        omni = new MockOmni(1_000_000 ether, user);
        nomina = new Nomina(address(omni), mintAuthority);
        vm.prank(mintAuthority);
        nomina.setMinter(minter);
    }

    function test_constructor() public {
        vm.expectRevert(Nomina.ZeroAddress.selector);
        new Nomina(address(0), mintAuthority);

        vm.expectRevert(Nomina.ZeroAddress.selector);
        new Nomina(address(omni), address(0));
    }

    function testMetadata() public view {
        assertEq(nomina.name(), "Nomina", "name mismatch");
        assertEq(nomina.symbol(), "NOM", "symbol mismatch");
        assertEq(nomina.decimals(), 18, "decimals mismatch");
        assertEq(nomina.totalSupply(), 0, "total supply mismatch");
        assertEq(nomina.OMNI(), address(omni), "omni mismatch");
        assertEq(nomina.mintAuthority(), mintAuthority, "mint authority mismatch");
        assertEq(nomina.minter(), minter, "minter mismatch");
        assertEq(nomina.CONVERSION_RATE(), 75, "conversion rate mismatch");
    }

    function testMintReverts() public {
        vm.expectRevert(Nomina.Unauthorized.selector);
        vm.prank(user);
        nomina.mint(user, 1 ether);
    }

    function testMint() public {
        vm.prank(minter);
        nomina.mint(user, 1 ether);

        assertEq(nomina.balanceOf(user), 1 ether, "balance mismatch");
        assertEq(nomina.totalSupply(), 1 ether, "total supply mismatch");
    }

    function testBurnReverts() public {
        // Must have balance to burn
        vm.expectRevert(ERC20.InsufficientBalance.selector);
        nomina.burn(1 ether);
    }

    function testBurn() public {
        vm.startPrank(user);
        omni.approve(address(nomina), 1 ether);
        nomina.convert(user, 1 ether);
        nomina.burn(1 ether);
        vm.stopPrank();

        // Ensure no conversion ratio is applied to burns
        assertEq(nomina.balanceOf(user), 74 ether, "balance mismatch");
        assertEq(nomina.totalSupply(), 74 ether, "total supply mismatch");
    }

    function testConvertReverts() public {
        // Token approval is required
        vm.expectRevert(SafeTransferLib.TransferFromFailed.selector);
        vm.prank(user);
        nomina.convert(user, 1 ether);

        // Token balance is required
        vm.startPrank(minter);
        omni.approve(address(nomina), type(uint256).max);
        vm.expectRevert(SafeTransferLib.TransferFromFailed.selector);
        nomina.convert(user, 1 ether);
        vm.stopPrank();

        // Burn on conversion is disallowed
        vm.expectRevert(Nomina.ZeroAddress.selector);
        vm.prank(user);
        nomina.convert(address(0), 1 ether);
    }

    function testConvert() public {
        vm.startPrank(user);
        omni.approve(address(nomina), 1 ether);
        nomina.convert(user, 1 ether);
        vm.stopPrank();

        // Ensure 1:75 ratio is preserved
        uint256 conversionRate = nomina.CONVERSION_RATE();
        assertEq(nomina.balanceOf(user), 1 ether * conversionRate, "balance mismatch");
        assertEq(nomina.totalSupply(), 1 ether * conversionRate, "total supply mismatch");
        assertEq(1 ether * conversionRate, 75 ether, "conversion rate mismatch");
    }

    function testSetMintAuthorityReverts() public {
        vm.expectRevert(Nomina.Unauthorized.selector);
        vm.prank(user);
        nomina.setMintAuthority(user);
    }

    function testSetMintAuthority() public {
        // Trigger pending transfer of mint authority
        vm.prank(mintAuthority);
        nomina.setMintAuthority(user);
        assertEq(nomina.pendingMintAuthority(), user, "pending mint authority mismatch");

        // Cancel pending transfer by setting the zero address
        vm.prank(mintAuthority);
        nomina.setMintAuthority(address(0));
        assertEq(nomina.pendingMintAuthority(), address(0), "pending mint authority mismatch");
    }

    function testAcceptMintAuthorityReverts() public {
        vm.expectRevert(Nomina.Unauthorized.selector);
        vm.prank(user);
        nomina.acceptMintAuthority();
    }

    function testAcceptMintAuthority() public {
        vm.prank(mintAuthority);
        nomina.setMintAuthority(user);

        vm.prank(user);
        nomina.acceptMintAuthority();

        assertEq(nomina.mintAuthority(), user, "mint authority mismatch");
    }

    function testSetMinterReverts() public {
        vm.expectRevert(Nomina.Unauthorized.selector);
        vm.prank(user);
        nomina.setMinter(user);
    }

    function testSetMinter() public {
        vm.prank(mintAuthority);
        nomina.setMinter(user);
        assertEq(nomina.minter(), user, "minter mismatch");
    }
}
