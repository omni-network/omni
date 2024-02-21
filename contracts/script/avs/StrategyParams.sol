// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { IStrategy } from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";
import { IOmniAVS } from "src/interfaces/IOmniAVS.sol";
import { EigenM2GoerliDeployments } from "test/avs/eigen/EigenM2GoerliDeployments.sol";

library StrategyParams {
    uint96 public constant STD_MULTIPLIER = 1e18; // OmniAVS.WEIGHTING_DIVISOR

    /// @notice Goerli strategy params
    function goerli() external pure returns (IOmniAVS.StrategyParams[] memory params) {
        params = new IOmniAVS.StrategyParams[](2);

        params[0] = IOmniAVS.StrategyParams({
            strategy: IStrategy(EigenM2GoerliDeployments.stETHStrategy),
            multiplier: STD_MULTIPLIER
        });

        params[0] = IOmniAVS.StrategyParams({
            strategy: IStrategy(EigenM2GoerliDeployments.stETHStrategy),
            multiplier: STD_MULTIPLIER
        });
    }

    /// @notice Mainnet strategy params
    function mainnet() external pure returns (IOmniAVS.StrategyParams[] memory) {
        revert("Not implemented");
    }
}
