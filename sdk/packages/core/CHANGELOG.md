# @omni-network/core

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
