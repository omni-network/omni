// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { IMerkleDistributor } from "../interfaces/IMerkleDistributor.sol";
import { LibBitmap } from "solady/src/utils/LibBitmap.sol";
import { MerkleProofLib } from "solady/src/utils/MerkleProofLib.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";

contract MerkleDistributor is IMerkleDistributor {
    using LibBitmap for LibBitmap.Bitmap;
    using SafeTransferLib for address;

    error AlreadyClaimed();
    error InvalidProof();

    address public immutable override token;
    bytes32 public immutable override merkleRoot;

    // This is a packed array of booleans.
    LibBitmap.Bitmap internal claimedBitMap;

    constructor(address token_, bytes32 merkleRoot_) {
        token = token_;
        merkleRoot = merkleRoot_;
    }

    function isClaimed(uint256 index) public view override returns (bool) {
        return claimedBitMap.get(index);
    }

    function claim(uint256 index, address account, uint256 amount, bytes32[] calldata merkleProof)
        public
        virtual
        override
    {
        if (isClaimed(index)) revert AlreadyClaimed();

        // Verify the merkle proof.
        bytes32 node = keccak256(abi.encodePacked(index, account, amount));
        if (!MerkleProofLib.verifyCalldata(merkleProof, merkleRoot, node)) revert InvalidProof();

        // Mark it claimed and send the token.
        claimedBitMap.set(index);
        token.safeTransfer(account, amount);

        emit Claimed(index, account, amount);
    }
}
