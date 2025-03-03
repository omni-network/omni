# Operator Roadmap

## ✅ Private Developer Omega Testnet

## ✅ Public Omega Testnet

## ✅ Public Omega Testnet + Whitelisted Validators

## ✅ Mainnet - Beta

## ✅ Mainnet - Magellan - Staking

See https://docs.omni.network/operate/magellan

## ⏳ Mainnet - Drake - Staking withdrawals

- Withdrawals: withdraw your stake and leave the validator set.

## 🗺️ Mainnet++

- X-Chain attestations: Enable rewards and penalties.
- ETH restaking: receive $ETH delegations from users via Eigenlayer
- After launching each of these phases, we’ll be removing the validator whitelist.
- The **top n*** validators of the registered set will be included in the active validator set.
- The precise formula for determining the “top n” will be released with this upgrade. For most blockchains, the formula is simply the n validators with the most native tokens staked/delegated to them. However, because Omni validators can stake/receive delegation in both OMNI and ETH, the formula used to compute validator power is slightly more complex and will depend on several factors like the amount of economic security currently derived from each asset, the desired ratio, and more.


💡 Please note that the current Omni AVS contract is deployed to mainnet, but will require an upgrade in order to support separation of validator & operator keys (in addition to a few other upgrades). This will require you to re-register your operator.
