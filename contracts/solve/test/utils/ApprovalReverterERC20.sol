// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { ERC20 } from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

/**
 * @title ApprovalReverterERC20
 * @notice ERC20 that reverts when a spender is set to zero.
 */
contract ApprovalReverterERC20 is ERC20 {
    error ZeroAmount();
    error ZeroAddress();

    constructor() ERC20("Approve Reverter Token", "REVERT") { }

    function mint(address to, uint256 amount) external {
        _mint(to, amount);
    }

    function burn(address from, uint256 amount) external {
        if (from != msg.sender) _spendAllowance(from, msg.sender, amount);
        _burn(from, amount);
    }

    function approve(address spender, uint256 amount) public override returns (bool) {
        if (spender == address(0)) revert ZeroAddress();
        if (amount == 0) revert ZeroAmount();
        return super.approve(spender, amount);
    }
}
