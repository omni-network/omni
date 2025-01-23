// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.26;

import { StablecoinUpgradeable } from "rlusd/contracts/StablecoinUpgradeable.sol";
import { StablecoinProxy } from "rlusd/contracts/StablecoinProxy.sol";

import { Lockbox } from "src/bridge/Lockbox.sol";
import { Bridge, IBridge } from "src/bridge/Bridge.sol";
import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import { Test } from "forge-std/Test.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { MockPortal } from "core/test/utils/MockPortal.sol";

contract TestBase is Test {
    StablecoinUpgradeable internal originalToken;
    StablecoinUpgradeable internal srcWrapper;
    StablecoinUpgradeable internal destWrapper;

    Lockbox internal srcLockbox;
    Bridge internal srcBridge;
    Bridge internal destBridge;

    MockPortal internal omni;

    uint64 internal constant SRC_CHAIN_ID = 1;
    uint64 internal constant DEST_CHAIN_ID = 2;
    uint64 internal constant DEFAULT_GAS_LIMIT = 140_000;
    uint256 internal constant INITIAL_USER_BALANCE = 1_000_000 ether;

    address internal user = makeAddr("user");
    address internal admin = makeAddr("admin");
    address internal minter = makeAddr("minter");
    address internal pauser = makeAddr("pauser");
    address internal upgrader = makeAddr("upgrader");
    address internal clawbacker = makeAddr("clawbacker");

    function setUp() public {
        deploy();
        configure();
        vm.chainId(SRC_CHAIN_ID);
    }

    function deploy() internal {
        omni = new MockPortal();
        _deployTokens();
        _deployInfra();
    }

    function configure() internal {
        _fundUser();
        _configureApprovals();
        _configureRoutes();
        _configurePermissions();
    }

    function mockBridge(
        Bridge origin,
        uint64 srcChainId,
        uint64 destChainId,
        bool wrap,
        address from,
        address to,
        uint256 value
    ) internal {
        address destination = origin.routes(destChainId);
        uint256 fee = origin.bridgeFee(destChainId);
        bytes memory data = abi.encodeCall(Bridge.receiveToken, (to, value));

        vm.chainId(srcChainId);
        vm.prank(from);
        vm.expectEmit(true, true, true, true);
        emit IBridge.TokenSent(destChainId, from, to, value);
        origin.sendToken{ value: fee }(destChainId, to, value, wrap);

        vm.chainId(destChainId);
        vm.expectEmit(true, true, true, true);
        emit IBridge.TokenReceived(srcChainId, to, value);
        omni.mockXCall(srcChainId, address(origin), destination, data, DEFAULT_GAS_LIMIT);

        vm.chainId(srcChainId);
    }

    function _deployTokens() internal {
        originalToken = _deployToken("Ripple USD", "RLUSD");
        srcWrapper = _deployToken("Bridged RLUSD (Omni)", "RLUSD.e");
        destWrapper = _deployToken("Bridged RLUSD (Omni)", "RLUSD.e");
    }

    function _deployInfra() internal {
        srcLockbox = _deployLockbox(address(originalToken), address(srcWrapper));
        srcBridge = _deployBridge(address(srcWrapper), address(srcLockbox));
        destBridge = _deployBridge(address(destWrapper), address(0));
    }

    function _deployToken(string memory name, string memory symbol) internal returns (StablecoinUpgradeable) {
        address impl = address(new StablecoinUpgradeable());
        bytes memory data = abi.encodeCall(
            StablecoinUpgradeable.initialize, (name, symbol, minter, admin, upgrader, pauser, clawbacker)
        );

        address proxy = address(new StablecoinProxy(impl, data));
        return StablecoinUpgradeable(proxy);
    }

    function _deployLockbox(address token, address wrapper) internal returns (Lockbox) {
        address impl = address(new Lockbox());
        bytes memory data = abi.encodeCall(Lockbox.initialize, (admin, pauser, token, wrapper));

        address proxy = address(new TransparentUpgradeableProxy(impl, admin, data));
        return Lockbox(proxy);
    }

    function _deployBridge(address token, address lockbox) internal returns (Bridge) {
        bytes memory data = abi.encodeCall(Bridge.initialize, (admin, pauser, address(omni), token, lockbox));
        address impl = address(new Bridge());

        address proxy = address(new TransparentUpgradeableProxy(impl, admin, data));
        return Bridge(proxy);
    }

    function _fundUser() internal {
        vm.deal(user, 1 ether);
        vm.prank(minter);
        originalToken.mint(user, INITIAL_USER_BALANCE);
    }

    function _configureApprovals() internal {
        vm.startPrank(user);

        // Approve source lockbox to wrap original tokens.
        originalToken.approve(address(srcLockbox), type(uint256).max);

        // Approve source bridge to transfer original tokens.
        originalToken.approve(address(srcBridge), type(uint256).max);

        // Approve both bridges to transfer wrapped tokens.
        srcWrapper.approve(address(srcBridge), type(uint256).max);
        destWrapper.approve(address(destBridge), type(uint256).max);

        vm.stopPrank();
    }

    function _configureRoutes() internal {
        uint64[] memory chainIds = new uint64[](1);
        address[] memory bridges = new address[](1);

        vm.startPrank(admin);

        chainIds[0] = DEST_CHAIN_ID;
        bridges[0] = address(destBridge);
        srcBridge.setRoutes(chainIds, bridges);

        chainIds[0] = SRC_CHAIN_ID;
        bridges[0] = address(srcBridge);
        destBridge.setRoutes(chainIds, bridges);

        vm.stopPrank();
    }

    function _configurePermissions() internal {
        vm.startPrank(admin);

        srcWrapper.grantRole(srcWrapper.MINTER_ROLE(), address(srcLockbox));
        srcWrapper.grantRole(srcWrapper.CLAWBACKER_ROLE(), address(srcLockbox));

        srcWrapper.revokeRole(srcWrapper.MINTER_ROLE(), admin); // Assigned at initialization, unnecessary.
        srcWrapper.grantRole(srcWrapper.MINTER_ROLE(), address(srcBridge));
        srcWrapper.grantRole(srcWrapper.BURNER_ROLE(), address(srcBridge));

        destWrapper.revokeRole(destWrapper.MINTER_ROLE(), admin); // Assigned at initialization, unnecessary.
        destWrapper.grantRole(destWrapper.MINTER_ROLE(), address(destBridge));
        destWrapper.grantRole(destWrapper.BURNER_ROLE(), address(destBridge));

        vm.stopPrank();
    }
}
