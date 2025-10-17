// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.24;

import { IERC7683 } from "../erc7683/IERC7683.sol";
import { IPermit2 } from "@uniswap/permit2/src/interfaces/IPermit2.sol";
import { SolverNet } from "./SolverNet.sol";

library HashLib {
    /**
     * @notice Typehash for the OrderData struct.
     * @dev Deprecated typehash for the old OrderData struct, which is being replaced with OmniOrderData for more clarity in EIP-712 signing.
     */
    bytes32 internal constant OLD_ORDERDATA_TYPEHASH = keccak256(
        "OrderData(address owner,uint64 destChainId,Deposit deposit,Call[] calls,TokenExpense[] expenses)Deposit(address token,uint96 amount)Call(address target,bytes4 selector,uint256 value,bytes params)TokenExpense(address spender,address token,uint96 amount)"
    );

    /**
     * @notice Typehash for the OmniOrderData struct.
     * @dev Used for more clarity in EIP-712 signing.
     */
    bytes32 internal constant OMNIORDERDATA_TYPEHASH = keccak256(
        "OmniOrderData(address owner,uint64 destChainId,Deposit deposit,Call[] calls,TokenExpense[] expenses)Call(address target,bytes4 selector,uint256 value,bytes params)Deposit(address token,uint96 amount)TokenExpense(address spender,address token,uint96 amount)"
    );

    /**
     * @notice Typehash for the GaslessCrossChainOrder struct for witness data with structured OmniOrderData.
     */
    bytes32 internal constant GASLESS_ORDER_TYPEHASH = keccak256(
        "GaslessCrossChainOrder(address originSettler,address user,uint256 nonce,uint256 originChainId,uint32 openDeadline,uint32 fillDeadline,bytes32 orderDataType,OmniOrderData orderData)Call(address target,bytes4 selector,uint256 value,bytes params)Deposit(address token,uint96 amount)OmniOrderData(address owner,uint64 destChainId,Deposit deposit,Call[] calls,TokenExpense[] expenses)TokenExpense(address spender,address token,uint96 amount)"
    );

    /**
     * @notice Typehash for the Call struct.
     */
    bytes32 internal constant CALL_TYPEHASH =
        keccak256("Call(address target,bytes4 selector,uint256 value,bytes params)");

    /**
     * @notice Typehash for the Deposit struct.
     */
    bytes32 internal constant DEPOSIT_TYPEHASH = keccak256("Deposit(address token,uint96 amount)");

    /**
     * @notice Typehash for the TokenExpense struct.
     */
    bytes32 internal constant TOKEN_EXPENSE_TYPEHASH =
        keccak256("TokenExpense(address spender,address token,uint96 amount)");

    /**
     * @notice Type string for permit2 witness transfers.
     */
    string internal constant PERMIT2_TRANSFER_TYPE_STRING =
        "PermitWitnessTransferFrom(TokenPermissions permitted,address spender,uint256 nonce,uint256 deadline,";

    /**
     * @notice Type string for permit2 witness data with structured OmniOrderData.
     */
    string internal constant PERMIT2_WITNESS_TYPE_STRING =
        "GaslessCrossChainOrder witness)Call(address target,bytes4 selector,uint256 value,bytes params)Deposit(address token,uint96 amount)GaslessCrossChainOrder(address originSettler,address user,uint256 nonce,uint256 originChainId,uint32 openDeadline,uint32 fillDeadline,bytes32 orderDataType,OmniOrderData orderData)OmniOrderData(address owner,uint64 destChainId,Deposit deposit,Call[] calls,TokenExpense[] expenses)TokenExpense(address spender,address token,uint96 amount)TokenPermissions(address token,uint256 amount)";

    /**
     * @notice Typehash for the TokenPermissions struct.
     */
    bytes32 internal constant PERMIT2_TOKEN_PERMISSIONS_TYPEHASH =
        keccak256("TokenPermissions(address token,uint256 amount)");

    /**
     * @notice Typehash for the full permit2 witness transfer with structured OmniOrderData.
     */
    bytes32 internal constant PERMIT2_WITNESS_TYPEHASH =
        keccak256(abi.encodePacked(PERMIT2_TRANSFER_TYPE_STRING, PERMIT2_WITNESS_TYPE_STRING));

    /**
     * @notice Canonical permit2 contract address.
     */
    IPermit2 internal constant permit2 = IPermit2(0x000000000022D473030F116dDEE9F6B43aC78BA3);

    /**
     * @notice Generate witness hash for permit2 witness transfers using structured OrderData.
     * @param order GaslessCrossChainOrder to hash as witness data.
     * @param orderData Decoded OrderData from order.orderData.
     * @return witnessHash The EIP-712 compliant hash of the gasless order with structured data.
     */
    function witnessHash(IERC7683.GaslessCrossChainOrder memory order, SolverNet.OrderData memory orderData)
        internal
        pure
        returns (bytes32)
    {
        return keccak256(
            abi.encode(
                GASLESS_ORDER_TYPEHASH,
                order.originSettler,
                order.user,
                order.nonce,
                order.originChainId,
                order.openDeadline,
                order.fillDeadline,
                order.orderDataType,
                hashOmniOrderData(orderData)
            )
        );
    }

    /**
     * @notice Generate witness hash for permit2 witness transfers using structured OrderData.
     * @dev This function is used with calldata types for efficiency. (decoded OrderData is never stored as calldata)
     * @param order GaslessCrossChainOrder to hash as witness data.
     * @param orderData Decoded OrderData from order.orderData.
     * @return witnessHash The EIP-712 compliant hash of the gasless order with structured data.
     */
    function witnessHashCalldata(IERC7683.GaslessCrossChainOrder calldata order, SolverNet.OrderData memory orderData)
        internal
        pure
        returns (bytes32)
    {
        return keccak256(
            abi.encode(
                GASLESS_ORDER_TYPEHASH,
                order.originSettler,
                order.user,
                order.nonce,
                order.originChainId,
                order.openDeadline,
                order.fillDeadline,
                order.orderDataType,
                hashOmniOrderData(orderData)
            )
        );
    }

    /**
     * @notice Hash an OrderData struct for EIP-712 compliance.
     * @param orderData OrderData to hash.
     * @return hash The EIP-712 compliant hash of the OrderData.
     */
    function hashOmniOrderData(SolverNet.OrderData memory orderData) internal pure returns (bytes32) {
        return keccak256(
            abi.encode(
                OMNIORDERDATA_TYPEHASH,
                orderData.owner,
                orderData.destChainId,
                hashDeposit(orderData.deposit),
                hashCalls(orderData.calls),
                hashTokenExpenses(orderData.expenses)
            )
        );
    }

    /**
     * @notice Hash a Deposit struct for EIP-712 compliance.
     * @param deposit Deposit to hash.
     * @return hash The EIP-712 compliant hash of the Deposit.
     */
    function hashDeposit(SolverNet.Deposit memory deposit) internal pure returns (bytes32) {
        return keccak256(abi.encode(DEPOSIT_TYPEHASH, deposit.token, deposit.amount));
    }

    /**
     * @notice Hash an array of Call structs for EIP-712 compliance.
     * @param calls Array of calls to hash.
     * @return hash The EIP-712 compliant hash of the calls array.
     */
    function hashCalls(SolverNet.Call[] memory calls) internal pure returns (bytes32) {
        // Deterministic hash for empty array
        if (calls.length == 0) {
            return keccak256("");
        }

        bytes32[] memory callHashes = new bytes32[](calls.length);
        for (uint256 i; i < calls.length; ++i) {
            callHashes[i] = keccak256(
                abi.encode(
                    CALL_TYPEHASH, calls[i].target, calls[i].selector, calls[i].value, keccak256(calls[i].params)
                )
            );
        }
        return keccak256(abi.encodePacked(callHashes));
    }

    /**
     * @notice Hash an array of TokenExpense structs for EIP-712 compliance.
     * @param expenses Array of token expenses to hash.
     * @return hash The EIP-712 compliant hash of the expenses array.
     */
    function hashTokenExpenses(SolverNet.TokenExpense[] memory expenses) internal pure returns (bytes32) {
        // Deterministic hash for empty array
        if (expenses.length == 0) {
            return keccak256("");
        }

        bytes32[] memory expenseHashes = new bytes32[](expenses.length);
        for (uint256 i; i < expenses.length; ++i) {
            expenseHashes[i] = keccak256(
                abi.encode(TOKEN_EXPENSE_TYPEHASH, expenses[i].spender, expenses[i].token, expenses[i].amount)
            );
        }
        return keccak256(abi.encodePacked(expenseHashes));
    }

    /**
     * @notice Generate digest for gasless order with structured OrderData.
     * @param order GaslessCrossChainOrder to hash as witness data.
     * @param orderData Decoded OrderData from order.orderData.
     * @param inbox Inbox contract address.
     * @return digest The EIP-712 compliant hash of the gasless order with structured data.
     */
    function gaslessOrderDigest(
        IERC7683.GaslessCrossChainOrder memory order,
        SolverNet.OrderData memory orderData,
        address inbox
    ) internal view returns (bytes32) {
        // Hash TokenPermissions
        bytes32 tokenPermissionsHash = keccak256(
            abi.encode(PERMIT2_TOKEN_PERMISSIONS_TYPEHASH, orderData.deposit.token, orderData.deposit.amount)
        );

        // Create final struct hash
        bytes32 structHash = keccak256(
            abi.encode(
                PERMIT2_WITNESS_TYPEHASH,
                tokenPermissionsHash,
                inbox,
                order.nonce,
                order.openDeadline,
                witnessHash(order, orderData)
            )
        );

        // Create EIP-712 hash
        bytes32 hash = keccak256(abi.encodePacked("\x19\x01", permit2.DOMAIN_SEPARATOR(), structHash));

        return hash;
    }

    /**
     * @notice Generate digest for gasless order with structured OrderData.
     * @dev This function is used with calldata types for efficiency. (decoded OrderData is never stored as calldata)
     * @param order GaslessCrossChainOrder to hash as witness data.
     * @param orderData Decoded OrderData from order.orderData.
     * @param inbox Inbox contract address.
     * @return digest The EIP-712 compliant hash of the gasless order with structured data.
     */
    function gaslessOrderDigestCalldata(
        IERC7683.GaslessCrossChainOrder calldata order,
        SolverNet.OrderData memory orderData,
        address inbox
    ) internal view returns (bytes32) {
        // Hash TokenPermissions
        bytes32 tokenPermissionsHash = keccak256(
            abi.encode(PERMIT2_TOKEN_PERMISSIONS_TYPEHASH, orderData.deposit.token, orderData.deposit.amount)
        );

        // Create final struct hash
        bytes32 structHash = keccak256(
            abi.encode(
                PERMIT2_WITNESS_TYPEHASH,
                tokenPermissionsHash,
                inbox,
                order.nonce,
                order.openDeadline,
                witnessHashCalldata(order, orderData)
            )
        );

        // Create EIP-712 hash
        bytes32 hash = keccak256(abi.encodePacked("\x19\x01", permit2.DOMAIN_SEPARATOR(), structHash));

        return hash;
    }
}
