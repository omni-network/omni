---
sidebar_position: 3
---

# Implementation

<figure>
  <img src="/img/staking.png" alt="Staking" />
  <figcaption>*Restaking ETH in Omni and Ethereum using EigenLayer AVS*</figcaption>
</figure>

The Omni staking contract is implemented on the Omni EVM. It tracks each validator’s stake and delegations, facilitates rewards distribution, and handles slashing events. The Omni AVS contract is implemented on Ethereum. It registers Omni as an application with the EigenLayer protocol and allows operators to opt-in to providing validation services to the Omni blockchain. Finally, Omni Portal contracts are implemented on the Omni EVM and all connected rollup VMs. These contracts maintain a copy of the Omni validator set’s stake, delegations, and voting power.

To communicate staking events, Omni leverages its own `XMsg` format. Validators monitor the Omni staking contract on the Omni EVM for stake changes and the Omni AVS contract on Ethereum for delegation changes. These changes alert the Omni validator set to update its voting power details and pass the updates to Omni Portal contracts on every connected rollup VM. Finally, rewards and slashing events are initiated by the Omni validator set and are delivered to the Omni staking contract for execution.
