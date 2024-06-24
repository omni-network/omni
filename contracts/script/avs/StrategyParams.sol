// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { IStrategy } from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";
import { IOmniAVS } from "src/interfaces/IOmniAVS.sol";
import { EigenM2Deployments } from "./EigenM2Deployments.sol";

/**
 * @title StrategyParams
 * @dev Defines OmniAVS strategy params for different chains
 */
library StrategyParams {
    /// @notice standar strategy multiplier, matches OmniAVS.STRATEGY_WEIGHTING_DIVISOR
    uint96 public constant STD_MULTIPLIER = 1e18;

    /// @notice Mainnet strategy params
    function mainnet() internal pure returns (IOmniAVS.StrategyParam[] memory params) {
        params = new IOmniAVS.StrategyParam[](5);
        params[0] = _stdStrategyParam(EigenM2Deployments.cbETHStrategy);
        params[1] = _stdStrategyParam(EigenM2Deployments.stETHStrategy);
        params[2] = _stdStrategyParam(EigenM2Deployments.rETHStrategy);
        params[3] = _stdStrategyParam(EigenM2Deployments.wBETHStrategy);
        params[4] = _stdStrategyParam(EigenM2Deployments.beaconETHStrategy);
    }

    function _stdStrategyParam(address strategy) internal pure returns (IOmniAVS.StrategyParam memory) {
        return IOmniAVS.StrategyParam({ strategy: IStrategy(strategy), multiplier: STD_MULTIPLIER });
    }
}
