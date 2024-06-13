// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { PortalRegistry } from "src/xchain/PortalRegistry.sol";
import { MockPortal } from "test/utils/MockPortal.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";
import { OmniPortal } from "src/xchain/OmniPortal.sol";
import { ConfLevel } from "src/libraries/ConfLevel.sol";
import { Test } from "forge-std/Test.sol";

contract PortalRegistry_Test is Test {
    // copied from PortalRegistry.sol
    event PortalRegistered(uint64 indexed chainId, address indexed addr, uint64 deployHeight, uint64[] shards);

    PortalRegistryHarness reg;
    address owner;

    function setUp() public {
        owner = makeAddr("owner");
        reg = new PortalRegistryHarness(owner);
    }

    function test_register() public {
        PortalRegistry.Deployment memory dep;

        // no zero address
        vm.expectRevert("PortalRegistry: no zero addr");
        vm.prank(owner);
        reg.register(dep);

        // no zero chain ID
        dep.addr = makeAddr("addr");
        dep.chainId = 0;
        vm.expectRevert("PortalRegistry: no zero chain ID");
        vm.prank(owner);
        reg.register(dep);

        // Must have shards
        dep.chainId = 1;
        vm.expectRevert("PortalRegistry: no shards");
        vm.prank(owner);
        reg.register(dep);

        // Must have valid shards
        dep.shards = new uint64[](2);
        dep.shards[0] = 12_341_234;
        dep.shards[1] = 56_785_678;
        vm.expectRevert("PortalRegistry: invalid shard");
        vm.prank(owner);
        reg.register(dep);

        // success
        dep.shards[0] = ConfLevel.Finalized;
        dep.shards[1] = ConfLevel.Latest;
        vm.expectEmit();
        emit PortalRegistered(dep.chainId, dep.addr, dep.deployHeight, dep.shards);
        vm.prank(owner);
        reg.register(dep);

        assertEq(reg.get(dep.chainId).chainId, dep.chainId);
        assertEq(reg.get(dep.chainId).addr, dep.addr);
        assertEq(reg.get(dep.chainId).deployHeight, dep.deployHeight);
        assertEq(reg.get(dep.chainId).shards[0], dep.shards[0]);
        assertEq(reg.get(dep.chainId).shards[1], dep.shards[1]);
        assertEq(reg.list().length, 1);

        // cannot register the same chain twice
        vm.expectRevert("PortalRegistry: already set");
        vm.prank(owner);
        reg.register(dep);

        // can register multiple chains
        for (uint64 i = 2; i <= 5; i++) {
            dep = _deployment(i);

            vm.expectEmit();
            emit PortalRegistered(i, dep.addr, dep.deployHeight, dep.shards);

            vm.prank(owner);
            reg.register(dep);

            assertEq(reg.get(dep.chainId).chainId, dep.chainId);
            assertEq(reg.get(dep.chainId).addr, dep.addr);
            assertEq(reg.get(dep.chainId).deployHeight, dep.deployHeight);
            assertEq(reg.get(dep.chainId).shards[0], dep.shards[0]);
            assertEq(reg.get(dep.chainId).shards[1], dep.shards[1]);
            assertEq(reg.list().length, i);
        }
    }

    function _deployment(uint64 chainId) internal returns (PortalRegistry.Deployment memory) {
        PortalRegistry.Deployment memory dep = PortalRegistry.Deployment({
            chainId: chainId,
            addr: makeAddr(string(abi.encodePacked("portal", chainId))),
            deployHeight: chainId * 1234,
            shards: new uint64[](2)
        });

        dep.shards[0] = ConfLevel.Finalized;
        dep.shards[1] = ConfLevel.Latest;

        return dep;
    }
}

/**
 * @dev Wrapper around PortalRegistry that adds a constructor.
 */
contract PortalRegistryHarness is PortalRegistry {
    constructor(address _owner) {
        _transferOwnership(_owner);
    }
}
