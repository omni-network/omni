// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { GenesisStake } from "src/token/GenesisStake.sol";
import { ProxyAdmin } from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {
    ITransparentUpgradeableProxy
} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { EIP1967Helper } from "script/utils/EIP1967Helper.sol";
import { Test } from "forge-std/Test.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { VmSafe } from "forge-std/Vm.sol";

contract GenesisStakeUpgradeTest is Test {
    using SafeTransferLib for address;

    GenesisStake internal genesisStake = GenesisStake(0xD2639676dA3dEA5491d27DA19340556b3a7d58B8);

    ProxyAdmin internal proxyAdmin;
    address internal admin;
    address internal owner;
    address internal token;
    uint256 internal unbondingPeriod;
    uint256 internal userBalance;
    uint256 internal userUnstakedAt;
    bool internal isOpen;

    address internal rewardsDistributor = makeAddr("rewardsDistributor");
    address internal user = makeAddr("user");

    function run() public {
        (VmSafe.CallerMode mode,,) = vm.readCallers();
        require(mode == VmSafe.CallerMode.None, "no broadcast");

        _setup();
        _openStakeUnstakeClose();
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
        unbondingPeriod = genesisStake.unbondingPeriod();
    }

    // Open staking, user stake, warp past timelock, user unstake, and then close staking
    // This eliminates the need to retrieve a valid staker address
    function _openStakeUnstakeClose() internal {
        vm.startPrank(owner);
        if (!genesisStake.isOpen()) genesisStake.open();
        vm.stopPrank();

        deal(token, user, 1000 ether);
        vm.startPrank(user);
        token.safeApprove(address(genesisStake), type(uint256).max);
        genesisStake.stake(1000 ether);
        vm.warp(block.timestamp + unbondingPeriod + 1);
        genesisStake.unstake();
        vm.stopPrank();

        vm.prank(owner);
        genesisStake.close();
    }

    // Cache values relevant for testing post-upgrade
    function _cacheState() internal {
        userBalance = genesisStake.balanceOf(user);
        userUnstakedAt = genesisStake.unstakedAt(user);
        isOpen = genesisStake.isOpen();
    }

    // Deploy the implementation contract and upgrade the proxy without reinitialization
    function _upgrade() internal {
        address impl = address(new GenesisStake(token, rewardsDistributor));

        vm.prank(admin);
        proxyAdmin.upgradeAndCall(ITransparentUpgradeableProxy(address(genesisStake)), impl, "");
    }

    // Check the state of the contract after the upgrade
    function _checkState() internal view {
        assertEq(address(genesisStake.token()), token, "token mismatch");
        assertEq(genesisStake.rewardsDistributor(), rewardsDistributor, "rewardsDistributor mismatch");
        assertEq(genesisStake.unbondingPeriod(), unbondingPeriod, "unbondingPeriod mismatch");
        assertEq(genesisStake.balanceOf(user), userBalance, "balanceOf(user) mismatch");
        assertEq(genesisStake.unstakedAt(user), userUnstakedAt, "unstakedAt(user) mismatch");
        assertEq(genesisStake.isOpen(), isOpen, "isOpen mismatch");
    }

    // Test the newly introduced migrateStake function
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
