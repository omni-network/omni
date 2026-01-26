// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { SolverNetInbox } from "src/SolverNetInbox.sol";
import { SolverNetOutbox } from "src/SolverNetOutbox.sol";
import { SolverNetExecutor } from "src/SolverNetExecutor.sol";
import { ISolverNetInbox } from "src/interfaces/ISolverNetInbox.sol";
import { ISolverNetOutbox } from "src/interfaces/ISolverNetOutbox.sol";

import { IERC20 } from "@openzeppelin/contracts/interfaces/IERC20.sol";
import { IOmniPortal } from "core/src/interfaces/IOmniPortal.sol";
import { IMailbox } from "@hyperlane-xyz/core/contracts/interfaces/IMailbox.sol";
import { MockERC20 } from "test/utils/MockERC20.sol";
import { MaxTransferToken } from "test/utils/MaxTransferToken.sol";
import { FeeOnTransferToken } from "test/utils/FeeOnTransferToken.sol";
import { MockVault } from "test/utils/MockVault.sol";
import { MockMultiTokenVault } from "test/utils/MockMultiTokenVault.sol";
import { MockPortal } from "core/test/utils/MockPortal.sol";
import { MockHyperlaneEnvironment } from "test/utils/hyperlane/MockHyperlaneEnvironment.sol";

import { IERC7683 } from "src/erc7683/IERC7683.sol";
import { SolverNet } from "src/lib/SolverNet.sol";
import { HashLib } from "src/lib/HashLib.sol";

import { Test, console2 } from "forge-std/Test.sol";
import { TestStorage } from "./TestStorage.sol";
import { Ownable } from "solady/src/auth/Ownable.sol";
import { Receiver } from "solady/src/accounts/Receiver.sol";
import { FixedPointMathLib } from "solady/src/utils/FixedPointMathLib.sol";
import { AddrUtils } from "src/lib/AddrUtils.sol";
import { Create3 } from "core/src/deploy/Create3.sol";
import { IPermit2, ISignatureTransfer } from "@uniswap/permit2/src/interfaces/IPermit2.sol";

/**
 * @title TestBase
 * @dev Shared test utils / fixtures.
 */
contract TestBase is Test, TestStorage, MockHyperlaneEnvironment {
    using AddrUtils for address;
    using AddrUtils for bytes32;

    SolverNetInbox inbox;
    SolverNetOutbox outbox;
    SolverNetExecutor executor;

    MockERC20 token1;
    MockERC20 token2;
    MaxTransferToken maxTransferToken;
    FeeOnTransferToken feeOnTransferToken;

    MockVault nativeVault;
    MockVault erc20Vault;
    MockMultiTokenVault multiTokenVault;

    MockPortal portal;
    Create3 create3;
    IPermit2 permit2 = IPermit2(0x000000000022D473030F116dDEE9F6B43aC78BA3);

    uint64 srcChainId = 1;
    uint64 destChainId = 2;

    address user;
    uint256 userPk;

    address solver = makeAddr("solver");
    address proxyAdmin = makeAddr("proxy-admin-owner");

    uint96 defaultAmount = 100 ether;
    uint32 defaultOpenDeadline = uint32(block.timestamp + 1 minutes);
    uint32 defaultFillDeadline = uint32(block.timestamp + 1 hours);
    uint32 defaultFillBuffer = 6 hours;

    modifier prankUser() {
        vm.startPrank(user);
        _;
        vm.stopPrank();
    }

    constructor() MockHyperlaneEnvironment(uint32(srcChainId), uint32(destChainId)) { }

    function setUp() public virtual {
        (user, userPk) = makeAddrAndKey("user");

        token1 = new MockERC20("Token 1", "TKN1");
        token2 = new MockERC20("Token 2", "TKN2");
        maxTransferToken = new MaxTransferToken();
        feeOnTransferToken = new FeeOnTransferToken();

        nativeVault = new MockVault(address(0));
        erc20Vault = new MockVault(address(token2));
        multiTokenVault = new MockMultiTokenVault();
        portal = new MockPortal();
        create3 = new Create3();
        vm.etch(address(permit2), PERMIT2_CODE);

        address expectedInboxAddr = create3.getDeployed(address(this), keccak256("inbox"));
        address expectedOutboxAddr = create3.getDeployed(address(this), keccak256("outbox"));
        address expectedExecutorAddr = create3.getDeployed(address(this), keccak256("executor"));

        deploySolverNetInbox(expectedInboxAddr);
        deploySolverNetOutbox(expectedOutboxAddr, expectedExecutorAddr);
        deploySolverNetExecutor(expectedExecutorAddr, expectedOutboxAddr);

        initializeInbox();
        initializeOutbox();
        setRoutes(ISolverNetOutbox.Provider.OmniCore);

        vm.chainId(srcChainId);
    }

    // Helper functions

    function assertResolvedEq(
        IERC7683.ResolvedCrossChainOrder memory resolved1,
        IERC7683.ResolvedCrossChainOrder memory resolved2
    ) internal pure {
        assertEq(keccak256(abi.encode(resolved1)), keccak256(abi.encode(resolved2)), "resolved orders are not equal");
    }

    function fillHash(bytes32 orderId, bytes memory originData) internal pure returns (bytes32) {
        SolverNet.FillOriginData memory fillOriginData = abi.decode(originData, (SolverNet.FillOriginData));
        return keccak256(abi.encode(orderId, fillOriginData));
    }

    function fundUser(SolverNet.OrderData memory orderData) internal {
        SolverNet.Deposit memory deposit = orderData.deposit;
        address token = deposit.token;
        uint96 amount = deposit.amount;

        if (amount > 0) {
            if (token == address(0)) {
                vm.deal(user, amount);
            } else {
                vm.startPrank(user);
                MockERC20(token).mint(user, amount);
                MockERC20(token).approve(address(inbox), type(uint256).max);
                vm.stopPrank();
            }
        }
    }

    function fundUser(SolverNet.OrderData memory orderData, bool gasless) internal {
        SolverNet.Deposit memory deposit = orderData.deposit;
        address token = deposit.token;
        uint96 amount = deposit.amount;

        if (amount > 0) {
            if (token == address(0)) {
                vm.deal(user, amount);
            } else {
                vm.startPrank(user);
                MockERC20(token).mint(user, amount);
                // Gasless orders do not need approvals
                if (!gasless) MockERC20(token).approve(address(inbox), type(uint256).max);
                vm.stopPrank();
            }
        }
    }

    function fundSolver(SolverNet.OrderData memory orderData, uint256 fillFees) internal {
        SolverNet.Call[] memory calls = orderData.calls;
        SolverNet.TokenExpense[] memory expenses = orderData.expenses;

        uint256 nativeValue;
        for (uint256 i; i < calls.length; ++i) {
            SolverNet.Call memory call = calls[i];
            if (call.value > 0) nativeValue += call.value;
        }
        if (nativeValue + fillFees > 0) vm.deal(solver, nativeValue + fillFees);

        for (uint256 i; i < expenses.length; ++i) {
            SolverNet.TokenExpense memory expense = expenses[i];
            address token = expense.token;
            uint96 amount = expense.amount;

            if (amount > 0) {
                vm.prank(solver);
                MockERC20(token).approve(address(outbox), type(uint256).max);
                MockERC20(token).mint(solver, amount);
            }
        }
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
        returns (SolverNet.TokenExpense memory)
    {
        return SolverNet.TokenExpense({ spender: spender, token: token, amount: amount });
    }

    function getOrder(uint256 fillDeadline, SolverNet.OrderData memory orderData)
        internal
        pure
        returns (IERC7683.OnchainCrossChainOrder memory)
    {
        return IERC7683.OnchainCrossChainOrder({
            fillDeadline: uint32(fillDeadline),
            orderDataType: HashLib.OLD_ORDERDATA_TYPEHASH,
            orderData: abi.encode(orderData)
        });
    }

    function getOrder(
        address owner,
        uint64 chainId,
        uint32 fillDeadline,
        address depositToken,
        uint96 depositAmount,
        SolverNet.Call[] memory calls,
        SolverNet.TokenExpense[] memory expenses
    ) internal pure returns (SolverNet.OrderData memory, IERC7683.OnchainCrossChainOrder memory) {
        SolverNet.OrderData memory orderData = SolverNet.OrderData({
            owner: owner,
            destChainId: chainId,
            deposit: SolverNet.Deposit({ token: depositToken, amount: depositAmount }),
            calls: calls,
            expenses: expenses
        });

        IERC7683.OnchainCrossChainOrder memory order = IERC7683.OnchainCrossChainOrder({
            fillDeadline: fillDeadline, orderDataType: HashLib.OLD_ORDERDATA_TYPEHASH, orderData: abi.encode(orderData)
        });

        return (orderData, order);
    }

    function getGaslessOrder(
        address depositor,
        address owner,
        uint256 nonce,
        uint64 chainId,
        uint32 openDeadline,
        uint32 fillDeadline,
        address depositToken,
        uint96 depositAmount,
        SolverNet.Call[] memory calls,
        SolverNet.TokenExpense[] memory expenses
    ) internal view returns (SolverNet.OrderData memory, IERC7683.GaslessCrossChainOrder memory) {
        SolverNet.OrderData memory orderData = SolverNet.OrderData({
            owner: owner,
            destChainId: chainId,
            deposit: SolverNet.Deposit({ token: depositToken, amount: depositAmount }),
            calls: calls,
            expenses: expenses
        });

        IERC7683.GaslessCrossChainOrder memory order = IERC7683.GaslessCrossChainOrder({
            originSettler: address(inbox),
            user: depositor,
            nonce: nonce,
            originChainId: block.chainid,
            openDeadline: openDeadline,
            fillDeadline: fillDeadline,
            orderDataType: HashLib.OMNIORDERDATA_TYPEHASH,
            orderData: abi.encode(orderData)
        });

        return (orderData, order);
    }

    function getNativeForNativeVaultOrder(uint256 depositAmount, uint256 expenseAmount)
        internal
        view
        returns (SolverNet.OrderData memory, IERC7683.OnchainCrossChainOrder memory)
    {
        SolverNet.Deposit memory deposit = SolverNet.Deposit({ token: address(0), amount: uint96(depositAmount) });

        SolverNet.Call[] memory calls = new SolverNet.Call[](1);
        calls[0] = SolverNet.Call({
            target: address(nativeVault),
            selector: MockVault.deposit.selector,
            value: depositAmount,
            params: abi.encode(user, expenseAmount)
        });

        SolverNet.TokenExpense[] memory expenses = new SolverNet.TokenExpense[](0);

        SolverNet.OrderData memory orderData = SolverNet.OrderData({
            owner: address(0), destChainId: destChainId, deposit: deposit, calls: calls, expenses: expenses
        });

        IERC7683.OnchainCrossChainOrder memory order = IERC7683.OnchainCrossChainOrder({
            fillDeadline: defaultFillDeadline,
            orderDataType: HashLib.OLD_ORDERDATA_TYPEHASH,
            orderData: abi.encode(orderData)
        });

        return (orderData, order);
    }

    function getErc20ForErc20VaultOrder(uint256 depositAmount, uint256 expenseAmount)
        internal
        view
        returns (SolverNet.OrderData memory, IERC7683.OnchainCrossChainOrder memory)
    {
        SolverNet.Deposit memory deposit = SolverNet.Deposit({ token: address(token1), amount: uint96(depositAmount) });

        SolverNet.Call[] memory calls = new SolverNet.Call[](1);
        calls[0] = SolverNet.Call({
            target: address(erc20Vault),
            selector: MockVault.deposit.selector,
            value: 0,
            params: abi.encode(user, expenseAmount)
        });

        SolverNet.TokenExpense[] memory expenses = new SolverNet.TokenExpense[](1);
        expenses[0] = SolverNet.TokenExpense({
            spender: address(erc20Vault), token: address(token2), amount: uint96(expenseAmount)
        });

        SolverNet.OrderData memory orderData = SolverNet.OrderData({
            owner: address(0), destChainId: destChainId, deposit: deposit, calls: calls, expenses: expenses
        });

        IERC7683.OnchainCrossChainOrder memory order = IERC7683.OnchainCrossChainOrder({
            fillDeadline: defaultFillDeadline,
            orderDataType: HashLib.OLD_ORDERDATA_TYPEHASH,
            orderData: abi.encode(orderData)
        });

        return (orderData, order);
    }

    function getArbitraryVaultOrder(
        address depositToken,
        uint96 depositAmount,
        address[] memory expenseTokens,
        uint96[] memory expenseAmounts
    ) internal view returns (SolverNet.OrderData memory, IERC7683.OnchainCrossChainOrder memory) {
        require(expenseTokens.length == expenseAmounts.length, "array length mismatch");

        SolverNet.Deposit memory deposit = SolverNet.Deposit({ token: depositToken, amount: depositAmount });

        SolverNet.Call[] memory calls = new SolverNet.Call[](expenseTokens.length);
        for (uint256 i; i < expenseTokens.length; ++i) {
            calls[i] = SolverNet.Call({
                target: expenseTokens[i] == address(0) ? address(nativeVault) : address(erc20Vault),
                selector: MockVault.deposit.selector,
                value: expenseTokens[i] == address(0) ? expenseAmounts[i] : 0,
                params: abi.encode(user, expenseAmounts[i])
            });
        }

        uint256 expenseLength;
        for (uint256 i; i < expenseTokens.length; ++i) {
            if (expenseTokens[i] != address(0)) ++expenseLength;
        }

        uint256 bias;
        SolverNet.TokenExpense[] memory expenses = new SolverNet.TokenExpense[](expenseLength);
        for (uint256 i; i < expenseTokens.length; ++i) {
            if (expenseTokens[i] == address(0)) {
                ++bias;
                continue;
            }

            expenses[i - bias] = SolverNet.TokenExpense({
                spender: address(erc20Vault), token: expenseTokens[i], amount: expenseAmounts[i]
            });
        }

        SolverNet.OrderData memory orderData = SolverNet.OrderData({
            owner: address(0), destChainId: destChainId, deposit: deposit, calls: calls, expenses: expenses
        });

        IERC7683.OnchainCrossChainOrder memory order = IERC7683.OnchainCrossChainOrder({
            fillDeadline: defaultFillDeadline,
            orderDataType: HashLib.OLD_ORDERDATA_TYPEHASH,
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
            uint256 paramsLength = call.params.length;
            unchecked {
                // 5000 gas for the two slots that hold target, selector, and value.
                // 2500 gas per params slot (1 per function argument) used (minimum of 1 slot).
                callsGas += 5000 + (FixedPointMathLib.divUp(call.params.length + 32, 32) * 2500);
                callsGas += (3 * FixedPointMathLib.divUp(paramsLength, 32))
                    + FixedPointMathLib.mulDivUp(paramsLength, paramsLength, 524_288);
            }
        }

        // 2500 gas for TokenExpense array length SLOAD + cost of reading each expense.
        uint256 expensesGas = 2500;
        unchecked {
            expensesGas += fillData.expenses.length * 5000;
        }

        return uint64(metadataGas + callsGas + expensesGas + 100_000); // 100k base gas limit
    }

    // Setup functions

    function deploySolverNetInbox(address expectedInboxAddr) internal {
        address mailbox = address(mailboxes[uint32(srcChainId)]);
        address impl = address(new SolverNetInbox(address(portal), mailbox));
        inbox = SolverNetInbox(
            create3.deploy(
                keccak256("inbox"),
                abi.encodePacked(
                    type(TransparentUpgradeableProxy).creationCode, abi.encode(impl, proxyAdmin, bytes(""))
                )
            )
        );
        require(address(inbox) == expectedInboxAddr, "inbox address mismatch");
    }

    function deploySolverNetOutbox(address expectedOutboxAddr, address expectedExecutorAddr) internal {
        address mailbox = address(mailboxes[uint32(destChainId)]);
        address impl = address(new SolverNetOutbox(expectedExecutorAddr, address(portal), mailbox));
        outbox = SolverNetOutbox(
            create3.deploy(
                keccak256("outbox"),
                abi.encodePacked(
                    type(TransparentUpgradeableProxy).creationCode, abi.encode(impl, proxyAdmin, bytes(""))
                )
            )
        );
        require(address(outbox) == expectedOutboxAddr, "outbox address mismatch");
    }

    function deploySolverNetExecutor(address expectedExecutorAddr, address expectedOutboxAddr) internal {
        executor = SolverNetExecutor(
            payable(create3.deploy(
                    keccak256("executor"),
                    abi.encodePacked(type(SolverNetExecutor).creationCode, abi.encode(address(expectedOutboxAddr)))
                ))
        );
        require(address(executor) == expectedExecutorAddr, "executor address mismatch");
    }

    function initializeInbox() internal {
        inbox.initialize(address(this), solver);
    }

    function initializeOutbox() internal {
        outbox.initialize(address(this), solver);
    }

    function setRoutes(ISolverNetOutbox.Provider provider) internal {
        // Configure inbox
        uint64[] memory chainIds = new uint64[](2);
        chainIds[0] = srcChainId;
        chainIds[1] = destChainId;
        address[] memory outboxes = new address[](2);
        outboxes[0] = address(outbox);
        outboxes[1] = address(outbox);
        inbox.setOutboxes(chainIds, outboxes);

        // Configure outbox
        chainIds = new uint64[](1);
        chainIds[0] = srcChainId;
        ISolverNetOutbox.InboxConfig[] memory configs = new ISolverNetOutbox.InboxConfig[](1);
        configs[0] = ISolverNetOutbox.InboxConfig({ inbox: address(inbox), provider: provider });
        outbox.setInboxes(chainIds, configs);
    }

    function assertStatus(bytes32 orderId, ISolverNetInbox.Status status) internal view {
        (, ISolverNetInbox.OrderState memory state,) = inbox.getOrder(orderId);

        uint8 expect = uint8(status);
        uint8 actual = uint8(state.status);

        if (status == ISolverNetInbox.Status.Pending) assertEq(expect, actual, "order should be pending");
        if (status == ISolverNetInbox.Status.Claimed) assertEq(expect, actual, "order should be claimed");
        if (status == ISolverNetInbox.Status.Rejected) assertEq(expect, actual, "order should be rejected");
        if (status == ISolverNetInbox.Status.Closed) assertEq(expect, actual, "order should be closed");
        if (status == ISolverNetInbox.Status.Filled) assertEq(expect, actual, "order should be filled");
        if (status == ISolverNetInbox.Status.Invalid) revert("invalid status");
    }
}
