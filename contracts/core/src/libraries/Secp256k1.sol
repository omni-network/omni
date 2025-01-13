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
     * @notice Returns 33 byte compressed public key for the given x, y coordinates
     * @param x The x coordinate of the public key
     * @param y The y coordinate of the public key
     */
    function compress(uint256 x, uint256 y) internal pure returns (bytes memory) {
        bytes32 _y = bytes32(y);

        // Set prefix based on y coordinate's parity
        // 0x02 for even y, 0x03 for odd y
        uint8 prefix = uint8(_y[31] & 0x01) + 0x02;

        // Concatenate prefix and x coordinate
        return abi.encodePacked(bytes1(prefix), x);
    }

    /**
     * @notice Returns x,y coordinates of an compressed public key
     *         Reverts if the public key is invalid
     * @param pubkey The 33 byte compressed public key
     */
    function decompress(bytes calldata pubkey) internal pure returns (uint256, uint256) {
        require(pubkey.length == 33, "Secp256k1: pubkey not 33 bytes");
        require(pubkey[0] == 0x02 || pubkey[0] == 0x03, "Secp256k1: invalid pubkey prefix");

        // Extract x coordinate
        uint256 x;
        assembly {
            let ptr := add(pubkey.offset, 1)
            x := calldataload(ptr)
        }

        // Derive y coordinate
        uint256 y = EllipticCurve.deriveY(uint8(pubkey[0]), x, AA, BB, PP);

        require(onCurve(x, y), "Secp256k1: pubkey not on curve");

        return (x, y);
    }

    /**
     * @notice Return true if x, y are on the secp256k1 curve
     */
    function onCurve(uint256 x, uint256 y) internal pure returns (bool) {
        return EllipticCurve.isOnCurve(x, y, AA, BB, PP);
    }

    /**
     * @notice Revert if the public key is invalid
     * @param pubkey The compressed public key to validate
     */
    function verify(bytes calldata pubkey) internal pure {
        decompress(pubkey);
    }

    /**
     * @notice Convert public key coordinates to an Ethereum address
     * @param x The x coordinate of the public key
     * @param y The y coordinate of the public key
     * @return The Ethereum address corresponding to the public key
     */
    function pubkeyToAddress(uint256 x, uint256 y) internal pure returns (address) {
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
