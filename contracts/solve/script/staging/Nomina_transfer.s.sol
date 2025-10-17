// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { SolverNetStagingFixtures } from "../fixtures/SolverNetStagingFixtures.sol";
import { SolverNet } from "src/lib/SolverNet.sol";
import { IERC7683 } from "src/erc7683/IERC7683.sol";

contract Staging_Nomina_transfer is SolverNetStagingFixtures {
    uint96 internal constant defaultAmount = 10_000 ether;

    function run() public {
        IERC7683.OnchainCrossChainOrder memory order = _getOrder(defaultAmount, address(0));
        bool isApproved = _checkApprovals(defaultAmount);

        // Send order, approve tokens if needed
        vm.startBroadcast();
        if (!isApproved) _setApprovals();
        inbox.open(order);
        vm.stopBroadcast();
    }

    function run(uint96 amount) public {
        IERC7683.OnchainCrossChainOrder memory order = _getOrder(amount, address(0));
        bool isApproved = _checkApprovals(amount);

        // Send order, approve tokens if needed
        vm.startBroadcast();
        if (!isApproved) _setApprovals();
        inbox.open(order);
        vm.stopBroadcast();
    }

    function run(uint96 amount, address recipient) public {
        IERC7683.OnchainCrossChainOrder memory order = _getOrder(amount, recipient);
        bool isApproved = _checkApprovals(amount);

        // Send order, approve tokens if needed
        vm.startBroadcast();
        if (!isApproved) _setApprovals();
        inbox.open(order);
        vm.stopBroadcast();
    }

    function _getOrder(uint96 amount, address recipient)
        internal
        view
        returns (IERC7683.OnchainCrossChainOrder memory)
    {
        // Get order, validate it, and check for token approval
        IERC7683.OnchainCrossChainOrder memory order = _getSolverNetOrder(amount, recipient);
        require(inbox.validate(order) == true, "Order is invalid");
        return order;
    }

    function _getSolverNetOrder(uint96 amount, address recipient)
        internal
        view
        returns (IERC7683.OnchainCrossChainOrder memory)
    {
        SolverNet.Deposit memory deposit = SolverNet.Deposit({ token: address(nom), amount: amount });

        SolverNet.Call[] memory call = new SolverNet.Call[](1);
        call[0] = SolverNet.Call({
            target: recipient != address(0) ? recipient : msg.sender, selector: bytes4(""), value: amount, params: ""
        });

        SolverNet.OrderData memory orderData = SolverNet.OrderData({
            owner: msg.sender,
            destChainId: portal.omniChainId(),
            deposit: deposit,
            calls: call,
            expenses: new SolverNet.TokenExpense[](0)
        });

        return IERC7683.OnchainCrossChainOrder({
            fillDeadline: uint32(block.timestamp + 1 hours),
            orderDataType: ORDERDATA_TYPEHASH,
            orderData: abi.encode(orderData)
        });
    }

    function _checkApprovals(uint96 amount) internal view returns (bool) {
        return nom.allowance(msg.sender, address(inbox)) >= amount;
    }

    function _setApprovals() internal {
        nom.approve(address(inbox), type(uint256).max);
    }
}
