// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.24;

import { IERC7683 } from "../erc7683/IERC7683.sol";
import { IPermit2 } from "@uniswap/permit2/src/interfaces/IPermit2.sol";
import { SolverNet } from "./SolverNet.sol";

library HashLibV2 {
    /**
     * @notice Typehash for the OrderData struct.
     */
    bytes32 internal constant OLD_ORDERDATA_TYPEHASH = keccak256(
        "OrderData(address owner,uint64 destChainId,Deposit deposit,Call[] calls,TokenExpense[] expenses)Deposit(address token,uint96 amount)Call(address target,bytes4 selector,uint256 value,bytes params)TokenExpense(address spender,address token,uint96 amount)"
    );

    /**
     * @notice Typehash for the OmniOrderData struct.
     */
    bytes32 internal constant OMNIORDERDATA_TYPEHASH = keccak256(
        "OmniOrderData(address owner,uint64 destChainId,Deposit deposit,Call[] calls,TokenExpense[] expenses)Call(address target,bytes4 selector,uint256 value,bytes params)Deposit(address token,uint96 amount)TokenExpense(address spender,address token,uint96 amount)"
    );

    /**
     * @notice Typehash for the GaslessCrossChainOrder struct for witness data.
     */
    bytes32 internal constant GASLESS_ORDER_TYPEHASH = keccak256(
        "GaslessCrossChainOrder(address originSettler,address user,uint256 nonce,uint256 originChainId,uint32 openDeadline,uint32 fillDeadline,bytes32 orderDataType,bytes orderData)"
    );

    /**
     * @notice Type string for permit2 witness transfers.
     */
    string internal constant PERMIT2_TRANSFER_TYPE_STRING =
        "PermitWitnessTransferFrom(TokenPermissions permitted,address spender,uint256 nonce,uint256 deadline,";

    /**
     * @notice Type string for permit2 witness data.
     */
    string internal constant PERMIT2_WITNESS_TYPE_STRING =
        "GaslessCrossChainOrder witness)GaslessCrossChainOrder(address originSettler,address user,uint256 nonce,uint256 originChainId,uint32 openDeadline,uint32 fillDeadline,bytes32 orderDataType,bytes orderData)TokenPermissions(address token,uint256 amount)";

    /**
     * @notice Typehash for the TokenPermissions struct.
     */
    bytes32 internal constant PERMIT2_TOKEN_PERMISSIONS_TYPEHASH =
        keccak256("TokenPermissions(address token,uint256 amount)");

    /**
     * @notice Typehash for the full permit2 witness transfer.
     */
    bytes32 internal constant PERMIT2_WITNESS_TYPEHASH =
        keccak256(abi.encodePacked(PERMIT2_TRANSFER_TYPE_STRING, PERMIT2_WITNESS_TYPE_STRING));

    IPermit2 internal constant permit2 = IPermit2(0x000000000022D473030F116dDEE9F6B43aC78BA3);

    /**
     * @notice Generate witness hash for permit2 witness transfers using EIP-712 structured data.
     * @param order GaslessCrossChainOrder to hash as witness data.
     * @return witnessHash The EIP-712 compliant hash of the gasless order.
     */
    function witnessHash(IERC7683.GaslessCrossChainOrder memory order) internal pure returns (bytes32) {
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
                keccak256(order.orderData)
            )
        );
    }

    /**
     * @notice Generate witness hash for permit2 witness transfers using EIP-712 structured data.
     * @dev This function is used with calldata types for efficiency.
     * @param order GaslessCrossChainOrder to hash as witness data.
     * @return witnessHash The EIP-712 compliant hash of the gasless order.
     */
    function witnessHashCalldata(IERC7683.GaslessCrossChainOrder calldata order) internal pure returns (bytes32) {
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
                keccak256(order.orderData)
            )
        );
    }

    /**
     * @notice Generate digest for gasless order.
     * @param order GaslessCrossChainOrder to hash as witness data.
     * @param deposit Deposit to hash as witness data.
     * @param inbox Inbox contract address.
     * @return digest The EIP-712 compliant hash of the gasless order.
     */
    function gaslessOrderDigest(
        IERC7683.GaslessCrossChainOrder memory order,
        SolverNet.Deposit memory deposit,
        address inbox
    ) internal view returns (bytes32) {
        // Hash TokenPermissions
        bytes32 tokenPermissionsHash =
            keccak256(abi.encode(PERMIT2_TOKEN_PERMISSIONS_TYPEHASH, deposit.token, deposit.amount));

        // Create final struct hash
        bytes32 structHash = keccak256(
            abi.encode(
                PERMIT2_WITNESS_TYPEHASH,
                tokenPermissionsHash,
                inbox,
                order.nonce,
                order.openDeadline,
                witnessHash(order)
            )
        );

        // Create EIP-712 hash
        bytes32 hash = keccak256(abi.encodePacked("\x19\x01", permit2.DOMAIN_SEPARATOR(), structHash));

        return hash;
    }

    /**
     * @notice Generate digest for gasless order.
     * @dev This function is used with calldata types for efficiency.
     * @param order GaslessCrossChainOrder to hash as witness data.
     * @param deposit Deposit to hash as witness data.
     * @param inbox Inbox contract address.
     * @return digest The EIP-712 compliant hash of the gasless order.
     */
    function gaslessOrderDigestCalldata(
        IERC7683.GaslessCrossChainOrder calldata order,
        SolverNet.Deposit calldata deposit,
        address inbox
    ) internal view returns (bytes32) {
        // Hash TokenPermissions
        bytes32 tokenPermissionsHash =
            keccak256(abi.encode(PERMIT2_TOKEN_PERMISSIONS_TYPEHASH, deposit.token, deposit.amount));

        // Create final struct hash
        bytes32 structHash = keccak256(
            abi.encode(
                PERMIT2_WITNESS_TYPEHASH,
                tokenPermissionsHash,
                inbox,
                order.nonce,
                order.openDeadline,
                witnessHashCalldata(order)
            )
        );

        // Create EIP-712 hash
        bytes32 hash = keccak256(abi.encodePacked("\x19\x01", permit2.DOMAIN_SEPARATOR(), structHash));

        return hash;
    }
}
