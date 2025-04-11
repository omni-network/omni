---
title: Rocket Pool Deposit Demo
---

# Rocket Pool rETH Deposit Example

This example demonstrates depositing ETH from an L2 into Rocket Pool on L1 (Ethereum mainnet or Holesky testnet) using Omni SolverNet.

Since the standard Rocket Pool `deposit` function is payable and credits rETH to `msg.sender`, this example utilizes the [`withExecAndTransfer`](../../sdk/utils/withExecAndTransfer) utility to ensure the user receives the rETH on L1.

:::info Note
This example uses a simplified hook structure (`useL2Deposit`) for demonstration. The core logic involves `useQuote` (implicitly), `useOrder`, and `withExecAndTransfer`. Adapt addresses and chain IDs as needed.
:::

## Custom Hook (`useL2Deposit.ts`)

```typescript
import { useOpenOrder, withExecAndTransfer, useQuote, useOrder, useOmniContracts } from '@omni-network/react' // Assuming hooks are available here
import { zeroAddress } from 'viem'
import { useAccount } from 'wagmi'
import { useMemo } from 'react'

// --- Configuration (Adapt as needed) ---
const ROCKETPOOL_DEPOSIT_ADDRESS = "0x320f3aAB9405e38b955178BBe75c477dECBA0C27" // Holesky Example
const RETH_ADDRESS = "0x7322c24752f79c05FFD1E2a6FCB97020C1C264F1" // Holesky Example

const rocketPoolAbi = [
  {
    type: 'function',
    inputs: [],
    name: 'deposit',
    outputs: [],
    stateMutability: 'payable',
  },
] as const
// -------------------------------------

type UseL2DepositParams = {
  destChainId: number // e.g., holesky.id
  srcChainId: number  // e.g., baseSepolia.id
  depositAmount: bigint // Amount of ETH to deposit from L2
}

export function useL2Deposit(params: UseL2DepositParams) {
  const { address: userAddress, isConnected } = useAccount();

  // 1. Quote: How much ETH expense on L1 for the L2 ETH deposit?
  const quote = useQuote({
    srcChainId: params.srcChainId,
    destChainId: params.destChainId,
    deposit: {
      isNative: true,
      amount: params.depositAmount
    },
    expense: {
      isNative: true // Rocketpool deposit() is payable
    },
    mode: "expense",
    enabled: isConnected && params.depositAmount > 0n,
  });

  const quotedDepositAmt = quote.data?.deposit.amount ?? 0n;
  const quotedExpenseAmt = quote.data?.expense.amount ?? 0n;

  // 2. Get Middleman Address
  const contracts = useOmniContracts(params.destChainId);
  const middlemanAddress = contracts.data?.middleman ?? zeroAddress;

  // 3. Prepare Wrapped Call
  const middlemanCall = useMemo(() => {
    if (!userAddress || middlemanAddress === zeroAddress || !quote.isSuccess) return undefined;

    const originalRocketpoolCall = {
        target: ROCKETPOOL_DEPOSIT_ADDRESS,
        value: quotedExpenseAmt, // ETH value for the deposit
        abi: rocketPoolAbi,
        functionName: 'deposit',
    };

    return withExecAndTransfer({
        middlemanAddress: middlemanAddress,
        call: originalRocketpoolCall,
        transfer: {
            to: userAddress, // Final recipient of rETH
            token: RETH_ADDRESS, // The token credited by Rocketpool deposit
        },
    });
  }, [userAddress, middlemanAddress, quotedExpenseAmt, quote.isSuccess]);

  // 4. Configure Order
  const order = useOrder({
    srcChainId: params.srcChainId,
    destChainId: params.destChainId,
    deposit: {
      isNative: true,
      amount: quotedDepositAmt,
    },
    expense: {
      isNative: true,
      amount: quotedExpenseAmt,
      spender: middlemanAddress, // Middleman spends the ETH
    },
    calls: middlemanCall ? [middlemanCall] : [],
    validateEnabled: quote.isSuccess && contracts.isSuccess && middlemanCall != null,
  });

  // Expose order controls and status
  return {
    mutation: order.txMutation, // For UI feedback on source tx
    waitOrder: order.waitForTx,  // For source tx confirmation
    orderStatus: order.status,
    openL2Deposit: order.open, // Function to trigger the deposit
    isReady: order.isReady && order.validation?.status === 'accepted',
    validation: order.validation,
    quote,
  };
}

```

## Usage in Component

```tsx
import React, { useState } from 'react';
import { useL2Deposit } from './useL2Deposit'; // Import the custom hook
import { parseEther, formatEther } from 'viem';
import { baseSepolia, holesky } from 'viem/chains';
import { useAccount } from 'wagmi';

function RocketpoolDepositComponent() {
  const [amountStr, setAmountStr] = useState('0.01');
  const { isConnected } = useAccount();

  const depositAmount = useMemo(() => {
    try {
      return parseEther(amountStr as `${number}`);
    } catch {
      return 0n;
    }
  }, [amountStr]);

  const {
    openL2Deposit,
    orderStatus,
    isReady,
    validation,
    quote
  } = useL2Deposit({
    srcChainId: baseSepolia.id,
    destChainId: holesky.id,
    depositAmount: depositAmount,
  });

  const quotedDepositAmt = quote.data?.deposit.amount ?? 0n;
  const quotedExpenseAmt = quote.data?.expense.amount ?? 0n;

  const isLoading = orderStatus === 'opening' || orderStatus === 'open';

  return (
    <div style={{ border: '1px solid #ccc', padding: '1rem', marginTop: '1rem' }}>
      <h2>Rocket Pool Deposit (L2 ETH -&gt; L1 rETH)</h2>
      <label>
        Amount ETH to Deposit (from Base Sepolia):
        <input
          type="text"
          value={amountStr}
          onChange={(e) => setAmountStr(e.target.value)}
          disabled={isLoading || !isConnected}
        />
      </label>

      {!isConnected && <p>Connect wallet to proceed.</p>}

      {isConnected && depositAmount > 0n && (
        <>
          {quote.isLoading && <p>Getting quote...</p>}
          {quote.isSuccess && <p>Quote: Deposit {formatEther(quotedDepositAmt)} ETH to receive rETH for {formatEther(quotedExpenseAmt)} ETH value.</p>}
          {quote.isError && <p style={{color: 'red'}}>Quote Error: {quote.error.message}</p>}

          {validation?.status === 'pending' && <p>Validating...</p>}
          {validation?.status === 'rejected' && <p style={{color: 'red'}}>Rejected: {validation.rejectDescription}</p>}

          <button onClick={openL2Deposit} disabled={!isReady || isLoading} style={{marginTop: '1rem'}}>
            {isLoading ? 'Processing...' : 'Deposit via Omni'}
          </button>
          <p style={{ marginTop: '1rem' }}>Status: <strong>{orderStatus}</strong></p>
          {orderStatus === 'filled' && <p style={{color: 'green'}}>✅ Success!</p>}
        </>
      )}
    </div>
  );
}

export default RocketpoolDepositComponent;
```

## Key Concepts

*   **`withExecAndTransfer`:** Essential because Rocket Pool's `deposit` sends rETH to `msg.sender`.
*   **Middleman:** The hook fetches the `SolverNetMiddleman` address on the destination chain (L1).
*   **Wrapped Call:** The `deposit` call is wrapped, specifying the `rETH` address as the token to be transferred to the user after the deposit.
*   **Order Configuration:**
    *   `deposit`: Native ETH from the source L2.
    *   `expense`: Native ETH required by the `deposit` function on L1, with the `spender` set to the `middlemanAddress`.
    *   `calls`: Contains the single wrapped call generated by `withExecAndTransfer`.
