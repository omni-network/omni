---
sidebar_position: 2 # Second in Concepts
title: "The Solution: Intents & SolverNet"
---

# User-Centric Actions via Intents

Intents shift the focus from: *moving assets between chains*, to: *achieving user goals across chains*.

Instead of telling the system *how* to achieve something (bridge 1000 USDC from Arbitrum to Base, then call function Y on contract Z), a user simply expresses their desired outcome (their **intent**): "I want to deposit 1000 USDC into Vault Z on Base, using my funds on Arbitrum."

## How we use intents

1.  **User Expresses Intent -** In your frontend, the user expresses an action (e.g. clicking "Deposit" or "Stake"). Using our SDK helps you formulate this intent.
2.  **Intent Broadcast -** The intent (e.g. "deposit 1000 USDC from Arbitrum into target contract on Base") is broadcast to a network of sophisticated actors called **_Solvers_**.
3.  **Solvers Compete -** Solvers are kind of like specialized market makers. They maintain capital across chains and understand how to execute actions efficiently. They compete to fulfill the user's intent quickly and cheaply.
4.  **Solver Executes -** The winning solver takes on the risk and complexity:
    *   They use their *own funds* on the destination chain (Base) to make the deposit into Vault V *immediately* on the user's behalf.
    *   They simultaneously handle the bridging/settlement of the user's funds from the source chain (Arbitrum) to reimburse themselves.
5.  **User Goal Achieved -** From the user's perspective, they clicked "Deposit," and within seconds, their deposit is confirmed in Vault V on Base. They never had to leave the application, find a bridge, or manage assets across chains.

**Analogy: Food Delivery**

Think of ordering food online. Your *intent* is "I want a pizza delivered to my house." You don't specify the route the driver should take, which car they use, or how they handle traffic. You express your goal, and the delivery service (the Solver) figures out the optimal way to fulfill it.
