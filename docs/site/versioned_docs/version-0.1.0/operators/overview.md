---
sidebar_position: 1
id: overview
---

# Network Connection Details

The Omni Overdrive Testnet currently connects 3 different rollup networks, uniting Ethereum's L2 ecosystem. When building applications that operate across multiple rollups, you will need the following network connection details to deploy smart contracts.

:::note

If you connected to Omni Origins (Testnet V1), please remove that configuration from your wallet before adding Omni Overdrive (Testnet V2).

:::

"Network name" is used in functions in the Omni protocol like `isTxFrom()` and `isTxFromOneOf()`, read more.

## Omni Overdrive

- RPC: https://testnet.omni.network
- Chain ID: 165
- Explorer: https://testnet.explorer.omni.network

## Integrated Rollups

### Arbitrum Goerli

- Chain ID: 421613
- RPC: https://goerli-rollup.arbitrum.io/rpc
- Explorer: https://goerli.arbiscan.io
- Network name: "arbitrum-goerli"
- Portal address: 0xcbbc5Da52ea2728279560Dca8f4ec08d5F829985

### Scroll Sepolia

- Chain ID: 534351
- RPC: https://sepolia-rpc.scroll.io/
- Explorer: https://sepolia-blockscout.scroll.io
- Network name: "scroll-sepolia"
- Portal address: 0xcbbc5Da52ea2728279560Dca8f4ec08d5F829985

### Linea Goerli

- Chain ID: 59140
- RPC: https://rpc.goerli.linea.build
- Explorer: https://explorer.goerli.linea.build
- Network name: "linea-goerli"
- Portal address: 0xcbbc5Da52ea2728279560Dca8f4ec08d5F829985

### Optimism Goerli

- Chain ID: 420
- RPC: https://goerli.optimism.io
- Explorer: https://goerli-optimism.etherscan.io
- Network name: "optimism-goerli"
- Portal address: 0xcbbc5Da52ea2728279560Dca8f4ec08d5F829985
