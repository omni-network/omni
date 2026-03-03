// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.24;

import { Script } from "forge-std/Script.sol";
import { Nomina } from "src/token/Nomina.sol";
import { NominaMintLock } from "src/token/NominaMintLock.sol";

contract DeployNominaMintLock is Script {
    function run() public returns (NominaMintLock) {
        Nomina nomina = Nomina(0x6e6F6d696e61decd6605bD4a57836c5DB6923340);

        vm.broadcast();
        NominaMintLock lock = new NominaMintLock(nomina);

        return lock;
    }
}
