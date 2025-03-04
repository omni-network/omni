// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

library SolverNet {
    // OrderData describes an intent for a cross-chain action between a source and a destination chain.
    struct OrderData {
        address owner;
        uint64 destChainId;
        Deposit deposit;
        Call[] calls;
        TokenExpense[] expenses;
    }

    struct Order {
        Header header;
        Deposit deposit;
        Call[] calls;
        TokenExpense[] expenses;
    }

    struct Header {
        address owner;
        uint64 destChainId;
        uint32 fillDeadline;
    }

    // Deposit contains the token amount `X + fee` with `X` being the amount
    // that needs to be transferred/deposited on the destination chain and `fee`
    // being the incentive paid to the solver for the intent execution.
    // The deposit is paid by the user on the source chain.
    struct Deposit {
        address token;
        uint96 amount;
    }

    // Call encodes a contract call or a native transfer that the solver will execute on the
    // destination chain. For native ETH transfers, the call is a simple transfer and its
    // `value` field denotes the amount to be transferred.
    struct Call {
        address target;
        bytes4 selector;
        uint256 value;
        bytes params;
    }

    // TokenExpense contains token amounts the solver needs to have on the destination chain balance
    // when executing `calls`. For a native ETH transfers, expenses are not needed and are
    // inferred from the `calls` field of the order data.
    struct TokenExpense {
        address spender;
        address token;
        uint96 amount;
    }

    struct FillOriginData {
        uint64 srcChainId;
        uint64 destChainId;
        uint32 fillDeadline;
        Call[] calls;
        TokenExpense[] expenses;
    }
}
