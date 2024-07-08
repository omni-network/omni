// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { IStrategy } from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";
import { IOmniAVS } from "src/interfaces/IOmniAVS.sol";
import { EigenM2HoleskyDeployments } from "./eigen/EigenM2HoleskyDeployments.sol";

/**
 * @title StrategyParams
 * @dev Defines OmniAVS strategy params for different chains
 */
library StrategyParams {
    /// @notice standar strategy multiplier, matches OmniAVS.STRATEGY_WEIGHTING_DIVISOR
    uint96 public constant STD_MULTIPLIER = 1e18;

    /// @notice EigenLayer's canonical, virtual beacon chain ETH strategy
    address public constant BEACON_CHAIN_ETH_STRATEGY = 0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0;

    /// @notice Holesky strategy params
    function holesky() external pure returns (IOmniAVS.StrategyParam[] memory params) {
        params = new IOmniAVS.StrategyParam[](4);

        params[0] = IOmniAVS.StrategyParam({
            strategy: IStrategy(EigenM2HoleskyDeployments.stETHStrategy),
            multiplier: STD_MULTIPLIER
        });

        params[1] = IOmniAVS.StrategyParam({
            strategy: IStrategy(EigenM2HoleskyDeployments.rETHStrategy),
            multiplier: STD_MULTIPLIER
        });

        params[2] = IOmniAVS.StrategyParam({
            strategy: IStrategy(EigenM2HoleskyDeployments.WETHStrategy),
            multiplier: STD_MULTIPLIER
        });

        params[3] =
            IOmniAVS.StrategyParam({ strategy: IStrategy(BEACON_CHAIN_ETH_STRATEGY), multiplier: STD_MULTIPLIER });
    }
}
