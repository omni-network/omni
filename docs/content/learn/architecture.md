# Architecture

## Core Network (Cross-rollup messaging + Omni EVM)

<img src="/img/core_architecture.png" width="500px"/>

Omni is a network purpose built for fixing fragmentation across the Ethereum ecosystem. This is accomplished by combining a secure messaging protocol for cross-rollup communications with a computational environment (the Omni EVM). Given that Omni is purpose built for solving fragmentation across Ethereum’s rollup ecosystem, it derives its security from Ethereum L1 through the use of re-staking.

To achieve this, Omni was designed with a novel protocol architecture, Octane, that runs a consensus and execution engine in parallel. While Octane was developed specifically for Omni’s use case, it is open sourced and is being used by protocols that have raised over $150M, such as [Story Protocol](https://www.story.foundation/).

Omni uses Octane to combine the EVM with the [CometBFT](https://docs.cometbft.com/v0.38/) (formerly Tendermint) consensus engine, providing fast consensus on every rollup network connected with Omni. Within Octane, the Engine API separates the execution environment from the consensus engine, preventing transactions from interfering with blockchain consensus. Octane is the first consensus implementation of the EngineAPI besides Ethereum itself (the Beacon Chain). [ABCI 2.0](https://docs.cometbft.com/v1.0/spec/abci/) complements the Engine API by providing a programmable interface for high-performance consensus engines like CometBFT.

<img src="/img/octane_architecture.png" width="500px"/>

More information on Octane can be found in the following video:

<video controls width="500px">
  <!-- TODO replace with actual video -->
  <source src="/img/omni_1.mp4"/>
</video>


## Solver Network

<img src="/img/solver_architecture.png" width="500px"/>

### The Challenge of Cross-Chain Applications

The rapid expansion of rollups in the Ethereum ecosystem has made it increasingly difficult for developers to easily access users and liquidity across these chains. While cross-chain applications have often been viewed as the solution, the industry still faces major adoption obstacles, despite significant investments in interoperability infrastructure. Our discussions with hundreds of developers have revealed two primary challenges:

1. **Extended Development Timelines**

    Building a cross-chain application is resource-intensive, requiring up to four months for design, implementation, and security audits. In a fast-evolving industry where most applications are still pursuing product-market fit, this long development cycle is impractical.

2. **Higher Security Risks**

    Cross-chain implementations introduce complex smart contracts, increasing vulnerability to hacks. If an interoperability protocol is compromised, all applications relying on it are at risk. This heightened risk has discouraged widespread adoption of cross-chain applications.


These challenges have stagnated cross-chain application adoption and left the fragmentation issue unresolved.

### Omni's Intent-Centric Solution

Omni leverages intent-based interactions and solver networks to offer developers a simpler, more secure alternative to traditional cross-chain applications. Our architecture enables:

- **Easy Integration:** With a single frontend SDK integration, developers can access users and liquidity across Ethereum rollups without modifying their smart contracts. This process now takes only a few days, rather than multiple months.
- **Seamless User Experience:** Users interact with applications on different rollups without needing to bridge assets, change wallets, or worry about gas fees across networks. This architecture is fully compatible with existing wallet setups.

By abstracting away cross-chain complexity, Omni’s solver network lets both developers and users operate without facing the traditional barriers of bridging and multi-rollup deployments.

Here’s how Omni’s solver network makes applications that are only deployed on one rollup available to users across all rollups:

1. **Intent Submission**

    The user deposits tokens into a smart contract on Arbitrum and includes a payload (their "intent")—a function call they want to execute on Ethereum. This intent emits an event.

2. **Solver Action**

    Solvers monitor for these events, and immediately upon identifying one, they execute the function call on the destination platform, using the tokens as specified.

3. **Completion Acknowledgment**

    After successful execution on the destination, a cross-rollup message is sent back to the origin rollup via Omni, confirming that the intent was fulfilled. The lockbox contract on the origin then releases the funds to the solver. These funds include a small fee for the solver’s services.


<img src="/img/solver_model.png" width="500px"/>

As a result:

- The user is debited funds on Arbitrum.
- The user has deposited into the protocol on Ethereum, and has a balance there.
- The solver no longer has those tokens on Ethereum, but it has been credited tokens on Arbitrum.

Through this process, Omni handles all cross-chain communication, freeing users from interacting with multiple rollups and allowing developers to bypass the need for complex cross-chain contract deployment.

Currently we are working with a select few teams that all have over $500m TVL as our initial build partners. If your team would like to be considered as an early build partner as we iterate and expand access to the SDK, please reach out to Omni’s founders, [Austin King](https://x.com/0xASK) and [Tyler Tarsi](https://x.com/ttarsi_), on Twitter (X).
