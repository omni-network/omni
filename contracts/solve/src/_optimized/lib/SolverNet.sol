// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

library SolverNet {
    struct OrderData {
        Header header;
        Deposit deposit;
        Call[] calls;
        Expense[] expenses;
    }

    struct Header {
        address owner;
        uint64 destChainId;
        uint32 fillDeadline;
    }

    struct Call {
        address target;
        bytes4 selector;
        uint256 value;
        bytes params;
    }

    struct Deposit {
        address token;
        uint96 amount;
    }

    struct Expense {
        address spender;
        address token;
        uint96 amount;
    }

    struct FillOriginData {
        uint64 srcChainId;
        uint64 destChainId;
        uint32 fillDeadline;
        Call[] calls;
        Expense[] expenses;
    }
}
