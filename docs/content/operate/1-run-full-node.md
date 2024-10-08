# Run a Full Node

Currently, anyone can run a node on Omega Testnet. Stay tuned for running full nodes on mainnet.

## Omni Omega Testnet

Please note that if you're running an Omega full node, you will need to redeploy it several times, as the network is being redeployed with new features frequently.

### Quick Start

The simplest way to run a full node is with the following commands:

```bash
# Install the Omni CLI
curl -sSfL https://raw.githubusercontent.com/omni-network/omni/main/scripts/install_omni_cli.sh | sh -s

# init geth and halo
omni operator init-nodes --network=omega --moniker=foo --clean

# start geth and helo
cd ~/.omni/omega
docker compose up
```

Congrats, you're running a full node!

### Details

What's actually happening here?

- First, you're installing the `omni` cli. We've packaged up several flows into this CLI to make running a node easier for operators.
- The `init-nodes` command is used to generate genesis files and config files and docker compose in the `~/.omni/<network>` directory.
- `docker compose up -d` spins up docker containers `halovisor` and `geth`.

Note that this is the preferred way to run an omni node and results in a production quality deployment (as long as docker is started on startup of the machine).

### Node Requirements

| Category         | Recommendation                                                               |
|------------------|------------------------------------------------------------------------------|
| Cores            | 4                                                                            |
| Bandwidth        | 100 Mbps                                                                     |
| RAM              | 16GB                                                                         |
| SSD Hard Disk    | 500 GB                                                                       |
| Docker           | 24.0.7                                                                       |
| Operating System | `linux/amd64`                                                                |
| Inbound ports    | Enabled for cometBFT (`tcp://26656`) and Geth (`tcp://30303`, `udp://30303`) |

### `halo` Deployment Instructions

Note that `halo` is a CosmosSDK application which requires a specific binary version to run at each network upgrade height.
CosmosSDK uses [Cosmovisor](https://docs.cosmos.network/main/build/tooling/cosmovisor) to manage the binary versioning and swapping at the correct height.

There are basically three ways to run `halo`:

1. **🥇`omniops/halovisor:<latest>` docker container**
    - Simply run the latest version of `halovisor` docker container. It will automatically detect and run the correct halo binary version using cosmovisor internally.
    - E.g. `omniops/halovisor:v0.9.0` contains the `halo:v0.8.1` and `halo:v0.9.0` binaries and will automatically switch at the correct height.
    - It only requires a single docker volume mount: `-v ~/.omni/<network>/halo:/halo`
    - It will persist the cosmovisor “current” binary symlink to: `halo/halovisor-current`
2. **🥈Standard Cosmovisor with `halo` binaries**
    - Install and configure stock-standard CosmosSDK Cosmovisor with `halo` binaries, see docs [here](https://docs.cosmos.network/main/build/tooling/cosmovisor#setup) and [here](https://docs.archway.io/validators/running-a-node/cosmovisor) and [here](https://docs.junonetwork.io/validators/setting-up-cosmovisor). This will also automatically swap the “current” binary at the correct height.
    - The binaries versions to use are:
        - `genesis: halo:v0.8.1`
        - `upgrade/1_uluwatu: halo:v0.9.0`
    - Suggested env vars:
        - `ENV DAEMON_ALLOW_DOWNLOAD_BINARIES=false`
        - `ENV DAEMON_RESTART_AFTER_UPGRADE=true`
        - `ENV UNSAFE_SKIP_BACKUP=true`
    - The folder structure should be:
    ```bash
    ~/.omni/<network>/halo # $DEAMONHOME
      ├─ data/...
      ├─ config/...
      ├─ cosmovisor/
      │ ├── genesis/bin/$DEAMONNAME # halo:v0.8.1
      │ ├── upgrades/1_uluwatu/bin/$DEAMONNAME # halo:v0.9.0
    ```
3. **🥉Manual binary/docker swapping**
    - Swapping halo binary version manually is also an option.
    - When `halo:v0.8.1` reaches the network upgrade height, it will stall.
    - Stop it, and replace it with `halo:v0.9.0`
    - Start the node and it should catch up and continue processing the chain.
    - Note this will include downtime and is therefore not advised for validators as will negatively impact validator performance.

See the [Operator FAQ](./5-faq.md)  for details on `halovisor vs halo` and `docker vs binaries`
