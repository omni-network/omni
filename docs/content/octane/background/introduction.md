---
sidebar_position: 1
---

# Introduction

> _Octane: The Integrated Blockchain Standard_

Octane is a next-generation modular framework for the EVM.

Octane is the standard for building integrated blockchains, allowing builders to combine the best modular components into a single, unified system. Engineered for compatibility, developers can easily leverage the latest execution, consensus, and data availability technologies without disrupting their networks. This standard enables the creation of novel L1 networks and decentralized sequencers for L2s and L3s.

Developers can use octane to run:

- any EVM execution client that implements the [EngineAPI](https://hackmd.io/@danielrachi/engine_api)
- any consensus client that implements [ABCI 2.0](https://github.com/cometbft/cometbft/tree/main/spec/abci)

Octane is the first consensus implementation of the EngineAPI besides Ethereum itself (the Beacon Chain).

It is designed and built to simplify the integration of the EVM into any blockchain application.

<details>
<summary>**Octane Explained with Legos**</summary>

Modular blockchain components are like Lego pieces. You can mix and match Legos to create different structures, just like you can mix and match modular components to create different modular stacks.

The Octane framework is like a set of instructions for building a specific Lego structure. Imagine you want to build a car – you need to assemble your Lego pieces according to the instructions to end up with a functioning car. If you don’t like the color of your car, you can follow the same instructions with different colored Lego pieces, and your car will still work as intended. Or, if you find new wheels that make your car go faster, you can use these instead, as long as you attach them in the same way as the old wheels. You can mix and match the pieces, just make sure you follow the instructions for assembly.

Similarly, Octane is a set of instructions for building integrated blockchains (blockchains made up of smaller, modular components). You can mix and match different modular blockchain components as long as they’re assembled according to Octane’s instructions. Using Octane, builders can combine the latest virtual machines, consensus engines, and node clients in a standardized manner, ensuring that their blockchain always serves its purpose for the end user.

</details>
