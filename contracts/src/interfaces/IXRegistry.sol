// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

interface IXRegistry {
    event ContractAdded(uint64 indexed chainId, string indexed name, address addr, uint256 deployHeight);

    /**
     * @notice Contract deployment information for a chain within the Omni network.
     * @custom:field addr           The address of the deployed contract
     * @custom:field deployHeight   The block height at which the contract was deployed
     */
    struct DeployInfo {
        address addr;
        uint256 deployHeight;
    }

    function addContract(uint64 chainId, string calldata name, address addr, uint256 deployHeight) external;

    function getContract(uint64 chainId, string calldata name) external view returns (DeployInfo memory);

    function isRegistered(uint64 chainId, string calldata name) external view returns (bool);

    function addPortal(uint64 chainId, address addr, uint256 deployHeight) external;

    function getPortal(uint64 chainId) external view returns (DeployInfo memory);

    function isPortalRegistered(uint64 chainId) external view returns (bool);
}
