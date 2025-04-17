---
sidebar_position: 1
title: What is SolverNet?
---

# What is SolverNet?

SolverNet is a network of solvers that execute intent-based actions on behalf of users. Its primary interface for developers is the Omni SDK, which allows devs to integrate SolverNet functionality into their frontend. An application may be deployed on one chain, but users can still use the app even if their funds are held on a different chain.

## Why SolverNet for Developers?

SolverNet empowers developers to **take their apps cross-chain in days**, giving them access to a wider user base at an accelerated pace.

### Integration in Days, not Months

Cross-chain integrations typically require months of development time. SolverNet removes these roadblocks with a plug-and-play SDK that gives developers the ability to scale their app across multiple networks in just days. This means faster deployment, quicker access to new users, and more time spent on improving products rather than worrying about cross-chain complexities.

Adding cross-chain functionality often requires only **~100 lines of TypeScript code** using the Omni SDK. You can see practical examples in the [Examples section](/guides/examples/symbiotic.mdx) (link points to the first example).

### No Smart Contract Changes Needed

Integrating SolverNet doesn't require any changes to existing smart contracts. This helps mitigate risk, saves developers time, and ensures that there is zero disruption to an app's existing architecture. You can stay laser-focused on product iteration.

### A Great User Experience in 1 Click

Users can interact with your protocol from anywhere in just a few seconds **and in 1 click**, without needing to bridge, swap, or leave your app.

## Why SolverNet for Users?

It's no secret that even today, many blockchain-based apps aren't easy to use. Users are often stuck:

*   Searching for a bridge
*   Manually moving assets between different chains
*   Switching networks
*   Making complex gas fee calculations

This means that the user experience is slow, complicated, and frustrating, often leading to **user attrition**. SolverNet handles all of this complexity in the background without users even realizing it – no more external bridges, no more losing users to third-party services.

## Rapid, Frictionless Growth for All Apps

For developers building the next generation of apps, speed and user experience are everything. This is why SolverNet is built to rapidly expand an app's user base while allowing devs to focus on building, not debugging, cross-chain complexities.

Instead of being locked into a single network, devs can onboard users from across the **crypto ecosystem** within days, not months. Faster growth and an improved user experience with minimal dev resources – that's the new standard SolverNet is setting for apps.

### SDK Examples

*   **[Basic ETH Deposit Example](/guides/basic-deposit.md):** Deposit ETH from Base Sepolia to Holesky.
*   **[EigenLayer Restake Example](/guides/examples/eigenlayer.mdx):** Deposit ETH from Base Sepolia into EigenLayer's stETH strategy on Holesky, showcasing signature handling.
*   **[Handling Contracts Without "onBehalfOf"](/guides/contracts-without-onbehalfof.mdx):** Using `withExecAndTransfer` (for contracts that credit `msg.sender`).
