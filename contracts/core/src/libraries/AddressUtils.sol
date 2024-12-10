// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

library AddressUtils {
    /**
     * @notice Convert an address to a bytes32
     * @param a The address value to convert
     * @return b The bytes32 representation of the address
     */
    function toBytes32(address a) internal pure returns (bytes32) {
        return bytes32(uint256(uint160(a)));
    }

    /**
     * @notice Convert a bytes32 to an address
     * @param b The bytes32 value to convert
     * @return a The address representation of the bytes32
     *
     * @custom:note Should we return a bool to indicate if its > type(uint160).max, e.g. a non-evm address?
     */
    function toAddress(bytes32 b) internal pure returns (address) {
        return address(uint160(uint256(b)));
    }
}
