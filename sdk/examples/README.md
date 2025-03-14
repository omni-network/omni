# Omni SDK Integration Guide

This example demonstrates how to integrate the Omni SDK into a React application to enable cross-chain ETH bridging with minimal code.

## Quick Start

1. Install the Omni SDK package:
```bash
npm install @omni/sdk
# or
yarn add @omni/sdk
# or
pnpm add @omni/sdk
```

2. Add Omni SDK provider to your application:
```tsx
import { OmniProvider } from '@omni/sdk';

function App() {
  return (
    <OmniProvider env="testnet">
      <YourApp />
    </OmniProvider>
  );
}
```

3. Use the Omni SDK hooks in your components:
```tsx
import { useQuote, useOrder, useGetOrderStatus } from '@omni/sdk';
```

## Key Integration Points

Our example application (`bridge-eth.tsx`) highlights six key SDK integration points:

1. **Import SDK Components**: Import the essential hooks and components
2. **Configure Cross-Chain Quote**: Set up quote parameters for the bridge transaction
3. **Create Order Parameters**: Define the transaction details with useOrder
4. **Track Transaction Status**: Monitor bridge progress with useGetOrderStatus
5. **Execute the Bridge**: Initiate the cross-chain transaction with orderResult.open()
6. **Root Provider Setup**: Wrap your application with OmniProvider

## Core SDK Hooks

### `useQuote`

Get a quote for a cross-chain transfer:

```tsx
// SDK INTEGRATION POINT #2 - Get cross-chain quote
const quoteResult = useQuote({
  srcChainId: sourceChainId,
  destChainId: destinationChainId,
  mode: quoteMode === 'deposit' ? 'expense' : 'deposit',
  deposit: {
    isNative: true,
    amount: parseEther(amount || '0')
  },
  expense: {
    isNative: true,
    amount: parseEther(amount || '0')
  },
  enabled: !!address && parseFloat(amount || '0') > 0
});
```

### `useOrder`

Create and execute a bridge transaction:

```tsx
// SDK INTEGRATION POINT #3 - Setup order parameters
const orderResult = useOrder({
  srcChainId: sourceChainId,
  destChainId: destinationChainId,
  owner: address,
  deposit: {
    token: undefined, // For native ETH
    amount: quoteResult.isSuccess ? quoteResult.deposit.amount : 0n
  },
  calls: [
    {
      target: (address || '0x') as `0x${string}`, // Send to same address by default
      value: quoteResult.isSuccess
        ? (quoteMode === 'deposit'
          ? quoteResult.expense.amount  // "Send Exact" mode
          : parseEther(amount || '0'))  // "Receive Exact" mode
        : parseEther(amount || '0'),
    }
  ],
  expense: {
    token: undefined,
    amount: 0n
  },
  validateEnabled: quoteResult.isSuccess
});

// Execute the transaction
const txHash = await orderResult.open();
```

### `useGetOrderStatus`

Track the status of a bridge transaction:

```tsx
// SDK INTEGRATION POINT #4 - Track transaction status
const orderStatus = useGetOrderStatus({
  destChainId: destinationChainId,
  orderId,
  srcChainId: sourceChainId
});

// Check status
if (orderStatus.status === 'filled') {
  // Transaction completed!
}
```

## Best Practices

1. **Clear Error Handling**: Reset errors when user makes changes to inputs to prevent confusion
2. **Chain Validation**: Auto-switch or prompt users to switch to the correct chain
3. **Balance Checks**: Verify sufficient balance before allowing transaction submission
4. **User Feedback**: Provide clear status updates throughout the bridging process
5. **Input Validation**: Validate all user inputs before attempting bridge operations

## Documentation

For full API documentation, visit [docs.omni.network](https://docs.omni.network).
