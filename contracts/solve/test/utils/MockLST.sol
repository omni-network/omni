// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.24;

import { MockERC20 } from "./MockERC20.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";

contract MockLST {
    using SafeTransferLib for address;

    error InsufficientBalance();

    MockERC20 public immutable token;

    constructor() {
        token = new MockERC20("RocketPool Staked ETH", "rETH");
    }

    function deposit() external payable {
        if (msg.value > 0) token.mint(msg.sender, msg.value);
    }

    function withdraw(address to, uint256 amount) external {
        if (amount > 0) {
            token.burn(msg.sender, amount);
            to.safeTransferETH(amount);
        }
    }
}
