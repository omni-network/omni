// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

/**
 * @title Secp256k1
 * @dev Defines secp256k1 utilities
 */
library Secp256k1 {
    /**
     * @notice Convert a validator's public key to its Ethereum address.
     * @param pubkey 64 byte uncompressed secp256k1 public key (no 0x04 prefix)
     */
    function pubkeyToAddress(bytes memory pubkey) internal pure returns (address) {
        require(pubkey.length == 64, "Secp256k1: invalid pubkey length");
        return address(uint160(uint256(keccak256(pubkey))));
    }
}
