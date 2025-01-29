// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

interface ISolverNet {
    /**
     * @notice Deposit backing an order.
     * @param token   Token address, 0x0 for native.
     * @param amount  Token amount.
     */
    struct Deposit {
        bytes32 token;
        uint256 amount;
    }

    /**
     * @notice Token expense required to fill an order.
     * @param token    Token address
     * @param spender  Address spending the token.
     * @param amount   Token amount.
     */
    struct TokenExpense {
        bytes32 token;
        bytes32 spender;
        uint256 amount;
    }

    /**
     * @notice Call to execute on a destination chain.
     * @param chainId   The ID of the chain on which the call should be executed.
     * @param target    Target contract address.
     * @param value     Value to send to the target.
     * @param data      Calldata to send to the target.
     * @param expenses  Expenses required to fund the call.
     */
    struct Call {
        uint64 chainId;
        bytes32 target;
        uint256 value;
        bytes data;
        TokenExpense[] expenses;
    }

    /**
     * @notice SolverNet ERC-7683 order data.
     *         Restricted to single call on a destination chain.
     * @param owner     Address allowed to cancel the order. address(0) for msg.sender of inbox.open(...)
     * @param call      Call to execute on.
     * @param deposits  Deposits payed by user, locked on source chain. Awarded to solver on fill.
     */
    struct OrderData {
        address owner;
        Call call;
        Deposit[] deposits;
    }

    /**
     * @notice SolverNet ERC-7683 fill instruction origin data.
     * @param srcChainId    Chain ID on which the order was opened.
     * @param fillDeadline  Deadline for the fill.
     * @param call          Call to execute
     */
    struct FillOriginData {
        uint64 srcChainId;
        uint40 fillDeadline;
        Call call;
    }
}
