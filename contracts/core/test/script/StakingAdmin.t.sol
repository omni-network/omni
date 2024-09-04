// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { InitializableHelper } from "script/utils/InitializableHelper.sol";
import { EIP1967Helper } from "script/utils/EIP1967Helper.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";
import { StakingAdmin } from "script/admin/StakingAdmin.s.sol";
import { Staking } from "src/octane/Staking.sol";
import { Test } from "forge-std/Test.sol";

contract StakingAdmin_Test is Test {
    address admin = makeAddr("admin");
    bool enableAllowList = true;

    // @dev Etchs a TransparentUpgradeableProxy at Predeploys.Staking,
    //      with a Staking implementation
    function setUp() public {
        address impl = address(new Staking());

        // tmp proxy, code etched to predeploy address
        address tmp = address(new TransparentUpgradeableProxy(impl, admin, ""));
        vm.etch(Predeploys.Staking, tmp.code);

        EIP1967Helper.setImplementation(Predeploys.Staking, impl);
        EIP1967Helper.setAdmin(Predeploys.Staking, EIP1967Helper.getAdmin(tmp));

        vm.etch(tmp, "");
        vm.resetNonce(tmp);

        Staking(Predeploys.Staking).initialize(admin, enableAllowList);
    }

    function test_setAllowlist() public {
        address[] memory allowlist = new address[](2);
        allowlist[0] = makeAddr("val1");
        allowlist[1] = makeAddr("val2");

        Staking staking = Staking(Predeploys.Staking);
        StakingAdmin a = new StakingAdmin();
        a.setAllowlist(admin, allowlist);

        assertEq(staking.allowlist(), allowlist);
        assertEq(staking.isAllowedValidator(allowlist[0]), true);
        assertEq(staking.isAllowedValidator(allowlist[1]), true);
    }

    function test_setAllowlistEnabled() public {
        Staking staking = Staking(Predeploys.Staking);

        StakingAdmin a = new StakingAdmin();
        a.setAllowlistEnabled(admin, !enableAllowList);

        assertEq(staking.isAllowlistEnabled(), !enableAllowList);
    }

    function test_upgrade() public {
        StakingAdmin a = new StakingAdmin();

        address deployer = makeAddr("deployer");
        vm.deal(deployer, 1 ether);

        // just assert no errors, this runs full Staking test suite
        a.upgrade(admin, deployer);
    }
}
