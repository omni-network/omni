// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.30;

/**
 * @title Predeploys
 * @notice Halo predeploy addresses (match halo/genutil/evm/predeploys.go)
 */
library Predeploys {
    uint256 internal constant NamespaceSize = 1024;
    address internal constant NominaNamespace = 0x121e240000000000000000000000000000000000;
    address internal constant OctaneNamespace = 0xCCcCCc0000000000000000000000000000000000;

    // nomina predeploys
    address internal constant PortalRegistry = 0x121E240000000000000000000000000000000001;
    address internal constant NominaBridgeNative = 0x121E240000000000000000000000000000000002;
    address internal constant WNomina = 0x121E240000000000000000000000000000000003;

    // octane predeploys
    address internal constant Staking = 0xCCcCcC0000000000000000000000000000000001;
    address internal constant Slashing = 0xCccCCC0000000000000000000000000000000002;
    address internal constant Upgrade = 0xccCCcc0000000000000000000000000000000003;
    address internal constant Distribution = 0xCcCcCC0000000000000000000000000000000004;

    function namespaces() internal pure returns (address[] memory ns) {
        ns = new address[](2);
        ns[0] = NominaNamespace;
        ns[1] = OctaneNamespace;
    }

    /**
     * @notice Return true if `addr` is not proxied
     */
    function notProxied(address addr) internal pure returns (bool) {
        return addr == WNomina;
    }

    /**
     * @notice Return implementation address for a proxied predeploy
     */
    function impl(address addr) internal pure returns (address) {
        require(isPredeploy(addr), "Predeploys: not a predeploy");
        require(!notProxied(addr), "Predeploys: not proxied");

        // max uint160 is odd, which gives us unique implementation for each predeploy
        return address(type(uint160).max - uint160(addr));
    }

    /**
     * @notice Return true if `addr` is an active predeploy
     */
    function isActivePredeploy(address addr) internal pure returns (bool) {
        return addr == PortalRegistry || addr == NominaBridgeNative || addr == WNomina || addr == Staking
            || addr == Slashing || addr == Upgrade || addr == Distribution;
    }

    /**
     * @notice Return true if `addr` is in some predeploy namespace
     */
    function isPredeploy(address addr) internal pure returns (bool) {
        return (uint160(addr) >> 10 == uint160(NominaNamespace) >> 10)
            || (uint160(addr) >> 10 == uint160(OctaneNamespace) >> 10);
    }
}
