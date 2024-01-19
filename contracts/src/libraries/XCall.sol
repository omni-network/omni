// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

import { Bytes } from "./Bytes.sol";
import { XTypes } from "./XTypes.sol";

/**
 * @title XCall
 * @dev Defines logic for executing an XMsg on its destination chain, and capturing the XReceipt
 */
library XCall {
    struct ExecOpts {
        /// @dev Maximum amount of return data to track from xmsg execution in its receipt
        uint64 maxReturnDataSize;
        /// @dev Error message to return if an xmsg runs out of gas
        string outOfGasErrorMsg;
    }

    /**
     * @notice Execute an XMsg
     * @param xmsg XMsg to execute
     * @param opts Execution options
     * @return receipt XReceipt of xmsg execution
     */
    function exec(XTypes.Msg memory xmsg, address relayer, ExecOpts memory opts)
        internal
        returns (XTypes.Receipt memory)
    {
        // execute xmsg, tracking gas used
        uint256 gasUsed = gasleft();
        (bool success, bytes memory returnData) = xmsg.to.call{ gas: xmsg.gasLimit }(xmsg.data); // solhint-disable-line avoid-low-level-calls
        gasUsed = gasUsed - gasleft();

        // if necessary, trim returnData to maxReturnDataSize
        if (returnData.length > opts.maxReturnDataSize) {
            returnData = Bytes.slice(returnData, 0, opts.maxReturnDataSize);
        }

        // if an xmsg runs out of gas, returnData will be empty, so we add a useful error message manually
        if (gasUsed >= xmsg.gasLimit) {
            returnData = abi.encodeWithSignature("Error(string)", opts.outOfGasErrorMsg);
        }

        return XTypes.Receipt(
            xmsg.sourceChainId, xmsg.destChainId, xmsg.streamOffset, gasUsed, relayer, success, returnData
        );
    }
}
