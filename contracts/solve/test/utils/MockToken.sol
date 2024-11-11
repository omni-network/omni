// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { ERC20 } from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

/**
 * @title MockToken
 * @notice ERC20 with public mints.
 */
contract MockToken is ERC20 {
    constructor() ERC20("MockToken", "MTK") { }

    function mint(address to, uint256 amount) external {
        _mint(to, amount);
    }
}
