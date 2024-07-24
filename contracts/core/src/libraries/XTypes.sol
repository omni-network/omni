// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

/**
 * @title XTypes
 * @dev Defines xchain types, core to Omni's xchain messaging protocol. These
 *      types mirror those defined in omni/lib/xchain/types.go.
 */
library XTypes {
    /**
     * @notice A cross chain message - the product of an xcall. This matches the XMsg type used
     *        throughout Omni's cross-chain messaging protocol. Msg is used to construct and verify
     *        XSubmission merkle trees / proofs.
     * @custom:field destChainId    Chain ID of the destination chain
     * @custom:field shardId        Shard ID of the XStream (first byte is the confirmation level)
     * @custom:field offset         Monotonically incremented offset of Msg in source -> dest Stream
     * @custom:field sender         msg.sender of xcall on source chain
     * @custom:field to             Target address to call on destination chain
     * @custom:field data           Data to provide to call on destination chain
     * @custom:field gasLimit       Gas limit to use for call execution on destination chain
     */
    struct Msg {
        uint64 destChainId;
        uint64 shardId;
        uint64 offset;
        address sender;
        address to;
        bytes data;
        uint64 gasLimit;
    }

    /**
     * @notice Msg context exposed during its execution to consuming xapps.
     * @custom:field sourceChainId  Chain ID of the source chain
     * @custom:field sender         msg.sender of xcall on source chain
     */
    struct MsgContext {
        uint64 sourceChainId;
        address sender;
    }

    /**
     * @notice BlockHeader of an XBlock.
     * @custom:field sourceChainId      Chain ID of the source chain
     * @custom:field sourceChainId      Chain ID of the Omni consensus chain
     * @custom:field confLevel          Confirmation level of the cross chain block
     * @custom:field offset             Offset of the cross chain block
     * @custom:field sourceBlockHeight  Height of the source chain block
     * @custom:field sourceBlockHash    Hash of the source chain block
     */
    struct BlockHeader {
        uint64 sourceChainId;
        uint64 consensusChainId;
        uint8 confLevel;
        uint64 offset;
        uint64 sourceBlockHeight;
        bytes32 sourceBlockHash;
    }

    /**
     * @notice The required parameters to submit xmsgs to an OmniPortal. Constructed by the relayer
     *         by watching Omni's consensus chain, and source chain blocks.
     * @custom:field attestationRoot  Merkle root of xchain block (XBlockRoot), attested to and signed by validators
     * @custom:field validatorSetId   Unique identifier of the validator set that attested to this root
     * @custom:field blockHeader      Block header, identifies xchain block
     * @custom:field msgs             Messages to execute
     * @custom:field proof            Multi proof of block header and messages, proven against attestationRoot
     * @custom:field proofFlags       Multi proof flags
     * @custom:field signatures       Array of validator signatures of the attestationRoot, and their public keys
     */
    struct Submission {
        bytes32 attestationRoot;
        uint64 validatorSetId;
        BlockHeader blockHeader;
        Msg[] msgs;
        bytes32[] proof;
        bool[] proofFlags;
        SigTuple[] signatures;
    }

    /**
     * @notice A tuple of a validator's ethereum address and signature over some digest.
     * @custom:field validatorAddr  Validator ethereum address
     * @custom:field signature      Validator signature over some digest; Ethereum 65 bytes [R || S || V] format.
     */
    struct SigTuple {
        address validatorAddr;
        bytes signature;
    }

    /**
     * @notice An Omni validator, specified by their etheruem address and voting power.
     * @custom:field addr   Validator ethereum address
     * @custom:field power  Validator voting power
     */
    struct Validator {
        address addr;
        uint64 power;
    }

    /**
     * @notice A chain in the "omni network" specified by its chain ID and supported shards.
     * @custom:field chainId  Chain ID
     * @custom:field shards   Supported shards
     */
    struct Chain {
        uint64 chainId;
        uint64[] shards;
    }
}
