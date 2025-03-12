// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { Test } from "forge-std/Test.sol";
import { VmSafe } from "forge-std/Vm.sol";

import { SolverNetInbox } from "solve/src/SolverNetInbox.sol";
import { SolverNetOutbox } from "solve/src/SolverNetOutbox.sol";
import { SolverNetMiddleman } from "solve/src/SolverNetMiddleman.sol";
import { SolverNetExecutor } from "solve/src/SolverNetExecutor.sol";

contract SolverNetPostUpgradeTest is Test {
    SolverNetInbox inbox;
    SolverNetOutbox outbox;
    SolverNetMiddleman middleman;
    SolverNetExecutor executor;

    address owner;

    function runInbox(address addr) public {
        (VmSafe.CallerMode mode,,) = vm.readCallers();
        require(mode == VmSafe.CallerMode.None, "no broadcast");

        _setupInbox(addr);
    }

    function runOutbox(address addr) public {
        (VmSafe.CallerMode mode,,) = vm.readCallers();
        require(mode == VmSafe.CallerMode.None, "no broadcast");

        _setupOutbox(addr);
    }

    function runMiddleman(address addr) public {
        (VmSafe.CallerMode mode,,) = vm.readCallers();
        require(mode == VmSafe.CallerMode.None, "no broadcast");

        _setupMiddleman(addr);
    }

    function runExecutor(address addr) public {
        (VmSafe.CallerMode mode,,) = vm.readCallers();
        require(mode == VmSafe.CallerMode.None, "no broadcast");

        _setupExecutor(addr);
    }

    function _setupInbox(address addr) internal {
        inbox = SolverNetInbox(addr);
        owner = inbox.owner();
    }

    function _setupOutbox(address addr) internal {
        outbox = SolverNetOutbox(addr);
        owner = outbox.owner();
    }

    function _setupMiddleman(address addr) internal {
        middleman = SolverNetMiddleman(payable(addr));
    }

    function _setupExecutor(address addr) internal {
        executor = SolverNetExecutor(payable(addr));
    }
}
