// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { ConfLevel } from "../libraries/ConfLevel.sol";

/**
 * @title PortalRegistry
 * @notice Registry for OmniPortal deployments. Predeployed on Omni's EVM.
 * @dev This contract is predeployed, and requires storage slots to be set in genesis.
 *      Genesis storage slots must:
 *          - set _owner on proxy
 *          - set _initialized on proxy to 1, to disable the initializer
 *          - set _initialized on implementation to 255, to disabled all initializers
 */
contract PortalRegistry is OwnableUpgradeable {
    /**
     * @notice Emitted when a new OmniPortal deployment is registered.
     */
    event PortalRegistered(uint64 indexed chainId, address indexed addr, uint64 deployHeight, uint64[] shards);

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
     * @custom:field chainId            The chain ID of the deployment.
     * @custom:field addr               The address of the deployment.
     * @custom:field deployHeight       The height at which the deployment was deployed.
     * @custom:field shards             Supported shards of the deployment.
     */
    struct Deployment {
        uint64 chainId;
        address addr;
        uint64 deployHeight;
        uint64[] shards;
    }

    /**
     * @notice Get the OmniPortal deployment for a chain.
     */
    function get(uint64 chainId) external view returns (Deployment memory) {
        return deployments[chainId];
    }

    /**
     * @notice Get the OmniPortal address for a chain.
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
    function bulkRegister(Deployment[] calldata deps) external payable onlyOwner {
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
        require(dep.addr != address(0), "PortalRegistry: no zero addr");
        require(dep.chainId > 0, "PortalRegistry: no zero chain ID");
        require(dep.shards.length > 0, "PortalRegistry: no shards");

        // TODO: allow multiple deployments per chain?
        require(deployments[dep.chainId].addr == address(0), "PortalRegistry: already set");

        // only allow ConfLevel shards
        for (uint64 i = 0; i < dep.shards.length; i++) {
            uint64 shard = dep.shards[i];
            require(shard == uint8(shard) && ConfLevel.isValid(uint8(shard)), "PortalRegistry: invalid shard");
        }

        deployments[dep.chainId] = dep;
        chainIds.push(dep.chainId);

        emit PortalRegistered(dep.chainId, dep.addr, dep.deployHeight, dep.shards);
    }
}
