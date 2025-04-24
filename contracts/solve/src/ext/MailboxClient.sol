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

    IInterchainSecurityModule public immutable interchainSecurityModule;

    uint32 public immutable localDomain;

    // ============ Mutable Storage ============
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
        interchainSecurityModule = mailbox.defaultIsm();
        localDomain = mailbox.localDomain();
    }

    // ============ Internal functions ============
    /**
     * @notice Dispatch a message to a destination domain.
     * @param _destinationDomain The destination domain.
     * @param _target The target address.
     * @param _value The value to send with the message.
     * @param _messageBody The message body.
     * @param _gasLimit The gas limit.
     */
    function _dispatch(
        uint32 _destinationDomain,
        address _target,
        uint256 _value,
        bytes memory _messageBody,
        uint256 _gasLimit
    ) internal returns (bytes32) {
        bytes memory _hookMetadata = StandardHookMetadata.overrideGasLimit(_gasLimit);
        return mailbox.dispatch{ value: _value }(_destinationDomain, _target.toBytes32(), _messageBody, _hookMetadata);
    }

    /**
     * @notice Quote the dispatch fee for a message to a destination domain.
     * @param _destinationDomain The destination domain.
     * @param _target The target address.
     * @param _messageBody The message body.
     * @param _gasLimit The gas limit.
     */
    function _quoteDispatch(uint32 _destinationDomain, address _target, bytes memory _messageBody, uint256 _gasLimit)
        internal
        view
        returns (uint256)
    {
        bytes memory _hookMetadata = StandardHookMetadata.overrideGasLimit(_gasLimit);
        return mailbox.quoteDispatch(_destinationDomain, _target.toBytes32(), _messageBody, _hookMetadata);
    }
}
