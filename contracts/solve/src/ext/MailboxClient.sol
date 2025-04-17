// SPDX-License-Identifier: MIT OR Apache-2.0
pragma solidity ^0.8.24;

/*@@@@@@@       @@@@@@@@@
 @@@@@@@@@       @@@@@@@@@
  @@@@@@@@@       @@@@@@@@@
   @@@@@@@@@       @@@@@@@@@
    @@@@@@@@@@@@@@@@@@@@@@@@@
     @@@@@  HYPERLANE  @@@@@@@
    @@@@@@@@@@@@@@@@@@@@@@@@@
   @@@@@@@@@       @@@@@@@@@
  @@@@@@@@@       @@@@@@@@@
 @@@@@@@@@       @@@@@@@@@
@@@@@@@@@       @@@@@@@@*/

// ============ Internal Imports ============
import { IMailbox } from "@hyperlane-xyz/core/contracts/interfaces/IMailbox.sol";
import { IPostDispatchHook } from "@hyperlane-xyz/core/contracts/interfaces/hooks/IPostDispatchHook.sol";
import { IInterchainSecurityModule } from "@hyperlane-xyz/core/contracts/interfaces/IInterchainSecurityModule.sol";
import { Message } from "@hyperlane-xyz/core/contracts/libs/Message.sol";
import { AddrUtils } from "../lib/AddrUtils.sol";
import { PackageVersioned } from "@hyperlane-xyz/core/contracts/PackageVersioned.sol";
import { StandardHookMetadata } from "@hyperlane-xyz/core/contracts/hooks/libs/StandardHookMetadata.sol";

// ============ External Imports ============
import { OwnableRoles } from "solady/src/auth/OwnableRoles.sol";
import { Initializable } from "solady/src/utils/Initializable.sol";

abstract contract MailboxClient is OwnableRoles, Initializable, PackageVersioned {
    using Message for bytes;
    using AddrUtils for address;

    // ============ Events ============
    event HookSet(address _hook);
    event IsmSet(address _ism);

    // ============ Immutable Storage ============
    IMailbox public immutable mailbox;

    uint32 public immutable localDomain;

    // ============ Mutable Storage ============
    IPostDispatchHook public hook;

    IInterchainSecurityModule public interchainSecurityModule;

    uint256[48] private __GAP; // gap for upgrade safety

    /**
     * @notice Only accept messages from a Hyperlane Mailbox contract
     */
    modifier onlyMailbox() {
        if (msg.sender != address(mailbox)) revert Unauthorized();
        _;
    }

    constructor(address _mailbox) {
        mailbox = IMailbox(_mailbox);
        localDomain = mailbox.localDomain();
    }

    // ============ External functions ============
    /**
     * @notice Sets the address of the application's custom hook.
     * @param _hook The address of the hook contract.
     */
    function setHook(address _hook) public virtual onlyOwner {
        hook = IPostDispatchHook(_hook);
        emit HookSet(_hook);
    }

    /**
     * @notice Sets the address of the application's custom interchain security module.
     * @param _module The address of the interchain security module contract.
     */
    function setInterchainSecurityModule(address _module) public onlyOwner {
        interchainSecurityModule = IInterchainSecurityModule(_module);
        emit IsmSet(_module);
    }

    // ============ Internal functions ============
    /**
     * @notice Dispatch a message to a destination domain.
     * @param _destinationDomain The destination domain.
     * @param _target The target address.
     * @param _value The value to send with the message.
     * @param _messageBody The message body.
     * @param _gasLimit The gas limit.
     * @param _hook The hook to use.
     */
    function _dispatch(
        uint32 _destinationDomain,
        address _target,
        uint256 _value,
        bytes memory _messageBody,
        uint256 _gasLimit,
        IPostDispatchHook _hook
    ) internal returns (bytes32) {
        bytes memory _hookMetadata = StandardHookMetadata.overrideGasLimit(_gasLimit);
        // Use default hook if none is configured
        if (address(_hook) == address(0)) {
            return
                mailbox.dispatch{ value: _value }(_destinationDomain, _target.toBytes32(), _messageBody, _hookMetadata);
        } else {
            return mailbox.dispatch{ value: _value }(
                _destinationDomain, _target.toBytes32(), _messageBody, _hookMetadata, _hook
            );
        }
    }

    /**
     * @notice Quote the dispatch fee for a message to a destination domain.
     * @param _destinationDomain The destination domain.
     * @param _target The target address.
     * @param _messageBody The message body.
     * @param _gasLimit The gas limit.
     * @param _hook The hook to use.
     */
    function _quoteDispatch(
        uint32 _destinationDomain,
        address _target,
        bytes memory _messageBody,
        uint256 _gasLimit,
        IPostDispatchHook _hook
    ) internal view returns (uint256) {
        bytes memory _hookMetadata = StandardHookMetadata.overrideGasLimit(_gasLimit);
        // Use default hook if none is configured
        if (address(_hook) == address(0)) {
            return mailbox.quoteDispatch(_destinationDomain, _target.toBytes32(), _messageBody, _hookMetadata);
        } else {
            return mailbox.quoteDispatch(_destinationDomain, _target.toBytes32(), _messageBody, _hookMetadata, _hook);
        }
    }
}
