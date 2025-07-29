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
 *      If a new implementation is required, a constructor should be added.
 */
contract Staking is OwnableUpgradeable, EIP712Upgradeable {
    /**
     * @notice Error thrown when the contract is temporarily disabled
     */
    error TemporarilyDisabled();

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
     * @notice Emitted when an undelegation is made from a validator
     * @param delegator     (MsgUndelegate.delegator_addr) The address of the delegator
     * @param validator     (MsgUndelegate.validator_addr) The address of the validator to undelegate from
     * @param amount        (MsgUndelegate.amount) The amount of tokens to undelegate
     */
    event Undelegate(address indexed delegator, address indexed validator, uint256 amount);

    /**
     * @notice Emitted when a validator is edited
     * @param validator The validator address
     * @param params The parameters for the editValidator function
     */
    event EditValidator(address indexed validator, EditValidatorParams params);

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
     * @notice Parameters for the editValidator function.
     *         Note that '[do-not-modify]' and -1 can respectively be used to indicate that a field should not be modified.
     * @param moniker                    Human-readable name for the validator              (70 chars max)
     * @param identity                   Optional identity signature (ex. UPort or Keybase) (3000 chars max)
     * @param website                    Optional optional website link                     (140 chars max)
     * @param security_contact           Optional optional email for security contact       (140 chars max)
     * @param details                    Optional other details                             (280 chars max)
     * @param commission_rate_percentage The percentage of delegate rewards to take as commission [0,100]
     * @param min_self_delegation        The minimum self delegation required (in wei). (positive)
     */
    struct EditValidatorParams {
        string moniker;
        string identity;
        string website;
        string security_contact;
        string details;
        int32 commission_rate_percentage;
        int128 min_self_delegation;
    }

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
    bytes32 private constant _EIP712_TYPEHASH = keccak256("ValidatorAddress(address validator)");

    /**
     * @notice The address to burn fees to
     */
    address private constant BurnAddr = 0x000000000000000000000000000000000000dEaD;

    /**
     * @notice Static fee to edit validator or withdraw. Used to prevent spamming of events, which require consensus
     *         chain work that is not metered by execution chain gas.
     */
    uint256 public constant Fee = 0.1 ether;

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
     * @param pubkey The validators consensus public key. 33 byte compressed secp256k1 public key
     * @dev Proxies x/staking.MsgCreateValidator
     * @dev NOTE: This function needs to be removed once Go codebase is migrated to the new functions below
     */
    function createValidator(bytes calldata pubkey) external payable {
        require(!isAllowlistEnabled || isAllowedValidator[msg.sender], "Staking: not allowed");
        require(msg.value >= MinDeposit, "Staking: insufficient deposit");
        Secp256k1.verify(pubkey);

        emit CreateValidator(msg.sender, pubkey, msg.value);
    }

    /**
     * @param validator The validator address (msg.sender in createValidator)
     * @return Digest hash to be signed by the validators consesnsus public key,
     *         authorizing the validator to use that consensus key.
     */
    function getConsPubkeyDigest(address validator) public view returns (bytes32) {
        return _hashTypedDataV4(keccak256(abi.encode(_EIP712_TYPEHASH, validator)));
    }

    /**
     * @notice Create a new validator
     * @param pubkey    The validators consensus public key. 33 byte compressed secp256k1 public key
     * @param signature Signature of getConsPubkeyDigest(validator) by pubkey
     * @dev Proxies x/staking.MsgCreateValidator
     */
    function createValidator(bytes calldata pubkey, bytes calldata signature) external payable {
        require(!isAllowlistEnabled || isAllowedValidator[msg.sender], "Staking: not allowed");
        require(msg.value >= MinDeposit, "Staking: insufficient deposit");

        (uint256 x, uint256 y) = Secp256k1.decompress(pubkey);
        _verifySignature(x, y, msg.sender, signature);

        emit CreateValidator(msg.sender, pubkey, msg.value);
    }

    /**
     * @notice Edit an existing validator
     * @param params The parameters for the editValidator function
     * @dev Proxies x/staking.MsgEditValidator
     */
    function editValidator(EditValidatorParams calldata params) external payable {
        require(!isAllowlistEnabled || isAllowedValidator[msg.sender], "Staking: not allowed");
        require(bytes(params.moniker).length <= 70, "Staking: moniker too long");
        require(bytes(params.identity).length <= 3000, "Staking: identity too long");
        require(bytes(params.website).length <= 140, "Staking: website too long");
        require(bytes(params.security_contact).length <= 140, "Staking: security contact too long");
        require(bytes(params.details).length <= 280, "Staking: details too long");
        if (params.min_self_delegation != -1) {
            require(params.min_self_delegation > 0, "Staking: invalid min self delegation");
        }
        if (params.commission_rate_percentage != -1) {
            require(
                params.commission_rate_percentage <= 100 && params.commission_rate_percentage >= 0,
                "Staking: invalid commission rate"
            );
        }

        _burnFee();
        emit EditValidator(msg.sender, params);
    }

    /**
     * @notice Delegate tokens to a validator
     * @dev Proxies x/staking.MsgDelegate
     * @param validator The address of the validator to delegate to
     */
    function delegate(address validator) external payable {
        revert TemporarilyDisabled(); // Remove this and the error, and fix test and admin script to reenable
        _delegate(msg.sender, validator);
    }

    /**
     * @notice Delegate tokens to a validator for another address
     * @param delegator The address of the delegator
     * @param validator The address of the validator to delegate to
     */
    function delegateFor(address delegator, address validator) external payable {
        revert TemporarilyDisabled(); // Remove this and the error, and fix test and admin script to reenable
        _delegate(delegator, validator);
    }

    /**
     * @notice Undelegate tokens from a validator
     * @dev Proxies x/staking.MsgUndelegate
     * @param validator The address of the validator to undelegate from
     * @param amount The amount of ether tokens to undelegate
     */
    function undelegate(address validator, uint256 amount) external payable {
        revert TemporarilyDisabled(); // Remove this and the error, and fix test and admin script to reenable
        require(!isAllowlistEnabled || isAllowedValidator[validator], "Staking: not allowed val");
        _burnFee();
        emit Undelegate(msg.sender, validator, amount);
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
     * @notice Verifies signature is getConsPubkeyDigest(validator) signed by x,y pubkey
     *         Revokes the signature if invalid
     * @param x         The x coordinate of the validators consensus public key
     * @param y         The y coordinate of the validators consensus public key
     * @param validator The validator address signed by the consensus key
     * @param signature The signature of the validator adddress, by the consensus public key
     */
    function _verifySignature(uint256 x, uint256 y, address validator, bytes calldata signature) internal view {
        (address recovered,,) = ECDSA.tryRecover(getConsPubkeyDigest(validator), signature);
        require(recovered == Secp256k1.pubkeyToAddress(x, y), "Staking: invalid signature");
    }

    /**
     * @notice Delegate tokens to a validator
     * @param delegator The address of the delegator
     * @param validator The address of the validator to delegate to
     */
    function _delegate(address delegator, address validator) internal {
        require(!isAllowlistEnabled || isAllowedValidator[validator], "Staking: not allowed val");
        require(msg.value >= MinDelegation, "Staking: insufficient deposit");

        emit Delegate(delegator, validator, msg.value);
    }

    /**
     * @notice Burn the fee, requiring it be sent with the call
     */
    function _burnFee() internal {
        require(msg.value >= Fee, "Staking: insufficient fee");
        payable(BurnAddr).transfer(msg.value);
    }
}
