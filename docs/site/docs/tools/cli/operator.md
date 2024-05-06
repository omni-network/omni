---
sidebar_position: 3
---

# Operator Commands

## Registering as an Operator

Registers EigenLayer operator to Omni using the `operator.yml` file.

```bash
omni operator register --config-file ~/path/to/operator.yaml
```

Performs a contract call to the RPC URL found in the configuration file to register the operator (using the address variables in the file).
