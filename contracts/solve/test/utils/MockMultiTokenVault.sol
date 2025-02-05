// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";

/// @dev This contract is purely used for testing purposes and should not be deployed in production.
contract MockMultiTokenVault {
    using SafeTransferLib for address;

    error IncorrectAmount();
    error InsufficientBalance();
    error ArrayLengthMismatch();

    mapping(address depositor => mapping(address token => uint256 balance)) public balances;

    constructor() { }

    function deposit(address onBehalfOf, address[] calldata tokens, uint256[] calldata amounts) external payable {
        if (tokens.length != amounts.length) revert ArrayLengthMismatch();

        for (uint256 i; i < tokens.length; ++i) {
            if (tokens[i] == address(0)) {
                if (msg.value != amounts[i]) revert IncorrectAmount();
                balances[onBehalfOf][tokens[i]] += amounts[i];
            } else {
                tokens[i].safeTransferFrom(msg.sender, address(this), amounts[i]);
                balances[onBehalfOf][tokens[i]] += amounts[i];
            }
        }
    }

    function withdraw(address token, address to, uint256 amount) external {
        if (balances[msg.sender][token] < amount) revert InsufficientBalance();

        balances[msg.sender][token] -= amount;

        if (token == address(0)) {
            to.safeTransferETH(amount);
        } else {
            token.safeTransfer(to, amount);
        }
    }
}
