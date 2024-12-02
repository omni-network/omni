// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { Pledge } from "../../src/pledge/Pledge.sol";
import { Script, console2 } from "forge-std/Script.sol";
import { ICreateX } from "createx/src/ICreateX.sol";

contract PledgeScript is Script {
    ICreateX public createX = ICreateX(0xba5Ed099633D3B313e4D5F7bdc1305d3c28ba5Ed);
    address public omni = 0x5e9A8Aa213C912Bf54C86bf64aDB8ed6A79C04d1;
    bytes32 public salt = 0xa779fc675db318dab004ab8d538cb320d0013f42009066ee4061894802758139;

    function getInitCodeHash() public view returns (bytes32) {
        bytes memory bytecode = type(Pledge).creationCode;
        bytes memory initCode = abi.encodePacked(bytecode, abi.encode(omni));
        return keccak256(initCode);
    }

    function deploy() external {
        bytes memory bytecode = type(Pledge).creationCode;
        bytes memory initCode = abi.encodePacked(bytecode, abi.encode(omni));

        vm.startBroadcast();
        address pledge = createX.deployCreate2(salt, initCode);
        vm.stopBroadcast();
        //solhint-disable-next-line no-console
        console2.log("Pledge deployed at", pledge);
    }
}
