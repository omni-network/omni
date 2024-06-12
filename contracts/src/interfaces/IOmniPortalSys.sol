// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import { XTypes } from "../libraries/XTypes.sol";

/**
 * @title IOmniPortalSys
 * @notice Defines syscall functions internal to Omni's cross-chain messaging protocol
 */
interface IOmniPortalSys {
    /**
     * @notice Emitted when a new validator set is added
     * @param setId Validator set ID
     */
    event ValidatorSetAdded(uint64 indexed setId);

    /**
     * @notice Add a new validator set.
     * @dev Only callable via xcall from Omni's consensus chain
     * @param valSetId      Validator set id
     * @param validators    Validator set
     */
    function addValidatorSet(uint64 valSetId, XTypes.Validator[] calldata validators) external;

    /**
     * @notice Initialize a source chain's in stream validator set
     * @dev Only callable from xregistry
     * @param shards        Shards supported by the source chain
     */
    function setShards(uint64[] calldata shards) external;
}
