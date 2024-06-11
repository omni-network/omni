// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { Omni } from "src/token/Omni.sol";
import { Test } from "forge-std/Test.sol";

contract Omni_Test is Test {
    function test_constructor() public {
        uint256 initialSupply = 1e24;
        address recipient = makeAddr("receipient");
        Omni token = new Omni(initialSupply, recipient);

        assertEq(token.totalSupply(), initialSupply);
        assertEq(token.balanceOf(recipient), initialSupply);
        assertEq(token.symbol(), "OMNI");
        assertEq(token.name(), "Omni Network");
    }
}
