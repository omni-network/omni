---
sidebar_position: 3 # Third in Concepts
title: Intent Mechanism
---

# The Intent Fulfillment Mechanism

SolverNet processes cross-chain intents through a mechanism designed for security and low latency:

1.  **Intent Creation & Fund Escrow (Source Chain):**
    *   The user initiates an action via an Omni SDK-enabled frontend.
    *   The SDK translates this into a structured **intent** (the desired state change on the destination chain) and prompts the user for a signature.
    *   The user's funds required for the action are locked in a secure **escrow contract** on the source chain via a transaction that **includes the intent payload** as part of its data.
    *   This transaction serves as the user's commitment and proof of available funds.

2.  **Intent Discovery & Solver Selection:**
    *   Once the source chain transaction confirms, the confirmed intent needs to be discoverable by potential solvers.
    *   Solvers monitor for fulfillable intents (e.g., via a shared mempool, API endpoint, or future P2P network).
    *   Solvers evaluate the intent and associated costs (execution gas, settlement costs, required profit).
    *   A selection process (e.g., an off-chain auction based on price/speed, or direct routing) determines the winning solver responsible for execution.

3.  **Pre-Confirmation Execution (Destination Chain):**
    *   The winning solver uses their **own capital** on the **destination chain** to execute the transaction(s) specified by the user's intent *immediately*.
    *   This delivers a low-latency experience, as the execution happens before slower cross-chain settlement.

4.  **Settlement Proof Generation & Submission (Destination -> Source):**
    *   After successful execution on the destination chain, the solver initiates the process to prove this execution back on the source chain.
    *   A **cross-chain message** containing proof or attestation of the successful execution is relayed back to the **source chain** via an underlying messaging protocol (e.g., Omni Core).

5.  **Proof Verification & Solver Reimbursement (Source Chain):**
    *   The **escrow contract** on the source chain receives the solver's submitted proof.
    *   The contract cryptographically verifies the proof against the destination chain's state (or state roots relayed via the messaging protocol).
    *   Upon successful verification, the escrow contract releases the user's originally locked funds to the solver.

This flow ensures that user funds are secured until their intent is verifiably executed, while solvers are compensated for providing the capital and execution service, enabling fast and secure cross-chain interactions.
