// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.24;

import { IERC7683 } from "../erc7683/IERC7683.sol";
import { SolverNet } from "./SolverNet.sol";

library HashLib {
    /**
     * @notice Type for the Deposit struct.
     */
    bytes internal constant DEPOSIT_TYPE = abi.encodePacked("Deposit(address token,uint96 amount)");

    /**
     * @notice Typehash for the Deposit struct.
     */
    bytes32 internal constant DEPOSIT_TYPEHASH = keccak256(DEPOSIT_TYPE);

    /**
     * @notice Type for the Call struct.
     */
    bytes internal constant CALL_TYPE =
        abi.encodePacked("Call(address target,bytes4 selector,uint256 value,bytes params)");

    /**
     * @notice Typehash for the Call struct.
     */
    bytes32 internal constant CALL_TYPEHASH = keccak256(CALL_TYPE);

    /**
     * @notice Type for the TokenExpense struct.
     */
    bytes internal constant TOKENEXPENSE_TYPE =
        abi.encodePacked("TokenExpense(address spender,address token,uint96 amount)");

    /**
     * @notice Typehash for the TokenExpense struct.
     */
    bytes32 internal constant TOKENEXPENSE_TYPEHASH = keccak256(TOKENEXPENSE_TYPE);

    /**
     * @notice Type for the OrderData struct.
     */
    bytes internal constant ORDERDATA_TYPE = abi.encodePacked(
        "OrderData(address owner,uint64 destChainId,Deposit deposit,Call[] calls,TokenExpense[] expenses)"
    );

    /**
     * @notice Type for the OmniOrderData struct.
     */
    bytes internal constant OMNIORDERDATA_TYPE = abi.encodePacked(
        "OmniOrderData(address owner,uint64 destChainId,Deposit deposit,Call[] calls,TokenExpense[] expenses)"
    );

    /**
     * @notice Previous full EIP-712 type for the OrderData struct.
     * @dev Included to maintain backwards compatibility.
     */
    bytes internal constant FULL_ORDERDATA_TYPE =
        abi.encodePacked(ORDERDATA_TYPE, DEPOSIT_TYPE, CALL_TYPE, TOKENEXPENSE_TYPE);

    /**
     * @notice Full EIP-712 type for the OmniOrderData struct.
     */
    bytes internal constant FULL_OMNIORDERDATA_TYPE =
        abi.encodePacked(OMNIORDERDATA_TYPE, CALL_TYPE, DEPOSIT_TYPE, TOKENEXPENSE_TYPE);

    /**
     * @notice Old typehash for the full OrderData struct.
     * @dev Included to maintain backwards compatibility.
     */
    bytes32 internal constant FULL_ORDERDATA_TYPEHASH = keccak256(FULL_ORDERDATA_TYPE);

    /**
     * @notice Typehash for the full OmniOrderData struct.
     */
    bytes32 internal constant FULL_OMNIORDERDATA_TYPEHASH = keccak256(FULL_OMNIORDERDATA_TYPE);

    /**
     * @notice Type for the GaslessCrossChainOrder struct.
     */
    bytes internal constant GASLESS_ORDER_TYPE = abi.encodePacked(
        "GaslessCrossChainOrder(address originSettler,address user,uint256 nonce,uint256 originChainId,uint32 openDeadline,uint32 fillDeadline,bytes32 orderDataType,OmniOrderData orderData)"
    );

    /**
     * @notice Full EIP-712 type for the GaslessCrossChainOrder struct.
     * @dev "Omni" prefix is added so wallets can show that the signature is for an Omni order.
     */
    bytes internal constant EIP712_GASLESS_ORDER_TYPE =
        abi.encodePacked(GASLESS_ORDER_TYPE, CALL_TYPE, DEPOSIT_TYPE, OMNIORDERDATA_TYPE, TOKENEXPENSE_TYPE);

    /**
     * @notice EIP-712 typehash for the GaslessCrossChainOrder struct.
     */
    bytes32 internal constant EIP712_GASLESS_ORDER_TYPEHASH = keccak256(EIP712_GASLESS_ORDER_TYPE);

    /**
     * @notice Type for the PermitWitnessTransferFrom struct.
     */
    string internal constant PERMIT_WITNESS_TRANSFER_FROM_TYPE_STUB =
        "PermitWitnessTransferFrom(TokenPermissions permitted,address spender,uint256 nonce,uint256 deadline,";

    /**
     * @notice Type for the TokenPermissions struct.
     */
    string internal constant TOKEN_PERMISSIONS_TYPE = "TokenPermissions(address token,uint256 amount)";

    /**
     * @notice Type for the Permit2 witness.
     */
    string internal constant PERMIT2_WITNESS_TYPE = "GaslessCrossChainOrder witness)";

    /**
     * @notice Full EIP-712 type for the Permit2 order.
     * @dev "Omni" prefix is added so wallets can show that the signature is for an Omni order.
     */
    string internal constant PERMIT2_ORDER_TYPE = string(
        abi.encodePacked(
            PERMIT2_WITNESS_TYPE,
            CALL_TYPE,
            DEPOSIT_TYPE,
            GASLESS_ORDER_TYPE,
            OMNIORDERDATA_TYPE,
            TOKENEXPENSE_TYPE,
            TOKEN_PERMISSIONS_TYPE
        )
    );

    /**
     * @dev Hashes a deposit.
     * @param deposit Deposit to hash.
     * @return _ Hashed deposit.
     */
    function hashDeposit(SolverNet.Deposit memory deposit) internal pure returns (bytes32) {
        return keccak256(abi.encode(DEPOSIT_TYPEHASH, deposit.token, deposit.amount));
    }

    /**
     * @dev Hashes a call.
     * @param call Call to hash.
     * @return _ Hashed call.
     */
    function hashCall(SolverNet.Call memory call) internal pure returns (bytes32) {
        return keccak256(abi.encode(CALL_TYPEHASH, call.target, call.selector, call.value, keccak256(call.params)));
    }

    /**
     * @dev Hashes an array of calls.
     * @param calls Calls to hash.
     * @return _ Hashed calls.
     */
    function hashCalls(SolverNet.Call[] memory calls) internal pure returns (bytes32) {
        bytes32[] memory callHashes = new bytes32[](calls.length);
        for (uint256 i; i < calls.length; ++i) {
            callHashes[i] = hashCall(calls[i]);
        }
        return keccak256(abi.encodePacked(callHashes));
    }

    /**
     * @dev Hashes a token expense.
     * @param expense Token expense to hash.
     * @return _ Hashed token expense.
     */
    function hashTokenExpense(SolverNet.TokenExpense memory expense) internal pure returns (bytes32) {
        return keccak256(abi.encode(TOKENEXPENSE_TYPEHASH, expense.spender, expense.token, expense.amount));
    }

    /**
     * @dev Hashes an array of token expenses.
     * @param expenses Token expenses to hash.
     * @return _ Hashed token expenses.
     */
    function hashTokenExpenses(SolverNet.TokenExpense[] memory expenses) internal pure returns (bytes32) {
        bytes32[] memory expenseHashes = new bytes32[](expenses.length);
        for (uint256 i; i < expenses.length; ++i) {
            expenseHashes[i] = hashTokenExpense(expenses[i]);
        }
        return keccak256(abi.encodePacked(expenseHashes));
    }

    /**
     * @dev Hashes the order data.
     * @param orderData Order data to hash.
     * @return _ Hashed order data.
     */
    function hashOrderData(SolverNet.OmniOrderData memory orderData) internal pure returns (bytes32) {
        return keccak256(
            abi.encode(
                FULL_OMNIORDERDATA_TYPEHASH,
                orderData.owner,
                orderData.destChainId,
                hashDeposit(orderData.deposit),
                hashCalls(orderData.calls),
                hashTokenExpenses(orderData.expenses)
            )
        );
    }

    /**
     * @dev Hashes a gasless order.
     * @param order GaslessCrossChainOrder to hash.
     * @param orderDataHash Hashed order data.
     * @return _ Hashed gasless order.
     */
    function hashGaslessOrder(IERC7683.GaslessCrossChainOrder memory order, bytes32 orderDataHash)
        internal
        pure
        returns (bytes32)
    {
        return keccak256(
            abi.encode(
                EIP712_GASLESS_ORDER_TYPEHASH,
                order.originSettler,
                order.user,
                order.nonce,
                order.originChainId,
                order.openDeadline,
                order.fillDeadline,
                order.orderDataType,
                orderDataHash
            )
        );
    }
}
