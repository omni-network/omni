---
sidebar_position: 2
---

# Tokenomics & Utility

This section covers the supply & distribution of **\$OMNI**'s tokenomics.

## Supply

### Validator Rewards and Inflation

Incentivizing Omni validators is crucial to the safety of the network. Initial validator rewards will come from the Ecosystem Development category (29.5% of supply) and are used to reward validators for staking restaked **\$ETH** or staking **\$OMNI**. The proportion of rewards distributed for restaked **\$ETH** and staked **\$OMNI** will be dynamic, allowing the network to incentivize more or less of a given asset over time. For example, if the network requires more restaked $ETH, then restaked **\$ETH** would have a higher reward rate than staked **\$OMNI** and vice versa.

Currently, EigenLayer does not support payments to operators (Omni validators) for their restaked assets. It also does not yet support slashing for these restaked assets. Due to these limitations, the Omni network cannot safely support restaked **\$ETH** as part of its security system. Support will be added when these conditions are met.

Today, **\$OMNI** stakers are receiving 6% APR in rewards for helping to bootstrap the network’s security. When Omni mainnet v1 is live, the Omni Foundation will release a new validator rewards schedule with a fixed amount of **\$OMNI** distributed each month. These rewards will be split dynamically between **\$OMNI** stakers and **\$ETH** restakers depending on market conditions as described above. When governance is fully decentralized, the community will be able to modify this rewards schedule and the parameters that determine the dynamic rewards split.

Initial estimations project the allocated validator rewards from the Ecosystem Development category to be used for 3 years after Omni mainnet v1. Validators will always need to be incentivized with newly issued rewards and governance will need to determine this perpetual  issuance rate. However, this does not mean **\$OMNI** will be inflationary — if the burn rate outweighs the rewards issued to validators, the system will become net-deflationary.

## Utility

**\$OMNI** can be used as the gas payment mechanism for all Omni-related transactions or as a Proof-of-Stake security mechanism. Tokens used for transaction payments are burned. Staking or delegating these tokens earns staking rewards and governance rights.

### Gas Payments

**\$OMNI** functions as the gas token for two types of transactions: cross-rollup transactions and Omni EVM transactions.

1. **Cross-Rollup Transactions**

    Cross-rollup transactions (`XMsgs`) can be transmitted between any Ethereum rollups. The user pays for source network gas, destination network gas, and relayer fees when a cross-rollup transaction is initiated on a source network. In V1 of the protocol, users will only be able to pay fees using the native gas asset from the source rollup network. In V2, users will be able to pay fees using any token.

    In V1, the Omni Foundation will operate the relayer and collect fees as two independent processes. In V2, Omni will implement an in-protocol mechanism for payments to relayers, transforming **\$OMNI** into a gas abstraction primitive. Under this new model, relayers bid for the right to relay messages in **\$OMNI**. The winning relayer receives the fee paid by the user on the source network while the **\$OMNI** paid by the relayer is burned. The winning bid amount will be equal to the transaction’s gas cost on the destination rollup network + the relayer’s service fee.

2. **Omni EVM Transactions**

    Similar to other Layer 1 blockchains, **\$OMNI** is the native gas token that powers all transactions on the Omni EVM. The Octane standard allows the Omni EVM to natively inherit all of the EVM’s historical upgrades. This includes support for EIP-1559 which introduces deflationary burn mechanics into the system. Like Ethereum, Omni EVM transaction fees are split into base fees and priority fees. Base fees are set at the network level and are burned, creating deflationary pressure for the token. Priority fees are set by the user and are given to validators. Priority fees open the door for an MEV marketplace to develop similar to Ethereum and other smart contract networks.

### Staking

In addition to its role as a gas asset, **\$OMNI** is also used to generate economic security in the network’s Proof-of-Stake security model. Since this security model is a dual staking model, the total cryptoeconomic security of the network is equal to the value of staked **\$OMNI** + restaked **\$ETH**. Staked tokens can be slashed if a validator misbehaves or is offline for an extended period of time, resulting in all slashed tokens being burned.

Validators that stake tokens and users that delegate tokens are eligible to receive staking rewards (currently 6% APR) and airdrops given to the Omni Foundation. Additionally, when the network decentralizes governance, staked **\$OMNI** will provide governance rights to users proportional to the amount of tokens a user has staked or delegated. These governance rights will allow users to propose changes to the protocol and vote on others’ proposals. Early staking is live on the Omni [staking portal.](https://claims.omni.network/)
