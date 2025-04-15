---
sidebar_position: 3 # Or adjust if needed after restructuring
title: "Developer Workflow: Single-Chain Deployment"
---

# Developer Workflow: Single-Chain Deployment

SolverNet, powered by [intents](/concepts/the-solution.md), enables a new, simplified paradigm for decentralized application development: **deploy on a single chain, make it accessible from anywhere.**

## The New Model: Single Chain + Intents

This approach offers a superior alternative to previous multi-chain strategies:

1.  **Deploy Smart Contracts to ONE Chain:** Choose a single blockchain (L1 or L2) that best aligns with your project's values, security requirements, ecosystem preferences, or technical needs.
    *   Need Ethereum's security? Deploy on L1.
    *   Value Optimism's public goods funding? Deploy on OP Mainnet.
    *   Want proximity to high DeFi activity? Choose Base or Arbitrum.
    *   This keeps your backend simple, secure, and easy to manage.
2.  **Integrate the Omni SDK into Your Frontend:** Add the `@omni-network/react` package and use the hooks (`useQuote`, `useOrder`) in your application's user interface.
    *   This integration typically takes only a few hours and adds minimal code (often just **~100 lines of TypeScript**).
    *   Requires **zero changes** to your existing smart contracts.
    *   No additional audits needed for your backend.

**Result:**

*   **Global Accessibility:** Users on *any* supported chain can interact with your application seamlessly, directly from your frontend, without needing to bridge manually.
*   **Native UX:** The cross-chain interaction happens in seconds with minimal clicks.
*   **Simplified Development:** Your team focuses on building core product features on your chosen home chain.
*   **Unified Liquidity & State:** Your application logic and state remain consistent on a single chain.
*   **Faster Iteration:** Avoid the delays and complexities of multi-chain deployments or re-architectures.

This "Single Chain + Intents" model allows you to achieve the benefits of being multichain (wide user access, tapping into diverse liquidity) without the traditional drawbacks. You can build on the chain you believe in and let SolverNet handle the rest.
