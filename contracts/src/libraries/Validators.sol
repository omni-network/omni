// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

import { ECDSA } from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

/**
 * @title Validators
 * @dev Defines validator data types and quorum verification logic.
 */
library Validators {
    /// @dev Threshold scale factor, used when checking if an amount of power
    ///     exceeds a percentage of the total power, without loss of precision.
    uint16 private constant _THRESHOLD_SCALE_FACTOR = 1000;

    struct SigTuple {
        /// @dev Validator ethereum address
        address validatorAddr;
        /// @dev Validator signature over some digest; Ethereum 65 bytes [R || S || V] format.
        bytes signature;
    }

    struct Validator {
        /// @dev Validator ethereum address
        address addr;
        /// @dev Validator voting power
        uint64 power;
    }

    /**
     * @dev Verifies that the given percentage of the total power has signed the given digest.
     * @param digest Signed hash
     * @param sigs Signatures to verify
     * @param validators Maps validator addresses to their voting power
     * @param totalPower Total voting power
     * @param pct Percentage of total power required to reach quorum, 0-100
     */
    function verifyQuorum(
        bytes32 digest,
        SigTuple[] calldata sigs,
        mapping(address => uint64) storage validators,
        uint64 totalPower,
        uint8 pct
    ) internal view returns (bool) {
        uint64 power;
        SigTuple calldata sig;

        for (uint256 i = 0; i < sigs.length; i++) {
            sig = sigs[i];

            require(_isUnique(sig.validatorAddr, sigs, i + 1), "OmniPortal: duplicate validator");

            if (_verifySig(sig, digest)) power += validators[sig.validatorAddr];
            if (_exceedsThreshold(power, totalPower, pct)) return true;
        }

        return false;
    }

    /// @dev Verifies that the given SigTuple.sig is valid ECDSA signature over the given digest, for SigTuple.addr.
    function _verifySig(SigTuple calldata sig, bytes32 digest) internal pure returns (bool) {
        return ECDSA.recover(digest, sig.signature) == sig.validatorAddr;
    }

    /// @dev Verifies that the given power exceeds the given percentage of the total power.
    function _exceedsThreshold(uint64 power, uint64 totalPower, uint64 pct) private pure returns (bool) {
        return (power * _THRESHOLD_SCALE_FACTOR) >= (totalPower * _THRESHOLD_SCALE_FACTOR * pct) / 100;
    }

    /// @dev Verifies that the given address is unique in the given array of SigTuples, starting at the given index.
    ///      Starting index is used to avoid checking the same address twice.
    function _isUnique(address addr, SigTuple[] calldata sigs, uint256 start) private pure returns (bool) {
        for (uint256 i = start; i < sigs.length; i++) {
            if (sigs[i].validatorAddr == addr) return false;
        }
        return true;
    }
}
