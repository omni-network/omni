// SPDX-License-Identifier: AGPL-3.0
pragma solidity ^0.8.12;

import { ProxyAdmin as OZProxyAdmin } from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

/**
 * @title ProxyAdmin
 * @notice Wrapper around OpenZeppelin's ProxyAdmin that allows for setting the owner on deployment.
 * @dev This allows us to use Create3 to deploy
 */
contract ProxyAdmin is OZProxyAdmin {
    constructor(address owner) OZProxyAdmin() {
        _transferOwnership(owner);
    }
}
