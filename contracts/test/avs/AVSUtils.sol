// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import { ISignatureUtils } from "eigenlayer-contracts/src/contracts/interfaces/ISignatureUtils.sol";
import { IStrategy } from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";
import { IDelegationManager } from "eigenlayer-contracts/src/contracts/interfaces/IDelegationManager.sol";

import { EigenPodManagerHarness } from "./eigen/EigenPodManagerHarness.sol";
import { IOmniAVS } from "src/interfaces/IOmniAVS.sol";
import { AVSBase } from "./AVSBase.sol";

contract AVSUtils is AVSBase {
    /// Canonical, virtual beacon chain ETH strategy
    address constant beaconChainETHStrategy = 0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0;

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

    /// @dev deposit into a random strategy, that is part of the OmniAVS strategy params
    function _depositIntoSupportedStrategy(address staker, uint256 shares) internal {
        IOmniAVS.StrategyParams[] memory params = omniAVS.strategyParams();
        uint256 index = uint256(keccak256(abi.encodePacked(staker))) % params.length;
        _depositIntoStrategy(staker, shares, address(params[index].strategy));
    }

    /// @dev deposit into an that is NOT part of the OmniAVS strategy params
    function _depositIntoUnsupportedStrategy(address staker, uint256 shares) internal {
        IOmniAVS.StrategyParams[] memory params = omniAVS.strategyParams();

        // check that unsupportedStrategy is not part of the strategy params
        for (uint256 i = 0; i < params.length; i++) {
            require(
                address(params[i].strategy) != address(unsupportedStrat),
                "AVSUtils: unsupportedStrat should not be in strategy params"
            );
        }

        _depositIntoStrategy(staker, shares, address(unsupportedStrat));
    }

    /// @dev deposit into the provided strategy
    function _depositIntoStrategy(address staker, uint256 shares, address strategy) internal {
        if (strategy == beaconChainETHStrategy) {
            _depositBeaconEth(staker, shares);
            return;
        }

        IStrategy strat = IStrategy(strategy);

        // when running fork tests, some strategies (like stETH), do not map underlying tokens to shares 1:1
        // so we need to figure out how much underlying to deposit to get the correct amount of shares
        IERC20 underlying = strat.underlyingToken();
        uint256 underlyingAmt = strat.sharesToUnderlying(shares);

        // sometimes underlyingToShares(sharesToUnderlying(x)) != x (for some strategies like stETH)
        // so we keep incrementing underlyingAmt until sharesToUnderlying(underlyingAmt) == shares
        while (shares != strat.underlyingToShares(underlyingAmt)) {
            underlyingAmt = underlyingAmt + 1;
        }

        _testDepositToStrategy(staker, underlyingAmt, underlying, strat);
    }

    /// @dev Deposit beacon eth
    function _depositBeaconEth(address staker, uint256 amount) internal {
        eigenPodManager.updatePodOwnerShares(staker, int256(amount));
    }
}
