// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { XApp } from "src/pkg/XApp.sol";
import { ConfLevel } from "src/libraries/ConfLevel.sol";

/**
 * @title PingPong
 * @notice A contract that pingpongs xmsgs between two chains
 */
contract PingPong is XApp {
    /**
     * @notice Gas limit used for a single pingpong xcall
     */
    uint64 public constant GAS_LIMIT = 200_000;

    /**
     * @notice Emitted when the pingpong loop is done
     * @param id            Ping pong id
     * @param destChainID   The destination chain id
     * @param to            The address the PingPong contract on the destination chain
     * @param times         The number of pingpong loops completed
     */
    event Done(string id, uint64 destChainID, address to, uint64 times);

    /**
     * @notice Emitted when a ping is received
     * @param id            Ping pong id
     * @param srcChainID    The source chain of the ping
     * @param from          The address of the sender of the ping
     * @param n             The number of xcalls left to make
     */
    event Ping(string id, uint64 srcChainID, address from, uint64 n);

    constructor(address portal) XApp(portal, ConfLevel.Latest) { }

    /**
     * @notice Start the pingpong xmsg loop
     * @param id            Ping pong id
     * @param destChainID   The destination chain id
     * @param conf          Confirmation level on which to xcall
     * @param to            The address the PingPong contract on the destination chain
     * @param times         The number of times to pingpong (times == 1 means once there and back)
     */
    function start(string calldata id, uint64 destChainID, uint8 conf, address to, uint64 times) external {
        require(times > 0, "PingPong: times must be > 0");
        _xpingpong(id, destChainID, conf, to, times, times * 2 - 1);
    }

    /**
     * @notice The pingpong xmsg loop
     * @param id        The ping pong id
     * @param conf      Confirmation level on which to xcall
     * @param times     The pingpongs in the loop
     * @param n The     number of xcalls left to make
     */
    function pingpong(string calldata id, uint8 conf, uint64 times, uint64 n) external xrecv {
        require(isXCall(), "PingPong: not an omni xcall");

        emit Ping(id, xmsg.sourceChainId, xmsg.sender, n);

        if (n == 0) {
            emit Done(id, xmsg.sourceChainId, xmsg.sender, times);
            return;
        }

        _xpingpong(id, xmsg.sourceChainId, conf, xmsg.sender, times, n - 1);
    }

    function _xpingpong(string calldata id, uint64 destChainID, uint8 conf, address to, uint64 times, uint64 n)
        internal
    {
        xcall(destChainID, conf, to, abi.encodeWithSelector(this.pingpong.selector, id, conf, times, n), GAS_LIMIT);
    }

    receive() external payable { }
}
