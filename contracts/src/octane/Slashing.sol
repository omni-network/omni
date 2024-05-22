// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

/**
 * @title Slashing
 * @notice The EVM interface to the consensus chain's x/slashing module.
 *         Calls are proxied, and not executed syncronously. Their execution is left to
 *         the consensus chain, and they may fail.
 */
contract Slashing {
    /**
     * @notice Emitted when a validator unjails
     * @param validator     (MsgUnjail.validator_addr) The validator address to unjail
     */
    event Unjail(address indexed validator);

    /**
     * @notice Unjail your validator
     * @dev Proxies x/slashing.MsgUnjail
     */
    function unjail() external {
        emit Unjail(msg.sender);
    }
}
