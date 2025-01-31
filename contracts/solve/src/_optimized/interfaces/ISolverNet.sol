// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

interface ISolverNet {
    struct OrderData {
        address owner;
        Call call;
        Values values;
        Deposit deposit;
        Expense[] expenses;
    }

    struct Call {
        uint64 chainId;
        address target;
        bytes4 selector;
        bytes callParams;
    }

    struct Values {
        uint96 nativeTip;
        uint96 callValue;
        uint32 openDeadline;
        uint32 fillDeadline;
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
        uint96 callValue;
        address target;
        bytes callData;
        Expense[] expenses;
    }
}
