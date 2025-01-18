// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { LockboxUpgradeable } from "src/bridge/LockboxUpgradeable.sol";
import { BridgeUpgradeable } from "src/bridge/BridgeUpgradeable.sol";
import { TokenUpgradeable } from "src/bridge/TokenUpgradeable.sol";
import { Proxy } from "src/bridge/Proxy.sol";

import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { SolverNetInbox } from "../../../solve/src/ERC7683/SolverNetInbox.sol";
import { SolverNetOutbox } from "../../../solve/src/ERC7683/SolverNetOutbox.sol";
import { MockPortal } from "core/test/utils/MockPortal.sol";

import { ILockboxUpgradeable } from "src/bridge/interfaces/ILockboxUpgradeable.sol";
import { IBridgeUpgradeable } from "src/bridge/interfaces/IBridgeUpgradeable.sol";
import { ITokenUpgradeable } from "src/bridge/interfaces/ITokenUpgradeable.sol";
import { IERC7683 } from "../../../solve/src/ERC7683/interfaces/IERC7683.sol";
import { IOriginSettler, ISolverNet } from "../../../solve/src/ERC7683/interfaces/ISolverNetInbox.sol";
import { IDestinationSettler } from "../../../solve/src/ERC7683/interfaces/IDestinationSettler.sol";

import { Test } from "forge-std/Test.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { AddrUtils } from "../../../solve/src/ERC7683/lib/AddrUtils.sol";

contract TestBase is Test {
    using SafeTransferLib for address;
    using AddrUtils for address;

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

    SolverNetInbox internal solverNetInbox;
    SolverNetOutbox internal solverNetOutbox;

    MockPortal internal omni;

    uint16 internal constant FAST_BRIDGE_FEE = 50;
    uint64 internal constant DEFAULT_GAS_LIMIT = 105_000;
    uint256 internal constant INITIAL_USER_BALANCE = 1_000_000 ether;
    uint256 internal constant INITIAL_SOLVER_BALANCE = INITIAL_USER_BALANCE * 3;

    bytes32 internal constant SOLVERNET_ORDER_TYPEHASH = keccak256(
        "OrderData(Call call,Deposit[] deposits)Call(uint64 chainId,bytes32 target,uint256 value,bytes data,TokenExpense[] expenses)TokenExpense(bytes32 token,bytes32 spender,uint256 amount)Deposit(bytes32 token,uint256 amount)"
    );

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

        solverNetInbox = _deploySolverNetInbox();
        solverNetOutbox = _deploySolverNetOutbox();

        solverNetInbox.initialize(admin, solver, address(omni), address(solverNetOutbox));
        solverNetOutbox.initialize(admin, solver, address(omni), address(solverNetInbox));

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

    function configureBridges() internal prankUser(admin) {
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

        vm.chainId(sourceChainId);
        srcBridge.sendToken{ value: fee }(destChainId, srcToken, to, value);

        vm.chainId(destChainId);
        omni.mockXCall(sourceChainId, address(srcBridge), destBridge, data, DEFAULT_GAS_LIMIT);
    }

    function mockBridgeIntent(
        IBridgeUpgradeable srcBridge,
        uint64 sourceChainId,
        uint64 destChainId,
        address srcToken,
        address to,
        uint256 value
    ) internal {
        uint32 fillDeadline = uint32(block.timestamp + 5 minutes);

        vm.chainId(sourceChainId);
        vm.startPrank(user);
        IERC7683.OnchainCrossChainOrder memory order =
            _getSolverNetOrder(srcBridge, sourceChainId, destChainId, srcToken, to, value, fillDeadline);
        IERC7683.ResolvedCrossChainOrder memory resolvedOrder = solverNetInbox.resolve(order);

        bytes32 orderId = resolvedOrder.orderId;
        address destToken = srcBridge.tokenRoutes(srcToken, destChainId);
        uint256 fee = srcBridge.bridgeIntentFee(value);
        uint256 fillFee = solverNetOutbox.fillFee(sourceChainId);
        bytes32 fillHash = _fillHash(orderId, resolvedOrder.fillInstructions[0].originData);
        bytes memory data = abi.encodeCall(SolverNetInbox.markFilled, (orderId, fillHash));

        srcBridge.sendTokenIntent(destChainId, srcToken, to, value, fillDeadline);
        vm.stopPrank();

        vm.startPrank(solver);
        solverNetInbox.accept(orderId);

        vm.chainId(destChainId);
        destToken.safeApprove(address(solverNetOutbox), value - fee);
        IDestinationSettler(address(solverNetOutbox)).fill{ value: fillFee }(
            orderId, resolvedOrder.fillInstructions[0].originData, bytes("")
        );

        vm.chainId(sourceChainId);
        omni.mockXCall(destChainId, address(solverNetOutbox), address(solverNetInbox), data, DEFAULT_GAS_LIMIT);

        solverNetInbox.claim(orderId, solver);
        vm.stopPrank();
    }

    function _deploySolverNetInbox() internal returns (SolverNetInbox) {
        address impl = address(new SolverNetInbox());
        return SolverNetInbox(address(new TransparentUpgradeableProxy(impl, admin, bytes(""))));
    }

    function _deploySolverNetOutbox() internal returns (SolverNetOutbox) {
        address impl = address(new SolverNetOutbox());
        return SolverNetOutbox(address(new TransparentUpgradeableProxy(impl, admin, bytes(""))));
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
                        BridgeUpgradeable.initialize,
                        (
                            address(omni),
                            address(solverNetInbox),
                            address(solverNetOutbox),
                            address(lockbox),
                            admin,
                            upgrader,
                            pauser,
                            FAST_BRIDGE_FEE
                        )
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

    function _getSolverNetOrder(
        IBridgeUpgradeable srcBridge,
        uint64 sourceChainId,
        uint64 destChainId,
        address srcToken,
        address to,
        uint256 value,
        uint32 fillDeadline
    ) internal view returns (IERC7683.OnchainCrossChainOrder memory order) {
        address destBridge = srcBridge.bridgeRoutes(destChainId);
        address destToken = srcBridge.tokenRoutes(srcToken, destChainId);
        uint256 solverFee = srcBridge.bridgeIntentFee(value);
        uint256 amount = value - solverFee;
        bytes memory data = abi.encodeCall(BridgeUpgradeable.receiveTokenIntent, (sourceChainId, destToken, to, amount));

        ISolverNet.TokenExpense[] memory tokenExpense = new ISolverNet.TokenExpense[](1);
        tokenExpense[0] =
            ISolverNet.TokenExpense({ token: destToken.toBytes32(), spender: destBridge.toBytes32(), amount: amount });

        ISolverNet.Call memory call = ISolverNet.Call({
            chainId: destChainId,
            target: destBridge.toBytes32(),
            value: 0,
            data: data,
            expenses: tokenExpense
        });

        ISolverNet.Deposit[] memory deposits = new ISolverNet.Deposit[](1);
        deposits[0] = ISolverNet.Deposit({ token: destToken.toBytes32(), amount: value });

        ISolverNet.OrderData memory orderData = ISolverNet.OrderData({ call: call, deposits: deposits });

        order = IERC7683.OnchainCrossChainOrder({
            fillDeadline: fillDeadline,
            orderDataType: SOLVERNET_ORDER_TYPEHASH,
            orderData: abi.encode(orderData)
        });

        return order;
    }

    function _fillHash(bytes32 orderId, bytes memory originData) internal pure returns (bytes32) {
        return keccak256(abi.encode(orderId, originData));
    }
}
