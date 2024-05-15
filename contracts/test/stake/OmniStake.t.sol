// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OmniStake } from "src/protocol/OmniStake.sol";
import { Test, Vm } from "forge-std/Test.sol";

/**
 * @title OmniStake_Test
 * @notice Test suite for the OmniStake contract
 */
contract OmniStake_Test is Test {
    /**
     * @notice Emitted when a registers a validator key with some stake.
     * @dev Matches OmniStake.ValidatorRegistered event
     */
    event ValidatorRegistered(address delegator, bytes validatorPubKey, uint256 deposit);

    /**
     * @notice Emitted when a delegator increases their stake.
     * @dev Matches OmniStake.Deposit event
     */
    event Deposit(address delegator, uint256 amount);

    OmniStake stake;

    function setUp() public {
        stake = new OmniStake();
    }

    /// @dev Test that a delegate can register a validator
    function test_register_succeds() public {
        Vm.Wallet memory val = vm.createWallet("val");
        address delegator = makeAddr("delegator");

        uint256 deposit = stake.MIN_REGISTER_DEPOSIT() + 1 ether;
        vm.deal(delegator, deposit);

        bytes32 salt = keccak256("salt");
        uint256 expiry = block.timestamp + 1 days;

        bytes32 digestHash = stake.validatorRegistrationDigestHash(delegator, _pubkey(val), salt, expiry);

        OmniStake.SignatureWithSaltAndExpiry memory sig =
            OmniStake.SignatureWithSaltAndExpiry({ signature: _sign(val, digestHash), salt: salt, expiry: expiry });

        vm.expectEmit();
        emit ValidatorRegistered(delegator, _pubkey(val), deposit);

        vm.prank(delegator);
        stake.register{ value: deposit }(_pubkey(val), sig);
    }

    /// @dev Test that the pubkey of a deposit must match the sender
    function test_register_invalidSignature_reverts() public {
        Vm.Wallet memory val = vm.createWallet("val");
        address delegator = makeAddr("delegator");

        uint256 deposit = stake.MIN_REGISTER_DEPOSIT() + 1 ether;
        vm.deal(delegator, deposit);

        bytes32 salt = keccak256("salt");
        uint256 expiry = block.timestamp + 1 days;

        // use invalid digest
        bytes32 digestHash = valRegDigestHash_wrongDomain(delegator, _pubkey(val), salt, expiry);

        OmniStake.SignatureWithSaltAndExpiry memory sig =
            OmniStake.SignatureWithSaltAndExpiry({ signature: _sign(val, digestHash), salt: salt, expiry: expiry });

        vm.expectRevert("OmniStake: invalid val signature");
        vm.prank(delegator);
        stake.register{ value: deposit }(_pubkey(val), sig);
    }

    // @dev Test that a deposit with an invalid pubkey reverts
    function test_register_invalidPubKey_reverts() public {
        Vm.Wallet memory val = vm.createWallet("test val");
        address delegator = makeAddr("delegator");

        uint256 deposit = stake.MIN_REGISTER_DEPOSIT() + 1 ether;
        vm.deal(delegator, deposit);

        bytes32 salt = keccak256("salt");
        uint256 expiry = block.timestamp + 1 days;

        // use invalid digest
        bytes32 digestHash = valRegDigestHash_wrongDomain(delegator, _pubkey(val), salt, expiry);

        OmniStake.SignatureWithSaltAndExpiry memory sig =
            OmniStake.SignatureWithSaltAndExpiry({ signature: _sign(val, digestHash), salt: salt, expiry: expiry });

        bytes memory invalidPubKey = bytes.concat(hex"04", _pubkey(val));

        vm.expectRevert("Secp256k1: invalid pubkey length");
        vm.prank(delegator);
        stake.register{ value: deposit }(invalidPubKey, sig);
    }

    /// @dev Test that using a validator signature with a spent salt reverts
    function test_register_spentSalt_reverts() public {
        // register once
        Vm.Wallet memory val = vm.createWallet("val");
        address delegator = makeAddr("delegator");

        uint256 deposit = stake.MIN_REGISTER_DEPOSIT() + 1 ether;
        vm.deal(delegator, deposit);

        bytes32 salt = keccak256("salt");
        uint256 expiry = block.timestamp + 1 days;

        bytes32 digestHash = stake.validatorRegistrationDigestHash(delegator, _pubkey(val), salt, expiry);

        OmniStake.SignatureWithSaltAndExpiry memory sig =
            OmniStake.SignatureWithSaltAndExpiry({ signature: _sign(val, digestHash), salt: salt, expiry: expiry });

        vm.expectEmit();
        emit ValidatorRegistered(delegator, _pubkey(val), deposit);

        vm.prank(delegator);
        stake.register{ value: deposit }(_pubkey(val), sig);

        // register again, with a new delegator
        delegator = makeAddr("delegator2");
        vm.deal(delegator, deposit);

        vm.expectRevert("OmniStake: spent salt");
        stake.register{ value: deposit }(_pubkey(val), sig);
    }

    /// @dev Test that registering with an expired signature reverts
    function test_register_signatureExpired_reverts() public {
        Vm.Wallet memory val = vm.createWallet("val");
        address delegator = makeAddr("delegator");

        uint256 deposit = stake.MIN_REGISTER_DEPOSIT() + 1 ether;
        vm.deal(delegator, deposit);

        bytes32 salt = keccak256("salt");
        uint256 expiry = block.timestamp + 1 days;

        // warp ahead of expiry
        vm.warp(block.timestamp + expiry + 1 days);

        bytes32 digestHash = stake.validatorRegistrationDigestHash(delegator, _pubkey(val), salt, expiry);

        OmniStake.SignatureWithSaltAndExpiry memory sig =
            OmniStake.SignatureWithSaltAndExpiry({ signature: _sign(val, digestHash), salt: salt, expiry: expiry });

        vm.expectRevert("OmniStake: expired signature");
        vm.prank(delegator);
        stake.register{ value: deposit }(_pubkey(val), sig);
    }

    /// @dev Test that a valid deposit succeeds
    function test_deposit_succeeds() public {
        address delegator = makeAddr("delegator");
        uint256 amt = 16 ether;
        vm.deal(delegator, amt);

        vm.expectEmit();
        emit Deposit(delegator, amt);

        // NOTE: the consensus chain verifies if the delegator has an active validator
        // if they do not, the deposit is lost. But stake.deposit() will succeed regardless

        vm.prank(delegator);
        stake.deposit{ value: amt }();
    }

    /// @dev Test that a deposit below OmniStake.MIN_DEPOSIT reverts
    function test_deposit_belowMinimum_reverts() public {
        address delegator = makeAddr("delegator");
        uint256 amt = stake.MIN_DEPOSIT() - 1;
        vm.deal(delegator, amt);

        vm.expectRevert("OmniStake: deposit below min");
        vm.prank(delegator);
        stake.deposit{ value: amt }();
    }

    /// @dev Helper function to convert a wallet to a secp256k1 64 uncompressed public key (without 0x04 prefix)
    function _pubkey(Vm.Wallet memory w) internal pure returns (bytes memory) {
        return abi.encode(w.publicKeyX, w.publicKeyY);
    }

    /// @dev Helper to sign digest hash with a wallet
    function _sign(Vm.Wallet memory w, bytes32 hash) internal pure returns (bytes memory) {
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(w.privateKey, hash);
        return abi.encodePacked(r, s, v);
    }

    /**
     * Utils repuposed from OmniStake.sol
     */

    /// @dev Matches OmniStake.DOMAIN_TYPEHASH
    bytes32 public constant DOMAIN_TYPEHASH =
        keccak256("EIP712Domain(string name,uint256 chainId,address verifyingContract)");

    /// @dev Matches OmniStake.VALIDATOR_REGISTRATION_TYPEHASH
    bytes32 public constant VALIDATOR_REGISTRATION_TYPEHASH =
        keccak256("ValidatorRegistration(address delegator,bytes validatorPubKey,bytes32 salt,uint256 expiry)");

    /// @dev Differs from OmniStake.validatorRegistrationDigestHash(), because domainSeparator() is different
    function valRegDigestHash_wrongDomain(address delegator, bytes memory valPubKey, bytes32 salt, uint256 expiry)
        public
        view
        returns (bytes32)
    {
        bytes32 structHash = keccak256(abi.encode(VALIDATOR_REGISTRATION_TYPEHASH, delegator, valPubKey, salt, expiry));
        return keccak256(abi.encodePacked("\x19\x01", domainSeparator(), structHash));
    }

    // @dev Differs from OmniStake.domainSeparator() because uses address(this)
    function domainSeparator() public view returns (bytes32) {
        return keccak256(abi.encode(DOMAIN_TYPEHASH, keccak256(bytes("OmniStake")), block.chainid, address(this)));
    }
}
