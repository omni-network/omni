// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { Test } from "forge-std/Test.sol";
import { OmniPortal } from "src/protocol/OmniPortal.sol";

contract Harness is OmniPortal {
    function lastXCallID() external view returns (bytes32) {
        return _lastXCallID;
    }

    function setLastCallID(bytes32 id) public {
        _lastXCallID = id;
    }

    function parseLastXCallID() public view returns (uint32 xcallOffset_, uint64 xblockOffset_, uint128 sourceBlock_) {
        return _parseLastXCallID();
    }

    function packLastCallID(uint32 _xcallOffset, uint64 _xblockOffset, uint128 _sourceBlockNumber)
        public
        pure
        returns (bytes32)
    {
        return _packLastCallID(_xcallOffset, _xblockOffset, _sourceBlockNumber);
    }
}

contract OmniPortal_lastXCallID_Test is Test {
    Harness harness;

    function setUp() public {
        harness = new Harness();
    }

    function test_parse_pack() public {
        harness.setLastCallID(harness.packLastCallID(1, 2, 3));

        (uint32 xcallOffset, uint64 xblockOffset, uint128 sourceBlock) = harness.parseLastXCallID();
        assertTrue(xcallOffset == 1);
        assertTrue(xblockOffset == 2);
        assertTrue(sourceBlock == 3);

        harness.setLastCallID(harness.packLastCallID(4, 5, 6));

        (xcallOffset, xblockOffset, sourceBlock) = harness.parseLastXCallID();
        assertTrue(xcallOffset == 4);
        assertTrue(xblockOffset == 5);
        assertTrue(sourceBlock == 6);
    }
}
