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

    function test_xsubmi_addValidator_succeeds() public {
        _testGasSubmitXBlock("addValSet2", broadcastChainId);
    }

    function test_singleExec() public {
        XTypes.Submission memory xsub = readXSubmission("guzzle5", portal.chainId(), genesisValSetId);
        XTypes.Msg memory xmsg;

        for (uint256 i = 0; i < xsub.msgs.length; i++) {
            xmsg = xsub.msgs[i];

            uint256 gasStart = gasleft();
            portal.exec(xmsg);
            uint256 gasUsed = gasStart - gasleft();

            console.log("exec single");
            console.log("offset: ", xmsg.streamOffset);
            console.log("non-xmsg gas used: ", gasUsed - xmsg.gasLimit);
        }
    }

    function _testGasSubmitXBlock(string memory name) internal {
        _testGasSubmitXBlock(name, portal.chainId());
    }

    function _testGasSubmitXBlock(string memory name, uint64 destChainId) internal {
        XTypes.Submission memory xsub = readXSubmission(name, destChainId, genesisValSetId);

        uint64 sourceChainId = xsub.blockHeader.sourceChainId;
        uint64 expectedOffset = xsub.msgs[xsub.msgs.length - 1].streamOffset;

        uint256 totalXMsgGasLimit;
        for (uint256 i = 0; i < xsub.msgs.length; i++) {
            totalXMsgGasLimit += xsub.msgs[i].gasLimit;
        }

        uint256 gasStart = gasleft();
        portal.xsubmit(xsub);
        uint256 gasUsed = gasStart - gasleft();

        console.log("xsubmit - ", name);
        console.log("num signatures: ", xsub.signatures.length);
        console.log("num xmsgs: ", xsub.msgs.length);
        console.log("non-xmsg gas used: ", gasUsed - totalXMsgGasLimit);
        console.log("non-xmsg gas per xmsg: ", (gasUsed - totalXMsgGasLimit) / xsub.msgs.length);

        assertEq(portal.inXStreamOffset(sourceChainId), expectedOffset);
        assertEq(portal.inXStreamBlockHeight(sourceChainId), xsub.blockHeader.blockHeight);
    }
}
