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
     * @param owner The intiial owner of the contract
     * @param omni The Omni portal contract
     * @param omniChainId The Omni chain id
     * @param minimumOperatorStake The minimum operator stake, not including delegations
     * @param maxOperatorCount The maximum operator count
     * @param strategyParams List of accepted strategies and their multipliers
     */
    function initialize(
        address owner,
        IOmniPortal omni,
        uint64 omniChainId,
        uint96 minimumOperatorStake,
        uint32 maxOperatorCount,
        IOmniAVS.StrategyParams[] calldata strategyParams
    ) external;

    /**
     * @notice Set the Omni portal contract.
     * @dev Only the owner can call this function.
     * @param omni The Omni portal contract
     */
    function setOmniPortal(IOmniPortal omni) external;

    /**
     * @notice Set the Omni chain id.
     * @dev Only the owner can call this function.
     * @param omniChainId The Omni chain id
     */
    function setOmniChainId(uint64 omniChainId) external;

    /**
     * @notice Set the strategy parameters.
     * @dev Only the owner can call this function.
     * @param strategyParams The strategy parameters
     */
    function setStrategyParams(IOmniAVS.StrategyParams[] calldata strategyParams) external;

    /**
     * @notice Sets the metadata URI for the AVS
     * @dev Matches eigenlayer-middleware ServiceManagerBase.setMetadataURI
     * @param metadataURI is the metadata URI for the AVS
     * @dev only callable by the owner
     */
    function setMetadataURI(string memory metadataURI) external;
}
