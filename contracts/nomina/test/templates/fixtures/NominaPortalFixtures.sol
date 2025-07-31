// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.30;

import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { NominaPortal } from "src/xchain/NominaPortal.sol";
import { PortalHarness } from "test/xchain/common/PortalHarness.sol";
import { MockFeeOracle } from "test/utils/MockFeeOracle.sol";
import { XSubGen } from "test/utils/XSubGen.sol";
import { Test } from "forge-std/Test.sol";

/// @dev NominaPortal test fixtures
contract NominaPortalFixtures is Test {
    uint8 constant xsubValsetCutoff = 10;
    uint16 constant xmsgMaxDataSize = 20_000;
    uint16 constant xreceiptMaxErrorSize = 256;
    uint64 constant xmsgMaxGasLimit = 5_000_000;
    uint64 constant xmsgMinGasLimit = 21_000;
    uint64 constant nominaChainId = 166;
    uint64 constant nominaCChainID = 1_000_166;
    uint64 constant initialCChainMsgOffset = 1;
    uint64 constant initialCChainBlockOffset = 1;
    uint64 constant genesisValsetId = 1;

    uint64 constant broadcastChainId = 0;
    address constant cChainSender = address(0);
    address constant virtualPortalAddress = address(0);

    address owner = makeAddr("owner");
    PortalHarness portal;
    XSubGen xsubgen;

    function setUp() public {
        xsubgen = new XSubGen();

        MockFeeOracle feeOracle = new MockFeeOracle(1 gwei);
        address impl = address(new PortalHarness());

        portal = PortalHarness(
            address(
                new TransparentUpgradeableProxy(
                    impl,
                    owner,
                    abi.encodeWithSelector(
                        NominaPortal.initialize.selector,
                        NominaPortal.InitParams(
                            owner,
                            address(feeOracle),
                            nominaChainId,
                            nominaCChainID,
                            xmsgMaxGasLimit,
                            xmsgMinGasLimit,
                            xmsgMaxDataSize,
                            xreceiptMaxErrorSize,
                            xsubValsetCutoff,
                            initialCChainMsgOffset,
                            initialCChainBlockOffset,
                            genesisValsetId,
                            xsubgen.getVals(genesisValsetId)
                        )
                    )
                )
            )
        );

        xsubgen.setPortal(address(portal));
    }
}
