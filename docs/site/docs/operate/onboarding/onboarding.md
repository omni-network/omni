---
sidebar_position: 1
---

# Becoming an Omni Operator

## Running a full node

Currently, anyone can run a node on Omni Omega Testnet.

The simplest way to run a full node is with the following commands:

```bash
# Install the Omni CLI (alternate instructions here: https://docs.omni.network/tools/cli/)
curl -sSfL https://raw.githubusercontent.com/omni-network/omni/main/scripts/install_omni_cli.sh | sh -s

# init geth and halo
omni operator init-nodes --network=omega --moniker=foo --clean

# start geth and helo
cd ~/.omni/omega
docker compose up
```

Congrats, you're running a full node!

## AVS Registration

### Registering as an Operator

Operators can currently register with Omni's Eigenlayer AVS contract. Please note that **\$ETH** delegation will only be enable in the mainnet staking upgrade alongside other features like **\$OMNI** delegation, validator withdrawals, and more.

You can follow Eigenlayer's instructions [here](https://docs.eigenlayer.xyz/eigenlayer/operator-guides/operator-installation) to register your operator key with their contracts.

Then you can run the following command to register within the Omni AVS contract:

```bash
omni operator register --config-file ~/path/to/operator.yaml
```
