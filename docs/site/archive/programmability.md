---
sidebar_position: 4
id: programmability
---

# Cross-Rollup Programmability

The Omni Network does not only facilitate communication across rollups -- it also enables programmability across rollups. This is a brand new, incredibly powerful programming paradigm that over time will become the de facto method of decentralized application development. We did not build Omni so that people could just pass messages across rollups, we created Omni to facilitate an entirely new programming paradigm that allows developers to think globally, not locally.

## Global, not local

Below I will enumerate some simple examples of what becomes possible in this new programming paradigm. This is not to illustrate a finite set of new features, this is to help you as the reader understand the power of a Turing complete programming interface which has global awareness of all rollups and the ability to directly update those rollups itself.

## Simple Message Passing

This is the boring use case. A majority of readers are likely familiar with this as it does not require programmability and is the functionality offered by simple interoperability protocols like LayerZero.

Example: Alice is on Arbitrum where she has 1 **\$ETH**. She sends that **\$ETH** to Optimism using a simple message passing protocol. Now Alice has 1 **\$ETH** on Optimism.

This is important functionality -- however, it is largely commoditized and uninteresting for modern application development. Our perspective is that users should never have to think about low level operations like this. Developers should build applications that abstract away this complexity from end users. This lowers friction for users and gives developers a greater addressable market as their apps will be default accessible everywhere.

## Dynamic Routing

Expanding beyond what has been previously possible - let's imagine that a developer wished to create a global margin account protocol. The aim of this global margin account would be to allow a user to control their perpetual future positions on $GMX, collateralize lending positions on Optimism, and mint stablecoins on Polygon's zkEVM.

With today's simple message passing protocols, this would be an incredibly arduous process to maintain. Due to risk of liquidation, users would need to consistently monitor their positions in each location, bridge funds to each, and top up their collateral accounts in case they are approaching liquidation. Through Omni this all can be abstracted away into a single transaction.

Instead of monitoring each position, a user can make a simple transaction stating "Ensure I am collateralized at least 200% on my perp, lending and stablecoin positions across all rollups." This transaction would then automatically be picked up by the Omni Network. Instead of passing a static message along, Omni can run custom logic to decide on how many messages get passed, the parameters inside those messages, and where those messages go alongside an automation protocol like Gelato or Autonomy.

Concretely - let's imagine the user is collateralized 300% on their stablecoin position, 190% on their  GMX position and 130% on their lending position. The Omni Network could automatically run logic to understand nothing needs to be done on the stablecoin position, but funds should be topped up on the GMX and lending collateral accounts. It could then programmatically compute the amount that needs to be sent to Arbitrum and Optimism to reach 200% collateralization on these positions and dynamically route messages to these rollups updating the positions to reflect the user request.

In this process the user made a single transaction - the Omni Network took care of the rest.

## Global By Default

Taking this to its logical conclusion we arrive at a place where all apps are global by default. A "global by default" programming paradigm reduces UX friction without compromising on functionality. In addition, it substantially expands the addressable market of users for developers building applications.

_Simple example:_ A developer wishes to conduct an NFT mint. They have built a strong community and would like to set a cap of 10,000 mints. Why would they ever limit themselves to a single rollup and their users? With Omni, now they can conduct their NFT mint in a "global by default" fashion where all of the users across the entire Ethereum ecosystem can participate, substantially expanding their likelihood of success. This is a simple, yet illustrative example of the power of this programming paradigm.

_More powerful example:_ The power of this is even more clear when you consider applications which depend on network effects. Lending markets are a great example - the greater the supply, the cheaper it is to borrow. Through building a lending market that is global by default, there is a clear opportunity to aggregate liquidity across all rollups, creating the largest supply side in the ecosystem and therefore the cheapest borrow rates. This also improves the borrower experience beyond just cost. In this scenario a user would not be confined to borrow within the confines of the rollup where they created their collateral account. They could post collateral on Arbitrum and borrow on zkSync.

Applications that are global by default will outcompete siloed applications over time due to their inherent economic advantages. If you are interested in taking advantage of this property to build a new protocol, please join our Discord to access further resources and meet other builders in the ecosystem.
