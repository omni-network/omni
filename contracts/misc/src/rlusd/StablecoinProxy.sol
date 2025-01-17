// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import { ERC1967Proxy } from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

/**
 * @dev Proxy to interact with {StablecoinUpgradeable}, {StablecoinLockboxUpgradeable} and {StablecoinBridge}.
 *
 */
contract StablecoinProxy is ERC1967Proxy {
    constructor(address _delegate, bytes memory _data) ERC1967Proxy(_delegate, _data) { }

    /**
     * @dev Returns the implementation address of the contract that executes a transaction.
     */
    function getImplementation() public view returns (address) {
        return _implementation();
    }
}
