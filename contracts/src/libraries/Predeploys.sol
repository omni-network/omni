// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

/**
 * @title Predeploys
 * @notice Halo predeploy addresses (match halo/genutil/evm/predeploys.go)
 */
library Predeploys {
    // proxy admin for all predeploys
    address internal constant ProxyAdmin = 0xaAaAaAaaAaAaAaaAaAAAAAAAAaaaAaAaAaaAaaAa;

    // omni predeploys
    address internal constant PortalRegistry = 0x121E240000000000000000000000000000000001;
    address internal constant OmniBridgeNative = 0x121E240000000000000000000000000000000002;
    address internal constant WOmni = 0x121E240000000000000000000000000000000003;
    address internal constant EthStakeInbox = 0x121E240000000000000000000000000000000004;

    // octane predeploys
    address internal constant Staking = 0xCCcCcC0000000000000000000000000000000001;
    address internal constant Slashing = 0xCccCCC0000000000000000000000000000000002;
}
