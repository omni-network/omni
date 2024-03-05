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
     * @notice Emitted when an operator is added to the allowlist.
     * @param operator The operator
     */
    event OperatorAllowed(address operator);

    /**
     * @notice Emitted when an operator is removed from the allowlist.
     * @param operator The operator
     */
    event OperatorDisallowed(address operator);

    /**
     * @notice Emitted when the allowlist is enabled.
     */
    event AllowlistEnabled();

    /**
     * @notice Emitted when the allowlist is disabled.
     */
    event AllowlistDisabled();

    /**
     * @notice Initialize the Omni AVS admin contract.
     * @param owner             Intiial contract owner
     * @param omni              Omni portal contract
     * @param omniChainId       Omni chain id
     * @param ethStakeInbox     EthStakeInbox contract address, on Omni
     * @param strategyParams    List of accepted strategies and their multipliers
     */
    function initialize(
        address owner,
        IOmniPortal omni,
        uint64 omniChainId,
        address ethStakeInbox,
        IOmniAVS.StrategyParam[] calldata strategyParams
    ) external;

    /**
     * @notice Set the Omni portal contract.
     * @param portal The Omni portal contract
     */
    function setOmniPortal(IOmniPortal portal) external;

    /**
     * @notice Set the Omni chain id.
     * @param chainID The Omni chain id
     */
    function setOmniChainId(uint64 chainID) external;

    /**
     * @notice Set the EthStakeInbox contract address.
     * @param inbox The EthStakeInbox contract address
     */
    function setEthStakeInbox(address inbox) external;

    /**
     * @notice Set the strategy parameters.
     * @param params The strategy parameters
     */
    function setStrategyParams(IOmniAVS.StrategyParam[] calldata params) external;

    /**
     * @notice Set the xcall gas limits.
     * @param base          The base xcall gas limit
     * @param perValidator  The per-validator additional xcall gas limit
     */
    function setXCallGasLimits(uint64 base, uint64 perValidator) external;

    /**
     * @notice Returns true if the operator is in the allowlist.
     * @param operator The operator to check
     */
    function isInAllowlist(address operator) external view returns (bool);

    /**
     * @notice Add an operator to the allowlist.
     * @param operator The operator to add
     */
    function addToAllowlist(address operator) external;

    /**
     * @notice Remove an operator from the allowlist.
     * @param operator The operator to remove
     */
    function removeFromAllowlist(address operator) external;

    /**
     * @notice Enable the allowlist.
     */
    function enableAllowlist() external;

    /**
     * @notice Disable the allowlist.
     */
    function disableAllowlist() external;

    /**
     * @notice Pause the contract.
     */
    function pause() external;

    /**
     * @notice Unpause the contract.
     */
    function unpause() external;
}
