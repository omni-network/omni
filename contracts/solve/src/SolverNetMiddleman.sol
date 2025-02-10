// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { Initializable } from "solady/src/utils/Initializable.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";

contract SolverNetMiddleman is Initializable {
    using SafeTransferLib for address;

    error CallFailed();

    constructor() {
        _disableInitializers();
    }

    function executeAndTransfer(address token, address to, address target, bytes calldata data) external payable {
        (bool success,) = target.call{ value: msg.value }(data);
        if (!success) revert CallFailed();

        if (token == address(0)) SafeTransferLib.safeTransferAllETH(to);
        else token.safeTransferAll(to);
    }

    receive() external payable { }
    fallback() external payable { }
}
