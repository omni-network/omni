// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import { IFeeOracle } from "../../src/interfaces/IFeeOracle.sol";

/**
 * @title MockFeeOracle
 * @notice A mock fee orcale, used by MockPortal
 */
contract MockFeeOracle is IFeeOracle {
    uint256 public fee;

    constructor(uint256 fee_) {
        fee = fee_;
    }

    function feeFor(uint64, bytes calldata, uint64) external view returns (uint256) {
        return fee;
    }
}
