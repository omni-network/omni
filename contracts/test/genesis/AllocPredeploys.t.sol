// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { AllocPredeploys } from "script/genesis/AllocPredeploys.s.sol";
import { Test } from "forge-std/Test.sol";
import { Process } from "./utils/Process.sol";

/**
 * @title AllocPredeploys_Test
 * @notice Test suite for AllocPredeploys script.
 * @dev We inherit from AllocPredeploys so that vm.stateDump() is called from this contract,
 *      which keeps this contract's state out of the state dump.
 */
contract AllocPredeploys_Test is Test, AllocPredeploys {
    /**
     * @notice Tests predeploy allocs, asserting the number of allocs is expected.
     */
    function test_allocs() public {
        address admin = makeAddr("admin");
        string memory output = tmpfile();

        this.runWithCfg(
            AllocPredeploys.Config({ admin: admin, chainId: 165, enableStakingAllowlist: false, output: output })
        );

        uint256 expected = 0;
        expected += 1024 * 2; // namespace size * 2
        expected += 1; // ProxyAdmin
        expected += 4; // predeploy implementations (excl. not prodiex WOmni and ProxyAdmin)
        expected += 15; // preinstalls
        expected += 1; // 4788 deployer account (nonce set to 1)

        assertEq(expected, getJSONKeyCount(output), "key count check");

        deleteFile(output);
    }

    //////////////////////////////////////////////////////////////////////////////
    //                      FS / JSON Utils                                     //
    //////////////////////////////////////////////////////////////////////////////

    /**
     * @notice Creates a temp file and returns the path to it.
     */
    function tmpfile() internal returns (string memory) {
        string[] memory commands = new string[](3);
        commands[0] = "bash";
        commands[1] = "-c";
        commands[2] = "mktemp";
        bytes memory result = Process.run(commands);
        return string(result);
    }

    /**
     * @notice Deletes a file at a given filesystem path.
     */
    function deleteFile(string memory path) internal {
        string[] memory commands = new string[](3);
        commands[0] = "bash";
        commands[1] = "-c";
        commands[2] = string.concat("rm ", path);
        Process.run({ _command: commands, _allowEmpty: true });
    }

    /**
     * @notice Returns the number of top level keys in a JSON object at a given file path.
     */
    function getJSONKeyCount(string memory path) internal returns (uint256) {
        string[] memory commands = new string[](3);
        commands[0] = "bash";
        commands[1] = "-c";
        commands[2] = string.concat("jq 'keys | length' < ", path, " | xargs cast abi-encode 'f(uint256)'");
        return abi.decode(Process.run(commands), (uint256));
    }
}
