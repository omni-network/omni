---
sidebar_position: 1
title: Getting Started
---

# Getting Started with the Omni SDK

:::info Alpha SDK
The SDK is currently in alpha - expect potential breaking changes.
:::

## Overview

The Omni SDK contains React hooks for interfacing with Omni SolverNet, your gateway to enabling any transaction, on any chain, directly from your application frontend.

## Prerequisites

You'll need to have `wagmi` and `react-query` setup in your project already. If you don't, you can follow the instructions here: [Wagmi Getting Started Guide](https://wagmi.sh/react/getting-started).

## Installation

Once your project meets the prerequisites, install the SDK package:

```bash
pnpm install @omni-network/react
# or
yarn add @omni-network/react
# or
npm install @omni-network/react
```

## Setup

You need to wrap your application with the `OmniProvider`. Ensure it is placed *inside* your `WagmiProvider` and `QueryClientProvider`:

```tsx title="App.tsx / main.tsx"
import { WagmiProvider } from 'wagmi'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { OmniProvider } from '@omni-network/react'

const queryClient = new QueryClient()

function AppWrapper() {
  return (
    <WagmiProvider config={wagmiConfig}> {/* Your wagmi config */}
      <QueryClientProvider client={queryClient}>
        <OmniProvider env="testnet"> {/* 'testnet' or 'mainnet' */}
          <App />
        </OmniProvider>
      </QueryClientProvider>
    </WagmiProvider>
  )
}
```

Key points:

*   Provide your `wagmi` configuration to `WagmiProvider`.
*   The `env` prop in `OmniProvider` specifies the target Omni network environment. Use `testnet` for development and testing, and `mainnet` for production.

Now you're ready to use the Omni SDK hooks in your application components!
