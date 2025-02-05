// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { SolverNetInbox } from "src/SolverNetInbox.sol";
import { SolverNetOutbox } from "src/SolverNetOutbox.sol";

import { MockERC20 } from "test/utils/MockERC20.sol";
import { MockVault } from "test/utils/MockVault.sol";
import { MockMultiTokenVault } from "test/utils/MockMultiTokenVault.sol";
import { MockPortal } from "core/test/utils/MockPortal.sol";

import { IERC7683 } from "src/erc7683/IERC7683.sol";
import { SolverNet } from "src/lib/SolverNet.sol";

import { Test, console2 } from "forge-std/Test.sol";
import { FixedPointMathLib } from "solady/src/utils/FixedPointMathLib.sol";
import { AddrUtils } from "src/lib/AddrUtils.sol";

/**
 * @title TestBase
 * @dev Shared test utils / fixtures.
 */
contract TestBase is Test {
    using AddrUtils for address;
    using AddrUtils for bytes32;

    SolverNetInbox inbox;
    SolverNetOutbox outbox;

    MockERC20 token1;
    MockERC20 token2;

    MockVault nativeVault;
    MockVault erc20Vault;
    MockMultiTokenVault multiTokenVault;

    MockPortal portal;

    uint64 srcChainId = 1;
    uint64 destChainId = 2;

    address user = makeAddr("user");
    address solver = makeAddr("solver");
    address proxyAdmin = makeAddr("proxy-admin-owner");

    uint96 defaultAmount = 100 ether;
    uint32 defaultFillDeadline = uint32(block.timestamp + 1 hours);

    bytes32 internal constant ORDER_DATA_TYPEHASH = keccak256(
        "OrderData(address owner,uint64 destChainId,Deposit deposit,Call[] calls,Expense[] expenses)Deposit(address token,uint96 amount)Call(address target,bytes4 selector,uint256 value,bytes params)Expense(address spender,address token,uint96 amount)"
    );

    modifier prankUser() {
        vm.startPrank(user);
        _;
        vm.stopPrank();
    }

    function setUp() public virtual {
        token1 = new MockERC20("Token 1", "TKN1");
        token2 = new MockERC20("Token 2", "TKN2");

        nativeVault = new MockVault(address(0));
        erc20Vault = new MockVault(address(token2));
        multiTokenVault = new MockMultiTokenVault();
        portal = new MockPortal();

        inbox = deploySolverNetInbox();
        outbox = deploySolverNetOutbox();
        initializeInbox();
        initializeOutbox();

        vm.chainId(srcChainId);
    }

    // Helper functions

    function fillHash(bytes32 orderId, bytes memory originData) internal pure returns (bytes32) {
        return keccak256(abi.encode(orderId, originData));
    }

    function getVaultCall(address vault, uint256 callValue, address depositRecipient, uint256 depositAmount)
        internal
        pure
        returns (SolverNet.Call memory)
    {
        return SolverNet.Call({
            target: vault,
            selector: MockVault.deposit.selector,
            value: callValue,
            params: abi.encode(depositRecipient, depositAmount)
        });
    }

    function getExpense(address spender, address token, uint96 amount)
        internal
        pure
        returns (SolverNet.Expense memory)
    {
        return SolverNet.Expense({ spender: spender, token: token, amount: amount });
    }

    function getOrder(
        address owner,
        uint64 chainId,
        uint32 fillDeadline,
        address depositToken,
        uint96 depositAmount,
        SolverNet.Call[] memory calls,
        SolverNet.Expense[] memory expenses
    ) internal pure returns (SolverNet.OrderData memory, IERC7683.OnchainCrossChainOrder memory) {
        SolverNet.OrderData memory orderData = SolverNet.OrderData({
            owner: owner,
            destChainId: chainId,
            deposit: SolverNet.Deposit({ token: depositToken, amount: depositAmount }),
            calls: calls,
            expenses: expenses
        });

        IERC7683.OnchainCrossChainOrder memory order = IERC7683.OnchainCrossChainOrder({
            fillDeadline: fillDeadline,
            orderDataType: ORDER_DATA_TYPEHASH,
            orderData: abi.encode(orderData)
        });

        return (orderData, order);
    }

    function getGasLimit(bytes memory originData) internal pure returns (uint64) {
        SolverNet.FillOriginData memory fillData = abi.decode(originData, (SolverNet.FillOriginData));

        // 2500 gas for the Metadata struct SLOAD.
        uint256 metadataGas = 2500;

        // 2500 gas for Call array length SLOAD + dynamic cost of reading each call.
        uint256 callsGas = 2500;
        for (uint256 i; i < fillData.calls.length; ++i) {
            SolverNet.Call memory call = fillData.calls[i];
            unchecked {
                // 5000 gas for the two slots that hold target, selector, and value.
                // 2500 gas per params slot (1 per function argument) used (minimum of 1 slot).
                callsGas += 5000 + (FixedPointMathLib.divUp(call.params.length + 32, 32) * 2500);
            }
        }

        // 2500 gas for Expense array length SLOAD + cost of reading each expense.
        uint256 expensesGas = 2500;
        unchecked {
            expensesGas += fillData.expenses.length * 5000;
        }

        return uint64(metadataGas + callsGas + expensesGas + 100_000); // 100k base gas limit
    }

    // Setup functions

    function deploySolverNetInbox() internal returns (SolverNetInbox) {
        address impl = address(new SolverNetInbox());
        return SolverNetInbox(address(new TransparentUpgradeableProxy(impl, proxyAdmin, bytes(""))));
    }

    function deploySolverNetOutbox() internal returns (SolverNetOutbox) {
        address impl = address(new SolverNetOutbox());
        return SolverNetOutbox(address(new TransparentUpgradeableProxy(impl, proxyAdmin, bytes(""))));
    }

    function initializeInbox() internal {
        inbox.initialize(address(this), solver, address(portal));

        uint64[] memory chainIds = new uint64[](1);
        chainIds[0] = destChainId;
        address[] memory outboxes = new address[](1);
        outboxes[0] = address(outbox);
        inbox.setOutboxes(chainIds, outboxes);
    }

    function initializeOutbox() internal {
        outbox.initialize(address(this), solver, address(portal));

        uint64[] memory chainIds = new uint64[](1);
        chainIds[0] = srcChainId;
        address[] memory inboxes = new address[](1);
        inboxes[0] = address(inbox);
        outbox.setInboxes(chainIds, inboxes);
    }
}
