---
sidebar_position: 3
id: building
---

# Building Cross-Rollup Apps

Omni is only possible because of a series of recent innovations in blockchain infrastructure. This section provides a high level overview of the "How It Works" and the components that work together to establish the final product.

For a more detailed overview, see the [Protocol](../protocol/introduction/introduction.md) section.

## Global EVM-Compatible Application Development

Omni offers a programmable, EVM-compatible layer that enables the development of cross-rollup applications, utilizing fast settlement and secure interoperability.

It integrates seamlessly with rollups, allowing developers to access state, messages, and liquidity, simplifying the creation of decentralized applications. Omni's approach sets a new standard, enabling developers to leverage the advantages of rollups cost-effectively while ensuring broad accessibility for Web3 users.

To jump right into developing cross-chain applications read further about developing using Omni [here](../develop/introduction.md).

## The Omni Native Token

The native token is integral to Omni's ecosystem, serving three primary functions:

- **acting as gas for transactions** on the Omni EVM and cross-chain applications,
- facilitating a **global gas marketplace for cross-rollup transactions**, and
- **enabling network decentralization** by allowing stakeholders to vote independently of ETH.

This design allows Omni to manage DDoS attacks, provide "gas liquidity" for seamless transactions across rollups, and progressively decentralize governance. This approach ensures backward compatibility with local gas tokens through an abstracted gas marketplace, highlighting the native token's role in enhancing the network's functionality and security.

For detailed information on how the native token is utilized for gas abstraction and supports global application development, refer to the protocol [Fees](../protocol/xmessages/fees/fees.md) section.

## Restaking ETH

The Omni Network consists of validators who restake **\$ETH** and monitor the state of rollups. These validators relay state updates from one domain to others and provide crypto-economic assurance of validity. Read more about economic security in the [Omni Security Model section](../protocol/security/implementation.md).

This means that the Omni Network is going to be the first platform that provides developers with a global view of state from all rollups, making cross-rollup application development extremely simple. All of this is made possible through the unique insight to use Ethereum's validator set to aggregate a global perspective of Ethereum's L2 ecosystem.

Read on how to delegate restaked **\$ETH** to Omni with EigenLayer in the [Delegating ETH](./delegate.md) section and read more on how Omni handles restaking in the [Restaking](../protocol/security/restaking.md) protocol section.
