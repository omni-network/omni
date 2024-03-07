---
sidebar_position: 2
---

# Engine API

The Engine API allows the Omni execution layer (Omni EVM) to mirror the functionality and design of Ethereum L1’s execution layer (EVM). Users send transactions to the Omni EVM mempool and execution clients share those transactions through the execution layer’s peer-to-peer (P2P) node network. For each block of transactions, a node’s execution client computes the state transition function for the Omni EVM and shares the output with its `halo` client through the Engine API. Since the transaction mempool resides on the execution layer, Omni can scale activity without congesting the CometBFT mempool used by validators on the consensus layer. Previously, using the CometBFT mempool to handle transaction requests caused the network to become overloaded and resulted in liveness disruptions. Similar challenges have been observed in other projects that adopted a comparable approach.
