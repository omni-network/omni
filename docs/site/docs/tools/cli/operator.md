---
sidebar_position: 3
---

# Operator Commands

## Start a full node

```bash
omni operator init-nodes --network=omega --moniker=foo --clean
cd ~/.omni/omega
docker compose up
```

## Registering as an Operator

Registers EigenLayer operator to Omni using the `operator.yml` file.

```bash
omni operator register --config-file ~/path/to/operator.yaml
```

Performs a contract call to the RPC URL found in the configuration file to register the operator (using the address variables in the file).
