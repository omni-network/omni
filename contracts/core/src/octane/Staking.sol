// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { EIP712Upgradeable } from "@openzeppelin/contracts-upgradeable/utils/cryptography/EIP712Upgradeable.sol";
import { ECDSA } from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import { Secp256k1 } from "../libraries/Secp256k1.sol";

/**
 * @title Staking
 * @notice The EVM interface to the consensus chain's x/staking module.
 *         Calls are proxied, and not executed synchronously. Their execution is left to
 *         the consensus chain, and they may fail.
 * @dev This contract is predeployed, and requires storage slots to be set in genesis.
 *      initialize(...) is called pre-deployment, in script/genesis/AllocPredeploys.s.sol
 *      Initializers on the implementation are disabled via manual storage updates, rather than in a constructor.
 *      If an new implementation is required, a constructor should be added.
 */
contract Staking is OwnableUpgradeable, EIP712Upgradeable {
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
     * @notice Emitted when a validator is allowed to create a validator
     * @param validator     The validator address
     */
    event ValidatorAllowed(address indexed validator);

    /**
     * @notice Emitted when a validator is disallowed to create a validator
     * @param validator     The validator address
     */
    event ValidatorDisallowed(address indexed validator);

    /**
     * @notice Emitted when the allowlist is enabled
     */
    event AllowlistEnabled();

    /**
     * @notice Emitted when the allowlist is disabled
     */
    event AllowlistDisabled();

    /**
     * @notice The minimum deposit required to create a validator
     */
    uint256 public constant MinDeposit = 100 ether;

    /**
     * @notice The minimum delegation required to delegate to a validator
     */
    uint256 public constant MinDelegation = 1 ether;

    /**
     * @notice EIP-712 typehash
     */
    bytes32 private constant _EIP712_TYPEHASH = keccak256("ValidatorPublicKey(bytes32 x,bytes32 y)");

    /**
     * @notice True if the validator allowlist is enabled.
     */
    bool public isAllowlistEnabled;

    /**
     * @notice True if the validator address is allowed to create a validator.
     */
    mapping(address => bool) public isAllowedValidator;

    constructor() {
        _disableInitializers();
    }

    /**
     * @notice Initialize the contract, used for fresh deployment
     */
    function initialize(address owner_, bool isAllowlistEnabled_) public initializer {
        __Ownable_init(owner_);
        __EIP712_init("Staking", "1");
        isAllowlistEnabled = isAllowlistEnabled_;
    }

    /**
     * @notice Original initializer when first deployed publicly
     */
    function initializeV1(address owner_, bool isAllowlistEnabled_) public initializer {
        __Ownable_init(owner_);
        isAllowlistEnabled = isAllowlistEnabled_;
    }

    /**
     * @notice Initializer for upgrade deployment, unnecessary for fresh deployment
     */
    function initializeV2() public reinitializer(2) {
        __EIP712_init("Staking", "1");
    }

    /**
     * @notice Create a new validator
     * @param pubkey The validators consensus public key. 33 bytes compressed secp256k1 public key
     * @dev Proxies x/staking.MsgCreateValidator
     * @dev NOTE: This function needs to be removed once Go codebase is migrated to the new functions below
     */
    function createValidator(bytes calldata pubkey) external payable {
        require(!isAllowlistEnabled || isAllowedValidator[msg.sender], "Staking: not allowed");
        require(msg.value >= MinDeposit, "Staking: insufficient deposit");
        require(Secp256k1.verifyPubkey(pubkey), "Staking: invalid pubkey");

        emit CreateValidator(msg.sender, pubkey, msg.value);
    }

    /**
     * @param x The x coordinate of the validators consensus public key
     * @param y The y coordinate of the validators consensus public key
     * @return Digest hash to be signed by the validators public key
     */
    function getValidatorPubkeyDigest(bytes32 x, bytes32 y) external view returns (bytes32) {
        return _hashTypedDataV4(keccak256(abi.encode(_EIP712_TYPEHASH, x, y)));
    }

    /**
     * @notice Create a new validator
     * @param x The x coordinate of the validators consensus public key
     * @param y The y coordinate of the validators consensus public key
     * @param signature The signature of the validators consensus public key
     * @dev Proxies x/staking.MsgCreateValidator
     */
    function createValidator(bytes32 x, bytes32 y, bytes calldata signature) external payable {
        require(!isAllowlistEnabled || isAllowedValidator[msg.sender], "Staking: not allowed");
        require(msg.value >= MinDeposit, "Staking: insufficient deposit");
        require(Secp256k1.verifyPubkey(x, y), "Staking: invalid pubkey");
        require(_verifySignature(x, y, signature), "Staking: invalid signature");

        bytes memory pubkey = Secp256k1.compressPublicKey(x, y);
        emit CreateValidator(msg.sender, pubkey, msg.value);
    }

    /**
     * @notice Increase your validators self delegation.
     * @dev Proxies x/staking.MsgDelegate
     */
    function delegate(address delegator, address validator) external payable {
        require(!isAllowlistEnabled || isAllowedValidator[validator], "Staking: not allowed val");
        require(msg.value >= MinDelegation, "Staking: insufficient deposit");

        emit Delegate(delegator, validator, msg.value);
    }

    //////////////////////////////////////////////////////////////////////////////
    //                                  Admin                                   //
    //////////////////////////////////////////////////////////////////////////////

    /**
     * @notice Enable the validator allowlist
     */
    function enableAllowlist() external onlyOwner {
        isAllowlistEnabled = true;
        emit AllowlistEnabled();
    }

    /**
     * @notice Disable the validator allowlist
     */
    function disableAllowlist() external onlyOwner {
        isAllowlistEnabled = false;
        emit AllowlistDisabled();
    }

    /**
     * @notice Add validators to allow list
     */
    function allowValidators(address[] calldata validators) external onlyOwner {
        for (uint256 i = 0; i < validators.length; i++) {
            isAllowedValidator[validators[i]] = true;
            emit ValidatorAllowed(validators[i]);
        }
    }

    /**
     * @notice Remove validators from allow list
     */
    function disallowValidators(address[] calldata validators) external onlyOwner {
        for (uint256 i = 0; i < validators.length; i++) {
            isAllowedValidator[validators[i]] = false;
            emit ValidatorDisallowed(validators[i]);
        }
    }

    //////////////////////////////////////////////////////////////////////////////
    //                                 Internal                                 //
    //////////////////////////////////////////////////////////////////////////////

    /**
     * @notice Verify a signature matches a secp256k1 public key
     * @param x The x coordinate of the validators consensus public key
     * @param y The y coordinate of the validators consensus public key
     * @param signature The signature of the validators consensus public key
     */
    function _verifySignature(bytes32 x, bytes32 y, bytes calldata signature) internal view returns (bool) {
        bytes32 digest = _hashTypedDataV4(keccak256(abi.encode(_EIP712_TYPEHASH, x, y)));
        (address recovered,,) = ECDSA.tryRecover(digest, signature);
        address pubKeyAddress = Secp256k1.pubkeyToAddress(x, y);
        return recovered == pubKeyAddress;
    }
}
