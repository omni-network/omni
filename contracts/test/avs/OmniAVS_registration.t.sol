// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { ISignatureUtils } from "eigenlayer-contracts/src/contracts/interfaces/ISignatureUtils.sol";
import { IOmniAVS } from "src/interfaces/IOmniAVS.sol";
import { Base } from "./common/Base.sol";

/**
 * @title OmniAVS_registration_Test
 * @dev Test suite for the AVS registration functionality
 */
contract OmniAVS_registration_Test is Base {
    /// @dev Test that an operator can register
    function test_registerOperator_succeeds() public {
        address operator = _operator(0);

        _addToAllowlist(operator);
        _registerAsOperator(operator);
        _depositIntoSupportedStrategy(operator, minOperatorStake);
        _registerOperatorWithAVS(operator);

        IOmniAVS.Operator[] memory operators = omniAVS.operators();

        assertEq(operators.length, 1);
        assertEq(operators[0].operator, operator);
        assertEq(operators[0].validatorPubKey, _valPubKey(operator));
    }

    /// @dev Test that an operator can register, be ejected, than register again
    function test_registerOperator_registerTwice_succeeds() public {
        // register once
        address operator = _operator(0);
        _addToAllowlist(operator);
        _registerAsOperator(operator);
        _depositIntoSupportedStrategy(operator, minOperatorStake);

        bytes memory valPubKey = _valPubKey(operator);
        bytes32 salt1 = keccak256("test");
        ISignatureUtils.SignatureWithSaltAndExpiry memory valsig = _validatorSig(operator, salt1);
        ISignatureUtils.SignatureWithSaltAndExpiry memory opsig = _operatorSig(operator, salt1);

        vm.prank(operator);
        omniAVS.registerOperator(valPubKey, valsig, opsig);

        // assert operator is registerd
        IOmniAVS.Operator[] memory operators = omniAVS.operators();
        assertEq(operators.length, 1);
        assertEq(operators[0].operator, operator);
        assertEq(operators[0].validatorPubKey, _valPubKey(operator));
        assertEq(omniAVS.validatorPubKey(operator), valPubKey);

        // eject operator
        vm.prank(omniAVSOwner);
        omniAVS.ejectOperator(operator);

        // try registering again with same salt
        vm.expectRevert("OmniAVS: spent salt");
        vm.prank(operator);
        omniAVS.registerOperator(valPubKey, valsig, opsig);

        // register again, requires new salt
        bytes32 salt2 = keccak256("test2");
        valsig = _validatorSig(operator, salt2);
        opsig = _operatorSig(operator, salt2);

        // re-register
        vm.prank(operator);
        omniAVS.registerOperator(valPubKey, valsig, opsig);

        // assert operator is registerd
        operators = omniAVS.operators();
        assertEq(operators.length, 1);
        assertEq(operators[0].operator, operator);
        assertEq(operators[0].validatorPubKey, _valPubKey(operator));
        assertEq(omniAVS.validatorPubKey(operator), valPubKey);
    }

    // @dev Test that another operator cannot register an active validator key
    function test_registerOperator_activeValKey_reverts() public {
        // register operator
        address operator0 = _operator(0);
        _addToAllowlist(operator0);
        _registerAsOperator(operator0);
        _depositIntoSupportedStrategy(operator0, minOperatorStake);
        _registerOperatorWithAVS(operator0);

        // attempt to register a new operator with same pubkey
        address operator1 = _operator(1);
        _addToAllowlist(operator1);
        _registerAsOperator(operator1);
        _depositIntoSupportedStrategy(operator1, minOperatorStake);

        // uses operator0 validator key
        bytes memory valPubKey = _valPubKey(operator0);

        ISignatureUtils.SignatureWithSaltAndExpiry memory valsig = ISignatureUtils.SignatureWithSaltAndExpiry({
            signature: new bytes(0),
            salt: keccak256("test"),
            expiry: block.timestamp + 1 days
        });

        bytes32 validatorRegistrationDigestHash = omniAVS.validatorRegistrationDigestHash({
            operator: operator1, // but use operator1 in registration digest
            valPubKey: valPubKey,
            salt: valsig.salt,
            expiry: valsig.expiry
        });

        valsig.signature = _sign(_valPrivKey(operator0), validatorRegistrationDigestHash);

        // use operator1 for AVSDirectory operator signature
        ISignatureUtils.SignatureWithSaltAndExpiry memory opsig = _operatorSig(operator1);

        // try registering again with same salt
        vm.expectRevert("OmniAVS: pubkey already active");
        vm.prank(operator1);
        omniAVS.registerOperator(valPubKey, valsig, opsig);
    }

    /// @dev Test that a  validator signature salt cannot be double spent
    function test_registerOperator_spentSalt_reverts() public {
        // register opeator
        address operator = _operator(0);
        _addToAllowlist(operator);
        _registerAsOperator(operator);
        _depositIntoSupportedStrategy(operator, minOperatorStake);

        bytes memory valPubKey = _valPubKey(operator);
        bytes32 salt = keccak256("test");
        ISignatureUtils.SignatureWithSaltAndExpiry memory valsig = _validatorSig(operator, salt);
        ISignatureUtils.SignatureWithSaltAndExpiry memory opsig = _operatorSig(operator);

        vm.prank(operator);
        omniAVS.registerOperator(valPubKey, valsig, opsig);
        assertEq(omniAVS.validatorPubKey(operator), valPubKey);

        // deregister operator
        vm.prank(omniAVSOwner);
        omniAVS.ejectOperator(operator);

        // try registering again with same salt
        vm.expectRevert("OmniAVS: spent salt");
        vm.prank(operator);
        omniAVS.registerOperator(valPubKey, valsig, opsig);
    }

    /// @dev Test that using a validator signature after it's expiry reverts
    function test_registerOperator_expiredSig_reverts() public {
        address operator = _operator(0);
        _addToAllowlist(operator);
        _registerAsOperator(operator);
        _depositIntoSupportedStrategy(operator, minOperatorStake);

        bytes memory valPubKey = _valPubKey(operator);
        bytes32 salt = keccak256("test");
        ISignatureUtils.SignatureWithSaltAndExpiry memory valsig = _validatorSig(operator, salt);
        ISignatureUtils.SignatureWithSaltAndExpiry memory opsig = _operatorSig(operator);

        vm.warp(valsig.expiry + 1 days);
        vm.expectRevert("OmniAVS: signature expired");
        vm.prank(operator);
        omniAVS.registerOperator(valPubKey, valsig, opsig);
    }

    function test_registerOperator_invalidDigestHash_reverts() public {
        address operator = _operator(0);
        _addToAllowlist(operator);
        _registerAsOperator(operator);
        _depositIntoSupportedStrategy(operator, minOperatorStake);

        bytes memory valPubKey = _valPubKey(operator);

        ISignatureUtils.SignatureWithSaltAndExpiry memory valsig = ISignatureUtils.SignatureWithSaltAndExpiry({
            signature: new bytes(0),
            salt: keccak256("test"),
            expiry: block.timestamp + 1 days
        });

        bytes32 validatorRegistrationDigestHash = valRegDigestHash_wrongDomainSeperator({
            operator: operator,
            valPubKey: valPubKey,
            salt: valsig.salt,
            expiry: valsig.expiry
        });

        valsig.signature = _sign(_valPrivKey(operator), validatorRegistrationDigestHash);

        ISignatureUtils.SignatureWithSaltAndExpiry memory opsig = _operatorSig(operator);

        vm.expectRevert("OmniAVS: invalid val signature");
        vm.prank(operator);
        omniAVS.registerOperator(valPubKey, valsig, opsig);
    }

    /// @dev That that an operator cannot register with a pubkey that did not sign the registration digest
    function test_registerOperator_invalidSignature_reverts() public {
        address operator1 = _operator(0);
        address operator2 = _operator(1);

        _addToAllowlist(operator1);
        _registerAsOperator(operator1);
        _depositIntoSupportedStrategy(operator1, minOperatorStake);

        ISignatureUtils.SignatureWithSaltAndExpiry memory val1Sig = _validatorSig(operator1);
        ISignatureUtils.SignatureWithSaltAndExpiry memory op1Sig = _operatorSig(operator1);

        vm.expectRevert("OmniAVS: invalid val signature");
        vm.prank(operator1);

        // use operator2 pubkey
        omniAVS.registerOperator(_valPubKey(operator2), val1Sig, op1Sig);
    }

    /// @dev Test that an operator cannot register with an invalid pubkey
    function test_registerOperator_invalidPubkey_reverts() public {
        address operator = _operator(0);

        _addToAllowlist(operator);
        _registerAsOperator(operator);
        _depositIntoSupportedStrategy(operator, minOperatorStake);

        bytes memory valPubKey = bytes.concat(hex"04", _valPubKey(operator));

        ISignatureUtils.SignatureWithSaltAndExpiry memory valsig = _validatorSig(operator);
        ISignatureUtils.SignatureWithSaltAndExpiry memory opsig = _operatorSig(operator);

        vm.expectRevert("Secp256k1: invalid pubkey length");
        vm.prank(operator);
        omniAVS.registerOperator(valPubKey, valsig, opsig);
    }

    /**
     * Utils & constants copied from OmniAVS.
     */

    /// @dev OmniAVS.DOMAIN_TYPEHASH
    bytes32 public constant DOMAIN_TYPEHASH =
        keccak256("EIP712Domain(string name,uint256 chainId,address verifyingContract)");

    /// @dev OmniAVS.VALIDATOR_REGISTRATION_TYPEHASH
    bytes32 public constant VALIDATOR_REGISTRATION_TYPEHASH =
        keccak256("ValidatorRegistration(address operator,bytes validatorPubKey,bytes32 salt,uint256 expiry)");

    /// @dev Same as OmniAVS.validatorRegistrationDigestHash(...), but uses a different domainSeparator()
    function valRegDigestHash_wrongDomainSeperator(
        address operator,
        bytes memory valPubKey,
        bytes32 salt,
        uint256 expiry
    ) public view returns (bytes32) {
        bytes32 structHash = keccak256(abi.encode(VALIDATOR_REGISTRATION_TYPEHASH, operator, valPubKey, salt, expiry));
        return keccak256(abi.encodePacked("\x19\x01", domainSeparator(), structHash));
    }

    // @dev This domain separator will be different than OmniAVS.domainSeparator(), because it includes address
    function domainSeparator() public view returns (bytes32) {
        return keccak256(abi.encode(DOMAIN_TYPEHASH, keccak256(bytes("OmniAVS")), block.chainid, address(this)));
    }
}
