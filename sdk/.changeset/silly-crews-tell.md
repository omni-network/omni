---
"@omni-network/react": minor
"@omni-network/core": minor
---

- Replace dependency on RPC endpoint eth_newFilter (due to not being widely supported) by switching watchDidFill to watch blocks.
- Includes an update to the watchDidFill function signature, replacing `onLogs` with `onFill`.
- Includes an update to the useWatchDidFill function signature, replacing `orderId` with `resolvedOrder`.
- Includes an update to the `generateOrder` function signature, requiring `outboxAddress` to be passed in.
- Use `watchDidFill` in `generateOrder`, to propagate the destTxHash.
