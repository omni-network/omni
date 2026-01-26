// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import { Initializable } from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import { AccessControlUpgradeable } from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import { PausableUpgradeable } from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import { XAppUpgradeable } from "core/src/pkg/XAppUpgradeable.sol";
import { IBridge } from "./interfaces/IBridge.sol";

import { ConfLevel } from "core/src/libraries/ConfLevel.sol";
import { TypeMax } from "core/src/libraries/TypeMax.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { ITokenOps } from "./interfaces/ITokenOps.sol";
import { ILockbox } from "./interfaces/ILockbox.sol";
import { IERC20 } from "@openzeppelin/contracts/interfaces/IERC20.sol";

contract Bridge is Initializable, AccessControlUpgradeable, PausableUpgradeable, XAppUpgradeable, IBridge {
    using SafeTransferLib for address;

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                         CONSTANTS                          */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    // keccak256("PAUSER");
    bytes32 public constant PAUSER_ROLE = 0x539440820030c4994db4e31b6b800deafd503688728f932addfe7a410515c14c;
    // keccak256("UNPAUSER");
    bytes32 public constant UNPAUSER_ROLE = 0x82b32d9ab5100db08aeb9a0e08b422d14851ec118736590462bf9c085a6e9448;
    // keccak256("CONFIGURER");
    bytes32 public constant CONFIGURER_ROLE = 0x527e2c92bb6983874717bce74818faf5a9be45b6e85909ee478af653c6d98755;
    // keccak256("AUTHORIZER");
    bytes32 public constant AUTHORIZER_ROLE = 0x94858e5561d6a33fcce848f16acfe1514fe5166e32b456aff42d7fb50e8c52ad;

    /**
     * @dev Gas limit for xcalls to bridges without a lockbox. (~125k suggested unless transfer is nonstandard)
     */
    uint64 private immutable RECEIVE_DEFAULT_GAS_LIMIT;

    /**
     * @dev Gas limit for xcalls to bridges with a lockbox. (~200k suggested unless clawback is nonstandard)
     */
    uint64 private immutable RECEIVE_LOCKBOX_GAS_LIMIT;

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                          STORAGE                           */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Address of the token being bridged.
     */
    address public token;

    /**
     * @dev Lockbox (if assigned) indicating where the token is wrapped.
     */
    address public lockbox;

    /**
     * @dev Mapping of destination chainId to bridge contract and config.
     */
    mapping(uint64 chainId => Route) private _routes;

    /**
     * @dev Mapping of destination chainId to pending route updates.
     */
    mapping(uint64 chainId => Route) private _pendingRoutes;

    /**
     * @dev Track claimable amount per address, which increases when `token` reverts in `receiveToken`.
     */
    mapping(address => uint256) public claimable;

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                         MODIFIERS                          */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Modifier to restrict `receiveToken` access to bridge contracts.
     */
    modifier onlyBridge() {
        if (msg.sender == address(omni)) {
            if (_routes[xmsg.sourceChainId].bridge != xmsg.sender) {
                revert Unauthorized(xmsg.sourceChainId, xmsg.sender);
            }
        } else {
            revert Unauthorized(uint64(block.chainid), msg.sender);
        }
        _;
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                        CONSTRUCTOR                         */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    constructor(uint64 receiveDefaultGasLimit, uint64 receiveLockboxGasLimit) {
        RECEIVE_DEFAULT_GAS_LIMIT = receiveDefaultGasLimit;
        RECEIVE_LOCKBOX_GAS_LIMIT = receiveLockboxGasLimit;
        _disableInitializers();
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                        INITIALIZER                         */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    function initialize(
        address admin_,
        address configurer_,
        address authorizer_,
        address pauser_,
        address unpauser_,
        address omni_,
        address token_,
        address lockbox_
    ) external initializer {
        // Validate required inputs are not zero addresses.
        if (admin_ == address(0)) revert ZeroAddress();
        if (configurer_ == address(0)) revert ZeroAddress();
        if (authorizer_ == address(0)) revert ZeroAddress();
        if (pauser_ == address(0)) revert ZeroAddress();
        if (unpauser_ == address(0)) revert ZeroAddress();
        if (omni_ == address(0)) revert ZeroAddress();
        if (token_ == address(0)) revert ZeroAddress();

        // Initialize everything and grant roles.
        __AccessControl_init();
        __Pausable_init();
        __XApp_init(omni_, ConfLevel.Finalized);
        _grantRole(DEFAULT_ADMIN_ROLE, admin_);
        _grantRole(CONFIGURER_ROLE, configurer_);
        _grantRole(AUTHORIZER_ROLE, authorizer_);
        _grantRole(PAUSER_ROLE, pauser_);
        _grantRole(UNPAUSER_ROLE, unpauser_);
        token = token_;

        // Give lockbox relevant approvals to handle deposits and withdrawals if necessary.
        if (lockbox_ != address(0)) {
            lockbox = lockbox_;
            ILockbox(lockbox_).token().safeApproveWithRetry(lockbox_, type(uint256).max);
        }
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                       VIEW FUNCTIONS                       */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Returns the bridge address and config for a given destination chainId.
     * @param destChainId The chainId of the destination chain.
     * @return bridge     The bridge address.
     * @return hasLockbox Whether the bridge has a lockbox.
     */
    function getRoute(uint64 destChainId) external view returns (address bridge, bool hasLockbox) {
        Route memory route = _routes[destChainId];
        if (route.bridge == address(0)) revert InvalidRoute(destChainId);
        return (route.bridge, route.hasLockbox);
    }

    /**
     * @dev Returns the fee for bridging a token to a destination chain.
     * @param destChainId The chainId of the destination chain.
     * @return fee        The fee paid to the `OmniPortal` contract.
     */
    function bridgeFee(uint64 destChainId) external view returns (uint256 fee) {
        Route memory route = _routes[destChainId];
        if (route.bridge == address(0)) revert InvalidRoute(destChainId);
        return feeFor(
            destChainId,
            abi.encodeCall(Bridge.receiveToken, (TypeMax.Address, TypeMax.Uint256)),
            route.hasLockbox ? RECEIVE_LOCKBOX_GAS_LIMIT : RECEIVE_DEFAULT_GAS_LIMIT
        );
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                      BRIDGE FUNCTIONS                      */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Bridges a token to a destination chain.
     * @param destChainId The chainId of the destination chain.
     * @param to          The address of the recipient.
     * @param value       The amount of tokens to bridge.
     * @param wrap        Whether to wrap the token first.
     * @param refundTo    The address to refund any excess payment to.
     */
    function sendToken(uint64 destChainId, address to, uint256 value, bool wrap, address refundTo)
        external
        payable
        whenNotPaused
    {
        _validateSend(destChainId, to, value, wrap, refundTo);
        _sendToken(destChainId, to, value, wrap, refundTo);
    }

    /**
     * @dev Receives a token from a bridge contract and mints/unwraps it to the recipient.
     * @param to    The address of the recipient.
     * @param value The amount of tokens to mint/unwrap.
     */
    function receiveToken(address to, uint256 value) external xrecv onlyBridge {
        _receiveToken(to, value);
    }

    /**
     * @dev Attempts to transfer claimable tokens to the recipient.
     * @param addr The address to retry the transfer for.
     */
    function claimFailedReceive(address addr) external whenNotPaused {
        uint256 value = claimable[addr];
        if (value == 0) revert NoClaimable();

        delete claimable[addr];

        if (lockbox == address(0)) {
            ITokenOps(token).mint(addr, value);
        } else {
            ITokenOps(token).mint(address(this), value);
            ILockbox(lockbox).withdrawTo(addr, value);
        }

        emit RetrySuccessful(msg.sender, addr, value);
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                      ADMIN FUNCTIONS                       */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Configures bridge routes for given chainIds.
     * @param chainIds The chainIds to configure.
     * @param routes   The bridges addresses and configs to configure.
     */
    function configureRoutes(uint64[] calldata chainIds, Route[] calldata routes) external onlyRole(CONFIGURER_ROLE) {
        if (chainIds.length != routes.length) revert ArrayLengthMismatch();
        for (uint256 i = 0; i < chainIds.length; i++) {
            _pendingRoutes[chainIds[i]] = routes[i];
            emit RouteConfigured(chainIds[i], routes[i].bridge, routes[i].hasLockbox);
        }
    }

    /**
     * @dev Authorizes pending bridge routes, manual specification prevents frontrunning.
     * @param chainIds       The chainIds for pending routes to authorize.
     * @param expectedRoutes The expected routes to authorize.
     */
    function authorizeRoutes(uint64[] calldata chainIds, Route[] calldata expectedRoutes)
        external
        onlyRole(AUTHORIZER_ROLE)
    {
        for (uint256 i = 0; i < chainIds.length; i++) {
            Route memory pendingRoute = _pendingRoutes[chainIds[i]];
            if (
                pendingRoute.bridge != expectedRoutes[i].bridge
                    || pendingRoute.hasLockbox != expectedRoutes[i].hasLockbox
            ) revert InvalidRoute(chainIds[i]);
            _routes[chainIds[i]] = pendingRoute;
            delete _pendingRoutes[chainIds[i]];
            emit RouteAuthorized(chainIds[i], pendingRoute.bridge, pendingRoute.hasLockbox);
        }
    }

    /**
     * @dev Pauses performing crosschain transfers.
     */
    function pause() external onlyRole(PAUSER_ROLE) {
        _pause();
    }

    /**
     * @dev Unpauses performing crosschain transfers.
     */
    function unpause() external onlyRole(UNPAUSER_ROLE) {
        _unpause();
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                     INTERNAL FUNCTIONS                     */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Validates the outbound transfer of tokens.
     * @param destChainId The chainId of the destination chain.
     * @param to          The address of the recipient.
     * @param value       The amount of tokens to transfer.
     * @param wrap        Whether the token is being wrapped.
     * @param refundTo    The address to refund any excess payment to.
     */
    function _validateSend(uint64 destChainId, address to, uint256 value, bool wrap, address refundTo) internal view {
        if (_routes[destChainId].bridge == address(0)) revert InvalidRoute(destChainId);
        if (to == address(0)) revert ZeroAddress();
        if (value == 0) revert ZeroAmount();
        if (wrap && lockbox == address(0)) revert CannotWrap();
        if (refundTo == address(0)) revert ZeroAddress();
    }

    /**
     * @dev Handles retrieving tokens from the sender and prepares for crosschain transfer.
     * @param destChainId The chainId of the destination chain.
     * @param to          The address of the recipient.
     * @param value       The amount of tokens to transfer.
     * @param wrap        Whether the token is being wrapped.
     * @param refundTo    The address to refund any excess payment to.
     */
    function _sendToken(uint64 destChainId, address to, uint256 value, bool wrap, address refundTo) internal {
        // Cache addresses for gas optimization.
        address _token = token;
        address _lockbox = lockbox;

        // Retrieve tokens from the sender, and deposit them into the lockbox for wrapping if necessary.
        if (wrap) {
            ILockbox(_lockbox).token().safeTransferFrom(msg.sender, address(this), value);
            ILockbox(_lockbox).depositTo(msg.sender, value);
        }

        // Burn the tokens.
        ITokenOps(_token).clawback(msg.sender, value);

        // Send the tokens to the destination chain.
        _omniTransfer(destChainId, to, value, refundTo);
    }

    /**
     * @dev Handles the crosschain transfer of tokens.
     * @param destChainId The chainId of the destination chain.
     * @param to          The address of the recipient.
     * @param value       The amount of tokens to transfer.
     * @param refundTo    The address to refund any excess payment to.
     */
    function _omniTransfer(uint64 destChainId, address to, uint256 value, address refundTo) internal {
        Route memory route = _routes[destChainId];
        bytes memory data = abi.encodeCall(Bridge.receiveToken, (to, value));
        uint256 fee = xcall(
            destChainId, route.bridge, data, route.hasLockbox ? RECEIVE_LOCKBOX_GAS_LIMIT : RECEIVE_DEFAULT_GAS_LIMIT
        );

        if (msg.value < fee) revert InsufficientPayment();
        if (msg.value > fee) refundTo.safeTransferETH(msg.value - fee);

        emit TokenSent(destChainId, msg.sender, to, value);
    }

    /**
     * @dev Handles the receipt of tokens from the source chain.
     * @param to    The address of the recipient.
     * @param value The amount of tokens to receive.
     */
    function _receiveToken(address to, uint256 value) internal {
        // Cache addresses for gas optimization.
        address _token = token;
        address _lockbox = lockbox;
        bool success;

        if (_lockbox == address(0)) {
            // If no lockbox is set, just mint the wrapped tokens to the recipient.
            success = _tryCatchTokenMint(_token, to, value, false);
        } else {
            // If a lockbox is set, mint the wrapped tokens to the bridge contract.
            success = _tryCatchTokenMint(_token, to, value, true);
            // Attempt withdrawal from the lockbox, but transfer the wrapped tokens to the recipient if it fails.
            if (success) success = _tryCatchLockboxWithdrawal(_token, _lockbox, to, value);
        }

        emit TokenReceived(xmsg.sourceChainId, to, value, success);
    }

    /**
     * @dev Attempts to mint tokens, but caches the tokens for the recipient if it reverts.
     * @param _token       The address of the token.
     * @param to           The address of the recipient.
     * @param value        The amount of tokens to mint.
     * @param intermediate Whether to attempt to mint to the bridge contract directly.
     */
    function _tryCatchTokenMint(address _token, address to, uint256 value, bool intermediate) internal returns (bool) {
        try ITokenOps(_token).mint(intermediate ? address(this) : to, value) { }
        catch {
            _incrementClaimable(to, value);
            emit TokenMintFailed(_token, to, value);
            return false;
        }
        return true;
    }

    /**
     * @dev Attempts to withdraw tokens from the lockbox, sends tokens to the recipient if it fails.
     * @param _token    The address of the token.
     * @param _lockbox  The address of the lockbox.
     * @param to        The address of the recipient.
     * @param value     The amount of tokens to withdraw.
     */
    function _tryCatchLockboxWithdrawal(address _token, address _lockbox, address to, uint256 value)
        internal
        returns (bool)
    {
        try ILockbox(_lockbox).withdrawTo(to, value) { }
        catch {
            try IERC20(_token).transfer(to, value) returns (bool success) {
                if (!success) {
                    _incrementClaimable(to, value);
                    emit TokenTransferFailed(_token, to, value);
                    return false;
                }
            } catch {
                _incrementClaimable(to, value);
                emit LockboxWithdrawalFailed(_lockbox, to, value);
                return false;
            }
        }
        return true;
    }

    function _incrementClaimable(address to, uint256 value) internal {
        unchecked {
            claimable[to] += value;
        }
    }
}
