// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.24;

import { ERC20 } from "solady/src/tokens/ERC20.sol";
import { FixedPointMathLib } from "solady/src/utils/FixedPointMathLib.sol";

/**
 * @title FeeOnTransferToken
 * @notice ERC20 that charges a fee on transfer.
 */
contract FeeOnTransferToken is ERC20 {
    string private constant _name = "FeeOnTransferToken";
    string private constant _symbol = "FOT";

    uint256 public constant FEE_RATE_BPS = 100; // 1% fee (100 basis points)
    address public feeCollector = address(0xdead);

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
        return _transferWithFee(msg.sender, to, amount);
    }

    function transferFrom(address from, address to, uint256 amount) public override returns (bool) {
        if (from != msg.sender) _spendAllowance(from, msg.sender, amount);
        return _transferWithFee(from, to, amount);
    }

    function _transferWithFee(address from, address to, uint256 amount) internal returns (bool) {
        if (amount == 0) {
            _transfer(from, to, 0);
            return true;
        }

        uint256 fee = FixedPointMathLib.mulDivUp(amount, FEE_RATE_BPS, 10_000);
        uint256 amountAfterFee = amount - fee;

        if (fee > 0) {
            _transfer(from, feeCollector, fee);
        }
        _transfer(from, to, amountAfterFee);

        return true;
    }
}
