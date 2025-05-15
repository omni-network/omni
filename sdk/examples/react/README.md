# Omni SDK React Example

This example app demonstrates usage of the Omni SDK React hooks, allowing you to use Omni's SolverNet in your React app to support cross-chain transactions.

The example runs on testnet, using the baseSepolia and holesky chains, but feel free to update the config as you see fit!

## Running the example app

```bash
pnpm install
pnpm dev
```

## NOTE:

Fetching the destination chain transaction hash may fail if you're using a public RPC provider. Internally, we use viems watchContractEvent to listen for the `Fill` event, but this isn't reliable when working with public RPCs. It's strongly recommended to use a private RPC provider anyway, to prevent instability or rate limiting issues.
