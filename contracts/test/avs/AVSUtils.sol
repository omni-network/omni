// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { ISignatureUtils } from "eigenlayer-contracts/src/contracts/interfaces/ISignatureUtils.sol";
import { IStrategy } from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";
import { IDelegationManager } from "eigenlayer-contracts/src/contracts/interfaces/IDelegationManager.sol";

import { AVSBase } from "./AVSBase.sol";

contract AVSUtils is AVSBase {
    // map addr to private key
    mapping(address => uint256) _pks;

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
        // don't matter
        ISignatureUtils.SignatureWithSaltAndExpiry memory sig = ISignatureUtils.SignatureWithSaltAndExpiry({
            signature: new bytes(0),
            salt: keccak256(abi.encodePacked(operator)),
            expiry: block.timestamp + 1 days
        });

        bytes32 operatorRegistrationDigestHash = avsDirectory.calculateOperatorAVSRegistrationDigestHash({
            operator: operator,
            avs: address(omniAVS),
            salt: sig.salt,
            expiry: sig.expiry
        });

        sig.signature = _sign(operator, operatorRegistrationDigestHash);

        vm.prank(operator);
        omniAVS.registerOperatorToAVS(operator, sig);
    }

    /// @dev add an operator to the allowlist
    function _addToAllowlist(address operator) internal {
        vm.prank(omniAVSOwner);
        omniAVS.addToAllowlist(operator);
    }

    /// @dev remove an operator from the allowlist
    function _removeFromAllowlist(address operator) internal {
        vm.prank(omniAVSOwner);
        omniAVS.removeFromAllowlist(operator);
    }

    /// @dev set allow list
    function _setAllowlist(address[] memory allowlist) internal {
        vm.prank(omniAVSOwner);
        omniAVS.setAllowlist(allowlist);
    }

    /// @dev deregister an operator from OmniAVS
    function _deregisterOperatorFromAVS(address operator) internal {
        vm.prank(operator);
        omniAVS.deregisterOperatorFromAVS(operator);
    }

    /// @dev create an operator address
    function _operator(uint256 index) internal returns (address) {
        return _addr("operator", index);
    }

    /// @dev create a delegator address
    function _delegator(uint256 index) internal returns (address) {
        return _addr("delegator", index);
    }

    /// @dev create a namespaced address
    function _addr(string memory namespace, uint256 index) internal returns (address) {
        (address addr, uint256 pk) = makeAddrAndKey(string(abi.encodePacked(namespace, index)));
        _pks[addr] = pk;
        return addr;
    }

    /// @dev sign a digest
    function _sign(address signer, bytes32 digest) internal view returns (bytes memory) {
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(_pks[signer], digest);
        return abi.encodePacked(r, s, v);
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
