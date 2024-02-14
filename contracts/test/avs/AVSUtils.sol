// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { ISignatureUtils } from "eigenlayer-contracts/src/contracts/interfaces/ISignatureUtils.sol";
import { IStrategy } from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";
import { IDelegationManager } from "eigenlayer-contracts/src/contracts/interfaces/IDelegationManager.sol";

import { BN254 } from "eigenlayer-middleware/src/libraries/BN254.sol";
import { IBLSApkRegistry } from "eigenlayer-middleware/src/interfaces/IBLSApkRegistry.sol";
import { OperatorStateRetriever } from "eigenlayer-middleware/src/OperatorStateRetriever.sol";

import { AVSBase } from "./AVSBase.sol";

contract AVSUtils is AVSBase {
    using BN254 for BN254.G1Point;

    // map stake -> strategy -> amount
    mapping(address => mapping(address => uint256)) internal _deposited;

    // map operator addr -> operator id
    mapping(address => bytes32) internal _operatorIds;

    function _operator(uint8 index) internal pure returns (address) {
        return _addr("operator", index);
    }

    function _delegator(uint8 index) internal pure returns (address) {
        return _addr("delegator", index);
    }

    function _addr(string memory namespace, uint8 index) internal pure returns (address) {
        return address(uint160(uint256(keccak256(abi.encodePacked(namespace, index)))));
    }

    function _registerAsOperator(address operator) internal {
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });

        _testRegisterAsOperator(operator, operatorDetails);
    }

    function _registerOperatorWithAVS(address operator) internal {
        // only one quorum
        bytes memory quorumNumbers = hex"00";

        // don't matter
        string memory socket = "12.34.56.78:123";
        ISignatureUtils.SignatureWithSaltAndExpiry memory emptySignature;
        IBLSApkRegistry.PubkeyRegistrationParams memory emptyPubkeyRegistrationParams;
        BN254.G1Point memory pubkey = BN254.hashToG1(keccak256(abi.encodePacked("pubkey", operator)));
        blsApkRegistry.setBLSPublicKey(operator, pubkey);

        _operatorIds[operator] = pubkey.hashG1Point();

        vm.prank(operator);

        // signature & pubkey registration verification are skipped, see test/avs/eigen/DelegationManagerMock.sol
        registryCoordinator.registerOperator(quorumNumbers, socket, emptyPubkeyRegistrationParams, emptySignature);
    }

    function _depositWeth(address staker, uint256 amount) internal {
        _testDepositWeth(staker, amount);
        _deposited[staker][address(wethStrat)] += amount;
    }

    function _wethDeposits(address staker) internal view returns (uint256) {
        return _deposited[staker][address(wethStrat)];
    }

    function _depositEigen(address staker, uint256 amount) internal {
        _testDepositEigen(staker, amount);
        _deposited[staker][address(eigenStrat)] += amount;
    }

    function _eigenDeposits(address staker) internal view returns (uint256) {
        return _deposited[staker][address(eigenStrat)];
    }
}
