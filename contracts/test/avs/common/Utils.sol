// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import { ISignatureUtils } from "eigenlayer-contracts/src/contracts/interfaces/ISignatureUtils.sol";
import { IStrategy } from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";
import { IDelegationManager } from "eigenlayer-contracts/src/contracts/interfaces/IDelegationManager.sol";

import { IOmniAVS } from "src/interfaces/IOmniAVS.sol";

import { EigenPodManagerHarness } from "./eigen/EigenPodManagerHarness.sol";
import { Fixtures } from "./Fixtures.sol";
import { Vm } from "forge-std/Vm.sol";

/**
 * @title Utils
 * @dev Common utilities for AVS tests
 */
contract Utils is Fixtures {
    // map addr to private key
    mapping(address => Vm.Wallet) _wallets;

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
        omniAVS.registerOperator(_pubkey(operator), sig);
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

    /// @dev disable the allowlist
    function _disableAllowlist() internal {
        vm.prank(omniAVSOwner);
        omniAVS.disableAllowlist();
    }

    /// @dev enable the allowlist
    function _enableAllowlist() internal {
        vm.prank(omniAVSOwner);
        omniAVS.enableAllowlist();
    }

    /// @dev eject an operator from OmniAVS
    function _ejectOperatorFromAVS(address operator) internal {
        vm.prank(omniAVSOwner);
        omniAVS.ejectOperator(operator);
    }

    /// @dev create an operator address
    function _operator(uint256 index) internal returns (address) {
        return _addr("operator", index);
    }

    /// @dev create a delegator address
    function _delegator(uint256 index) internal returns (address) {
        return _addr("delegator", index);
    }

    /// @dev create a namespaced address, save the addresses wallet
    function _addr(string memory namespace, uint256 index) internal returns (address) {
        Vm.Wallet memory w = vm.createWallet(uint256(keccak256(abi.encode(namespace, index))));
        _wallets[w.addr] = w;
        return w.addr;
    }

    /// @dev return the public key of an operator
    function _pubkey(address operator) internal view returns (bytes memory) {
        Vm.Wallet memory w = _wallets[operator];
        return abi.encodePacked(w.publicKeyX, w.publicKeyY);
    }

    /// @dev sign a digest
    function _sign(address signer, bytes32 digest) internal view returns (bytes memory) {
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(_wallets[signer].privateKey, digest);
        return abi.encodePacked(r, s, v);
    }

    /// @dev deposit into a random strategy, that is part of the OmniAVS strategy params
    function _depositIntoSupportedStrategy(address staker, uint256 shares) internal {
        IOmniAVS.StrategyParam[] memory params = omniAVS.strategyParams();
        uint256 index = uint256(keccak256(abi.encodePacked(staker))) % params.length;
        _depositIntoStrategy(staker, shares, address(params[index].strategy));
    }

    /// @dev deposit into an that is NOT part of the OmniAVS strategy params
    function _depositIntoUnsupportedStrategy(address staker, uint256 shares) internal {
        IOmniAVS.StrategyParam[] memory params = omniAVS.strategyParams();

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

    /**
     *
     * Utils repurposed from eignlayer-contracts/src/test/EigenLayerTestHelper.sol
     * Each util both performs and verifies a specific action (hence _test prefix)
     *
     */

    /**
     * @notice Register 'sender' as an operator, setting their 'OperatorDetails' in DelegationManager to 'operatorDetails', verifies
     * that the storage of DelegationManager contract is updated appropriately
     *
     * @param sender is the address being registered as an operator
     * @param operatorDetails is the `sender`'s OperatorDetails struct
     */
    function _testRegisterAsOperator(address sender, IDelegationManager.OperatorDetails memory operatorDetails)
        internal
    {
        vm.startPrank(sender);
        string memory emptyStringForMetadataURI;
        delegation.registerAsOperator(operatorDetails, emptyStringForMetadataURI);
        assertTrue(delegation.isOperator(sender), "testRegisterAsOperator: sender is not a operator");

        assertTrue(
            keccak256(abi.encode(delegation.operatorDetails(sender))) == keccak256(abi.encode(operatorDetails)),
            "_testRegisterAsOperator: operatorDetails not set appropriately"
        );

        assertTrue(delegation.isDelegated(sender), "_testRegisterAsOperator: sender not marked as actively delegated");
        vm.stopPrank();
    }

    /**
     * @notice Deposits `amountToDeposit` of `underlyingToken` from address `sender` into `stratToDepositTo`.
     * *If*  `sender` has zero shares prior to deposit, *then* checks that `stratToDepositTo` is correctly added to their `stakerStrategyList` array.
     *
     * @param sender The address to spoof calls from using `vm.startPrank(sender)`
     * @param amountToDeposit Amount of WETH that is first *transferred from this contract to `sender`* and then deposited by `sender` into `stratToDepositTo`
     */
    function _testDepositToStrategy(
        address sender,
        uint256 amountToDeposit,
        IERC20 underlyingToken,
        IStrategy stratToDepositTo
    ) internal returns (uint256 amountDeposited) {
        // deposits will revert when amountToDeposit is 0
        vm.assume(amountToDeposit > 0);

        // whitelist the strategy for deposit, in case it wasn't before
        {
            vm.startPrank(strategyManager.strategyWhitelister());
            IStrategy[] memory _strategy = new IStrategy[](1);
            bool[] memory _thirdPartyTransfersForbiddenValues = new bool[](1);
            _strategy[0] = stratToDepositTo;
            strategyManager.addStrategiesToDepositWhitelist(_strategy, _thirdPartyTransfersForbiddenValues);
            vm.stopPrank();
        }

        uint256 operatorSharesBefore = strategyManager.stakerStrategyShares(sender, stratToDepositTo);
        uint256 expectedSharesOut = stratToDepositTo.underlyingToShares(amountToDeposit);

        deal(address(underlyingToken), sender, amountToDeposit);
        vm.startPrank(sender);
        underlyingToken.approve(address(strategyManager), type(uint256).max);
        strategyManager.depositIntoStrategy(stratToDepositTo, underlyingToken, amountToDeposit);
        amountDeposited = amountToDeposit;

        //check if depositor has never used this strat, that it is added correctly to stakerStrategyList array.
        if (operatorSharesBefore == 0) {
            // check that strategy is appropriately added to dynamic array of all of sender's strategies
            assertTrue(
                strategyManager.stakerStrategyList(sender, strategyManager.stakerStrategyListLength(sender) - 1)
                    == stratToDepositTo,
                "_testDepositToStrategy: stakerStrategyList array updated incorrectly"
            );
        }

        // check that the shares out match the expected amount out
        assertEq(
            strategyManager.stakerStrategyShares(sender, stratToDepositTo) - operatorSharesBefore,
            expectedSharesOut,
            "_testDepositToStrategy: actual shares out should match expected shares out"
        );

        vm.stopPrank();
    }

    /**
     * @notice tries to delegate from 'staker' to 'operator', verifies that staker has at least some shares
     * delegatedShares update correctly for 'operator' and delegated status is updated correctly for 'staker'
     * @param staker the staker address to delegate from
     * @param operator the operator address to delegate to
     */
    function _testDelegateToOperator(address staker, address operator) internal {
        //staker-specific information
        (IStrategy[] memory delegateStrategies, uint256[] memory delegateShares) = strategyManager.getDeposits(staker);

        uint256 numStrats = delegateShares.length;
        assertTrue(numStrats != 0, "_testDelegateToOperator: delegating from address with no deposits");
        uint256[] memory inititalSharesInStrats = new uint256[](numStrats);
        for (uint256 i = 0; i < numStrats; ++i) {
            inititalSharesInStrats[i] = delegation.operatorShares(operator, delegateStrategies[i]);
        }

        vm.startPrank(staker);
        ISignatureUtils.SignatureWithExpiry memory signatureWithExpiry;
        delegation.delegateTo(operator, signatureWithExpiry, bytes32(0));
        vm.stopPrank();

        assertTrue(
            delegation.delegatedTo(staker) == operator,
            "_testDelegateToOperator: delegated address not set appropriately"
        );
        assertTrue(delegation.isDelegated(staker), "_testDelegateToOperator: delegated status not set appropriately");

        for (uint256 i = 0; i < numStrats; ++i) {
            uint256 operatorSharesBefore = inititalSharesInStrats[i];
            uint256 operatorSharesAfter = delegation.operatorShares(operator, delegateStrategies[i]);
            assertTrue(
                operatorSharesAfter == (operatorSharesBefore + delegateShares[i]),
                "_testDelegateToOperator: delegatedShares not increased correctly"
            );
        }
    }
}
