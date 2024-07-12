// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

library Admins {
    address internal constant devnet = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266;
    address internal constant staging = 0x4891925c4f13A34FC26453FD168Db80aF3273014;
    address internal constant omega = 0xEAD625eB2011394cdD739E91Bf9D51A7169C22F5;

    function forNetwork(string calldata network) internal pure returns (address) {
        bytes32 hash = keccak256(abi.encodePacked(network));

        if (hash == keccak256(abi.encodePacked("devnet"))) return devnet;
        if (hash == keccak256(abi.encodePacked("staging"))) return staging;
        if (hash == keccak256(abi.encodePacked("omega"))) return omega;
        revert("Admins: invalid network");
    }
}
