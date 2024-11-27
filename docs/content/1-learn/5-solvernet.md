# SolverNet

SolverNet is Omni’s solution to the challenges of cross-chain application development and user experience. By leveraging an intent-based architecture and a competitive solver network, SolverNet simplifies cross-rollup execution while maintaining security and scalability.

<img src="/img/solver_architecture.jpg" width="500px"/>

## The Challenge of Cross-Chain Applications

The rapid expansion of rollups in the Ethereum ecosystem has made it increasingly difficult for developers to easily access users and liquidity across these chains. While cross-chain applications have often been viewed as the solution, the industry still faces major adoption obstacles, despite significant investments in interoperability infrastructure. Our discussions with hundreds of developers have revealed two primary challenges:

1. **Extended Development Timelines**

    Building a cross-chain application is resource-intensive, requiring up to four months for design, implementation, and security audits. In a fast-evolving industry where most applications are still pursuing product-market fit, this long development cycle is impractical.

2. **Higher Security Risks**

    Cross-chain implementations introduce complex smart contracts, increasing vulnerability to hacks. If an interoperability protocol is compromised, all applications relying on it are at risk. This heightened risk has discouraged widespread adoption of cross-chain applications.

These challenges have stagnated cross-chain application adoption and left the fragmentation issue unresolved.

## Omni's Intent-Centric Solution

Omni SolverNet leverages intent-based interactions and solver networks to offer developers a simpler, more secure alternative to traditional cross-chain applications. Our architecture enables:

- **Easy Integration:**: Developers can access users and liquidity across rollups through a single frontend SDK without modifying their smart contracts. Integration timelines shrink from months to days.
- **Seamless User Experience:** Users interact with applications on different rollups without needing to bridge assets, change wallets, or worry about gas fees across networks. This architecture is fully compatible with existing wallet setups.

By delegating cross-chain execution to solvers, Omni lets both developers and users operate without the usual complexities of bridging and multi-rollup deployments.

Here’s how Omni’s solver network makes applications that are only deployed on one rollup available to users across all rollups:

1. **Intent Submission**

    A user deposits tokens into a smart contract on their rollup of choice (e.g., Arbitrum) and includes a payload specifying the action they want to execute on a destination rollup (e.g., Ethereum). This payload is emitted as an event.

2. **Solver Action**

    Solvers monitor intent events in real time. Upon detecting an executable intent, they simulate the transaction to ensure validity and then execute the specified function on the target rollup using the required tokens.

3. **Completion Acknowledgment**

    Once the execution is complete, the solver sends a cross-rollup message via Omni Core, confirming that the intent has been fulfilled. The origin rollup’s lockbox contract releases the user’s deposited funds to the solver, along with a small execution fee.

<img src="/img/solver_model.jpg" width="500px"/>

### Results of a Successful Intent

- The user is debited funds on Arbitrum.
- The user has deposited into the protocol on Ethereum, and has a balance there.
- The solver no longer has those tokens on Ethereum, but it has been credited tokens on Arbitrum.

Through this process, Omni handles all cross-chain communication, freeing users from interacting with multiple rollups and allowing developers to bypass the need for complex cross-chain contract deployment.

## Early Adoption and Next Steps

Currently, SolverNet is being tested with select launch partners that manage over $500M in Total Value Locked (TVL). These collaborations are helping refine SolverNet’s architecture for broader adoption.

 If your team would like to be considered as an early build partner as we iterate and expand access to the SDK, please reach out to Omni’s founders, [Austin King](https://x.com/0xASK) and [Tyler Tarsi](https://x.com/ttarsi_), on Twitter (X).
