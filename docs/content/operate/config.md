# Configure a Full Node

This document provides a detailed step-by-step process describing how to manually setup and configure an Omni node.
It reproduces what the `omni operators init-nodes` command does automatically.
See [Run a Full Node](run-full-node.md) docs for the automated process.

The resulting `~/.omni/<network>` folder has the following structure:
```
~/.omni/<network>/     # Node home folder
 ├─ halo/              # Consensus client
 ├─ geth/              # Execution client
 ├─ compose.yaml        # Docker compose
```

It is a valid configuration for running a production Omni full node.

The supported `<network>` are:
- `omega` testnet, chain id [164](https://chainlist.org/chain/164)
- `mainnet` (experimental), chain id [166](https://chainlist.org/chain/166)

Note that validator nodes require additional config as specified in the [Run a Validator](validator.md).

## Configure halo

Halo is a CosmosSDK application. Configuring it is therefore very similar to any
CosmosSDK application, learn more from their [Running a Node](https://docs.cosmos.network/main/user/run-node/run-node).

The `~/.omni/<network>/halo` folder has the following structure:
```
~/.omni/<network>/halo/
 ├─ config/
 │ ├── halo.toml                # Halo configuration file, similar to Cosmos app.toml
 │ ├── config.toml              # Standard CometBFT configuration file
 │ ├── genesis.json             # The <network> consensus chain genesis file
 │ ├── node_key.json            # Standard CometBFT P2P node identity key
 │ ├── priv_validator_key.json  # Standard CometBFT consensus private key
 ├─ data/
 │ ├─ priv_validator_state.json # Standard CometBFT slashing DB (validator last signed state)
 │ ├─ voter_state.json          # Omni xchain slashing DB (last signed state)
```

### halo.toml
`halo.toml` is Halo's CosmosSDK application configuration file.
It is similar to the `app.toml` file in other CosmosSDK applications.
It is located in `~/.omni/<network>/halo/config/halo.toml`.

Download the latest default `halo.toml` [here](https://github.com/omni-network/omni/blob/main/halo/config/testdata/default_halo.toml).

The following required fields are empty by default and must be populated:
- `network`: Omni network to participate in; `mainnet` or `omega`
- `engine-endpoint`: Omni execution client Engine API http endpoint
- `engine-jwt-file`: Omni execution client JWT file used for authentication. See details [here](https://geth.ethereum.org/docs/faq#what-is-jwtsecret).

### config.toml
`config.toml` is the CometBFT configuration file. It is located in `~/.omni/<network>/halo/config/config.toml`.
Learn more from CometBFT [docs](https://docs.cometbft.com/v0.38/core/configuration).

Download the latest default omni `config.toml` [here](https://github.com/omni-network/omni/blob/main/halo/cmd/testdata/default_config.toml).

Ensure the following fields are populated:
- `moniker`: A custom human readable name for this node
- `mempool.type = "nop"`: Omni consensus chain doesn't have a mempool
- `proxy_app = ""`: Only built-in ABCI app supported
- `p2p.external_address`: Externally advertised address for incoming P2P connections.
- `p2p.seeds`: Omni consensus seed nodes to connect to, see omega's [here](https://github.com/omni-network/omni/blob/main/lib/netconf/omega/consensus-seeds.txt).
- `p2p.persistent_peers`: Can configure this with same seed nodes above.

### Configure CometBFT State Sync
When syncing a new full node from scratch, state sync can be configured to speed up the process.
Learn more from the CometBFT docs [here](https://docs.cometbft.com/v0.34/core/state-sync).

The Omega Omni nodes serving https://consensus.omega.omni.network/ are configured to create snapshots every 100 blocks.

First, obtain the trusted height and hash for the latest snapshot (100th block):
```bash
# Get the latest height
LATEST=$(curl -s https://consensus.omega.omni.network/commit | jq -r '.result.signed_header.header.height')
echo "LATEST=$LATEST"

# Calculate the snapshot height
let "SNAPSHOT_HEIGHT=$LATEST / 100 * 100"
echo "SNAPSHOT_HEIGHT=$SNAPSHOT_HEIGHT"

# Get the snapshot hash
SNAPSHOT_HASH=$(curl -s https://consensus.omega.omni.network/commit\?height\=$SNAPSHOT_HEIGHT | jq -r '.result.signed_header.commit.block_id.hash')
echo "SNAPSHOT_HASH=$SNAPSHOT_HASH"
```

Then, configure state sync in `~/.omni/<network>/halo/config/config.toml`:
- `statesync.enable = true`: Enable state sync
- `statesync.rpc_servers = "https://consensus.omega.omni.network,https://consensus.omega.omni.network"`: Two trusted RPC servers required. Can duplicate omni consensus RPC.
- `statesync.trust_height`: Set to above `SNAPSHOT_HEIGHT`
- `statesync.trust_hash`: Set to above `SNAPSHOT_HASH`

### genesis.json
`genesis.json` is the Omni consensus chain genesis file for the specific network (omega vs mainnet).
It is located in `~/.omni/<network>/halo/config/genesis.json`.

Download the Omega consensus `genesis.json` [here](https://github.com/omni-network/omni/blob/main/lib/netconf/omega/consensus-genesis.json).

### node_key.json
`node_key.json` is the CometBFT P2P node identity key.
It is located in `~/.omni/<network>/halo/config/node_key.json`.

`halo` will automatically generate this file on startup if it doesn't exist.

It can be generated via the `cometbft` cli, see installation instructions [here](https://docs.cometbft.com/v0.38/guides/install).
```bash
cometbft gen-node-key --home ~/.omni/<network>/halo
```

### priv_validator_key.json and state files
`priv_validator_key.json` is the CometBFT consensus private key.
It is located in `~/.omni/<network>/halo/config/priv_validator_key.json`.

`priv_validator_state.json` and `voter_state.json` are the CometBFT and Omni xchain slashing DBs respectively.
They are located in `~/.omni/<network>/halo/data/`.

> Note that Omni requires secp256k1 consensus private keys, not the default ed25519 keys. The cometBFT cli only generates ed25519 keys so cannot be used.

The `omni` CLI can be used to generate a secp256k1 consensus private key and associated state files:
```bash
omni operator create-consensus-key --home ~/.omni/<network>/halo
```

## Configure geth
Omni uses stock standard geth as the execution client.
Learn more from the [geth docs](https://geth.ethereum.org/docs/fundamentals/config-files).

The `~/.omni/<network>/geth` folder has the following structure:
```
~/.omni/<network>/geth
 ├─ config.toml                # Geth configuration file
 ├─ genesis.json               # The <network> execution chain genesis file
 ├─ geth/
 │ ├── nodekey                 # Geth P2P node identity key
 │ ├── jwtsecret               # Geth JWT secret file for auth RPC
```

### config.toml
`config.toml` is the geth configuration file.
It is located in `~/.omni/<network>/geth/config.toml`.

Download the latest default omni `config.toml` [here](https://github.com/omni-network/omni/blob/main/e2e/app/geth/testdata/default_config.toml).

Ensure the following fields are populated:
- `NetworkId`: Omni network ID; Omega testnet is `164`, mainnet is `165`.
- `SyncMode = "full"`: Only full sync supported, snap sync not supported (yet).
- `Node.P2P.Bootnodes`:Omni execution seed nodes to connect to, see omega's [here](https://github.com/omni-network/omni/blob/main/lib/netconf/omega/execution-seeds.txt).

### genesis.json
`genesis.json` is the Omni execution chain genesis file for the specific network (omega vs mainnet).
It is located in `~/.omni/<network>/geth/genesis.json`.

Download the Omega execution `genesis.json` [here](https://github.com/omni-network/omni/blob/main/lib/netconf/omega/execution-genesis.json).

`geth` **_MUST_** be initialized with this genesis file before starting for the first time, see geth [docs](https://geth.ethereum.org/docs/fundamentals/private-network#initializing-geth-database) to learn more:
```bash
geth --state.scheme=path init --datadir ~/.omni/<network>/geth/ ~/.omni/<network>/geth/genesis.json
```

### nodekey
`nodekey` is the geth P2P node identity key.
It is located in `~/.omni/<network>/geth/geth/nodekey`.

`geth` will automatically generate this file during `geth init` or on startup if it doesn't exist.

### jwtsecret
`jwtsecret` is the geth JWT secret file for auth RPC.
It is located in `~/.omni/<network>/geth/jwtsecret`.

A path to this file (or copy of) must be provided in the `halo.toml` `engine-jwt-file` field.

`geth` will automatically generate this file during `geth init` on startup if it doesn't exist.

## Configure Docker Compose
The preferred way to run an Omni node is via Docker Compose as described in the [Run a Full Node](run-full-node.md#halo-deployment-options) docs.
The docker `compose.yaml` file is located in `~/.omni/<network>/compose.yaml`.

Download the latest template `compose.yaml` [here](https://github.com/omni-network/omni/blob/main/cli/cmd/compose.yaml.tpl).

Ensure the following template fields are replaced:
- `{{.HaloTag}}`: The latest `omniops/halovisor` docker image tag, e.g. `v0.99.0`, see [releases](https://github.com/omni-network/omni/releases)
- `{{.GethTag}}`: The `ethereum/client-go` version supported by the above omni release, see release notes, e.g. `v1.99.0`.
- `{{.GethVerbosity}}`: Geth logging level, `3` is recommended for info level.
- `{{ if .GethArchive }}- --gcmode=archive{{ end }}`: Remove this line if not an archive node.
