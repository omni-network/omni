// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { Secp256k1 } from "../libraries/Secp256k1.sol";

/**
 * @title OmniStake
 * @notice The deposit contract for OMNI-staked validators.
 */
contract OmniStake {
    /**
     * @notice Emitted when a user deposits funds into the contract.
     * @param pubkey    64 byte uncompressed secp256k1 public key (no 0x04 prefix)
     * @param amount    Funds deposited
     */
    event Deposit(bytes pubkey, uint256 amount);

    /**
     * @notice Deposit OMNI. This is the entry point for validator staking. The consensus chain is
     *         notified of the deposit, and manages stake accounting / validator onboarding.
     * @param pubkey 64 byte uncompressed secp256k1 public key (no 0x04 prefix)
     */
    function deposit(bytes memory pubkey) external payable {
        require(msg.value >= 1 ether, "OmniStake: deposit amt too low");
        require(msg.value < type(uint64).max, "OmniStake: deposit amt too high");
        require(msg.sender == Secp256k1.pubkeyToAddress(pubkey), "OmniStake: pubkey not sender");
        emit Deposit(pubkey, msg.value);
    }
}
