// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { XRegistry } from "src/protocol/XRegistry.sol";
import { PortalRegistry } from "src/protocol/PortalRegistry.sol";
import { MockPortal } from "test/utils/MockPortal.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";
import { OmniPortal } from "src/protocol/OmniPortal.sol";
import { XRegistryReplica } from "src/protocol/XRegistryReplica.sol";
import { XRegistryNames } from "src/libraries/XRegistryNames.sol";

import { Test } from "forge-std/Test.sol";

contract PortalRegistry_Test is Test {
    MockPortal portal;
    XRegistry xreg;
    PortalRegistry preg;

    address owner = makeAddr("owner");

    address replica1 = makeAddr("replica1");
    address replica2 = makeAddr("replica2");
    address replica3 = makeAddr("replica3");

    address portal1 = makeAddr("portal1");
    address portal2 = makeAddr("portal2");
    address portal3 = makeAddr("portal3");

    mapping(uint64 => address) replicas;

    PortalRegistry.Deployment dep1 =
        PortalRegistry.Deployment({ chainId: 1, addr: portal1, deployHeight: 1234, finalizationStrat: "finalized" });

    PortalRegistry.Deployment dep2 =
        PortalRegistry.Deployment({ chainId: 2, addr: portal2, deployHeight: 5678, finalizationStrat: "finalized" });

    PortalRegistry.Deployment dep3 =
        PortalRegistry.Deployment({ chainId: 3, addr: portal3, deployHeight: 91_011, finalizationStrat: "finalized" });

    function setUp() public {
        xreg = XRegistry(Predeploys.XRegistry);
        preg = PortalRegistry(Predeploys.PortalRegistry);

        XRegistry tmpXReg = new XRegistry();
        PortalRegistry tmpPReg = new PortalRegistry();

        vm.etch(Predeploys.XRegistry, address(tmpXReg).code);
        vm.etch(Predeploys.PortalRegistry, address(tmpPReg).code);

        vm.store(Predeploys.XRegistry, 0, bytes32(uint256(uint160(owner))));
        vm.store(Predeploys.PortalRegistry, 0, bytes32(uint256(uint160(owner))));

        require(xreg.owner() == owner, "XRegistry owner not set");
        require(preg.owner() == owner, "PortalRegistry owner not set");

        portal = new MockPortal();

        vm.startPrank(owner);
        xreg.setPortal(address(portal));
        xreg.setReplica(1, replica1);
        xreg.setReplica(2, replica2);
        xreg.setReplica(3, replica3);
        vm.stopPrank();

        vm.deal(owner, 1000 ether);

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

        vm.startPrank(owner);
        preg.register{ value: fee }(dep1);

        assertEq(xreg.get(1, XRegistryNames.OmniPortal, Predeploys.PortalRegistry), portal1);

        //
        // register deployment 2
        //
        fee = preg.registrationFee(dep2);

        vm.startPrank(owner);

        // chain 1 gets new chain 2 portal
        _expectXCall({ destChainId: 1, depChainId: 2, addr: portal2 });

        // chain 2 gets existing chain 1 portal
        _expectXCall({ destChainId: 2, depChainId: 1, addr: portal1 });

        preg.register{ value: fee }(dep2);

        //
        // register deployment 3
        //
        fee = preg.registrationFee(dep3);

        vm.startPrank(owner);

        // chain 1 gets new chain 3 portal
        _expectXCall({ destChainId: 1, depChainId: 3, addr: portal3 });

        // chain 2 gets new chain 3 portal
        _expectXCall({ destChainId: 2, depChainId: 3, addr: portal3 });

        // chain 3 gets existing chain 1 portal
        _expectXCall({ destChainId: 3, depChainId: 1, addr: portal1 });

        // chain 3 gets existing chain 2 portal
        _expectXCall({ destChainId: 3, depChainId: 2, addr: portal2 });

        preg.register{ value: fee }(dep3);
    }

    /**
     * @notice Expect that the correct portal.xcall(...) was made, to set the portal address on a replica.
     * @param destChainId   The chainId on which to set the deployment.
     * @param depChainId    The chainId of the deployment.
     * @param addr          The deployment address.
     */
    function _expectXCall(uint64 destChainId, uint64 depChainId, address addr) internal {
        bytes memory data = abi.encodeWithSelector(
            XRegistryReplica.set.selector, depChainId, XRegistryNames.OmniPortal, Predeploys.PortalRegistry, addr
        );

        uint256 fee = portal.feeFor(destChainId, data, xreg.XSET_GAS_LIMIT());

        vm.expectCall(
            address(portal),
            fee,
            abi.encodeWithSignature(
                "xcall(uint64,address,bytes,uint64)", destChainId, replicas[destChainId], data, xreg.XSET_GAS_LIMIT()
            )
        );
    }
}
