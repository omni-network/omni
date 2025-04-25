// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.13;

import "./MockMailbox.sol";
import "./TestInterchainGasPaymaster.sol";
import "@hyperlane-xyz/core/contracts/test/TestIsm.sol";

contract MockHyperlaneEnvironment {
    uint32 public originDomain;
    uint32 public destinationDomain;
    uint32 public invalidDomain;

    mapping(uint32 => MockMailbox) public mailboxes;
    mapping(uint32 => TestInterchainGasPaymaster) public igps;
    mapping(uint32 => IInterchainSecurityModule) public isms;

    constructor(uint32 _originDomain, uint32 _destinationDomain) {
        originDomain = _originDomain;
        destinationDomain = _destinationDomain;
        invalidDomain = _destinationDomain + 1;

        MockMailbox originMailbox = new MockMailbox(_originDomain);
        MockMailbox destinationMailbox = new MockMailbox(_destinationDomain);
        MockMailbox invalidMailbox = new MockMailbox(invalidDomain);

        originMailbox.addRemoteMailbox(_destinationDomain, destinationMailbox);
        originMailbox.addRemoteMailbox(invalidDomain, invalidMailbox);
        destinationMailbox.addRemoteMailbox(_originDomain, originMailbox);

        isms[originDomain] = new TestIsm();
        isms[invalidDomain] = new TestIsm();
        isms[destinationDomain] = new TestIsm();

        originMailbox.setDefaultIsm(address(isms[originDomain]));
        invalidMailbox.setDefaultIsm(address(isms[invalidDomain]));
        destinationMailbox.setDefaultIsm(address(isms[destinationDomain]));

        originMailbox.transferOwnership(msg.sender);
        invalidMailbox.transferOwnership(msg.sender);
        destinationMailbox.transferOwnership(msg.sender);

        mailboxes[_originDomain] = originMailbox;
        mailboxes[invalidDomain] = invalidMailbox;
        mailboxes[_destinationDomain] = destinationMailbox;
    }

    function processNextPendingMessage() public payable {
        mailboxes[destinationDomain].processNextInboundMessage{ value: msg.value }();
    }

    function processNextPendingMessageFromDestination() public {
        mailboxes[originDomain].processNextInboundMessage();
    }
}
