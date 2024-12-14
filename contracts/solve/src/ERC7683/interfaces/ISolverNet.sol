// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { IERC7683 } from "./IERC7683.sol";

interface ISolverNet is IERC7683 {
    /**
     * @notice Details of a call to be executed on another chain.
     * @param target      Address of the target contract on the destination chain.
     * @param value       Amount of native currency to send with the call.
     * @param callData    Encoded data to be sent with the call.
     */
    struct Call {
        bytes32 target;
        uint256 value;
        bytes callData;
    }

    struct FillOriginData {
        uint64 srcChainId;
        uint64 destChainId;
        Call[] calls;
        Output[] prereqs;
    }
}
