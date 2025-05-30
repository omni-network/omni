// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.24;

import { ERC20 } from "solady/src/tokens/ERC20.sol";

/**
 * @title ApprovalReverterERC20
 * @notice ERC20 that reverts when a spender is set to zero.
 */
contract ApprovalReverterERC20 is ERC20 {
    error ZeroAmount();
    error ZeroAddress();

    string private constant _name = "Approve Reverter Token";
    string private constant _symbol = "REVERT";

    function name() public pure override returns (string memory) {
        return _name;
    }

    function symbol() public pure override returns (string memory) {
        return _symbol;
    }

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
