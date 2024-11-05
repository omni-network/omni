// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

interface IInbox {
    function markFulfilled(bytes32 id, bytes32 callHash, address creditTo) external;
}
