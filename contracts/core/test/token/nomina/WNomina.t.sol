// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { Test } from "forge-std/Test.sol";
import { WNomina } from "src/token/nomina/WNomina.sol";

contract WNomina_Test is Test {
    WNomina public wnomina;

    function setUp() public {
        wnomina = new WNomina();
    }

    receive() external payable { }

    function testMetadata() public view {
        assertEq(wnomina.name(), "Wrapped Nomina", "name mismatch");
        assertEq(wnomina.symbol(), "WNOM", "symbol mismatch");
        assertEq(wnomina.decimals(), 18, "decimals mismatch");
    }

    function testDeposit() public {
        wnomina.deposit{ value: 1 ether }();
        assertEq(wnomina.balanceOf(address(this)), 1 ether);
        assertEq(address(wnomina).balance, 1 ether);

        wnomina.depositTo{ value: 1 ether }(address(this));
        assertEq(wnomina.balanceOf(address(this)), 2 ether);
        assertEq(address(wnomina).balance, 2 ether);

        wnomina.depositTo{ value: 1 ether }(address(0xbeef));
        assertEq(wnomina.balanceOf(address(0xbeef)), 1 ether);
        assertEq(address(wnomina).balance, 3 ether);

        (bool success,) = address(wnomina).call{ value: 1 ether }("");
        assertTrue(success);
        assertEq(wnomina.balanceOf(address(this)), 3 ether);
        assertEq(address(wnomina).balance, 4 ether);

        (success,) = address(wnomina).call{ value: 1 ether }(abi.encodeWithSelector(IFallback.inverseBrother.selector));
        assertTrue(success);
        assertEq(wnomina.balanceOf(address(this)), 4 ether);
        assertEq(address(wnomina).balance, 5 ether);
    }

    function testWithdraw() public {
        testDeposit();

        wnomina.withdraw(2 ether);
        assertEq(wnomina.balanceOf(address(this)), 2 ether);
        assertEq(address(wnomina).balance, 3 ether);

        wnomina.withdrawTo(address(this), 1 ether);
        assertEq(wnomina.balanceOf(address(this)), 1 ether);
        assertEq(address(wnomina).balance, 2 ether);

        wnomina.withdrawTo(address(0xbeef), 1 ether);
        assertEq(wnomina.balanceOf(address(this)), 0);
        assertEq(wnomina.balanceOf(address(0xbeef)), 1 ether);
        assertEq(address(wnomina).balance, 1 ether);
    }
}

interface IFallback {
    function inverseBrother() external view returns (bool);
}
