// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { ERC20 } from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

/**
 * @title MockERC20
 * @notice ERC20 with public mints.
 */
contract MockERC20 is ERC20 {
    constructor(string memory name, string memory symbol)
        ERC20(string.concat(name, " (Mock)"), string.concat(symbol, " (Mock)"))
    { }

    function mint(address to, uint256 amount) external {
        _mint(to, amount);
    }

    function burn(address from, uint256 amount) external {
        if (from != msg.sender) _spendAllowance(from, msg.sender, amount);
        _burn(from, amount);
    }
}
