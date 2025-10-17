// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { GenesisStakeV2 } from "src/token/GenesisStakeV2.sol";
import { ProxyAdmin } from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {
    ITransparentUpgradeableProxy
} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { EIP1967Helper } from "script/utils/EIP1967Helper.sol";
import { Test } from "forge-std/Test.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { VmSafe } from "forge-std/Vm.sol";

contract GenesisStakeV2UpgradeTest is Test {
    using SafeTransferLib for address;

    GenesisStakeV2 internal genesisStake = GenesisStakeV2(0xD2639676dA3dEA5491d27DA19340556b3a7d58B8);

    ProxyAdmin internal proxyAdmin;
    address internal admin;
    address internal owner;
    address internal token;
    uint256 internal userBalance;

    address internal rewardsDistributor = makeAddr("rewardsDistributor");
    address internal user = makeAddr("user");

    function run() public {
        (VmSafe.CallerMode mode,,) = vm.readCallers();
        require(mode == VmSafe.CallerMode.None, "no broadcast");

        _setup();
        _cacheState();
        _upgrade();
        _checkState();
        _testMigrateStake();
    }

    // Retrieve values relevant for testing
    function _setup() internal {
        proxyAdmin = ProxyAdmin(EIP1967Helper.getAdmin(address(genesisStake)));
        admin = proxyAdmin.owner();
        owner = genesisStake.owner();

        token = address(genesisStake.token());
    }

    // Cache values relevant for testing post-upgrade
    function _cacheState() internal {
        userBalance = genesisStake.balanceOf(user);
    }

    // Deploy the implementation contract and upgrade the proxy without reinitialization
    function _upgrade() internal {
        address impl = address(new GenesisStakeV2(token, rewardsDistributor));

        vm.prank(admin);
        proxyAdmin.upgradeAndCall(ITransparentUpgradeableProxy(address(genesisStake)), impl, "");
    }

    // Check the state of the contract after the upgrade
    function _checkState() internal view {
        assertEq(address(genesisStake.token()), token, "token mismatch");
        assertEq(genesisStake.rewardsDistributor(), rewardsDistributor, "rewardsDistributor mismatch");
        assertEq(genesisStake.balanceOf(user), userBalance, "balanceOf(user) mismatch");
    }

    // Test the migrateStake function
    function _testMigrateStake() internal {
        vm.startPrank(rewardsDistributor);

        // migrateStake should return if the user has no balance
        genesisStake.migrateStake(address(0));
        assertEq(token.balanceOf(rewardsDistributor), 0, "migrateStake improperly transferred tokens");

        // migrateStake should return the amount of tokens migrated
        genesisStake.migrateStake(user);
        assertEq(token.balanceOf(rewardsDistributor), userBalance, "migrateStake improperly transferred user's stake");

        vm.stopPrank();
    }
}
