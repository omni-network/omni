// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import { IStrategy } from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";
import { ISignatureUtils } from "eigenlayer-contracts/src/contracts/interfaces/ISignatureUtils.sol";

/**
 * @title IOmniAVS
 * @notice Interface for the Omni AVS contract. It is responsible for syncing Omni AVS operator
 *         stake and delegations with the Omni chain.
 */
interface IOmniAVS {
    /**
     * @notice Emitted when an operator is added to the OmniAVS.
     * @param operator The address of the operator
     * @param valPubKey The operator's 64 byte uncompressed secp256k1 validator public key
     */
    event OperatorAdded(address indexed operator, bytes valPubKey);

    /**
     * @notice Emitted when an operator is removed from the OmniAVS.
     * @param operator The address of the operator
     */
    event OperatorRemoved(address indexed operator);

    /**
     * @notice Struct representing an OmniAVS operator
     * @custom:field operator           The operator's ethereum address
     * @custom:field validatorPubKey    The operator's 64 byte uncompressed secp256k1 validator public key
     * @custom:field delegated          The total amount delegated, not including operator stake
     * @custom:field staked             The total amount staked by the operator, not including delegations
     */
    struct Operator {
        address operator;
        bytes validatorPubKey;
        uint96 delegated;
        uint96 staked;
    }

    /**
     * @notice Represents a single supported strategy.
     * @custom:field strategy   The strategy contract
     * @custom:field multiplier The stake multiplier, to weight strategy against others
     */
    struct StrategyParam {
        IStrategy strategy;
        uint96 multiplier;
    }

    /**
     * @notice Returns the fee required for syncWithOmni(), for the current operator set.
     */
    function feeForSync() external view returns (uint256);

    /**
     * @notice Sync OmniAVS operator stake & delegations with Omni chain.
     */
    function syncWithOmni() external payable;

    /**
     * @notice Returns the currrent list of operator registered as OmniAVS.
     */
    function operators() external view returns (Operator[] memory);

    /**
     * @notice Returns the current strategy parameters.
     */
    function strategyParams() external view returns (StrategyParam[] memory);

    /**
     * @notice Register an operator with the AVS. Forwards call to EigenLayer' AVSDirectory.
     * @param validatorPubKey       The operator's 64 byte uncompressed secp256k1 validator public key.
     * @param validatorSignature    The operator's validator pubkey registration signature, with salt and expiry.
     *                              Signature must match `validatorPubKey`
     * @param operatorSignature     The operator's AVS registration signature, with salt and expiry.
     *                              Signed must match `msg.sender`
     */
    function registerOperator(
        bytes calldata validatorPubKey,
        ISignatureUtils.SignatureWithSaltAndExpiry calldata validatorSignature,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) external;

    /**
     * @notice Returns the digest hash to be signed by an opeator's validator key on registration.
     * @param operator      The operator's ethereum address
     * @param valPubKey     The operator's 64 byte uncompressed secp256k1 validator public key
     * @param salt          A salt unique to this registration
     * @param expiry        The timestamp at which this registration expires
     */
    function validatorRegistrationDigestHash(address operator, bytes calldata valPubKey, bytes32 salt, uint256 expiry)
        external
        view
        returns (bytes32);

    /**
     * @notice Check if an operator can register to the AVS.
     *         Returns true, with no reason, if the operator can register to the AVS.
     *         Returns false, with a reason, if the operator cannot register to the AVS.
     * @dev This function is intented to be called off-chain.
     * @param operator The operator to check
     * @return canRegister True if the operator can register, false otherwise
     * @return reason      The reason the operator cannot register. Empty if canRegister is true.
     */
    function canRegister(address operator) external view returns (bool, string memory);
}
