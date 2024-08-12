---
sidebar_position: 1
---

# Restaking to Omni

## Enhancing Security with Staked **\$ETH**

The Omni protocol introduces a groundbreaking approach to blockchain security through the innovative use of restaked **\$ETH**. This process involves validators and delegators working in tandem to verify protocol messages, manage reward and slashing events, and maintain the integrity of the validator set. Here, we explore how restaking contributes to the robust security of the Omni network.

## The Role of Validators and Delegators

In the Omni network, **validators** have the critical job of verifying the authenticity of protocol messages and managing the network's staking balances and voting powers. **Delegators**, on the other hand, support the network by delegating their restaked **\$ETH** to validators of their choice, thereby contributing to the overall security and efficiency of the network.

## How Restaking Works

Omni's staking implementation is supported by a suite of smart contracts designed to facilitate the staking process:

- **Omni Staking Contract**: Implemented on the Omni EVM, this contract tracks each validator’s stake and delegations, distributes rewards, and manages slashing events.
- **Omni AVS Contract**: Located on Ethereum, it registers Omni as an application with the EigenLayer protocol, enabling operators to provide validation services to the Omni blockchain.
- **Omni Portal Contracts**: Deployed on the Omni EVM and connected rollup VMs, these contracts maintain a record of the Omni validator set’s stake, delegations, and voting power.

## Securing the Omni Chain with Restaked **\$ETH**

Restaking **\$ETH** to secure the Omni chain leverages the existing cryptoeconomic security of the Ethereum network. By allowing Ethereum stakers to restake their **\$ETH** within the Omni ecosystem, Omni can tap into Ethereum's substantial security budget, significantly enhancing its own security without the limitations faced by other interoperability solutions.

### Dual Staking Mechanism

Omni employs a dual staking mechanism that combines restaked **\$ETH** and staked **\$OMNI** tokens, offering a layered security model. This approach not only allows Omni to inherit Ethereum's vast security resources but also enables it to scale its security budget over time as more OMNI tokens are staked.

### Communication and Updates

Omni uses its XMsg format to communicate staking events. Validators and delegators stay informed about stake changes and delegation updates through the Omni staking and AVS contracts, ensuring the network remains responsive and secure.

## The Impact of Restaking

The restaking model allows Omni to establish a new standard in protocol safety, offering security guarantees that are orders of magnitude higher than existing solutions. It facilitates a secure, decentralized environment for validators and delegators to contribute to the network's integrity, paving the way for a more interconnected and resilient blockchain ecosystem.

By integrating restaked **\$ETH** into its security model, Omni is setting a precedent for how blockchain networks can leverage existing assets to bootstrap and enhance their security, driving forward the vision of a secure, interconnected blockchain future.
