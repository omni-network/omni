// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

import { ECDSA } from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

/**
 * @title Validators
 * @dev Defines validator data types and quorum verification logic.
 */
library Validators {
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
     * @param thresholdNumerator Numerator of the quorum threshold. Ex: 2/3 -> 2
     * @param thresholdDenominator Denominator of the quorum threshold. Ex: 2/3 -> 3
     */
    function verifyQuorum(
        bytes32 digest,
        SigTuple[] calldata sigs,
        mapping(address => uint64) storage validators,
        uint64 totalPower,
        uint8 thresholdNumerator,
        uint8 thresholdDenominator
    ) internal view returns (bool) {
        uint64 votedPower;
        SigTuple calldata sig;

        for (uint256 i = 0; i < sigs.length; i++) {
            sig = sigs[i];

            require(_isUnique(sig.validatorAddr, sigs, i + 1), "OmniPortal: duplicate validator");

            if (_verifySig(sig, digest)) votedPower += validators[sig.validatorAddr];
            if (_exceedsThreshold(votedPower, totalPower, thresholdNumerator, thresholdDenominator)) return true;
        }

        return false;
    }

    /// @dev Verifies that the given SigTuple.sig is valid ECDSA signature over the given digest, for SigTuple.addr.
    function _verifySig(SigTuple calldata sig, bytes32 digest) internal pure returns (bool) {
        return ECDSA.recover(digest, sig.signature) == sig.validatorAddr;
    }

    /// @dev Verifies that the given power exceeds the given percentage of the total power.
    function _exceedsThreshold(uint64 votedPower, uint64 totalPower, uint8 numerator, uint8 denominator)
        private
        pure
        returns (bool)
    {
        return votedPower > totalPower * numerator / denominator;
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
