---
sidebar_position: 2
---

# Run a Full Node

Currently, anyone can run a node on Omega Testnet. Stay tuned for running full nodes on mainnet.

## Omni Omega Testnet

Please note that if you're running an Omega full node, you will need to redeploy it several times, as the network is being redeployed with new features frequently.

### Quick Start

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

<details>
<summary>Known Issue **"validator does not exist"**</summary>

Please note if you see this error that it is a known issue, it is sporadic and resolves itself after a couple tries. If you're interested to follow along the solution (or give it a shot yourself!), you can follow along [here](https://github.com/omni-network/omni/issues/1524).

</details>

### Details

What's actually happening here?

- First, you're installing the `omni` cli. We've packaged up several flows into this CLI to make running a node easier for operators.
- The `init-nodes` command is used to generate genesis files and node initialization configs in your `~/.omni/` directory.
- `docker compose up` spins up a docker container for `halo` and `geth`, and starts your node.

### Node Requirements

| Category | Recommendation |
| --- | --- |
| Cores | 4 |
| Bandwidth | 100 Mbps |
| RAM | 16GB |
| SSD Hard Disk | 500 GB |
| Docker | 24.0.7 |
| Operating System | Linux/macOS (arm/64) |

Inbound ports will be enabled for cometBFT (tcp://266567) and Geth (tcp://30303, udp://30303)
