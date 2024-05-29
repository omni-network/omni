// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { XRegistryNames } from "../libraries/XRegistryNames.sol";
import { Predeploys } from "../libraries/Predeploys.sol";

/**
 * @title XRegistryBase
 * @notice Base contract for XRegistry and XRegistryReplica. Contains common storage, setters, and views.
 */
contract XRegistryBase {
    /**
     * @dev Mapping of chain ID -> keccak256(name, registrant) -> address.
     */
    mapping(uint64 => mapping(bytes32 => address)) internal _addrs;

    /**
     * @notice Return true if a contract with `name` has been registered by `registrant` for `chainId`.
     * @param chainId       The chain ID of the registration.
     * @param name          The name of the contract.
     * @param registrant    The address of the registrant.
     */
    function has(uint64 chainId, string memory name, address registrant) external view returns (bool) {
        return _get(chainId, name, registrant) != address(0);
    }

    /**
     * @notice Return the address of the contract with `name` registered by `registrant` for `chainId`.
     * @param chainId       The chain ID of the registration.
     * @param name          The name of the contract.
     * @param registrant    The address of the registrant.
     */
    function get(uint64 chainId, string memory name, address registrant) external view returns (address) {
        return _get(chainId, name, registrant);
    }

    function _get(uint64 chainId, string memory name, address registrant) internal view returns (address) {
        return _addrs[chainId][_pack(name, registrant)];
    }

    function _set(uint64 chainId, string memory name, address registrant, address addr) internal {
        _addrs[chainId][_pack(name, registrant)] = addr;
    }

    function _pack(string memory name, address registrant) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(name, registrant));
    }

    /**
     * @dev Return true if `name` and `registrant` are "OmniPortal" and PortalRegistry predpeloy, respectively.
     *      This marks a portal registration internal to Omni's protocol.
     */
    function _isPortal(string memory name, address registrant) internal pure returns (bool) {
        return _pack(name, registrant) == _pack(XRegistryNames.OmniPortal, Predeploys.PortalRegistry);
    }
}
