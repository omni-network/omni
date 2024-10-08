// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { OmniBridgeL1Fixtures } from "./fixtures/OmniBridgeL1Fixtures.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";

/**
 * @title OmniBridgeL1_Test
 * @notice Template test suite for src/token/OmniBridgeL1.sol
 * @dev Available fixtures (see OmniBridgeL1Fixtures.sol):
 *
 *   OmniBridgeL1 l1bridge; // the L1 bridge contract
 *   Omni token;            // the OMNI token
 *   MockPortal portal;     // mock portal - use to make mock xcalls
 *   address owner;         // the contract owner
 *   address bank;          // initial token supply recipient
 *
 * Tips:
 *
 *  Use portal.mockXCall({
 *      sourceChainId: portal.omniChainId(),
 *      sender: Predeploys.OmniBridgeNative,
 *      to: address(l1bridge),
 *      ...
 *  }) to test xcalls to the l1 bridge contract.
 */
contract OmniBridgeL1_Test is OmniBridgeL1Fixtures {
    function test_stub() public { }
}
