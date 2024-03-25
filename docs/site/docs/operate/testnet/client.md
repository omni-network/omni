---
sidebar_position: 4
---

# Run the Client

This is the component that participates in network validation. Our consensus clients track delegations and stake from our AVS contracts.

Similar to Ethereum, Omni validators run 2 components: our consensus client, `halo`, and an EVM execution client `geth`, `erigon`, `nethermind`, etc. The EVM execution client does not require any modifications.

## Obtain $OMNI

To run the client, you will need **\$OMNI**. You can obtain **\$OMNI** by reaching out to the team.

## Initialize and Run the Client

:::warning Feature not yet enabled

This feature is not yet enabled. Please do not attempt to run it.

:::

After adding the `halo` binary to your system, you can initialize and run the client with the following command:

```bash
halo init --network testnet
```

This command initializes the setup, creating several configuration files required for running the node. You can then run the client with the following command:

```bash
halo run
```
