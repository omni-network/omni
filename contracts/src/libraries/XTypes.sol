// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import { Validators } from "./Validators.sol";

/**
 * @title XTypes
 * @dev Defines xchain types, core to Omni's xchain messaging protocol. These
 *      types mirror those defined in omni/lib/xchain/types.go.
 */
library XTypes {
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
    }

    struct BlockHeader {
        /// @dev Chain ID of the source chain
        uint64 sourceChainId;
        /// @dev Height of the source chain block
        uint64 blockHeight;
        /// @dev Hash of the source chain block
        bytes32 blockHash;
    }

    struct Submission {
        /// @dev Merkle root of xchain block (XBlockRoot), attested to and signed by validators
        bytes32 attestationRoot;
        /// @dev Unique identifier of the validator set that attested to this root
        uint64 validatorSetId;
        /// @dev Block header, identifies xchain block
        BlockHeader blockHeader;
        /// @dev Messages to execute
        Msg[] msgs;
        /// @dev Multi proof of block header and messages, proven against attestationRoot
        bytes32[] proof;
        /// @dev Multi proof flags
        bool[] proofFlags;
        /// @dev Array of validator signatures of the attestationRoot, and their public keys
        Validators.SigTuple[] signatures;
    }

    /// @dev Zero value for Msg
    function zeroMsg() internal pure returns (Msg memory) {
        return Msg(0, 0, 0, address(0), address(0), "", 0);
    }
}
