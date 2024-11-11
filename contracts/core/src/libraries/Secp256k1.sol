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
     * @notice Validate a compressed secp256k1 public key
     * @param compressedPubKey The compressed public key to validate
     * @return True if the public key is valid, false otherwise
     */
    function validatePubkey(bytes calldata compressedPubKey) internal pure returns (bool) {
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
}
