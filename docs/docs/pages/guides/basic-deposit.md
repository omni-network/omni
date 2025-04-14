---
sidebar_position: 1
title: Basic Cross-Chain Deposit
---

# Guide: Basic Cross-Chain Deposit

This guide walks through the standard process of enabling a user to deposit assets from one chain into a contract on another chain using the Omni SDK, assuming the target contract supports depositing on behalf of a user (e.g., via an `onBehalfOf` parameter).

**Scenario:** A user has wstETH on Base Sepolia and wants to deposit it into a vault contract on Holesky. The vault contract has a `deposit(address onBehalfOf, uint256 amount)` function.

## Prerequisites

*   Omni SDK installed and `OmniProvider` set up ([Getting Started](/sdk/getting-started.mdx)).
*   `wagmi` and `@tanstack/react-query` configured.
*   Target contract ABI and address known.
*   Source and destination chain IDs and token addresses known.

## Steps

1.  **Import necessary hooks and utilities:**

```tsx
import React, { useState } from 'react';
import { useAccount } from 'wagmi';
import { useQuote, useOrder } from '@omni-network/react';
import { parseEther, formatEther, type Abi, type Address } from 'viem';
import { baseSepolia, holesky } from 'viem/chains';

// --- Configuration (Replace with your actual values) ---

// Vault contract ABI with the deposit function
const vaultABI = [
  {
    inputs: [
      { internalType: 'address', name: 'onBehalfOf', type: 'address' },
      { internalType: 'uint256', name: 'amount', type: 'uint256' },
    ],
    name: 'deposit',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
] as const;

// Vault contract address on the destination chain (Holesky)
const vaultAddress = '0xYourVaultContractAddressOnHolesky' as const;

// Token addresses
const sourceTokenAddress = '0x6319df7c227e34B967C1903A08a698A3cC43492B' as const; // wstETH on Base Sepolia
const destTokenAddress = '0x8d09a4502Cc8Cf1547aD300E066060D043f6982D' as const;   // wstETH on Holesky

// Chain IDs
const sourceChainId = baseSepolia.id;
const destChainId = holesky.id;
// ------------------------------------------------------
```

2.  **Create the React Component:**
    Set up state for the deposit amount input.

```tsx
function BasicDepositForm() {
  const [depositAmountStr, setDepositAmountStr] = useState<string>('0.1');
  const { address: userAddress, isConnected } = useAccount();

  // Convert input string to bigint, handle potential errors
  const depositAmount = React.useMemo(() => {
    try {
      return parseEther(depositAmountStr as `${number}`);
    } catch {
      return 0n; // Handle invalid input
    }
  }, [depositAmountStr]);

  // ... (hooks will go here)

  return (
    <div>
      <h2>Deposit wstETH from Base Sepolia to Holesky Vault</h2>
      <label>
        Amount to Deposit (wstETH on Base Sepolia):
        <input
          type="text"
          value={depositAmountStr}
          onChange={(e) => setDepositAmountStr(e.target.value)}
        />
      </label>

      {!isConnected && <p>Please connect your wallet.</p>}

      {/* ... (Quote and Order logic will go here) ... */}
    </div>
  );
}
```

3.  **Implement `useQuote`:**
    Fetch a quote to determine the corresponding expense amount on the destination chain for the user's desired deposit amount.

```tsx
  // Inside BasicDepositForm component
  const quote = useQuote({
    srcChainId: sourceChainId,
    destChainId: destChainId,
    deposit: {
      isNative: false,
      token: sourceTokenAddress,
      amount: depositAmount // Use the state variable
    },
    expense: {
      isNative: false,
      token: destTokenAddress
    },
    mode: "expense", // We specify deposit, quote the expense
    enabled: isConnected && depositAmount > 0n, // Only run if connected and amount > 0
  });

  // Get the exact amounts from the successful quote
  const quotedDepositAmt = quote.data?.deposit.amount ?? 0n;
  const quotedExpenseAmt = quote.data?.expense.amount ?? 0n;
```

4.  **Implement `useOrder`:**
    Configure the order using the amounts from the successful quote and define the destination call.

```tsx
  // Inside BasicDepositForm component
  const order = useOrder({
    srcChainId: sourceChainId,
    destChainId: destChainId,
    deposit: {
      amount: quotedDepositAmt, // Use quoted amount
      token: sourceTokenAddress,
    },
    expense: {
      amount: quotedExpenseAmt, // Use quoted amount
      token: destTokenAddress,
      spender: vaultAddress // The vault contract will spend the solver's funds
    },
    calls: [
      {
        target: vaultAddress,
        abi: vaultABI,
        functionName: 'deposit',
        // Pass the user's address and the *quoted expense amount*
        args: [userAddress!, quotedExpenseAmt], // Ensure userAddress is defined
      }
    ],
    // Only enable validation/order if the quote was successful and user is connected
    validateEnabled: quote.isSuccess && isConnected && userAddress != null,
  });

  const {
    open: openOrder, // Rename to avoid conflict if needed
    status: orderStatus,
    validation,
    isReady,
    isTxPending,
    error: orderError,
  } = order;

  // Determine if the button should be enabled
  const canOpen = isReady && validation?.status === 'accepted' && !isTxPending;
```

5.  **Add UI Elements for Feedback and Action:**
    Display quote status, validation status, order status, and the deposit button.

```tsx
  // Inside the return statement of BasicDepositForm
  return (
    <div>
      {/* ... Input field ... */}

      {isConnected && depositAmount > 0n && (
        <>
          {quote.isLoading && <p>Fetching quote...</p>}
          {quote.isError && <p>Quote Error: {quote.error.message}</p>}
          {quote.isSuccess && (
            <p>
              Quote: Deposit {formatEther(quotedDepositAmt)} Base wstETH
              to spend {formatEther(quotedExpenseAmt)} Holesky wstETH.
            </p>
          )}

          {validation?.status === 'pending' && <p>Validating order...</p>}
          {validation?.status === 'rejected' && (
            <p>Order Rejected: {validation.rejectReason} - {validation.rejectDescription}</p>
          )}
          {validation?.status === 'accepted' && <p>✅ Order Validated</p>}

          <button
            disabled={!canOpen}
            onClick={() => openOrder?.()} // Call the open function from useOrder
            style={{ marginTop: '10px' }}
          >
            {isTxPending ? 'Opening Order...' : 'Deposit via Omni'}
          </button>

          <p style={{ marginTop: '10px' }}>Order Status: <strong>{orderStatus}</strong></p>
          {orderError && <p style={{ color: 'red' }}>Order Error: {orderError.message}</p>}
        </>
      )}
    </div>
  );
}

export default BasicDepositForm;
```

## Complete Code Example

```tsx
import React, { useState, useMemo } from 'react';
import { useAccount } from 'wagmi';
import { useQuote, useOrder } from '@omni-network/react';
import { parseEther, formatEther, type Abi, type Address, zeroAddress } from 'viem';
import { baseSepolia, holesky } from 'viem/chains';

// --- Configuration (Replace with your actual values) ---
const vaultABI = [
  {
    inputs: [
      { internalType: 'address', name: 'onBehalfOf', type: 'address' },
      { internalType: 'uint256', name: 'amount', type: 'uint256' },
    ],
    name: 'deposit',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
] as const;
const vaultAddress = '0xYourVaultContractAddressOnHolesky' as const; // Replace
const sourceTokenAddress = '0x6319df7c227e34B967C1903A08a698A3cC43492B' as const;
const destTokenAddress = '0x8d09a4502Cc8Cf1547aD300E066060D043f6982D' as const;
const sourceChainId = baseSepolia.id;
const destChainId = holesky.id;
// ------------------------------------------------------

function BasicDepositForm() {
  const [depositAmountStr, setDepositAmountStr] = useState<string>('0.1');
  const { address: userAddress, isConnected } = useAccount();

  const depositAmount = useMemo(() => {
    try {
      return parseEther(depositAmountStr as `${number}`);
    } catch {
      return 0n;
    }
  }, [depositAmountStr]);

  const quote = useQuote({
    srcChainId: sourceChainId,
    destChainId: destChainId,
    deposit: { isNative: false, token: sourceTokenAddress, amount: depositAmount },
    expense: { isNative: false, token: destTokenAddress },
    mode: "expense",
    enabled: isConnected && depositAmount > 0n,
  });

  const quotedDepositAmt = quote.data?.deposit.amount ?? 0n;
  const quotedExpenseAmt = quote.data?.expense.amount ?? 0n;

  const order = useOrder({
    srcChainId: sourceChainId,
    destChainId: destChainId,
    deposit: {
      amount: quotedDepositAmt,
      token: sourceTokenAddress,
    },
    expense: {
      amount: quotedExpenseAmt,
      token: destTokenAddress,
      spender: vaultAddress,
    },
    calls: [
      {
        target: vaultAddress,
        abi: vaultABI,
        functionName: 'deposit',
        args: [userAddress ?? zeroAddress, quotedExpenseAmt], // Provide default for type safety
      }
    ],
    validateEnabled: quote.isSuccess && isConnected && userAddress != null,
  });

  const {
    open: openOrder,
    status: orderStatus,
    validation,
    isReady,
    isTxPending,
    error: orderError,
  } = order;

  const canOpen = isReady && validation?.status === 'accepted' && !isTxPending;

  return (
    <div>
      <h2>Deposit wstETH from Base Sepolia to Holesky Vault</h2>
      <label>
        Amount to Deposit (wstETH on Base Sepolia):
        <input
          type="text"
          value={depositAmountStr}
          onChange={(e) => setDepositAmountStr(e.target.value)}
          disabled={isTxPending || (orderStatus !== 'initializing' && orderStatus !== 'ready')}
        />
      </label>

      {!isConnected && <p>Please connect your wallet.</p>}

      {isConnected && depositAmount > 0n && (
        <div style={{ marginTop: '15px' }}>
          {quote.isLoading && <p>Fetching quote...</p>}
          {quote.isError && <p>Quote Error: {quote.error.message}</p>}
          {quote.isSuccess && (
            <p>
              Quote: Deposit {formatEther(quotedDepositAmt)} Base wstETH
              to spend {formatEther(quotedExpenseAmt)} Holesky wstETH.
            </p>
          )}

          {validation?.status === 'pending' && <p>Validating order...</p>}
          {validation?.status === 'rejected' && (
            <p>Order Rejected: {validation.rejectReason} - {validation.rejectDescription}</p>
          )}
          {validation?.status === 'accepted' && <p>✅ Order Validated</p>}

          <button
            disabled={!canOpen}
            onClick={() => openOrder?.()}
            style={{ marginTop: '10px' }}
          >
            {isTxPending ? 'Opening Order...' : 'Deposit via Omni'}
          </button>

          <p style={{ marginTop: '10px' }}>Order Status: <strong>{orderStatus}</strong></p>
          {orderError && <p style={{ color: 'red' }}>Order Error: {orderError.message}</p>}
        </div>
      )}
    </div>
  );
}

export default BasicDepositForm;
```

## Key Considerations

*   **Error Handling:** Add more robust error handling for invalid inputs, quote failures, validation rejections, and transaction errors.
*   **User Feedback:** Provide clear feedback to the user during each stage (quoting, validation, opening, tracking status).
*   **Configuration:** Ensure all addresses, ABIs, and chain IDs are correct for your specific use case.
*   **Token Approvals (ERC20 Deposits):** For deposits involving ERC20 tokens (like wstETH in this example), standard token approval prerequisites apply. Before the user can successfully call `openOrder`, your application must ensure that the Omni escrow contract on the source chain has sufficient allowance to transfer the `deposit.amount` of the `deposit.token`. This typically involves checking the current allowance and prompting the user to send an `approve` transaction if the allowance is insufficient.
*   **`userAddress`:** Make sure the `userAddress` is available and passed correctly, especially within the `args` of the destination `call`. Handle the case where the wallet is not connected.
*   **Amount Formatting:** Use `parseEther` and `formatEther` (or equivalent for different decimals) carefully when handling user input and displaying amounts.
