# Operator Roadmap

## ✅ Omega Testnet - 2024

## ✅ Mainnet - Beta - 2024

## ✅ Mainnet - Magellan - Staking - Q1 2025

See [Network Upgrade - Magellan](https://docs.omni.network/operate/magellan)

## ⏳ Mainnet - Drake - Staking withdrawals - Q2 2025

## 🗺️ Mainnet++

- X-Chain attestations: Enable rewards and penalties.
- ETH restaking: receive $ETH delegations from users via Eigenlayer.
- Remove the validator whitelist.
- The **top n*** validators of the registered set will be included in the active validator set.
- The precise formula for determining the “top n” will be released with this upgrade. For most blockchains, the formula is simply the n validators with the most native tokens staked/delegated to them. However, because Omni validators can stake/receive delegation in both OMNI and ETH, the formula used to compute validator power is slightly more complex and will depend on several factors like the amount of economic security currently derived from each asset, the desired ratio, and more.

💡 Please note that the current Omni AVS contract is deployed to mainnet, but will require an upgrade in order to support separation of validator & operator keys (in addition to a few other upgrades). This will require you to re-register your operator.
