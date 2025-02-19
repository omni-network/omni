// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.26;

import "../TestBase.sol";

contract GeneralBridgeTest is TestBase {
    function test_initialize_reverts() public {
        address impl = address(new Bridge(DEFAULT_RECEIVE_GAS_LIMIT, DEFAULT_RECEIVE_LOCKBOX_GAS_LIMIT));
        bridgeWithLockbox = Bridge(address(new TransparentUpgradeableProxy(impl, admin, "")));

        // `admin` cannot be zero address.
        vm.expectRevert(IBridge.ZeroAddress.selector);
        bridgeWithLockbox.initialize({
            admin_: address(0),
            authorizer_: address(0),
            pauser_: address(0),
            unpauser_: address(0),
            omni_: address(0),
            token_: address(0),
            lockbox_: address(0)
        });

        // `authorizer` cannot be zero address.
        vm.expectRevert(IBridge.ZeroAddress.selector);
        bridgeWithLockbox.initialize({
            admin_: admin,
            authorizer_: address(0),
            pauser_: address(0),
            unpauser_: address(0),
            omni_: address(0),
            token_: address(0),
            lockbox_: address(0)
        });

        // `pauser` cannot be zero address.
        vm.expectRevert(IBridge.ZeroAddress.selector);
        bridgeWithLockbox.initialize({
            admin_: admin,
            authorizer_: authorizer,
            pauser_: address(0),
            unpauser_: address(0),
            omni_: address(0),
            token_: address(0),
            lockbox_: address(0)
        });

        // `unpauser` cannot be zero address.
        vm.expectRevert(IBridge.ZeroAddress.selector);
        bridgeWithLockbox.initialize({
            admin_: admin,
            authorizer_: authorizer,
            pauser_: pauser,
            unpauser_: address(0),
            omni_: address(0),
            token_: address(0),
            lockbox_: address(0)
        });

        // `omni` cannot be zero address.
        vm.expectRevert(IBridge.ZeroAddress.selector);
        bridgeWithLockbox.initialize({
            admin_: admin,
            authorizer_: authorizer,
            pauser_: pauser,
            unpauser_: unpauser,
            omni_: address(0),
            token_: address(0),
            lockbox_: address(0)
        });

        // `token` cannot be zero address.
        vm.expectRevert(IBridge.ZeroAddress.selector);
        bridgeWithLockbox.initialize({
            admin_: admin,
            authorizer_: authorizer,
            pauser_: pauser,
            unpauser_: unpauser,
            omni_: address(omni),
            token_: address(0),
            lockbox_: address(0)
        });

        // Initialization works with all fields populated.
        bridgeWithLockbox.initialize({
            admin_: admin,
            authorizer_: authorizer,
            pauser_: pauser,
            unpauser_: unpauser,
            omni_: address(omni),
            token_: address(token),
            lockbox_: address(lockbox)
        });

        // Initialization can also occur without a `lockbox` if none is present.
        bridgeNoLockbox = Bridge(address(new TransparentUpgradeableProxy(impl, admin, "")));
        bridgeNoLockbox.initialize({
            admin_: admin,
            authorizer_: authorizer,
            pauser_: pauser,
            unpauser_: unpauser,
            omni_: address(omni),
            token_: address(token),
            lockbox_: address(0)
        });
    }

    function test_setRoutes_reverts() public prank(admin) {
        uint64[] memory chainIds = new uint64[](1);
        chainIds[0] = DEST_CHAIN_ID + 1;
        IBridge.Route[] memory routes;

        // reverts if array lengths mismatch
        vm.expectRevert(IBridge.ArrayLengthMismatch.selector);
        bridgeWithLockbox.configureRoutes(chainIds, routes);
        routes = new IBridge.Route[](1);
        routes[0] = IBridge.Route({ bridge: makeAddr("newBridge"), hasLockbox: false });

        // Configuration successful with valid inputs.
        bridgeWithLockbox.configureRoutes(chainIds, routes);
    }

    function test_setRoutes_succeeds() public prank(admin) {
        uint64[] memory chainIds = new uint64[](1);
        chainIds[0] = DEST_CHAIN_ID + 1;
        IBridge.Route[] memory routes = new IBridge.Route[](1);
        routes[0] = IBridge.Route({ bridge: makeAddr("newBridge"), hasLockbox: false });

        bridgeWithLockbox.configureRoutes(chainIds, routes);
    }

    function test_authorizeRoutes_succeeds() public {
        test_setRoutes_succeeds();

        uint64[] memory chainIds = new uint64[](1);
        chainIds[0] = DEST_CHAIN_ID + 1;
        vm.prank(authorizer);
        bridgeWithLockbox.authorizeRoutes(chainIds);
    }
}
