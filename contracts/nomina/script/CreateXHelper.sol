// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.24;

import { Script, console2 } from "forge-std/Script.sol";
import { Nomina } from "src/token/Nomina.sol";

contract CreateXHelper is Script {
    function nom_staging() public pure returns (bytes memory, bytes32) {
        address omni = 0x73cC960fb6705e9a6A3d9EAf4De94a828CFa6d2a;
        address mintAuthority = 0xE0cF003AC27FaeC91f107E3834968A601842e9c6;

        bytes memory initCode = abi.encodePacked(type(Nomina).creationCode, abi.encode(omni, mintAuthority));
        bytes32 initCodeHash = keccak256(initCode);

        // solhint-disable-next-line no-console
        console2.logBytes(initCode);
        // solhint-disable-next-line no-console
        console2.logBytes32(initCodeHash);

        return (initCode, initCodeHash);
    }
}
