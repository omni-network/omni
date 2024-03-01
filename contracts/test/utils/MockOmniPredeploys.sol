// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

/**
 * @title MockOmniPredeploys
 * @notice Mock redeploy addresses for Omni contracts
 *
 * TODO: Replace with actual predeploys, and move to src/libraries/OmniPredeploys.sol
 */
library MockOmniPredeploys {
    address public constant ETH_STAKE_INBOX = address(12_345);
}
