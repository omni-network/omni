---
sidebar_position: 1
title: useQuote
---

# `useQuote`

The `useQuote` hook is your first step in using Omni SolverNet. It fetches the real-time cost for a potential cross-chain action by calculating the relationship between:

*   The **Deposit**: The asset and amount **the user pays** on the **source chain** to initiate the action.
*   The **Expense**: The asset and amount the **solver spends** on the **destination chain** to execute your desired action (e.g., the amount deposited into a target vault).

Based on the `mode` you select, the hook calculates either:

1.  The required `deposit` for a fixed `expense`.
2.  The output `expense` for a fixed `deposit`.

This quote accounts for solver fees and current market conditions.

## Usage

Import the hook from `@omni-network/react`:

```tsx
import { useQuote } from '@omni-network/react'
import { parseEther } from 'viem'
import { baseSepolia, holesky } from 'viem/chains'

// Example token addresses
const holeskyWSTETH = '0x8d09a4502Cc8Cf1547aD300E066060D043f6982D' as const
const baseSepoliaWSTETH = '0x6319df7c227e34B967C1903A08a698A3cC43492B' as const

function MyComponent() {
  const quote = useQuote({
    // ... configuration
  });

  // ... rest of component
}
```

## Parameters

Crucially, a quote requires fixing *either* the `deposit` amount *or* the `expense` amount. You tell the hook which one you're fixing using the `mode` parameter.

The hook accepts a configuration object with the following properties.

| Prop          | Type                                   | Required | Description                                                                                                                                |
| ------------- | -------------------------------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `srcChainId`  | `number`                               | Yes      | The chain ID of the source network (where the user initiates the action and provides the `deposit`).                                       |
| `destChainId` | `number`                               | Yes      | The chain ID of the destination network (where the action occurs and the `expense` is spent).                                               |
| `deposit`     | `QuoteAsset & { amount?: bigint }`     | Yes      | Describes the asset being deposited on the source chain. Provide `amount` only if using `mode: 'expense'`.                                     |
| `expense`     | `QuoteAsset & { amount?: bigint }`     | Yes      | Describes the asset being spent on the destination chain. Provide `amount` only if using `mode: 'deposit'`.                                     |
| `mode`        | `'deposit' \| 'expense'`                | Yes      | Determines which amount to calculate: <br/>- `'deposit'`: Provide fixed `expense.amount`, calculate `deposit.amount`. <br/>- `'expense'`: Provide fixed `deposit.amount`, calculate `expense.amount`. |
| `enabled`     | `boolean`                              | No       | Defaults to `true`. Set to `false` to disable fetching the quote.                                                                          |

### `QuoteAsset` Type

```typescript
type QuoteAsset =
  | { isNative: true; token?: never } // For native ETH
  | { isNative?: false; token: Address } // For ERC20 tokens
```

## Examples

### Quoting Expense Amount

To find out how much `wstETH` can be spent on Holesky for a deposit of 0.1 `wstETH` from Base Sepolia:

```tsx
const quote = useQuote({
  srcChainId: baseSepolia.id,
  destChainId: holesky.id,
  deposit: { isNative: false, token: baseSepoliaWSTETH, amount: parseEther("0.1") },
  expense: { isNative: false, token: holeskyWSTETH }, // No amount specified
  mode: "expense", // Calculate expense amount
  enabled: true, // Fetch quote when component mounts
})

if (quote.isSuccess) {
  console.log(`Depositing ${quote.data.deposit.amount} yields ${quote.data.expense.amount} on destination`);
}
```

### Quote Deposit Amount

To find out how much `wstETH` needs to be deposited on Base Sepolia to spend exactly 0.1 `wstETH` on Holesky:

```tsx
const quote = useQuote({
  srcChainId: baseSepolia.id,
  destChainId: holesky.id,
  deposit: { isNative: false, token: baseSepoliaWSTETH }, // No amount specified
  expense: { isNative: false, token: holeskyWSTETH, amount: parseEther("0.1") },
  mode: "deposit", // Calculate deposit amount
  enabled: true,
})

if (quote.isSuccess) {
  console.log(`Spending ${quote.data.expense.amount} requires depositing ${quote.data.deposit.amount} on source`);
}
```

## Return Value

`useQuote` returns a standard query result object from [`@tanstack/react-query`](https://tanstack.com/query/latest/docs/react/reference/useQuery). Consult their documentation for all available properties.

Key properties include:

*   `data`: The quote result if the query is successful.
    *   `deposit`: `{ amount: bigint, token?: Address, isNative: boolean }`
    *   `expense`: `{ amount: bigint, token?: Address, isNative: boolean }`
*   `isSuccess`: `true` if the quote was fetched successfully.
*   `isLoading`: `true` while the quote is being fetched.
*   `isError`: `true` if there was an error fetching the quote.
*   `error`: The error object if `isError` is `true`.

Use the `data` property (specifically `data.deposit.amount` and `data.expense.amount`) to inform the parameters for the [`useOrder`](/sdk/hooks/useOrder.md) hook.
