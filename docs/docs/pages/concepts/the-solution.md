---
sidebar_position: 2 # Second in Concepts
title: "The Solution: Intents & SolverNet"
---

# The Solution: User-Centric Actions via Intents

Intents shift the focus from *moving assets* between chains to *achieving user goals* across chains.

Instead of telling the system *how* to achieve something (bridge X amount of ETH from A to B, then call function Y on contract Z), a user simply expresses their desired outcome (their **intent**): "I want to deposit 0.1 ETH into Vault V on Base, using my funds currently on Arbitrum."

## How SolverNet Uses Intents:

1.  **User Expresses Intent:** Through your application's frontend (using the Omni SDK), the user clicks a button like "Deposit" or "Stake". The SDK helps formulate this intent.
2.  **Intent Broadcast:** The intent (e.g., "deposit X amount of Token A from source chain S into target contract T on destination chain D") is broadcast to a network of sophisticated actors called **Solvers**.
3.  **Solvers Compete:** Solvers are like specialized market makers. They maintain capital across various chains and understand how to execute actions efficiently. They compete to fulfill the user's intent quickly and cheaply.
4.  **Solver Executes:** The winning solver takes on the complexity:
    *   They use their *own funds* on the destination chain (Base) to make the deposit into Vault V *immediately* on the user's behalf.
    *   They simultaneously handle the bridging/settlement of the user's funds from the source chain (Arbitrum) to reimburse themselves.
5.  **User Goal Achieved:** From the user's perspective, they clicked "Deposit," and within seconds, their deposit is confirmed in Vault V on Base. They never had to leave the application, find a bridge, or manage assets across chains.

### Analogy: Food Delivery

Think of ordering food online. Your *intent* is "I want a pizza delivered to my house." You don't specify the route the driver should take, which car they use, or how they handle traffic. You express your goal, and the delivery service (the Solver) figures out the optimal way to fulfill it.
