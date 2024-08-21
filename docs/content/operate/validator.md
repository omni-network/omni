---
sidebar_position: 1
---

# Run a Validator

## Validator Requirements

### Running full nodes for supported chains

In the Omni Network, **validators must also run full nodes for supported chains**, since validators attest to cross chain blocks (`XBlocks`). Mainnet v1 includes: Ethereum, Arbitrum, Optimism, and Base. If you are interested in becoming a validator, you must have robust infrastructure supporting full nodes for each of these chains.

### Staking

Validators must stake **\$OMNI**. Validators can also opt into receiving **\$ETH** delegations once the Eigenlayer integration is complete.

## Validator Whitelist

Omni currently has a validator whitelist. A future network upgrade will enable permissionless validator registration.

## Mainnet Staking Network Upgrades

There will be several network upgrades to enable various validator / staking features. Some of these features include:

- Withdrawals: similar to the Beacon Chain launch, validators will not initially be able to withdrawal their $OMNI stake.
- Delegations: validators can receive **\$OMNI** delegations.
- **\$ETH** Restaking: validators can opt into receiving restaked **\$ETH** delegations, pending Eigenlayer slashing.
- Permissionless validator registration: anyone can register, and collect delegations to be included in the active set.

## AVS Registration

Operators can currently register with Omni's Eigenlayer AVS contract. Please note that **\$ETH** delegation counting for validator power will only be enabled in the mainnet staking upgrades listed above, and is not available in mainnet v1. The full Eigenlayer integration is pending **\$ETH** slashing being available as a feature in Eigenlayer (if **\$ETH** isn't slashable, it can't count for economic security yet).

You can follow Eigenlayer's instructions [here](https://docs.eigenlayer.xyz/eigenlayer/operator-guides/operator-installation) to register your operator key with their contracts.

Then you can run the following command to register within the Omni AVS contract:

```bash
omni operator register --config-file ~/path/to/operator.yaml
```

## Node Requirements

| Category | Recommendation |
| --- | --- |
| Cores | 4 |
| Bandwidth | 100 Mbps |
| RAM | 16GB |
| SSD Hard Disk | 500 GB |
| Docker | 24.0.7 |
| Operating System | Linux/macOS (arm/64) |

Inbound ports will be enabled for cometBFT (tcp://266567) and Geth (tcp://30303, udp://30303)
