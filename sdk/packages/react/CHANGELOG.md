# @omni-network/react

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
