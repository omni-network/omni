# Omni Core

Omni Core is the foundational layer of the Omni Orderflow Engine, purpose-built to address fragmentation in Ethereum’s rollup-centric ecosystem. It combines secure cross-rollup messaging with a computational environment, the Omni EVM, to enable seamless interactions across rollups while maintaining Ethereum-native security.

<img src="/img/core_architecture.jpg" width="500px"/>

## Key Features of Omni Core

### Cross Chain Messaging

A secure, decentralized, and low-latency protocol for cross-chain communication.

Omni is secured by a validator set (initially whitelisted, but later permissionless) by staked OMNI and restaked ETH.

It has 2 confirmation strategies – `fast` and `finalized`, which have different finality guarantees and latency.

### Octane EVM

Omni was designed with a novel protocol architecture, Octane, that runs a consensus and execution engine in parallel. While Octane was developed specifically for Omni’s use case, it is open sourced and is being used by protocols that have raised over $150M, such as [Story Protocol](https://www.story.foundation/).

#### How Octane Works

- **CometBFT Consensus Engine**: Omni uses [CometBFT (formerly Tendermint)](https://docs.cometbft.com/v0.38/), a proven, high-performance consensus engine with fast finality at scale.
- **Engine API**: Separates the execution environment from the consensus engine, preventing transactions from interfering with blockchain consensus. Octane is the first consensus implementation of the Engine API besides Ethereum itself (the Beacon Chain).
- **ABCI 2.0**: complements the Engine API by providing a programmable interface for high-performance consensus engines like CometBFT.

<img src="/img/octane_architecture.jpg"/>

## Learn More

More information on Octane can be found in the following video:

<div style={{ position: 'relative', paddingBottom: '56.25%', height: 0, overflow: 'hidden', maxWidth: '100%', }}>
  <iframe
    src="https://www.youtube.com/embed/hrGgvypAMvA"
    style={{ position: 'absolute', top: 0, left: 0, width: '100%', height: '100%' }}
    frameBorder="0"
    allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
    allowFullScreen
  ></iframe>
</div>
