// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

library SolverNet {
    /**
     * @notice OrderData is the SolverNet's ERC7683 order data encoding, used when opening an order.
     *  It describes the calls to execute, the expenses required, and the deposit backing the order.
     *
     *  Note that expenses only includes ERC20 expenses, native expenses are inferred from the calls.
     *  This reduces the order data size, saving gas.
     *
     *  Order expenses (both token expenses and required call values) should be exact. If expenses are underestimated,
     *  the order will be rejected. If expenses are overestimated, the excess will be refunded to the solver.
     *
     *  Order deposits must cover expenses + solver fees. If they do not, the order will be rejected. To
     *  determine required deposits, use Omni's solver /quote API, or Omni's SolverNet SDK. (TODO link to docs)
     *
     * @custom:field owner          The address that can close the order. Defaults to msg.sender if zero.
     * @custom:field destChainId    The chain on which to execute calls.
     * @custom:field deposit        The deposit paid by the user on the source chain when opening the order.
     * @custom:field calls          The calls to be executed on the destination chain.
     * @custom:field expenses       The token expenses required to fill the order, paid by the solver.
     */
    struct OrderData {
        address owner;
        uint64 destChainId;
        Deposit deposit;
        Call[] calls;
        TokenExpense[] expenses;
    }

    /**
     * @notice Order is a convenience struct that fully describes an order.
     * @dev It is not written to storage, but rather built on view.
     */
    struct Order {
        Header header;
        Deposit deposit;
        Call[] calls;
        TokenExpense[] expenses;
    }

    /**
     * @notice Header is an "order header", containing order data not related to deposit, calls, or expenses.
     * @dev This struct is used to pack loose order data fields into a single slot.
     */
    struct Header {
        address owner;
        uint64 destChainId;
        uint32 fillDeadline;
    }

    /**
     * @notice Deposit describes a deposit (native or ERC20) paid by the user on the source chain when opening an order.
     * @custom:field token   The address of the token, address(0) if native.
     * @custom:field amount  The amount of the token deposited.
     */
    struct Deposit {
        address token;
        uint96 amount;
    }

    /**
     * @notice Call describes a call to execute.
     *  If the call is a native transfer, `target` is the recipient address, and `selector` / `params` are empty.
     * @dev Full call data is built from `selector` and `params`: abi.encodePacked(call.selector, call.params).
     * @custom:field target    The address to execute the call against
     * @custom:field selector  The function selector to call
     * @custom:field value     The amount of native token to send with the call
     * @custom:field params    The call parameters
     */
    struct Call {
        address target;
        bytes4 selector;
        uint256 value;
        bytes params;
    }

    /**
     * @notice TokenExpense describes an ERC20 expense to be paid by the solver on destination chain when filling an
     *  order. Native expenses are inferred from the calls, and are not included in the order data.
     * @custom:field spender  The address that will do token.transferFrom(...) on fill. Required to set allowance.
     * @custom:field token    The address of the token on the destination chain.
     * @custom:field amount   The amount of the token to be spent.
     */
    struct TokenExpense {
        address spender;
        address token;
        uint96 amount;
    }

    /**
     * @notice FillOriginData is the SolverNet's ERC7683 fill instruction data encoding, used when filling an order.
     * @dev The hash(orderId, fillOriginData) is used to prove that a solver properly filled an order.
     * @custom:field srcChainId     The chain on which the order was opened.
     * @custom:field destChainId    The chain on which to execute calls.
     * @custom:field fillDeadline   The deadline by which the order must be filled.
     * @custom:field calls          The calls to execute on the destination chain.
     * @custom:field expenses       The token expenses required to fill the order, paid by the solver.
     */
    struct FillOriginData {
        uint64 srcChainId;
        uint64 destChainId;
        uint32 fillDeadline;
        Call[] calls;
        TokenExpense[] expenses;
    }
}
