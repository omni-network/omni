---
sidebar_position: 1 # Keep it first in Concepts
title: "The Problem"
---

# The Problem

Interacting with applications on different blockchains (Ethereum L1, L2s - like Base, Arbitrum, Optimism etc.) is terrible UX:

1.  **Find the Dapp:** Discover the application you want to use (e.g. deployed on Base).
2.  **Check Funds:** Check your balance on this chain the app is deployed to.
3.  **Find a Bridge:** Search for a reliable bridge to move funds.
4   **Go to Bridge:** Navigate away from the app you want to use, to said bridge.
4.  **Approve ERC20** Approve bridge to spend ERC20.
5.  **Approve Transaction** Approve the transaction to send funds, paying fees.
4.  **Wait for finality:** Wait for your funds to be bridged.
5.  **Return to Dapp:** Return to the application you wanted to use on Base.
6.  **Interact:** Finally, _finally_, use the application.

This multi-step process is not good UX, it creates friction, time delays, and causes us to lose users - we need to do better!

**Existing Solutions**

While widgets can embed bridging within a dapp's UI, they don't eliminate the underlying issues - the user still experiences multiple transaction steps (approvals, bridge), delays waiting for cross-chain settlement (often bottlenecked by source chain finality or message passing), and potential bridge-specific risks.

For application developers - choice becomes:

*   **Single-chain deployment:** Sacrificing user reach and forcing the pain point onto users.
*   **Multi-chain deployment:** Fragmenting liquidity, increasing maintenance and costs, and diluting network effects.
*   **Complex cross-chain architecture:** Requiring months of development, significant audit costs, and slower product iterations.
*   **Bridge Widget:** As discussed, this doesn't resolve the underlying issue.
