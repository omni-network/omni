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
    DistributionHarness distribution;

    function setUp() public {
        owner = makeAddr("owner");
        validator = makeAddr("validator");
        validators.push(validator);
        distribution = new DistributionHarness(owner);
    }

    function test_withdraw() public {
        uint256 fee = 0.1 ether;

        vm.expectRevert("Distribution: insufficient fee");
        distribution.withdraw{ value: 0 }(owner);

        vm.deal(owner, fee);
        vm.prank(owner);
        vm.expectEmit();
        emit Withdraw(owner, validator);
        distribution.withdraw{ value: fee }(validator);
    }
}

/**
 * @title DistributionHarness
 * @notice Wrapper around Distribution.sol that allows setting owner in constructor
 */
contract DistributionHarness is Distribution {
    constructor(address _owner) {
        _transferOwnership(_owner);
    }
}
