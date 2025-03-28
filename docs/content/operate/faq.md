# FAQ

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
- Note that before starting geth, it must first be initialised with the relevant `execution-genesis.json` file via `geth init` (see [omega](https://github.com/omni-network/omni/tree/main/lib/netconf/omega) and [mainnet](https://github.com/omni-network/omni/tree/main/lib/netconf/mainnet)).

### What syncing options are supported?
Omni has a 2 chain architecture, similar to ethereum PoS.
This means that both the execution chain (geth) and the consensus chain (halo) has to be synced.
Each chain however supports different syncing protocols of which not all are supported by Omni yet.

Halo (via CometBFT) supports ["comet fast sync" and "comet state sync"](https://docs.cometbft.com/v0.34/core/state-sync).
Although Geth technically supports ["ethereum full sync" and "ethereum snap sync"](https://geth.ethereum.org/docs/fundamentals/sync-modes), only "full sync" is supported by Omni ("snap sync" is not supported).

The following three options are currently supported:
1. Halo "fast sync" + Geth "full sync"
   - Syncs both chains from genesis.
   - This is the slowest option.

2. Halo "state sync" + Geth "full sync"
   - CometBFT state sync protocol is configured which syncs the consensus chain from a recent snapshot.
   - This provides a "full sync" target to Geth, which then performs a full sync from genesis.
   - This is faster than option 1.
   - See mode details [here](config.md#configure-cometbft-state-sync).

3. Restoring Node Snapshots
   - Omni provides node snapshot (tarballs) which are disk backups of both `halo` and `geth` data directories.
   - These node snapshots are taken daily from Omni's archive nodes and have a larger memory footprint (see [Hardware Requirements](run-full-node.md#hardware-requirements)).
   - This results in the `geth` snapshots utilizing the legacy `hash` state scheme instead of the newer and more memory efficient `path` state scheme.
   - When spinning up a full node pruning of the archive data (intermediate historical states) is to be expected.
   - The nodes will replay blocks from the snapshot to the current block height.
   - This is the fastest option.
   - See more details [here](config.md#restoring-node-snapshots).

### How long does Geth fullsync take?
The time it takes for Geth to do a fullsync (snapsync is not supported yet) depends largely on your disk write throughput.
Syncing speed is directly proportional to write performance.
We observed the following numbers in practice:
- 200 MB/s disk write speed achieves approximately 40 blocks per second.
- 1.5 GB/s disk write speed achieves around 460 blocks per second.

To estimate your full sync time, check the current chain height (see [omniscan.network](https://omniscan.network)) and divide this number by your blocks-per-second rate to get a rough duration in seconds.
You can estimate your blocks-per-second rate by monitoring the Geth logs: the log line starting with `Imported new potential chain segment` contains the `number` parameter denoting the last imported block height.

### How to check if the validator is ready?
A node is considered ready when it is healthy and functions as intended.
This means both the consensus and execution chains are fully synchronized, there are execution and consensus P2P peers connected, and new blocks are being produced.
The node will be deemed unhealthy while it is syncing, if P2P peers drop to zero, or if consensus is stalled.

The node readiness can be checked as follows:
- For Docker deployments, check the `docker ps` output for `healthy` or `unhealthy` status.
- For binary deployments, visit `http://<NODE_IP>:26660/ready`. The JSON response shows multiple system metrics. Any of these metric being unhealthy leads to a `503` status, indicating the node is not ready.

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

Use `--network=mainnet` for mainnet.

### What is the difference between L and F on the dashboard?
- Each validator votes for each support chain twice: once for latest blocks (L) and once for finalized blocks (F). This allows users of our xchain protocol to decide if they want to wait for chain finalization (strong security and exactly once guarantees) or if they want fast xchain messages with latest (strong security but no delivery guarantees due to risk of reorg).

### What are the validation duties of a validator?
- Omni validators have two duty types to perform:
  1. Normal cosmos/cometBFT consensus : UptimeC on dashboard
  2. XChain votes and attestations: UptimeX on dashboard

### When and how is a "rollback" performed?
CosmosSDK chains, including `halo` supports the `rollback` command that rolls back the CosmosSDK and CometBFT state by one height.

A state rollback is performed to recover from an incorrect application state transition, when CometBFT has persisted an incorrect app hash and is thus unable to make progress. Rollback overwrites a state at height `n` with the state at height `n - 1`. `Halo` application state also rolls back to height `n - 1`. No blocks are removed, so upon restarting the transactions in block `n` will be re-executed against the application.

Assuming that the `halo` is deploying with docker-compose and runs inside the `omniops/halovisor` container:
```
docker compose down
docker compose run halo run rollback
docker compose up -d
```
Note pure `halo` command is just `rollback`. But when running via `halovisor` one needs to specify `run rollback`, with the `run` being the Cosmovisor command to execute halo binary commands.

### How do validators vote?
- Validators vote in the Omni consensus chain, using a new feature of CosmosSDK/CometBFT called [“vote extensions”](https://docs.cosmos.network/main/build/abci/vote-extensions)

### When will ETH delegation be available?

**\$ETH** delegation counting for validator power is not available initially on Mainnet and will only be enabled after a staking upgrade. The full Eigenlayer integration is pending **\$ETH** slashing being available as a feature in Eigenlayer (if **\$ETH** isn't slashable, it can't count for economic security yet).

Operators can still register with Omni's Eigenlayer AVS contract following Eigenlayer's instructions [here](https://docs.eigenlayer.xyz/eigenlayer/operator-guides/operator-installation) to register operator key with their contracts and then run the following command to register within the Omni AVS contract:
```bash
omni operator avs-register --config-file ~/path/to/operator.yaml
# previously `omni operator register` (v0.11.0 or earlier)
```

<aside>
💡 Please note that the current Omni AVS contract is deployed, but will require an upgrade in order to support separation of validator & operator keys. This will require you to re-register your operator.
</aside>

### Which tokens can be staked?

Validators must stake native **\$OMNI**. Validators can also opt into receiving **\$ETH** delegations once the Eigenlayer integration is complete.

### What is the validator whitelist?

Omni currently has a validator whitelist. The whitelist applies to both native **\$OMNI** staking and **\$ETH** staking via the Omni AVS. A future network upgrade will enable permissionless validator registration.

### How can I deregister from the old AVS contract?

To deregister from the old AVS contract, you need to build the CLI from source using `main`. Alternatively, you can wait for release `v0.12.0 (commit: ead6061)` of the Omni CLI, which will support AVS deregistration.

```bash
git clone https://github.com/omni-network/omni.git
cd omni
make setup
make install-cli
omni operator avs-deregister --config-file ~/path/to/operator.yaml
```

### What are the planned staking upgrades?

There will be several network upgrades to enable various validator / staking features. Some of these features include:

- Withdrawals: similar to the Beacon Chain launch, validators will not initially be able to withdrawal their $OMNI stake.
- Delegations: validators can receive **\$OMNI** delegations.
- X-chain rewards and penalties
- **\$ETH** Restaking: validators can opt into receiving restaked **\$ETH** delegations, pending Eigenlayer slashing.
- Permissionless validator registration: anyone can register, and collect delegations to be included in the active set.

### Does Halo support HSMs, TKMS or Horcrux?
As mentioned [above](#what-are-the-validation-duties-of-a-validator), Halo validators have two duties to perform. The interface used by external signer services like HSMs, TKMS or Horcrux only supports standard CometBFT signatures, it doesn't support Omni XChain votes and attestations.

Halo therefore only supports simple `priv_validator_key.json` private key files for signing both CometBFT and XChain votes.
