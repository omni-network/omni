// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { XRegistry } from "src/xchain/XRegistry.sol";
import { PortalRegistry } from "src/xchain/PortalRegistry.sol";
import { MockPortal } from "test/utils/MockPortal.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";
import { OmniPortal } from "src/xchain/OmniPortal.sol";
import { XRegistryBase } from "src/xchain/XRegistryBase.sol";
import { XRegistryReplica } from "src/xchain/XRegistryReplica.sol";
import { XRegistryNames } from "src/libraries/XRegistryNames.sol";
import { ConfLevel } from "src/libraries/ConfLevel.sol";

import { Test } from "forge-std/Test.sol";

contract PortalRegistry_Test is Test {
    MockPortalHarness portal;
    XRegistryHarness xreg;
    PortalRegistryHarness preg;

    address owner = makeAddr("owner");

    address replica1 = makeAddr("replica1");
    address replica2 = makeAddr("replica2");
    address replica3 = makeAddr("replica3");

    address portal1 = makeAddr("portal1");
    address portal2 = makeAddr("portal2");
    address portal3 = makeAddr("portal3");

    mapping(uint64 => address) replicas;

    PortalRegistry.Deployment dep1 = PortalRegistry.Deployment({
        chainId: 1,
        addr: portal1,
        deployHeight: 1234,
        shards: new uint64[](0) // added in setUp
     });

    PortalRegistry.Deployment dep2 = PortalRegistry.Deployment({
        chainId: 2,
        addr: portal2,
        deployHeight: 5678,
        shards: new uint64[](0) // added in setUp
     });

    PortalRegistry.Deployment dep3 = PortalRegistry.Deployment({
        chainId: 3,
        addr: portal3,
        deployHeight: 91_011,
        shards: new uint64[](0) // added in setUp
     });

    function setUp() public {
        xreg = XRegistryHarness(Predeploys.XRegistry);
        preg = PortalRegistryHarness(Predeploys.PortalRegistry);

        uint64[] memory shards = new uint64[](2);
        shards[0] = ConfLevel.Finalized;
        shards[1] = ConfLevel.Latest;

        dep1.shards = shards;
        dep2.shards = shards;
        dep3.shards = shards;

        address tmpXReg = address(new XRegistryHarness());
        address tmpPReg = address(new PortalRegistryHarness());

        vm.etch(Predeploys.XRegistry, tmpXReg.code);
        vm.etch(Predeploys.PortalRegistry, tmpPReg.code);

        xreg.initialize(owner);
        preg.initialize(owner);

        require(xreg.owner() == owner, "XRegistry owner not set");
        require(preg.owner() == owner, "PortalRegistry owner not set");

        portal = new MockPortalHarness();

        vm.startPrank(owner);
        xreg.setPortal(address(portal));
        xreg.setReplica(1, replica1);
        xreg.setReplica(2, replica2);
        xreg.setReplica(3, replica3);
        vm.stopPrank();

        vm.deal(owner, 10_000 ether);

        assertTrue(xreg.isSupportedChain(1), "Chain 1 not supported");
        assertTrue(xreg.isSupportedChain(2), "Chain 2 not supported");
        assertTrue(xreg.isSupportedChain(3), "Chain 3 not supported");

        assertEq(xreg.replicas(1), replica1, "Replica 1 not set");
        assertEq(xreg.replicas(2), replica2, "Replica 2 not set");
        assertEq(xreg.replicas(3), replica3, "Replica 3 not set");

        replicas[1] = replica1;
        replicas[2] = replica2;
        replicas[3] = replica3;
    }

    function test_register() public {
        //
        // register deployment 1
        //
        uint256 fee = preg.registrationFee(dep1);

        // chain 1 gets itself
        _expectXCall({ destChainId: 1, depChainId: 1, dep: XRegistryBase.Deployment(dep1.addr, abi.encode(dep1.shards)) });

        vm.startPrank(owner);
        preg.register{ value: fee }(dep1);

        assertEq(xreg.get(1, XRegistryNames.OmniPortal, Predeploys.PortalRegistry).addr, portal1);
        assertEq(xreg.get(1, XRegistryNames.OmniPortal, Predeploys.PortalRegistry).metadata, abi.encode(dep1.shards));

        //
        // register deployment 2
        //
        fee = preg.registrationFee(dep2);

        vm.startPrank(owner);

        // chain 2 gets itself
        _expectXCall({ destChainId: 2, depChainId: 2, dep: XRegistryBase.Deployment(dep2.addr, abi.encode(dep2.shards)) });

        // chain 1 gets new chain 2 portal
        _expectXCall({ destChainId: 1, depChainId: 2, dep: XRegistryBase.Deployment(dep2.addr, abi.encode(dep2.shards)) });

        // chain 2 gets existing chain 1 portal
        _expectXCall({ destChainId: 2, depChainId: 1, dep: XRegistryBase.Deployment(dep1.addr, abi.encode(dep1.shards)) });

        preg.register{ value: fee }(dep2);

        //
        // register deployment 3
        //
        fee = preg.registrationFee(dep3);

        vm.startPrank(owner);

        // chain 3 gets itself
        _expectXCall({ destChainId: 3, depChainId: 3, dep: XRegistryBase.Deployment(dep3.addr, abi.encode(dep3.shards)) });

        // chain 1 gets new chain 3 portal
        _expectXCall({ destChainId: 1, depChainId: 3, dep: XRegistryBase.Deployment(dep3.addr, abi.encode(dep3.shards)) });

        // chain 2 gets new chain 3 portal
        _expectXCall({ destChainId: 2, depChainId: 3, dep: XRegistryBase.Deployment(dep3.addr, abi.encode(dep3.shards)) });

        // chain 3 gets existing chain 1 portal
        _expectXCall({ destChainId: 3, depChainId: 1, dep: XRegistryBase.Deployment(dep1.addr, abi.encode(dep1.shards)) });

        // chain 3 gets existing chain 2 portal
        _expectXCall({ destChainId: 3, depChainId: 2, dep: XRegistryBase.Deployment(dep2.addr, abi.encode(dep2.shards)) });

        preg.register{ value: fee }(dep3);
    }

    /**
     * @notice Expect that the correct portal.xcall(...) was made, to set the portal address on a replica.
     * @param destChainId   The chainId on which to set the deployment.
     * @param depChainId    The chainId of the deployment.
     * @param dep           The deployment to set.
     */
    function _expectXCall(uint64 destChainId, uint64 depChainId, XRegistry.Deployment memory dep) private {
        bytes memory data = abi.encodeWithSelector(
            XRegistryReplica.set.selector, depChainId, XRegistryNames.OmniPortal, Predeploys.PortalRegistry, dep
        );

        uint256 fee = portal.feeFor(destChainId, data, xreg.XSET_PORTAL_GAS_LIMIT());

        vm.expectCall(
            address(portal),
            fee,
            abi.encodeWithSignature(
                "xcall(uint64,uint8,address,bytes,uint64)",
                destChainId,
                ConfLevel.Finalized,
                replicas[destChainId],
                data,
                xreg.XSET_PORTAL_GAS_LIMIT()
            )
        );
    }
}

/**
 * @dev Wrapper around MockPortal (a user facing test utility) that adds initSourceChain, required by XRegistry.
 */
contract MockPortalHarness is MockPortal {
    function initSourceChain(uint64 _sourceChainId, uint64[] memory _shards) public {
        // do nothing
    }
}

/**
 * @dev Wrapper around XRegistry, that adds initializer.
 */
contract XRegistryHarness is XRegistry {
    function initialize(address _owner) external initializer {
        __Ownable_init();
        _transferOwnership(_owner);
    }
}

/**
 * @dev Wrapper around PortalRegistry, that adds initializer.
 */
contract PortalRegistryHarness is PortalRegistry {
    function initialize(address _owner) external initializer {
        __Ownable_init();
        _transferOwnership(_owner);
    }
}
