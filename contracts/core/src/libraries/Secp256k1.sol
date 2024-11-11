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
     * @notice Prefix added to Ethereum signed messages
     */
    string internal constant ETH_SIGN_PREFIX = "\x19Ethereum Signed Message:\n32";

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
    function publicKeyToAddress(bytes32 x, bytes32 y) internal pure returns (address) {
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

    /**
     * @notice Get the Ethereum signed message hash of a public key
     * @param x The x coordinate of the public key
     * @param y The y coordinate of the public key
     * @return The Ethereum signed message hash
     */
    function getEthSignedMessageHash(bytes32 x, bytes32 y) internal pure returns (bytes32) {
        // Use full uncompressed public key as the message
        bytes memory message = abi.encodePacked(hex"04", x, y);
        bytes32 ethSignedMessageHash =
            keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", keccak256(message)));
        return ethSignedMessageHash;
    }

    /**
     * @notice Verify a signature against a public key
     * @param x The x coordinate of the public key
     * @param y The y coordinate of the public key
     * @param signature The 65-byte signature (r,s,v)
     * @return True if signature is valid, false otherwise
     */
    function verifySignature(bytes32 x, bytes32 y, bytes calldata signature) internal pure returns (bool) {
        require(signature.length == 65, "Secp256k1: invalid signature length");
        bytes32 ethSignedMessageHash = getEthSignedMessageHash(x, y);

        // Extract signature components
        bytes32 r;
        bytes32 s;
        uint8 v;
        assembly {
            r := calldataload(signature.offset)
            s := calldataload(add(signature.offset, 32))
            v := byte(0, calldataload(add(signature.offset, 64)))
        }

        // v needs to be 0 or 1 for secp256k1 (not 27/28 like Ethereum)
        require(v == 0 || v == 1, "Secp256k1: invalid v value");

        // Add 27 to make it compatible with ecrecover
        v += 27;

        // Recover the signer's Ethereum address
        address recovered = ecrecover(ethSignedMessageHash, v, r, s);
        require(recovered != address(0), "Secp256k1: ecrecover failed");

        // Convert the provided public key to an Ethereum address
        address pubKeyAddress = publicKeyToAddress(x, y);

        // Compare the addresses
        return recovered == pubKeyAddress;
    }
}
