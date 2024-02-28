---
sidebar_position: 2
id: restaking
---

# ETH Restaking

Omni provides a missing, yet essential, infrastructure layer for Ethereum's rollup centric future. This is why, when designing the Omni Network, we asked ourselves "What would be the best way to empower developers to build applications across all rollups if it was part of the core Ethereum protocol?" The answer to this question is to utilize restaking, creating an auxiliary network consisting of existing Ethereum validators who assume additional responsibilities to facilitate this functionality.

Omni sources a global worldview of all rollup state through the coordination of restakers who participate in a set of consensus rules. In order to be eligible to participate in consensus of the network, the validators must stake **\$ETH** so that they have capital at stake that can be slashed in the event that they misbehave and do not follow the agreed upon consensus rules of the network.

## Restaking as a concept

The Omni validators actually "restake" their **\$ETH** -- meaning that they can simultaneously participate in consensus both for the core Ethereum network alongside the Omni network. This concept was invented by the team at EigenLayer and it is a mechanism that allows validators to increase their rewards, while simultaneously opening up an entirely new class of networks like Omni to be formed that further expand functionality of Ethereum network in a way that could not be achieved just through the deployment of smart contracts.

The creation of auxiliary networks that expand the functionality of Ethereum clearly presents great value to the ecosystem of Ethereum developers, but it also empowers the creators of these auxiliary networks with an opportunity to use a credibly neutral, difficult to manipulate and large market cap token as a form of pure collateral to bootstrap new networks.

## How Omni uses restaking

Omni Validators restake **\$ETH** in order to join the list of nodes that participate in the operation of the network. Upon restaking their capital, it is in their best interest to follow the outlined rules of consensus. If they do not, their **\$ETH** will be slashed and they will lose money. If they participate as expected, they will earn rewards through transaction fees submitted to the network. Therefore, through leveraging restaking we have a way to coordinate entities across the world who do not need to know or trust one another to all operate software in a way that provides a net societal benefit. Let's walk through a simple example.

Alice has $100 worth of **\$ETH** on Arbitrum and would like to borrow some $USDC against that position to buy a token for a new project that just launched on Optimism. Using Omni, she could escrow her funds on Arbitrum, have that confirmation relayed to Optimism, and borrow $50 on Optimism while using her $100 of **\$ETH** on Arbitrum as collateral. Alice can now use that $50 to invest in the new project that she is excited about.

This is a simple, yet illustrative example of the power of Omni -- especially because this can be abstracted away from Alice. She might not care about which rollup her money is on, she might just want to participate in the decentralized web and Omni allows application developers to abstract all this away from end users to give them the best experiences the crypto industry has to offer without burdening them with technical requirements.

The way that the Omni Network secures this transition of data from Arbitrum to Optimism is restaking. The validators in the network would be monitoring Arbitrum and see Alice's transaction. It is in all of their interest to truthfully report the data of this transaction and they will do this by signing a statement declaring their view on this transaction that can be associated with their address that has restaked **\$ETH**. If they lie, ultimately Arbitrum will post this data to layer 1 Ethereum and then anybody will be able to take the false, yet signed statement by a dishonest validator, submit it to the restaking contract and slash the dishonest validator.

Validators restake **\$ETH** and attest to state updates that happen in rollups. They are incentivized to participate in this network through rewards that manifest from transaction fees. They are disincentivized from lying because ultimately the truth of all data will be posted down from rollups to layer 1 Ethereum where the validator will have their restaked **\$ETH** slashed. It is simple, elegant, and scalable.

This is how the Omni Network is secured. The simplicity is intentional. We are building a network to secure the future of the crypto industry. There is no space for convoluted data structures, complex multi-agent interactions or third party dependencies. This is the only model that can keep pace as we onboard hundreds of millions of people across the world into an open, permissionless economy.
