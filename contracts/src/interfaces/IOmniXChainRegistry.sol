// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

interface IOmniXChainRegistry {
    event ChainAdded(uint64 indexed chainId, string name, address portal, uint256 deployHeight);

    struct Chain {
        uint64 chainId;
        string name; // not totally necessary, but could be useful
        address portal;
        uint256 deployHeight;
    }

    function isRegistered(uint64 chainId) external view returns (bool);

    function addChain(Chain calldata chain) external;

    function getChains() external view returns (Chain[] memory);

    function getChain(uint64 chainId) external view returns (Chain memory);
}
