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

contract Bridge is Initializable, AccessControlUpgradeable, PausableUpgradeable, XAppUpgradeable, IBridge {
    using SafeTransferLib for address;

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                         CONSTANTS                          */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    // keccak256("PAUSER")
    bytes32 public constant PAUSER_ROLE = 0x539440820030c4994db4e31b6b800deafd503688728f932addfe7a410515c14c;

    /**
     * @dev Gas limit for xcalls to bridges without a lockbox.
     */
    uint64 internal constant RECEIVE_DEFAULT_GAS_LIMIT = 125_000;

    /**
     * @dev Gas limit for xcalls to bridges with a lockbox.
     */
    uint64 internal constant RECEIVE_LOCKBOX_GAS_LIMIT = 180_000;

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

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                         MODIFIERS                          */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Modifier to restrict `receiveToken` access to bridge contracts.
     */
    modifier onlyBridge() {
        if (msg.sender == address(omni)) {
            if (_routes[xmsg.sourceChainId].bridge != xmsg.sender) revert Unauthorized(xmsg.sourceChainId, xmsg.sender);
        } else {
            revert Unauthorized(uint64(block.chainid), msg.sender);
        }
        _;
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                        CONSTRUCTOR                         */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    constructor() {
        _disableInitializers();
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                        INITIALIZER                         */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    function initialize(address admin_, address pauser_, address omni_, address token_, address lockbox_)
        external
        initializer
    {
        // Validate required inputs are not zero addresses.
        if (admin_ == address(0)) revert ZeroAddress();
        if (pauser_ == address(0)) revert ZeroAddress();
        if (omni_ == address(0)) revert ZeroAddress();
        if (token_ == address(0)) revert ZeroAddress();

        // Initialize everything and grant roles.
        __AccessControl_init();
        __Pausable_init();
        __XApp_init(omni_, ConfLevel.Finalized);
        _grantRole(DEFAULT_ADMIN_ROLE, admin_);
        _grantRole(PAUSER_ROLE, pauser_);

        // Set configured values.
        token = token_;
        if (lockbox_ != address(0)) lockbox = lockbox_;

        // Give lockbox relevant approvals to handle deposits and withdrawals if necessary.
        if (lockbox_ != address(0)) {
            ILockbox(lockbox_).token().safeApproveWithRetry(lockbox_, type(uint256).max);
            token_.safeApproveWithRetry(lockbox_, type(uint256).max);
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
        return (route.bridge, route.hasLockbox);
    }

    /**
     * @dev Returns the fee for bridging a token to a destination chain.
     * @param destChainId The chainId of the destination chain.
     * @return fee        The fee paid to the `OmniPortal` contract.
     */
    function bridgeFee(uint64 destChainId) external view returns (uint256 fee) {
        Route memory route = _routes[destChainId];
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
     */
    function sendToken(uint64 destChainId, address to, uint256 value, bool wrap) external payable whenNotPaused {
        _validateSend(destChainId, to, value, wrap);
        _sendToken(destChainId, to, value, wrap);
    }

    /**
     * @dev Receives a token from a bridge contract and mints/unwraps it to the recipient.
     * @param to    The address of the recipient.
     * @param value The amount of tokens to mint/unwrap.
     */
    function receiveToken(address to, uint256 value) external xrecv onlyBridge {
        _receiveToken(to, value);
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                      ADMIN FUNCTIONS                       */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Sets bridge routes for given chainIds.
     * @param chainIds The chainIds to configure.
     * @param routes   The bridges addresses and configs to configure.
     */
    function setRoutes(uint64[] calldata chainIds, Route[] calldata routes) external onlyRole(DEFAULT_ADMIN_ROLE) {
        if (chainIds.length != routes.length) revert ArrayLengthMismatch();
        for (uint256 i = 0; i < chainIds.length; i++) {
            if (routes[i].bridge == address(0)) revert ZeroAddress();
            _routes[chainIds[i]] = routes[i];
            emit RouteConfigured(chainIds[i], routes[i].bridge, routes[i].hasLockbox);
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
    function unpause() external onlyRole(PAUSER_ROLE) {
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
     */
    function _validateSend(uint64 destChainId, address to, uint256 value, bool wrap) internal view {
        if (_routes[destChainId].bridge == address(0)) revert InvalidRoute(destChainId);
        if (to == address(0)) revert ZeroAddress();
        if (value == 0) revert ZeroAmount();
        if (wrap && lockbox == address(0)) revert CannotWrap();
    }

    /**
     * @dev Handles retrieving tokens from the sender and prepares for crosschain transfer.
     * @param destChainId The chainId of the destination chain.
     * @param to          The address of the recipient.
     * @param value       The amount of tokens to transfer.
     * @param wrap        Whether the token is being wrapped.
     */
    function _sendToken(uint64 destChainId, address to, uint256 value, bool wrap) internal {
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
        _omniTransfer(destChainId, to, value);
    }

    /**
     * @dev Handles the crosschain transfer of tokens.
     * @param destChainId The chainId of the destination chain.
     * @param to          The address of the recipient.
     * @param value       The amount of tokens to transfer.
     */
    function _omniTransfer(uint64 destChainId, address to, uint256 value) internal {
        Route memory route = _routes[destChainId];
        bytes memory data = abi.encodeCall(Bridge.receiveToken, (to, value));
        uint256 fee = xcall(
            destChainId, route.bridge, data, route.hasLockbox ? RECEIVE_LOCKBOX_GAS_LIMIT : RECEIVE_DEFAULT_GAS_LIMIT
        );

        if (msg.value < fee) revert InsufficientPayment();
        if (msg.value > fee) msg.sender.safeTransferETH(msg.value - fee);

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

        if (_lockbox == address(0)) {
            // If no lockbox is set, just mint the wrapped tokens to the recipient.
            ITokenOps(_token).mint(to, value);
        } else {
            // If a lockbox is set, mint the wrapped tokens to the bridge contract.
            ITokenOps(_token).mint(address(this), value);

            // Attempt withdrawal from the lockbox, but transfer the wrapped tokens to the recipient if it fails.
            try ILockbox(_lockbox).withdrawTo(to, value) { }
            catch {
                _token.safeTransfer(to, value);
                emit LockboxWithdrawalFailed(_lockbox, to, value);
            }
        }

        emit TokenReceived(xmsg.sourceChainId, to, value);
    }
}
