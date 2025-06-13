# @omni-network/core

## 0.3.2

### Patch Changes

- 4c5fe3b: add support for SVM assets in API responses

## 0.3.1

### Patch Changes

- fecf07b: Introduce getRejection core action and corresponding react hook. This propagates the tx hash of the reject event log and the reject reason.

## 0.3.0

### Minor Changes

- d4f17ec: - Replace dependency on RPC endpoint eth_newFilter (due to not being widely supported) by switching watchDidFill to watch blocks.
  - Includes an update to the watchDidFill function signature, replacing `onLogs` with `onFill`.
  - Includes an update to the `generateOrder` function signature, requiring `outboxAddress` to be passed in.
  - Use `watchDidFill` in `generateOrder`, to propagate the destTxHash.

## 0.2.1

### Patch Changes

- 60b097a: add support for debug flag in order validation call

## 0.2.0

### Minor Changes

- 1be3019: Introduce getAssets core API and useOmniAssets react hook
- 736f1ae: Removed isNative flag from the quote inputs, deposit.token and expense.token are always optional now. By default we assume native when token is not provided, but we also support tokens in which case you need to provide the address.

  This cleans up the interface for consumers as the isNative flag was proving awkward.

- b48a06c: Replace middleman logic with executor

### Patch Changes

- b0a2600: Use the new useWatchDidFill hook in useOrder to propagate the destTxHash, but fallback to didFill in case consumers rely on public RPCs
- 90f880a: Add gas options to `sendOrder` and related functions parameters

## 0.1.3

### Patch Changes

- f623ee5: Rollback removing middleman

## 0.1.2

- Add `watchDidFill` API - used to watch for order fill on destination chain, and retrieve the destination tx hash.
- Remove middleman from the `withExecAndTransfer` utility.

## 0.1.1

- Add `watchDidFill` API - used to watch for order fill on destination chain, and retrieve the destination tx hash.

## 0.1.0

Initial version.
