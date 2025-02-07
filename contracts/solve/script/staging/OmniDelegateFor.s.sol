// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { Script } from "forge-std/Script.sol";
import { LibString } from "solady/src/utils/LibString.sol";
import { JSONParserLib } from "solady/src/utils/JSONParserLib.sol";
import { IERC20 } from "@openzeppelin/contracts/interfaces/IERC20.sol";
import { IOmniPortal } from "core/src/interfaces/IOmniPortal.sol";
import { IERC7683 } from "src/erc7683/IERC7683.sol";
import { ISolverNetInbox } from "src/interfaces/ISolverNetInbox.sol";
import { IStaking } from "core/src/interfaces/IStaking.sol";
import { SolverNet } from "src/lib/SolverNet.sol";

contract OmniDelegateFor is Script {
    IERC20 internal omni;
    IOmniPortal internal portal;
    ISolverNetInbox internal inbox;
    IStaking internal constant staking = IStaking(0xCCcCcC0000000000000000000000000000000001);

    bytes32 internal constant ORDERDATA_TYPEHASH = keccak256(
        "OrderData(address owner,uint64 destChainId,Deposit deposit,Call[] calls,Expense[] expenses)Deposit(address token,uint96 amount)Call(address target,bytes4 selector,uint256 value,bytes params)Expense(address spender,address token,uint96 amount)"
    );
    address internal constant validator1 = 0xD6CD71dF91a6886f69761826A9C4D123178A8d9D;
    address internal constant validator2 = 0x9C7bf21f72CA34af89F620D27E0F18C4366b88c6;
    uint96 internal constant amount = 100 ether;

    function run() public {
        _setUp();
        IERC7683.OnchainCrossChainOrder memory order = _getOrder();
        bool isApproved = _checkApprovals();

        // Send order, approve tokens if needed
        vm.startBroadcast();
        if (!isApproved) _setApprovals();
        inbox.open(order);
        vm.stopBroadcast();
    }

    function _setUp() internal {
        string memory stagingAddrsJson = _getStagingAddresses();
        _setStagingAddresses(stagingAddrsJson);
    }

    function _getOrder() internal view returns (IERC7683.OnchainCrossChainOrder memory) {
        // Get order, validate it, and check for token approval
        IERC7683.OnchainCrossChainOrder memory order = _getSolverNetOrder();
        require(inbox.validate(order) == true, "Order is invalid");
        return order;
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
        JSONParserLib.Item memory portalItem = JSONParserLib.at(object, '"portal"');
        JSONParserLib.Item memory inboxItem = JSONParserLib.at(object, '"solvernetinbox"');
        /* solhint-enable quotes */

        string memory omniAddr = JSONParserLib.value(omniItem);
        omniAddr = LibString.slice(omniAddr, 1, bytes(omniAddr).length - 1);
        omni = IERC20(vm.parseAddress(omniAddr));

        string memory portalAddr = JSONParserLib.value(portalItem);
        portalAddr = LibString.slice(portalAddr, 1, bytes(portalAddr).length - 1);
        portal = IOmniPortal(vm.parseAddress(portalAddr));

        string memory inboxAddr = JSONParserLib.value(inboxItem);
        inboxAddr = LibString.slice(inboxAddr, 1, bytes(inboxAddr).length - 1);
        inbox = ISolverNetInbox(vm.parseAddress(inboxAddr));
    }

    function _getSolverNetOrder() internal view returns (IERC7683.OnchainCrossChainOrder memory) {
        SolverNet.Deposit memory deposit = SolverNet.Deposit({ token: address(omni), amount: amount });

        SolverNet.Call[] memory call = new SolverNet.Call[](1);
        call[0] = SolverNet.Call({
            target: address(staking),
            selector: staking.delegateFor.selector,
            value: amount,
            params: abi.encodeCall(IStaking.delegateFor, (msg.sender, block.number % 2 == 0 ? validator1 : validator2))
        });

        SolverNet.OrderData memory orderData = SolverNet.OrderData({
            owner: msg.sender,
            destChainId: portal.omniChainId(),
            deposit: deposit,
            calls: call,
            expenses: new SolverNet.Expense[](0)
        });

        return IERC7683.OnchainCrossChainOrder({
            fillDeadline: 0,
            orderDataType: ORDERDATA_TYPEHASH,
            orderData: abi.encode(orderData)
        });
    }

    function _checkApprovals() internal view returns (bool) {
        return omni.allowance(msg.sender, address(inbox)) >= amount;
    }

    function _setApprovals() internal {
        omni.approve(address(inbox), type(uint256).max);
    }
}
