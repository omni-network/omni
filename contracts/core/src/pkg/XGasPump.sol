// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import { OmniGasPump } from "../token/OmniGasPump.sol";

/**
 * @title XGasPump
 * @notice Abstract contract that makes it easy to swap ETH for OMNI on Omni.
 */
abstract contract XGasPump {
    event Refunded(address indexed recipient, uint256 amtETH, string reason);
    event FundedOMNI(address indexed recipient, uint256 ethPaid, uint256 omniReceived);

    OmniGasPump public immutable omniGasPump;

    constructor(address exchange) {
        omniGasPump = OmniGasPump(exchange);
    }

    /**
     * @notice Swap `amtETH` ETH for OMNI on Omni, funding `recipient`.
     *         Reverts if `amtETH` does not cover xcall fee, or is > max allowed swap.
     */
    function fillUp(address recipient, uint256 amtETH) internal {
        _fillUp(recipient, amtETH);
    }

    /**
     * @notice Fund `recipient` with `amtETH` worth of OMNI on Omni.
     *         If `amtETH` is not swappable, refund to `recipient`.
     */
    function fillUpOrRefund(address recipient, uint256 amtETH) internal {
        _fillUpOrRefund(recipient, recipient, amtETH);
    }

    /**
     * @notice Fund `recipient` with `amtETH` worth of OMNI on Omni.
     *         If `amtETH` is not swappable, refund to `refundTo`.
     */
    function fillUpOrRefund(address refundTo, address recipient, uint256 amtETH) internal {
        _fillUpOrRefund(refundTo, recipient, amtETH);
    }

    function _fillUpOrRefund(address refundTo, address recipient, uint256 amtETH) internal {
        (, bool succes, string memory reason) = omniGasPump.dryFillUp(amtETH);

        if (!succes) {
            emit Refunded(refundTo, amtETH, reason);
            payable(refundTo).transfer(amtETH);
            return;
        }

        _fillUp(recipient, amtETH);
    }

    function _fillUp(address recipient, uint256 amtETH) private {
        uint256 amtOMNI = omniGasPump.fillUp{ value: amtETH }(recipient);
        emit FundedOMNI(recipient, amtETH, amtOMNI);
    }
}
