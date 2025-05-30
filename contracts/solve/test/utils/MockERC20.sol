// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.24;

import { ERC20 } from "solady/src/tokens/ERC20.sol";

/**
 * @title MockERC20
 * @notice ERC20 with public mints.
 */
contract MockERC20 is ERC20 {
    string private _name;
    string private _symbol;

    constructor(string memory name_, string memory symbol_) {
        _name = string.concat(name_, " (Mock)");
        _symbol = string.concat(symbol_, " (Mock)");
    }

    function name() public view override returns (string memory) {
        return _name;
    }

    function symbol() public view override returns (string memory) {
        return _symbol;
    }

    function mint(address to, uint256 amount) external {
        _mint(to, amount);
    }

    function burn(address from, uint256 amount) external {
        if (from != msg.sender) _spendAllowance(from, msg.sender, amount);
        _burn(from, amount);
    }
}
