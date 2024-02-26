---
sidebar_position: 5
id: permissionless
---

# Permissionless Expansion

The Omni Network is designed to facilitate permissionless expansion. This is essential in our vision of creating an open and free global economy, but it is also simply pragmatic. In order to provide the most utility to end users, the Omni Network needs a way to be permissionlessly deployed across all rollups. We achieve this through portal contracts.

## Portal Contracts

Portal contracts are solidity contracts that effectively function as the doorway through which actions flows to and from the Omni Network. Given that they are solidity contracts, they can be deployed on any rollup network without requiring anything from the rollup itself. Validators in the Omni Network monitor updates to these contracts and relay these transactions into the Omni computing environment (see: ), and then propagate resulting actions to any relevant rollup.

## Global By Default

This also benefits developers in their ability to reach wider audiences. When designing applications that are global by default (covered in: [Cross-Rollup Programmability](./programmability)) the portal contract design makes it incredibly easy to propagate interfaces for their applications to any rollup they wish it to be available on as these interfaces just link directly into the portal contracts. This makes the process of expanding an application to a new rollup as simple as appending the name of the rollup in a YAML file. This furthers our design goal of building a decentralized application development platform that allows builders to easily access all users, not just those confined within a single rollup.
