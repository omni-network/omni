# Network Upgrade: Uluwatu

This guide describes the process to participate in the *critical coordinated* Omni Omega  “Uluwatu” network upgrade (hard fork).  (Omni network upgrades are named after iconic surf spots; Uluwatu is in Bali).

## TL;DR

- Simply ensure the `omniops/halovisor:v0.9.0` docker image is running **BEFORE** the upgrade height.
  - `halovisor:v0.9.0`: wraps cosmovisor with `halo:v0.8.0` and `halo:v0.9.0`
  - It will perform the binary switch automatically at the required block.
  - Note that `omniops/halovisor:v0.9.0` will be released in week of 1 Oct 2024
- Omega upgrade height: TBD
- Approximate upgrade date: 7~11 Oct 2024
- Version(s) supported before upgrade: `halo:v0.4.0 .. v0.8.0`
- Version required after upgrade: `halo:v0.9.0`  (not yet released)

> 🚧 Like any blockchain network upgrade (hard fork), nodes that do not upgrade will crash or stall.

## Details

The “uluwatu” upgrade is the first network upgrade (hard fork) planned for the Omni Omega network and is included in the `halo:v0.9.0` release.

The upgrade contains changes to `halo`’s `attest` module logic ensuring that attestations are only deleted when they exit the modified vote window. See [issue](https://github.com/omni-network/omni/issues/1787) and [PR](https://github.com/omni-network/omni/pull/1983) for details.

No changes to `geth` is required, this version is still compatible with `v1.14.8`. (Note that geth `v1.14.9` has been released but hasn’t published docker images, see [issue](https://github.com/ethereum/go-ethereum/issues/30469) for details.)

See the [Validator FAQ](https://www.notion.so/External-Running-a-Validator-on-Omni-Omega-f7e53b98e38b467cba0aad7d640c0556?pvs=21)  for details on `halovisor vs halo` and `docker vs binaries`

## `halo` Deployment Instructions

There are basically three ways to run a `halo`:

1. **🥇`omniops/halovisor:<latest>` docker container**
    - Simply run the latest version of `halovisor` docker container **BEFORE** the upgrade height. It will automatically detect and run the correct halo binary version using cosmovisor internally.
    - E.g. `omniops/halovisor:v0.9.0` contains the `halo:v0.8.0` and `halo:v0.9.0` binaries and will automatically switch at the correct height.
    - It only requires a single docker volume mount: `-v ~/.omni/omega/halo:/halo`
    - It will persist the cosmovisor “current” binary symlink to: `/halo/halovisor-current`
2. **🥈Standard Cosmovisor with `halo` binaries**
    - Install and configure stock-standard CosmosSDK Cosmovisor with `halo` binaries, see docs [here](https://docs.cosmos.network/main/build/tooling/cosmovisor#setup) and [here](https://docs.archway.io/validators/running-a-node/cosmovisor) and [here](https://docs.junonetwork.io/validators/setting-up-cosmovisor). This will also automatically swap the “current” binary at the correct height.
    - The binaries versions to use are:
        - `genesis: halo:v0.8.0`
        - `upgrade/1_uluwatu: halo:v0.9.0`
    - Suggested env vars:
        - `ENV DAEMON_ALLOW_DOWNLOAD_BINARIES=false`
        - `ENV DAEMON_RESTART_AFTER_UPGRADE=true`
    - The folder structure should be:

    ```bash
    ~/.omni/omega/halo # $DEAMONHOME
      ├─ data/...
      ├─ config/...
      ├─ cosmovisor/
      │ ├── genesis/bin/$DEAMONNAME # halo:v0.8.0
      │ ├── upgrades/1_uluwatu/bin/$DEAMONNAME # halo:v0.9.0
    ```
3. **🥉Manual binary/docker swapping**
    - Swapping halo binary version manually is also an option.
    - When `halo:v0.8.0` reaches the network upgrade height, it will stall.
    - Stop it, and replace it with `halo:v0.9.0`
    - Start the node and it should catch up and continue processing the chain.
    - Note this will include downtime and is therefore not advised for validators as will negatively impact validator performance.
