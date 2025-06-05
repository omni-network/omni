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
| `client`          | `Client`                               | Yes      | The `viem` client used to send the transaction. This client must have an `account` attached to it.  |
| `outboxAddress`   | `Hex`                               | Yes      | The address of the outbox contract, retrieved using the [`getContracts` function](/sdk/core/getContracts).  |
| `orderId`         | `Hex`                          | Yes      | Order identifier defined on a [`ResolvedOrder`](/sdk/core/waitForOrderOpen#resolvedorder) |
| `onFill`          | `(txHash: Hex) => void`                | Yes      | Callback that'll be invoked with the destination tx hash that emitted the `Filled` event.  |
| `pollingInterval` | `number`                               | No       | Polling interval in milliseconds, defaults to the `client` polling interval.  |
| `onError`         | `(error: Error) => void`               | No       | Optional callback that'll be invoked when an error occurs. |

## Return

### `unwatch`

Terminates watching of contract events.

```ts
() => void
```

## Example

```ts
import { watchDidFill } from '@omni-network/core'

const { unwatch, status, destTxHash } = await watchDidFill({
  client: viemWalletClient,
  outboxAddress: contracts.outbox,
  orderId: resolvedOrder.orderId,
  onFill: (txHash) => {
    console.log('Fill tx hash found:', txHash)
  },
})

// Stop watching for the Filled event
unwatch()
```
