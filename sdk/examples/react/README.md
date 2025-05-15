# Omni SDK React Example

This example app demonstrates usage of the Omni SDK React hooks, allowing you to use Omni's SolverNet in your React app to support cross-chain transactions.

The example runs on testnet, using the baseSepolia and holesky chains, but feel free to update the config as you see fit!

## Running the example app

```bash
pnpm install
pnpm dev
```

## NOTE:

Occasionally, watching the destination 'Fill' event will fail if using a public RPC provider. This is due to the event not being parsed correctly, but the order will still be filled and we'll set the status as such.
