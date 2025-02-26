# Network Upgrade: Magellan

This guide outlines the process for participating in the *critical coordinated* Omni “Magellan” network upgrade (hard fork).
(Omni network upgrades are now named after famous explorers; Ferdinand Magellan led the first expedition to circumnavigate the globe.)

## TL;DR

- Ensure the `omniops/halovisor:v0.13.0` Docker image is running **BEFORE** the upgrade height.
  - `halovisor:v0.13.0` wraps Cosmovisor with:
    - `halo:v0.8.1` for `genesis`
    - `halo:v0.12.0` for `1_uluwatu`
    - `halo:v0.13.0` for `2_magellan`
  - It will automatically switch the binary at the required block.
- Upgrade name (for Cosmovisor): `2_magellan`
- Omega upgrade height: `TBD` (1 PM UTC, 27 Feb 2025)
- Mainnet upgrade height: `TBD` (1 PM UTC, 3 Mar 2025)
- Supported versions before upgrade (`1_uluwatu`): `halo:v0.11.0 .. v0.12.0`
- Required version after upgrade (`2_magellan`): `halo:v0.13.0`

> 🚧 Like any blockchain network upgrade (hard fork), nodes that do not upgrade will crash or stall.

## Details

The `2_magellan` upgrade is the second planned network upgrade (hard fork) for the Omni network and is included in the `halo:v0.13.0` release.

### Key Changes:
- Enables **permissionless delegation** of OMNI to validators, earning 11% staking rewards.
  - *Withdrawals are not yet supported and will be introduced in the next upgrade.*
- **Consensus improvements:**
  - Protobuf encoding of EVM payloads in consensus blocks for improved performance and security.
  - Simplified EVM event processing for enhanced efficiency.
  - Enqueuing of automatic staking withdrawals to EVM (*not yet processed*).
  - Staking event buffer (up to 12h) to prevent validator set thrashing.

### Geth Compatibility:
- No changes are required to `geth`; this version remains compatible with `v1.14.13`.
- The latest `geth` version `v1.15.3` (at time of writing) is **not supported** due to a [regression](https://github.com/ethereum/go-ethereum/issues/31208) affecting DNS bootnode support.

### Additional Resources:
- See [Run a Full Node](./run-full-node.md#halo-deployment-options) for details on running `halo` with Cosmovisor.
- Check the [Operator FAQ](./faq.md) for more details on `halovisor vs halo` and `Docker vs binaries`.
