// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { MockToken } from "test/utils/MockToken.sol";
import { MockVault } from "test/utils/MockVault.sol";
import { Outbox } from "src/solve/Outbox.sol";
import { Inbox, IInbox } from "src/solve/Inbox.sol";
import { Solve } from "src/solve/Solve.sol";

import "test/xchain/common/Base.sol";
import { IOmniPortal } from "src/interfaces/IOmniPortal.sol";
import { ConfLevel } from "src/libraries/ConfLevel.sol";
import { XTypes } from "src/libraries/XTypes.sol";
import { TestXTypes } from "test/xchain/common/TestXTypes.sol";
import { Ownable } from "solady/src/auth/Ownable.sol";

/**
 * @title Outbox_fulfill_Test
 * @notice Test suite for solver Outbox.fulfill(...) and Inbox.markFulfilled(...) callback
 * @dev TODO: add fuzz / invariant tests
 */
contract Outbox_fulfill_Test is Base {
    Inbox inbox;
    Outbox outbox;

    MockToken token1;
    MockToken token2;
    MockVault vault1;
    MockVault vault2;

    address user = makeAddr("user");
    address solver = makeAddr("solver");

    function setUp() public override {
        super.setUp();

        vm.deal(user, 10 ether);
        vm.deal(solver, 10 ether);

        address pAdmin = makeAddr("proxy-admin-owner");
        inbox = Inbox(address(new TransparentUpgradeableProxy(address(new Inbox()), pAdmin, "")));
        outbox = Outbox(address(new TransparentUpgradeableProxy(address(new Outbox()), pAdmin, "")));

        inbox.initialize(address(this), solver, address(portal), address(outbox));
        outbox.initialize(address(this), solver, address(chainAPortal), address(inbox));

        token1 = new MockToken();
        token2 = new MockToken();
        vault1 = new MockVault(address(token1));
        vault2 = new MockVault(address(token2));

        outbox.setAllowedCall(address(vault1), MockVault.deposit.selector, true);
        outbox.setAllowedCall(address(vault2), MockVault.deposit.selector, true);
    }

    function test_fulfillFee() public {
        vm.chainId(chainAId);
        uint256 fee = outbox.fulfillFee(thisChainId);
        assertGt(fee, 0);
    }

    function test_fulfill_succeeds() public {
        Solve.Call memory call = _vaultCall(chainAId, address(vault1), user, 1 ether);

        // Fulfill the request as the solver
        vm.chainId(chainAId);
        vm.startPrank(solver);

        // Mint/approve token and prepare fulfill call parameters
        _mintAndApprove(address(token1), 1 ether);
        Solve.TokenPrereq[] memory prereqs = new Solve.TokenPrereq[](1);
        prereqs[0] = Solve.TokenPrereq({ token: address(token1), spender: address(vault1), amount: 1 ether });
        uint256 fee = outbox.fulfillFee(thisChainId);

        // Ensure fulfill call is executed properly
        bytes32 callHash = _getCallHash(bytes32(uint256(1)), thisChainId, call);
        vm.expectCall(call.target, call.data);
        vm.expectEmit(address(outbox));
        emit Outbox.Fulfilled(bytes32(uint256(1)), callHash, solver);
        outbox.fulfill{ value: fee }(bytes32(uint256(1)), thisChainId, solver, call, prereqs);
        vm.stopPrank();

        // Assert token balances
        assertEq(vault1.balances(user), 1 ether, "vault1.balances(user)");
        assertEq(token1.balanceOf(solver), 0, "token1.balanceOf(solver)");
        assertEq(token1.balanceOf(address(outbox)), 0, "token1.balanceOf(outbox)");
    }

    function test_markFulfilled_singleNative() public {
        // Prepare a request to deposit into a vault
        Solve.Call memory call = _vaultCall(chainAId, address(vault1), user, 1 ether);
        uint256 solverBalance = solver.balance;

        // Calculate user payment amount
        vm.chainId(chainAId);
        uint256 fulfillFee = outbox.fulfillFee(thisChainId);
        vm.chainId(thisChainId);
        uint256 nativeFee = inbox.suggestNativePayment(call, 200_000, 0.01 gwei, fulfillFee);

        _completeFulfillment(call, new Solve.TokenDeposit[](0), 1 ether, 1 ether);

        // Verify state
        assertEq(portal.inXMsgOffset(chainAId, ConfLevel.Finalized), uint64(1), "portal.inXMsgOffset");
        assertEq(portal.inXBlockOffset(chainAId, ConfLevel.Finalized), uint64(1), "portal.inXBlockOffset");
        assertEq(vault1.balances(user), 1 ether, "vault1.balances");
        assertEq(token1.balanceOf(solver), 0, "token1.balanceOf(solver)");
        assertEq(token1.balanceOf(address(outbox)), 0, "token1.balanceOf(outbox)");
        assertEq(solver.balance, solverBalance + 1 ether + nativeFee - fulfillFee, "solver.balance");
        assertEq(address(inbox).balance, 0, "inbox.balance");
    }

    function test_markFulfilled_singleToken() public {
        // Prepare a request to deposit into a vault
        Solve.Call memory call = _vaultCall(chainAId, address(vault2), user, 1 ether);
        uint256 solverBalance = solver.balance;

        // Calculate user payment amount
        vm.chainId(chainAId);
        uint256 fulfillFee = outbox.fulfillFee(thisChainId);
        vm.chainId(thisChainId);
        uint256 nativeFee = inbox.suggestNativePayment(call, 200_000, 0.01 gwei, fulfillFee);

        Solve.TokenDeposit[] memory deposits = new Solve.TokenDeposit[](1);
        deposits[0] = Solve.TokenDeposit({ token: address(token1), amount: 1 ether });
        _mintAndApprove(deposits);
        _completeFulfillment(call, deposits, 1 ether, 0);

        // Verify state
        assertEq(portal.inXMsgOffset(chainAId, ConfLevel.Finalized), uint64(1), "portal.inXMsgOffset");
        assertEq(portal.inXBlockOffset(chainAId, ConfLevel.Finalized), uint64(1), "portal.inXBlockOffset");
        assertEq(vault2.balances(user), 1 ether, "vault2.balances");
        assertEq(token1.balanceOf(solver), 1 ether, "token1.balanceOf(solver)");
        assertEq(token2.balanceOf(solver), 0, "token2.balanceOf(solver)");
        assertEq(token1.balanceOf(address(inbox)), 0, "token1.balanceOf(inbox)");
        assertEq(token2.balanceOf(address(outbox)), 0, "token2.balanceOf(outbox)");
        assertEq(solver.balance, solverBalance + nativeFee - fulfillFee, "solver.balance");
        assertEq(address(inbox).balance, 0, "inbox.balance");
    }

    function test_markFulfilled_nativeMultiToken() public {
        // Prepare a request to deposit into a vault
        Solve.Call memory call = _vaultCall(chainAId, address(vault1), user, 1 ether);
        uint256 solverBalance = solver.balance;

        // Calculate user payment amount
        vm.chainId(chainAId);
        uint256 fulfillFee = outbox.fulfillFee(thisChainId);
        vm.chainId(thisChainId);
        uint256 nativeFee = inbox.suggestNativePayment(call, 200_000, 0.01 gwei, fulfillFee);

        Solve.TokenDeposit[] memory deposits = new Solve.TokenDeposit[](2);
        deposits[0] = Solve.TokenDeposit({ token: address(token1), amount: 0.2 ether });
        deposits[1] = Solve.TokenDeposit({ token: address(token2), amount: 0.3 ether });
        _mintAndApprove(deposits);
        _completeFulfillment(call, deposits, 1 ether, 0.5 ether);

        // Verify state
        assertEq(portal.inXMsgOffset(chainAId, ConfLevel.Finalized), uint64(1), "portal.inXMsgOffset");
        assertEq(portal.inXBlockOffset(chainAId, ConfLevel.Finalized), uint64(1), "portal.inXBlockOffset");
        assertEq(vault1.balances(user), 1 ether, "vault1.balances");
        assertEq(token1.balanceOf(solver), 0.2 ether, "token1.balanceOf(solver)");
        assertEq(token2.balanceOf(solver), 0.3 ether, "token2.balanceOf(solver)");
        assertEq(token1.balanceOf(address(inbox)), 0, "token1.balanceOf(inbox)");
        assertEq(token2.balanceOf(address(inbox)), 0, "token2.balanceOf(inbox)");
        assertEq(token1.balanceOf(address(outbox)), 0, "token1.balanceOf(outbox)");
        assertEq(solver.balance, solverBalance + 0.5 ether + nativeFee - fulfillFee, "solver.balance");
        assertEq(address(inbox).balance, 0, "inbox.balance");
    }

    /// @dev Create a call to deposit into a vault.
    function _vaultCall(uint64 destChainId, address vault, address onBehalfOf, uint256 amount)
        internal
        pure
        returns (Solve.Call memory)
    {
        bytes memory data = abi.encodeWithSelector(MockVault.deposit.selector, onBehalfOf, amount);
        return Solve.Call({ destChainId: destChainId, target: vault, value: 0, data: data });
    }

    /// @dev Used by the user to get tokens for the inbox.
    function _mintAndApprove(Solve.TokenDeposit[] memory deposits) internal {
        vm.startPrank(user);
        for (uint256 i = 0; i < deposits.length; i++) {
            MockToken(deposits[i].token).approve(address(inbox), deposits[i].amount);
            MockToken(deposits[i].token).mint(user, deposits[i].amount);
        }
        vm.stopPrank();
    }

    /// @dev Used by the solver to get tokens for the outbox.
    function _mintAndApprove(address token, uint256 amount) internal {
        MockToken(token).approve(address(outbox), amount);
        MockToken(token).mint(solver, amount);
    }

    /// @dev Hash of the call to be included in the xmsg
    function _getCallHash(bytes32 id, uint64 sourceChainId, Solve.Call memory call) internal pure returns (bytes32) {
        return keccak256(abi.encode(id, sourceChainId, call));
    }

    /// @dev Create an xmsg to deliver the markFulfilled callback
    function _getXMsg(bytes32 id, uint64 sourceChainId, uint64 offset, Solve.Call memory call)
        internal
        view
        returns (XTypes.Msg memory)
    {
        return XTypes.Msg({
            destChainId: thisChainId,
            shardId: ConfLevel.Finalized,
            offset: offset,
            sender: address(outbox),
            to: address(inbox),
            data: abi.encodeWithSignature(
                "markFulfilled(bytes32,bytes32,address)", id, _getCallHash(id, sourceChainId, call), solver
            ),
            gasLimit: 200_000
        });
    }

    function _completeFulfillment(
        Solve.Call memory call,
        Solve.TokenDeposit[] memory deposits,
        uint256 depositAmount,
        uint256 paymentAmount
    ) internal {
        // Retrieve fulfill fee
        vm.chainId(chainAId);
        uint256 fulfillFee = outbox.fulfillFee(thisChainId);

        // Create the request as the user
        vm.chainId(thisChainId);
        uint256 nativeFee = inbox.suggestNativePayment(call, 200_000, 0.01 gwei, fulfillFee);
        vm.prank(user);
        bytes32 id = inbox.request{ value: paymentAmount + nativeFee }(call, deposits);

        // Accept the request as the solver
        vm.prank(solver);
        inbox.accept(id);

        // Fulfill the request as the solver
        vm.chainId(chainAId);
        vm.startPrank(solver);
        address token = MockVault(call.target).collateral();
        _mintAndApprove(token, depositAmount);
        Solve.TokenPrereq[] memory prereqs = new Solve.TokenPrereq[](1);
        prereqs[0] = Solve.TokenPrereq({ token: token, spender: call.target, amount: depositAmount });
        outbox.fulfill{ value: outbox.fulfillFee(thisChainId) }(id, thisChainId, solver, call, prereqs);
        vm.stopPrank();

        // Build and perform XSubmission to deliver markFulfilled callback
        vm.chainId(thisChainId);
        XTypes.Msg[] memory xmsgs = new XTypes.Msg[](1);
        xmsgs[0] = _getXMsg(id, thisChainId, uint64(1), call);
        TestXTypes.Block memory xblock = _xblock(chainAId, ConfLevel.Finalized, uint64(1), xmsgs);
        XTypes.Submission memory xsub =
            makeXSub(1, xblock.blockHeader, xblock.msgs, msgFlagsForDest(xblock.msgs, thisChainId));
        bytes32 callHash = _getCallHash(bytes32(uint256(1)), thisChainId, call);

        // Perform and validate XSubmission
        vm.prank(relayer);
        expectCalls(xsub.msgs);
        vm.expectEmit(true, true, true, true, address(inbox));
        emit IInbox.Fulfilled(bytes32(uint256(1)), callHash, solver);
        vm.expectEmit(true, true, true, false, address(portal));
        emit IOmniPortal.XReceipt(chainAId, ConfLevel.Finalized, uint64(1), 0, relayer, true, bytes(""));
        portal.xsubmit(xsub);

        // Claim the request as the solver
        vm.prank(solver);
        inbox.claim(id);
    }
}
