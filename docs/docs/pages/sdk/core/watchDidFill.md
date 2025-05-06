# `watchDidFill`

The `watchDidFill` waits for an order to be `filled` on the destination chain. This involves waiting for a `Filled` event to be emitted on the destination outbox contract, and returns the destination transaction hash.

## Usage

`import { watchDidFill } from '@omni-network/core'`

```tsx
import { watchDidFill } from '@omni-network/core'

const { unwatch, status, destTxHash } = await watchDidFill({
  // ... params
});
```

## Parameters

| Prop              | Type                                   | Required | Description                                                                                                                                |
| -------------     | -------------------------------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `client`          | `Client`                               | Yes      | Chain ID of the destination chain (where the action occurs and the `expense` is spent on behalf of the user).  |
| `outboxAddress`   | `number`                               | Yes      | The address of the inbox contract, retrieved using the [`getContracts` function](/sdk/core/getContracts).  |
| `orderId`         | `0x${string}`                          | Yes      | Order identifier defined on a [`ResolvedOrder`](/sdk/core/waitForOrderOpen#resolvedorder) |
| `onLogs`          | `(logs: Log[]) => void`                | Yes      | Callback that'll be invoked with the logs of the `Filled` event.  |
| `pollingInterval` | `number`                               | No       | Polling interval in milliseconds, defaults to the `client` polling interval.  |
| `onError`         | `(error: Error) => void`               | No       | Optional callback that'll be invoked when an error occurs. |


## Return

`watchDidFill` returns three values:

### `Status`

Describes the state of the watch API.

```ts
'pending' | 'filled' | 'idle'
```

### `destTxHash`

The transaction hash on the destination chain that fills the order.

```ts
`Hex` | `0x${string}`
```

### `unwatch`

Terminates watching of contract events.

```ts
() => void
```

## Example

```ts
import {watchDidFill } from '@omni-network/core'

const { unwatch, status, destTxHash } = await watchDidFill({
  client: viemWalletClient,
  outboxAddress: contracts.outbox,
  orderId: resolvedOrder.orderId,
  onLogs: (logs) => {
    console.log('Filled event found:', logs[0].transactionHash)
  },
})

// Stop watching for the Filled event
unwatch()
```
