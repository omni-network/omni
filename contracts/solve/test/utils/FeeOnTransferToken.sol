// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { ERC20 } from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import { FixedPointMathLib } from "solady/src/utils/FixedPointMathLib.sol";

/**
 * @title FeeOnTransferToken
 * @notice ERC20 that charges a fee on transfer.
 */
contract FeeOnTransferToken is ERC20 {
    uint256 public constant FEE_RATE_BPS = 100; // 1% fee (100 basis points)
    address public feeCollector = address(0xdead);

    constructor() ERC20("FeeOnTransferToken", "FOT") { }

    function mint(address to, uint256 amount) external {
        _mint(to, amount);
    }

    function burn(address from, uint256 amount) external {
        if (from != msg.sender) _spendAllowance(from, msg.sender, amount);
        _burn(from, amount);
    }

    function _update(address from, address to, uint256 value) internal override {
        if (from == address(0) || to == address(0) || value == 0) {
            super._update(from, to, value);
        } else {
            uint256 fee = FixedPointMathLib.mulDivUp(value, FEE_RATE_BPS, 10_000);
            uint256 amountAfterFee = value - fee;

            if (fee > 0) super._update(from, feeCollector, fee);
            super._update(from, to, amountAfterFee);
        }
    }
}
