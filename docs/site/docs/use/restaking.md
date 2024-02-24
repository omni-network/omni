---
sidebar_position: 3
---

# Restaking

Omni accepts delegated ETH to [EigenLayer](https://www.eigenlayer.xyz/) AVS contracts.

## Overview

The Omni protocol’s staking implementation involves validators and delegators. Validators are responsible for verifying the authenticity of the protocol’s messages, triggering reward and slashing events, and tracking the staking balances and voting powers of the entire validator set. Delegators are solely responsible for delegating re-staked ETH to their selected validators. These staking processes are facilitated by several smart contracts within the Omni staking protocol and are shown below.

### From Single-Chain Staking to Multi-Chain

<figure>
  <img src="/img/staking.png" alt="Staking" />
  <figcaption>*Restaking ETH in Omni and Ethereum using EigenLayer AVS*</figcaption>
</figure>

The Omni staking contract is implemented on the Omni EVM. It tracks each validator’s stake and delegations, facilitates rewards distribution, and handles slashing events. The Omni AVS contract is implemented on Ethereum L1. It registers Omni as an application with the EigenLayer protocol and allows operators to opt-in to providing validation services to the Omni blockchain. Finally, Omni Portal contracts are implemented on the Omni EVM and all connected rollup VMs. These contracts maintain a copy of the Omni validator set’s stake, delegations, and voting power.

To communicate staking events, Omni leverages its own `XMsg` format. Validators monitor the Omni staking contract on the Omni EVM for stake changes and the Omni AVS contract on Ethereum L1 for delegation changes. These changes alert the Omni validator set to update its voting power details and pass the updates to Omni Portal contracts on every connected rollup VM. Finally, rewards and slashing events are initiated by the Omni validator set and are delivered to the Omni staking contract for execution.

## Omni Security Model

In the Omni network, validators play a pivotal role by managing the Omni EVM and facilitating cross-network message requests among external rollup VMs. To guarantee the integrity of validator actions, the Omni protocol incorporates a [dual staking mechanism](https://www.blog.eigenlayer.xyz/dual-staking/), securing the network through a combination of re-staked ETH and staked OMNI. A significant advantage of this approach is Omni can immediately access Ethereum’s 30.6M ETH security budget. This provides Omni with an unmatched level of security assurance among interoperability protocols, setting a new standard for network safety and stability.

One of the challenges with introducing a dual security model is ensuring augmentative security. Let us propose an architecture for dual staking that does not achieve this: modular dual staking. For each consensus game in the network, Token A stakers reach quorum on the outcome while Token B stakers independently come to their own quorum on the outcome. If one of these independent networks does not reach quorum, the vote does not pass.

Each network possesses its own security function: $s_a(A), s_b(B)$. Under this model, the cost to violate safety $S$ of the network is defined as:

$$
S = \frac{2}{3} \cdot \min(s_a(\text{totalAStaked}), s_b(\text{totalBstaked}))
$$

Additionally, the cost to violate liveness $L$ of the network is defined as:

$$
L = \frac{1}{3} \cdot \min(s_a(\text{totalAStaked}), s_b(\text{totalBstaked}))
$$

As we can see, the total cryptoeconomic security of a network under the modular dual staking model is determined by the *minimum* security derived from Token A or Token B. Effectively this approach favors implementation simplicity over security.

The Omni protocol implements its security model according to a native dual staking model. Instead of using two independent networks that reach quorum separately, Omni treats both as one set. Thus, the cost to violate safety $S$ of Omni is defined as:

$$
S = \frac{2}{3} \left( S_{\text{ETH}}(\text{totalETHstaked}) + S_{\text{OMNI}}(\text{totalOMNIstaked}) \right)
$$

Additionally, the cost to violate liveness $L$ of Omni is defined as:

$$
L = \frac{1}{3} \left( S_{\text{ETH}}(\text{totalETHstaked}) + S_{\text{OMNI}}(\text{totalOMNIstaked}) \right)
$$

Omni’s implementation of this native dual staking model enables greater security but comes with more implementation complexity. Specifically, the protocol will implement separate functions for mapping existing stake to voting power and for mapping existing stake to staking rewards. This will allow the protocol to dynamically incentivize validators to contribute more security from either ETH or OMNI depending on market conditions. Altogether, Omni is able to bootstrap an unparalleled security budget from the outset using re-staked ETH while establishing a mechanism for its total security to grow over time with the addition of staked OMNI.

<!-- TODO: inclue below

- the diagrams currently do not reflect the titles —> both show a validator operating for ETH + Omni, the second one just shows rehypothecation towards other networks
- This page will be extremely important sales collateral for winning people over — we should have clear diagrams differentiating what restaking empowers us to achieve compared to previous generation solutions that did not leverage Ethereum security
- People should see a literal picture that you only need 70 IQ to understand “wow Omni really is like 10x more secure than anything else”


-->
