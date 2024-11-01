// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

library Solve {
    enum Status {
        Open,
        Accepted,
        Cancelled,
        Rejected,
        Fulfilled,
        Paid
    }

    /**
     * @notice A request to execute a call on another chain, backed by a deposit.
     * @param id            ID for the request, globally unique per inbox.
     * @param updatedAt     Timestamp request status was last updated.
     * @param from          Address of the user who created the request.
     * @param fulfilledBy   Address of the solver that fulfilled the request.
     * @param status        Request status (open, accepted, cancelled, rejected, fulfilled, paid).
     * @param call          Details of the call to be executed on another chain.
     * @param deposits      Array of deposits backing the request.
     */
    struct Request {
        bytes32 id;
        uint40 updatedAt;
        Status status;
        address from;
        address fulfilledBy;
        Call call;
        Deposit[] deposits;
    }

    /**
     * @notice Details of a call to be executed on another chain.
     * @param destChainId  ID of the destination chain.
     * @param value        Amount of native currency to send with the call.
     * @param target       Address of the target contract on the destination chain.
     * @param data         Encoded data to be sent with the call.
     */
    struct Call {
        uint64 destChainId;
        address target;
        uint256 value;
        bytes data;
    }

    /**
     * @notice Details of a deposit backing a request.
     * @param isNative  Whether the deposit is in native currency.
     * @param token     Address of the token, address(0) if native.
     * @param amount    Deposit amount.
     */
    struct Deposit {
        bool isNative;
        address token;
        uint256 amount;
    }

    /**
     * @notice Details of a token deposit backing a request.
     * @dev Not stored, only used in opening a request.
     * @param token  Address of the token.
     * @param amount Deposit amount.
     */
    struct TokenDeposit {
        address token;
        uint256 amount;
    }
}
