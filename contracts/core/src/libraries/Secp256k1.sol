// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { EllipticCurve } from "elliptic-curve-solidity/contracts/EllipticCurve.sol";

/**
 * @title Secp256k1
 * @dev Utility lib to validate compressed secp256k1 public keys
 */
library Secp256k1 {
    /**
     * @notice Curve parameter a
     */
    uint256 public constant AA = 0;

    /**
     * @notice Curve parameter b
     */
    uint256 public constant BB = 7;

    /**
     * @notice Prime field modulus
     */
    uint256 public constant PP = 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F;

    /**
     * @notice Compress a public key
     * @param x The x coordinate of the public key
     * @param y The y coordinate of the public key
     * @return compressedPubKey The compressed public key
     */
    function compressPublicKey(bytes32 x, bytes32 y) internal pure returns (bytes memory) {
        // Set prefix based on y coordinate's parity
        // 0x02 for even y, 0x03 for odd y
        uint8 prefix = uint8(y[31] & 0x01) + 0x02;

        // Concatenate prefix and x coordinate
        bytes memory compressedPubKey = abi.encodePacked(bytes1(prefix), x);
        return compressedPubKey;
    }

    /**
     * @notice Verify a secp256k1 public key is on the curve
     * @param x The x coordinate of the public key
     * @param y The y coordinate of the public key
     * @return True if the public key is valid, false otherwise
     */
    function verifyPubkey(bytes32 x, bytes32 y) internal pure returns (bool) {
        return EllipticCurve.isOnCurve(uint256(x), uint256(y), AA, BB, PP);
    }

    /**
     * @notice Validate a compressed secp256k1 public key
     * @param compressedPubKey The compressed public key to validate
     * @return True if the public key is valid, false otherwise
     */
    function verifyPubkey(bytes calldata compressedPubKey) internal pure returns (bool) {
        require(compressedPubKey.length == 33, "Staking: invalid pubkey length");
        require(compressedPubKey[0] == 0x02 || compressedPubKey[0] == 0x03, "Staking: invalid pubkey prefix");

        // Extract x coordinate
        uint256 x;
        assembly {
            let ptr := add(compressedPubKey.offset, 1)
            x := calldataload(ptr)
        }

        // Derive y coordinate
        uint256 y = EllipticCurve.deriveY(uint8(compressedPubKey[0]), x, AA, BB, PP);

        // Verify the derived point lies on the curve
        return EllipticCurve.isOnCurve(x, y, AA, BB, PP);
    }

    /**
     * @notice Convert public key coordinates to an Ethereum address
     * @param x The x coordinate of the public key
     * @param y The y coordinate of the public key
     * @return The Ethereum address corresponding to the public key
     */
    function pubkeyToAddress(bytes32 x, bytes32 y) internal pure returns (address) {
        bytes memory pubKey = new bytes(64);
        assembly {
            // Store x and y coordinates
            mstore(add(pubKey, 32), x)
            mstore(add(pubKey, 64), y)
        }

        // Hash the public key and keep last 20 bytes
        bytes32 hash = keccak256(pubKey);
        return address(uint160(uint256(hash)));
    }
}
