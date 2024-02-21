// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

// solhint-disable const-name-snakecase

library EigenM2GoerliDeployments {
    // core
    address internal constant DelegationManager = 0x1b7b8F6b258f95Cf9596EabB9aa18B62940Eb0a8;
    address internal constant EigenPodManager = 0xa286b84C96aF280a49Fe1F40B9627C2A2827df41;
    address internal constant StrategyManager = 0x779d1b5315df083e3F9E94cB495983500bA8E907;
    address internal constant Slasher = 0xD11d60b669Ecf7bE10329726043B3ac07B380C22;
    address internal constant AVSDirectory = 0x0AC9694c271eFbA6059e9783769e515E8731f935;
    address internal constant BeaconChainETHSTrategy = 0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0;

    // strategies
    address internal constant stETHStrategy = 0xB613E78E2068d7489bb66419fB1cfa11275d14da;
    address internal constant stETH = 0x1643E812aE58766192Cf7D2Cf9567dF2C37e9B7F;
    address internal constant rETHStrategy = 0x879944A8cB437a5f8061361f82A6d4EED59070b5;
    address internal constant rETH = 0x178E141a0E3b34152f73Ff610437A7bf9B83267A;
}
