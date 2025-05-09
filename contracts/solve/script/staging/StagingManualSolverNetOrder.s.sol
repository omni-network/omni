// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.24;

import { SolverNetStagingFixtures } from "../SolverNetStagingFixtures.sol";
// import { console2 } from "forge-std/console2.sol";
import { SolverNet } from "src/lib/SolverNet.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { FixedPointMathLib } from "solady/src/utils/FixedPointMathLib.sol";
import { IERC20 } from "@openzeppelin/contracts/interfaces/IERC20.sol";
import { IERC7683 } from "src/erc7683/IERC7683.sol";

contract StagingManualSolverNetOrderScript is SolverNetStagingFixtures {
    using SafeTransferLib for address;

    uint256 internal callAmount = 0.01 ether;
    uint256 internal msgValue = FixedPointMathLib.fullMulDivUp(callAmount, 10_030, 10_000);
    address internal depositToken = address(0);
    uint96 internal depositAmount = uint96(msgValue);
    address internal owner = 0xA779fC675Db318dab004Ab8D538CB320D0013F42;
    address internal target = owner;
    uint64 internal destChainId = 11_155_111;

    function run() public {
        IERC7683.OnchainCrossChainOrder memory order = _getOrder();
        bool isApproved = _checkApprovals();

        // Send order, approve tokens if needed
        vm.startBroadcast();
        if (!isApproved) _setApprovals();
        inbox.open{ value: msgValue }(order);
        vm.stopBroadcast();
    }

    function _getOrder() internal view returns (IERC7683.OnchainCrossChainOrder memory) {
        // Get order, validate it, and check for token approval
        IERC7683.OnchainCrossChainOrder memory order = _getSolverNetOrder();
        require(inbox.validate(order) == true, "Order is invalid");
        return order;
    }

    function _getSolverNetOrder() internal view returns (IERC7683.OnchainCrossChainOrder memory) {
        SolverNet.Deposit memory deposit = SolverNet.Deposit({ token: depositToken, amount: depositAmount });

        SolverNet.Call[] memory call = new SolverNet.Call[](1);
        call[0] = SolverNet.Call({ target: target, selector: hex"00000000", value: callAmount, params: "" });

        SolverNet.OrderData memory orderData = SolverNet.OrderData({
            owner: owner,
            destChainId: destChainId,
            deposit: deposit,
            calls: call,
            expenses: new SolverNet.TokenExpense[](0)
        });

        return IERC7683.OnchainCrossChainOrder({
            fillDeadline: uint32(block.timestamp + 1 days),
            orderDataType: ORDERDATA_TYPEHASH,
            orderData: abi.encode(orderData)
        });
    }

    function _checkApprovals() internal view returns (bool) {
        if (depositToken == address(0)) return true;
        return IERC20(depositToken).allowance(owner, address(inbox)) >= depositAmount;
    }

    function _setApprovals() internal {
        depositToken.safeApprove(address(inbox), type(uint256).max);
    }
}

interface IScatterNFT {
    struct Auth {
        bytes32 key;
        bytes32[] proof;
    }

    function mintTo(Auth calldata auth, uint256 quantity, address to, address affiliate, bytes calldata signature)
        external
        payable;
}
