// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { MerkleProof } from "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";
import { XTypes } from "./XTypes.sol";

/**
 * @title XBlockMerkleProof
 * @dev Library for verifying XBlock merkle proofs
 */
library XBlockMerkleProof {
    /// @dev Domain separation tag for XBlockHeaders, prepended to leaves before hashing and signing.
    uint8 internal constant DST_XBLOCK_HEADER = 1;

    /// @dev Domain separation tag for XMsgs, prepended to leaves before hashing and signing.
    uint8 internal constant DST_XMSG = 2;

    /**
     * @notice Verifies that the provided xmsgs & multi proof produce an xmsg merkle root that, when
     *         combined with the xblock header, produces the provided root.
     * @param root          The root of the nested xblock merkle tree (xblock header + xmsg merkle root).
     * @param blockHeader   Xblock header.
     * @param msgs          Xmsgs to verify.
     * @param msgProof      Xmsg merkle proof.
     * @param msgProofFlags Xmsg merkle proof flags.
     * @return              True if the msgs, msg proof & block header are valid, agsinst the provided root.
     */
    function verify(
        bytes32 root,
        XTypes.BlockHeader calldata blockHeader,
        XTypes.Msg[] calldata msgs,
        bytes32[] calldata msgProof,
        bool[] calldata msgProofFlags
    ) internal pure returns (bool) {
        bytes32[] memory rootProof = new bytes32[](1);
        rootProof[0] = MerkleProof.processMultiProofCalldata(msgProof, msgProofFlags, _msgLeaves(msgs));
        return MerkleProof.verify(rootProof, root, _blockHeaderLeaf(blockHeader));
    }

    /// @dev Convert xmsgs to leaf hashes
    function _msgLeaves(XTypes.Msg[] calldata msgs) private pure returns (bytes32[] memory) {
        bytes32[] memory leaves = new bytes32[](msgs.length);

        for (uint256 i = 0; i < msgs.length; i++) {
            leaves[i] = _leafHash(DST_XMSG, abi.encode(msgs[i]));
        }

        return leaves;
    }

    /// @dev Convert xblock header to leaf hash
    function _blockHeaderLeaf(XTypes.BlockHeader calldata blockHeader) private pure returns (bytes32) {
        return _leafHash(DST_XBLOCK_HEADER, abi.encode(blockHeader));
    }

    /// @dev Double hash leaves, as recommended by OpenZeppelin, to prevent second preimage attacks
    ///      Leaves must be double hashed in tree / proof construction
    ///      Callers must specify the domain separation tag of the leaf, which will be hashed in
    function _leafHash(uint8 dst, bytes memory leaf) private pure returns (bytes32) {
        return keccak256(bytes.concat(keccak256(abi.encodePacked(dst, leaf))));
    }
}
