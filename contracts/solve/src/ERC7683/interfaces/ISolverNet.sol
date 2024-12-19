// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

interface ISolverNet {
    /**
     * @notice ERC20 pre-requisite for an order to be filled on another chain.
     * @param token    Address of the token.
     * @param spender  Address authorized to spend the token on the destination chain.
     * @param amount   Amount of the token to be approved.
     */
    struct TokenPrereq {
        bytes32 token;
        bytes32 spender;
        uint256 amount;
    }

    /**
     * @notice Details of a call to be executed on another chain.
     * @param target      Address of the target contract on the destination chain.
     * @param value       Amount of native currency to send with the call.
     * @param callData    Encoded data to be sent with the call.
     */
    struct Call {
        bytes32 target;
        uint256 value;
        bytes callData;
    }

    /**
     * @notice Data to be sent to the destination chain to parameterize a fill.
     * @param srcChainId    ID of the source chain.
     * @param destChainId   ID of the destination chain.
     * @param tokenPrereqs  Array of token pre-requisites (including natives) for the request to be filled on the destination chain.
     * @param call          Call to be executed by the outbox on the destination chain.
     */
    struct SolverNetIntent {
        uint64 srcChainId;
        uint64 destChainId;
        TokenPrereq[] tokenPrereqs;
        Call call;
    }
}
