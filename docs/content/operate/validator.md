---
title: Run a Validator
---

# Run a Validator

This guide describes the process to configure and register a validator on **mainnet** and **Omega** testnet.

## Pre-requisites

- Synced Omni full node (`halo`+ `geth` ) on the [latest release](https://github.com/omni-network/omni/releases/latest). Otherwise follow the [Run a Full Node](./run-full-node.md) guide first.
- Ethereum L1 and multiple L2 RPC endpoints, required for cross-chain validation duties. See [Mainnet](../build/mainnet.md) and [Omega Testnet](../build/omega.md) for more details.

## Summary

This guide will take you through the following steps:

1. **Update validator config** for `halo` and `geth`.
1. **Obtain the halo consensus public key** by running a docker compose command.
1. **Generate an operator address**, or use an existing Ethereum EOA account.
1. **Fund the operator address** with `native OMNI` (chain id [166](https://chainlist.org/chain/166) for `mainnet`) and [164](https://chainlist.org/chain/164) for `Omega`.
1. **Add the operator address** to the omni staking allow list.
1. **Run the`omni operator create-validator`**  CLI command to register the validator.

## Instructions

:::info
The config below is for the **Omega** testnet network. Adjust accordingly for the **mainnet** network.
:::

### 1. **Update validator config** for `halo` and `geth`.

#### 1.1 Configure RPC endpoints for supported chains
Halo nodes running as validators require RPC endpoints for L1 and L2 cross chain validation duties.
This must to be configured in under `[xchain.evm-rpc-endpoints]` in the halo config file `~/.omni/omega/halo/config/halo.toml`.
The chains currently required are: `arb_sepolia, base_sepolia, holesky, op_sepolia`.

```
#######################################################################
###                             X-Chain                             ###
#######################################################################

[xchain]

# Cross-chain EVM RPC endpoints to use for voting; only required for validators. One per supported EVM is required.
# It is strongly advised to operate fullnodes for each chain and NOT to use free public RPCs.
[xchain.evm-rpc-endpoints]

arb_sepolia = "http://my.arbitrum-sepolia.node:8545"
base_sepolia = "http://my.base-sepolia.node:8545"
holesky = "http://my.ethereum-holesky.node:8545"
op_sepolia = "http://my.op-sepolia.node:8545"
```

> Note that a halo node that is a validator will crash if all xchain RPC endpoints are not configured.
> Normal full nodes do not however need to connect to xchain RPCs, so this config is ignored by full nodes.

#### 1.2 Synchronize block building config
Halo the consensus client must be aligned with geth the execution client in terms of block building timing configuration.
Execution blocks must be built slightly faster (500ms) than halo waits before fetching them (600ms).

Ensure the following two fields are configured:
- `halo/config/halo.toml`: `evm-build-delay = "600ms"`
- `geth/config.toml`: `Eth.Miner.Recommit = 500000000` # 500ms

If `halo` or `geth` is already running, restart them to pickup the new config:
```
❯ docker compose restart
```

#### 1.3 Consensus timeouts
Omni aims to achieve fast 1s block times.
All validator MUST ensure that CometBFT is configured with a 1s commit timeout.
This must to be configured in under `[consensus.timeout_commit]` in the CometBFT config file `~/.omni/omega/halo/config/config.toml`.

```
# How long we wait after committing a block, before starting on the new
# height (this gives us a chance to receive some more precommits, even
# though we already have +2/3).
timeout_commit = "1s"
```


### 2. **Obtain the halo consensus public key**

> The “halo consensus key” (also known in Cosmos chains as the “[Tendermint/CometBFT consensus key](https://tutorials.cosmos.network/tutorials/9-path-to-prod/3-keys.html#what-validator-keys)”) is used to sign CometBFT blocks on an ongoing basis.
>

The `omni operators init-nodes` command generates a consensus private key in `~/.omni/omega/halo/config/priv_validator_key.json` .

When registering a validator, the associated public key must be provided which will enable this node as a validator.

The public key can be obtained via either of the following commands

```bash
# docker compose command in `~/.omni/omega` if `omni init-nodes` was used
❯ docker compose run halo run consensus-pubkey

# or

# bash commands on private key itself
❯ cat ~/.omni/omega/halo/config/priv_validator_key.json | jq -r .pub_key.value | base64 -d | xxd -ps -c33
```

The `docker compose` command outputs the following:

```jsx
INF running app args=["consensus-pubkey"] module=cosmovisor path=/halo/cosmovisor/genesis/bin/halo
Consensus public key: 02e47138b658317e8a9ce3fd59c4c41ede153cf2051de3bf9926bd6cfe839512f5
```

> Note the `omni operator create-consesus-key` CLI command can also be used to generate a new consensus key and state files.
>

**Remember to backup this consensus private key** if you haven’t already.

### 3. **Generate an operator address**

> The “operator address” (also known in Cosmos chains as the “v[alidator address](https://hub.cosmos.network/main/validators/validator-faq.html)” or “[validator operator application key](https://tutorials.cosmos.network/tutorials/9-path-to-prod/3-keys.html#what-validator-keys)” is a normal Ethereum EOA address that is used to publicly identify your validator. The private key associated with this address is used to register the validator, delegate, unbond and claim rewards in native `OMNI` on the Omni Omega EVM .
>

At this point, it is only possible to register a validator using the `omni` CLI which only supports insecure hex-encoded private key files.

To generate such a operator private key, run the following `omni` CLI command:

`$ omni operator create-operator-key`

This creates the `./operator-private-key-{ADDRESS}` file containing the hex-encoded private key. Note that the filename is suffixed with the operator address for easy identification.

```
❯ omni operator create-operator-key
🎉 Created operator private key                                                            type=insecure file=./operator-private-key-0x6e9C5F0Ad4739C746f4398faAf773A3503476b90 address=**0x6e9C5F0Ad4739C746f4398faAf773A3503476b90**
🚧 Remember to backup this key 🚧
```

‼️ **Remember to backup this operator private key** taking note of the associated **operator address**

### 4. Fund the operator address

When registering a validator, the operator address must self-delegate `100 OMNI` on the [Omni Omega](https://chainlist.org/chain/164) chain. It therefore needs to funded with slightly more than `100 OMNI` to pay for gas.

The [omega faucet](https://faucet.omni.network/) only provides a few OMNI at a time, so please reach out to the Omni team on Slack, providing your **operator address** and we will fund it with `200 OMNI`.

### 5. Add the operator address to Omni staking allow list

All staking actions (like validator registration) are submitted to the [omni staking predeploy contract](https://omega.omniscan.network/address/0xcccccc0000000000000000000000000000000001) on Omni Omega. The `CreateValidator` function has an allow list that prevents unexpected validator registrations. Only the omni team’s admin account can modify the allow list.

Send your **operator address** to the Omni team via Slack. We will add add it to the allow list and confirm when it has been added.

### 6. **Run `omni operator create-validator` CLI command**

The final step is to submit a `CreateValidator` transaction from the **operator address** to Omni Omega that registers a new validator on the Omni Omega Consensus chain.

Run the `omni operator create-validator` command, and pass in:

- The consensus public key
- The name of the network: `omega`
- The path of the operator private key file
- `100 OMNI` self-delegation amount

```bash
❯ omni operator create-validator \
  --consensus-pubkey-hex=02e47138b658317e8a9ce3fd59c4c41ede153cf2051de3bf9926bd6cfe839512f5 \
  --network=omega \
  --private-key-file=./operator-private-key-0x6e9C5F0Ad4739C746f4398faAf773A3503476b90 \
  --self-delegation 100
```

You should see an output similar to:

```jsx
🎉 Create-validator transaction sent and included on-chain
      link=https://omega.omniscan.network/tx/0x6c0643cae8e56772ee83dcec1aa9d06958529f0e6d3bd4ed55e509406db2e2ee
      block=235054
```

After about 10-20 blocks, the omni full node will automatically detect that it has been registered as a validator and start performing validator duties.

🎉You’re all done. Watch the Halo logs to ensure things are running smoothly.

## Troubleshooting

#### Halo crashes with an error:  `"flag --xchain-evm-rpc-endpoints empty/missing chain so cannot perform xchain voting duties"`

This means `halo.toml` `xchain.evm-rpc-endpoints` has not been properly configured with all the required chains and that the full node has detected it has been registered as a validator. As a validator, Halo node requires connectivity to L1 and L2 chains to perform xchain voting duties. If these are not specified halo crashes.

#### When running `omni operator create-validator` you see an error message `create-validator: load private key: invalid character '8' at end of key file`

This may happen if you specify your own operator key, and it's the wrong length or invalid format or ends on a newline. Check you don't have `0x` in front of the key, or create one with `omni operator create-operator-key` to ensure the one you are using is the same format/length.

    ```bash
    ❯ cat ./operator-private-key-0x6e9C5F0Ad4739C746f4398faAf773A3503476b90
    cb5b39fbaa81f55d00285f8014b9a65e24e4127fcd982ead4b91ef2ff2e65d41%
    ```
