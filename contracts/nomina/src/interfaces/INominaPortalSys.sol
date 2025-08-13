// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.30;

import { XTypes } from "src/libraries/XTypes.sol";

/**
 * @title INominaPortalSys
 * @notice Defines syscall functions internal to Nomina's cross-chain messaging protocol
 */
interface INominaPortalSys {
    /**
     * @notice Emitted when a new validator set is added
     * @param setId Validator set ID
     */
    event ValidatorSetAdded(uint64 indexed setId);

    /**
     * @notice Add a new validator set.
     * @dev Only callable via xcall from Nomina's consensus chain
     * @param valSetId      Validator set id
     * @param validators    Validator set
     */
    function addValidatorSet(uint64 valSetId, XTypes.Validator[] calldata validators) external;

    /**
     * @notice Set the network of supported chains & shards
     * @dev Only callable via xcall from Nomina's consensus chain
     * @param network_  The new network
     */
    function setNetwork(XTypes.Chain[] calldata network_) external;
}
