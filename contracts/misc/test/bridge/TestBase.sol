// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.26;

import { StablecoinUpgradeable } from "rlusd/contracts/StablecoinUpgradeable.sol";
import { StablecoinProxy } from "rlusd/contracts/StablecoinProxy.sol";

import { LockboxUpgradeable } from "src/bridge/LockboxUpgradeable.sol";
import { BridgeUpgradeable, IBridgeUpgradeable } from "src/bridge/BridgeUpgradeable.sol";
import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import { Test } from "forge-std/Test.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { MockPortal } from "core/test/utils/MockPortal.sol";

contract TestBase is Test {
    StablecoinUpgradeable internal originalToken;
    StablecoinUpgradeable internal srcWrapper;
    StablecoinUpgradeable internal destWrapper;

    LockboxUpgradeable internal srcLockbox;
    BridgeUpgradeable internal srcBridge;
    BridgeUpgradeable internal destBridge;

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
        BridgeUpgradeable origin,
        uint64 srcChainId,
        uint64 destChainId,
        bool wrap,
        address from,
        address to,
        uint256 value
    ) internal {
        address destination = origin.routes(destChainId);
        uint256 fee = origin.bridgeFee(destChainId);
        bytes memory data = abi.encodeCall(BridgeUpgradeable.receiveToken, (to, value));

        vm.chainId(srcChainId);
        vm.prank(from);
        vm.expectEmit(true, true, true, true);
        emit IBridgeUpgradeable.CrosschainTransfer(destChainId, from, to, value);
        origin.sendToken{ value: fee }(wrap, destChainId, to, value);

        vm.chainId(destChainId);
        vm.expectEmit(true, true, true, true);
        emit IBridgeUpgradeable.CrosschainReceive(srcChainId, to, value);
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
        srcBridge = _deployBridge(address(srcWrapper), address(originalToken), address(srcLockbox));
        destBridge = _deployBridge(address(destWrapper), address(0), address(0));
    }

    function _deployToken(string memory name, string memory symbol) internal returns (StablecoinUpgradeable) {
        address impl = address(new StablecoinUpgradeable());
        bytes memory data = abi.encodeCall(
            StablecoinUpgradeable.initialize, (name, symbol, minter, admin, upgrader, pauser, clawbacker)
        );

        address proxy = address(new StablecoinProxy(impl, data));
        return StablecoinUpgradeable(proxy);
    }

    function _deployLockbox(address token, address wrapper) internal returns (LockboxUpgradeable) {
        address impl = address(new LockboxUpgradeable());
        bytes memory data = abi.encodeCall(LockboxUpgradeable.initialize, (admin, pauser, token, wrapper));

        address proxy = address(new TransparentUpgradeableProxy(impl, admin, data));
        return LockboxUpgradeable(proxy);
    }

    function _deployBridge(address wrapper, address token, address lockbox) internal returns (BridgeUpgradeable) {
        bytes memory data =
            abi.encodeCall(BridgeUpgradeable.initialize, (admin, pauser, address(omni), wrapper, token, lockbox));
        address impl = address(new BridgeUpgradeable());

        address proxy = address(new TransparentUpgradeableProxy(impl, admin, data));
        return BridgeUpgradeable(proxy);
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
        srcBridge.configureBridges(chainIds, bridges);

        chainIds[0] = SRC_CHAIN_ID;
        bridges[0] = address(srcBridge);
        destBridge.configureBridges(chainIds, bridges);

        vm.stopPrank();
    }

    function _configurePermissions() internal {
        vm.startPrank(admin);

        srcWrapper.grantRole(srcWrapper.MINTER_ROLE(), address(srcLockbox));
        srcWrapper.grantRole(srcWrapper.BURNER_ROLE(), address(srcLockbox));

        srcWrapper.revokeRole(srcWrapper.MINTER_ROLE(), admin); // Assigned at initialization, unnecessary.
        srcWrapper.grantRole(srcWrapper.MINTER_ROLE(), address(srcBridge));
        srcWrapper.grantRole(srcWrapper.BURNER_ROLE(), address(srcBridge));

        destWrapper.revokeRole(destWrapper.MINTER_ROLE(), admin); // Assigned at initialization, unnecessary.
        destWrapper.grantRole(destWrapper.MINTER_ROLE(), address(destBridge));
        destWrapper.grantRole(destWrapper.BURNER_ROLE(), address(destBridge));

        vm.stopPrank();
    }
}
