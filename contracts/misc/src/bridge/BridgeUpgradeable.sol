// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import { Initializable } from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import { UUPSUpgradeable } from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import { AccessControlUpgradeable } from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import { PausableUpgradeable } from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import { XAppUpgradeable } from "core/src/pkg/XAppUpgradeable.sol";
import { IBridgeUpgradeable } from "./interfaces/IBridgeUpgradeable.sol";

import { ConfLevel } from "core/src/libraries/ConfLevel.sol";
import { TypeMax } from "core/src/libraries/TypeMax.sol";
import { AddrUtils } from "../../../solve/src/ERC7683/lib/AddrUtils.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { FixedPointMathLib } from "solady/src/utils/FixedPointMathLib.sol";

import { ITokenUpgradeable } from "./interfaces/ITokenUpgradeable.sol";
import { ILockboxUpgradeable } from "./interfaces/ILockboxUpgradeable.sol";
import { IERC7683 } from "../../../solve/src/ERC7683/interfaces/IERC7683.sol";
import { IOriginSettler, ISolverNet } from "../../../solve/src/ERC7683/interfaces/ISolverNetInbox.sol";
import { ISolverNetOutbox } from "../../../solve/src/ERC7683/interfaces/ISolverNetOutbox.sol";

contract BridgeUpgradeable is
    Initializable,
    UUPSUpgradeable,
    AccessControlUpgradeable,
    PausableUpgradeable,
    XAppUpgradeable,
    IBridgeUpgradeable
{
    using SafeTransferLib for address;
    using AddrUtils for address;

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                         CONSTANTS                          */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    //keccak256("UPGRADER")
    bytes32 public constant UPGRADER_ROLE = 0xa615a8afb6fffcb8c6809ac0997b5c9c12b8cc97651150f14c8f6203168cff4c;
    //keccak256("PAUSER")
    bytes32 public constant PAUSER_ROLE = 0x539440820030c4994db4e31b6b800deafd503688728f932addfe7a410515c14c;

    /**
     * @dev Fee denominator in basis points.
     */
    uint16 public constant FEE_DENOMINATOR = 10_000;

    /**
     * @dev Default gas limit for xcalls.
     */
    uint64 internal constant DEFAULT_GAS_LIMIT = 105_000;

    /**
     * @dev EIP-712 typehash for the SolverNet order data.
     */
    bytes32 internal constant SOLVERNET_ORDER_TYPEHASH = keccak256(
        "OrderData(Call call,Deposit[] deposits)Call(uint64 chainId,bytes32 target,uint256 value,bytes data,TokenExpense[] expenses)TokenExpense(bytes32 token,bytes32 spender,uint256 amount)Deposit(bytes32 token,uint256 amount)"
    );

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                          STORAGE                           */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev The stablecoin lockbox where native tokens are deposited.
     */
    ILockboxUpgradeable public lockbox;

    /**
     * @dev The SolverNetInbox contract, used for fast bridging.
     */
    address public solverNetInbox;

    /**
     * @dev The SolverNetOutbox contract, used for fast bridging.
     */
    address public solverNetOutbox;

    /**
     * @dev The fast bridge fee in basis points.
     */
    uint16 public fastBridgeFee;

    /**
     * @dev Mapping of destination chainId to bridge contract.
     */
    mapping(uint64 chainId => address bridge) public bridgeRoutes;

    /**
     * @dev Mapping of token to whether it is the native representation of an ERC20 token.
     */
    mapping(address token => bool) public isNativeToken;

    /**
     * @dev Mapping of source token to destination chainId to destination token.
     */
    mapping(address srcToken => mapping(uint64 destChainId => address destToken)) public tokenRoutes;

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                         MODIFIERS                          */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Modifier to restrict `receiveToken` access to bridge contracts.
     */
    modifier onlyBridge() {
        if (msg.sender == address(omni)) {
            if (bridgeRoutes[xmsg.sourceChainId] != xmsg.sender) revert Unauthorized(xmsg.sourceChainId, xmsg.sender);
        } else {
            revert Unauthorized(uint64(block.chainid), msg.sender);
        }
        _;
    }

    /**
     * @dev Modifier to restrict `receiveTokenIntent` access to the SolverNetOutbox executor.
     */
    modifier onlySolverNetOutbox() {
        if (msg.sender != ISolverNetOutbox(solverNetOutbox).executor()) {
            revert Unauthorized(uint64(block.chainid), msg.sender);
        }
        _;
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                        CONSTRUCTOR                         */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                        INITIALIZER                         */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev This method is used to initialize the contract with values that we want to use to bootstrap and run things.
     * The modifier initializer here helps us block initialization in the constructor so that we initialize value only
     * when deploying the proxy and not the contract itself. The initializer also tracks how many times this method is
     * called and it can only be called once.
     */
    function initialize(
        address omni_,
        address solverNetInbox_,
        address solverNetOutbox_,
        address lockbox_,
        address admin_,
        address upgrader_,
        address pauser_,
        uint16 fastBridgeFee_
    ) external initializer {
        if (omni_ == address(0)) revert ZeroAddress();
        if (solverNetInbox_ == address(0)) revert ZeroAddress();
        if (solverNetOutbox_ == address(0)) revert ZeroAddress();
        if (lockbox_ == address(0)) revert ZeroAddress();
        if (admin_ == address(0)) revert ZeroAddress();
        if (upgrader_ == address(0)) revert ZeroAddress();
        if (pauser_ == address(0)) revert ZeroAddress();
        if (fastBridgeFee_ > FEE_DENOMINATOR) revert BadInput();

        __UUPSUpgradeable_init();
        __AccessControl_init();
        __Pausable_init();
        __XApp_init(omni_, ConfLevel.Finalized);
        _grantRole(DEFAULT_ADMIN_ROLE, admin_);
        _grantRole(UPGRADER_ROLE, upgrader_);
        _grantRole(PAUSER_ROLE, pauser_);
        lockbox = ILockboxUpgradeable(lockbox_);
        solverNetInbox = solverNetInbox_;
        solverNetOutbox = solverNetOutbox_;
        fastBridgeFee = fastBridgeFee_;
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                       VIEW FUNCTIONS                       */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Returns the fee for bridging a token to a destination chain.
     * @param destChainId The chainId of the destination chain.
     * @return fee The fee paid to the `OmniPortal` contract.
     */
    function bridgeFee(uint64 destChainId) external view returns (uint256 fee) {
        return feeFor(
            destChainId,
            abi.encodeCall(BridgeUpgradeable.receiveToken, (TypeMax.Address, TypeMax.Address, TypeMax.Uint256)),
            DEFAULT_GAS_LIMIT
        );
    }

    /**
     * @dev Returns the fee for bridging a token to a destination chain via SolverNet intent.
     * @param value The amount of tokens to bridge.
     * @return fee The fee paid to the `OmniPortal` contract.
     */
    function bridgeIntentFee(uint256 value) external view returns (uint256 fee) {
        return FixedPointMathLib.fullMulDivUnchecked(value, fastBridgeFee, FEE_DENOMINATOR);
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                      BRIDGE FUNCTIONS                      */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Bridges a token to a destination chain.
     * @param destChainId The chainId of the destination chain.
     * @param token The address of the source token.
     * @param to The address of the recipient.
     * @param value The amount of tokens to bridge.
     */
    function sendToken(uint64 destChainId, address token, address to, uint256 value) external payable whenNotPaused {
        (bool isNative, address bridge, address destToken) = _getTransferInfo(token, destChainId);

        _validateTransfer(bridge, destToken, destChainId, to, value);
        _handleSend(token, value, isNative);

        bytes memory data = abi.encodeCall(BridgeUpgradeable.receiveToken, (destToken, to, value));
        uint256 fee = xcall(destChainId, bridge, data, DEFAULT_GAS_LIMIT);

        if (msg.value < fee) revert InsufficientFunds();
        if (msg.value > fee) msg.sender.safeTransferETH(msg.value - fee);

        emit TokenSent(destChainId, token, destToken, to, value);
    }

    /**
     * @dev Initiates a fast crosschain token transfer via SolverNet intent.
     * @param destChainId The chainId of the destination chain.
     * @param token The address of the source token.
     * @param to The address of the recipient.
     * @param value The amount of tokens to transfer.
     * @param fillDeadline The deadline for the fill.
     */
    function sendTokenIntent(uint64 destChainId, address token, address to, uint256 value, uint32 fillDeadline)
        external
        whenNotPaused
    {
        (, address bridge, address destToken) = _getTransferInfo(token, destChainId);

        _validateTransfer(bridge, destToken, destChainId, to, value);
        token.safeTransferFrom(msg.sender, address(this), value);
        token.safeApproveWithRetry(solverNetInbox, value);

        IERC7683.OnchainCrossChainOrder memory order =
            _createSolverNetOrder(destChainId, token, destToken, to, value, fillDeadline);
        IOriginSettler(solverNetInbox).open(order);

        emit TokenSentIntent(destChainId, token, destToken, to, value);
    }

    /**
     * @dev Receives a token from a bridge contract and mints it to the recipient.
     * @param token The address of the source token.
     * @param to The address of the recipient.
     * @param value The amount of tokens to mint.
     */
    function receiveToken(address token, address to, uint256 value) external xrecv onlyBridge {
        _handleReceive(token, to, value);
        emit TokenReceived(xmsg.sourceChainId, token, to, value);
    }

    /**
     * @dev Receives a token from the SolverNetOutbox executor and mints it to the recipient.
     * @param srcChainId The chainId of the source chain.
     * @param token The address of the source token.
     * @param to The address of the recipient.
     * @param value The amount of tokens to mint.
     */
    function receiveTokenIntent(uint64 srcChainId, address token, address to, uint256 value)
        external
        xrecv
        onlySolverNetOutbox
    {
        token.safeTransferFrom(ISolverNetOutbox(solverNetOutbox).executor(), to, value);
        emit TokenReceivedIntent(srcChainId, token, to, value);
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                      ADMIN FUNCTIONS                       */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Configures bridges for a given chainId.
     * @param chainIds The chainIds to configure.
     * @param bridges The bridges to configure.
     */
    function configureBridges(uint64[] calldata chainIds, address[] calldata bridges)
        external
        onlyRole(DEFAULT_ADMIN_ROLE)
    {
        if (chainIds.length != bridges.length) revert ArrayLengthMismatch();
        for (uint256 i = 0; i < chainIds.length; i++) {
            bridgeRoutes[chainIds[i]] = bridges[i];
            emit BridgeConfigured(chainIds[i], bridges[i]);
        }
    }

    /**
     * @dev Configures token routes for a given source token and destination chainId.
     * @param srcTokens The source tokens to configure.
     * @param destChainIds The destination chainIds to configure.
     * @param destTokens The destination tokens to configure.
     */
    function configureTokens(
        address[] calldata srcTokens,
        bool[] calldata isNative,
        uint64[] calldata destChainIds,
        address[] calldata destTokens
    ) external onlyRole(DEFAULT_ADMIN_ROLE) {
        if (srcTokens.length != destChainIds.length || srcTokens.length != destTokens.length) {
            revert ArrayLengthMismatch();
        }
        for (uint256 i = 0; i < srcTokens.length; i++) {
            if (destChainIds[i] == block.chainid) revert InvalidChainId();

            tokenRoutes[srcTokens[i]][destChainIds[i]] = destTokens[i];
            if (isNative[i]) isNativeToken[srcTokens[i]] = true;

            emit TokenConfigured(srcTokens[i], destChainIds[i], destTokens[i], isNative[i]);
        }
    }

    /**
     * @dev Sets the fast bridge fee.
     * @param fastBridgeFee_ The new fast bridge fee.
     */
    function setFastBridgeFee(uint16 fastBridgeFee_) external onlyRole(DEFAULT_ADMIN_ROLE) {
        fastBridgeFee = fastBridgeFee_;
        emit FastBridgeFeeSet(fastBridgeFee_);
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                     INTERNAL FUNCTIONS                     */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Retrieves the transfer information for a given token and destination chainId.
     * @param token The address of the source token.
     * @param destChainId The chainId of the destination chain.
     * @return isNative Whether the token is the native representation of an ERC20 token.
     * @return bridge The address of the bridge.
     * @return destToken The address of the destination token.
     */
    function _getTransferInfo(address token, uint64 destChainId)
        internal
        view
        returns (bool isNative, address bridge, address destToken)
    {
        isNative = isNativeToken[token];
        bridge = bridgeRoutes[destChainId];
        destToken = tokenRoutes[token][destChainId];
    }

    /**
     * @dev Validates the transfer of tokens.
     * @param bridge The address of the bridge.
     * @param destToken The address of the destination token.
     * @param destChainId The chainId of the destination chain.
     * @param to The address of the recipient.
     * @param value The amount of tokens to transfer.
     */
    function _validateTransfer(address bridge, address destToken, uint64 destChainId, address to, uint256 value)
        internal
        pure
    {
        if (bridge == address(0) || destToken == address(0)) revert InvalidTokenRoute(destToken, destChainId);
        if (to == address(0)) revert ZeroAddress();
        if (value == 0) revert ZeroAmount();
    }

    /**
     * @dev Handles the transfer/burn of tokens from the sender through the bridge.
     * @param token The address of the source token.
     * @param value The amount of tokens to transfer.
     * @param isNative Whether the token is the native representation of an ERC20 token.
     */
    function _handleSend(address token, uint256 value, bool isNative) internal {
        if (isNative) {
            token.safeTransferFrom(msg.sender, address(this), value);
            token.safeApproveWithRetry(address(lockbox), value);
            lockbox.deposit(token, value);
        } else {
            token.safeTransferFrom(msg.sender, address(this), value);
            ITokenUpgradeable(token).burn(value);
        }
    }

    /**
     * @dev Handles the withdrawal/mint of tokens through the bridge to the recipient.
     * @param token The address of the source token.
     * @param to The address of the recipient.
     * @param value The amount of tokens to withdraw.
     */
    function _handleReceive(address token, address to, uint256 value) internal {
        bool isNative = isNativeToken[token];

        if (isNative) {
            lockbox.withdrawTo(token, to, value);
        } else {
            ITokenUpgradeable(token).mint(to, value);
        }
    }

    /**
     * @dev Creates a SolverNet order for fast bridging.
     * @param destChainId The chainId of the destination chain.
     * @param destToken The address of the destination token.
     * @param to The address of the recipient.
     * @param value The amount of tokens to bridge.
     * @param fillDeadline The deadline for the order to be filled.
     * @return order The SolverNet order.
     */
    function _createSolverNetOrder(
        uint64 destChainId,
        address srcToken,
        address destToken,
        address to,
        uint256 value,
        uint32 fillDeadline
    ) internal view returns (IERC7683.OnchainCrossChainOrder memory order) {
        address destBridge = bridgeRoutes[destChainId];
        uint256 solverFee = FixedPointMathLib.fullMulDivUnchecked(value, fastBridgeFee, FEE_DENOMINATOR);
        uint256 amount = value - solverFee;
        bytes memory data =
            abi.encodeCall(BridgeUpgradeable.receiveTokenIntent, (uint64(block.chainid), destToken, to, amount));

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
        deposits[0] = ISolverNet.Deposit({ token: srcToken.toBytes32(), amount: value });

        ISolverNet.OrderData memory orderData = ISolverNet.OrderData({ call: call, deposits: deposits });

        order = IERC7683.OnchainCrossChainOrder({
            fillDeadline: fillDeadline,
            orderDataType: SOLVERNET_ORDER_TYPEHASH,
            orderData: abi.encode(orderData)
        });

        return order;
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                    OVERRIDDEN FUNCTIONS                    */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * An overridden method from {UUPSUpgradeable} which defines the permissions for authorizing an upgrade to a
     * new implementation.
     */
    function _authorizeUpgrade(address newImplementation) internal virtual override onlyRole(UPGRADER_ROLE) { }
}
