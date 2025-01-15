// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

contract Reverter {
    receive() external payable {
        revert();
    }

    fallback() external payable {
        revert();
    }
}
