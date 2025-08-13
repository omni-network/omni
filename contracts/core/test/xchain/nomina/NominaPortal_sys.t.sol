// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { NominaPortalFixtures } from "test/templates/fixtures/nomina/NominaPortalFixtures.sol";
import { NominaPortal } from "src/xchain/nomina/NominaPortal.sol";
import { XTypes } from "src/libraries/XTypes.sol";
import { ConfLevel } from "src/libraries/ConfLevel.sol";
import { console2 as console } from "forge-std/console2.sol";

/**
 * @title NominaPortal_sys_Test
 * @dev Test of NominaPortal sys calls
 */
contract NominaPortal_sys_Test is NominaPortalFixtures {
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

        XTypes.BlockHeader memory xheader = xsubgen.makeXHeader(nominaCChainID, ConfLevel.Finalized);
        XTypes.Msg[] memory msgs = new XTypes.Msg[](1);
        msgs[0] = _sysXMsg(abi.encodeCall(NominaPortal.setNetwork, (network)));

        vm.chainId(1);
        portal.xsubmit(xsubgen.makeXSub(1, xheader, msgs, xsubgen.msgFlagsForDest(msgs, broadcastChainId)));

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

    /// @dev Test that adding a valset of max size is below reasonable gas limits
    function test_addValidatorSet_maxGas() public {
        uint256 maxVals = 30; // matches the max validators in the NominaPortal contract
        uint64 power = 100;

        XTypes.Validator[] memory validators = new XTypes.Validator[](maxVals);
        for (uint256 i = 1; i < validators.length + 1; i++) {
            validators[i - 1] = XTypes.Validator({ addr: address(uint160(i)), power: power });
        }

        XTypes.BlockHeader memory xheader = xsubgen.makeXHeader(nominaCChainID, ConfLevel.Finalized);
        XTypes.Msg[] memory msgs = new XTypes.Msg[](1);
        msgs[0] = _sysXMsg(abi.encodeCall(NominaPortal.addValidatorSet, (2, validators)));

        XTypes.Submission memory xsub =
            xsubgen.makeXSub(1, xheader, msgs, xsubgen.msgFlagsForDest(msgs, broadcastChainId));

        vm.chainId(1);
        uint256 gasBefore = gasleft();
        portal.xsubmit(xsub);
        uint256 gasUsed = gasBefore - gasleft();
        console.log("Gas used: ", gasUsed);

        assertLt(gasUsed, 1_000_000); // assert under 1M gas, well below block gas limits
        assertEq(portal.valSetTotalPower(2), power * maxVals);
    }

    /// @dev Test syscalls (xcalls to VirtualPortalAddress) are properly authorized
    function test_syscall_auth() public {
        uint64 destChainId = 1;
        vm.chainId(destChainId);

        XTypes.Submission memory xsub;
        bytes memory data = abi.encodeCall(NominaPortal.setNetwork, (new XTypes.Chain[](0))); // must be known syscall
        XTypes.BlockHeader memory xheader = xsubgen.makeXHeader(nominaCChainID, ConfLevel.Finalized);
        XTypes.Msg[] memory xmsgs = new XTypes.Msg[](1);
        xmsgs[0] = _sysXMsg(data);

        // source chain must be nominaCChainID
        xheader.sourceChainId = 1234;
        xmsgs[0].offset = 1; // changing source chain id changes ofset required
        xsub = xsubgen.makeXSub(1, xheader, xmsgs, xsubgen.msgFlagsForDest(xmsgs, broadcastChainId));
        vm.expectRevert("NominaPortal: invalid syscall");
        portal.xsubmit(xsub);
        xheader.sourceChainId = nominaCChainID;
        xmsgs[0].offset = 2; // cchain initialized with offset 1

        // shard must be broadcast shard
        xmsgs[0].shardId = 1234;
        xmsgs[0].offset = 1; // changing shard id changes ofset required
        xsub = xsubgen.makeXSub(1, xheader, xmsgs, xsubgen.msgFlagsForDest(xmsgs, broadcastChainId));
        vm.expectRevert("NominaPortal: invalid syscall");
        portal.xsubmit(xsub);
        xmsgs[0].shardId = ConfLevel.toBroadcastShard(ConfLevel.Finalized);
        xmsgs[0].offset = 2; // cchain initialized with offset 1

        // must be to broadcast chain
        xmsgs[0].destChainId = destChainId;
        xsub = xsubgen.makeXSub(1, xheader, xmsgs, xsubgen.msgFlagsForDest(xmsgs, destChainId));
        vm.expectRevert("NominaPortal: invalid syscall");
        vm.chainId(destChainId);
        portal.xsubmit(xsub);
        xmsgs[0].destChainId = broadcastChainId;

        // sender must be cChainSender
        xmsgs[0].sender = address(1234);
        xsub = xsubgen.makeXSub(1, xheader, xmsgs, xsubgen.msgFlagsForDest(xmsgs, broadcastChainId));
        vm.expectRevert("NominaPortal: invalid syscall");
        portal.xsubmit(xsub);
        xmsgs[0].sender = cChainSender;

        // data must be a known syscall
        xmsgs[0].data = abi.encodePacked("not a known syscall");
        xsub = xsubgen.makeXSub(1, xheader, xmsgs, xsubgen.msgFlagsForDest(xmsgs, broadcastChainId));
        vm.expectRevert("NominaPortal: invalid syscall");
        portal.xsubmit(xsub);

        // if xmsg.to != VirtualPortalAddress, the xmsg must not have been broadcast from the cchain
        xmsgs[0].to = address(1234);
        xsub = xsubgen.makeXSub(1, xheader, xmsgs, xsubgen.msgFlagsForDest(xmsgs, broadcastChainId));
        vm.expectRevert("NominaPortal: invalid xcall");
        portal.xsubmit(xsub);

        // just changing source chain id is not enough to bypass the check
        xheader.sourceChainId = 1234;
        xmsgs[0].offset = 1;
        xsub = xsubgen.makeXSub(1, xheader, xmsgs, xsubgen.msgFlagsForDest(xmsgs, broadcastChainId));
        vm.expectRevert("NominaPortal: invalid xcall");
        portal.xsubmit(xsub);

        // changing dest chain id is not enough
        xmsgs[0].destChainId = destChainId;
        xsub = xsubgen.makeXSub(1, xheader, xmsgs, xsubgen.msgFlagsForDest(xmsgs, destChainId));
        vm.expectRevert("NominaPortal: invalid xcall");
        portal.xsubmit(xsub);

        // changing shard id is not enough
        xmsgs[0].shardId = ConfLevel.Finalized;
        xsub = xsubgen.makeXSub(1, xheader, xmsgs, xsubgen.msgFlagsForDest(xmsgs, destChainId));
        vm.expectRevert("NominaPortal: invalid xcall");
        portal.xsubmit(xsub);

        // changing source chain, dest chain / shard and sender is enough
        xmsgs[0].sender = address(1234);
        xsub = xsubgen.makeXSub(1, xheader, xmsgs, xsubgen.msgFlagsForDest(xmsgs, destChainId));
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
