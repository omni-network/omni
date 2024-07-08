// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import { IOmniAVS } from "avs/interfaces/IOmniAVS.sol";

/**
 * @title IEthStakeInbox
 * @dev Interface for EthStakeInbox predeployed contract, that receives operator stake updates from OmniAVS
 *      TOOD: implement EthStakeInbox (name TBD)
 */
interface IEthStakeInbox {
    /// @dev Syncs operator state with OmniAVS. Only callable by XMsg from OmniAVS
    function sync(IOmniAVS.Operator[] calldata operators) external;
}
