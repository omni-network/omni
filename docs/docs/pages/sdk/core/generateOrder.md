# `generateOrder`

The `generateOrder` function implements a common flow to open and track the status of an order on a SolverNet inbox contract, by combining single-purpose functions:

1. Validating the order, using [`validateOrder`](/sdk/core/validateOrder)
2. Sending the order transaction, using [`sendOrder`](/sdk/core/sendOrder)
3. Waiting for the order to be open on the blockchain, using [`waitForOrderOpen`](/sdk/core/waitForOrderOpen)
4. Waiting for the order reach a terminal state, using [`waitForOrderClose`](/sdk/core/waitForOrderClose)

## Usage

`import { generateOrder } from '@omni-network/core'`

```tsx
import { generateOrder } from '@omni-network/core'

const generator = await generateOrder({
  // ... params
});
```

## Arguments

The `generateOrder` function uses a single parameters object as argument:

| Prop                | Type                                 | Required | Description                                                                                                                         |
| ------------------- | ------------------------------------ | -------- | ----------------------------------------------------------------------------------------------------------------------------------- |
| `client`        | `Client`                             | Yes      | The `viem` client used to send the transaction. This client must have an `account` attached to it.                                                                          |
| `inboxAddress`       | `Address`                             | Yes      | The address of the inbox contract, retrieved using the [`getContracts` function](/sdk/core/getContracts).                                                                     |
| `order`           | `Order`                         | Yes      | [Order parameters](/sdk/core/validateOrder#1-order-parameters-required) |
| `pollinginterval`       | `number`                             | No      | Polling interval in milliseconds, defaults to the `client` polling interval.                                                                     |
| `environment`           | `Environment | string`                         | No      | SolverNet environment to use, either `mainnet` (default) or `testnet`. |

## Return

`generateOrder` returns an [AsyncGenerator](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/AsyncGenerator) of `OrderState` objects through the order flow.

### `OrderState`

```ts
type OrderState =
  | { status: 'valid'; txHash?: never; order?: never }
  | { status: 'sent'; txHash: Hex; order?: never }
  | { status: 'open'; txHash: Hex; order: ResolvedOrder }
  | { status: TerminalStatus; txHash: Hex; order: ResolvedOrder }
```

## Example

```ts
import { type OrderState, getContracts, generateOrder } from '@omni-network/core'

const contracts = await getContracts()
const generator = generateOrder({
  client: viemWalletClient,
  inboxAddress: contracts.inbox,
  order: orderParams,
})

let orderState: OrderState | undefined
for await (orderState of generator) {
  console.log('current order state', orderState)
}
console.log('terminal order state', orderState)
```
