// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { Script } from "forge-std/Script.sol";
import { LibString } from "solady/src/utils/LibString.sol";
import { JSONParserLib } from "solady/src/utils/JSONParserLib.sol";
import { IERC20 } from "@openzeppelin/contracts/interfaces/IERC20.sol";
import { IOmniPortal } from "core/src/interfaces/IOmniPortal.sol";
import { ISolverNetInbox } from "src/interfaces/ISolverNetInbox.sol";

contract SolverNetStagingFixtures is Script {
    IERC20 internal omni;
    IERC20 internal nom;
    IOmniPortal internal portal;
    ISolverNetInbox internal inbox;

    bytes32 internal constant ORDERDATA_TYPEHASH = keccak256(
        "OrderData(address owner,uint64 destChainId,Deposit deposit,Call[] calls,TokenExpense[] expenses)Deposit(address token,uint96 amount)Call(address target,bytes4 selector,uint256 value,bytes params)TokenExpense(address spender,address token,uint96 amount)"
    );

    function setUp() public {
        string memory stagingAddrsJson = _getStagingAddresses();
        _setStagingAddresses(stagingAddrsJson);
    }

    function _getStagingAddresses() internal returns (string memory) {
        string[] memory inputs = new string[](3);
        inputs[0] = "go";
        inputs[1] = "run";
        inputs[2] = "../../scripts/stagingaddrs/main.go";

        bytes memory stagingAddrsJson = vm.ffi(inputs);
        return string(stagingAddrsJson);
    }

    function _setStagingAddresses(string memory stagingAddrsJson) internal {
        JSONParserLib.Item memory object = JSONParserLib.parse(stagingAddrsJson);
        /* solhint-disable quotes */
        JSONParserLib.Item memory omniItem = JSONParserLib.at(object, '"token"');
        JSONParserLib.Item memory nomItem = JSONParserLib.at(object, '"nom"');
        JSONParserLib.Item memory portalItem = JSONParserLib.at(object, '"portal"');
        JSONParserLib.Item memory inboxItem = JSONParserLib.at(object, '"solvernetinbox"');
        /* solhint-enable quotes */

        string memory omniAddr = JSONParserLib.value(omniItem);
        omniAddr = LibString.slice(omniAddr, 1, bytes(omniAddr).length - 1);
        omni = IERC20(vm.parseAddress(omniAddr));

        string memory nomAddr = JSONParserLib.value(nomItem);
        nomAddr = LibString.slice(nomAddr, 1, bytes(nomAddr).length - 1);
        nom = IERC20(vm.parseAddress(nomAddr));

        string memory portalAddr = JSONParserLib.value(portalItem);
        portalAddr = LibString.slice(portalAddr, 1, bytes(portalAddr).length - 1);
        portal = IOmniPortal(vm.parseAddress(portalAddr));

        string memory inboxAddr = JSONParserLib.value(inboxItem);
        inboxAddr = LibString.slice(inboxAddr, 1, bytes(inboxAddr).length - 1);
        inbox = ISolverNetInbox(vm.parseAddress(inboxAddr));
    }
}
