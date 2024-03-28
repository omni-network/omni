// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

/**
 * @title Predeploys
 * @notice Define addresses of contracts predeployed to Omni's EVM
 * @dev Must match predeploys defined in halo/genutil/evm/predeploys
 */
library Predeploys {
    address public constant ProxyAdmin = 0x121E240000000000000000000000000000000001;
    address public constant OmniStake = 0x121E240000000000000000000000000000000002;
    // address public constant EthStakeInbox = 0x121E240000000000000000000000000000000003; TODO: implement
    address public constant GlobalXRegistry = 0x121E240000000000000000000000000000000004;
}
