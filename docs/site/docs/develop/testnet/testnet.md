---
sidebar_position: 1
---

# Omega Testnet

This page will walk through how to scaffold, test, deploy and interact with a Solidity program added to the Omni Omega testnet.

## Scaffold

You can scaffold a new [Foundry](https://github.com/foundry-rs/foundry) template project with `forge` by running following from your new project directory:

```bash
forge init --template https://github.com/omni-network/omni-forge-template.git
```

## Test

You can test your project by running:

```forge
forge test
```

## Deploy

### Have a Funded Account

Before you deploy your project to any network, you will need native tokens for the network you'll deploy to. See [the resources page](./resources.md) for more information on how to get testnet tokens.

### Override Variables

#### `portalAddress`

The value for `portalAddress` should be updated to the address of the deployed Omni portals. This can be found in the relevant addresses page for the Omni Omega testnet which will be shared soon.

### Deployment

You can deploy your project by running:

```bash
forge script script/XGreeter.s.sol --rpc-url <OMNI_TESTNET_RPC_URL> --private-key <YOUR_PRIVATE_KEY> --broadcast
```

You can deploy your project to any of the supported testnet networks found in the [testnet resources page](./resources.md).

## Interact

You can also interact with your deployed contract by using the `cast send` command, or creating a script and invoke it using:

```bash
forge script script/YourNewScript.s.sol --broadcast --rpc-url <OMNI_TESTNET_RPC_URL> --private-key <YOUR_PRIVATE_KEY>
```

### Monitor Transactions

You can monitor Omni transactions by using the Omni Omega Explorer [found in the resources section](./resources.md).
