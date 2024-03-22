// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

// solhint-disable const-name-snakecase

library EigenM2GoerliDeployments {
    // core
    address internal constant PauserRegistry = 0x2588f9299871a519883D92dcd5092B4A0Cf70010;
    address internal constant DelegationManager = 0x1b7b8F6b258f95Cf9596EabB9aa18B62940Eb0a8;
    address internal constant EigenPodManager = 0xa286b84C96aF280a49Fe1F40B9627C2A2827df41;
    address internal constant StrategyManager = 0x779d1b5315df083e3F9E94cB495983500bA8E907;
    address internal constant Slasher = 0xD11d60b669Ecf7bE10329726043B3ac07B380C22;
    address internal constant AVSDirectory = 0x0AC9694c271eFbA6059e9783769e515E8731f935;

    // strategies
    address internal constant stETHStrategy = 0xB613E78E2068d7489bb66419fB1cfa11275d14da;
    address internal constant rETHStrategy = 0x879944A8cB437a5f8061361f82A6d4EED59070b5;
    address internal constant wBETHStrategy = 0xD89dc4C40d901D4622C203Fb8808e6e7C7fcE164;
    address internal constant LsETHStrategy = 0xa9DC3c93ae59B8d26AF17Ae63c96Be78793537A9;
    address internal constant ankrETHStrategy = 0x98b47798B68b734af53c930495595729E96cdB8E;
    address internal constant ETHxStrategy = 0x5d1E9DC056C906CBfe06205a39B0D965A6Df7C14;
    address internal constant mETHSTrategy = 0x1755d34476BB4DaEd726ee4a81E8132dF00F9b14;
}
