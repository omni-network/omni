// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

library Solve {
    /**
     * @notice Status of a request.
     */
    enum Status {
        Invalid,
        Pending,
        Accepted,
        Rejected,
        Reverted,
        Fulfilled,
        Claimed
    }

    /**
     * @notice Details of a status update.
     * @param status    Status of the request.
     * @param timestamp Timestamp of the status update.
     */
    struct StatusUpdate {
        Status status;
        uint40 timestamp;
    }

    /**
     * @notice A request to execute a call on another chain, backed by a deposit.
     * @param id            ID for the request, globally unique per inbox.
     * @param updatedAt     Timestamp request status was last updated.
     * @param from          Address of the user who created the request.
     * @param acceptedBy    Address of the solver that accepted the request.
     * @param status        Request status (open, accepted, cancelled, rejected, fulfilled, paid).
     * @param call          Details of the call to be executed on another chain.
     * @param deposits      Array of deposits backing the request.
     * @param updateHistory Array of status updates including timestamps.
     */
    struct Request {
        bytes32 id;
        uint40 updatedAt;
        Status status;
        address from;
        address acceptedBy;
        Call call;
        Deposit[] deposits;
        StatusUpdate[] updateHistory;
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

    /**
     * @notice Details of a token pre-requisite for a call.
     * @dev Not stored, only used in opening a request.
     * @param token    Address of the token.
     * @param spender  Address of the spender.
     * @param amount   Transfer and approval amount.
     */
    struct TokenPrereq {
        address token;
        address spender;
        uint256 amount;
    }
}
