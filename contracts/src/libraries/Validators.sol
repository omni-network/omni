// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

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
     * @param sigs Signatures to verify, must be sorted by validator address
     * @param validators Maps validator addresses to their voting power
     * @param totalPower Total voting power
     * @param qNumerator Numerator of the quorum threshold. Ex: 2/3 -> 2
     * @param qDenominator Denominator of the quorum threshold. Ex: 2/3 -> 3
     */
    function verifyQuorum(
        bytes32 digest,
        SigTuple[] calldata sigs,
        mapping(address => uint64) storage validators,
        uint64 totalPower,
        uint8 qNumerator,
        uint8 qDenominator
    ) internal view returns (bool) {
        uint64 votedPower;
        SigTuple calldata sig;

        for (uint256 i = 0; i < sigs.length; i++) {
            sig = sigs[i];

            if (i > 0) {
                SigTuple memory prev = sigs[i - 1];
                require(sig.validatorAddr != prev.validatorAddr, "OmniPortal: duplicate validator");
                require(sig.validatorAddr > prev.validatorAddr, "OmniPortal: sigs not sorted");
            }

            if (_isValidSig(sig, digest)) votedPower += validators[sig.validatorAddr];
            if (_isQuorum(votedPower, totalPower, qNumerator, qDenominator)) return true;
        }

        return false;
    }

    /// @dev True if SigTuple.sig is a valid ECDSA signature over the given digest for SigTuple.addr, else false.
    function _isValidSig(SigTuple calldata sig, bytes32 digest) internal pure returns (bool) {
        return ECDSA.recover(digest, sig.signature) == sig.validatorAddr;
    }

    /// @dev True if votedPower exceeds the quorum threshold of numerator/denominator, else false.
    function _isQuorum(uint64 votedPower, uint64 totalPower, uint8 numerator, uint8 denominator)
        private
        pure
        returns (bool)
    {
        return votedPower > totalPower * numerator / denominator;
    }
}
