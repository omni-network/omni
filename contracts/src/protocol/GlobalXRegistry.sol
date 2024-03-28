// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

import { OmniPortal } from "./OmniPortal.sol";
import { XRegistryBase } from "./XRegistryBase.sol";
import { ContractNames } from "../libraries/ContractNames.sol";
import { IXRegistry } from "../interfaces/IXRegistry.sol";

/**
 * @title GlobalXRegistry
 * @notice Global registry of all deployed contracts on all chains. Controls local XRegistries at
 *         each portal. It functions as the local XRegistry for portal deployed on Omni.
 */
contract GlobalXRegistry is IXRegistry, OwnableUpgradeable, XRegistryBase {
    uint64 internal constant XADD_CONTRACT_GAS_LIMIT = 200_000;

    function initialize(address owner) external initializer {
        _transferOwnership(owner);
        __Ownable_init();
    }

    function syncXRegistry(uint64 destChainId, string calldata contractName) public {
        for (uint256 i = 0; i < _chainIds.length; i++) {
            uint64 chainId = _chainIds[i];

            // skip this chain
            if (chainId == _portal.chainId()) continue;

            address xregistry = _mustGetContract(chainId, ContractNames.XRegistry).addr;
        }
    }


    function addContract(uint64 chainId, string calldata name, address addr, uint256 deployHeight) external onlyOwner {
        _addContract(chainId, name, addr, deployHeight);
    }

    function addXRegistry(uint64 chainId, address addr, uint256 deployHeight) external onlyOwner {
        _addContract(chainId, ContractNames.XRegistry, addr, deployHeight);
    }

    function addPortal(uint64 chainId, address addr, uint256 deployHeight) external onlyOwner {
        _addContract(chainId, ContractNames.OmniPortal, addr, deployHeight);
    }

    function _addContract(uint64 chainId, string calldata name, address addr, uint256 deployHeight) internal {
        require(address(_portal) != address(0), "XRegistry: portal not set");
        require(!isRegistered(chainId, name), "XRegistry: already registered"); // allow overwriting?
        require(addr != address(0), "XRegistry: portal is zero addr");
        require(deployHeight > 0, "XRegistry: deployHeight is zero");

        // register contract locally
        _registerContract(chainId, name, addr, deployHeight);

        // add contract on all registered portals
        _xAddContract(chainId, name, addr, deployHeight);

        emit ContractAdded(chainId, name, addr, deployHeight);
    }

    // add new chain at all currently registered portals
    function _xAddContract(uint64 chainId, string calldata name, address addr, uint256 deployHeight) internal {
        for (uint256 i = 0; i < _chainIds.length; i++) {
            uint64 destChainId = _chainIds[i];

            // skip this chain
            if (destChainId == _portal.chainId()) continue;

            address xregistry = _mustGetContract(destChainId, ContractNames.XRegistry).addr;

            _portal.xcall(
                destChainId,
                xregistry,
                abi.encodeWithSelector(_portal.addContract.selector, chainId, name, addr, deployHeight),
                XADD_CONTRACT_GAS_LIMIT
            );
        }
    }

    function _xAddContract
}
