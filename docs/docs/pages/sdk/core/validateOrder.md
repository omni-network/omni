# `validateOrder`

The `validateOrder` function validates order parameters with the SolverNet server to check if an order could be handled by SolverNet.

## Usage

`import { validateOrder } from '@omni-network/core'`

```tsx
import { validateOrder } from '@omni-network/core'

const validation = await validateOrder({
  // ... params
});
```

## Arguments

The `validateOrder` function uses the following arguments:

### 1. Order parameters (required)

| Prop                | Type                                 | Required | Description                                                                                                                         |
| ------------------- | ------------------------------------ | -------- | ----------------------------------------------------------------------------------------------------------------------------------- |
| `srcChainId`        | `number`                             | Yes      | The chain ID of the source chain. Must match `getQuote`.                                                                          |
| `destChainId`       | `number`                             | Yes      | The chain ID of the destination chain. Must match `getQuote`.                                                                     |
| `deposit`           | `Deposit`                         | Yes      | Describes the asset and amount being deposited on the source chain (paid by the user) - taken from `quote.deposit`. |
| `expense`           | `Expense`  | Yes      | Describes the asset, amount, and spender on the destination chain (paid by the solver) - taken from `quote.expense`.            |
| `calls`             | `Call[]`                             | Yes      | An array of contract calls to be executed on the destination chain by the solver.                                                     |
| `validateEnabled`   | `boolean`                            | No       | Defaults to `true`. Enables pre-validation of the order with Omni. Use this to validate your calls. Recommended to set based on `quote.isSuccess`.

### 2. Environment (optional)

SolverNet environment to use, either `mainnet` (default) or `testnet`.

## Types

### `Deposit`

Describes the deposit parameter.

```typescript
type Deposit = {
  readonly token?: Address
  readonly amount: bigint
}
```

### `Expense`

Describes the expense parameter.

```typescript
type Expense = {
  readonly spender?: Address
  readonly token?: Address
  readonly amount: bigint
}
```

### `Order`

Describes the order to placed by the user.

```typescript
export type Order<abis extends OptionalAbis> = {
  readonly owner?: Address
  readonly srcChainId?: number
  readonly destChainId: number
  readonly fillDeadline?: number
  readonly calls: Calls<abis>
  readonly deposit: Deposit
  readonly expense: Expense
}
```

### `Call`

Describes a contract interaction on the destination chain.

```typescript
type Call = {
  target: Address;      // The contract address to call
  abi: Abi;             // The ABI of the target contract
  functionName: string; // The function to call - type inferred from the abi
  args?: unknown[];     // Arguments for the function call - type inferred from the abi
  value?: bigint;       // ETH value to send with the call (optional)
}
```

## Return

`validateOrder` returns the [Promise](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise) of a `ValidationResponse` object.

Note that the Promise will resolve based on receiving a success HTTP response from the SolverNet server. This response may indicate that the order is invalid.

### `ValidationResponse`

```typescript
type ValidationResponse = {
  accepted?: boolean
  rejected?: boolean
  error?: {
    code: number
    message: string
  }
  rejectReason?: string
  rejectDescription?: string
}
```

## Examples

### Successful validation

```ts
import { validateOrder } from '@omni-network/core'

const validation = await validateOrder({...})
if (validation.accepted) {
  // Order parameters are correct
}
```

### Rejected validation

```ts
import { validateOrder } from '@omni-network/core'

const validation = await validateOrder({...})
if (validation.rejected) {
  console.error(`Validation rejected order with reason: ${validation.rejectReason} - ${validation.rejectDescription}`)
}
```

### Server error

```ts
import { validateOrder } from '@omni-network/core'

const validation = await validateOrder({...})
if (validation.error) {
  console.error(`Server error ${validation.error.code} attempting to validate order: ${validation.error.message}`)
}
```

### Other exceptions

```ts
import { validateOrder } from '@omni-network/core'

try {
  const validation = await validateOrder({...})
} catch (e) {
  // Network error or other exception
}
```
