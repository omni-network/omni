// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

// solhint-disable const-name-snakecase

library EigenM2Deployments {
    // See https://github.com/Layr-Labs/eigenlayer-contracts?tab=readme-ov-file#deployments

    // core
    address internal constant PauserRegistry = 0x0c431C66F4dE941d089625E5B423D00707977060;
    address internal constant DelegationManager = 0x39053D51B77DC0d36036Fc1fCc8Cb819df8Ef37A;
    address internal constant EigenPodManager = 0x91E677b07F7AF907ec9a428aafA9fc14a0d3A338;
    address internal constant StrategyManager = 0x858646372CC42E1A627fcE94aa7A7033e7CF075A;
    address internal constant Slasher = 0xD92145c07f8Ed1D392c1B88017934E301CC1c3Cd;
    address internal constant AVSDirectory = 0x135DDa560e946695d6f155dACaFC6f1F25C1F5AF;

    // strategies
    address internal constant cbETHStrategy = 0x54945180dB7943c0ed0FEE7EdaB2Bd24620256bc;
    address internal constant stETHStrategy = 0x93c4b944D05dfe6df7645A86cd2206016c51564D;
    address internal constant rETHStrategy = 0x1BeE69b7dFFfA4E2d53C2a2Df135C388AD25dCD2;
    address internal constant ETHxStrategy = 0x9d7eD45EE2E8FC5482fa2428f15C971e6369011d;
    address internal constant ankrETHStrategy = 0x13760F50a9d7377e4F20CB8CF9e4c26586c658ff;
    address internal constant OETHStrategy = 0xa4C637e0F704745D182e4D38cAb7E7485321d059;
    address internal constant osETHStrategy = 0x57ba429517c3473B6d34CA9aCd56c0e735b94c02;
    address internal constant swETHStrategy = 0x0Fe4F44beE93503346A3Ac9EE5A26b130a5796d6;
    address internal constant wBETHStrategy = 0x7CA911E83dabf90C90dD3De5411a10F1A6112184;
    address internal constant sfrxETHStrategy = 0x8CA7A5d6f3acd3A7A8bC468a8CD0FB14B6BD28b6;
    address internal constant LsETHStrategy = 0xAe60d8180437b5C34bB956822ac2710972584473;
    address internal constant mETHSTrategy = 0x298aFB19A105D59E74658C4C334Ff360BadE6dd2;
    address internal constant beaconETHStrategy = 0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0;
}
