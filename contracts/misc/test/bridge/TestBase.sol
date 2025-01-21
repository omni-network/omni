// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.26;

import { StablecoinUpgradeable } from "RLUSD-Implementation/contracts/StablecoinUpgradeable.sol";
import { LockboxUpgradeable } from "src/bridge/LockboxUpgradeable.sol";
import { BridgeUpgradeable } from "src/bridge/BridgeUpgradeable.sol";
import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { Proxy } from "src/bridge/Proxy.sol";

import { ILockboxUpgradeable } from "src/bridge/interfaces/ILockboxUpgradeable.sol";
import { IBridgeUpgradeable } from "src/bridge/interfaces/IBridgeUpgradeable.sol";

import { Test } from "forge-std/Test.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { MockPortal } from "core/test/utils/MockPortal.sol";

contract TestBase is Test {
    LockboxUpgradeable internal lockboxImpl;
    BridgeUpgradeable internal bridgeImpl;
    StablecoinUpgradeable internal tokenImpl;

    ILockboxUpgradeable internal lockboxSrc;
    ILockboxUpgradeable internal lockboxA;
    ILockboxUpgradeable internal lockboxB;

    IBridgeUpgradeable internal bridgeSrc;
    IBridgeUpgradeable internal bridgeA;
    IBridgeUpgradeable internal bridgeB;

    StablecoinUpgradeable internal tokenSrc;
    StablecoinUpgradeable internal tokenA;
    StablecoinUpgradeable internal tokenB;

    MockPortal internal omni;

    uint64 internal constant DEFAULT_GAS_LIMIT = 105_000;
    uint256 internal constant INITIAL_USER_BALANCE = 1_000_000 ether;

    uint64 internal constant srcChainId = 1;
    uint64 internal constant destChainIdA = 2;
    uint64 internal constant destChainIdB = 3;

    address internal user = makeAddr("user");
    address internal solver = makeAddr("solver");
    address internal admin = makeAddr("admin");
    address internal upgrader = makeAddr("upgrader");
    address internal pauser = makeAddr("pauser");
    address internal clawbacker = makeAddr("clawbacker");

    modifier prankUser(address addr) {
        vm.startPrank(addr);
        _;
        vm.stopPrank();
    }

    function setUp() public {
        deploy();
        configureBridges();

        vm.deal(user, 1 ether);
        vm.deal(solver, 1 ether);

        vm.prank(admin);
        tokenSrc.mint(user, INITIAL_USER_BALANCE);
    }

    function deploy() internal {
        omni = new MockPortal();

        lockboxImpl = new LockboxUpgradeable();
        bridgeImpl = new BridgeUpgradeable();
        tokenImpl = new StablecoinUpgradeable();

        lockboxSrc = _deployLockbox();
        lockboxA = _deployLockbox();
        lockboxB = _deployLockbox();

        bridgeSrc = _deployBridge(address(lockboxSrc));
        bridgeA = _deployBridge(address(lockboxA));
        bridgeB = _deployBridge(address(lockboxB));

        tokenSrc = _deployToken("Ripple USD", "RLUSD", admin);
        tokenA = _deployToken("Bridged RLUSD (Omni)", "RLUSD.e", address(bridgeA));
        tokenB = _deployToken("Bridged RLUSD (Omni)", "RLUSD.e", address(bridgeB));

        vm.startPrank(admin);
        tokenA.grantRole(tokenA.BURNER_ROLE(), address(bridgeA));
        tokenB.grantRole(tokenB.BURNER_ROLE(), address(bridgeB));
        vm.stopPrank();
    }

    function configureBridges() internal prankUser(admin) {
        uint64[] memory destChainIds = new uint64[](2);
        address[] memory destTokens = new address[](2);
        address[] memory destBridges = new address[](2);

        destChainIds[0] = destChainIdA;
        destTokens[0] = address(tokenA);
        destBridges[0] = address(bridgeA);

        destChainIds[1] = destChainIdB;
        destTokens[1] = address(tokenB);
        destBridges[1] = address(bridgeB);

        _configureBridge(bridgeSrc, tokenSrc, destChainIds, destTokens, destBridges);

        destChainIds[0] = destChainIdB;
        destTokens[0] = address(tokenB);
        destBridges[0] = address(bridgeB);

        destChainIds[1] = srcChainId;
        destTokens[1] = address(tokenSrc);
        destBridges[1] = address(bridgeSrc);

        _configureBridge(bridgeA, tokenA, destChainIds, destTokens, destBridges);

        destChainIds[0] = srcChainId;
        destTokens[0] = address(tokenSrc);
        destBridges[0] = address(bridgeSrc);

        destChainIds[1] = destChainIdA;
        destTokens[1] = address(tokenA);
        destBridges[1] = address(bridgeA);

        _configureBridge(bridgeB, tokenB, destChainIds, destTokens, destBridges);
    }

    function mockBridge(
        IBridgeUpgradeable srcBridge,
        uint64 sourceChainId,
        uint64 destChainId,
        address srcToken,
        address to,
        uint256 value
    ) internal {
        address destBridge = srcBridge.bridgeRoutes(destChainId);
        address destToken = srcBridge.tokenRoutes(srcToken, destChainId);
        uint256 fee = srcBridge.bridgeFee(destChainId);
        bytes memory data = abi.encodeCall(BridgeUpgradeable.receiveToken, (destToken, to, value));

        vm.chainId(sourceChainId);
        srcBridge.sendToken{ value: fee }(destChainId, srcToken, to, value);

        vm.chainId(destChainId);
        omni.mockXCall(sourceChainId, address(srcBridge), destBridge, data, DEFAULT_GAS_LIMIT);
    }

    function _deployLockbox() internal returns (ILockboxUpgradeable) {
        return ILockboxUpgradeable(
            address(
                new Proxy(
                    address(lockboxImpl), abi.encodeCall(LockboxUpgradeable.initialize, (admin, upgrader, pauser))
                )
            )
        );
    }

    function _deployBridge(address lockbox) internal returns (IBridgeUpgradeable) {
        return IBridgeUpgradeable(
            address(
                new Proxy(
                    address(bridgeImpl),
                    abi.encodeCall(
                        BridgeUpgradeable.initialize, (address(omni), address(lockbox), admin, upgrader, pauser)
                    )
                )
            )
        );
    }

    function _deployToken(string memory name, string memory symbol, address mintAuthority)
        internal
        returns (StablecoinUpgradeable)
    {
        return StablecoinUpgradeable(
            address(
                new Proxy(
                    address(tokenImpl),
                    abi.encodeCall(
                        StablecoinUpgradeable.initialize,
                        (name, symbol, mintAuthority, admin, upgrader, pauser, clawbacker)
                    )
                )
            )
        );
    }

    function _configureBridge(
        IBridgeUpgradeable bridge,
        StablecoinUpgradeable token,
        uint64[] memory destChainIds,
        address[] memory destTokens,
        address[] memory destBridges
    ) internal {
        bool isNative = bridge == bridgeSrc ? true : false;
        bridge.configureBridges(destChainIds, destBridges);
        bridge.configureToken(address(token), isNative, destChainIds, destTokens);
    }
}
