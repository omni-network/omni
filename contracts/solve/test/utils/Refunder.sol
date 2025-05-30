// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.24;

import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";

contract Refunder {
    using SafeTransferLib for address;

    function _refund() private {
        msg.sender.safeTransferETH(msg.value);
    }

    receive() external payable {
        _refund();
    }

    fallback() external payable {
        _refund();
    }
}
