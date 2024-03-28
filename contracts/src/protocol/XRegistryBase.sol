// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OmniPortal } from "../protocol/OmniPortal.sol";
import { ContractNames } from "../libraries/ContractNames.sol";
import { IXRegistry } from "../interfaces/IXRegistry.sol";

/**
 * @title XRegistryBase
 * @notice Logic and state shared between XRegistry and GlobalXRegistry contracts.
 */
abstract contract XRegistryBase is IXRegistry {
    /**
     * @notice The omni portal contract.
     */
    OmniPortal internal _portal;

    /**
     * @notice All chainIds with registered portals contracts
     */
    uint64[] internal _chainIds;

    /**
     *
     * @notice All deployed and registered contracts within the network
     * @dev Maps chainId => keccak256(name) => DeployInfo
     *      Name hashed to allow for contract identifiers that are not just string names (i.e.
     *      hashes namespaced by registrant)
     */
    mapping(uint64 => mapping(bytes32 => DeployInfo)) internal _contracts;

    /**
     * @notice Returns the DeployInfo of the contract with `name` on `chainId`.
     */
    function getContract(uint64 chainId, string memory name) public view override returns (DeployInfo memory) {
        return _contracts[chainId][_pack(name)];
    }

    /**
     * @notice Returns true if a deployment contract with `name` has been registered for `chainId`.
     */
    function isRegistered(uint64 chainId, string memory name) public view returns (bool) {
        return getContract(chainId, name).addr != address(0);
    }

    /**
     * @notice Returns the DeployInfo of the omni portal contract on `chainId`.
     */
    function getPortal(uint64 chainId) public view override returns (DeployInfo memory) {
        return getContract(chainId, ContractNames.OmniPortal);
    }

    /**
     * @notice Returns true if a portal contract has been registered for `chainId`.
     */
    function isPortalRegistered(uint64 chainId) public view override returns (bool) {
        return isRegistered(chainId, ContractNames.OmniPortal);
    }

    /**
     * @dev Hashes a string into a bytes32.
     */
    function _pack(string memory name) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(name));
    }

    function _mustGetContract(uint64 chainId, string memory name) internal view returns (DeployInfo memory) {
        DeployInfo memory info = getContract(chainId, name);
        require(info.addr != address(0), "XRegistry: not registered");
        return info;
    }

    /**
     * @dev Writes the contract deploy info to storage.
     */
    function _registerContract(uint64 chainId, string memory name, address addr, uint256 deployHeight) internal {
        _contracts[chainId][_pack(name)] = DeployInfo({ addr: addr, deployHeight: deployHeight });
    }

    /**
     * @dev Writes the portal contract deploy info to storage.
     */
    function _registerPortal(uint64 chainId, address addr, uint256 deployHeight) internal {
        _registerContract(chainId, ContractNames.OmniPortal, addr, deployHeight);
    }
}
