// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { Quorum } from "src/libraries/Quorum.sol";
import { XTypes } from "src/libraries/XTypes.sol";

/**
 * @title QuorumVerifier
 * @dev Helper contract to allow us to test Quorum.verify(...) without modifying it's function signature,
 *      which expects storage and calldata arguments.
 */
contract QuorumVerifier {
    /// @dev map valset id to total power
    mapping(uint64 => uint64) totalPower;

    /// @dev map valset id to addr to power
    mapping(uint64 => mapping(address => uint64)) valset;

    /**
     * @dev Verify a quorum of signatures, forwards call to Quorum.verify(...)
     */
    function verify(
        bytes32 digest,
        XTypes.SigTuple[] calldata sigs,
        uint64 valSetId,
        uint8 qNumerator,
        uint8 qDenominator
    ) public view returns (bool) {
        return Quorum.verify(digest, sigs, valset[valSetId], totalPower[valSetId], qNumerator, qDenominator);
    }

    /**
     * @dev Helper to allow caller to set validator sets.
     */
    function setValset(uint64 valSetId, XTypes.Validator[] calldata vals) public {
        uint64 power = 0;

        for (uint256 i = 0; i < vals.length; i++) {
            valset[valSetId][vals[i].addr] = vals[i].power;
            power += vals[i].power;
        }

        totalPower[valSetId] = power;
    }

    /**
     * @dev Helper to allow caller to update validator power.
     */
    function updateValidatorPower(uint64 valSetId, address addr, uint64 power) public {
        totalPower[valSetId] -= valset[valSetId][addr];
        valset[valSetId][addr] = power;
        totalPower[valSetId] += power;
    }
}
