# `waitForOrderOpen`

The `waitForOrderOpen` waits for an order to be open on a SolverNet inbox contract.

## Usage

`import { waitForOrderOpen } from '@omni-network/core'`

```tsx
import { waitForOrderOpen } from '@omni-network/core'

const resolvedOrder = await waitForOrderOpen({
  // ... params
});
```

## Parameters

| Prop                | Type                                 | Required | Description                                                                                                                         |
| ------------------- | ------------------------------------ | -------- | ----------------------------------------------------------------------------------------------------------------------------------- |
| `client`        | `Client`                             | Yes      | The `viem` client used to read from the blockchain.                                                                         |

| `txHash`           | `Hex`                         | Yes      | Transaction hash returned by calling the [`sendOrder` function](/sdk/core/sendOrder) |
| `pollinginterval`       | `number`                             | No      | Polling interval in milliseconds, defaults to the `client` polling interval.                                                                     |

## Return

`waitForOrderOpen` returns the [Promise](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise) of a `ResolvedOrder` object.

### `ResolvedOrder`

```ts
type ResolvedOrder = {
  user: Address;
  originChainId: bigint;
  openDeadline: number;
  fillDeadline: number;
  maxSpent: Output[];
  minReceived: Output[];
  fillInstructions: FillInstruction[];
}
```

### `Output`

```ts
type Output = {
  token: Address;
  amount: bigint;
  recipient: Address;
  chainId: bigint;
}
```

### `FillInstruction`

```ts
type FillInstruction = {
  destinationChainId: bigint;
  destinationSettler: Address;
  originData: Hex;
}
```

## Example

```ts
import { sendOrder, waitForOrderOpen } from '@omni-network/core'

const txHash = await sendOrder({
  client: viemClient,
  inboxAddress: contracts.inbox,
  order: orderParams
})
const resolvedOrder = await waitForOrderOpen({
  client: viemClient,
  txHash: txHash,
})
```
