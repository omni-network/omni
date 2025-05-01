// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { Script } from "forge-std/Script.sol";
// import { console2 } from "forge-std/console2.sol";
import { SolverNet } from "src/lib/SolverNet.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
// import { FixedPointMathLib } from "solady/src/utils/FixedPointMathLib.sol";
import { IERC20 } from "@openzeppelin/contracts/interfaces/IERC20.sol";
import { IERC7683 } from "src/erc7683/IERC7683.sol";
import { ISolverNetInbox } from "src/interfaces/ISolverNetInbox.sol";

contract ManualSolverNetOrderScript is Script {
    using SafeTransferLib for address;

    ISolverNetInbox internal mainnetInbox = ISolverNetInbox(0x8FCFcd0B4Fa2cc2965a3c7F27995B0A43F210dB8);

    ISolverNetInbox internal inbox;

    uint256 internal callAmount = 0.0077 ether;
    uint256 internal msgValue = 0;
    address internal depositToken = 0x0b2C639c533813f4Aa9D7837CAf62653d097Ff85;
    uint96 internal depositAmount = 14_000_000;
    address internal target = 0xC5759Dd58057808568C698fFDc4Ad5548D17e75a;
    address internal owner = 0xA779fC675Db318dab004Ab8D538CB320D0013F42;
    uint64 internal destChainId = 1;

    bytes32 internal constant ORDERDATA_TYPEHASH = keccak256(
        "OrderData(address owner,uint64 destChainId,Deposit deposit,Call[] calls,TokenExpense[] expenses)Deposit(address token,uint96 amount)Call(address target,bytes4 selector,uint256 value,bytes params)TokenExpense(address spender,address token,uint96 amount)"
    );

    function setUp() public {
        if (block.chainid == 1 || block.chainid == 10 || block.chainid == 8453 || block.chainid == 11_155_420) {
            inbox = mainnetInbox;
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
        IScatterNFT.Auth memory auth = IScatterNFT.Auth({ key: bytes32(0), proof: new bytes32[](0) });
        uint256 quantity = 1;
        address affiliate = address(0);
        bytes memory signature = new bytes(0);

        SolverNet.Deposit memory deposit = SolverNet.Deposit({ token: depositToken, amount: depositAmount });

        SolverNet.Call[] memory call = new SolverNet.Call[](1);
        call[0] = SolverNet.Call({
            target: target,
            selector: IScatterNFT.mintTo.selector,
            value: callAmount,
            params: abi.encode(auth, quantity, owner, affiliate, signature)
        });

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
