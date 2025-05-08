# Omni SDK Core Example

This example script demonstrates usage of the Omni SDK Core functions, allowing you to use Omni's SolverNet in your servers or scripts to support cross-chain transactions.

The example runs on testnet, using the Base Sepolia and Holesky chains, but feel free to update the config as you see fit!

## Running the example

### Prerequisites

```bash
pnpm install
pnpm run build
```

### Running the script

The script transfers 0.01 ETH from Base Sepolia to Holesky using the account defined using the `WALLET_PRIVATE_KEY` environment variable. This environment variable must contain an hexadecimal-encoded private key starting with `0x`.

```bash
WALLET_PRIVATE_KEY=0x123... pnpm start
```
