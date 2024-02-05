// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

import { IOmniPortal } from "src/protocol/interfaces/IOmniPortal.sol";
import { XTypes } from "src/protocol/libraries/XTypes.sol";

/**
 * @title PingPong
 * @notice A contract that pingpongs xmsgs between two chains
 */
contract PingPong {
    IOmniPortal public omni;

    /**
     * @notice Emitted when the pingpong loop is done
     * @param destChainID The destination chain id
     * @param to The address the PingPong contract on the destination chain
     * @param times The number of pingpong loops completed
     */
    event Done(uint64 destChainID, address to, uint64 times);

    constructor(IOmniPortal _omni) {
        omni = _omni;
    }

    /**
     * @notice Start the pingpong xmsg loop
     * @param destChainID The destination chain id
     * @param to The address the PingPong contract on the destination chain
     * @param times The number of times to pingpong (times == 1 means once there and back)
     */
    function start(uint64 destChainID, address to, uint64 times) external {
        require(times > 0, "PingPong: times must be > 0");
        _xpingpong(destChainID, to, times, times * 2 - 1);
    }

    /**
     * @notice The pingpong xmsg loop
     * @param times The pingpongs in the loop
     * @param n The number of xcalls left to make
     */
    function pingpong(uint64 times, uint64 n) external {
        require(omni.isXCall() && msg.sender == address(omni), "Pong: not an omni xcall");

        XTypes.Msg memory xmsg = omni.xmsg();

        if (n == 0) {
            emit Done(xmsg.sourceChainId, xmsg.sender, times);
            return;
        }

        _xpingpong(xmsg.sourceChainId, xmsg.sender, times, n - 1);
    }

    function _xpingpong(uint64 destChainID, address to, uint64 times, uint64 n) internal {
        omni.xcall(destChainID, to, abi.encodeWithSelector(this.pingpong.selector, times, n));
    }
}
