// SPDX-License-Identifier: MIT OR Apache-2.0
pragma solidity >=0.8.0;

// ============ Internal Imports ============
import { InterchainGasPaymaster } from "./InterchainGasPaymaster.sol";

contract TestInterchainGasPaymaster is InterchainGasPaymaster {
    uint256 public constant gasPrice = 10;

    constructor() {
        initialize(msg.sender, msg.sender);
    }

    function quoteGasPayment(uint32, uint256 gasAmount) public pure override returns (uint256) {
        return gasPrice * gasAmount;
    }

    function getDefaultGasUsage() public pure returns (uint256) {
        return DEFAULT_GAS_USAGE;
    }
}
