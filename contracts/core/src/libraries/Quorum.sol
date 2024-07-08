// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { ECDSA } from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import { XTypes } from "./XTypes.sol";

/**
 * @title Quorom
 * @dev Defines quorum verification logic.
 */
library Quorum {
    /**
     * @notice Verifies that the given percentage of the total power has signed the given digest.
     * @param digest        Signed hash
     * @param sigs          Signatures to verify, must be sorted by validator address
     * @param validators    Maps validator addresses to their voting power
     * @param totalPower    Total voting power
     * @param qNumerator    Numerator of the quorum threshold. Ex: 2/3 -> 2
     * @param qDenominator  Denominator of the quorum threshold. Ex: 2/3 -> 3
     */
    function verify(
        bytes32 digest,
        XTypes.SigTuple[] calldata sigs,
        mapping(address => uint64) storage validators,
        uint64 totalPower,
        uint8 qNumerator,
        uint8 qDenominator
    ) internal view returns (bool) {
        uint64 votedPower;
        XTypes.SigTuple calldata sig;

        for (uint256 i = 0; i < sigs.length; i++) {
            sig = sigs[i];

            if (i > 0) {
                XTypes.SigTuple memory prev = sigs[i - 1];
                require(sig.validatorAddr != prev.validatorAddr, "Quorum: duplicate validator");
                require(sig.validatorAddr > prev.validatorAddr, "Quorum: sigs not sorted");
            }

            require(_isValidSig(sig, digest), "Quorum: invalid signature");

            votedPower += validators[sig.validatorAddr];

            if (_isQuorum(votedPower, totalPower, qNumerator, qDenominator)) return true;
        }

        return false;
    }

    /// @dev True if SigTuple.sig is a valid ECDSA signature over the given digest for SigTuple.addr, else false.
    function _isValidSig(XTypes.SigTuple calldata sig, bytes32 digest) internal pure returns (bool) {
        return ECDSA.recover(digest, sig.signature) == sig.validatorAddr;
    }

    /// @dev True if votedPower exceeds the quorum threshold of numerator/denominator, else false.
    function _isQuorum(uint64 votedPower, uint64 totalPower, uint8 numerator, uint8 denominator)
        private
        pure
        returns (bool)
    {
        return votedPower > (totalPower * numerator) / denominator;
    }
}
