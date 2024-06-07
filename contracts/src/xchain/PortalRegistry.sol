// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { XRegistry } from "./XRegistry.sol";
import { XRegistryBase } from "./XRegistryBase.sol";
import { Predeploys } from "../libraries/Predeploys.sol";
import { XRegistryNames } from "../libraries/XRegistryNames.sol";
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
     * @notice The XRegistry predeploy.
     */
    XRegistry public constant xregistry = XRegistry(Predeploys.XRegistry);

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
        uint64[] memory chainIds = xregistry.chainIds();

        // xregistry.chainIds() is set before portals are registered
        // we therefore need to filter out the zero addresses.

        uint256 nonZeros = 0;
        for (uint64 i = 0; i < chainIds.length; i++) {
            if (deployments[chainIds[i]].addr != address(0)) {
                nonZeros++;
            }
        }

        Deployment[] memory deps = new Deployment[](nonZeros);
        for (uint64 i = 0; i < chainIds.length; i++) {
            if (deployments[chainIds[i]].addr != address(0)) {
                deps[i] = deployments[chainIds[i]];
            }
        }

        return deps;
    }

    /**
     * @notice Register a new OmniPortal deployment.
     * @dev Zero height deployments are allowed for now, as we use them for "private" chains.
     *      TODO: require non-zero height when e2e flow is updated to reflect real deploy heights.
     */
    function register(Deployment calldata dep) external payable onlyOwner {
        require(!isRegistered(dep.chainId), "PortalRegistry: already set");
        require(dep.addr != address(0), "PortalRegistry: zero address");
        require(dep.shards.length > 0, "PortalRegistry: no shards");

        // only allow ConfLevel shards
        for (uint64 i = 0; i < dep.shards.length; i++) {
            uint64 shard = dep.shards[i];
            require(shard == uint8(shard) && ConfLevel.isValid(uint8(shard)), "PortalRegistry: invalid shard");
        }

        XRegistryBase.Deployment memory xdep =
            XRegistryBase.Deployment({ addr: dep.addr, metadata: abi.encode(dep.shards) });

        require(
            msg.value >= xregistry.registrationFee(dep.chainId, XRegistryNames.OmniPortal, xdep),
            "PortalRegistry: insufficient fee"
        );

        xregistry.register{ value: msg.value }(dep.chainId, XRegistryNames.OmniPortal, xdep);
        deployments[dep.chainId] = dep;

        emit PortalRegistered(dep.chainId, dep.addr, dep.deployHeight, dep.shards);
    }

    /**
     * @notice Calculate the fee to register a new OmniPortal deployment.
     */
    function registrationFee(Deployment calldata deployment) external view returns (uint256) {
        return xregistry.registrationFee(
            deployment.chainId,
            XRegistryNames.OmniPortal,
            XRegistryBase.Deployment({ addr: deployment.addr, metadata: abi.encode(deployment.shards) })
        );
    }

    /**
     * @notice Check if a deployment is registered.
     */
    function isRegistered(uint64 chainId) public view returns (bool) {
        return deployments[chainId].addr != address(0);
    }
}
