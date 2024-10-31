# Run a Full Node

## Quick Start

The simplest way to run a full node is with the following commands:

```bash
# install the omni cli (or download https://github.com/omni-network/omni/releases/latest)
curl -sSfL https://raw.githubusercontent.com/omni-network/omni/main/scripts/install_omni_cli.sh | bash -s

# init halo and geth
omni operator init-nodes --network=omega --moniker=foo

# start halo and geth
cd ~/.omni/omega
docker compose up
```

Congrats, you're running a full node!

For the upcoming mainnet, replace the `omega` network with `mainnet`.

## Details

### What's actually happening here?
- First, you're installing the `omni` CLI which contains tooling to manage a node.
- The `omni operator init-nodes` command generates config files, genesis files, and docker compose in `~/.omni/<network>`.
- `docker compose up -d` spins up the `halovisor` and `geth` containers.

### What is the Omni Node software stack?
- The Omni architecture is similar to Ethereum PoS in that it consists of two chains: an execution chain and a consensus chain.
- The execution chain is implemented by running the latest version of `geth` . Note that Omni doesn’t fork geth, we use the stock standard version, just with a custom Omni execution genesis file.
- The consensus chain is implemented by running `halo` which is a CosmosSDK application chain. Halo connects to geth via the [EngineAPI](https://geth.ethereum.org/docs/interacting-with-geth/rpc#engine-api).
- Running an Omni full node therefore consists of running both `halo` and `geth`.
- For step-by-step instructions to manually configuring a full node, see [Configure a Full Node](config.md)

### Hardware Requirements

| Category         | Recommendation                                                               |
|------------------|------------------------------------------------------------------------------|
| Cores            | 4                                                                            |
| Bandwidth        | 100 Mbps                                                                     |
| RAM              | 16GB                                                                         |
| SSD Hard Disk    | 500 GB                                                                       |
| Docker           | 24.0.7                                                                       |
| Operating System | `linux/amd64`                                                                |
| Inbound ports    | Enabled for cometBFT (`tcp://26656`) and Geth (`tcp://30303`, `udp://30303`) |

### `halo` Deployment Options

Note that `halo` is a CosmosSDK application which requires a specific binary version to run at each network upgrade height.
CosmosSDK uses [Cosmovisor](https://docs.cosmos.network/main/build/tooling/cosmovisor) to manage the binary versioning and swapping at the correct height.

There are three ways to run `halo`, listed in order of preference:

1. **🥇 Halovisor docker container**
    - Simply run the `omniops/halovisor:<latest>` docker container.
    - It combines multiple `halo` versions with `cosmovisor` for automatic network upgrades.
    - E.g. `omniops/halovisor:v0.9.0` contains the `halo:v0.8.1` and `halo:v0.9.0` binaries and will automatically switch at the correct height.
    - It only requires a single docker volume mount: `-v ~/.omni/<network>/halo:/halo`
    - It will persist the cosmovisor “current” binary symlink to: `halo/halovisor-current`
    - It will persist the cosmovisor “current” upgrade info to: `halo/halovisor-upgradeinfo.json`

2. **🥈 Cosmovisor with halo binaries**
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

3. **🥉 Manual binary/docker swapping**
    - Swapping halo binary or docker version manually is also an option.
    - When `halo:v0.8.1` reaches the network upgrade height, it will stall.
    - Stop it, and replace it with `halo:v0.9.0`
    - Start the node and it should catch up and continue processing the chain.
    - Note this will include downtime and is therefore not advised for validators as will negatively impact validator performance.

See the [Operator FAQ](./faq.md)  for details on `halovisor vs halo` and `docker vs binaries`
