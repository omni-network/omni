// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { LockboxUpgradeable } from "src/bridge/LockboxUpgradeable.sol";
import { BridgeUpgradeable } from "src/bridge/BridgeUpgradeable.sol";
import { TokenUpgradeable } from "src/bridge/TokenUpgradeable.sol";
import { Proxy } from "src/bridge/Proxy.sol";

import { ILockboxUpgradeable } from "src/bridge/interfaces/ILockboxUpgradeable.sol";
import { IBridgeUpgradeable } from "src/bridge/interfaces/IBridgeUpgradeable.sol";
import { ITokenUpgradeable } from "src/bridge/interfaces/ITokenUpgradeable.sol";

import { Test } from "forge-std/Test.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { MockPortal } from "core/test/utils/MockPortal.sol";

contract TestBase is Test {
    LockboxUpgradeable internal lockboxImpl;
    BridgeUpgradeable internal bridgeImpl;
    TokenUpgradeable internal tokenImpl;

    ILockboxUpgradeable internal lockboxSrc;
    ILockboxUpgradeable internal lockboxA;
    ILockboxUpgradeable internal lockboxB;

    IBridgeUpgradeable internal bridgeSrc;
    IBridgeUpgradeable internal bridgeA;
    IBridgeUpgradeable internal bridgeB;

    ITokenUpgradeable internal tokenSrc;
    ITokenUpgradeable internal tokenA;
    ITokenUpgradeable internal tokenB;

    MockPortal internal omni;

    uint64 internal constant DEFAULT_GAS_LIMIT = 105_000;
    uint256 internal constant INITIAL_BALANCE = 1_000_000 ether;

    uint64 internal constant srcChainId = 1;
    uint64 internal constant destChainIdA = 2;
    uint64 internal constant destChainIdB = 3;

    address internal user = makeAddr("user");
    address internal admin = makeAddr("admin");
    address internal upgrader = makeAddr("upgrader");
    address internal pauser = makeAddr("pauser");
    address internal clawbacker = makeAddr("clawbacker");

    modifier prankUser() {
        vm.startPrank(user);
        _;
        vm.stopPrank();
    }

    modifier prankAdmin() {
        vm.startPrank(admin);
        _;
        vm.stopPrank();
    }

    function setUp() public {
        deploy();
        configureBridges();

        vm.deal(user, 1 ether);
        vm.prank(admin);
        tokenSrc.mint(user, INITIAL_BALANCE);
    }

    function deploy() internal {
        omni = new MockPortal();

        lockboxImpl = new LockboxUpgradeable();
        bridgeImpl = new BridgeUpgradeable();
        tokenImpl = new TokenUpgradeable();

        lockboxSrc = _deployLockbox();
        lockboxA = _deployLockbox();
        lockboxB = _deployLockbox();

        bridgeSrc = _deployBridge(address(lockboxSrc));
        bridgeA = _deployBridge(address(lockboxA));
        bridgeB = _deployBridge(address(lockboxB));

        tokenSrc = _deployToken("Ripple USD", "RLUSD", admin);
        tokenA = _deployToken("Bridged RLUSD (Omni)", "RLUSD.e", address(bridgeA));
        tokenB = _deployToken("Bridged RLUSD (Omni)", "RLUSD.e", address(bridgeB));
    }

    function configureBridges() internal prankAdmin {
        uint64[] memory destChainIds = new uint64[](2);
        address[] memory destBridges = new address[](2);
        address[] memory destTokens = new address[](2);

        destChainIds[0] = destChainIdA;
        destBridges[0] = address(bridgeA);
        destTokens[0] = address(tokenA);

        destChainIds[1] = destChainIdB;
        destBridges[1] = address(bridgeB);
        destTokens[1] = address(tokenB);

        _configureBridge(bridgeSrc, tokenSrc, destChainIds, destBridges, destTokens);

        destChainIds[0] = destChainIdB;
        destBridges[0] = address(bridgeB);
        destTokens[0] = address(tokenB);

        destChainIds[1] = srcChainId;
        destBridges[1] = address(bridgeSrc);
        destTokens[1] = address(tokenSrc);

        _configureBridge(bridgeA, tokenA, destChainIds, destBridges, destTokens);

        destChainIds[0] = srcChainId;
        destBridges[0] = address(bridgeSrc);
        destTokens[0] = address(tokenSrc);

        destChainIds[1] = destChainIdA;
        destBridges[1] = address(bridgeA);
        destTokens[1] = address(tokenA);

        _configureBridge(bridgeB, tokenB, destChainIds, destBridges, destTokens);
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

        srcBridge.sendToken{ value: fee }(destChainId, srcToken, to, value);
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
        returns (ITokenUpgradeable)
    {
        return ITokenUpgradeable(
            address(
                new Proxy(
                    address(tokenImpl),
                    abi.encodeCall(
                        TokenUpgradeable.initialize, (name, symbol, mintAuthority, admin, upgrader, pauser, clawbacker)
                    )
                )
            )
        );
    }

    function _configureBridge(
        IBridgeUpgradeable bridge,
        ITokenUpgradeable token,
        uint64[] memory destChainIds,
        address[] memory destBridges,
        address[] memory destTokens
    ) internal {
        bridge.configureBridges(destChainIds, destBridges);

        address[] memory srcTokens = new address[](destTokens.length);
        bool[] memory isNative = new bool[](destTokens.length);

        for (uint256 i; i < srcTokens.length; ++i) {
            srcTokens[i] = address(token);
            if (token == tokenSrc) isNative[i] = true;
        }

        bridge.configureTokens(srcTokens, isNative, destChainIds, destTokens);
    }
}
