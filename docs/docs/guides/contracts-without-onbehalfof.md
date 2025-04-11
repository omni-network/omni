---
sidebar_position: 2
title: Handling Contracts Without "onBehalfOf"
---

# Guide: Handling Contracts Without "onBehalfOf" (`withExecAndTransfer`)

This guide demonstrates how to use the `withExecAndTransfer` utility for cross-chain interactions with destination contracts that **do not** support acting `onBehalfOf` a user. Instead, these contracts might credit assets (like ERC20 tokens, NFTs, or even native ETH) directly to the `msg.sender` of the transaction.

**Scenario:** A user has native ETH on Base Sepolia and wants to deposit it into a `SampleVault` contract on Holesky. This vault has a payable `deposit()` function that returns vault shares (an ERC20 token) directly to the `msg.sender`.

## Problem

If we directly call `SampleVault.deposit()` via SolverNet, the `msg.sender` on Holesky will be an **Omni Network contract** executing the solver's instructions, not the user's address. The Omni contract would receive the vault shares, not the user who initiated the deposit.

## Solution: `withExecAndTransfer`

We wrap the `SampleVault.deposit()` call using `withExecAndTransfer`. This tells the Omni system (via the solver) to:

1.  Call the `SolverNetMiddleman` contract on Holesky.
2.  The `SolverNetMiddleman` then calls `SampleVault.deposit()` with the required ETH (`value`).
3.  After the `deposit` call returns, the `SolverNetMiddleman` (which was the `msg.sender` for the vault call and thus received the vault shares) transfers those specific vault shares to the original user's address.

## Prerequisites

*   Omni SDK installed and `OmniProvider` set up ([Getting Started](../sdk/getting-started.md)).
*   `wagmi` and `@tanstack/react-query` configured.
*   Target contract (`SampleVault`) ABI and address known.
*   The address of the asset (ERC20 token, NFT, or `zeroAddress` for ETH) returned/credited by the target contract is known.
*   Source and destination chain IDs known.

## Steps

1.  **Import necessary hooks and utilities:**

```tsx
import React, { useState, useMemo } from 'react';
import { useAccount } from 'wagmi';
import {
  useQuote,
  useOrder,
  useOmniContracts, // Hook to get Middleman address
  withExecAndTransfer // The utility function
} from '@omni-network/react';
import { parseEther, formatEther, type Abi, type Address, zeroAddress } from 'viem';
import { baseSepolia, holesky } from 'viem/chains';

// --- Configuration (Replace with your actual values) ---

// ABI for the SampleVault's payable deposit function
const sampleVaultABI = [
  {
    inputs: [],
    name: 'deposit',
    outputs: [], // Assume it doesn't explicitly return the token address
    stateMutability: 'payable',
    type: 'function',
  },
] as const;

// Vault contract address on the destination chain (Holesky)
const vaultAddress = '0xYourSampleVaultAddressOnHolesky' as const;

// Address of the ERC20 token credited by the vault
const vaultTokenAddress = '0xYourVaultTokenAddressOnHolesky' as const;

// Chain IDs
const sourceChainId = baseSepolia.id;
const destChainId = holesky.id;
// ------------------------------------------------------
```

2.  **Create the React Component:**
    Set up state for the deposit amount (native ETH in this case).

```tsx
function TokenizedDepositForm() {
  const [depositAmountStr, setDepositAmountStr] = useState<string>('0.1');
  const { address: userAddress, isConnected } = useAccount();

  const depositAmount = useMemo(() => {
    try {
      return parseEther(depositAmountStr as `${number}`);
    } catch {
      return 0n;
    }
  }, [depositAmountStr]);

  // ... (hooks will go here)

  return (
    <div>
      <h2>Deposit ETH from Base Sepolia to Tokenized Holesky Vault</h2>
      <label>
        Amount to Deposit (ETH on Base Sepolia):
        <input
          type="text"
          value={depositAmountStr}
          onChange={(e) => setDepositAmountStr(e.target.value)}
        />
      </label>

      {!isConnected && <p>Please connect your wallet.</p>}

      {/* ... (Quote, Contracts, Order logic) ... */}
    </div>
  );
}
```

3.  **Implement `useQuote`:**
    Quote the required source chain deposit (native ETH) for the desired destination chain expense (native ETH to be sent with the vault's `deposit()` call).

```tsx
  // Inside TokenizedDepositForm component
  const quote = useQuote({
    srcChainId: sourceChainId,
    destChainId: destChainId,
    deposit: {
      isNative: true, // Depositing native ETH from source
      amount: depositAmount
    },
    expense: {
      isNative: true // Vault deposit() is payable, expects native ETH
    },
    mode: "expense",
    enabled: isConnected && depositAmount > 0n,
  });

  const quotedDepositAmt = quote.data?.deposit.amount ?? 0n; // ETH to deposit on Base Sepolia
  const quotedExpenseAmt = quote.data?.expense.amount ?? 0n; // ETH to be spent on Holesky (value for vault call)
```

4.  **Fetch Middleman Address using `useOmniContracts`:**

```tsx
  // Inside TokenizedDepositForm component
  const contracts = useOmniContracts(destChainId); // Pass destination chain ID
  const middlemanAddress = contracts.data?.middleman ?? zeroAddress;
```

5.  **Prepare the Wrapped Call using `withExecAndTransfer`:**
    Define the original call to the vault and then wrap it.

```tsx
  // Inside TokenizedDepositForm component
  const middlemanCall = useMemo(() => {
    if (!userAddress || middlemanAddress === zeroAddress) return undefined;

    // 1. Define the original call to the vault
    const originalVaultCall = {
        target: vaultAddress,
        abi: sampleVaultABI,
        functionName: 'deposit',
        value: quotedExpenseAmt, // ETH for payable function
    };

    // 2. Wrap the call
    return withExecAndTransfer({
        middlemanAddress: middlemanAddress,
        call: originalVaultCall,
        transfer: {
            // The token the middleman should transfer *after* executing the call
            // This is the asset credited to msg.sender by the 'deposit' function
            token: vaultTokenAddress,
            // The final recipient of the credited asset
            to: userAddress,
        }
    });
  }, [userAddress, middlemanAddress, vaultAddress, sampleVaultABI, quotedExpenseAmt, vaultTokenAddress]);
```

6.  **Implement `useOrder`:**
    Configure the order, passing the `middlemanCall` into the `calls` array. Note that the `spender` for the `expense` is now the `middlemanAddress`.

```tsx
  // Inside TokenizedDepositForm component
  const order = useOrder({
    srcChainId: sourceChainId,
    destChainId: destChainId,
    deposit: {
      isNative: true,
      amount: quotedDepositAmt, // User deposits native ETH
    },
    expense: {
      isNative: true,
      amount: quotedExpenseAmt, // Solver spends native ETH (to call the vault)
      spender: middlemanAddress // Middleman receives ETH from solver to make the call
    },
    // Pass the wrapped call generated by withExecAndTransfer
    // Ensure middlemanCall is defined before enabling
    calls: middlemanCall ? [middlemanCall] : [],
    validateEnabled: quote.isSuccess && contracts.isSuccess && middlemanCall != null,
  });

  const {
    open: openOrder,
    status: orderStatus,
    validation,
    isReady,
    isTxPending,
    error: orderError,
  } = order;

  // Enable button only if quote, contracts, and middleman call are ready
  const canOpen = isReady && validation?.status === 'accepted' && !isTxPending && middlemanCall != null;
```

7.  **Add UI Elements:**
    Similar to the basic deposit, display status and the action button.

```tsx
  // Inside the return statement of TokenizedDepositForm
  return (
    <div>
      {/* ... Input field ... */}

      {isConnected && depositAmount > 0n && (
        <div style={{ marginTop: '15px' }}>
          {quote.isLoading && <p>Fetching quote...</p>}
          {contracts.isLoading && <p>Fetching contracts...</p>}
          {/* ... Other status messages ... */}

          {middlemanCall === undefined && contracts.isSuccess && <p>Preparing wrapped call...</p>}

          <button
            disabled={!canOpen}
            onClick={() => openOrder?.()}
            style={{ marginTop: '10px' }}
          >
            {isTxPending ? 'Opening Order...' : 'Deposit ETH via Omni (Tokenized)'}
          </button>

          <p style={{ marginTop: '10px' }}>Order Status: <strong>{orderStatus}</strong></p>
          {/* ... Error messages ... */}
        </div>
      )}
    </div>
  );
}

export default TokenizedDepositForm;
```

## Key Considerations

*   **`spender`:** When using `withExecAndTransfer`, the `spender` in the `useOrder` expense configuration **must** be the `middlemanAddress` for the destination chain.
*   **`transfer.token`:** This is the address of the asset you expect the *middleman* to receive (because it was the `msg.sender` of the wrapped `call`) and subsequently transfer to the `transfer.to` address. Use `zeroAddress` if the credited asset is native ETH.
*   **`call.value`:** If the original target function is `payable`, ensure the correct ETH amount (usually `quotedExpenseAmt`) is set in the `value` field of the `call` object passed *into* `withExecAndTransfer`.
*   **Dependencies:** Ensure hooks like `useOmniContracts` and calculations like `middlemanCall` complete successfully before enabling `useOrder` or the final action button.
