// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import { IOmniPortal } from "./IOmniPortal.sol";
import { IOmniAVS } from "./IOmniAVS.sol";

/**
 * @title OmniAVSAdmin
 * @notice Omni AVS admin internface.
 */
interface IOmniAVSAdmin {
    /**
     * @notice Initialize the Omni AVS admin contract.
     * @param owner_ The intiial owner of the contract
     * @param omni_ The Omni portal contract
     * @param omniChainId_ The Omni chain id
     */
    function initialize(
        address owner_,
        IOmniPortal omni_,
        uint64 omniChainId_,
        uint96 minimumStake_,
        uint32 maxOperatorCount_,
        IOmniAVS.StrategyParams[] calldata strategyParams_
    ) external;

    /**
     * @notice Set the Omni portal contract.
     * @dev Only the owner can call this function.
     * @param omni_ The Omni portal contract
     */
    function setOmniPortal(IOmniPortal omni_) external;

    /**
     * @notice Set the Omni chain id.
     * @dev Only the owner can call this function.
     * @param omniChainId_ The Omni chain id
     */
    function setOmniChainId(uint64 omniChainId_) external;

    /**
     * @notice Set the strategy parameters.
     * @dev Only the owner can call this function.
     * @param strategyParams_ The strategy parameters
     */
    function setStrategyParams(IOmniAVS.StrategyParams[] calldata strategyParams_) external;
}
