---
sidebar_position: 7
---

# Omni Integrated Consensus

Omni introduces a groundbreaking approach to blockchain consensus known as **Integrated Consensus**. This innovative method allows the Omni network to simultaneously process consensus on its Ethereum Virtual Machine (EVM) and manage cross-chain data or messages, known as XBlocks. This dual-path consensus mechanism enhances network performance, ensuring swift and secure validation of transactions and cross-chain interactions.

## How Integrated Consensus Works

Integrated Consensus in Omni consists of two main sub-processes, each handling a distinct aspect of network operations:

### Omni EVM Consensus

1. **Block Proposal**: Validators take turns proposing new blocks. The proposing validator's Halo client requests the latest block from the Omni EVM through the Engine API.
2. **Block Construction**: The execution client compiles a block from pending transactions and sends the block header back to the Halo client.
3. **Block Inclusion**: The proposed block is wrapped as a CometBFT transaction and added to the consensus layer block, ready for network validation.
4. **Validation**: Other validators use the Engine API to verify the proposed block's validity by running the state transition function on the block header.

### `XBlock` Consensus

1. **XMsg Event Monitoring**: Each validator's `halo` client monitors rollup VMs for `XMsg` events emitted by Portal contracts.
2. **`XBlock` Creation**: Validators construct `XBlock`s containing `XMsg`s for each rollup VM block that includes cross-chain messages.
3. **Attestation**: After cross-chain message data is finalized on Ethereum L1, validators attest to the hash of the corresponding `XBlock`, incorporating these attestations into the current consensus layer block.

### Finalizing Blocks

The combined efforts of validating Omni EVM transactions and attesting to `XBlock` hashes culminate in a unified consensus vote. If a proposed block garners approval from two-thirds of the validator set, it is finalized and added to the Omni chain. These finalized blocks encompass both the Omni EVM transactions and attested XBlock hashes, enabling full node operators of specific rollup VMs to reconstruct and verify XBlock contents.

## The Significance of Integrated Consensus

Omni's Integrated Consensus represents a significant leap forward in blockchain technology, enabling:

- **Efficient Processing**: By separating the consensus processes for the Omni EVM and `XBlock`s, Omni achieves rapid state transitions and block finalization.
- **Enhanced Security and Reliability**: The dual-path consensus strengthens the network's integrity, providing robust security for both on-chain transactions and cross-chain communications.
- **Global Scalability**: This consensus model paves the way for a scalable and interoperable blockchain ecosystem, where diverse networks can communicate and transact seamlessly.

Omni's adoption of Integrated Consensus underscores its commitment to fostering a more connected and efficient blockchain world, where the boundaries between different chains and technologies blur, creating a unified and powerful platform for decentralized applications and services.

For a deeper dive into Omni's Integrated Consensus and its impact on blockchain interoperability, [visit the protocol section](../../protocol/introduction/introduction.md).
