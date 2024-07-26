// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { XTypes } from "src/libraries/XTypes.sol";
import { IOmniPortal } from "src/interfaces/IOmniPortal.sol";

import { Base } from "./common/Base.sol";
import { Counter } from "./common/Counter.sol";
import { Vm } from "forge-std/Vm.sol";
import { console } from "forge-std/console.sol";

/**
 * @title OmniPortal_xsubmit_gas_Test
 * @dev Test exploring gas usage of xsubmit and dependent functions.
 */
contract OmniPortal_xsubmit_gas_Test is Base {
    function test_xsubmit_guzzle1_succeeds() public {
        _testGasSubmitXBlock("guzzle1");
    }

    function test_xsubmit_guzzle5_succeeds() public {
        _testGasSubmitXBlock("guzzle5");
    }

    function test_xsubmit_guzzle10_succeeds() public {
        _testGasSubmitXBlock("guzzle10");
    }

    function test_xsubmit_guzzle25_succeeds() public {
        _testGasSubmitXBlock("guzzle25");
    }

    function test_xsubmit_guzzle50_succeeds() public {
        _testGasSubmitXBlock("guzzle50");
    }

    function test_xsubmit_addValidator_succeeds() public {
        _testGasSubmitXBlock("addValSet2", broadcastChainId);
    }

    function test_singleExec() public {
        XTypes.Submission memory xsub = readXSubmission({ name: "guzzle5", destChainId: thisChainId });
        XTypes.Msg memory xmsg;

        for (uint256 i = 0; i < xsub.msgs.length; i++) {
            xmsg = xsub.msgs[i];

            uint256 gasStart = gasleft();
            vm.chainId(xmsg.destChainId);
            portal.exec(_xheader(xmsg, xsub.blockHeader.sourceChainId), xmsg);
            uint256 gasUsed = gasStart - gasleft();

            console.log("exec single");
            console.log("offset: ", xmsg.offset);
            console.log("non-xmsg gas used: ", gasUsed - xmsg.gasLimit);
        }
    }

    function _testGasSubmitXBlock(string memory name) internal {
        _testGasSubmitXBlock(name, thisChainId);
    }

    function _testGasSubmitXBlock(string memory name, uint64 destChainId) internal {
        XTypes.Submission memory xsub = readXSubmission(name, destChainId, genesisValSetId);

        uint64 sourceChainId = xsub.blockHeader.sourceChainId;
        uint64 shardId = xsub.msgs[xsub.msgs.length - 1].shardId;
        uint64 expectedOffset = xsub.msgs[xsub.msgs.length - 1].offset;

        uint256 totalXMsgGasLimit;
        for (uint256 i = 0; i < xsub.msgs.length; i++) {
            totalXMsgGasLimit += xsub.msgs[i].gasLimit;
        }

        uint256 gasStart = gasleft();
        vm.chainId(destChainId);
        portal.xsubmit(xsub);
        uint256 gasUsed = gasStart - gasleft();

        console.log("xsubmit - ", name);
        console.log("num signatures: ", xsub.signatures.length);
        console.log("num xmsgs: ", xsub.msgs.length);
        console.log("non-xmsg gas used: ", gasUsed - totalXMsgGasLimit);
        console.log("non-xmsg gas per xmsg: ", (gasUsed - totalXMsgGasLimit) / xsub.msgs.length);

        assertEq(portal.inXMsgOffset(sourceChainId, shardId), expectedOffset);
        assertEq(portal.inXBlockOffset(sourceChainId, shardId), xsub.blockHeader.offset);
    }

    // @dev Helper to create a XBlock header for an xmsg
    function _xheader(XTypes.Msg memory xmsg, uint64 sourceChainId) internal pure returns (XTypes.BlockHeader memory) {
        return XTypes.BlockHeader({
            sourceChainId: sourceChainId,
            consensusChainId: omniCChainID,
            confLevel: uint8(xmsg.shardId),
            offset: 1,
            sourceBlockHeight: 100,
            sourceBlockHash: bytes32(0)
        });
    }
}
