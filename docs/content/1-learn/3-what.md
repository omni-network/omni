# What is Omni?

Omni is a chain abstraction layer, allowing developers to make their application available to users and liquidity across all rollups, while only deploying their application on a single rollup. Omni is making the Ethereum ecosystem feel like a single chain for users — no manual bridging, switching RPCs, managing gas, or wallet upgrades are needed to interact with applications powered by Omni. To simplify both user and developer experiences, Omni delegates cross-chain complexity to sophisticated third parties through Omni’s solver network.

Multiple layers of abstraction combine to form the **Omni Orderflow Engine**.

At the heart of the Omni Orderflow Engine are two core primitives:

## Omni Core

Omni Core is the base layer of the Omni Orderflow Engine. It provides:

- Cross Chain Messaging: A protocol for decentralized and fast cross-chain communication.
- Omni EVM: A computational layer designed for global computation across chains.
- Ethereum-Native Security: Through restaking, Omni Core derives its security directly from Ethereum’s base layer.

Core enables developers to deploy their applications on a single rollup and seamlessly interact with users and liquidity across the entire Ethereum ecosystem.

## SolverNet

SolverNet builds on Omni Core by introducing multi-chain intent-based execution. It allows developers to:

- Offload complex cross-chain logic to professional solvers.
- Enable users to perform actions on any rollup without needing to bridge, manage gas, or switch networks.
- Focus on building applications while SolverNet handles cross-chain execution.

Solvers in the network compete to execute user intents, ensuring efficient and cost-effective transactions.

<img src="/img/architecture.jpg" width="500px"/>

## How does it work?

Here's a quick breakdown of how the Omni Orderflow Engine works.

1. User Intent Submission:

   A user declares a desired action (intent) on their preferred rollup. This action might involve interacting with a protocol on another rollup.

2. Solver Execution:

   Solvers detect the intent, bid to execute it on the destination, and provide "just in time" liquidity to the destination rollup by depositing funds on behalf of the user.

3. Cross-Rollup Settlement:

   Using Omni Core, the solvers provide proff of execution of the user's intent, allowing the escrow contract on the origin rollup to release the users deposited funds.

## Why the Omni Orderflow Engine Matters

The Omni Orderflow Engine bridges the gap between Ethereum’s rollup-centric scalability and the user experience of a monolithic chain. By combining Omni Core and SolverNet, it simplifies the complexities of cross-rollup development and ensures a seamless experience for end users.

Explore the following sections to dive deeper into Omni Core and SolverNet:

- Omni Core: Learn how Omni Core powers cross-rollup messaging and computation.
- SolverNet: Discover how SolverNet enables intent-based execution across rollups.

Omni’s solver network has timelock escrow contracts on every supported rollup — these allow users to deposit funds and declare intents for actions they want completed on applications deployed on other rollups.

After a user deposits funds into escrow and declares their intent on the source rollup, a solver provides “just in time liquidity” on the destination rollup by depositing capital into the target application on behalf of the user.

The solver then provides proof of completing the user’s intent through Omni’s interoperability network, allowing the escrow contract on the origin rollup to release the user’s deposited funds to the solver.
