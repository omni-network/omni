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
     * @param target    Target contract address.
     * @param value     Value to send to the target.
     * @param data      Calldata to send to the target.
     */
    struct Call {
        bytes32 target;
        uint256 value;
        bytes data;
    }

    /**
     * @notice SolverNet ERC-7683 order data.
     *         Restricted to single call on a destination chain.
     * @param owner         Address allowed to cancel the order. address(0) for msg.sender of inbox.open(...)
     * @param destChainId   Chain ID on which the order needs to be filled.
     * @param calls         Calls to execute on.
     * @param deposits      Deposits payed by user, locked on source chain. Awarded to solver on fill.
     * @param expenses      Expenses required to fund the calls.
     */
    struct OrderData {
        address owner;
        uint64 destChainId;
        Call[] calls;
        Deposit[] deposits;
        TokenExpense[] expenses;
    }

    /**
     * @notice SolverNet ERC-7683 fill instruction origin data.
     * @param srcChainId    Chain ID on which the order was opened.
     * @param destChainId   Chain ID on which the order needs to be filled.
     * @param fillDeadline  Deadline for the fill.
     * @param calls         Calls to execute.
     * @param expenses      Expenses required to fund the calls.
     */
    struct FillOriginData {
        uint64 srcChainId;
        uint64 destChainId;
        uint40 fillDeadline;
        Call[] calls;
        TokenExpense[] expenses;
    }
}
