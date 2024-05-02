// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { MockPortal } from "test/utils/MockPortal.sol";
import { PingPong } from "src/examples/PingPong.sol";

import { Test } from "forge-std/Test.sol";

/**
 * @title PingPong_Test
 * @dev Test of PingPong
 */
contract PingPong_Test is Test {
    MockPortal public portal;
    PingPong public pp;

    function setUp() public {
        portal = new MockPortal();
        pp = new PingPong(address(portal));

        // pp contract makes xcalls, and needs funds
        vm.deal(address(pp), 1 ether);
    }

    /// @dev Test gas useage of pingpong with recv
    function test_pingpong_recv_gas() public {
        portal.mockXCall(
            1, // destChainID
            msg.sender, // sender
            address(pp), // destContract
            abi.encodeWithSelector(pp.pingpong.selector, 10, 10), // calldata
            pp.GAS_LIMIT()
        );
    }

    /// @dev Test gas useage of pingpong without recv
    function test_pingpong_norecv_gas() public {
        portal.mockXCall(
            1, // destChainID
            msg.sender, // sender
            address(pp), // destContract
            abi.encodeWithSelector(pp.pingpong_norecv.selector, 10, 10), // calldata
            pp.GAS_LIMIT()
        );
    }

    /// @dev Test gas useage of pingpong with recv twice
    function test_pingpong_recv_twice_gas() public {
        portal.mockXCall(
            1, // destChainID
            msg.sender, // sender
            address(pp), // destContract
            abi.encodeWithSelector(pp.pingpong.selector, 10, 10), // calldata
            pp.GAS_LIMIT()
        );
        portal.mockXCall(
            1, // destChainID
            msg.sender, // sender
            address(pp), // destContract
            abi.encodeWithSelector(pp.pingpong.selector, 10, 10), // calldata
            pp.GAS_LIMIT()
        );
    }

    /// @dev Test gas useage of pingpong without recv twice
    function test_pingpong_norecv_twice_gas() public {
        portal.mockXCall(
            1, // destChainID
            msg.sender, // sender
            address(pp), // destContract
            abi.encodeWithSelector(pp.pingpong_norecv.selector, 10, 10), // calldata
            pp.GAS_LIMIT()
        );
        portal.mockXCall(
            1, // destChainID
            msg.sender, // sender
            address(pp), // destContract
            abi.encodeWithSelector(pp.pingpong_norecv.selector, 10, 10), // calldata
            pp.GAS_LIMIT()
        );
    }
}
