// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { Quorum } from "src/libraries/Quorum.sol";
import { Test } from "forge-std/Test.sol";
import { XTypes } from "src/libraries/XTypes.sol";
import { QuorumVerifier } from "./utils/QuorumVerifier.sol";

contract Quorum_Test is Test {
    // map valset id to array of validators
    mapping(uint64 => XTypes.Validator[]) validators;

    // map validator address to private key
    mapping(address => uint256) privKeys;

    QuorumVerifier quorum;

    function setUp() public {
        quorum = new QuorumVerifier();
    }

    function test_verify_succeeds() public {
        _initValset({ valSetId: 1, numVals: 3 });

        bytes32 digest = keccak256(abi.encodePacked("hello"));
        XTypes.SigTuple[] memory sigs = _getSignatures({ digest: digest, valSetId: 1 });

        bool verified = quorum.verify({ digest: digest, sigs: sigs, valSetId: 1, qNumerator: 2, qDenominator: 3 });
        assertTrue(verified);
    }

    function test_verify_largeValset_succeeds() public {
        _initValset({ valSetId: 1, numVals: 100 });

        bytes32 digest = keccak256(abi.encodePacked("hello"));
        XTypes.SigTuple[] memory sigs = _getSignatures({ digest: digest, valSetId: 1 });

        bool verified = quorum.verify({ digest: digest, sigs: sigs, valSetId: 1, qNumerator: 2, qDenominator: 3 });
        assertTrue(verified);
    }

    function test_verify_noQuorum_fails() public {
        _initValset({ valSetId: 1, numVals: 3 });

        bytes32 digest = keccak256(abi.encodePacked("hello"));

        // only get first 2 signatures
        XTypes.SigTuple[] memory sigs = _getSignatures({ digest: digest, valSetId: 1, numSigs: 2 });

        // require 3/4 quorum
        bool verified = quorum.verify({ digest: digest, sigs: sigs, valSetId: 1, qNumerator: 3, qDenominator: 4 });

        assertFalse(verified);
    }

    function test_verify_powerBasedQourum_succeeds() public {
        // each val will be given 100 power
        _initValset({ valSetId: 1, numVals: 10 });

        // we update val1 and 2 to have 1000 power each, making them quorum
        quorum.updateValidatorPower({ valSetId: 1, addr: validators[1][0].addr, power: 1000 });
        quorum.updateValidatorPower({ valSetId: 1, addr: validators[1][1].addr, power: 1000 });

        bytes32 digest = keccak256(abi.encodePacked("hello"));

        // only get first 2 signatures
        XTypes.SigTuple[] memory sigs = _getSignatures({ digest: digest, valSetId: 1, numSigs: 2 });

        bool verified = quorum.verify({ digest: digest, sigs: sigs, valSetId: 1, qNumerator: 2, qDenominator: 3 });
        assertTrue(verified);
    }

    function test_verify_invalidSignature_reverts() public {
        _initValset({ valSetId: 1, numVals: 3 });

        bytes32 digest = keccak256(abi.encodePacked("hello"));
        XTypes.SigTuple[] memory sigs = _getSignatures({ digest: digest, valSetId: 1 });

        // modify a signature
        sigs[0].signature = _sign(keccak256(abi.encodePacked("world")), privKeys[validators[1][0].addr]);

        vm.expectRevert("Quorum: invalid signature");
        quorum.verify({ digest: digest, sigs: sigs, valSetId: 1, qNumerator: 2, qDenominator: 3 });
    }

    function test_verify_duplicateValidator_reverts() public {
        _initValset({ valSetId: 1, numVals: 3 });

        bytes32 digest = keccak256(abi.encodePacked("hello"));
        XTypes.SigTuple[] memory sigs = _getSignatures({ digest: digest, valSetId: 1 });

        // make a duplicate validator
        sigs[1].validatorAddr = sigs[0].validatorAddr;

        vm.expectRevert("Quorum: duplicate validator");
        quorum.verify({ digest: digest, sigs: sigs, valSetId: 1, qNumerator: 2, qDenominator: 3 });
    }

    function test_verify_sigsNotSorted_reverts() public {
        _initValset({ valSetId: 1, numVals: 3 });

        bytes32 digest = keccak256(abi.encodePacked("hello"));
        XTypes.SigTuple[] memory sigs = _getSignatures({ digest: digest, valSetId: 1 });

        // swap first two sigs
        (sigs[0], sigs[1]) = (sigs[1], sigs[0]);

        vm.expectRevert("Quorum: sigs not sorted");
        quorum.verify({ digest: digest, sigs: sigs, valSetId: 1, qNumerator: 2, qDenominator: 3 });
    }

    function test_verify_noSigs_fails() public {
        _initValset({ valSetId: 1, numVals: 3 });

        bytes32 digest = keccak256(abi.encodePacked("hello"));
        XTypes.SigTuple[] memory sigs = new XTypes.SigTuple[](0);

        bool verified = quorum.verify({ digest: digest, sigs: sigs, valSetId: 1, qNumerator: 2, qDenominator: 3 });

        assertFalse(verified);
    }

    function test_verify_allInvalidSigs_reverts() public {
        _initValset({ valSetId: 1, numVals: 3 });

        bytes32 digest = keccak256(abi.encodePacked("hello"));

        // use different digest
        XTypes.SigTuple[] memory sigs = _getSignatures({ digest: keccak256(abi.encodePacked("world")), valSetId: 1 });

        vm.expectRevert("Quorum: invalid signature");
        quorum.verify({ digest: digest, sigs: sigs, valSetId: 1, qNumerator: 2, qDenominator: 3 });
    }

    //////////////////////////////////////////////////////////////////////////////
    //                              Helpers                                     //
    //////////////////////////////////////////////////////////////////////////////

    /**
     * @dev Initialize a validator set with numVals validators. Sets validator set in QuorumVerifier contract.
     */
    function _initValset(uint64 valSetId, uint64 numVals) internal {
        for (uint64 i = 0; i < numVals; i++) {
            (address addr, uint256 privkey) = makeAddrAndKey(string(abi.encode("val", i)));
            privKeys[addr] = privkey;

            XTypes.Validator memory val = XTypes.Validator(addr, 100);
            validators[valSetId].push(val);
        }

        quorum.setValset(valSetId, validators[valSetId]);
    }

    /**
     * @dev Get a list of signatures for a given digest and validator set.
     */
    function _getSignatures(bytes32 digest, uint64 valSetId) internal view returns (XTypes.SigTuple[] memory) {
        return _getSignatures(digest, valSetId, validators[valSetId].length);
    }

    /**
     * @dev Get a list of signatures for a given digest and validator set, only returning the first
     *      `numSigs` signatures.
     */
    function _getSignatures(bytes32 digest, uint64 valSetId, uint256 numSigs)
        internal
        view
        returns (XTypes.SigTuple[] memory sigs)
    {
        XTypes.Validator[] storage vals = validators[valSetId];
        sigs = new XTypes.SigTuple[](numSigs);

        for (uint256 i = 0; i < numSigs; i++) {
            sigs[i] = XTypes.SigTuple(vals[i].addr, _sign(digest, privKeys[vals[i].addr]));
        }

        _sortSigsByAddr(sigs);
    }

    function _sign(bytes32 digest, uint256 pk) internal pure returns (bytes memory) {
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(pk, digest);
        return abi.encodePacked(r, s, v);
    }

    /**
     * @dev Sort sigs by validator address. XTypes.verifyQuorum expects sigs to be sorted.
     *      Func is not really 'pure', because it modifies the input array in place. But it does not depend on contract state.
     */
    function _sortSigsByAddr(XTypes.SigTuple[] memory sigs) internal pure {
        uint256 n = sigs.length;
        for (uint256 i = 0; i < n - 1; i++) {
            for (uint256 j = 0; j < n - i - 1; j++) {
                if (sigs[j].validatorAddr > sigs[j + 1].validatorAddr) {
                    (sigs[j], sigs[j + 1]) = (sigs[j + 1], sigs[j]);
                }
            }
        }
    }
}
