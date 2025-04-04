// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.24;

interface IPermit2 {
    function DOMAIN_SEPARATOR() external view returns (bytes32);
    function allowance(address user, address token, address spender)
        external
        view
        returns (uint160 amount, uint48 expiration, uint48 nonce);
}
