// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import { Initializable } from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import { AccessControlUpgradeable } from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import { PausableUpgradeable } from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import { XAppUpgradeable } from "core/src/pkg/XAppUpgradeable.sol";
import { IBridgeUpgradeable } from "./interfaces/IBridgeUpgradeable.sol";

import { ConfLevel } from "core/src/libraries/ConfLevel.sol";
import { TypeMax } from "core/src/libraries/TypeMax.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { IMintBurn } from "./interfaces/IMintBurn.sol";
import { ILockboxUpgradeable } from "./interfaces/ILockboxUpgradeable.sol";

contract BridgeUpgradeable is
    Initializable,
    AccessControlUpgradeable,
    PausableUpgradeable,
    XAppUpgradeable,
    IBridgeUpgradeable
{
    using SafeTransferLib for address;

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                         CONSTANTS                          */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    // keccak256("PAUSER")
    bytes32 public constant PAUSER_ROLE = 0x539440820030c4994db4e31b6b800deafd503688728f932addfe7a410515c14c;

    /**
     * @dev Default gas limit for xcalls.
     */
    uint64 internal constant DEFAULT_GAS_LIMIT = 140_000;

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                          STORAGE                           */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Address of the token being wrapped.
     */
    address public token;

    /**
     * @dev Address of the token wrapper being bridged.
     */
    address public wrapper;

    /**
     * @dev Lockbox (if assigned) indicating where the wrapper's original tokens are stored.
     */
    address public lockbox;

    /**
     * @dev Mapping of destination chainId to bridge contract.
     */
    mapping(uint64 chainId => address bridge) private _routes;

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                         MODIFIERS                          */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Modifier to restrict `receiveToken` access to bridge contracts.
     */
    modifier onlyBridge() {
        if (msg.sender == address(omni)) {
            if (_routes[xmsg.sourceChainId] != xmsg.sender) revert Unauthorized(xmsg.sourceChainId, xmsg.sender);
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

    function initialize(
        address admin_,
        address pauser_,
        address omni_,
        address wrapper_,
        address token_,
        address lockbox_
    ) external initializer {
        // Validate required inputs are not zero addresses.
        if (admin_ == address(0)) revert ZeroAddress();
        if (pauser_ == address(0)) revert ZeroAddress();
        if (omni_ == address(0)) revert ZeroAddress();
        if (wrapper_ == address(0)) revert ZeroAddress();

        // If either token or lockbox is set, both must be set.
        if (token_ != address(0) || lockbox_ != address(0)) {
            if (token_ == address(0)) revert BadInput();
            if (lockbox_ == address(0)) revert BadInput();
        }

        // Initialize everything and grant roles.
        __AccessControl_init();
        __Pausable_init();
        __XApp_init(omni_, ConfLevel.Finalized);
        _grantRole(DEFAULT_ADMIN_ROLE, admin_);
        _grantRole(PAUSER_ROLE, pauser_);

        // Set configured values.
        wrapper = wrapper_;
        if (token_ != address(0)) token = token_;
        if (lockbox_ != address(0)) lockbox = lockbox_;

        // Give lockbox relevant approvals to handle deposits and withdrawals if necessary.
        if (lockbox_ != address(0)) {
            token_.safeApproveWithRetry(lockbox_, type(uint256).max);
            wrapper_.safeApproveWithRetry(lockbox_, type(uint256).max);
        }
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                       VIEW FUNCTIONS                       */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Returns the bridge address for a given destination chainId.
     * @param destChainId The chainId of the destination chain.
     * @return bridge     The bridge address.
     */
    function routes(uint64 destChainId) external view returns (address bridge) {
        return _routes[destChainId];
    }

    /**
     * @dev Returns the fee for bridging a token to a destination chain.
     * @param destChainId The chainId of the destination chain.
     * @return fee        The fee paid to the `OmniPortal` contract.
     */
    function bridgeFee(uint64 destChainId) external view returns (uint256 fee) {
        return feeFor(
            destChainId,
            abi.encodeCall(BridgeUpgradeable.receiveToken, (TypeMax.Address, TypeMax.Uint256)),
            DEFAULT_GAS_LIMIT
        );
    }

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                      BRIDGE FUNCTIONS                      */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Bridges a token to a destination chain.
     * @param wrap        Whether to wrap the token first.
     * @param destChainId The chainId of the destination chain.
     * @param to          The address of the recipient.
     * @param value       The amount of tokens to bridge.
     */
    function sendToken(bool wrap, uint64 destChainId, address to, uint256 value) external payable whenNotPaused {
        _validateSend(wrap, destChainId, to, value);
        _sendToken(wrap, destChainId, to, value);
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
     * @dev Configures bridges for given chainIds.
     * @param chainIds    The chainIds to configure.
     * @param bridgeAddrs The bridges addresses to configure.
     */
    function configureBridges(uint64[] calldata chainIds, address[] calldata bridgeAddrs)
        external
        onlyRole(DEFAULT_ADMIN_ROLE)
    {
        if (chainIds.length != bridgeAddrs.length) revert ArrayLengthMismatch();
        for (uint256 i = 0; i < chainIds.length; i++) {
            _routes[chainIds[i]] = bridgeAddrs[i];
            emit BridgeConfigured(chainIds[i], bridgeAddrs[i]);
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
     * @param wrap        Whether the token is being wrapped.
     * @param destChainId The chainId of the destination chain.
     * @param to          The address of the recipient.
     * @param value       The amount of tokens to transfer.
     */
    function _validateSend(bool wrap, uint64 destChainId, address to, uint256 value) internal view {
        if (wrap && token == address(0)) revert CannotWrap();
        if (_routes[destChainId] == address(0)) revert InvalidRoute(destChainId);
        if (to == address(0)) revert ZeroAddress();
        if (value == 0) revert ZeroAmount();
    }

    /**
     * @dev Handles retrieving tokens from the sender and prepares for crosschain transfer.
     * @param wrap        Whether the token is being wrapped.
     * @param destChainId The chainId of the destination chain.
     * @param to          The address of the recipient.
     * @param value       The amount of tokens to transfer.
     */
    function _sendToken(bool wrap, uint64 destChainId, address to, uint256 value) internal {
        address _wrapper = wrapper; // Cache address for gas optimization.

        // Retrieve tokens from the sender, and deposit them into the lockbox for wrapping if necessary.
        if (wrap) {
            token.safeTransferFrom(msg.sender, address(this), value);
            ILockboxUpgradeable(lockbox).deposit(value);
        } else {
            _wrapper.safeTransferFrom(msg.sender, address(this), value);
        }

        // Burn the wrapped tokens.
        IMintBurn(_wrapper).burn(value);

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
        bytes memory data = abi.encodeCall(BridgeUpgradeable.receiveToken, (to, value));
        uint256 fee = xcall(destChainId, _routes[destChainId], data, DEFAULT_GAS_LIMIT);

        if (msg.value < fee) revert InsufficientPayment();
        if (msg.value > fee) msg.sender.safeTransferETH(msg.value - fee);

        emit CrosschainTransfer(destChainId, msg.sender, to, value);
    }

    /**
     * @dev Handles the receipt of tokens from the source chain.
     * @param to    The address of the recipient.
     * @param value The amount of tokens to receive.
     */
    function _receiveToken(address to, uint256 value) internal {
        // Cache addresses for gas optimization.
        address _lockbox = lockbox;
        address _wrapper = wrapper;

        if (_lockbox == address(0)) {
            // If no lockbox is set, just mint the wrapped tokens to the recipient.
            IMintBurn(_wrapper).mint(to, value);
        } else {
            // If a lockbox is set, mint the wrapped tokens to the bridge contract.
            IMintBurn(_wrapper).mint(address(this), value);

            // Attempt withdrawal from the lockbox, but transfer the wrapped tokens to the recipient if it fails.
            try ILockboxUpgradeable(_lockbox).withdrawTo(to, value) { }
            catch {
                _wrapper.safeTransfer(to, value);
                emit LockboxWithdrawalFailed(_lockbox, to, value);
            }
        }

        emit CrosschainReceive(xmsg.sourceChainId, to, value);
    }
}
