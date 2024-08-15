---
sidebar_position: 2
---

# Node Types

## Validators

- stake **\$OMNI** and **\$ETH**, and receive delegations for cryptoeconomic security
- build, propose, and attest to new blocks
- attest to cross chain blocks (`XBlocks`)
- receive **\$OMNI** rewards for performing their duties
- receive penalties (slashes) if they do not perform their duties

## Full Nodes

- do block by block validation of the network
- verify block body and state data for each block (not cross chain block data)
- store the last `n` chain states
- serve network data to other nodes
- can serve RPC endpoints for users, developers, and services

## Archive Nodes

- full nodes that never delete any downloaded data
- store chain state for every historical block

## Seed Nodes

- maintain an address book of peers on the p2p network
- help new nodes bootstrap p2p connections
