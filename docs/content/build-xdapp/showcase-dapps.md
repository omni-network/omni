---
sidebar_position: 3
---

import GitHubCodeBlock from '@site/src/components/GitHubCodeBlock/GitHubCodeBlock';

# Showcase



## xstake

 An [example xchain staking app](https://github.com/omni-netwpr/xstake) that allows ERC20 deposits on multiple chains.

 <GitHubCodeBlock url="https://github.com/omni-network/xstake/blob/main/src/XStaker.sol"/>


## Ethereum to Omni bridge

OMNI is an ERC-20 token on Layer 1, and we understand that developers and users need a seamless way to transfer it to the Omni EVM for gas payments. To address this, we've developed a bridge that utilizes the Omni protocol, ensuring smooth and secure token transfers.


### Omni to Ethereum
<GitHubCodeBlock url="https://github.com/omni-network/omni/blob/main/contracts/core/src/token/OmniBridgeL1.sol"/>

### Ethereum to Omni
<GitHubCodeBlock url="https://github.com/omni-network/omni/blob/main/contracts/core/src/token/OmniBridgeNative.sol" />


## What else?

:::tip
Get inspired by the following ideas. Looking forward to seeing your PRs on https://github.com/omni-network/awesome-omni
:::

<div class="grid-wrapper">
  <div class="grid-rfp">
    <div class="rfp-grid-cell">
      ### NFT Community Backed Stablecoin
      Take a heavily community / social-backed asset and use it as a stablecoin. $CULT (Remelia NFT family) or other strong NFT community-backed token could be used as initial collateral source.
    </div>
    <div class="rfp-grid-cell">
      ### Real Estate Index Synthetics / Perps
      GMX-style perps that track commercial and residential real estate indices from major global cities. Pooled collateral model.
    </div>
    <div class="rfp-grid-cell">
      ### TikTok Perps
      GMX-style perps that track hashtags, sounds, follower counts on TikTok clone (or pulled from actual TikTok API) application.
    </div>
    <div class="rfp-grid-cell">
      ### Twitch Livestream Spot / Perps Betting
      Betting markets for Twitch livestreams – i.e., Fortnite already allows players to bet on their own number of kills in a match. Could create live lines / markets, revenue share / portion of total pool goes back to streamer.
    </div>
    <div class="rfp-grid-cell">
      ### Weather Betting
      Spot / leverage / perps trading on weather – i.e., will it rain in X city on Y day (futures), will it rain in X and Y cities on the same day (parlay), over/under on temperatures, will temperature be above/below X on Y day of the week.
    </div>
    <div class="rfp-grid-cell">
      ### Rocket Game
      Game which displays an ascending rocket that explodes at a “random” point in time with an increasing multiplier / reward until explosion. Allow for social live betting on the side for greater social coordination.
    </div>
    <div class="rfp-grid-cell">
      ### Synthetic Asset Management
      GMX-style spot / perps on commonly traded ETFs and mutual funds or foreign stocks / commodities.
    </div>
    <div class="rfp-grid-cell">
      ### Google Search Predictions
      Trending search markets on Google, binary predictions on if X or Y word/phrase will have more searches on Z day, week, month.
    </div>
    <div class="rfp-grid-cell">
      ### VC / Angel Perps
      Perps / synthetics that bucket VC-portfolio tokens or angel allocations – i.e., Paradigm = $UNI, $DYDX, Panterra = $INJ, $OMNI, Multicoin = $SOL, $SEI.
    </div>
    <div class="rfp-grid-cell">
      ### Frictionless NFTs
      Launchpad for [xERC-721s](https://www.openliquidity.org/research-eips/eip-7611) as mechanism to most easily create chain-agnostic NFTs / POAPs, SBTs, in-game assets.
    </div>
    <div class="rfp-grid-cell">
      ### Universal Launchpad
      Chain-agnostic crowdfunding launchpad platform that pulls deposits from any L2 and creates tokens that can be traded across all L2s.
    </div>
    <div class="rfp-grid-cell">
      ### Chain-Agnostic Casino
      Chain-agnostic casino games that allow users to deposit collateral and participate in on-chain games from any Ethereum L2 without the need to bridge, manage multiple assets, wallets, etc.
    </div>
    <div class="rfp-grid-cell">
      ### Ultra Simple Binary Markets
      Single-market app that pulls deposits from every chain – best suited for one massive event (i.e., presidential election). Painfully simple frontend with deposit contracts.
    </div>
  </div>
</div>
