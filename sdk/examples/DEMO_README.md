# Omni SDK Bridge Demo

This is a simple React application demonstrating how to use the Omni SDK to bridge ETH between chains with a clean, user-friendly interface.

## Module System

This project uses ECMAScript Modules (ESM) as specified by `"type": "module"` in package.json.

## Features

The demo includes the following features:
- Connect to your wallet through MetaMask or other web3 wallets
- Select source and destination chains from multiple testnet options
- Choose between "Send Exact" or "Receive Exact" modes
- Real-time bridging quotes with automatic fee calculation
- Submit and track bridge transactions
- Clear error handling with automatic clearing on user input changes
- Network auto-switching to match source chain

## Prerequisites

Before running this demo, make sure you have:

1. Node.js installed (v16 or newer)
2. A web3 wallet like MetaMask installed in your browser
3. Some testnet ETH on at least one of the supported testnets

## Setup

To run the demo application:

1. First, install the dependencies:

```bash
cd sdk/examples
npm install
# or
pnpm install
```

2. Make sure the Omni SDK is properly built. If you're working with the local SDK, you might need to build it first:

```bash
# From the root of the SDK
npm run build
```

3. Start the development server:

```bash
npm run dev
```

4. Open your browser and navigate to `http://localhost:3000`

## Using the Demo

1. **Connect Wallet**: Click the "Connect Wallet" button and approve the connection in your wallet
2. **Select Networks**: Choose source and destination chains from the dropdown menus
3. **Select Mode**: Choose between "Send Exact" (specify how much to send) or "Receive Exact" (specify how much to receive)
4. **Enter Amount**: Input the amount of ETH to bridge
5. **Review Quote**: Check the quote summary showing the amount you'll send, receive, and the fee
6. **Bridge ETH**: Click the "Bridge ETH" button and approve the transaction in your wallet
7. **Track Status**: Monitor the transaction status in the order status section

## Bridge Flow Explained

1. **Quote Calculation**: The SDK calculates the optimal path and fees for your bridge transaction
2. **Network Switching**: If needed, the app will prompt your wallet to switch to the correct source chain
3. **Transaction Creation**: When you click "Bridge ETH", the app creates and signs a transaction on the source chain
4. **Status Tracking**: After submission, the app monitors the transaction across chains until completion

## Supported Chains

This demo supports bridging between:
- Ethereum Holesky (Chain ID: 17000)
- Arbitrum Sepolia (Chain ID: 421614)
- Base Sepolia (Chain ID: 84532)
- Optimism Sepolia (Chain ID: 11155420)

## Understanding the Code

The demo is designed to clearly highlight SDK integration points:

1. The demo imports Omni SDK hooks at the top of the file
2. SDK integration points are marked with clear comments
3. The core functionality is organized into self-contained sections
4. Error handling and user input management are separated from SDK logic

## Troubleshooting

- **Connection Issues**: Make sure your wallet is installed and unlocked
- **Transaction Errors**: Check that you have sufficient balance for both the bridge amount and gas fees
- **Network Switching**: If prompted to switch networks, approve the request in your wallet
- **Quote Unavailable**: Ensure the amount is greater than 0 and the wallet is connected
- **Stuck Transactions**: Bridge operations can take several minutes to complete, particularly on testnets

## Resources

- [Omni Network Documentation](https://docs.omni.network/)
- [Wagmi Documentation](https://wagmi.sh/)
- [Viem Documentation](https://viem.sh/)
- [Vite Documentation](https://vitejs.dev/)
