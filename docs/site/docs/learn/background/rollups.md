---
sidebar_position: 3
---

# Rollups

Rollups are a type of layer 2 solution that aim to enhance the scalability and efficiency of blockchain networks. They do so by processing transactions outside the main blockchain (off-chain) and then posting the transaction data (on-chain) in a compressed form. This method significantly reduces the strain on the network, allowing for faster and cheaper transactions. Rollups are pivotal to the Omni protocol's ability to facilitate seamless cross-rollup communication.

There are mainly two types of rollups that differ in how transaction data is processed and verified: **Optimistic Rollups** and **ZK-Rollups**.

## Optimistic Rollups

Optimistic Rollups assume transactions are valid by default and only perform computation, and hence incur expenses, when a transaction is challenged. They offer a significant increase in scalability by running computation off-chain and posting the transaction data on-chain. Users or "watchers" can challenge a transaction if they believe it to be fraudulent.

[Learn more about Optimistic Rollups](https://ethereum.org/en/developers/docs/scaling/optimistic-rollups/)

## ZK-Rollups

ZK-Rollups, or Zero-Knowledge Rollups, bundle hundreds of transfers into a single transaction. Instead of publishing all data on the chain, they generate a cryptographic proof known as a SNARK (Succinct Non-interactive Argument of Knowledge). This proof confirms the validity of all transactions in the bundle. ZK-Rollups are highly efficient, as they allow the network to confirm the legitimacy of transactions without executing them fully.

[Discover more about ZK-Rollups](https://ethereum.org/en/developers/docs/scaling/zk-rollups/)
