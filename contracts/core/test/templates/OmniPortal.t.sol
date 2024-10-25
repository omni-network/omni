// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { OmniPortalFixtures } from "./fixtures/OmniPortalFixtures.sol";
import { Counter } from "test/xchain/common/Counter.sol";
import { ConfLevel } from "src/libraries/ConfLevel.sol";
import { XTypes } from "src/libraries/XTypes.sol";

/**
 * @title OmniPortal_Test
 * @notice Template test suite for src/xchain/OmniPortal.sol
 * @dev Available fixtures (see OmniPortalFixtures.sol):
 *
 *  PortalHarness portal;   // the OmniPortal contract, with additional public setters
 *  XSubGen xsubgen;        // a utility contract for generating XSubmissions
 *  address owner;          // the owner address for the portal
 *
 * See test_example() for an example test case using XSubGen.
 */
contract OmniPortal_Test is OmniPortalFixtures {
    /**
     * @notice Example test case for OmniPortal contract.
     * @dev This test case shows how to use xsubgen (XSubGen) to create a submission
     *      and submit it to the portal.
     */
    function test_example() public {
        Counter counter = new Counter(portal); // new Counter to increment via xmsg
        uint64 srcChainId = 1; // source chain id for the xsubmission
        uint64 valSetId = 1; // use genesis valset for the xsubmission
        uint64 thisChainId = uint64(block.chainid); // destChainId for each xmsg
        uint64 shardId = uint64(ConfLevel.Finalized); // shardId for each xmsg (shard == conf level)
        address sender = makeAddr("sender"); // xmsg sender

        // mock block header
        XTypes.BlockHeader memory xheader = xsubgen.makeXHeader(srcChainId, ConfLevel.Finalized);

        // xmsgs to Counter
        XTypes.Msg[] memory msgs = new XTypes.Msg[](3);
        msgs[0] = XTypes.Msg({
            destChainId: thisChainId,
            shardId: shardId,
            offset: 1,
            sender: sender,
            to: address(counter),
            data: abi.encodeCall(Counter.increment, ()),
            gasLimit: 100_000
        });
        msgs[1] = XTypes.Msg({
            destChainId: thisChainId,
            shardId: uint64(ConfLevel.Finalized),
            offset: 2,
            sender: sender,
            to: address(counter),
            data: abi.encodeCall(Counter.increment, ()),
            gasLimit: 100_000
        });
        msgs[2] = XTypes.Msg({
            destChainId: thisChainId - 1, // some other chain - should not be included in submission
            shardId: shardId,
            offset: 1,
            sender: sender,
            to: address(counter),
            data: abi.encodeCall(Counter.increment, ()),
            gasLimit: 100_000
        });

        // select which xmsgs to include in submission (exclude xmsg to "some other chain")
        bool[] memory msgFlags = new bool[](3);
        msgFlags[0] = true;
        msgFlags[1] = true;
        msgFlags[2] = false;

        // make and submit the xsubmission
        XTypes.Submission memory xsub = xsubgen.makeXSub(valSetId, thisChainId, xheader, msgs, msgFlags);
        portal.xsubmit(xsub);

        // check that the counter has been incremented
        assertEq(counter.count(), 2);
    }
}
