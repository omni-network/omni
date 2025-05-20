// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.24;

import { Script } from "forge-std/Script.sol";
// import { console2 } from "forge-std/console2.sol";
import { SolverNet } from "src/lib/SolverNet.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { FixedPointMathLib } from "solady/src/utils/FixedPointMathLib.sol";
import { IERC20 } from "@openzeppelin/contracts/interfaces/IERC20.sol";
import { IERC7683 } from "src/erc7683/IERC7683.sol";
import { ISolverNetInbox } from "src/interfaces/ISolverNetInbox.sol";

contract ManualSolverNetOrderScript is Script {
    using SafeTransferLib for address;

    ISolverNetInbox internal mainnetInbox = ISolverNetInbox(0x8FCFcd0B4Fa2cc2965a3c7F27995B0A43F210dB8);
    ISolverNetInbox internal omegaInbox = ISolverNetInbox(0x7EA0CeB70D5Df75a730E6cB3EADeC12EfdFe80a1);

    address internal mainnetExecutor = 0xf92Dd37ae11F2CCb4de9355BEcEd42Deb4158815;
    address internal omegaExecutor = 0x2b6bf280897ccCBEf827E8546CbD4d28367a8196;

    ISolverNetInbox internal inbox;
    address internal executor;

    uint256 internal callAmount = 0;
    uint256 internal msgValue = 0;
    address internal depositToken = 0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48;
    address internal expenseToken = 0x09Bc4E0D864854c6aFB6eB9A9cdF58aC190D0dF9;
    uint96 internal expenseAmount = 5_000_000;
    uint96 internal depositAmount = uint96(FixedPointMathLib.fullMulDivUp(expenseAmount, 10_030, 10_000));
    address internal target = 0x09Bc4E0D864854c6aFB6eB9A9cdF58aC190D0dF9;
    address internal owner = 0xA779fC675Db318dab004Ab8D538CB320D0013F42;
    uint64 internal destChainId = 5000;

    bytes32 internal constant ORDERDATA_TYPEHASH = keccak256(
        "OrderData(address owner,uint64 destChainId,Deposit deposit,Call[] calls,TokenExpense[] expenses)Deposit(address token,uint96 amount)Call(address target,bytes4 selector,uint256 value,bytes params)TokenExpense(address spender,address token,uint96 amount)"
    );

    function setUp() public {
        if (
            block.chainid == 1 || block.chainid == 10 || block.chainid == 5000 || block.chainid == 8453
                || block.chainid == 11_155_420
        ) {
            inbox = mainnetInbox;
            executor = mainnetExecutor;
        } else if (
            block.chainid == 17_000 || block.chainid == 84_532 || block.chainid == 421_614
                || block.chainid == 11_155_111 || block.chainid == 11_155_420
        ) {
            inbox = omegaInbox;
            executor = omegaExecutor;
        }
    }

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
        call[0] = SolverNet.Call({
            target: target,
            selector: IERC20.transfer.selector,
            value: 0,
            params: abi.encode(owner, expenseAmount)
        });

        SolverNet.TokenExpense[] memory expenses = new SolverNet.TokenExpense[](1);
        expenses[0] = SolverNet.TokenExpense({ spender: address(0), token: expenseToken, amount: expenseAmount });

        SolverNet.OrderData memory orderData = SolverNet.OrderData({
            owner: owner,
            destChainId: destChainId,
            deposit: deposit,
            calls: call,
            expenses: expenses
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
