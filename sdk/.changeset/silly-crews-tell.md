---
"@omni-network/react": minor
"@omni-network/core": minor
---

Replace dependency on RPC endpoint eth_newFilter (due to not being widely supported) by switching watchDidFill to watch blocks. Including an update to the watchDidFill function signature, replacing `onLogs` with `onFill`.
