// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

/**
 * @title Slashing
 * @notice The EVM interface to the consensus chain's x/slashing module.
 *         Calls are proxied, and not executed syncronously. Their execution is left to
 *         the consensus chain, and they may fail.
 * @dev This contract is predeployed as an upgradable proxy, though currently has no storage.
 *      It therefoes does not need to be Initializeable. Initializeable should be added when
 *      initialization logic is required.
 */
contract Slashing {
    /**
     * @notice Emitted when a validator unjails
     * @param validator     (MsgUnjail.validator_addr) The validator address to unjail
     */
    event Unjail(address indexed validator);

    /**
     * @notice The address to burn fees to
     */
    address private constant BurnAddr = 0x000000000000000000000000000000000000dEaD;

    /**
     * @notice Static fee to unjail. Used to prevent spamming of Unjail events, which require consensus
     *         chain work that is not metered by execution chain gas.
     */
    uint256 public constant Fee = 0.1 ether;

    /**
     * @notice Unjail your validator
     * @dev Proxies x/slashing.MsgUnjail
     */
    function unjail() external payable {
        _burnFee();
        emit Unjail(msg.sender);
    }

    /**
     * @notice Burn the fee, requiring it be sent with the call
     */
    function _burnFee() internal {
        require(msg.value >= Fee, "Slashing: insufficient fee");
        payable(BurnAddr).transfer(msg.value);
    }
}
