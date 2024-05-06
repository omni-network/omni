---
sidebar_position: 2
---

# Developer Commands

## Scaffolding a Templated Project

The `developer new` command is used to scaffold a new project from a template. The command can be called from your new project directory with:

```bash
omni developer new
```

This will scaffold a new Forge project with an [`XGreeter`](../../develop/xapp/example.md#xgreeter-contract) contract
accompanied by simple mocked testing and a multi-chain deployment script.

## Starting a Local Devnet

:::info Docker Required

Docker is required to run the local development network. If you don't have Docker installed, you can download it from the [official Docker website](https://docs.docker.com/get-docker/).

:::

The `devnet start` command is used to start a local development network. The command can be called from your project directory with:

```bash
omni devnet start
```

This will start a local development network with a single Omni validator node and Omni EVM, two (rollup) anvil nodes and a relayer node.

All EVM nodes have the same rich accounts as anvils nodes regularly do.

Configuration files are created in `~/.omni/devnet`.

## Get Local Devnet Information

The `devnet info` command is used to obtain information about the local development network. The command can be called from your project directory with:

```bash
omni devnet info
```

This will display information about the local development network, including the chain ID, RPC URLs, and Portal contract addresses.

## Clean Up Local Devnet

The `devnet clean` command is used to clean up the local development network. The command can be called from your project directory with:

```bash
omni devnet clean
```

This will stop and remove the local development network containers and clean up the configuration files.
