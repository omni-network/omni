// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";
import { XRegistry } from "./XRegistry.sol";
import { Predeploys } from "../libraries/Predeploys.sol";
import { XRegistryNames } from "../libraries/XRegistryNames.sol";

/**
 * @title PortalRegistry
 * @notice Registry for OmniPortal deployments. Predeployed on Omni's EVM.
 * @dev Using Ownable, rather than OwnableUpgradeable, because predeploys are initialized at genesis.
 */
contract PortalRegistry is Ownable {
    /**
     * @notice Emitted when a new OmniPortal deployment is registered.
     */
    event PortalRegistered(uint64 indexed chainId, address indexed addr, uint64 deployHeight, string finalizationStrat);

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
     * @custom:field finalizationStrat  The finalization strategy of the chain (e.g. "finalized")
     */
    struct Deployment {
        uint64 chainId;
        address addr;
        uint64 deployHeight;
        string finalizationStrat;
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

        Deployment[] memory deps = new Deployment[](chainIds.length);
        for (uint64 i = 0; i < chainIds.length; i++) {
            deps[i] = deployments[chainIds[i]];
        }

        return deps;
    }

    /**
     * @notice Register a new OmniPortal deployment.
     * @dev Zero height deployments are allowed for now, as we use them for "private" chains.
     *      TODO: require non-zero height when e2e flow is updated to reflect real deploy heights.
     */
    function register(Deployment calldata deployment) external payable onlyOwner {
        require(!isRegistered(deployment.chainId), "PortalRegistry: already set");
        require(deployment.addr != address(0), "PortalRegistry: zero address");
        require(bytes(deployment.finalizationStrat).length != 0, "PortalRegistry: empty strat");
        require(
            msg.value >= xregistry.registrationFee(deployment.chainId, XRegistryNames.OmniPortal, deployment.addr),
            "PortalRegistry: insufficient fee"
        );

        xregistry.register{ value: msg.value }(deployment.chainId, XRegistryNames.OmniPortal, deployment.addr);
        deployments[deployment.chainId] = deployment;

        emit PortalRegistered(
            deployment.chainId, deployment.addr, deployment.deployHeight, deployment.finalizationStrat
        );
    }

    /**
     * @notice Calculate the fee to register a new OmniPortal deployment.
     */
    function registrationFee(Deployment calldata deployment) external view returns (uint256) {
        return xregistry.registrationFee(deployment.chainId, XRegistryNames.OmniPortal, deployment.addr);
    }

    /**
     * @notice Check if a deployment is registered.
     */
    function isRegistered(uint64 chainId) public view returns (bool) {
        return deployments[chainId].addr != address(0);
    }
}
