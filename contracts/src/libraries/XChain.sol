// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

/**
 * @title XChain
 * @dev Defines xchain types, core to Omni's xchain messaging protocol. These
 *      types mirror those defined in omni/lib/xchain/types.go.
 */
library XChain {
    struct Msg {
        /// @dev Chain ID of the source chain
        uint64 sourceChainId;
        /// @dev Chain ID of the destination chain
        uint64 destChainId;
        /// @dev Monotonically incremented offset of Msg in source -> dest Stream
        uint64 streamOffset;
        /// @dev msg.sender of xcall on source chain
        address sender;
        /// @dev Target address to call on destination chain
        address to;
        /// @dev Data to provide to call on destination chain
        bytes data;
        /// @dev Gas limit to use for call execution on destination chain
        uint64 gasLimit;
        /// @dev Hash of the source chain transaction that emitted the message
        bytes32 txHash;
    }

    struct BlockHeader {
        /// @dev Chain ID of the source chain
        uint64 sourceChainId;
        /// @dev Height of the source chain block
        uint64 blockHeight;
        /// @dev Hash of the source chain block
        bytes32 blockHash;
    }

    struct SigTuple {
        /// @dev Validator public key - 33 bytes compressed secp256k1
        bytes validatorPubKey;
        /// @dev Signature of the attestationRoot
        bytes signature;
    }

    struct Submission {
        /// @dev Merkle root of xchain block, attested to and signed by Omni validators
        bytes32 attestationRoot;
        /// @dev Block header, proven against attestationRoot, identifies xchain block
        BlockHeader blockHeader;
        /// @dev Array of xchain messages in the block
        Msg[] msgs;
        /// @dev Multi proof of block header and messages, proven against attestationRoot
        bytes32[] proof;
        /// @dev Multi proof flags
        bool[] ProofFlags;
        /// @dev Array of Omni validator signatures of the attestationRoot
        SigTuple[] signatures;
    }
}
