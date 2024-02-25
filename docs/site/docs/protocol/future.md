---
sidebar_position: 4
---

# Future Work

Omni was intentionally designed not only to establish a new standard in the interoperability industry, but also to expand in both functionality and coverage in the future. Specifically, Omni will introduce novel finality mechanisms, support for alternative VM environments beyond the EVM, and support for alternative DA and consensus layers.

## Fast Finality Mechanisms

While Omni achieves its own finality in less than a second, it must still contend with Ethereum's ~12 minute finality window for processing incoming XMsg requests. The challenge lies in enabling Omni to bypass this delay for its users. A promising potential approach involves allowing re-staked ETH delegations to be used as insurance for cross-network transactions, shifting the risk of finality in exchange for a fee. Implementing such a strategy could enable Omni to fully shield users from the complexities of rollups, making the use of crypto applications across various rollups as straightforward as accessing websites across different data centers.

## Expanded VM Support

Omni will initially prioritize support for rollups that utilize the EVM, recognizing its significant network effects and widespread adoption across the crypto industry. However, as described in Section 3, Omni is designed to transcend any single VM architecture and will scale support for additional VM types as innovation at the VM layer continues. Omni will soon extend support to rollups utilizing alternative VMs, including the Move VM, the Solana VM, and CosmWasm VMs.

## Expanded DA and Consensus Support

In alignment with the strategy for embracing diverse VMs, Omni will initially focus on rollups leveraging Ethereum for DA and consensus. As the Omni ecosystem evolves, it will incorporate support for DA solutions and consensus networks. This expansion into various DA networks and VM environments will position Omni as the go-to interoperability provider for any rollup configuration.
