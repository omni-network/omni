# @omni-network/react

## 0.3.3

### Patch Changes

- 1332864: Bug fix in useOrder
- 4c5fe3b: add support for SVM assets in API responses
- Updated dependencies [4c5fe3b]
  - @omni-network/core@0.3.2

## 0.3.2

### Patch Changes

- fecf07b: Introduce getRejection core action and corresponding react hook. This propagates the tx hash of the reject event log and the reject reason.
- Updated dependencies [fecf07b]
  - @omni-network/core@0.3.1

## 0.3.1

### Patch Changes

- 3f8d985: Pass resolvedOrder to useGetOrderStatus so we can fetch destTxHash, and remove obselete use of useGetOrder in useGetOrderStatus.

## 0.3.0

### Minor Changes

- d4f17ec: - Replace dependency on RPC endpoint eth_newFilter (due to not being widely supported) by switching watchDidFill to watch blocks.
  - Includes an update to the useWatchDidFill function signature, replacing `orderId` with `resolvedOrder`.

### Patch Changes

- Updated dependencies [d4f17ec]
  - @omni-network/core@0.3.0

## 0.2.2

### Patch Changes

- 60b097a: add support for debug flag in order validation call
- Updated dependencies [60b097a]
  - @omni-network/core@0.2.1

## 0.2.1

### Patch Changes

- Added query options to relevant hooks

## 0.2.0

### Minor Changes

- 1be3019: Introduce getAssets core API and useOmniAssets react hook
- 736f1ae: Removed isNative flag from the quote inputs, deposit.token and expense.token are always optional now. By default we assume native when token is not provided, but we also support tokens in which case you need to provide the address.

  This cleans up the interface for consumers as the isNative flag was proving awkward.

### Patch Changes

- b0a2600: Use the new useWatchDidFill hook in useOrder to propagate the destTxHash, but fallback to didFill in case consumers rely on public RPCs
- Updated dependencies [b0a2600]
- Updated dependencies [90f880a]
- Updated dependencies [1be3019]
- Updated dependencies [736f1ae]
- Updated dependencies [b48a06c]
  - @omni-network/core@0.2.0

## 0.1.3

### Patch Changes

- f623ee5: Rollback removing middleman from the `withExecAndTransfer` utility.
- Updated dependencies [f623ee5]
  - @omni-network/core@0.1.3

## 0.1.2

- Add `watchDidFill` hook - used to watch for order fill on destination chain, and retrieve the destination tx hash.
- Remove middleman from the `withExecAndTransfer` utility.

## 0.1.1

- Add `watchDidFill` hook - used to watch for order fill on destination chain, and retrieve the destination tx hash.

## 0.1.0

### Breaking changes

- The `withExecAndTransfer` functions is no longer exported by the `@omni-network/react` package, and must instead be imported from `@omni-network/core`.
- Similarly, the following types must be imported from the core pages: `Order`, `OrderStatus`, `Quote`, `Quotable`.

### Other changes

- This version uses the new `@omni-network/core` package internally.

## 0.0.0-alpha.5

- 5a679b0: 🎉 Initial release 🎉

  Initial alpha release of the Omni Solvernet SDK. To get started, follow the guide in the readme.

  Package available at `@omni-network/react`.

  Exports:

  - `useQuote` hook for quoting an order, the result of which should form some of the input for `useOrder`
  - `useOrder` hook for opening and verifying orders
  - `useValidateOrder` hook for validating orders
  - `withExecAndTransfer` utility for executing an order where a contract doesn't have an `onBehalfOf` style parameter
  - `OmniProvider` context that needs to wrap root of apps using the SDK
  - `Order`, `Quote`, `OrderStatus`, `Quoteable` types
