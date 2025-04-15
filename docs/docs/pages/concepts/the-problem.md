---
sidebar_position: 1 # Keep it first in Concepts
title: "The Problem"
---

# The Problem

Historically, interacting with applications across different blockchains (Ethereum L1, L2s like Base, Arbitrum, Optimism, etc.) has been cumbersome:

1.  **Find the Dapp:** Discover the application you want to use (e.g., deployed on Base).
2.  **Check Funds:** Realize your funds are on a different chain (e.g., Ethereum L1).
3.  **Find a Bridge:** Search for a reliable bridge to move funds to Base.
4.  **Bridge Assets:** Execute the bridge transaction, pay fees, and wait for finality.
5.  **Return to Dapp:** Go back to the original application on Base.
6.  **Interact:** Finally use the application.

This multi-step process introduces significant friction, time delays, potential drop-off points for users, and user attrition.

While bridging widgets can embed the bridging step within a dapp's UI, they don't eliminate the underlying issues. The user still experiences multiple transaction steps (approvals, bridge), delays waiting for cross-chain settlement (often bottlenecked by source chain finality or message passing), and potential bridge-specific risks.

For developers, it meant choosing between suboptimal options:

*   **Single-chain deployment:** Sacrificing user reach and forcing users through the painful bridging process.
*   **Multi-chain deployment:** Fragmenting liquidity, increasing maintenance overhead, and diluting network effects.
*   **Complex cross-chain architecture:** Requiring months of development, significant audit costs, and slower product iteration.
