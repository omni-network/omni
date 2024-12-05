// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { ConfLevel } from "../libraries/ConfLevel.sol";

/**
 * @title PortalRegistry
 * @notice Registry for OmniPortal deployments. Predeployed on Omni's EVM.
 * @dev This contract is predeployed, and requires storage slots to be set in genesis.
 *      initialize(...) is called pre-deployment, in script/genesis/AllocPredeploys.s.sol
 *      Initializers on the implementation are disabled via manual storage updates, rather than in a constructor.
 *      If an new implementation is required, a constructor should be added.
 */
contract PortalRegistry is OwnableUpgradeable {
    /**
     * @notice Emitted when a new OmniPortal deployment is registered.
     */
    event PortalRegistered(
        uint64 indexed chainId,
        bytes32 indexed addr,
        uint64 deployHeight,
        uint64 attestInterval,
        uint64 blockPeriodNs,
        uint64[] shards,
        string name
    );

    /**
     * @notice A list of chain IDs that have registered OmniPortals.
     */
    uint64[] public chainIds;

    /**
     * @notice Portal deployments by chain ID.
     */
    mapping(uint64 => Deployment) public deployments;

    /**
     * @notice Deployment information for an OmniPortal.
     * @custom:field addr               The address of the deployment.
     * @custom:field chainId            The chain ID of the deployment.
     * @custom:field deployHeight       The height at which the deployment was deployed.
     * @custom:field attestInterval     The interval, in blocks, at which validators must attest, even if empty.
     * @custom:field blockPeriodNs      The block period of the chain deployed to, in nanoseconds.
     * @custom:field shards             Supported shards of the deployment.
     * @custom:field name               The name of the chain deployed to (ex "omni_evm", "ethereum")
     */
    struct Deployment {
        bytes32 addr;
        uint64 chainId;
        uint64 deployHeight;
        uint64 attestInterval;
        uint64 blockPeriodNs;
        uint64[] shards;
        string name;
    }

    constructor() {
        _disableInitializers();
    }

    function initialize(address owner_) public initializer {
        __Ownable_init(owner_);
    }

    /**
     * @notice Get the OmniPortal deployment for a chain.
     */
    function get(uint64 chainId) external view returns (Deployment memory) {
        return deployments[chainId];
    }

    /**
     * @notice List all registered OmniPortal deployments.
     */
    function list() external view returns (Deployment[] memory) {
        Deployment[] memory deps = new Deployment[](chainIds.length);
        for (uint64 i = 0; i < chainIds.length; i++) {
            deps[i] = deployments[chainIds[i]];
        }

        return deps;
    }

    /**
     * @notice Register a new OmniPortal deployment.
     */
    function register(Deployment calldata dep) external onlyOwner {
        _register(dep);
    }

    /**
     * @notice Register multiple OmniPortal deployments.
     */
    function bulkRegister(Deployment[] calldata deps) external onlyOwner {
        for (uint64 i = 0; i < deps.length; i++) {
            _register(deps[i]);
        }
    }

    /**
     * @notice Register an new OmniPortal deployment.
     * @dev Zero height deployments are allowed for now, as we use them for "private" chains.
     *      TODO: require non-zero height when e2e flow is updated to reflect real deploy heights.
     */
    function _register(Deployment calldata dep) internal {
        require(dep.addr != bytes32(0), "PortalRegistry: zero addr");
        require(dep.chainId > 0, "PortalRegistry: zero chain ID");
        require(dep.attestInterval > 0, "PortalRegistry: zero interval");
        require(dep.blockPeriodNs <= uint64(type(int64).max), "PortalRegistry: period too large");
        require(dep.blockPeriodNs > 0, "PortalRegistry: zero period");
        require(bytes(dep.name).length > 0, "PortalRegistry: no name");
        require(dep.shards.length > 0, "PortalRegistry: no shards");

        // TODO: allow multiple deployments per chain? maybe add a version reference?
        require(deployments[dep.chainId].addr == bytes32(0), "PortalRegistry: already set");

        // only allow ConfLevel shards
        for (uint64 i = 0; i < dep.shards.length; i++) {
            uint64 shard = dep.shards[i];
            require(shard == uint8(shard) && ConfLevel.isValid(uint8(shard)), "PortalRegistry: invalid shard");
        }

        deployments[dep.chainId] = dep;
        chainIds.push(dep.chainId);

        emit PortalRegistered(
            dep.chainId, dep.addr, dep.deployHeight, dep.attestInterval, dep.blockPeriodNs, dep.shards, dep.name
        );
    }
}
