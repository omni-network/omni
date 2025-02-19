# Why Omni?

## The Problem

The most glaring side effect of Ethereum’s pursuit of the rollup-centric roadmap is fragmentation of users and capital across the Ethereum ecosystem. This has degraded both the developer and user experiences on Ethereum.

### For developers

- Applications are siloed on individual rollups, limiting their ability to reach users across the ecosystem
- If a developer wants to reach users across the ecosystem, they have a couple of options:
    1. Deploy multiple copies of their application to multiple rollups, which means each deployment is isolated from the others, must be maintained independently. Congratulations, your users are fragmented, and your application doesn't benefit from network effects.
    2. Rearchitect their applications to be cross chain. This forces developers to reason about cross chain communication, finality risk, sequencer risk, platform risk, and longer audits. Plus, it turns smart contract development into a distributed systems problem! The whole point of the EVM is to create a simple VM abstraction over a complex distributed system.
    3. Give up and just deploy on a single rollup, limiting themselves to the users and capital on that rollup.
- None of these options are good, let alone desirable.

### For users

They must navigate application frontends, select networks, switch RPC endpoints, and manually bridge assets between networks, all while ensuring they have enough gas to complete transactions on the destination network.

It's a nightmare.

## The Solution

Omni simplifies and unifies the Ethereum ecosystem by abstracting away cross-chain complexity.

**For Developers:**

- Deploy your application on a single rollup.
- Instantly reach users and liquidity across the Ethereum ecosystem.
- No need to redesign or re-audit contracts for cross-chain functionality.
- No need to maintain multiple deployments of your application.
- No need to manage cross-chain logic in your contracts, or worry about cross chain security.

**For Users:**

- Seamless access to decentralized applications without bridging, network switching, or managing gas fees across chains.
- Omni makes the Ethereum ecosystem feel like a single, unified chain.

Omni ensures that the Ethereum ecosystem can scale without sacrificing usability or developer experience.
