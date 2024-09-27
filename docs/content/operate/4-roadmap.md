# Operator Roadmap

## ✅ Private Developer Omega Testnet

- This testnet is currently live.
- You can join the testnet as a full node.
- Instructions for joining as a full node can be found in our documentation.
- Please note that the network will be reset on occasion, so you will need to resync any nodes that you set up.

## ✅ Public Omega Testnet

- Omega testnet is now live.
- Initially Omega will launch with just validators run by the Omni team.

## ⏳ Public Omega Testnet + Whitelisted Validators

- You will be able to run a validator, archive node, or seed node.
- Validators must be whitelisted by team. Note that the AVS whitelist is not the same as the validator whitelist.
- If you will be running a validator on mainnet, we highly recommend doing it on Omega as well.
- To run a validator, you will need access to an RPC endpoint (full node) for all supported chains.
- Initial supported chains include: Ethereum Holesky, Arbitrum Sepolia, Optimism Sepolia, and Base Sepolia. Note that you’ll also need to run Ethereum Sepolia for the Arbitrum, Optimism and Base Sepolia full nodes.
- We recommend running a full node for supported chains, even for testnet, to ensure the stability of infrastructure in preparation for mainnet.
- To simulate mainnet, we will also be running a network upgrade during public testnet. We will need to coordinate this upgrade with you.

#### **Action items**

- Run a validator (see [instructions](./2-validator.md))
- Run a network upgrade
  - We will be reaching out to coordinate this once all operators are validating on Omega.

## 🗺️ Mainnet - Beta

- Pending audit completion, our mainnet v1 release will be the same as Omega.
- You will need to stake a minimum amount of 100 $OMNI tokens to be able to register your validator on mainnet.
- Slashing will be enabled for double signing. You may be jailed for inactivity, but can unjail your validator if you come back online.
- Staking rewards, withdrawals, delegations, and ETH restaking will NOT be enabled.
- You will need to run full nodes for: Ethereum, Arbitrum, Optimism, and Base (if you don’t already). Your Omni validator will need access to RPC endpoints for those chains.

## 🗺️ Mainnet - Staking Network Upgrades

After Mainnet Beta is live and stable, we will turn our sites to several staking features:

- Rewards: Enable staking rewards so delegators and validators accrue rewards.
- Withdrawals: withdraw your stake and leave the validator set.
- Delegations: receive native $OMNI delegations from the Omni foundation and other users.
- X-Chain attestations: Enable rewards and penalties.
- ETH restaking: receive $ETH delegations from users via Eigenlayer

We plan to batch release these in 3-4 network upgrades.

💡 Please note that the current Omni AVS contract is deployed to mainnet, but will require an upgrade in order to support separation of validator & operator keys (in addition to a few other updrades). This will require you to re-register your operator.

## 🗺️ Mainnet++

- After launching each of these phases, we’ll be removing the validator whitelist.
- The **top n*** validators of the registered set will be included in the active validator set.
- The precise formula for determining the “top n” will be released with this upgrade. For most blockchains, the formula is simply the n validators with the most native tokens staked/delegated to them. However, because Omni validators can stake/receive delegation in both $OMNI and $ETH, the formula used to compute validator power is slightly more complex and will depend on several factors like the amount of economic security currently derived from each asset, the desired ratio, and more.
