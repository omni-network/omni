// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { MerkleProof } from "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";
import { XTypes } from "./XTypes.sol";

/**
 * @title XBlockMerkleProof
 * @dev Library for verifying XBlock merkle proofs
 */
library XBlockMerkleProof {
    /**
     * @notice Verifies a multi merkle proof for the provided block header and messages, against the provided root.
     *      Msgs order must match the order used to construct the merkle proof.
     * @param root          The root of the xblock merkle tree, generally XSubmission.attestationRoot.
     * @param blockHeader   The xblock header.
     * @param msgs          The xmsgs to verify.
     * @param proof         The merkle proof.
     * @param proofFlags    The merkle proof flags.
     * @return              True if the proof is valid, false otherwise.
     */
    function verify(
        bytes32 root,
        XTypes.BlockHeader calldata blockHeader,
        XTypes.Msg[] calldata msgs,
        bytes32[] calldata proof,
        bool[] calldata proofFlags
    ) internal pure returns (bool) {
        return MerkleProof.multiProofVerify(proof, proofFlags, root, _leaves(blockHeader, msgs));
    }

    /// @dev Convert block header and msgs to leaf hashes
    function _leaves(XTypes.BlockHeader calldata blockHeader, XTypes.Msg[] calldata msgs)
        private
        pure
        returns (bytes32[] memory)
    {
        bytes32[] memory leaves = new bytes32[](msgs.length + 1);

        leaves[0] = _leafHash(abi.encode(blockHeader));
        for (uint256 i = 0; i < msgs.length; i++) {
            leaves[i + 1] = _leafHash(abi.encode(msgs[i]));
        }

        return leaves;
    }

    /// @dev Double hash leaves, as recommended by OpenZeppelin, to prevent second preimage attacks
    ///      Leaves must be double hashed in tree / proof construction
    function _leafHash(bytes memory leaf) private pure returns (bytes32) {
        return keccak256(bytes.concat(keccak256(leaf)));
    }
}
