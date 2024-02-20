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
     * @notice Returns list of operators that are allowed to register.
     */
    function allowlist() external view returns (address[] memory);

    /**
     * @notice Initialize the Omni AVS admin contract.
     * @param owner The intiial owner of the contract
     * @param omni The Omni portal contract
     * @param omniChainId The Omni chain id
     * @param minimumOperatorStake The minimum operator stake, not including delegations
     * @param maxOperatorCount The maximum operator count
     * @param allowlist The initial allowlist
     * @param strategyParams List of accepted strategies and their multipliers
     */
    function initialize(
        address owner,
        IOmniPortal omni,
        uint64 omniChainId,
        uint96 minimumOperatorStake,
        uint32 maxOperatorCount,
        address[] calldata allowlist,
        IOmniAVS.StrategyParams[] calldata strategyParams
    ) external;

    /**
     * @notice Set the Omni portal contract.
     * @dev Only the owner can call this function.
     * @param portal The Omni portal contract
     */
    function setOmniPortal(IOmniPortal portal) external;

    /**
     * @notice Set the Omni chain id.
     * @dev Only the owner can call this function.
     * @param chainID The Omni chain id
     */
    function setOmniChainId(uint64 chainID) external;

    /**
     * @notice Set the strategy parameters.
     * @dev Only the owner can call this function.
     * @param params The strategy parameters
     */
    function setStrategyParams(IOmniAVS.StrategyParams[] calldata params) external;

    /**
     * @notice Set the minimum operator stake.
     * @dev Only the owner can call this function.
     * @param stake The minimum operator stake, not including delegations
     */
    function setMinimumOperatorStake(uint96 stake) external;

    /**
     * @notice Set the maximum operator count.
     * @dev Only the owner can call this function.
     * @param count The maximum operator count
     */
    function setMaxOperatorCount(uint32 count) external;

    /**
     * @notice Set the xcall gas limits.
     * @dev Only the owner can call this function.
     * @param base The base xcall gas limit
     * @param perValidator The per-validator additional xcall gas limit
     */
    function setXcallGasLimits(uint256 base, uint256 perValidator) external;

    /**
     * @notice Set the allowlist.
     * @dev Only the owner can call this function.
     * @param allowlist The allowlist
     */
    function setAllowlist(address[] calldata allowlist) external;

    /**
     * @notice Add an operator to the allowlist.
     * @dev Only the owner can call this function.
     * @param operator The operator to add
     */
    function addToAllowlist(address operator) external;

    /**
     * @notice Remove an operator from the allowlist.
     * @dev Only the owner can call this function.
     * @param operator The operator to remove
     */
    function removeFromAllowlist(address operator) external;
}
