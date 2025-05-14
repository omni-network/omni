# `openOrder`

The `openOrder` function implements a common flow to open an order on a SolverNet inbox contract, by combining single-purpose functions:

1. Validating the order, using [`validateOrder`](/sdk/core/validateOrder.md)
2. Sending the order transaction, using [`sendOrder`](/sdk/core/sendOrder.mdx)
3. Waiting for the order to be open on the blockchain, using [`waitForOrderOpen`](/sdk/core/waitForOrderOpen.md)

## Usage

`import { openOrder } from '@omni-network/core'`

```tsx
import { openOrder } from '@omni-network/core'

const resolvedOrder = await openOrder({
  // ... params
});
```

## Parameters

| Prop                | Type                                 | Required | Description                                                                                                                         |
| ------------------- | ------------------------------------ | -------- | ----------------------------------------------------------------------------------------------------------------------------------- |
| `client`        | `Client`                             | Yes      | The `viem` client used to send the transaction. This client must have an `account` attached to it.                                                                          |
| `inboxAddress`       | `Address`                             | Yes      | The address of the inbox contract, retrieved using the [`getContracts` function](/sdk/core/getContracts).                                                                     |
| `order`           | `Order`                         | Yes      | [Order parameters](/sdk/core/validateOrder#1-order-parameters-required) |
| `pollinginterval`       | `number`                             | No      | Polling interval in milliseconds, defaults to the `client` polling interval.                                                                     |
| `environment`           | `Environment`                         | No      | SolverNet environment to use, either `mainnet` (default) or `testnet`. |

## Return

`openOrder` returns the [Promise](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise) of the created transaction hash as a [`ResolvedOrder` object](/sdk/core/waitForOrderOpen#resolvedorder).

## Example

```ts
import { getContracts, openOrder } from '@omni-network/core'

const orderParams = {...}

const contracts = await getContracts()
const resolvedOrder = await openOrder({
  client: viemWalletClient,
  inboxAddress: contracts.inbox,
  order: orderParams,
})
```
