// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { OmniPortal } from "src/xchain/OmniPortal.sol";
import { PortalHarness } from "test/xchain/common/PortalHarness.sol";
import { MockFeeOracle } from "test/utils/MockFeeOracle.sol";
import { XSubGen } from "test/utils/XSubGen.sol";
import { Test } from "forge-std/Test.sol";

/// @dev OmniPortal test fixtures
contract OmniPortalFixtures is Test {
    uint8 constant xsubValsetCutoff = 10;
    uint16 constant xmsgMaxDataSize = 20_000;
    uint16 constant xreceiptMaxErrorSize = 256;
    uint64 constant xmsgMaxGasLimit = 5_000_000;
    uint64 constant xmsgMinGasLimit = 21_000;
    uint64 constant omniChainId = 166;
    uint64 constant omniCChainID = 1_000_166;
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
                        OmniPortal.initialize.selector,
                        OmniPortal.InitParams(
                            owner,
                            address(feeOracle),
                            omniChainId,
                            omniCChainID,
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
