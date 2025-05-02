import { OmniProvider } from '@omni-network/react'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { baseSepolia, holesky } from 'viem/chains'
import { http, WagmiProvider, createConfig } from 'wagmi'
import { metaMask } from 'wagmi/connectors'

export const config = createConfig({
  chains: [baseSepolia, holesky],
  connectors: [metaMask()],
  transports: {
    [baseSepolia.id]: http(),
    [holesky.id]: http(),
  },
})

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      retry: 0,
    },
  },
})

export const Providers = ({ children }: { children: React.ReactNode }) => {
  return (
    <WagmiProvider config={config}>
      <QueryClientProvider client={queryClient}>
        <OmniProvider env="testnet">{children}</OmniProvider>
      </QueryClientProvider>
    </WagmiProvider>
  )
}
