// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { Versioned } from "@hyperlane-xyz/core/contracts/upgrade/Versioned.sol";
import { TypeCasts } from "@hyperlane-xyz/core/contracts/libs/TypeCasts.sol";
import { Message } from "@hyperlane-xyz/core/contracts/libs/Message.sol";
import { IMessageRecipient } from "@hyperlane-xyz/core/contracts/interfaces/IMessageRecipient.sol";
import {
    IInterchainSecurityModule,
    ISpecifiesInterchainSecurityModule
} from "@hyperlane-xyz/core/contracts/interfaces/IInterchainSecurityModule.sol";
import { Mailbox } from "./Mailbox.sol";
import { IPostDispatchHook } from "@hyperlane-xyz/core/contracts/interfaces/hooks/IPostDispatchHook.sol";

import { TestIsm } from "@hyperlane-xyz/core/contracts/test/TestIsm.sol";
import { TestPostDispatchHook } from "@hyperlane-xyz/core/contracts/test/TestPostDispatchHook.sol";

contract MockMailbox is Mailbox {
    using Message for bytes;

    uint32 public inboundUnprocessedNonce = 0;
    uint32 public inboundProcessedNonce = 0;

    mapping(uint32 => MockMailbox) public remoteMailboxes;
    mapping(uint256 => bytes) public inboundMessages;

    constructor(uint32 _domain) Mailbox(_domain) {
        TestIsm ism = new TestIsm();
        defaultIsm = ism;

        TestPostDispatchHook hook = new TestPostDispatchHook();
        defaultHook = hook;
        requiredHook = hook;
        hook.setFee(0.001 ether);

        _transferOwnership(msg.sender);
        _disableInitializers();
    }

    function addRemoteMailbox(uint32 _domain, MockMailbox _mailbox) external {
        remoteMailboxes[_domain] = _mailbox;
    }

    function dispatch(
        uint32 destinationDomain,
        bytes32 recipientAddress,
        bytes calldata messageBody,
        bytes calldata metadata,
        IPostDispatchHook hook
    ) public payable override returns (bytes32) {
        bytes memory message = _buildMessage(destinationDomain, recipientAddress, messageBody);
        bytes32 id = super.dispatch(destinationDomain, recipientAddress, messageBody, metadata, hook);

        MockMailbox _destinationMailbox = remoteMailboxes[destinationDomain];
        require(address(_destinationMailbox) != address(0), "Missing remote mailbox");
        _destinationMailbox.addInboundMessage(message);

        return id;
    }

    function addInboundMessage(bytes calldata message) external {
        inboundMessages[inboundUnprocessedNonce] = message;
        inboundUnprocessedNonce++;
    }

    function processNextInboundMessage() public payable {
        bytes memory _message = inboundMessages[inboundProcessedNonce];
        Mailbox(address(this)).process{ value: msg.value }("", _message);
        inboundProcessedNonce++;
    }

    function processInboundMessage(uint32 _nonce) public payable {
        bytes memory _message = inboundMessages[_nonce];
        Mailbox(address(this)).process{ value: msg.value }("", _message);
    }

    function buildMessage(
        uint32 destinationDomain,
        bytes32 senderAddress,
        bytes32 recipientAddress,
        bytes calldata messageBody
    ) public view returns (bytes memory) {
        return Message.formatMessage(
            VERSION, nonce, localDomain, senderAddress, destinationDomain, recipientAddress, messageBody
        );
    }
}
