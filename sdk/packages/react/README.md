# Omni SDK

## Note - The SDK is in alpha so expect breaking changes

## Overview

The Omni SDK contains React hooks for interfacing with Omni SolverNet, your gateway to any transaction, on any chain.

## Getting Started

The SDK in its current form has peer dependencies on `viem`, `wagmi`, and `react-query`. In future, we plan to build additional setups that don't have the same dependencies.

Note - given we're in alpha, we're still making improvements, particularly on e2e type safety, testing, and documentation. If you have any feedback or requests, please reach out to us (telegram below).

### Installation

1. You'll need to have `wagmi` and `react-query` setup in your project already. If you don't, you can follow the instructions [here](https://wagmi.sh/react/getting-started).

2. Once you're setup, install the SDK:
```bash
pnpm install @omni-network/react
```

### Usage

1. You'll need to wrap your app in the `OmniProvider`. Make sure to wrap it **_inside_** your `WagmiProvider` and `QueryClientProvider` provider:

```tsx
import { WagmiProvider } from 'wagmi'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { OmniProvider } from '@omni-network/react'

const queryClient = new QueryClient()

<WagmiProvider>
  <QueryClientProvider client={queryClient}>
    <OmniProvider env="testnet">
      <App />
    </OmniProvider>
  </QueryClientProvider>
</WagmiProvider>
```

Note - you need to supply an `env` prop, for now default to `testnet`.

2. Now you can start using the hooks! Let's build an example of an eth bridge from Base Sepolia to Arbitrum Sepolia.

First, we need to quote how much ETH we can receive on the destination chain for a given source chain deposit:

```tsx
import { useQuote } from '@omni-network/react'

function App() {
    // quote how much ArbSepolia eth we can get for 0.1 Eth on BaseSepolia
    const quote = useQuote({
        srcChainId: baseSepolia.id,
        destChainId: arbitrumSepolia.id,
        deposit: { isNative: true, amount: parseEther("0.1") },
        expense: { isNative: true, },
        mode: "expense", // quote expense amount
        enabled: true,
    })

    // ...
}
```

Now, we use that quote to inform the order we will open with Omni:

```tsx
import { useOrder, useQuote } from '@omni-network/react'

function App() {
   // ...
  const user = "0x...."
  const order = useOrder({
    srcChainId: baseSepolia.id,
    destChainId: arbitrumSepolia.id,

    // request ETH transfer of quoted expense to `user`
    calls: [
      {
        target: user,
        value: quote.isSuccess ? quote.expense.amount : 0n,
      }
    ],
    deposit: {
      amount: quote.isSuccess ? quote.deposit.amount : 0n,
    },
    expense: {
      amount: quote.isSuccess ? quote.expense.amount : 0n,
    },
    // when true, this will if check the order will be accepted by Omni, you can consume the result via validation
    validateEnabled: quote.isSuccess
  })
}

```

Finally, check the order is ready to be opened, open the order, and checks it's status:

```tsx
import { useOrder, useQuote } from '@omni-network/react'

function App() {
  // ...
  const {
    open,
    txHash,
    validation,
    txMutation,
    status,
    waitForTx,
    isError,
    isOpen,
    isTxPending,
    isValidated,
    isReady,
  } = order

  return (
    <div>
        <button disabled={!isReady} onClick={open}>Bridge</button>
        <p>Order status: {status}</p>
    </div>
  )
}


```

Order status lets you track the order's progress.

```tsx
export type OrderStatus =
  | 'initializing'
  | 'ready'
  | 'opening'
  | 'open'
  | 'closed'
  | 'rejected'
  | 'error'
  | 'filled'
```

And that's it! That's all you need to use SolverNet to bridge eth across L2s.

# Supported Assets

| Network | Chain | Asset | Contract Address | Min | Max |
|---------|-------|-------|-----------------|-----|-----|
| Mainnet | Ethereum, Base, Arbitrum One, Optimism | `ETH` | Native | 0.001 | 1 |
| Mainnet | Ethereum | `wstETH` | `0x7f39c581f595b53c5cb19bd0b3f8da6c935e2ca0` | 0.001 | 4 |
| Mainnet | Ethereum | `stETH` | `0xae7ab96520de3a18e5e111b5eaab095312d7fe84` | 0.001 | 4 |
| Mainnet | Base | `wstETH` | `0xc1cba3fcea344f92d9239c08c0568f6f2f0ee452` | 0.001 | 4 |
| Testnet | Holesky, Arb/Base/Op Sepolia | `ETH` | Native | 0.001 | 1 |
| Testnet | Holesky | `wstETH` | `0x8d09a4502cc8cf1547ad300e066060d043f6982d` | 0.001 | 1 |
| Testnet | Holesky | `stETH` | `0x3f1c547b21f65e10480de3ad8e19faac46c95034` | 0.001 | 1 |
| Testnet | Base Sepolia | `Mock wstETH` (mintable) | `0x6319df7c227e34B967C1903A08a698A3cC43492B` | 0.001 | 1 |

> **Note:** Currently limited to like-asset deposits (e.g., wstETH on Base → wstETH vault on Ethereum). Cross-asset swaps coming soon!

## Get in touch

You can reach out to us on telegram with any queries, feedback, or requests: [@omnidevsupport](https://t.me/omnidevsupport).

Of course feel free to open an issue or discussion here on github also.
