# Architecture

<img src="/img/staking.png" width="1000px"/>

## Consensus Engine + Execution Engine

Omni introduces an innovative approach to blockchain architecture, designed to enhance performance and scalability without compromising on security. At its core, Omni's architecture is split into two primary layers: the **Consensus Layer** and the **Execution Layer**. This dual-chain structure enables Omni to efficiently process transactions and manage global state across multiple networks.

## The Consensus Engine

At the heart of Omni's network lies the **Consensus Layer**, powered by the CometBFT consensus engine. This layer is where validators come together to agree on the state of the network, ensuring every transaction is valid and finalizing the global state across all connected networks within seconds. The key benefits include:

- **Immediate Transaction Finalization**: With CometBFT, once a transaction is included in a block, it's considered final—eliminating the need for additional confirmations.
- **Delegated Proof of Stake (DPoS)**: This mechanism allows users to delegate their restaked ETH directly to validators, bolstering the network's security.
- **Proven Robustness**: CometBFT is battle-tested and trusted in securing billions of dollars across various blockchain networks.

## The Execution Engine

Additional to the Consensus Layer is the **Execution Layer**, or the Omni EVM, mirroring the Ethereum L1 execution layer's functionality. This is where users' transactions are processed, and the following features stand out:

- **Scalable Transaction Processing**: The Omni EVM handles transactions in its mempool, enabling high throughput without overloading the network.
- **Compatibility with Ethereum Clients**: Omni leverages existing Ethereum execution clients (e.g., Geth, Besu, Erigon), ensuring stability and up-to-date features.
- **Dynamic Fee Mechanism**: Supporting [EIP-1559](https://eips.ethereum.org/EIPS/eip-1559), the Omni EVM allows for dynamic transaction fees and partial fee burning, optimizing network usage costs.

## Unifying Consensus and Execution

Omni's dual-chain architecture facilitates **Integrated Consensus**, allowing validators to simultaneously run consensus for the Omni EVM and cross-network messages. This innovative design, supported by tools like ABCI++ and the Engine API, makes Omni's sub-second finality possible. Validators efficiently attest to the state of external rollup VMs, ensuring seamless state transitions and unified state management.

By separating consensus operations from transaction execution, Omni effectively scales activity across its network and connected rollups. This architecture not only mitigates the risk of network congestion but also enhances the security and reliability of cross-network transactions.
