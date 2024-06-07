// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

/**
 * @title Staking
 * @notice The EVM interface to the consensus chain's x/staking module.
 *         Calls are proxied, and not executed syncronously. Their execution is left to
 *         the consensus chain, and they may fail.
 * @dev This contract is predeployed as an upgradable proxy, though currently has no storage.
 *      It therefoes does not need to be Initializeable. If storage is added, it will need to
 *      be Initializeable (in current v0.4.9 of OpenZeppelin). If we upgrade to  v5 of OpenZeppelin,
 *      we could wait to add Initializeable until initialization logic is required, as
 *      Initializeable storage is stored in a custom slot, not the first slots.
 */
contract Staking {
    /**
     * @notice Emitted when a validator is created
     * @param validator     (MsgCreateValidator.validator_addr) The address of the validator to create
     * @param pubkey        (MsgCreateValidator.pubkey) The validators consensus public key. 33 bytes compressed secp256k1 public key
     * @param deposit       (MsgCreateValidator.selfDelegation) The validators initial stake
     */
    event CreateValidator(address indexed validator, bytes pubkey, uint256 deposit);

    /**
     * @notice Emitted when a delegation is made to a validator
     * @param delegator     (MsgDelegate.delegator_addr) The address of the delegator
     * @param validator     (MsgDelegate.validator_addr) The address of the validator to delegate to
     * @param amount        (MsgDelegate.amount) The amount of tokens to delegate
     */
    event Delegate(address indexed delegator, address indexed validator, uint256 amount);

    /**
     * @notice The minimum deposit required to create a validator
     */
    uint256 public constant MIN_DEPOSIT = 100 ether;

    /**
     * @notice Create a new validator
     * @param pubkey The validators consensus public key. 33 bytes compressed secp256k1 public key
     * @dev Proxies x/staking.MsgCreateValidator
     */
    function createValidator(bytes calldata pubkey) external payable {
        require(pubkey.length == 33, "Staking: invalid pubkey length");
        require(msg.value >= MIN_DEPOSIT, "Staking: insufficient deposit");

        emit CreateValidator(msg.sender, pubkey, msg.value);
    }

    /**
     * @notice Increase your validators self delegation.
     *         NOTE: Only self delegations to existing validators are currently supported.
     *         If msg.sender is not a validator, the delegation will be lost.
     * @dev Proxies x/staking.MsgDelegate
     */
    function delegate(address validator) external payable {
        require(msg.value > 0, "Staking: insufficient deposit");

        // only support self delegation for now
        require(msg.sender == validator, "Staking: only self delegation");

        emit Delegate(msg.sender, validator, msg.value);
    }
}
