import { QueryClient, QueryClientProvider } from '@tanstack/react-query'

import { OmniProvider } from '../src/index.js'

export const MOCK_L1_ID = 1652
export const MOCK_L2_ID = 1654
export const ZERO_ADDRESS = '0x0000000000000000000000000000000000000000'

export function createQueryClient() {
  return new QueryClient({
    defaultOptions: {
      queries: {
        retry: false,
      },
    },
  })
}

export function ContextProvider({ children }: { children: React.ReactNode }) {
  return (
    <QueryClientProvider client={createQueryClient()}>
      <OmniProvider env="devnet">{children}</OmniProvider>
    </QueryClientProvider>
  )
}
