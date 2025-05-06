# `getQuote`

The `getQuote` function fetches the real-time cost for a potential cross-chain action by calculating the relationship between:

*   **Deposit**: The asset and amount **the user pays** on the **source chain** to initiate the action.
*   **Expense**: The asset and amount the **solver spends** on the **destination chain** to execute your desired action (e.g., the amount deposited into a target vault).

Based on the `mode` you select, the function calculates either:

1.  The required `deposit` for a fixed `expense`.
2.  The output `expense` for a fixed `deposit`.

This quote accounts for solver fees and current market conditions.

## Usage

`import { getQuote } from '@omni-network/core'`

```tsx
import { getQuote } from '@omni-network/core'

const quote = await getQuote({
  // ... params
});
```

## Parameters

Crucially, a quote requires fixing *either* the `deposit` amount *or* the `expense` amount. You tell the function which one you're fixing using the `mode` parameter.

The function accepts a configuration object with the following properties.

| Prop          | Type                                   | Required | Description                                                                                                                                |
| ------------- | -------------------------------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `srcChainId`  | `number`                               | Yes      | Chain ID of the source chain (where the user provides the `deposit`).                                       |
| `destChainId` | `number`                               | Yes      | Chain ID of the destination chain (where the action occurs and the `expense` is spent).                                               |
| `deposit`     | `QuoteAsset & { amount?: bigint }`     | Yes      | Aasset to deposit on the source chain. Provide `amount` only if using `mode: 'expense'`.                                     |
| `expense`     | `QuoteAsset & { amount?: bigint }`     | Yes      | Asset to spend on the destination chain. Provide `amount` only if using `mode: 'deposit'`.                                     |
| `mode`        | `'deposit' \| 'expense'`               | Yes      | Defines the direction of the quote. |
| `environment`           | `Environment | string`                         | No      | SolverNet environment to use, either `mainnet` (default) or `testnet`. |

## Types

### `QuoteAsset`

```typescript
type QuoteAsset =
  | { isNative: true; token?: never } // For native ETH
  | { isNative?: false; token: Address } // For ERC20 tokens
```

Describes `deposit` and `expense` paramaters shape.

### Quote

```typescript
export type Quote = {
  deposit: { token: Address; amount: bigint }
  expense: { token: Address; amount: bigint }
}
```

Describes a successful quote return.

## Return

`getQuote` returns the [Promise](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise) of a `quote` object

### quote

`{ deposit: { token: Address; amount: bigint }; expense: { token: Address; amount: bigint } }`

## Examples

### Quote Expense

To find out how much `wstETH` can be spent on Holesky for a deposit of 0.1 `wstETH` from Base Sepolia:

```tsx
const quote = await getQuote({
  srcChainId: baseSepolia.id,
  destChainId: holesky.id,
  deposit: { isNative: false, token: baseSepoliaWSTETH, amount: parseEther("0.1") },
  expense: { isNative: false, token: holeskyWSTETH }, // note - when mode: "expense" we don't supply expense.amount
  mode: "expense", // quote expense amount
})

console.log(`Depositing ${quote.deposit.amount} yields ${quote.expense.amount} on destination`)
```

### Quote Deposit

To find out how much `wstETH` needs to be deposited on Base Sepolia to spend exactly 0.1 `wstETH` on Holesky:

```tsx
const quote = await getQuote({
  srcChainId: baseSepolia.id,
  destChainId: holesky.id,
  deposit: { isNative: false, token: baseSepoliaWSTETH }, // note - when mode: "expense" we don't supply expense.amount
  expense: { isNative: false, token: holeskyWSTETH, amount: parseEther("0.1") },
  mode: "deposit", // quote deposit amount
})

console.log(`Spending ${quote.expense.amount} requires depositing ${quote.deposit.amount} on source`);
```
