// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { IOmniPortal } from "../interfaces/IOmniPortal.sol";
import { IOmniPortalSys } from "../interfaces/IOmniPortalSys.sol";
import { ConfLevel } from "../libraries/ConfLevel.sol";
import { XRegistryBase } from "./XRegistryBase.sol";
import { XRegistryReplica } from "./XRegistryReplica.sol";

/**
 * @title XRegistry
 * @notice Registry for cross-chain contract deployments. Replicated across each supported chain.
 * @dev This contract is predeployed, and requires storage slots to be set in genesis.
 *      Genesis storage slots must:
 *          - set _owner on proxy
 *          - set _initialized on proxy to 1, to disable the initializer
 *          - set _initialized on implementation to 255, to disabled all initializers
 */
contract XRegistry is OwnableUpgradeable, XRegistryBase {
    /**
     * @notice Register a contract deployment.
     */
    event ContractRegistered(
        uint64 indexed chainId, string indexed name, address indexed registrant, address addr, bytes metadata
    );

    /**
     * @notice xcall gas limit for XRegistryReplica.set
     */
    uint64 public constant XSET_GAS_LIMIT = 150_000;

    /**
     * @notice xcall gas limit for XRegistryReplica.set for portal registrations.
     */
    uint64 public constant XSET_PORTAL_GAS_LIMIT = 250_000;

    /**
     * @notice OmniPortal contract.
     */
    address public portal;

    /**
     * @notice Mapping of chain IDs to the XRegistryReplica address on that chain.
     */
    mapping(uint64 => address) public replicas;

    /**
     * @notice List of supported chain IDs (ie. chain ids with XRegistry replicas)
     */
    uint64[] internal _chainIds;

    /**
     * @notice Register a contract deployment with metadata.
     */
    function register(uint64 chainId, string calldata name, Deployment calldata deployment) external payable {
        _register(chainId, name, deployment);
    }

    /**
     * @notice Register a contract deployment.
     */
    function _register(uint64 chainId, string calldata name, Deployment calldata dep) internal {
        require(isSupportedChain(chainId), "XRegistry: chain not supported");
        require(msg.value >= _syncFee(chainId, name, msg.sender, dep), "XRegistry: insufficient fee");

        _set(chainId, name, msg.sender, dep);
        _sync(chainId, name, msg.sender, dep);

        if (_isPortalRegistration(name, msg.sender) && chainId == IOmniPortal(portal).chainId()) {
            uint64[] memory shards = abi.decode(dep.metadata, (uint64[]));
            IOmniPortalSys(portal).setShards(shards);
        }

        emit ContractRegistered(chainId, name, msg.sender, dep.addr, dep.metadata);
    }

    /**
     * @notice Calculate the fee to register a contract.
     */
    function registrationFee(uint64 chainId, string calldata name, Deployment calldata dep)
        external
        view
        returns (uint256)
    {
        return _syncFee(chainId, name, msg.sender, dep);
    }

    /**
     * @notice Syncs replicas with a new registration for contract `name` by `registrant`.
     *          - Adds all existing registrations of `name` by `registrant` to the replica on `chainId`,
     *          - Adds the new deployment on `chainId` to all replicas with existing registrations
     *            of `name` by `registrant`.
     *          - Add registration to its own chain's replica, so that is has access to its own metadata.
     */
    function _sync(uint64 chainId, string calldata name, address registrant, Deployment calldata dep) internal {
        // sync with self, so that this chain has its own metadata
        _xset(chainId, chainId, name, registrant, dep);

        for (uint256 i = 0; i < _chainIds.length; i++) {
            uint64 otherChainId = _chainIds[i];
            Deployment memory other = _get(otherChainId, name, registrant);

            // if this chain does not have a registration for `name` by `registrant`, do nothing
            if (other.addr == address(0)) continue;

            // don't sync with self, this is done before the loop
            if (otherChainId == chainId) continue;

            // set "other" registration on chain of new registration
            _xset(chainId, otherChainId, name, registrant, other);

            // set new registration on "other" chain
            _xset(otherChainId, chainId, name, registrant, dep);
        }
    }

    /**
     * @notice Calculate the fee to sync replicas with a new registration for contract `name` by `registrant`.
     */
    function _syncFee(uint64 chainId, string calldata name, address registrant, Deployment calldata dep)
        internal
        view
        returns (uint256)
    {
        // initial fee for sync with self
        uint256 fee = _xsetFee(chainId, chainId, name, registrant, dep);

        for (uint256 i = 0; i < _chainIds.length; i++) {
            uint64 otherChainId = _chainIds[i];
            Deployment memory other = _get(otherChainId, name, registrant);

            // if this chain does not have a registration for `name` by `registrant`, do nothing
            if (other.addr == address(0)) continue;

            // don't sync with self, this is done before the loop
            if (otherChainId == chainId) continue;

            // fee to set set "other" registration on chain of new registration
            fee += _xsetFee(chainId, otherChainId, name, registrant, other);

            // fee to set new registration on "other" chain
            fee += _xsetFee(otherChainId, chainId, name, registrant, dep);
        }

        return fee;
    }

    function _xset(uint64 destChainId, uint64 chainId, string calldata name, address registrant, Deployment memory dep)
        internal
    {
        IOmniPortal omni = IOmniPortal(portal);

        // don't xset to self, XRegistry acts as the XRegistryReplica on Omni
        if (destChainId == omni.chainId()) return;

        address replica = replicas[destChainId];
        require(replica != address(0), "XRegistry: unknown chain");

        bytes memory data = abi.encodeWithSelector(XRegistryReplica.set.selector, chainId, name, registrant, dep);
        uint64 gasLimit = _isPortalRegistration(name, registrant) ? XSET_PORTAL_GAS_LIMIT : XSET_GAS_LIMIT;
        uint256 fee = omni.feeFor(destChainId, data, gasLimit);

        omni.xcall{ value: fee }(destChainId, ConfLevel.Finalized, replica, data, gasLimit);
    }

    function _xsetFee(
        uint64 destChainId,
        uint64 chainId,
        string calldata name,
        address registrant,
        Deployment memory dep
    ) internal view returns (uint256) {
        IOmniPortal omni = IOmniPortal(portal);

        // don't xset on omni, XRegistry acts as the XRegistryReplica on Omni
        if (destChainId == omni.chainId()) return 0;

        address replica = replicas[destChainId];
        require(replica != address(0), "XRegistry: unknown chain");

        bytes memory data = abi.encodeWithSelector(XRegistryReplica.set.selector, chainId, name, registrant, dep);
        uint64 gasLimit = _isPortalRegistration(name, registrant) ? XSET_PORTAL_GAS_LIMIT : XSET_GAS_LIMIT;

        return omni.feeFor(destChainId, data, gasLimit);
    }

    /**
     * @notice Return the list of supported chain IDs.
     */
    function chainIds() public view returns (uint64[] memory) {
        return _chainIds;
    }

    /**
     * @notice Return true if the given chain is supported (i.e. has a replica)
     */
    function isSupportedChain(uint64 chainId) public view returns (bool) {
        return replicas[chainId] != address(0);
    }

    /**
     * @notice Set the address of the replica contract on the given chain.
     */
    function setReplica(uint64 chainId, address replica) public onlyOwner {
        require(replicas[chainId] == address(0), "XRegistry: replica already set");
        replicas[chainId] = replica;
        _chainIds.push(chainId);
    }

    /**
     * @notice Set the address of the OmniPortal contract.
     */
    function setPortal(address _portal) public onlyOwner {
        portal = _portal;
    }
}
