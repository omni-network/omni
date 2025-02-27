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

2. Now you can start using the hooks! Let's build an example of an eth bridge:

```tsx
import { useAccount } from 'wagmi'
import { useOrder } from '@omni-network/react'

const order = {
    srcChainId: 84532, 
    destChainId: 11155420,
    deposit: {
        amount: 10000000000000000,
        isNative: true
    }
    expense: {
        isNative: true,
    }
}

export function useBridge({ amount }: { amount: bigint }) {
  const { address: user } = useAccount()

  const quote = useQuote({
    ...order,
    mode: 'expense',
    deposit: {
        ...order.deposit,
        amount,
    }
    enabled: !!user && amount !== 0n
  })

  const {
    txHash,
    validation,
    txMutation,
    open,
    status,
    waitForTx,
    isError,
    isOpen,
    isTxPending,
    isValidated,
  } = useOrder({
    ...order,
    calls: [{
          target: user ?? '0x',
          value: quote.isSuccess ? quote.expense.amount : 0n,
      }],
      deposit: {
        amount: quote.isSuccess ? quote.deposit.amount : 0n,
      },
      expense: {
        amount: quote.isSuccess ? quote.expense.amount : 0n,
      },
      validateEnabled:
        !!user &&
        !!quote.query.data?.expense.amount &&
        !!quote.query.data?.deposit.amount,
  })

  return {
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
  }
}
```

Lets walk through what's happening here:

- `useAccount` gives us the address of the connected wallet
- `useQuote` internally calls our Solver API requesting a quote for the bridge. Note the `mode` param sets the direction of the quote. 
- The `expense` mode requires a `deposit` amount to be set. A quote will be returned describing the `deposit` and `expense` amounts.
- Deposit reflects the amount to be spent on the source chain.
- Expenses reflects the amount to be used for the calls on the destination chain.
- In this simple bridge, expense is simply the amount sent to the user (deposit - fees).
- `useOrder` is used to create an order. An open method is returned to asynchronously open an order.
- Validation reflects another call to our Solver API to validate the order, this can be helpful for verifying an order before opening it.
- `error` will give you information about any potential issues.
- `status` will give you the status of the order:

```tsx
export type OrderStatus =
  | 'idle'
  | 'opening'
  | 'open'
  | 'closed'
  | 'rejected'
  | 'error'
  | 'filled'
```

And that's it! That's all you need to use SolverNet to bridge eth across L2s.

## Get in touch

You can reach out to us on telegram with any queries, feedback, or requests: [@omnidevsupport](https://t.me/omnidevsupport)
