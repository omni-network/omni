// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

library AddrUtils {
    /**
     * @dev Convert address to bytes32.
     * @param a  Address to convert.
     */
    function toBytes32(address a) internal pure returns (bytes32) {
        return bytes32(uint256(uint160(a)));
    }

    /**
     * @dev Convert bytes32 to address.
     * @param b  Bytes32 to convert.
     */
    function toAddress(bytes32 b) internal pure returns (address) {
        return address(uint160(uint256(b)));
    }
}
