// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { Test } from "forge-std/Test.sol";
import { VmSafe } from "forge-std/Vm.sol";
import { Distribution } from "src/octane/Distribution.sol";

contract DistributionPostUpgradeTest is Test {
    Distribution distribution;

    function run(address _distribution) public {
        (VmSafe.CallerMode mode,,) = vm.readCallers();
        require(mode == VmSafe.CallerMode.None, "no broadcast");

        _setup(_distribution);
        _testTemporarilyDisabled();
    }

    function _setup(address _distribution) internal {
        distribution = Distribution(_distribution);
    }

    function _testTemporarilyDisabled() internal {
        vm.expectRevert(abi.encodeWithSelector(Distribution.TemporarilyDisabled.selector));
        distribution.withdraw(address(1));
    }
}
