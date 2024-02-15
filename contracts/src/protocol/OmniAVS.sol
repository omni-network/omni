// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.12;

import { OwnableUpgradeable } from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";

import { DelegationManager } from "eigenlayer-contracts/src/contracts/core/DelegationManager.sol";
import { IStrategy } from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";
import { IDelegationManager } from "eigenlayer-contracts/src/contracts/interfaces/IDelegationManager.sol";
import { ISignatureUtils } from "eigenlayer-contracts/src/contracts/interfaces/ISignatureUtils.sol";

import { OmniPredeploys } from "../libraries/OmniPredeploys.sol";
import { IOmniPortal } from "../interfaces/IOmniPortal.sol";
import { IOmniEthRestaking } from "../interfaces/IOmniEthRestaking.sol";
import { IOmniAVS } from "../interfaces/IOmniAVS.sol";
import { IOmniAVSAdmin } from "../interfaces/IOmniAVSAdmin.sol";

contract OmniAVS is IOmniAVS, IOmniAVSAdmin, OwnableUpgradeable {
    IDelegationManager public immutable delegation;

    IOmniAVS.StrategyParams[] public strategyParams;

    /// @notice Constant used as a divisor in calculating weights.
    uint256 public constant WEIGHTING_DIVISOR = 1e18;

    /// @dev Omni chain id, used to make xcalls to the Omni chain
    uint64 public omniChainId;

    /// @dev Omni portal contract, used to make xcalls to the Omni chain
    IOmniPortal public omni;

    /// @dev List of currently register operators, used to sync EigenCore
    address[] public operators;

    /// @dev Maximum number of operators that can be registered
    uint32 public maxOperatorCount;

    /// @dev Minimum stake required for an operator to register, not including delegations
    uint96 public minimumOperatorStake;

    constructor(IDelegationManager delegationManager_) {
        delegation = delegationManager_;
        _disableInitializers();
    }

    /// @inheritdoc IOmniAVSAdmin
    function initialize(
        address owner_,
        IOmniPortal omni_,
        uint64 omniChainId_,
        uint96 minimumOperatorStake_,
        uint32 maxOperatorCount_,
        StrategyParams[] calldata strategyParams_
    ) external initializer {
        _transferOwnership(owner_);
        omni = omni_;
        omniChainId = omniChainId_;
        minimumOperatorStake = minimumOperatorStake_;
        maxOperatorCount = maxOperatorCount_;
        _setStrategyParams(strategyParams_);
    }

    /**
     * Omni sync
     */

    /// @inheritdoc IOmniAVS
    function feeForSync() external view returns (uint256) {
        Validator[] memory vals = getValidators();
        return omni.feeFor(omniChainId, abi.encodeWithSelector(IOmniEthRestaking.sync.selector, vals));
    }

    /// @inheritdoc IOmniAVS
    function syncWithOmni() external payable {
        Validator[] memory vals = getValidators();
        omni.xcall{ value: msg.value }(
            omniChainId,
            OmniPredeploys.OMNI_ETH_RESTAKING,
            abi.encodeWithSelector(IOmniEthRestaking.sync.selector, vals)
        );
    }

    /**
     * ServiceManagerBase interface
     */

    /// @notice Forwards a call to EigenLayer's DelegationManager contract to confirm operator registration with the AVS
    /// @param operator The address of the operator to register.
    /// @param operatorSignature The signature, salt, and expiry of the operator's signature.
    function registerOperatorToAVS(
        address operator,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) public virtual {
        require(msg.sender == operator, "OmniAVS: only operator");
        require(operators.length < maxOperatorCount, "OmniAVS: max operators reached");
        require(_getStaked(operator) >= minimumOperatorStake, "OmniAVS: minimum stake not met");
        require(!_isOperator(operator), "OmniAVS: already an operator"); // we could let delegation.regsiterOperatorToAVS handle this, they do check

        delegation.registerOperatorToAVS(operator, operatorSignature);

        _addOperator(operator);
    }

    /// @notice Forwards a call to EigenLayer's DelegationManager contract to confirm operator deregistration from the AVS
    /// @param operator The address of the operator to deregister.
    function deregisterOperatorFromAVS(address operator) public virtual {
        require(msg.sender == operator, "OmniAVS: only operator");
        require(_isOperator(operator), "OmniAVS: not an operator");

        delegation.deregisterOperatorFromAVS(operator);
        _removeOperator(operator);
    }

    /// @notice Sets the metadata URI for the AVS
    /// @param _metadataURI is the metadata URI for the AVS
    /// @dev only callable by the owner
    function setMetadataURI(string memory _metadataURI) public virtual onlyOwner {
        delegation.updateAVSMetadataURI(_metadataURI);
    }

    /**
     * Admin controls
     */

    /// @inheritdoc IOmniAVSAdmin
    function setOmniPortal(IOmniPortal omni_) external onlyOwner {
        omni = omni_;
    }

    /// @inheritdoc IOmniAVSAdmin
    function setOmniChainId(uint64 omniChainId_) external onlyOwner {
        omniChainId = omniChainId_;
    }

    /// @inheritdoc IOmniAVSAdmin
    function setStrategyParams(StrategyParams[] calldata strategyParams_) external onlyOwner {
        _setStrategyParams(strategyParams_);
    }

    function _setStrategyParams(StrategyParams[] calldata strategyParams_) internal {
        delete strategyParams;
        for (uint256 i = 0; i < strategyParams_.length; i++) {
            strategyParams.push(strategyParams_[i]);
        }
    }

    /**
     * View functions
     */

    /// @inheritdoc IOmniAVS
    function getValidators() public view returns (Validator[] memory) {
        return _getValidators();
    }

    /**
     * Internal view functions
     */

    /// @dev Return current list of Validators, including their personal stake and delegated stake
    function _getValidators() internal view returns (Validator[] memory) {
        Validator[] memory vals = new Validator[](operators.length);

        for (uint256 i = 0; i < vals.length; i++) {
            address addr = operators[i];
            uint96 totalStaked;
            StrategyParams memory params;

            // get total opearator stake (their own stake + delegations)
            for (uint256 j = 0; j < strategyParams.length; j++) {
                params = strategyParams[j];

                // shares of the operator in the strategy
                uint256 sharesAmount = delegation.operatorShares(addr, params.strategy);

                // add the weight from the shares for this strategy to the total weight
                if (sharesAmount > 0) totalStaked += _weight(sharesAmount, params.multiplier);
            }

            uint96 staked = _getStaked(addr);
            uint96 delegated = totalStaked - staked;

            vals[i] = Validator(addr, delegated, staked);
        }

        return vals;
    }

    /// @dev Returns the total amount staked by the operator, not including deletations
    function _getStaked(address operator) internal view returns (uint96) {
        (IStrategy[] memory strategies, uint256[] memory shares) =
            DelegationManager(address(delegation)).getDelegatableShares(operator);

        uint96 staked;

        for (uint256 i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];
            uint256 sharesAmt = shares[i];

            // find the strategy params for the strategy
            StrategyParams memory params;
            for (uint256 j = 0; j < strategyParams.length; j++) {
                if (address(strategyParams[j].strategy) == address(strat)) {
                    params = strategyParams[j];
                    break;
                }
            }

            // if strategy is not found, do not consider it in stake
            if (address(params.strategy) == address(0)) continue;

            staked += _weight(sharesAmt, params.multiplier);
        }

        return staked;
    }

    function _weight(uint256 shares, uint96 multiplier) internal pure returns (uint96) {
        return uint96(shares * multiplier / WEIGHTING_DIVISOR);
    }

    /**
     * Internal functions.
     */

    /// @dev Adds an operator to the list of operators, does not check if operator already exists
    function _addOperator(address operator) internal {
        operators.push(operator);
    }

    /// @dev Removes an operator from the list of operators
    function _removeOperator(address operator) internal {
        for (uint256 i = 0; i < operators.length; i++) {
            if (operators[i] == operator) {
                operators[i] = operators[operators.length - 1];
                operators.pop();
                break;
            }
        }
    }

    /// @dev Returns true if the operator is in the list of operators
    function _isOperator(address operator) internal view returns (bool) {
        for (uint256 i = 0; i < operators.length; i++) {
            if (operators[i] == operator) {
                return true;
            }
        }
        return false;
    }
}
