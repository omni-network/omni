---
sidebar_position: 1
title: Run a Validator
---

# Run a Validator on Omni Omega Testnet

This guide describes the process to configure and register an Omni Omega validator.

## Pre-requisites

- Synced Omni full node (`halo`+ `geth` ) on the [latest release](https://github.com/omni-network/omni/releases/latest). Otherwise follow the “[run a full node](https://docs.omni.network/operate/run-full-node)” guide first.
- Ethereum L1 and multiple L2 RPC endpoints for the following chains. This is required for cross-chain validation duties:
    - [Ethereum Holesky](https://chainlist.org/chain/17000)
    - [Optimism Sepolia](https://chainlist.org/chain/11155420)
    - [Arbitrum Sepolia](https://chainlist.org/chain/421614)
    - [Base Sepolia](https://chainlist.org/chain/84532)


## Summary

This guide will take you through the following steps:

1. **Update `halo.toml`** config to include `xchain.evm-rpc-endpoints` for x-chain (cross-chain) votes/attestations (L1 + L2).
2. **Obtain the halo consensus public key** by running a docker compose command.
3. **Generate an operator address**, or use an existing Ethereum EOA account.
4. **Fund the operator address** with `200 OMNI` on the [Omni Omega](https://chainlist.org/chain/164) chain.
5. **Add the operator address** to the omni staking allow list.
6. **Run the`omni operator create-validator`**  CLI command to register the validator.


## Instructions

### 1. **Update `halo.toml`**

Halo nodes running as validators require RPC endpoints for L1 and L2 cross chain validation duties. This must to be configured under `[xchain.evm-rpc-endpoints]` in the halo config file `~/.omni/omega/halo/config/halo.toml`. The chains currently required are: `arb_sepolia, base_sepolia, holesky, op_sepolia`.

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
Normal full nodes do not however need to connect to xchain RPCs, so this config is ignored by full nodes.
>

If `halo` is already running, restart it to pickup the new config:

`$ docker compose restart halo`

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

 

‼️ **Remember to backup this consensus private key** if you haven’t already.

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

‼️ **Remember to backup this operator private key** taking note of the associated **operator address**

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

🎉 You’re all done. Watch the Halo logs to ensure things are running smoothly.

## Troubleshooting

#### Halo crashes with an error:  `"flag --xchain-evm-rpc-endpoints empty/missing chain so cannot perform xchain voting duties"`

This means `halo.toml` `xchain.evm-rpc-endpoints` has not been properly configured with all the required chains and that the full node has detected it has been registered as a validator. As a validator, Halo node requires connectivity to L1 and L2 chains to perform xchain voting duties. If these are not specified halo crashes.

#### When running `omni operator create-validator` you see an error message `create-validator: load private key: invalid character '8' at end of key file`

This may happen if you specify your own operator key, and it's the wrong length or invalid format or ends on a newline. Check you don't have `0x` in front of the key, or create one with `omni operator create-operator-key` to ensure the one you are using is the same format/length.

    ```bash
    ❯ cat ./operator-private-key-0x6e9C5F0Ad4739C746f4398faAf773A3503476b90
    cb5b39fbaa81f55d00285f8014b9a65e24e4127fcd982ead4b91ef2ff2e65d41%
    ```

## FAQ

### What is the Omni software stack?
- The Omni architecture is similar to Ethereum PoS in that it consists of two chains: an execution chain and a consensus chain.
- The execution chain is implemented by running the latest version of `geth` . Note that Omni doesn’t fork geth, we use the stock standard version, just with a custom Omni execution genesis.
- The consensus chain is implemented by running `halo` which is a CosmosSDK application chain. Halo connects to geth via the [EngineAPI](https://geth.ethereum.org/docs/interacting-with-geth/rpc#engine-api).
- Running a Omni full node therefore consists of running both `halo` and `geth`.

### Does Omni provide official docker images?
- Yes, see [omniops/halovisor](https://hub.docker.com/r/omniops/halovisor/tags?page_size=&ordering=&name=latest) and [ethereum/client-go](https://hub.docker.com/r/ethereum/client-go/tags?page_size=&ordering=&name=latest)
- Note that the `omni operator init-nodes` CLI command generates all the required config files, genesis files, keys and a docker compose file required to run `halo` and `geth` using docker compose. It also calls `geth init` with the Omni execution genesis file.

### What is the difference between the [omniops/halovisor](https://hub.docker.com/r/omniops/halovisor/tags?page_size=&ordering=&name=latest) and  [omniops/halo](https://hub.docker.com/r/omniops/halo/tags?page_size=&ordering=&name=latest) docker containers?

- The `omniops/halovisor` container combines [cosmovisor](https://docs.cosmos.network/v0.46/run-node/cosmovisor.html) with all halo binaries required for network upgrades.
- The `omniops/halo` container only contains a specific halo binary.
- It is strongly advised to run the latest `omniops/halovisor` version since this ensures your validator will automatically perform required network upgrades if and when they occur. The omni team will also communicate details of any planned network upgrades.
- Cosmos network upgrades require switching binary versions at the specific chain heights that network upgrades occur. The `halovisor` container handles this automatically.
- Both `omniops/halovisor:latest` and `omniops/halo:latest` point to the latest stable omni release as per the Omni monorepo [Github release page](https://github.com/omni-network/omni/releases).

### Can raw binaries be used instead of docker containers?
- Yes, the halo and geth binaries are available on their respective Github release pages.
- Note that setting up cosmovisor is strongly advised to support smooth network upgrades. See our [halovisor build scripts](https://github.com/omni-network/omni/tree/main/scripts/halovisor) for inspiration.
- Note that before starting geth, it must first be initialised with the Omni Omega [`execution-genesis.json`](https://github.com/omni-network/omni/tree/main/lib/netconf/omega) file via `geth init`.

### What is the XChain RPC request rate per validator?
- Each validator needs to attest to each source chain block twice, once for `latest` confirmation level, and once for `finalized` confirmation level.
- The validator does maximum 4 queries per block; once for the `block header`, once for `xmsg` event logs, once for `xreceipt` event logs plus some additional polling.
- So RPC request rate primarily depends on the block period per chain:
  1. `arb_sepolia` : 4 blocks/s  * 4 req/vote * 2 votes/block ~=  32 req/s
  2. `op_sepolia` : 0.5 blocks/s * … ~= 4 req/s
  3. `base_sepolia` : 0.5 blocks/s * … ~= 4 req/s
  4. `holesky` : 0.08 blocks/s ~= 1 req/s (due to additional polling of slow chains)
- Note that the above are ***average*** rates. ***Bursts*** of much higher rates must be supported since chains finalize blocks in large batches instead of continuously, e.g. +-2000 blocks every 8mins for `arb_sepolia`. The whole batch of finalized blocks need to be voted on as fast as possible, this results in very high query bursts (up to 200 req/s).
- Rate limiting of XChain RPC requests should therefore not be applied for best xchain validator performance.

### How to “unjail” a validator?
- The Cosmos staking module can “jail” a validator for inactivity which removes it from the active validator set. See more details [here](https://docs.cheqd.io/node/validator-guides/validator-guide/unjail).
- Note that Omni validators have two types of duties to perform:
  1. CometBFT consensus blocks (can be jailed for inactivity)
  2. XChain votes (cannot be jailed for inactivity yet).
- To “unjail” a validator, submit the following `unjail` transaction from the **operator address** to the Omni execution chain by using the `omni` CLI similar to `create-validator` above :

    ```bash
    ❯ omni operator unjail \
        --network=omega \
        --private-key-file=./operator-private-key-0x6e9C5F0Ad4739C746f4398faAf773A3503476b90
    ```
### What is the difference between L and F on the dashboard?
- Each validator votes for each support chain twice: once for latest blocks (L) and once for finalized blocks (F). This allows users of our xchain protocol to decide if they want to wait for chain finalization (strong security and exactly once guarantees) or if they want fast xchain messages with latest (strong security but no delivery guarantees due to risk of reorg).

### What are the validation duties of a validator?
- Omni validators have two duty types to perform:
  1. Normal cosmos/cometBFT consensus : UptimeC on dashboard
  2. XChain votes and attestations: UptimeX on dashboard

### How do validators vote?
- Validators vote in the Omni consensus chain, using a new feature of CosmosSDK/CometBFT called [“vote extensions”](https://docs.cosmos.network/main/build/abci/vote-extensions)

### When will ETH delegation be available?

**\$ETH** delegation counting for validator power is not available initially on Mainnet and will only be enabled after a staking upgrade. The full Eigenlayer integration is pending **\$ETH** slashing being available as a feature in Eigenlayer (if **\$ETH** isn't slashable, it can't count for economic security yet).

Operators can still register with Omni's Eigenlayer AVS contract following Eigenlayer's instructions [here](https://docs.eigenlayer.xyz/eigenlayer/operator-guides/operator-installation) to register operator key with their contracts and then run the following command to register within the Omni AVS contract:
```bash
omni operator register --config-file ~/path/to/operator.yaml
```

<aside>
💡 Please note that the current Omni AVS contract is deployed, but will require an upgrade in order to support separation of validator & operator keys. This will require you to re-register your operator.
</aside>

### Which tokens can be staked?

Validators must stake native **\$OMNI**. Validators can also opt into receiving **\$ETH** delegations once the Eigenlayer integration is complete.

### What is the validator whitelist?

Omni currently has a validator whitelist. The whitelist applies to both native **\$OMNI** staking and **\$ETH** staking via the Omni AVS. A future network upgrade will enable permissionless validator registration.

### What are the planned staking upgrades?

There will be several network upgrades to enable various validator / staking features. Some of these features include:

- Withdrawals: similar to the Beacon Chain launch, validators will not initially be able to withdrawal their $OMNI stake.
- Delegations: validators can receive **\$OMNI** delegations.
- X-chain rewards and penalties
- **\$ETH** Restaking: validators can opt into receiving restaked **\$ETH** delegations, pending Eigenlayer slashing.
- Permissionless validator registration: anyone can register, and collect delegations to be included in the active set.
