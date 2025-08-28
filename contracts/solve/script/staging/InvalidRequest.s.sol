// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { SolverNetStagingFixtures } from "../fixtures/SolverNetStagingFixtures.sol";
import { SolverNet } from "src/lib/SolverNet.sol";
import { IERC7683 } from "src/erc7683/IERC7683.sol";
import { IStaking } from "core/src/interfaces/IStaking.sol";

contract InvalidRequest is SolverNetStagingFixtures {
    IStaking internal constant staking = IStaking(0xCCcCcC0000000000000000000000000000000001);
    address internal constant validator1 = 0xD6CD71dF91a6886f69761826A9C4D123178A8d9D;
    address internal constant validator2 = 0x9C7bf21f72CA34af89F620D27E0F18C4366b88c6;
    uint96 internal constant defaultAmount = 100 ether;

    function run() public {
        IERC7683.OnchainCrossChainOrder memory order =
            _getOrder(defaultAmount, block.number % 2 == 0 ? validator1 : validator2);
        bool isApproved = _checkApprovals(defaultAmount);

        // Send order, approve tokens if needed
        vm.startBroadcast();
        if (!isApproved) _setApprovals();
        inbox.open(order);
        vm.stopBroadcast();
    }

    function _getOrder(uint96 amount, address validator) internal returns (IERC7683.OnchainCrossChainOrder memory) {
        // Get order, validate it, and check for token approval
        IERC7683.OnchainCrossChainOrder memory order = _getSolverNetOrder(amount, validator);
        require(inbox.validate(order) == true, "Order is invalid");
        return order;
    }

    function _getSolverNetOrder(uint96 amount, address validator)
        internal
        returns (IERC7683.OnchainCrossChainOrder memory)
    {
        uint256 rand = vm.randomUint(1, type(uint32).max);
        SolverNet.Deposit memory deposit = SolverNet.Deposit({ token: address(omni), amount: amount });

        SolverNet.Call[] memory call = new SolverNet.Call[](1);
        call[0] = SolverNet.Call({
            target: address(staking),
            selector: bytes4(uint32(rand)),
            value: amount,
            params: abi.encode(msg.sender, validator, rand)
        });

        SolverNet.OrderData memory orderData = SolverNet.OrderData({
            owner: msg.sender,
            destChainId: portal.omniChainId(),
            deposit: deposit,
            calls: call,
            expenses: new SolverNet.TokenExpense[](0)
        });

        return IERC7683.OnchainCrossChainOrder({
            fillDeadline: 0,
            orderDataType: ORDERDATA_TYPEHASH,
            orderData: abi.encode(orderData)
        });
    }

    function _checkApprovals(uint96 amount) internal view returns (bool) {
        return omni.allowance(msg.sender, address(inbox)) >= amount;
    }

    function _setApprovals() internal {
        omni.approve(address(inbox), type(uint256).max);
    }
}
