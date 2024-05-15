// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { ECDSA } from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import { Secp256k1 } from "../libraries/Secp256k1.sol";

/**
 * @title OmniStake
 * @notice The deposit contract for OMNI-staked validators.
 */
contract OmniStake {
    /**
     * @notice Emitted when a registers a validator key with some stake.
     * @param delegator         Address of the delegator, the account funding the deposit on behalf of
     *                          their validator key.
     * @param validatorPubKey   64 byte uncompressed secp256k1 validator public key (no 0x04 prefix)
     * @param deposit           Funds deposited
     */
    event ValidatorRegistered(address delegator, bytes validatorPubKey, uint256 deposit);

    /**
     * @notice Emitted when a delegator increases their stake.
     * @param delegator         The delegator's address
     * @param amount            The amount of deposited
     */
    event Deposit(address delegator, uint256 amount);

    /**
     * @notice The minimum deposit required to register a validator.
     */
    uint256 public constant MIN_REGISTER_DEPOSIT = 100 ether;

    /**
     * @notice The minimum deposit required to delegate to a validator..
     */
    uint256 public constant MIN_DEPOSIT = 1 ether;

    /**
     * @notice The EIP-712 typehash for the contract's domain
     */
    bytes32 public constant DOMAIN_TYPEHASH =
        keccak256("EIP712Domain(string name,uint256 chainId,address verifyingContract)");

    /**
     * @notice The EIP-712 typehash for the validator registration struct
     */
    bytes32 public constant VALIDATOR_REGISTRATION_TYPEHASH =
        keccak256("ValidatorRegistration(address delegator,bytes validatorPubKey,bytes32 salt,uint256 expiry)");

    struct SignatureWithSaltAndExpiry {
        bytes signature;
        bytes32 salt;
        uint256 expiry;
    }

    /**
     * @notice Tracks spent salts per each validator key.
     */
    mapping(bytes => mapping(bytes32 => bool)) internal _isValSaltSpent;

    /**
     * @notice Deposit OMNI and register a validator key.
     * @param validatorPubKey       64 byte uncompressed secp256k1 validator public key (no 0x04 prefix)
     * @param validatorSignature    Signature validatorRegistrationDigestHash() with salt and expiry
     */
    function register(bytes calldata validatorPubKey, SignatureWithSaltAndExpiry calldata validatorSignature)
        external
        payable
    {
        require(msg.value >= MIN_REGISTER_DEPOSIT, "OmniStake: deposit below min");
        require(validatorSignature.expiry >= block.timestamp, "OmniStake: expired signature");
        require(!_isValSaltSpent[validatorPubKey][validatorSignature.salt], "OmniStake: spent salt");
        require(
            ECDSA.recover(
                validatorRegistrationDigestHash(
                    msg.sender, validatorPubKey, validatorSignature.salt, validatorSignature.expiry
                ),
                validatorSignature.signature
            ) == Secp256k1.pubkeyToAddress(validatorPubKey),
            "OmniStake: invalid val signature"
        );

        _isValSaltSpent[validatorPubKey][validatorSignature.salt] = true;

        emit ValidatorRegistered(msg.sender, validatorPubKey, msg.value);
    }

    /**
     * @notice Delegate additional OMNI stake to an existing validator key.
     *         If the delegator (msg.sender) has not registered a validator key, the deposit is lost.
     */
    function deposit() external payable {
        require(msg.value >= MIN_DEPOSIT, "OmniStake: deposit below min");
        emit Deposit(msg.sender, msg.value);
    }

    /**
     * @notice Returns the digest hash to be signed by an opeator's validator key on registration.
     * @param delegator     The delegator's ethereum address
     * @param valPubKey     The delegator's 64 byte uncompressed secp256k1 validator public key
     * @param salt          A salt unique to this registration
     * @param expiry        The timestamp at which this registration expires
     */
    function validatorRegistrationDigestHash(address delegator, bytes calldata valPubKey, bytes32 salt, uint256 expiry)
        public
        view
        returns (bytes32)
    {
        bytes32 structHash = keccak256(abi.encode(VALIDATOR_REGISTRATION_TYPEHASH, delegator, valPubKey, salt, expiry));
        return keccak256(abi.encodePacked("\x19\x01", domainSeparator(), structHash));
    }

    /**
     * @notice Domain separator for EIP-712 signatures.
     */
    function domainSeparator() public view returns (bytes32) {
        return keccak256(abi.encode(DOMAIN_TYPEHASH, keccak256(bytes("OmniStake")), block.chainid, address(this)));
    }
}
