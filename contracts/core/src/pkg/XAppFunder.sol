// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import { IOmniExchange } from "../interfaces/IOmniExchange.sol";

contract XFunder {
    event OmniExchangeSet(address exchange);

    /**
     * @notice The OmniGasEx contract
     */
    IOmniExchange public exchange;

    function fund(address recipient, uint256 amtETH) internal returns (uint256) {
        uint64 defaultGasLimit = 100_000;
        uint256 fee = exchange.fundFee(recipient, amtETH, defaultGasLimit);
        require(msg.value >= fee + amtETH, "XApp: insufficient funds");
        exchange.fund{ value: fee + amtETH }(recipient, amtETH, defaultGasLimit);
        return fee + amtETH;
    }

    function fundOrRefund(address recipient, uint256 excess) internal {
        uint64 defaultGasLimit = 100_000;
        // Use max - as we don't know fee yet, and max will give us largest fee
        uint256 fee = exchange.fundFee(recipient, type(uint256).max, defaultGasLimit);

        // if not enough excess to cover fee, refund excess
        if (fee > excess) {
            payable(msg.sender).transfer(excess);
            return;
        }

        require(msg.value >= fee + excess, "XApp: insufficient funds");
        exchange.fund{ value: fee + excess }(recipient, excess - fee, defaultGasLimit);
    }

    function _setExchange(address _exchange) internal {
        require(_exchange != address(0), "XApp: zero addr");
        exchange = IOmniExchange(_exchange);
        emit OmniExchangeSet(_exchange);
    }
}
