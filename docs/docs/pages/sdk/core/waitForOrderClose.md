# `waitForOrderClose`

The `waitForOrderClose` waits for an order to be in a terminal state (`closed`, `filled` or `rejected`) on a SolverNet inbox contract.

## Usage

`import { waitForOrderClose } from '@omni-network/core'`

```tsx
import { waitForOrderClose } from '@omni-network/core'

const orderStatus = await waitForOrderClose({
  // ... params
});
```

## Parameters

| Prop                | Type                                 | Required | Description                                                                                                                         |
| ------------------- | ------------------------------------ | -------- | ----------------------------------------------------------------------------------------------------------------------------------- |
| `client`        | `Client`                             | Yes      | The `viem` client used to read from the blockchain.                                                                         |
| `inboxAddress`       | `Address`                             | Yes      | The address of the inbox contract, retrieved using the [`getContracts` function](/sdk/core/getContracts).                                                                     |
| `orderId`           | `Hex`                         | Yes      | Order identifier defined on a [`ResolvedOrder`](/sdk/core/waitForOrderOpen#resolvedorder) |
| `pollinginterval`       | `number`                             | No      | Polling interval in milliseconds, defaults to the `client` polling interval.                                                                     |
| `signal`       | `AbortSignal`                             | No      | Optional [`AbortSignal`](https://developer.mozilla.org/en-US/docs/Web/API/AbortSignal) that can be used to stop waiting.                                                                    |

## Return

`waitForOrderClose` returns the [Promise](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise) of a `TerminalStatus` string.

### `TerminalStatus`

```ts
type TerminalStatus = 'closed' | 'filled' | 'rejected'
```

## Example

```ts
import { sendOrder, waitForOrderClose, waitForOrderOpen } from '@omni-network/core'

const txHash = await sendOrder({
  client: viemClient,
  inboxAddress: contracts.inbox,
  order: orderParams
})
const resolvedOrder = await waitForOrderOpen({
  client: viemClient,
  txHash: txHash,
})
const orderStatus = await waitForOrderClose({
  client: viemClient,
  inboxAddress: contracts.inbox,
  orderId: resolvedOrder.orderId,
})
```
