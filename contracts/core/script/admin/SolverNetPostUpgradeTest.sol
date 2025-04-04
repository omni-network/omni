// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { Test } from "forge-std/Test.sol";
import { VmSafe } from "forge-std/Vm.sol";

import { SolverNetInbox, ISolverNetInbox } from "solve/src/SolverNetInbox.sol";
import { SolverNetOutbox } from "solve/src/SolverNetOutbox.sol";
import { SolverNetMiddleman } from "solve/src/SolverNetMiddleman.sol";
import { SolverNetExecutor } from "solve/src/SolverNetExecutor.sol";
import { IERC7683 } from "solve/src/erc7683/IERC7683.sol";
import { SolverNet } from "solve/src/lib/SolverNet.sol";

contract SolverNetPostUpgradeTest is Test {
    SolverNetInbox inbox;
    SolverNetOutbox outbox;
    SolverNetMiddleman middleman;
    SolverNetExecutor executor;

    bytes32 constant ORDERDATA_TYPEHASH = keccak256(
        "OrderData(address owner,uint64 destChainId,Deposit deposit,Call[] calls,TokenExpense[] expenses)Deposit(address token,uint96 amount)Call(address target,bytes4 selector,uint256 value,bytes params)TokenExpense(address spender,address token,uint96 amount)"
    );

    address owner;
    address user = makeAddr("user");

    function runInbox(address addr) public {
        (VmSafe.CallerMode mode,,) = vm.readCallers();
        require(mode == VmSafe.CallerMode.None, "no broadcast");

        _setupInbox(addr);
        _openOrder();
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

    function _openOrder() internal {
        SolverNet.Deposit memory deposit = SolverNet.Deposit({ token: address(0), amount: 1 ether });

        SolverNet.Call[] memory calls = new SolverNet.Call[](1);
        calls[0] = SolverNet.Call({ target: user, selector: bytes4(hex"00000000"), value: 1 ether, params: "" });

        SolverNet.OrderData memory orderData = SolverNet.OrderData({
            owner: user,
            destChainId: uint64(block.chainid == 1 ? 10 : 1),
            deposit: deposit,
            calls: calls,
            expenses: new SolverNet.TokenExpense[](0)
        });

        IERC7683.OnchainCrossChainOrder memory order = IERC7683.OnchainCrossChainOrder({
            fillDeadline: type(uint32).max,
            orderDataType: ORDERDATA_TYPEHASH,
            orderData: abi.encode(orderData)
        });

        bytes32 id = inbox.getNextOnchainOrderId(user);
        SolverNet.FillOriginData memory fillOriginData = SolverNet.FillOriginData({
            srcChainId: uint64(block.chainid),
            destChainId: uint64(block.chainid == 1 ? 10 : 1),
            fillDeadline: type(uint32).max,
            calls: calls,
            expenses: new SolverNet.TokenExpense[](0)
        });

        vm.deal(user, 1 ether);
        vm.prank(user);
        vm.expectEmit(address(inbox));
        emit ISolverNetInbox.FillOriginData(id, fillOriginData);
        inbox.open{ value: 1 ether }(order);
    }
}
