// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { SolverNetInboxV2 } from "src/SolverNetInboxV2.sol";
import { SolverNetOutbox } from "src/SolverNetOutbox.sol";
import { SolverNetMiddleman } from "src/SolverNetMiddleman.sol";
import { SolverNetExecutor } from "src/SolverNetExecutor.sol";
import { ISolverNetInboxV2 } from "src/interfaces/ISolverNetInboxV2.sol";
import { ISolverNetOutbox } from "src/interfaces/ISolverNetOutbox.sol";

import { IERC20 } from "@openzeppelin/contracts/interfaces/IERC20.sol";
import { MockERC20 } from "test/utils/MockERC20.sol";
import { MockVault } from "test/utils/MockVault.sol";
import { MockMultiTokenVault } from "test/utils/MockMultiTokenVault.sol";
import { MockPortal } from "core/test/utils/MockPortal.sol";
import { IPermit2, ISignatureTransfer } from "@uniswap/permit2/src/interfaces/IPermit2.sol";
import { IERC7683 } from "src/erc7683/IERC7683.sol";
import { SolverNet } from "src/lib/SolverNet.sol";
import { HashLib } from "src/lib/HashLib.sol";

import { Test, console2 } from "forge-std/Test.sol";
import { TestStorage } from "./TestStorage.sol";
import { Ownable } from "solady/src/auth/Ownable.sol";
import { FixedPointMathLib } from "solady/src/utils/FixedPointMathLib.sol";
import { AddrUtils } from "src/lib/AddrUtils.sol";

/**
 * @title TestBaseV2
 * @dev Shared test utils / fixtures.
 */
contract TestBaseV2 is Test, TestStorage {
    using AddrUtils for address;
    using AddrUtils for bytes32;

    SolverNetInboxV2 inbox;
    SolverNetOutbox outbox;
    SolverNetMiddleman middleman;
    SolverNetExecutor executor;

    MockERC20 token1;
    MockERC20 token2;

    MockVault nativeVault;
    MockVault erc20Vault;
    MockMultiTokenVault multiTokenVault;

    MockPortal portal;
    IPermit2 permit2;

    uint64 srcChainId = 1;
    uint64 destChainId = 2;

    address user;
    uint256 userPk;

    address solver = makeAddr("solver");
    address proxyAdmin = makeAddr("proxy-admin-owner");

    uint96 defaultAmount = 100 ether;
    uint32 defaultFillDeadline = uint32(block.timestamp + 1 hours);
    uint32 defaultFillBuffer = 6 hours;

    address internal constant PERMIT2 = 0x000000000022D473030F116dDEE9F6B43aC78BA3;

    bytes32 internal constant ORDERDATA_TYPEHASH = keccak256(
        "OrderData(address owner,uint64 destChainId,Deposit deposit,Call[] calls,TokenExpense[] expenses)Deposit(address token,uint96 amount)Call(address target,bytes4 selector,uint256 value,bytes params)TokenExpense(address spender,address token,uint96 amount)"
    );

    bytes32 internal constant OMNIORDERDATA_TYPEHASH = keccak256(
        "OmniOrderData(address owner,uint64 destChainId,Deposit deposit,Call[] calls,TokenExpense[] expenses)Call(address target,bytes4 selector,uint256 value,bytes params)Deposit(address token,uint96 amount)TokenExpense(address spender,address token,uint96 amount)"
    );

    modifier prankUser() {
        vm.startPrank(user);
        _;
        vm.stopPrank();
    }

    function setUp() public virtual {
        (user, userPk) = makeAddrAndKey("user");

        token1 = new MockERC20("Token 1", "TKN1");
        token2 = new MockERC20("Token 2", "TKN2");

        nativeVault = new MockVault(address(0));
        erc20Vault = new MockVault(address(token2));
        multiTokenVault = new MockMultiTokenVault();

        portal = new MockPortal();
        permit2 = etchPermit2();

        inbox = deploySolverNetInbox();
        outbox = deploySolverNetOutbox();
        middleman = new SolverNetMiddleman();
        executor = new SolverNetExecutor(address(outbox));
        initializeInbox();
        initializeOutbox();

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
        return keccak256(abi.encode(orderId, originData));
    }

    function fundUser(SolverNet.OrderData memory orderData) internal {
        SolverNet.Deposit memory deposit = orderData.deposit;
        address token = deposit.token;
        uint96 amount = deposit.amount;

        if (amount > 0) {
            if (token == address(0)) {
                vm.deal(user, amount);
            } else {
                vm.prank(user);
                MockERC20(token).approve(address(inbox), type(uint256).max);
                MockERC20(token).mint(user, amount);
            }
        }
    }

    function fundUser(SolverNet.Deposit memory deposit) internal {
        address token = deposit.token;
        uint96 amount = deposit.amount;

        if (amount > 0) {
            if (token == address(0)) {
                vm.deal(user, amount);
            } else {
                vm.prank(user);
                MockERC20(token).approve(address(inbox), type(uint256).max);
                MockERC20(token).mint(user, amount);
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

    function getPermit2Signature(IERC7683.GaslessCrossChainOrder memory order, SolverNet.OmniOrderData memory orderData)
        internal
        view
        returns (bytes memory)
    {
        // Required typehashes for Permit2's permitWitnessTransferFrom
        bytes32 TOKEN_PERMISSIONS_TYPEHASH = keccak256(bytes(HashLib.TOKEN_PERMISSIONS_TYPE));
        // Specific typehash for our witness type
        bytes32 PERMIT_WITNESS_TRANSFER_FROM_TYPEHASH = keccak256(
            abi.encodePacked(
                bytes(HashLib.PERMIT_WITNESS_TRANSFER_FROM_TYPE_STUB),
                bytes(HashLib.PERMIT2_ORDER_TYPE), // The specific witness type string used by SolverNetInboxV2
                bytes(HashLib.TOKEN_PERMISSIONS_TYPE)
            )
        );

        // 1. Calculate order data hash and the witness (full order hash)
        bytes32 orderDataHash = HashLib.hashOrderData(orderData);
        bytes32 witness = HashLib.hashGaslessOrder(order, orderDataHash);

        // 2. Calculate Permit2 Domain Separator
        bytes32 domainSeparator = permit2.DOMAIN_SEPARATOR();

        // 3. Hash the TokenPermissions struct
        bytes32 hashedTokenPermissions =
            keccak256(abi.encode(TOKEN_PERMISSIONS_TYPEHASH, orderData.deposit.token, orderData.deposit.amount));

        // 4. Hash the PermitWitnessTransferFrom struct
        bytes32 structHash = keccak256(
            abi.encode(
                PERMIT_WITNESS_TRANSFER_FROM_TYPEHASH,
                hashedTokenPermissions,
                address(inbox),
                order.nonce,
                order.openDeadline,
                witness
            )
        );

        // 5. Combine domain separator and struct hash according to EIP-712
        bytes32 digest = keccak256(abi.encodePacked("\x19\x01", domainSeparator, structHash));

        // 6. Sign the digest
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(userPk, digest);

        return abi.encodePacked(r, s, v);
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

    function getOnchainOrder(SolverNet.OrderData memory orderData)
        internal
        view
        returns (IERC7683.OnchainCrossChainOrder memory)
    {
        return IERC7683.OnchainCrossChainOrder({
            fillDeadline: defaultFillDeadline,
            orderDataType: ORDERDATA_TYPEHASH,
            orderData: abi.encode(orderData)
        });
    }

    function getGaslessOrder(uint256 nonce, SolverNet.OmniOrderData memory orderData)
        internal
        view
        returns (IERC7683.GaslessCrossChainOrder memory)
    {
        return IERC7683.GaslessCrossChainOrder({
            originSettler: address(inbox),
            user: user,
            nonce: nonce,
            originChainId: srcChainId,
            openDeadline: uint32(block.timestamp),
            fillDeadline: defaultFillDeadline,
            orderDataType: OMNIORDERDATA_TYPEHASH,
            orderData: abi.encode(orderData)
        });
    }

    function getOnchainOrder(
        address owner,
        address depositToken,
        uint96 depositAmount,
        SolverNet.Call[] memory calls,
        SolverNet.TokenExpense[] memory expenses
    ) internal view returns (SolverNet.OrderData memory, IERC7683.OnchainCrossChainOrder memory) {
        SolverNet.OrderData memory orderData = SolverNet.OrderData({
            owner: owner,
            destChainId: destChainId,
            deposit: SolverNet.Deposit({ token: depositToken, amount: depositAmount }),
            calls: calls,
            expenses: expenses
        });

        IERC7683.OnchainCrossChainOrder memory order = IERC7683.OnchainCrossChainOrder({
            fillDeadline: defaultFillDeadline,
            orderDataType: ORDERDATA_TYPEHASH,
            orderData: abi.encode(orderData)
        });

        return (orderData, order);
    }

    function getGaslessOrder(
        address owner,
        uint256 nonce,
        address depositToken,
        uint96 depositAmount,
        SolverNet.Call[] memory calls,
        SolverNet.TokenExpense[] memory expenses
    ) internal view returns (SolverNet.OmniOrderData memory, IERC7683.GaslessCrossChainOrder memory) {
        SolverNet.OmniOrderData memory orderData = SolverNet.OmniOrderData({
            owner: owner,
            destChainId: destChainId,
            deposit: SolverNet.Deposit({ token: depositToken, amount: depositAmount }),
            calls: calls,
            expenses: expenses
        });

        IERC7683.GaslessCrossChainOrder memory order = IERC7683.GaslessCrossChainOrder({
            originSettler: address(inbox),
            user: user,
            nonce: nonce,
            originChainId: srcChainId,
            openDeadline: uint32(block.timestamp),
            fillDeadline: defaultFillDeadline,
            orderDataType: OMNIORDERDATA_TYPEHASH,
            orderData: abi.encode(orderData)
        });

        return (orderData, order);
    }

    function getNativeForNativeVaultOnchainOrder(uint96 depositAmount, uint96 expenseAmount)
        internal
        view
        returns (SolverNet.OrderData memory, IERC7683.OnchainCrossChainOrder memory)
    {
        SolverNet.Deposit memory deposit = SolverNet.Deposit({ token: address(0), amount: depositAmount });

        SolverNet.Call[] memory calls = new SolverNet.Call[](1);
        calls[0] = SolverNet.Call({
            target: address(nativeVault),
            selector: MockVault.deposit.selector,
            value: depositAmount,
            params: abi.encode(user, expenseAmount)
        });

        SolverNet.TokenExpense[] memory expenses = new SolverNet.TokenExpense[](0);

        SolverNet.OrderData memory orderData = SolverNet.OrderData({
            owner: address(0),
            destChainId: destChainId,
            deposit: deposit,
            calls: calls,
            expenses: expenses
        });

        IERC7683.OnchainCrossChainOrder memory order = IERC7683.OnchainCrossChainOrder({
            fillDeadline: defaultFillDeadline,
            orderDataType: ORDERDATA_TYPEHASH,
            orderData: abi.encode(orderData)
        });

        return (orderData, order);
    }

    function getNativeForNativeVaultGaslessOrder(uint256 nonce, uint96 depositAmount, uint96 expenseAmount)
        internal
        view
        returns (SolverNet.OmniOrderData memory, IERC7683.GaslessCrossChainOrder memory)
    {
        SolverNet.Deposit memory deposit = SolverNet.Deposit({ token: address(0), amount: depositAmount });

        SolverNet.Call[] memory calls = new SolverNet.Call[](1);
        calls[0] = SolverNet.Call({
            target: address(nativeVault),
            selector: MockVault.deposit.selector,
            value: depositAmount,
            params: abi.encode(user, expenseAmount)
        });

        SolverNet.TokenExpense[] memory expenses = new SolverNet.TokenExpense[](0);

        SolverNet.OmniOrderData memory orderData = SolverNet.OmniOrderData({
            owner: user,
            destChainId: destChainId,
            deposit: deposit,
            calls: calls,
            expenses: expenses
        });

        IERC7683.GaslessCrossChainOrder memory order = IERC7683.GaslessCrossChainOrder({
            originSettler: address(inbox),
            user: user,
            nonce: nonce,
            originChainId: srcChainId,
            openDeadline: uint32(block.timestamp),
            fillDeadline: defaultFillDeadline,
            orderDataType: OMNIORDERDATA_TYPEHASH,
            orderData: abi.encode(orderData)
        });

        return (orderData, order);
    }

    function getErc20ForErc20VaultOnchainOrder(uint96 depositAmount, uint96 expenseAmount)
        internal
        view
        returns (SolverNet.OrderData memory, IERC7683.OnchainCrossChainOrder memory)
    {
        SolverNet.Deposit memory deposit = SolverNet.Deposit({ token: address(token1), amount: depositAmount });

        SolverNet.Call[] memory calls = new SolverNet.Call[](1);
        calls[0] = SolverNet.Call({
            target: address(erc20Vault),
            selector: MockVault.deposit.selector,
            value: 0,
            params: abi.encode(user, expenseAmount)
        });

        SolverNet.TokenExpense[] memory expenses = new SolverNet.TokenExpense[](1);
        expenses[0] =
            SolverNet.TokenExpense({ spender: address(erc20Vault), token: address(token2), amount: expenseAmount });

        SolverNet.OrderData memory orderData = SolverNet.OrderData({
            owner: address(0),
            destChainId: destChainId,
            deposit: deposit,
            calls: calls,
            expenses: expenses
        });

        IERC7683.OnchainCrossChainOrder memory order = IERC7683.OnchainCrossChainOrder({
            fillDeadline: defaultFillDeadline,
            orderDataType: ORDERDATA_TYPEHASH,
            orderData: abi.encode(orderData)
        });

        return (orderData, order);
    }

    function getErc20ForErc20VaultGaslessOrder(uint256 nonce, uint96 depositAmount, uint96 expenseAmount)
        internal
        view
        returns (SolverNet.OmniOrderData memory, IERC7683.GaslessCrossChainOrder memory)
    {
        SolverNet.Deposit memory deposit = SolverNet.Deposit({ token: address(token1), amount: depositAmount });

        SolverNet.Call[] memory calls = new SolverNet.Call[](1);
        calls[0] = SolverNet.Call({
            target: address(erc20Vault),
            selector: MockVault.deposit.selector,
            value: 0,
            params: abi.encode(user, expenseAmount)
        });

        SolverNet.TokenExpense[] memory expenses = new SolverNet.TokenExpense[](1);
        expenses[0] =
            SolverNet.TokenExpense({ spender: address(erc20Vault), token: address(token2), amount: expenseAmount });

        SolverNet.OmniOrderData memory orderData = SolverNet.OmniOrderData({
            owner: user,
            destChainId: destChainId,
            deposit: deposit,
            calls: calls,
            expenses: expenses
        });

        IERC7683.GaslessCrossChainOrder memory order = IERC7683.GaslessCrossChainOrder({
            originSettler: address(inbox),
            user: user,
            nonce: nonce,
            originChainId: srcChainId,
            openDeadline: uint32(block.timestamp),
            fillDeadline: defaultFillDeadline,
            orderDataType: OMNIORDERDATA_TYPEHASH,
            orderData: abi.encode(orderData)
        });

        return (orderData, order);
    }

    function getArbitraryVaultOnchainOrder(
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
                spender: address(erc20Vault),
                token: expenseTokens[i],
                amount: expenseAmounts[i]
            });
        }

        SolverNet.OrderData memory orderData = SolverNet.OrderData({
            owner: address(0),
            destChainId: destChainId,
            deposit: deposit,
            calls: calls,
            expenses: expenses
        });

        IERC7683.OnchainCrossChainOrder memory order = IERC7683.OnchainCrossChainOrder({
            fillDeadline: defaultFillDeadline,
            orderDataType: ORDERDATA_TYPEHASH,
            orderData: abi.encode(orderData)
        });

        return (orderData, order);
    }

    function getArbitraryVaultGaslessOrder(
        uint256 nonce,
        address depositToken,
        uint96 depositAmount,
        address[] memory expenseTokens,
        uint96[] memory expenseAmounts
    ) internal view returns (SolverNet.OmniOrderData memory, IERC7683.GaslessCrossChainOrder memory) {
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
                spender: address(erc20Vault),
                token: expenseTokens[i],
                amount: expenseAmounts[i]
            });
        }

        SolverNet.OmniOrderData memory orderData = SolverNet.OmniOrderData({
            owner: user,
            destChainId: destChainId,
            deposit: deposit,
            calls: calls,
            expenses: expenses
        });

        IERC7683.GaslessCrossChainOrder memory order = IERC7683.GaslessCrossChainOrder({
            originSettler: address(inbox),
            user: user,
            nonce: nonce,
            originChainId: srcChainId,
            openDeadline: uint32(block.timestamp),
            fillDeadline: defaultFillDeadline,
            orderDataType: OMNIORDERDATA_TYPEHASH,
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

    function etchPermit2() internal returns (IPermit2) {
        vm.etch(PERMIT2, PERMIT2_CODE);
        return IPermit2(PERMIT2);
    }

    function deploySolverNetInbox() internal returns (SolverNetInboxV2) {
        address impl = address(new SolverNetInboxV2());
        return SolverNetInboxV2(address(new TransparentUpgradeableProxy(impl, proxyAdmin, bytes(""))));
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
        outbox.initialize(address(this), solver, address(portal), address(executor));

        uint64[] memory chainIds = new uint64[](1);
        chainIds[0] = srcChainId;
        address[] memory inboxes = new address[](1);
        inboxes[0] = address(inbox);
        outbox.setInboxes(chainIds, inboxes);
    }

    function assertStatus(bytes32 orderId, ISolverNetInboxV2.Status status) internal view {
        (, ISolverNetInboxV2.OrderState memory state,) = inbox.getOrder(orderId);

        uint8 expect = uint8(status);
        uint8 actual = uint8(state.status);

        if (status == ISolverNetInboxV2.Status.Pending) assertEq(expect, actual, "order should be pending");
        if (status == ISolverNetInboxV2.Status.Claimed) assertEq(expect, actual, "order should be claimed");
        if (status == ISolverNetInboxV2.Status.Rejected) assertEq(expect, actual, "order should be rejected");
        if (status == ISolverNetInboxV2.Status.Closed) assertEq(expect, actual, "order should be closed");
        if (status == ISolverNetInboxV2.Status.Filled) assertEq(expect, actual, "order should be filled");
        if (status == ISolverNetInboxV2.Status.Invalid) revert("invalid status");
    }
}
