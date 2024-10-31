// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import { SafeERC20 } from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

contract Vault {
    using SafeERC20 for IERC20;

    address public immutable collateral;

    mapping(address depositor => uint256 balance) public balances;

    constructor(address newCollateral) {
        collateral = newCollateral;
    }

    function deposit(address onBehalfOf, uint256 amount) external {
        IERC20(collateral).safeTransferFrom(msg.sender, address(this), amount);
        balances[onBehalfOf] += amount;
    }

    function withdraw(address to, uint256 amount) external {
        balances[msg.sender] -= amount;
        IERC20(collateral).safeTransfer(to, amount);
    }
}
