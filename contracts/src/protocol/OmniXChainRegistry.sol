// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

import { OmniPortal } from "../protocol/OmniPortal.sol";
import { EnumerableUint64Set } from "../libraries/EnumerableUint64Set.sol";
import { IOmniXChainRegistry } from "../interfaces/IOmniXChainRegistry.sol";

contract OmniXChainRegistry is
    IOmniXChainRegistry,
    OwnableUpgradeable // use OwnableUpgradeable when proxies are used for predeploys
{
    using EnumerableUint64Set for EnumerableUint64Set.Set;

    // TODO: make configurable
    uint64 internal constant XADD_CHAIN_GAS_LIMIT = 200_000;
    uint64 internal constant XREMOVE_CHAIN_GAS_LIMIT = 200_000;

    OmniPortal internal _portal;

    // enumerable set of chainIds
    EnumerableUint64Set.Set internal _chainIds;

    // mapping of chainId to Chain struct
    mapping(uint64 => Chain) internal _chains;

    modifier whenPortalSet() {
        require(address(_portal) != address(0), "XRegistry: portal not set");
        _;
    }

    // Since this is predeployed, this will need to be called before genesis on a mock backend, and
    // storage needs to be copied to genesis
    //
    // For now, we call it right after geneis. We'll solve this when doing proxy predeploy namespace
    // task.
    function initialize(address owner) external initializer {
        _transferOwnership(owner);
        __Ownable_init();
    }

    // We need to deploy all of the portals, included one on the OmniEVM. So we will not know the
    // portal addresses and deploy heights until after genesis. So they cannot be in the
    // initializer. We also want to make sure only the owner can set portal / register chains.
    // So we do this post genesis

    // initPortal:
    //      - set the portal address
    //      - add the "omni" chain
    function initPortal(address portal, uint256 deployHeight) public onlyOwner {
        require(portal != address(0), "XRegistry: portal is zero addr");
        require(address(_portal) == address(0), "XRegistry: portal already set");
        require(deployHeight > 0, "XRegistry: deployHeight is zero");
        require(deployHeight < block.number, "XRegistry: future deployHeight");

        // set portal first, so _addChain can use it
        _portal = OmniPortal(portal);

        // add "omni" chain
        Chain memory omni = Chain(uint64(block.chainid), "omni", portal, deployHeight);

        // add it here
        _chainIds.add(omni.chainId);
        _chains[omni.chainId] = omni;

        // add it at the poral
        _portal.addChain(omni.chainId);
    }

    function getChains() external view returns (Chain[] memory) {
        Chain[] memory chains_ = new Chain[](_chainIds.length());
        for (uint256 i = 0; i < _chainIds.length(); i++) {
            chains_[i] = _chains[_chainIds.at(i)];
        }
        return chains_;
    }

    function getChain(uint64 chainId) external view returns (Chain memory) {
        return _chains[chainId];
    }

    function isRegistered(uint64 chainId) public view returns (bool) {
        return _chainIds.contains(chainId);
    }

    function addChain(Chain calldata c) external onlyOwner whenPortalSet {
        require(c.portal != address(0), "XRegistry: portal is zero addr");
        require(c.deployHeight > 0, "XRegistry: deployHeight is zero");

        // for existing chains, add new chain
        _xAddChain(c.chainId);

        // add to portal on this chain
        _portal.addChain(c.chainId);

        // for new chain, add existing chains
        _xInitChains(c.chainId, _chainIds.values());

        // add chain to registry
        require(_chainIds.add(c.chainId), "XRegistry: chain already exists");
        _chains[c.chainId] = c;

        emit ChainAdded(c.chainId, c.name, c.portal, c.deployHeight);
    }

    // add new chain at all currently registered portals
    function _xAddChain(uint64 chainId) internal {
        for (uint256 i = 0; i < _chainIds.length(); i++) {
            uint64 destChainId = _chainIds.at(i);

            // skip omni chain
            if (destChainId == _portal.chainId()) continue;

            _portal.xcall(
                destChainId,
                _portal.VIRTUAL_PORTAL_ADDRESS(),
                abi.encodeWithSelector(_portal.addChain.selector, chainId),
                XADD_CHAIN_GAS_LIMIT
            );
        }
    }

    // init chains at new chain
    function _xInitChains(uint64 portalChainId, uint64[] memory chainIds) internal {
        _portal.xcall(
            portalChainId,
            _portal.VIRTUAL_PORTAL_ADDRESS(),
            abi.encodeWithSelector(_portal.initChains.selector, chainIds),
            XADD_CHAIN_GAS_LIMIT
        );
    }
}
