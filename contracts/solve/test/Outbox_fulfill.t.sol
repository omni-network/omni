// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { MockToken } from "test/utils/MockToken.sol";
import { MockVault } from "test/utils/MockVault.sol";
import { SolveOutbox } from "src/SolveOutbox.sol";
import { ISolveInbox } from "src/SolveInbox.sol";
import { Solve } from "src/Solve.sol";

import { IOmniPortal } from "core/src/interfaces/IOmniPortal.sol";
import { ConfLevel } from "core/src/libraries/ConfLevel.sol";
import { MockPortal } from "core/test/utils/MockPortal.sol";
import { Ownable } from "solady/src/auth/Ownable.sol";
import { Test } from "forge-std/Test.sol";

/**
 * @title SolveOutbox_fulfill_test
 * @notice Test suite for SolveOutbox.fulfill(...) callback
 */
contract SolveOutbox_fulfill_test is Test {
    OutboxHarness outbox;

    MockToken token1;
    MockToken token2;
    MockVault vault1;
    MockVault vault2;

    MockPortal portal;

    address inbox = makeAddr("inbox");
    address user = makeAddr("user");
    address solver = makeAddr("solver");

    function setUp() public {
        vm.deal(user, 10 ether);
        vm.deal(solver, 10 ether);

        token1 = new MockToken();
        token2 = new MockToken();
        vault1 = new MockVault(address(token1));
        vault2 = new MockVault(address(token2));
        portal = new MockPortal();

        outbox = OutboxHarness(
            address(
                new TransparentUpgradeableProxy(
                    address(new OutboxHarness()),
                    makeAddr("proxy-admin-owner"),
                    abi.encodeCall(outbox.initialize, (address(this), solver, address(portal), address(inbox)))
                )
            )
        );

        outbox.setAllowedCall(address(vault1), MockVault.deposit.selector, true);
        outbox.setAllowedCall(address(vault2), MockVault.deposit.selector, true);
    }

    function test_fulfillFee() public view {
        assertGt(outbox.fulfillFee({ srcChainId: 1 }), 0);
    }

    function test_fulfill_reverts() public {
        Solve.Call memory call = vaultCall(address(vault1), user, 1 ether);
        bytes32 srcReqId = bytes32(uint256(1));
        uint64 srcChainId = 1;

        // only solver
        vm.prank(makeAddr("not-solver"));
        vm.expectRevert(Ownable.Unauthorized.selector);
        outbox.fulfill(srcReqId, srcChainId, call, new Solve.TokenPrereq[](0));

        vm.startPrank(solver);

        // only correct dest chain
        vm.chainId(call.destChainId + 1);
        vm.expectRevert(SolveOutbox.WrongDestChain.selector);
        outbox.fulfill(srcReqId, srcChainId, call, new Solve.TokenPrereq[](0));
        vm.chainId(call.destChainId);

        // only allowed calls
        Solve.Call memory notAllowed = randCall();
        vm.expectRevert(SolveOutbox.CallNotAllowed.selector);
        outbox.fulfill(srcReqId, srcChainId, notAllowed, new Solve.TokenPrereq[](0));

        // only if not already fulfilled
        outbox.setDidFullfill(srcReqId, srcChainId, call, true);
        vm.expectRevert(SolveOutbox.AreadyFulfilled.selector);
        outbox.fulfill(srcReqId, srcChainId, call, new Solve.TokenPrereq[](0));
        outbox.setDidFullfill(srcReqId, srcChainId, call, false);

        // call must succeed (will revert because prequisites are insufficient)
        vm.expectRevert(SolveOutbox.CallFailed.selector);
        outbox.fulfill(srcReqId, srcChainId, call, new Solve.TokenPrereq[](0));

        // prequisites must be approved
        Solve.TokenPrereq[] memory prereqs = new Solve.TokenPrereq[](1);
        prereqs[0] = Solve.TokenPrereq({ token: address(token1), spender: address(vault1), amount: 1 ether });
        vm.expectRevert(SafeTransferLib.TransferFromFailed.selector);
        outbox.fulfill(srcReqId, srcChainId, call, prereqs);

        // prequisites must be accurate
        prereqs[0].amount = 2 ether; // too much
        mintAndApprove(prereqs);
        vm.expectRevert(SolveOutbox.IncorrectPrereqs.selector);
        outbox.fulfill(srcReqId, srcChainId, call, prereqs);
        prereqs[0].amount = 1 ether;

        // requires xcall fee
        vm.expectRevert("XApp: insufficient funds");
        outbox.fulfill(srcReqId, srcChainId, call, prereqs);
    }

    function test_fulfill_succeeds() public {
        Solve.Call memory call = vaultCall(address(vault1), user, 1 ether);
        bytes32 srcReqId = bytes32(uint256(1));
        uint64 srcChainId = 1;
        bytes32 callHash = hashCall(srcReqId, srcChainId, call);
        uint256 fee = outbox.fulfillFee(srcChainId);

        vm.startPrank(solver);

        // token prequisites
        Solve.TokenPrereq[] memory prereqs = new Solve.TokenPrereq[](1);
        prereqs[0] = Solve.TokenPrereq({ token: address(token1), spender: address(vault1), amount: 1 ether });
        mintAndApprove(prereqs);

        // expect call to target
        vm.expectCall(call.target, call.data);

        // expect Inbox.markFulfilled xcall
        vm.expectCall(address(portal), markFulfilledXCalldata(srcChainId, srcReqId, callHash));

        // expect emit Fulfilled
        vm.expectEmit(address(outbox));
        emit SolveOutbox.Fulfilled(srcReqId, callHash, solver);

        // fulfill
        outbox.fulfill{ value: fee }(srcReqId, srcChainId, call, prereqs);
        vm.stopPrank();

        // assert call marked as fulfilled on outbox
        assertTrue(outbox.didFulfill(srcReqId, srcChainId, call), "outbox.didFulfill(srcReqId, srcChainId, call)");

        // assert token balances
        assertEq(vault1.balances(user), 1 ether, "vault1.balances(user)");
        assertEq(token1.balanceOf(solver), 0, "token1.balanceOf(solver)");
        assertEq(token1.balanceOf(address(outbox)), 0, "token1.balanceOf(outbox)");

        vm.stopPrank();
    }

    /// @dev Returns a call to deposit into a vault.
    function vaultCall(address vault, address onBehalfOf, uint256 amount) internal view returns (Solve.Call memory) {
        bytes memory data = abi.encodeCall(MockVault.deposit, (onBehalfOf, amount));
        return Solve.Call({ destChainId: uint64(block.chainid), target: vault, value: 0, data: data });
    }

    /// @dev Returns expected OmniPortal.xcall Inbox.markFulfilled calldata
    function markFulfilledXCalldata(uint64 srcChainId, bytes32 srcReqId, bytes32 callHash)
        internal
        view
        returns (bytes memory)
    {
        return abi.encodeCall(
            IOmniPortal.xcall,
            (
                srcChainId,
                ConfLevel.Finalized,
                inbox,
                abi.encodeCall(ISolveInbox.markFulfilled, (srcReqId, callHash)),
                125_000
            )
        );
    }

    /// @dev Returns a random call
    function randCall() internal returns (Solve.Call memory) {
        uint256 rand = vm.randomUint(1, 1000);
        return Solve.Call({
            destChainId: uint64(block.chainid),
            value: rand * 1 ether,
            target: address(uint160(rand)),
            data: abi.encode("data", rand)
        });
    }

    /// @dev Used by the solver to get tokens for the outbox.
    function mintAndApprove(Solve.TokenPrereq[] memory prereqs) internal {
        for (uint256 i = 0; i < prereqs.length; i++) {
            address token = prereqs[i].token;
            uint256 amount = prereqs[i].amount;
            MockToken(token).approve(address(outbox), amount);
            MockToken(token).mint(solver, amount);
        }
    }

    /// @dev Hash of the call to be included in the xmsg
    function hashCall(bytes32 id, uint64 sourceChainId, Solve.Call memory call) internal pure returns (bytes32) {
        return keccak256(abi.encode(id, sourceChainId, call));
    }
}

contract OutboxHarness is SolveOutbox {
    function setDidFullfill(bytes32 srcReqId, uint64 srcChainId, Solve.Call calldata call, bool fulfilled) external {
        fulfilledCalls[_callHash(srcReqId, srcChainId, call)] = fulfilled;
    }
}
