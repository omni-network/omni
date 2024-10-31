# Operator Roadmap

## ✅ Private Developer Omega Testnet

## ✅ Public Omega Testnet

## ✅ Public Omega Testnet + Whitelisted Validators

- Operators have been onboarded on Omega testnet. See [public dashbboard](https://omniops.grafana.net/d/ddycpqfje3pxcb/validator-dash-public)
- Operators have applied the [Uluwatu](uluwatu.md) network upgrade.

## ⏳ Mainnet - Beta

- Whitelisted operators will be asked to validate mainnet once ready.
- You will need to stake a minimum amount of 100 OMNI tokens initially.
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
- The precise formula for determining the “top n” will be released with this upgrade. For most blockchains, the formula is simply the n validators with the most native tokens staked/delegated to them. However, because Omni validators can stake/receive delegation in both OMNI and ETH, the formula used to compute validator power is slightly more complex and will depend on several factors like the amount of economic security currently derived from each asset, the desired ratio, and more.
