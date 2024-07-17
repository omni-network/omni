// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { ProxyAdmin } from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import { ConfLevel } from "src/libraries/ConfLevel.sol";
import { XTypes } from "src/libraries/XTypes.sol";
import { OmniPortal } from "src/xchain/OmniPortal.sol";
import { Admin } from "script/admin/Admin.s.sol";
import { EIP1967Helper } from "script/utils/EIP1967Helper.sol";
import { PortalHarness } from "test/xchain/common/PortalHarness.sol";
import { MockFeeOracle } from "test/utils/MockFeeOracle.sol";
import { Test } from "forge-std/Test.sol";

/**
 * @title Admin_Test
 * @notice Test suite for Admin script.
 */
contract Admin_Test is Test {
    // test chain Ids, used to set network and make test xcalls
    uint64 constant thisChainId = 1;
    uint64 constant thatChainId = 2;

    function test_pause_unpause() public {
        Admin a = new Admin();

        address admin = makeAddr("admin");
        address portal = deployPortal(admin);

        // no revert
        makeXCall(portal);

        // pause
        a.pausePortal(admin, portal);

        // should revert
        vm.expectRevert("OmniPortal: paused");
        makeXCall(portal);

        // unpause
        a.unpausePortal(admin, portal);

        // no revert
        makeXCall(portal);
    }

    function test_upgrade() public {
        Admin a = new Admin();

        address admin = makeAddr("admin");
        address deployer = makeAddr("deployer");
        address portal = deployPortal(admin);

        // no revert
        makeXCall(portal);

        address expectedImplAfter = vm.computeCreateAddress(deployer, 0);
        address proxyAdmin = EIP1967Helper.getAdmin(portal);
        bytes memory upgradeCalldata = new bytes(0);

        // upgrade
        vm.expectCall(
            proxyAdmin,
            abi.encodeWithSelector(ProxyAdmin.upgradeAndCall.selector, portal, expectedImplAfter, upgradeCalldata)
        );
        a.upgradePortal(admin, deployer, portal, upgradeCalldata);

        // check impl changed
        assertEq(expectedImplAfter, EIP1967Helper.getImplementation(portal));

        // no revert
        makeXCall(portal);
    }

    //////////////////////////////////////////////////////////////////////////////
    //                              Utils                                       //
    //////////////////////////////////////////////////////////////////////////////

    function getNetwork() internal pure returns (XTypes.Chain[] memory) {
        uint64[] memory shards = new uint64[](2);
        shards[0] = uint64(ConfLevel.Finalized);
        shards[1] = uint64(ConfLevel.Latest);

        XTypes.Chain[] memory network = new XTypes.Chain[](2);
        network[0] = XTypes.Chain({ chainId: thisChainId, shards: shards });
        network[1] = XTypes.Chain({ chainId: thatChainId, shards: shards });

        return network;
    }

    function deployPortal(address admin) internal returns (address) {
        XTypes.Validator[] memory validators = new XTypes.Validator[](1);
        validators[0] = XTypes.Validator({ addr: address(0x123), power: 1 });

        address impl = address(new PortalHarness());
        bytes memory initializer = abi.encodeWithSelector(
            OmniPortal.initialize.selector,
            OmniPortal.InitParams({
                owner: admin,
                feeOracle: address(new MockFeeOracle(1 gwei)),
                omniChainId: 166,
                omniCChainId: 1_000_166,
                xmsgMaxGasLimit: 5_000_000,
                xmsgMinGasLimit: 21_000,
                xmsgMaxDataSize: 20_000,
                xreceiptMaxErrorSize: 256,
                xsubValsetCutoff: 10,
                cChainXMsgOffset: 1,
                cChainXBlockOffset: 1,
                valSetId: 1,
                validators: validators
            })
        );
        address proxy = address(new TransparentUpgradeableProxy(impl, admin, initializer));

        // setNetwork, so we can make test calls
        vm.chainId(thisChainId);
        PortalHarness(proxy).setNetworkNoAuth(getNetwork());

        return address(proxy);
    }

    function makeXCall(address portal) internal {
        vm.deal(address(this), 1 gwei);

        uint8 conf = ConfLevel.Finalized;
        address to = address(0x1234);
        bytes memory data = abi.encodeWithSignature("test()");
        uint64 gasLimit = 100_000;

        vm.chainId(thisChainId);
        OmniPortal(portal).xcall{ value: 1 gwei }(thatChainId, conf, to, data, gasLimit);
    }
}
