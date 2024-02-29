// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import { IOmniAVS } from "./IOmniAVS.sol";

/**
 * @title IOmniEthRestaking
 * @dev Interface for OmniEthRestaking predeployed contract, that receives operator stake updates from OmniAVS
 *      TOOD: implement OmniEthRestaking (name TBD)
 */
interface IOmniEthRestaking {
    /// @dev Syncs operator state with OmniAVS. Only callable by XMsg from OmniAVS
    function sync(IOmniAVS.Operator[] calldata operators) external;
}
