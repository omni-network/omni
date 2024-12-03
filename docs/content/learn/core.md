# Core

Omni Core is the building blocks layer of the Omni Orderflow Engine, purpose-built for fast, secure cross-rollup communicaiton and computation. It combines secure cross-rollup messaging with the Omni EVM.

## Key Features of Omni Core

### Cross Chain Messaging

Omni Core implements a secure, decentralized, and low-latency protocol for cross-rollup communication build on the `XMsg` primitive.

- Users initiate cross-rollup function calls via Omni’s Portal contract.
- Validators securely process and attest to cross chain blocks (`XBlocks`) using CometBFT consensus.
- Relayers deliver validated messages to the destination rollup, where they are processed and executed.

#### Key Features

- **Decentralized**: powered by staked OMNI and restaked ETH.
- **Flexible and fast**: choose between `fast` and `finalized` finality.

### Octane EVM

Omni was designed with a novel protocol architecture, Octane, that runs a consensus and execution engine in parallel. While Octane was developed specifically for Omni’s use case, it is open sourced and is being used by protocols that have raised over $150M, such as [Story Protocol](https://www.story.foundation/).

Here's how Octane works:

- **CometBFT Consensus Engine**: Omni uses [CometBFT (formerly Tendermint)](https://docs.cometbft.com/v0.38/), a proven, high-performance consensus engine with fast finality at scale.
- **Engine API**: Separates the execution environment from the consensus engine, preventing transactions from interfering with blockchain consensus. Octane is the first consensus implementation of the Engine API besides Ethereum itself (the Beacon Chain).
- **ABCI 2.0**: complements the Engine API by providing a programmable interface for high-performance consensus engines like CometBFT.

<img src="/img/octane_architecture.jpg" width="750px"/>

## Learn More

More information on Octane can be found in the following video:

<div style={{ position: 'relative', paddingBottom: '56.25%', height: 0, overflow: 'hidden', maxWidth: '100%', }}>
  <iframe
    src="https://www.youtube.com/embed/hrGgvypAMvA"
    style={{ position: 'absolute', top: 0, left: 0, width: '80%', height: '80%' }}
    frameBorder="0"
    allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
    allowFullScreen
  ></iframe>
</div>
