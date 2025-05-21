---
sidebar_position: 1
title: useQuote
---

# `useQuote`

The `useQuote` hook is your first step in using Omni SolverNet. It fetches the real-time cost for a potential cross-chain action by calculating the relationship between:

*   **Deposit**: The asset and amount **the user pays** on the **source chain** to initiate the action.
*   **Expense**: The asset and amount the **solver spends** on the **destination chain** to execute your desired action (e.g., the amount deposited into a target vault).

Based on the `mode` you select, the hook calculates either:

1.  The required `deposit` for a fixed `expense`.
2.  The output `expense` for a fixed `deposit`.

This quote accounts for solver fees and current market conditions.

## Usage

`import { useQuote } from '@omni-network/react'`

```tsx
import { useQuote } from '@omni-network/react'

function Component() {
  const quote = useQuote({
    // ... params
  });
}
```

## Parameters

Crucially, a quote requires fixing *either* the `deposit` amount *or* the `expense` amount. You tell the hook which one you're fixing using the `mode` parameter.

The hook accepts a configuration object with the following properties.

| Prop          | Type                                   | Required | Description                                                                                                                                |
| ------------- | -------------------------------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `srcChainId`  | `number`                               | Yes      | Chain ID of the source chain (where the user provides the `deposit`).                                       |
| `destChainId` | `number`                               | Yes      | Chain ID of the destination chain (where the action occurs and the `expense` is spent).                                               |
| `deposit`     | `{ token?: Address; amount?: bigint }` | Yes if `mode: 'expense'` | Asset to deposit on the source chain. Provide `amount` only if using `mode: 'expense'`. Omit `token` or supply zero address if using native token (e.g. ETH).                                    |
| `expense`     | `{ token?: Address; amount?: bigint }` | Yes if `mode: 'deposit'` | Asset to spend on the destination chain. Provide `amount` only if using `mode: 'deposit'`. Omit `token` or supply zero address if using native token (e.g. ETH).                                     |
| `mode`        | `'deposit' \| 'expense'`               | Yes      | Defines the direction of the quote. |
| `enabled`     | `boolean`                              | No       | Defaults to `true`. Set to `false` to disable fetching the quote.                                                                          |
| `queryOpts`     | `UseQueryOptions<Quote, QuoteError>` | No       | React query options, we omit some keys here (`enabled`, `queryKey`, and `queryFn`) to prevent overriding some default behaviour. See [`useQuery`](https://tanstack.com/query/latest/docs/react/reference/useQuery) docs for available options. |

## Types

### Quote

```typescript
export type Quote = {
  deposit: { token: Address; amount: bigint }
  expense: { token: Address; amount: bigint }
}
```

Describes a successful quote return.

## Return

`useQuote` returns a quote and the query object from [`@tanstack/react-query`](https://tanstack.com/query/latest/docs/react/reference/useQuery). Consult their documentation for all available properties.

### isPending

`boolean`

Is `true` before a quote response is received.

### isSuccess

`boolean`

Is `true` when the quote call succeeds.

### isError

`boolean`

Is `true` when the quote call fails.

### quote

`{ deposit: { token: Address; amount: bigint } expense: { token: Address; amount: bigint } }`

Only available when `isSuccess` is `true`.

```typescript
if (res.isSuccess) res.quote.deposit
```

### error

Only available when `isError` is `true`.

```typescript
if (res.isError) res.error
```

Use the `quote` return (specifically `quote.deposit.amount` and `quote.expense.amount`) to inform the parameters for [`useOrder`](/sdk/hooks/useOrder.mdx).

## Examples

### Quote Expense

To find out how much `wstETH` can be spent on Holesky for a deposit of 0.1 `wstETH` from Base Sepolia:

```tsx
const quote = useQuote({
  srcChainId: baseSepolia.id,
  destChainId: holesky.id,
  deposit: { token: baseSepoliaWSTETH, amount: parseEther("0.1") },
  expense: { token: holeskyWSTETH }, // note - when mode: "expense" we don't supply expense.amount
  mode: "expense", // quote expense amount
  enabled: true,
})

if (quote.isSuccess) {
  console.log(`Depositing ${quote.deposit.amount} yields ${quote.expense.amount} on destination`);
}
```

### Quote Deposit

To find out how much `wstETH` needs to be deposited on Base Sepolia to spend exactly 0.1 `wstETH` on Holesky:

```tsx
const quote = useQuote({
  srcChainId: baseSepolia.id,
  destChainId: holesky.id,
  deposit: { token: baseSepoliaWSTETH }, // note - when mode: "expense" we don't supply expense.amount
  expense: { token: holeskyWSTETH, amount: parseEther("0.1") },
  mode: "deposit", // quote deposit amount
  enabled: true,
})

if (quote.isSuccess) {
  console.log(`Spending ${quote.data.expense.amount} requires depositing ${quote.data.deposit.amount} on source`);
}
```
