// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { Distribution } from "src/octane/Distribution.sol";
import { Test, Vm } from "forge-std/Test.sol";
import { Secp256k1 } from "src/libraries/Secp256k1.sol";

/**
 * @title Distribution_Test
 * @notice Test suite for Distribution.sol
 */
contract Distribution_Test is Test {
    /// @dev Matches Distribution.Withdraw event
    event Withdraw(address indexed delegator, address indexed validator);

    address owner;
    address validator;
    address[] validators;
    Distribution distribution;

    function setUp() public {
        owner = makeAddr("owner");
        validator = makeAddr("validator");
        validators.push(validator);
        distribution = new Distribution();
    }

    /*function test_withdraw() public {
        uint256 fee = 0.1 ether;

        vm.expectRevert("Distribution: insufficient fee");
        distribution.withdraw{ value: 0 }(owner);

        vm.deal(owner, fee);
        vm.prank(owner);
        vm.expectEmit();
        emit Withdraw(owner, validator);
        distribution.withdraw{ value: fee }(validator);
    }*/

    function test_temporarilyDisabled() public {
        vm.expectRevert(abi.encodeWithSelector(Distribution.TemporarilyDisabled.selector));
        distribution.withdraw(validator);
    }
}
