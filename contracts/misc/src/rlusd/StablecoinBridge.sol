// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import { Initializable } from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import { UUPSUpgradeable } from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import { AccessControlUpgradeable } from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import { PausableUpgradeable } from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import { XAppUpgradeable } from "core/src/pkg/XAppUpgradeable.sol";
import { IStablecoinBridgeUpgradeable } from "./interfaces/IStablecoinBridgeUpgradeable.sol";

import { ConfLevel } from "core/src/libraries/ConfLevel.sol";
import { TypeMax } from "core/src/libraries/TypeMax.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { IStablecoinUpgradeable } from "./interfaces/IStablecoinUpgradeable.sol";
import { IStablecoinLockboxUpgradeable } from "./interfaces/IStablecoinLockboxUpgradeable.sol";

contract StablecoinBridgeUpgradeable is
    Initializable,
    UUPSUpgradeable,
    AccessControlUpgradeable,
    PausableUpgradeable,
    XAppUpgradeable,
    IStablecoinBridgeUpgradeable
{
    using SafeTransferLib for address;

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                         CONSTANTS                          */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    //keccak256("UPGRADER")
    bytes32 public constant UPGRADER_ROLE = 0xa615a8afb6fffcb8c6809ac0997b5c9c12b8cc97651150f14c8f6203168cff4c;
    //keccak256("PAUSER")
    bytes32 public constant PAUSER_ROLE = 0x539440820030c4994db4e31b6b800deafd503688728f932addfe7a410515c14c;

    /**
     * @dev Default gas limit for xcalls.
     */
    uint64 internal constant DEFAULT_GAS_LIMIT = 100_000;

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                          STORAGE                           */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev The stablecoin lockbox where native tokens are deposited.
     */
    IStablecoinLockboxUpgradeable public lockbox;

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
    function initialize(address omni_, address lockbox_, address admin_, address upgrader_, address pauser_)
        external
        initializer
    {
        __UUPSUpgradeable_init();
        __AccessControl_init();
        __Pausable_init();
        __XApp_init(omni_, ConfLevel.Finalized);
        _grantRole(DEFAULT_ADMIN_ROLE, admin_);
        _grantRole(UPGRADER_ROLE, upgrader_);
        _grantRole(PAUSER_ROLE, pauser_);
        lockbox = IStablecoinLockboxUpgradeable(lockbox_);
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
            abi.encodeCall(
                StablecoinBridgeUpgradeable.receiveToken, (TypeMax.Address, TypeMax.Address, TypeMax.Uint256)
            ),
            DEFAULT_GAS_LIMIT
        );
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
        bool isNative = isNativeToken[token];
        address bridge = bridgeRoutes[destChainId];
        address destToken = tokenRoutes[token][destChainId];

        if (bridge == address(0) || destToken == address(0)) revert InvalidTokenRoute(token, destChainId);
        if (to == address(0)) revert ZeroAddress();
        if (value == 0) revert ZeroAmount();

        if (isNative) {
            token.safeTransferFrom(msg.sender, address(this), value);
            token.safeApproveWithRetry(address(lockbox), value);
            lockbox.deposit(token, value);
        } else {
            token.safeTransferFrom(msg.sender, address(this), value);
            IStablecoinUpgradeable(token).burn(value);
        }

        bytes memory data = abi.encodeCall(StablecoinBridgeUpgradeable.receiveToken, (destToken, to, value));
        uint256 fee = xcall(destChainId, bridge, data, DEFAULT_GAS_LIMIT);

        if (msg.value < fee) revert InsufficientFunds();
        if (msg.value > fee) msg.sender.safeTransferETH(msg.value - fee);

        emit TokenSent(destChainId, token, destToken, to, value);
    }

    /**
     * @dev Receives a token from a bridge contract and mints it to the recipient.
     * @param token The address of the source token.
     * @param to The address of the recipient.
     * @param value The amount of tokens to mint.
     */
    function receiveToken(address token, address to, uint256 value) external xrecv onlyBridge {
        bool isNative = isNativeToken[token];

        if (isNative) {
            lockbox.withdrawTo(token, to, value);
        } else {
            IStablecoinUpgradeable(token).mint(to, value);
        }

        emit TokenReceived(xmsg.sourceChainId, token, to, value);
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
            tokenRoutes[srcTokens[i]][destChainIds[i]] = destTokens[i];
            if (isNative[i]) isNativeToken[destTokens[i]] = true;
            emit TokenConfigured(srcTokens[i], destChainIds[i], destTokens[i], isNative[i]);
        }
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
