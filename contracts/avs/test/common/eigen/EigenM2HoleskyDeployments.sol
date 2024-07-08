// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

// solhint-disable const-name-snakecase

library EigenM2HoleskyDeployments {
    // core
    address internal constant PauserRegistry = 0x85Ef7299F8311B25642679edBF02B62FA2212F06;
    address internal constant DelegationManager = 0xA44151489861Fe9e3055d95adC98FbD462B948e7;
    address internal constant EigenPodManager = 0x30770d7E3e71112d7A6b7259542D1f680a70e315;
    address internal constant StrategyManager = 0xdfB5f6CE42aAA7830E94ECFCcAd411beF4d4D5b6;
    address internal constant Slasher = 0xcAe751b75833ef09627549868A04E32679386e7C;
    address internal constant AVSDirectory = 0x055733000064333CaDDbC92763c58BF0192fFeBf;

    // strategies
    address internal constant stETHStrategy = 0x7D704507b76571a51d9caE8AdDAbBFd0ba0e63d3;
    address internal constant rETHStrategy = 0x3A8fBdf9e77DFc25d09741f51d3E181b25d0c4E0;
    address internal constant WETHStrategy = 0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9;
}
