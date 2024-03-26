// SPDX-License-Identifier: AGPL-3.0
pragma solidity ^0.8.12;

import { CREATE3 } from "solmate/src/utils/CREATE3.sol";

/**
 * @title Create3
 * @notice Factory for deploying contracts to deterministic addresses via CREATE3 Enables deploying
 *         contracts using CREATE3. Each deployer (msg.sender) has its own namespace for deployed
 *         addresses.
 * @author zefram.eth
 * @custom:attribution zefram.eth (https://github.com/ZeframLou/create3-factory/blob/main/src/CREATE3Factory.sol)
 */
contract Create3 {
    /// @notice Maps salt to deployed height.
    mapping(bytes32 => uint256) internal _deployedHeight;

    /**
     * @notice Deploys a contract using CREATE3
     * @dev The provided salt is hashed together with msg.sender to generate the final salt
     * @param salt          The deployer-specific salt for determining the deployed contract's address
     * @param creationCode  The creation code of the contract to deploy
     * @return deployed     The address of the deployed contract
     */
    function deploy(bytes32 salt, bytes memory creationCode) external payable returns (address deployed) {
        // hash salt with the deployer address to give each deployer its own namespace
        salt = keccak256(abi.encodePacked(msg.sender, salt));
        deployed = CREATE3.deploy(salt, creationCode, msg.value);
        _deployedHeight[salt] = block.number;
        return deployed;
    }

    /**
     * @notice Predicts the address of a deployed contract
     * @dev The provided salt is hashed together with the deployer address to generate the final salt
     * @param deployer  The deployer account that will call deploy()
     * @param salt      The deployer-specific salt for determining the deployed contract's address
     * @return deployed The address of the contract that will be deployed
     */
    function getDeployedAddr(address deployer, bytes32 salt) external view returns (address deployed) {
        // hash salt with the deployer address to give each deployer its own namespace
        salt = keccak256(abi.encodePacked(deployer, salt));
        return CREATE3.getDeployed(salt);
    }

    /**
     * @notice Returns the block height at which the salt was used.
     * @dev The provided salt is hashed together with the deployer address to generate the final salt
     * @param deployer  The deployer account that will call deploy()
     * @param salt      The deployer-specific salt for determining the deployed contract's address
     * @return height The height at which the salt was used
     */
    function getDeployedHeight(address deployer, bytes32 salt) external view returns (uint256) {
        salt = keccak256(abi.encodePacked(deployer, salt));
        return _deployedHeight[salt];
    }
}
