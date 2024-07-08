// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { OwnableUpgradeable } from "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import { PausableUpgradeable } from "@openzeppelin-upgrades/contracts/security/PausableUpgradeable.sol";

import { IAVSDirectory } from "eigenlayer-contracts/src/contracts/interfaces/IAVSDirectory.sol";
import { IStrategy } from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";
import { ISignatureUtils } from "eigenlayer-contracts/src/contracts/interfaces/ISignatureUtils.sol";

import { Secp256k1 } from "./libraries/Secp256k1.sol";
import { IOmniAVS } from "./interfaces/IOmniAVS.sol";
import { IOmniAVSAdmin } from "./interfaces/IOmniAVSAdmin.sol";
import { IDelegationManager } from "./ext/IDelegationManager.sol";

import { IEthStakeInbox } from "core/interfaces/IEthStakeInbox.sol";
import { IOmniPortal } from "core/interfaces/IOmniPortal.sol";
import { ConfLevel } from "core/libraries/ConfLevel.sol";

import { OmniAVSStorage } from "./OmniAVSStorage.sol";

/**
 * @title OmniAVS
 * @notice Omni's AVS contract. It is responsible for facilitating registration / deregistration of
 *         EigenLayer operators, and for syncing operator delegations with the Omni chain.
 */
contract OmniAVS is IOmniAVS, IOmniAVSAdmin, OwnableUpgradeable, PausableUpgradeable, OmniAVSStorage {
    /// @notice Constant used as a divisor in calculating weights
    uint256 internal constant STRATEGY_WEIGHTING_DIVISOR = 1e18;

    /// @notice EigenLayer core DelegationManager
    IDelegationManager internal immutable _delegationManager;

    /// @notice EigenLayer core AVSDirectory
    IAVSDirectory internal immutable _avsDirectory;

    constructor(IDelegationManager delegationManager_, IAVSDirectory avsDirectory_) {
        _delegationManager = delegationManager_;
        _avsDirectory = avsDirectory_;
        _disableInitializers();
    }

    /**
     * @notice Initialize the Omni AVS admin contract.
     * @param owner_            Initial contract owner
     * @param omni_             Omni portal contract
     * @param omniChainId_      Omni chain ID
     * @param ethStakeInbox_    EthStakeInbox contract address
     * @param minOperatorStake_ Minimum operator stake
     * @param maxOperatorCount_ Maximum operator count
     * @param strategyParams_   List of accepted strategies and their multipliers
     * @param metadataURI_      Metadata URI for the AVS
     * @param allowlistEnabled_ Whether the allowlist is enabled
     */
    function initialize(
        address owner_,
        IOmniPortal omni_,
        uint64 omniChainId_,
        address ethStakeInbox_,
        uint96 minOperatorStake_,
        uint32 maxOperatorCount_,
        StrategyParam[] calldata strategyParams_,
        string calldata metadataURI_,
        bool allowlistEnabled_
    ) external initializer {
        _setOmniPortal(omni_);
        _setOmniChainId(omniChainId_);
        _setXCallGasLimits({ base: 75_000, perOperator: 50_000 });
        _setEthStakeInbox(ethStakeInbox_);
        _setMinOperatorStake(minOperatorStake_);
        _setMaxOperatorCount(maxOperatorCount_);
        _setStrategyParams(strategyParams_);

        // if necessary, enable allowlist
        if (allowlistEnabled_) {
            _enableAllowlist();
        }

        // if provided, set metadata URI
        if (bytes(metadataURI_).length > 0) {
            _avsDirectory.updateAVSMetadataURI(metadataURI_);
        }

        _transferOwnership(owner_);
    }

    //////////////////////////////////////////////////////////////////////////////
    //                          Operator Registration                           //
    //////////////////////////////////////////////////////////////////////////////

    /**
     * @notice Register an operator with the AVS. Forwards call to EigenLayer' AVSDirectory.
     * @param pubkey            64 byte uncompressed secp256k1 public key (no 0x04 prefix)
     *                          Pubkey must match operator's address (msg.sender)
     * @param operatorSignature The signature, salt, and expiry of the operator's signature.
     */
    function registerOperator(
        bytes calldata pubkey,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) external whenNotPaused {
        address operator = msg.sender;

        require(operator == Secp256k1.pubkeyToAddress(pubkey), "OmniAVS: pubkey != sender");
        require(!allowlistEnabled || _allowlist[operator], "OmniAVS: not allowed");
        require(!_isOperator(operator), "OmniAVS: already an operator");
        require(_operators.length < maxOperatorCount, "OmniAVS: max operators reached");
        require(_getTotalDelegations(operator) >= minOperatorStake, "OmniAVS: min stake not met");

        _addOperator(operator, pubkey);
        _avsDirectory.registerOperatorToAVS(operator, operatorSignature);

        emit OperatorAdded(operator);
    }

    /**
     * @notice Returns true if the operator is in the allowlist.
     * @param operator The operator to check
     */
    function isInAllowlist(address operator) external view returns (bool) {
        return _allowlist[operator];
    }

    /**
     * @notice Returns the EigenLayer AVSDirectory contract.
     * @dev Implemented to match IServiceManager interface - required for compatibility with
     *      eigenlayer frontend.
     */
    function avsDirectory() external view returns (address) {
        return address(_avsDirectory);
    }

    //////////////////////////////////////////////////////////////////////////////
    //                              Omni Sync                                   //
    //////////////////////////////////////////////////////////////////////////////

    /**
     * @notice Sync OmniAVS operator stake & delegations with Omni chain.
     */
    function syncWithOmni() external payable whenNotPaused {
        Operator[] memory ops = _getOperators();
        omni.xcall{ value: msg.value }(
            omniChainId,
            ConfLevel.Finalized,
            ethStakeInbox,
            abi.encodeWithSelector(IEthStakeInbox.sync.selector, ops),
            _xcallGasLimitFor(ops.length)
        );
    }

    /**
     * @notice Returns the fee required for syncWithOmni(), for the current operator set.
     */
    function feeForSync() external view returns (uint256) {
        Operator[] memory ops = _getOperators();
        return omni.feeFor(
            omniChainId, abi.encodeWithSelector(IEthStakeInbox.sync.selector, ops), _xcallGasLimitFor(ops.length)
        );
    }

    //////////////////////////////////////////////////////////////////////////////
    //                              AVS Views                                   //
    //////////////////////////////////////////////////////////////////////////////

    /**
     * @notice Returns the currrent list of operator registered as OmniAVS.
     *         Operator.addr        = The operator's ethereum address
     *         Operator.pubkey      = The operator's 64 byte uncompressed secp256k1 public key
     *         Operator.staked      = The total amount staked by the operator, not including delegations
     *         Operator.delegated   = The total amount delegated, not including operator stake
     */
    function operators() external view returns (Operator[] memory) {
        return _getOperators();
    }

    /**
     * @notice Returns the current strategy parameters. Strategy parameters determine which
     *         eigenlayer strateies the AVS considers when determining operator stake.
     */
    function strategyParams() external view returns (StrategyParam[] memory) {
        return _strategyParams;
    }

    /**
     * @notice Returns the list of strategies that the AVS supports for restaking.
     * @dev Implemented to match IServiceManager interface - required for compatibility with
     *      eigenlayer frontend.
     */
    function getRestakeableStrategies() external view returns (address[] memory) {
        return _getRestakeableStrategies();
    }

    /**
     * @notice Returns the list of strategies that the operator has potentially restaked on the AVS
     * @dev Implemented to match IServiceManager interface - required for compatibility with
     *      eigenlayer frontend.
     *
     *      This function is intended to be called off-chain
     *
     *      No guarantee is made on whether the operator has shares for a strategy. The off-chain
     *      service should do that validation separately. This matches the behavior defined in
     *      eigenlayer-middleware's ServiceManagerBase.
     *
     * @param operator The address of the operator to get restaked strategies for
     */
    function getOperatorRestakedStrategies(address operator) external view returns (address[] memory) {
        if (!_isOperator(operator)) return new address[](0);
        return _getRestakeableStrategies();
    }

    /**
     * @notice Check if an operator can register to the AVS.
     *         Returns true, with no reason, if the operator can register to the AVS.
     *         Returns false, with a reason, if the operator cannot register to the AVS.
     * @dev This function is intented to be called off-chain.
     * @param operator The operator to check
     * @return canRegister True if the operator can register, false otherwise
     * @return reason      The reason the operator cannot register. Empty if canRegister is true.
     */
    function canRegister(address operator) external view returns (bool, string memory) {
        if (!_delegationManager.isOperator(operator)) return (false, "not an operator");
        if (allowlistEnabled && !_allowlist[operator]) return (false, "not in allowlist");
        if (_isOperator(operator)) return (false, "already registered");
        if (_operators.length >= maxOperatorCount) return (false, "max operators reached");
        if (_getTotalDelegations(operator) < minOperatorStake) return (false, "min stake not met");
        return (true, "");
    }

    //////////////////////////////////////////////////////////////////////////////
    //                              Admin functions                             //
    //////////////////////////////////////////////////////////////////////////////

    /**
     * @notice Sets AVS metadata URI with the AVSDirectory.
     */
    function setMetadataURI(string memory metadataURI) external onlyOwner {
        _avsDirectory.updateAVSMetadataURI(metadataURI);
    }

    /**
     * @notice Set the Omni portal contract.
     * @param portal The Omni portal contract
     */
    function setOmniPortal(IOmniPortal portal) external onlyOwner {
        _setOmniPortal(portal);
    }

    /**
     * @notice Set the Omni chain ID.
     * @param chainId The Omni chain ID
     */
    function setOmniChainId(uint64 chainId) external onlyOwner {
        _setOmniChainId(chainId);
    }

    /**
     * @notice Set the EthStakeInbox contract address.
     * @param inbox The EthStakeInbox contract address
     */
    function setEthStakeInbox(address inbox) external onlyOwner {
        _setEthStakeInbox(inbox);
    }

    /**
     * @notice Set the strategy parameters.
     * @param params The strategy parameters
     */
    function setStrategyParams(StrategyParam[] calldata params) external onlyOwner {
        _setStrategyParams(params);
    }

    /**
     * @notice Set the minimum operator stake.
     * @param stake The minimum operator stake, not including delegations
     */
    function setMinOperatorStake(uint96 stake) external onlyOwner {
        _setMinOperatorStake(stake);
    }

    /**
     * @notice Set the maximum operator count.
     * @param count The maximum operator count
     */
    function setMaxOperatorCount(uint32 count) external onlyOwner {
        _setMaxOperatorCount(count);
    }

    /**
     * @notice Set the xcall gas limits.
     * @param base          The base xcall gas limit
     * @param perOperator   The per-operator additional xcall gas limit
     */
    function setXCallGasLimits(uint64 base, uint64 perOperator) external onlyOwner {
        _setXCallGasLimits(base, perOperator);
    }

    /**
     * @notice Add an operator to the allowlist.
     * @param operator The operator to add
     */
    function addToAllowlist(address operator) external onlyOwner {
        require(operator != address(0), "OmniAVS: zero address");
        require(!_allowlist[operator], "OmniAVS: already in allowlist");
        _allowlist[operator] = true;
        emit OperatorAllowed(operator);
    }

    /**
     * @notice Remove an operator from the allowlist.
     * @param operator The operator to remove
     */
    function removeFromAllowlist(address operator) external onlyOwner {
        require(_allowlist[operator], "OmniAVS: not in allowlist");
        _allowlist[operator] = false;
        emit OperatorDisallowed(operator);
    }

    /**
     * @notice Enable the allowlist.
     */
    function enableAllowlist() external onlyOwner {
        _enableAllowlist();
    }

    /**
     * @notice Disable the allowlist.
     */
    function disableAllowlist() external onlyOwner {
        _disableAllowlist();
    }

    /**
     * @notice Eject an operator from the AVS.
     */
    function ejectOperator(address operator) external onlyOwner {
        _deregisterOperator(operator);
    }

    /**
     * @notice Pause the contract.
     * @dev This pauses registerOperatorToAVS, deregisterOperatorFromAVS, and syncWithOmni.
     */
    function pause() external onlyOwner {
        _pause();
    }

    /**
     * @notice Unpause the contract.
     */
    function unpause() external onlyOwner {
        _unpause();
    }

    //////////////////////////////////////////////////////////////////////////////
    //                              Internal setters                            //
    //////////////////////////////////////////////////////////////////////////////

    /**
     * @notice Deregister an operator from the AVS. Forwards a call to EigenLayer's AVSDirectory.
     */
    function _deregisterOperator(address operator) private {
        require(_isOperator(operator), "OmniAVS: not an operator");

        _removeOperator(operator);
        _avsDirectory.deregisterOperatorFromAVS(operator);

        emit OperatorRemoved(operator);
    }

    /**
     * @notice Add an operator to internal AVS state (_operators, _operatorPubkeys)
     * @dev Does not check if operator already exists
     */
    function _addOperator(address operator, bytes calldata pubkey) private {
        _operators.push(operator);
        _operatorPubkeys[operator] = pubkey;
    }

    /**
     * @notice Removes an operator from internal AVS state (_operators, _operatorPubkeys)
     * @dev Does not check if operator exists
     */
    function _removeOperator(address operator) private {
        for (uint256 i = 0; i < _operators.length;) {
            if (_operators[i] == operator) {
                _operators[i] = _operators[_operators.length - 1];
                _operators.pop();
                break;
            }
            unchecked {
                i++;
            }
        }
        delete _operatorPubkeys[operator];
    }

    /**
     * @notice Set the Omni portal contract.
     * @param portal The Omni portal contract
     */
    function _setOmniPortal(IOmniPortal portal) private {
        require(address(portal) != address(0), "OmniAVS: no zero portal");
        omni = portal;
        emit OmniPortalSet(address(portal));
    }

    /**
     * @notice Set the Omni chain ID.
     * @param chainId The Omni chain ID
     */
    function _setOmniChainId(uint64 chainId) private {
        omniChainId = chainId;
        emit OmniChainIdSet(chainId);
    }

    /**
     * @notice Set the EthStakeInbox contract address.
     * @param inbox The EthStakeInbox contract address
     */
    function _setEthStakeInbox(address inbox) private {
        require(inbox != address(0), "OmniAVS: no zero inbox");
        ethStakeInbox = inbox;
        emit EthStakeInboxSet(inbox);
    }

    /**
     * @notice Set the minimum operator stake.
     * @param stake The minimum operator stake, not including delegations
     */
    function _setMinOperatorStake(uint96 stake) private {
        minOperatorStake = stake;
        emit MinOperatorStakeSet(stake);
    }

    /**
     * @notice Set the maximum operator count.
     * @param count The maximum operator count
     */
    function _setMaxOperatorCount(uint32 count) private {
        maxOperatorCount = count;
        emit MaxOperatorCountSet(count);
    }

    /**
     * @notice Set the xcall gas limits.
     * @param base          The base xcall gas limit
     * @param perOperator   The per-operator additional xcall gas limit
     */
    function _setXCallGasLimits(uint64 base, uint64 perOperator) private {
        xcallBaseGasLimit = base;
        xcallGasLimitPerOperator = perOperator;
        emit XCallGasLimitsSet(base, perOperator);
    }

    /**
     * @notice Enable the allowlist.
     */
    function _enableAllowlist() private {
        require(!allowlistEnabled, "OmniAVS: already enabled");
        allowlistEnabled = true;
        emit AllowlistEnabled();
    }

    /**
     * @notice Disable the allowlist.
     */
    function _disableAllowlist() private {
        require(allowlistEnabled, "OmniAVS: already disabled");
        allowlistEnabled = false;
        emit AllowlistDisabled();
    }

    /**
     * @notice Set the strategy parameters.
     * @param params The strategy parameters
     */
    function _setStrategyParams(StrategyParam[] calldata params) private {
        delete _strategyParams;

        for (uint256 i = 0; i < params.length;) {
            require(address(params[i].strategy) != address(0), "OmniAVS: no zero strategy");
            require(params[i].multiplier > 0, "OmniAVS: no zero multiplier");

            // ensure no duplicates
            for (uint256 j = i + 1; j < params.length;) {
                require(address(params[i].strategy) != address(params[j].strategy), "OmniAVS: no duplicate strategy");
                unchecked {
                    j++;
                }
            }

            _strategyParams.push(params[i]);
            unchecked {
                i++;
            }
        }

        emit StrategyParamsSet(params);
    }

    //////////////////////////////////////////////////////////////////////////////
    //                              Internal views                              //
    //////////////////////////////////////////////////////////////////////////////

    /**
     * @notice Returns the gas limit for OmniEthRestaking.sync xcall for some number of operators
     */
    function _xcallGasLimitFor(uint256 numOperators) internal view returns (uint64) {
        return uint64(numOperators) * xcallGasLimitPerOperator + xcallBaseGasLimit;
    }

    /**
     * @notice Returns true if the operator is in the list of operators
     */
    function _isOperator(address operator) private view returns (bool) {
        return _operatorPubkeys[operator].length > 0;
    }

    /**
     * @notice Return current list of Operators, including their personal stake and delegated stake
     */
    function _getOperators() internal view returns (Operator[] memory) {
        Operator[] memory ops = new Operator[](_operators.length);

        for (uint256 i = 0; i < ops.length;) {
            address operator = _operators[i];

            uint96 total = _getTotalDelegations(operator);
            uint96 staked = _getSelfDelegations(operator);

            // this should never happen, but just in case
            uint96 delegated = total > staked ? total - staked : 0;
            bytes memory pubkey = _operatorPubkeys[operator];

            ops[i] = Operator(operator, pubkey, delegated, staked);
            unchecked {
                i++;
            }
        }

        return ops;
    }

    /**
     * @notice Returns the operator's self-delegations
     * @param operator The operator address
     */
    function _getSelfDelegations(address operator) internal view returns (uint96) {
        (IStrategy[] memory strategies, uint256[] memory shares) = _delegationManager.getDelegatableShares(operator);

        uint96 staked;
        for (uint256 i = 0; i < strategies.length; i++) {
            IStrategy strat = strategies[i];

            // find the strategy params for the strategy
            StrategyParam memory params;
            for (uint256 j = 0; j < _strategyParams.length;) {
                if (address(_strategyParams[j].strategy) == address(strat)) {
                    params = _strategyParams[j];
                    break;
                }
                unchecked {
                    j++;
                }
            }

            // if strategy is not found, do not consider it in stake
            if (address(params.strategy) == address(0)) continue;

            staked += _weight(shares[i], params.multiplier);
        }

        return staked;
    }

    /**
     * @notice Returns total delegations to the operator, including self delegations
     * @param operator The operator address
     */
    function _getTotalDelegations(address operator) internal view returns (uint96) {
        uint96 total;
        StrategyParam memory params;

        for (uint256 i = 0; i < _strategyParams.length;) {
            params = _strategyParams[i];
            uint256 shares = _delegationManager.operatorShares(operator, params.strategy);

            total += _weight(shares, params.multiplier);
            unchecked {
                i++;
            }
        }

        return total;
    }

    /**
     * @notice Returns the weighted stake for shares with specified multiplier
     */
    function _weight(uint256 shares, uint96 multiplier) internal pure returns (uint96) {
        return uint96(shares * multiplier / STRATEGY_WEIGHTING_DIVISOR);
    }

    /**
     * @notice Returns the list of restakeable strategy addresses
     */
    function _getRestakeableStrategies() internal view returns (address[] memory) {
        address[] memory strategies = new address[](_strategyParams.length);
        for (uint256 i = 0; i < _strategyParams.length;) {
            strategies[i] = address(_strategyParams[i].strategy);
            unchecked {
                i++;
            }
        }
        return strategies;
    }
}
