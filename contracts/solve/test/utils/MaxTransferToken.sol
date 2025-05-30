// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.24;

import { ERC20 } from "solady/src/tokens/ERC20.sol";

/**
 * @title MaxTransferToken
 * @notice ERC20 that sends total balance on type(uint256).max transfer.
 */
contract MaxTransferToken is ERC20 {
    string private constant _name = "MaxTransferToken";
    string private constant _symbol = "MXT";

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

    function transfer(address to, uint256 amount) public override returns (bool) {
        if (amount >= type(uint96).max) amount = balanceOf(msg.sender);
        return super.transfer(to, amount);
    }

    function transferFrom(address from, address to, uint256 amount) public override returns (bool) {
        if (amount >= type(uint96).max) amount = balanceOf(from);
        return super.transferFrom(from, to, amount);
    }
}
