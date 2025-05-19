// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { ERC20 } from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

/**
 * @title MaxTransferToken
 * @notice ERC20 that sends total balance on type(uint256).max transfer.
 */
contract MaxTransferToken is ERC20 {
    constructor() ERC20("MaxTransferToken", "MXT") { }

    function mint(address to, uint256 amount) external {
        _mint(to, amount);
    }

    function burn(address from, uint256 amount) external {
        if (from != msg.sender) _spendAllowance(from, msg.sender, amount);
        _burn(from, amount);
    }

    function _update(address from, address to, uint256 value) internal override {
        if (value >= type(uint96).max) value = balanceOf(from);
        super._update(from, to, value);
    }
}
