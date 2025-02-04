// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";

contract MockVault {
    using SafeTransferLib for address;

    address public immutable collateral;

    mapping(address depositor => uint256 balance) public balances;

    constructor(address newCollateral) {
        collateral = newCollateral;
    }

    function deposit(address onBehalfOf, uint256 amount) external payable {
        if (collateral != address(0)) collateral.safeTransferFrom(msg.sender, address(this), amount);
        else require(msg.value == amount, "Invalid collateral");
        balances[onBehalfOf] += amount;
    }

    function withdraw(address to, uint256 amount) external {
        balances[msg.sender] -= amount;
        if (collateral != address(0)) collateral.safeTransfer(to, amount);
        else to.safeTransferETH(amount);
    }
}
