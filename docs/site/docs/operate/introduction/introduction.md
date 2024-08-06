---
sidebar_position: 1
---

# Introduction

A **node** is any instance of Omni client software that is connected to other instances in the peer to peer network.

Every node must run 2 clients:

- consensus client: `halo`
  - allows the network to agree on the state of cross chain blocks (`XBlocks`)
  - allows the network to agree (reach "consensus") on the state of the Omni execution client
- execution client: `geth`
  - listens to new transactions broadcasted in the network
  - executes transactions and builds blocks
  - keeps a database of execution state
