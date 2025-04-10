# React Hooks

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

You'll need to wrap your app in the `OmniProvider`. Make sure to wrap it **_inside_** your `WagmiProvider` and `QueryClientProvider` provider:

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

Now you can start using the hooks! Let's build an example of depositing a token into a vault on a different chain. We'll work with `wstETH`

```tsx
const holeskyWSTETH = '0x8d09a4502Cc8Cf1547aD300E066060D043f6982D' as const
const baseSepoliaWSTETH = '0x6319df7c227e34B967C1903A08a698A3cC43492B' as const
```

First, we need to quote how much `wstETH` we can spend on the destination chain for a given source chain deposit.

```tsx
import { useQuote } from '@omni-network/react'

function App() {
    // quote how much Holesky wstETH we can spend for a 0.1 wstETH deposit on Base Sepolia
    const quote = useQuote({
        srcChainId: baseSepolia.id,
        destChainId: holesky.id,
        deposit: { isNative: false, token: baseSepoliaWSTETH, amount: parseEther("0.1") },
        expense: { isNative: false, token: holeskyWSTETH },
        mode: "expense", // quote expense amount
        enabled: true,
    })

    // ...
}
```

Alternatively, we can quote how much `wstETH` we need to deposit for a certain spend.

```tsx
import { useQuote } from '@omni-network/react'

function App() {
    // quote how much BaseSepolia wstETH we need to deposit to spend 0.1 wstETH on Holesky
    const quote = useQuote({
        srcChainId: baseSepolia.id,
        destChainId: holesky.id,
        deposit: { isNative: false, token: baseSepoliaWSTETH },
        expense: { isNative: false, token: holeskyWSTETH, amount: parseEther("0.1") },
        mode: "deposit ", // quote deposit amount
        enabled: true,
    })

    // ...
}
```

Now, we use that quote to inform the order we will open with Omni. We'll use the quoted deposit / expense amounts. We'll also specify the calls we want to make on the destination chain. To describe that call, we'll need an abi.


```tsx
import { useOrder, useQuote } from '@omni-network/react'

// Vault.deposit(address onBehalfOf, uint256 amount)
const vaultABI = [
  {
    inputs: [
      { internalType: 'address', name: 'onBehalfOf', type: 'address' },
      { internalType: 'uint256', name: 'amount', type: 'uint256' },
    ],
    name: 'deposit',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
] as const

// your vault address
const vault = `0x...` as const

function App() {
    // quote how much Holesky wstETH we can spend for a 0.1 wstETH deposit on Base Sepolia
    const quote = useQuote({
        srcChainId: baseSepolia.id,
        destChainId: holesky.id,
        deposit: { isNative: false, token: baseSepoliaWSTETH, amount: parseEther("0.1") },
        expense: { isNative: false, token: holeskyWSTETH },
        mode: "expense", // quote expense amount
        enabled: true,
    })

    // address to deposit on behalf of
    const user = '0x...'

    // quoted amounts
    const depositAmt = quote.isSuccess ? quote.deposit.amount : 0n
    const expenseAmt = quote.isSuccess ? quote.expense.amount : 0n

    const order = useOrder({
        srcChainId: baseSepolia.id,
        destChainId: holesky.id,

        // amount to deposit on source chain, paid by connected account
        deposit: {
            amount: depositAmt,
            token: baseSepoliaWSTETH
        },

        // amount spent on the destination chain, paid by solver
        // spender = the address that will call token.transferFrom(...)
        expense: {
            amount: expenseAmt,
            token: holeskyWSTETH,
            spender: vault
        },

        // request Vault.deposit(user, expenseAmt)
        // Note we calldata also relies on quoted expense amount
        calls: [
          {
            target: vault,
            abi: vaultABI,
            functionName: 'deposit',
            args: [user, expenseAmt]
          }
        ],

        // when true, this will if check the order will be accepted by Omni
        // you can consume the result via order.validation
        validateEnabled: quote.isSuccess
    })
}

```

With the order defined, and quote successfull, the order will be validated with Omni. You can read the result at `order.validation`.

```tsx
order.validation?.status            // 'pending' | 'rejected'  | 'accepted' | 'rejected'
order.validation?.rejectReason      // string reason code (ex. "DestCallReverts")
order.validation?.rejectDescription // a longer description of the specific rejection reason
```


Note validation is best effort, and does not guarantee the order will be filled by the solver.
Once the order is validated, open the order and track it's status:

```tsx
import { useOrder, useQuote } from '@omni-network/react'

function App() {
    // ...

    const {
        open,
        status,
        validation,
        isReady,
    } = order

    const canOpen = isReady && validation?.status === 'accepted'

    return (
        <div>
            <button disabled={!canOpen} onClick={open}>Deposit</button>
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

The `useOrder` hook also returns additional information. The full list of properties is:

```tsx
const {
    // opens the order
    open,

    // order status
    status,

    // the order id, if open (status == 'open')
    orderId,
    isOpen,

    // validation result
    validation,
    isValidated,

    // error state (status == 'error')
    error,
    isError,

    // open tx state (status == 'opening')
    isTxPending,
    isTxSubmitted,
    txMutation,
    txHash,
    waitForTx,

    // ready to open (status == 'ready')
    // does not imply validation is accepted
    isReady,
} = order
```


And that's it! That's all you need to use SolverNet to bridge eth across L2s.

### withExecAndTransfer

Some target contracts do not have `onBehalfOf`-style methods. If they, instead, credit a tokenized position to `msg.sender`, you can use the `withExecAndTransfer` util. This util wraps your call in a call to our [SolverNetMiddleman](https://github.com/omni-network/omni/blob/main/contracts/solve/src/SolverNetMiddleman.sol). The wrapped call executes a call on your target, and transfers all earned tokens to some `to` address.

See ths solidity below code for reference.

```solidity
contract SolverNetMiddleman {
    // ...

    /**
     * @notice Execute a call and transfer any received tokens back to the recipient
     * @dev Intended to be used when interacting with contracts that don't allow us to specify a recipient
     * @param token  Token to transfer
     * @param to     Recipient address
     * @param target Call target address
     * @param data   Calldata for the call
     */
    function executeAndTransfer(address token, address to, address target, bytes calldata data)
        external
        payable
        nonReentrant
    {
        (bool success,) = target.call{ value: msg.value }(data);
        if (!success) revert CallFailed();

        if (token == address(0)) SafeTransferLib.safeTransferAllETH(to);
        else token.safeTransferAll(to);
    }

    // ...
}
```

To use `withExecAndTransfer`, specify your target call, the token to transfer post-call, and address to transfer to.

```typescript
// ABI for Vault.deposit{ value: amount }()
const tokenizedVaultABI = [
  {
    inputs: [],
    name: 'deposit',
    outputs: [],
    stateMutability: 'payable',
    type: 'function',
  },
] as const


// Your vault address
const vault = `0x...` as const


// Fetch SolverNetMiddleman contract address
const contracts = useOmniContracts()
const middlemanAddress = contracts.data?.middleman ?? zeroAddress

// Wrap Vault.deposit() with a middleman call.
const middlemanCall = withExecAndTransfer({
    middlemanAddress: middlemanAddress,
    call: {
        target: vault,
        abi: tokenizedVaultABI,
        functionName: 'deposit',
        value: expenseAmt, // from quote
    },
    transfer: {
        token: vault, // vault is address of token to transfer
        to: user,     // transfer all post-call tokens to this address
    }
})

// Pass the middleman call in to `useOrder`
const order = useOrder({
    srcChainId: baseSepolia.id,
    destChainId: holesky.id,
    deposit: { amount: depositAmt },
    expense: { amount: expenseAmt },
    calls: [middlemanCall],
    validateEnabled: quote.isSuccess,
})
```

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
