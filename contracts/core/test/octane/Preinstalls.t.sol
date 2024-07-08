// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { Preinstalls } from "src/octane/Preinstalls.sol";
import { Test, Vm } from "forge-std/Test.sol";

interface IEIP712 {
    function DOMAIN_SEPARATOR() external view returns (bytes32);
}

/**
 * @title Preinstalls_Test
 * @notice Test suite for Preinstalls.sol. Most of Preinstalls is just static bytecode.
 *         We only tests Permit2 templating.
 */
contract Preinstalls_Test is Test {
    /**
     * @notice Test getPermit2Code templating. This templating inserts immutable variables into the bytecode.
     */
    function test_getPermit2Code() public {
        bytes32 typeHash =
            keccak256(abi.encodePacked("EIP712Domain(string name,uint256 chainId,address verifyingContract)"));
        bytes32 nameHash = keccak256(abi.encodePacked("Permit2"));
        uint256 chainId = 165;
        bytes32 domainSeparator = keccak256(abi.encode(typeHash, nameHash, chainId, Preinstalls.Permit2));

        vm.etch(Preinstalls.Permit2, Preinstalls.getPermit2Code(chainId));

        vm.chainId(chainId);
        assertEq(IEIP712(Preinstalls.Permit2).DOMAIN_SEPARATOR(), domainSeparator);
    }
}
