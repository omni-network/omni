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

    /// @dev register an operator with eigenlayer core
    function _registerAsOperator(address operator) internal {
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });

        _testRegisterAsOperator(operator, operatorDetails);
    }

    /// @dev register an operator with OmniAVS
    function _registerOperatorWithAVS(address operator) internal {
        // only one quorum
        bytes memory quorumNumbers = hex"00";

        // don't matter
        string memory socket = "12.34.56.78:123";
        ISignatureUtils.SignatureWithSaltAndExpiry memory emptySignature;
        IBLSApkRegistry.PubkeyRegistrationParams memory emptyPubkeyRegistrationParams;
        BN254.G1Point memory pubkey = BN254.hashToG1(keccak256(abi.encodePacked("pubkey", operator)));
        blsApkRegistry.setBLSPublicKey(operator, pubkey);

        vm.prank(operator);

        // signature & pubkey registration verification are skipped, see test/avs/eigen/DelegationManagerMock.sol
        registryCoordinator.registerOperator(quorumNumbers, socket, emptyPubkeyRegistrationParams, emptySignature);
    }

    /// @dev create an operator address
    function _operator(uint256 index) internal pure returns (address) {
        return _addr("operator", index);
    }

    /// @dev create a delegator address
    function _delegator(uint256 index) internal pure returns (address) {
        return _addr("delegator", index);
    }

    /// @dev create a namespaced address
    function _addr(string memory namespace, uint256 index) internal pure returns (address) {
        return address(uint160(uint256(keccak256(abi.encodePacked(namespace, index)))));
    }

    /// @dev deposit WETH in eigen layer
    function _depositWeth(address staker, uint256 amount) internal {
        _testDepositWeth(staker, amount);
    }

    /// @dev deposit Eigen in eigen layer
    function _depositEigen(address staker, uint256 amount) internal {
        _testDepositEigen(staker, amount);
    }
}
