// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

/**
 * @title Slashing
 * @notice The EVM interface to the consensus chain's x/slashing module.
 *         Calls are proxied, and not executed syncronously. Their execution is left to
 *         the consensus chain, and they may fail.
 * @dev This contract is predeployed as an upgradable proxy, though currently has no storage.
 *      It therefoes does not need to be Initializeable. If storage is added, it will need to
 *      be Initializeable (in current v0.4.9 of OpenZeppelin). If we upgrade to  v5 of OpenZeppelin,
 *      we could wait to add Initializeable until initialization logic is required, as
 *      Initializeable storage is stored in a custom slot, not the first slots.
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
