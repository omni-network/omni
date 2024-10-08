// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { OmniBridgeNativeFixtures } from "./fixtures/OmniBridgeNativeFixtures.sol";

/**
 * @title OmniBridgeNative_Test
 * @notice Template test suite for src/token/OmniBridgeNative.sol
 * @dev Available fixtures (see OmniBridgeNativeFixtures.sol):
 *
 *  OmniBridgeNativeHarness nativebridge; // native bridge harness - exposes public setters
 *  MockPortal portal;                    // mock portal - use to make mock xcalls
 *  address owner;                        // the contract owner
 *  address l1bridge;                     // stub L1 bridge address
 *  uint64 l1ChainId;                     // L1 chain ID
 *
 * Tips:
 *
 *  Use portal.mockXCall({
 *      sourceChainId: l1ChainId,
 *      sender: l1bridge,
 *      to: address(nativebridge),
 *      ...
 *  }) to test xcalls to the native bridge contract.
 */
contract OmniBridgeNative_Test is OmniBridgeNativeFixtures {
    function test_stub() public { }
}
