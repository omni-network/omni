---
sidebar_position: 2
title: useOrder
---

# `useOrder`

After obtaining a valid quote with [`useQuote`](/sdk/hooks/useQuote), the `useOrder` hook is used to execute the actual cross-chain transaction via Omni SolverNet. It takes the details confirmed in the quote, validates them, and provides a function (`open`) to initiate the process and monitor its status.

## High-Level Usage Flow

```tsx
import { useQuote, useOrder } from '@omni-network/react';
// ... other imports (chains, ABIs, addresses)

function InitiateAction() {
  // 1. Get the quote
  const quote = useQuote({ /* ... config ... */ });

  // 2. Configure the order using the quote data
  const order = useOrder({
    srcChainId: /* ... */,
    destChainId: /* ... */,
    deposit: quote.data?.deposit, // Use quoted deposit
    expense: {
        ...quote.data?.expense, // Use quoted expense
        spender: /* Address needing approval on dest chain */
    },
    calls: [ /* ... contract call(s) on dest chain ... */ ],
    validateEnabled: quote.isSuccess, // when true, this will check if the order will be accepted by Omni. You can consume the result via order.validation

  });

  // 3. Open the order when ready and validated
  const canOpen = order.isReady && order.validation?.status === 'accepted';
  const handleOpen = () => {
    if (canOpen) {
      order.open?.(); // Call the open function provided by the hook
    }
  };

  // 4. Monitor order.status for UI updates
  // ... Render button, status messages, etc.
}
```

## Configuration

The hook accepts a configuration object:

| Prop                | Type                                 | Required | Description                                                                                                                         |
| ------------------- | ------------------------------------ | -------- | ----------------------------------------------------------------------------------------------------------------------------------- |
| `srcChainId`        | `number`                             | Yes      | The chain ID of the source network. Must match `useQuote`.                                                                          |
| `destChainId`       | `number`                             | Yes      | The chain ID of the destination network. Must match `useQuote`.                                                                     |
| `deposit`           | `OrderAsset`                         | Yes      | Describes the asset and amount being deposited on the source chain (paid by the connected user). Typically from `quote.data.deposit`. |
| `expense`           | `OrderAsset & { spender: Address }`  | Yes      | Describes the asset, amount, and spender on the destination chain (paid by the solver). Typically from `quote.data.expense`.            |
| `calls`             | `Call[]`                             | Yes      | An array of contract calls to be executed on the destination chain by the solver.                                                     |
| `validateEnabled`   | `boolean`                            | No       | Defaults to `true`. Enables pre-validation of the order with Omni. Recommended to set based on `quote.isSuccess`.                    |

### `OrderAsset` Type

```typescript
type OrderAsset =
  | { isNative: true; token?: never; amount: bigint } // Native ETH
  | { isNative?: false; token: Address; amount: bigint } // ERC20
```

### `Call` Type

Describes a contract interaction on the destination chain.

```typescript
import { type Abi } from 'viem'

type Call = {
  target: Address;       // The contract address to call
  abi: Abi;             // The ABI of the target contract
  functionName: string; // The function to call
  args?: any[];          // Arguments for the function call
  value?: bigint;        // ETH value to send with the call (optional)
}
```

## Example

Building on the `useQuote` example, let's define and use `useOrder` to deposit into a vault:

```tsx
import { useOrder, useQuote } from '@omni-network/react'
import { parseEther, type Abi } from 'viem'
import { baseSepolia, holesky } from 'viem/chains'

// Vault.deposit(address onBehalfOf, uint256 amount)
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
] as const

const vaultAddress = `0x...` as const // Your vault address
const userAddress = '0x...' as const // Address to deposit on behalf of

const holeskyWSTETH = '0x8d09a4502Cc8Cf1547aD300E066060D043f6982D' as const
const baseSepoliaWSTETH = '0x6319df7c227e34B967C1903A08a698A3cC43492B' as const

function DepositComponent() {
    // 1. Get quote
    const quote = useQuote({
        srcChainId: baseSepolia.id,
        destChainId: holesky.id,
        deposit: { isNative: false, token: baseSepoliaWSTETH, amount: parseEther("0.1") },
        expense: { isNative: false, token: holeskyWSTETH },
        mode: "expense",
        enabled: true,
    })

    // Get quoted amounts, defaulting to 0 if quote is not ready
    const depositAmt = quote.data?.deposit.amount ?? 0n
    const expenseAmt = quote.data?.expense.amount ?? 0n

    // 2. Configure the order
    const order = useOrder({
        srcChainId: baseSepolia.id,
        destChainId: holesky.id,
        deposit: {
            amount: depositAmt,
            token: baseSepoliaWSTETH
        },
        expense: {
            amount: expenseAmt,
            token: holeskyWSTETH,
            spender: vaultAddress
        },
        calls: [
          {
            target: vaultAddress,
            abi: vaultABI,
            functionName: 'deposit',
            args: [userAddress, expenseAmt]
          }
        ],
        validateEnabled: quote.isSuccess
    })

    const { open, status, validation, isReady, isTxPending } = order

    // Determine if the order can be opened
    const canOpen = isReady && validation?.status === 'accepted'

    return (
        <div>
            <h3>Deposit 0.1 Base Sepolia wstETH to Holesky Vault</h3>
            {quote.isLoading && <p>Getting quote...</p>}
            {quote.isError && <p>Error getting quote: {quote.error.message}</p>}
            {quote.isSuccess && (
              <p>Quoted: Deposit {depositAmt.toString()}, Spend {expenseAmt.toString()}</p>
            )}

            {validation?.status === 'pending' && <p>Validating order...</p>}
            {validation?.status === 'rejected' && (
              <p>Order rejected: {validation.rejectReason} - {validation.rejectDescription}</p>
            )}
            {validation?.status === 'accepted' && <p>Order validated successfully!</p>}

            <button disabled={!canOpen || isTxPending} onClick={() => open?.()}>
              {isTxPending ? 'Opening...' : 'Deposit'}
            </button>

            <p>Order Status: {status}</p>
        </div>
    )
}
```

## Return Value

The `useOrder` hook returns an object with the following properties:

| Property        | Type                                  | Description                                                                                                                            |
| --------------- | ------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `open`          | `(() => void) \| undefined`           | Function to initiate the order opening process (sends transaction on source chain). Undefined until the order is ready (`isReady`).          |
| `status`        | `OrderStatus`                         | The current status of the order lifecycle.                                                                                             |
| `orderId`       | `string \| undefined`                 | The unique identifier for the order, available once `status` is `'open'`.                                                              |
| `isOpen`        | `boolean`                             | `true` if the order has been successfully opened (`status === 'open'`).                                                                |
| `validation`    | `ValidationResult \| undefined`       | The result of the pre-validation check if `validateEnabled` is true.                                                                 |
| `isValidated`   | `boolean`                             | `true` if the validation check has completed (regardless of outcome).                                                                    |
| `error`         | `Error \| null`                        | Contains error information if `status` is `'error'`.                                                                                   |
| `isError`       | `boolean`                             | `true` if `status` is `'error'`.                                                                                                       |
| `isTxPending`   | `boolean`                             | `true` while the source chain transaction for opening the order is being submitted (part of `'opening'` status).                       |
| `isTxSubmitted` | `boolean`                             | `true` once the source chain transaction has been submitted (part of `'opening'` status).                                              |
| `txMutation`    | `UseMutationResult<...>`              | The raw mutation result object from `wagmi` for the opening transaction.                                                               |
| `txHash`        | `Hex \| undefined`                   | The hash of the source chain transaction used to open the order.                                                                       |
| `waitForTx`     | `UseWaitForTransactionReceiptReturnType` | The result object from `wagmi`'s `useWaitForTransactionReceipt` for the opening transaction.                                           |
| `isReady`       | `boolean`                             | `true` when the hook is initialized and ready to attempt opening the order (status is `'ready'`). Does *not* imply validation passed. |

### `OrderStatus` Type

```typescript
export type OrderStatus =
  | 'initializing'
  | 'ready'
  | 'opening'
  | 'open'
  | 'closed'
  | 'rejected'
  | 'error'
  | 'filled'
```

### `ValidationResult` Type

```typescript
type ValidationResult = {
  status: 'pending' | 'rejected' | 'accepted';
  rejectReason?: string;
  rejectDescription?: string;
}
```

Monitor the `status` property to provide feedback to the user throughout the cross-chain operation lifecycle.
