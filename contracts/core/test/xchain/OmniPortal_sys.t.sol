// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OmniPortalFixtures } from "test/templates/fixtures/OmniPortalFixtures.sol";
import { OmniPortal } from "src/xchain/OmniPortal.sol";
import { XTypes } from "src/libraries/XTypes.sol";
import { ConfLevel } from "src/libraries/ConfLevel.sol";

/**
 * @title OmniPortal_sys_Test
 * @dev Test of OmniPortal sys calls
 */
contract OmniPortal_sys_Test is OmniPortalFixtures {
    function test_setNetwork() public {
        uint64[] memory chain1Shards = new uint64[](2);
        chain1Shards[0] = 3;
        chain1Shards[1] = 4;

        uint64[] memory chain2Shards = new uint64[](2);
        chain2Shards[0] = 5;
        chain2Shards[1] = 6;

        XTypes.Chain memory chain1 = XTypes.Chain({ chainId: 1, shards: chain1Shards });
        XTypes.Chain memory chain2 = XTypes.Chain({ chainId: 2, shards: chain2Shards });

        XTypes.Chain[] memory network = new XTypes.Chain[](2);
        network[0] = chain1;
        network[1] = chain2;

        XTypes.BlockHeader memory xheader = xsubgen.makeXHeader(omniCChainID, ConfLevel.Finalized);
        XTypes.Msg[] memory msgs = new XTypes.Msg[](1);
        msgs[0] = _sysXMsg(abi.encodeCall(OmniPortal.setNetwork, (network)));

        vm.chainId(1);
        portal.xsubmit(xsubgen.makeXSub(xheader, msgs));

        XTypes.Chain[] memory read = portal.network();

        assertEq(read.length, 2);
        assertEq(read[0].chainId, chain1.chainId);
        assertEq(read[1].chainId, chain2.chainId);
        assertEq(read[0].shards.length, 2);
        assertEq(read[1].shards.length, 2);
        assertEq(read[0].shards[0], chain1Shards[0]);
        assertEq(read[0].shards[1], chain1Shards[1]);
        assertEq(read[1].shards[0], chain2Shards[0]);
        assertEq(read[1].shards[1], chain2Shards[1]);

        // chain 1 is not a supported dest
        assertFalse(portal.isSupportedDest(chain1.chainId));

        // chain 2 is a supported dest
        assertTrue(portal.isSupportedDest(chain2.chainId));

        // chain 1 shards are supported
        for (uint256 i = 0; i < chain1Shards.length; i++) {
            assertTrue(portal.isSupportedShard(chain1Shards[i]));
        }

        // chain 2 shards are not supported
        // NOTE: this works because chain 2 shards differ from chain 1 shards
        // They could have the same supported shards
        for (uint256 i = 0; i < chain2Shards.length; i++) {
            assertFalse(portal.isSupportedShard(chain2Shards[i]));
        }
    }

    /// @dev Test syscalls (xcalls to VirtualPortalAddress) are properly authorized
    function test_syscall_auth() public {
        uint64 destChainId = 1;
        vm.chainId(destChainId);

        XTypes.Submission memory xsub;
        bytes memory data = abi.encodeCall(OmniPortal.setNetwork, (new XTypes.Chain[](0))); // must be known syscall
        XTypes.BlockHeader memory xheader = xsubgen.makeXHeader(omniCChainID, ConfLevel.Finalized);
        XTypes.Msg[] memory xmsgs = new XTypes.Msg[](1);
        xmsgs[0] = _sysXMsg(data);

        // source chain must be omniCChainID
        xheader.sourceChainId = 1234;
        xmsgs[0].offset = 1; // changing source chain id changes ofset required
        xsub = xsubgen.makeXSub(xheader, xmsgs);
        vm.expectRevert("OmniPortal: invalid syscall");
        portal.xsubmit(xsub);
        xheader.sourceChainId = omniCChainID;
        xmsgs[0].offset = 2; // cchain initialized with offset 1

        // shard must be broadcast shard
        xmsgs[0].shardId = 1234;
        xmsgs[0].offset = 1; // changing shard id changes ofset required
        xsub = xsubgen.makeXSub(xheader, xmsgs);
        vm.expectRevert("OmniPortal: invalid syscall");
        portal.xsubmit(xsub);
        xmsgs[0].shardId = ConfLevel.toBroadcastShard(ConfLevel.Finalized);
        xmsgs[0].offset = 2; // cchain initialized with offset 1

        // must be to broadcast chain
        xmsgs[0].destChainId = destChainId;
        xsub = xsubgen.makeXSub(xheader, xmsgs);
        vm.expectRevert("OmniPortal: invalid syscall");
        vm.chainId(destChainId);
        portal.xsubmit(xsub);
        xmsgs[0].destChainId = broadcastChainId;

        // sender must be cChainSender
        xmsgs[0].sender = address(1234);
        xsub = xsubgen.makeXSub(xheader, xmsgs);
        vm.expectRevert("OmniPortal: invalid syscall");
        portal.xsubmit(xsub);
        xmsgs[0].sender = cChainSender;

        // data must be a known syscall
        xmsgs[0].data = abi.encodePacked("not a known syscall");
        xsub = xsubgen.makeXSub(xheader, xmsgs);
        vm.expectRevert("OmniPortal: invalid syscall");
        portal.xsubmit(xsub);

        // if xmsg.to != VirtualPortalAddress, the xmsg must not have been broadcast from the cchain
        xmsgs[0].to = address(1234);
        xsub = xsubgen.makeXSub(xheader, xmsgs);
        vm.expectRevert("OmniPortal: invalid xcall");
        portal.xsubmit(xsub);

        // just changing source chain id is not enough to bypass the check
        xheader.sourceChainId = 1234;
        xmsgs[0].offset = 1;
        xsub = xsubgen.makeXSub(xheader, xmsgs);
        vm.expectRevert("OmniPortal: invalid xcall");
        portal.xsubmit(xsub);

        // changing dest chain id is not enough
        xmsgs[0].destChainId = destChainId;
        xsub = xsubgen.makeXSub(xheader, xmsgs);
        vm.expectRevert("OmniPortal: invalid xcall");
        portal.xsubmit(xsub);

        // changing shard id is not enough
        xmsgs[0].shardId = ConfLevel.Finalized;
        xsub = xsubgen.makeXSub(xheader, xmsgs);
        vm.expectRevert("OmniPortal: invalid xcall");
        portal.xsubmit(xsub);

        // changing source chain, dest chain / shard and sender is enough
        xmsgs[0].sender = address(1234);
        xsub = xsubgen.makeXSub(xheader, xmsgs);
        portal.xsubmit(xsub);
    }

    function _sysXMsg(bytes memory data) internal pure returns (XTypes.Msg memory) {
        return _sysXMsg(data, 2);
    }

    function _sysXMsg(bytes memory data, uint64 offset) internal pure returns (XTypes.Msg memory) {
        return XTypes.Msg({
            destChainId: broadcastChainId,
            shardId: ConfLevel.toBroadcastShard(ConfLevel.Finalized),
            offset: offset,
            sender: cChainSender,
            to: virtualPortalAddress,
            data: data,
            gasLimit: 0
        });
    }
}
