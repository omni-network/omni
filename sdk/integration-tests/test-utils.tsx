import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { type RenderHookResult, renderHook } from '@testing-library/react'
import { http, type Chain, createWalletClient, publicActions } from 'viem'
import { privateKeyToAccount } from 'viem/accounts'
import {
  type Config,
  type CreateConnectorFn,
  WagmiProvider,
  createConfig,
  mock,
} from 'wagmi'

import { OmniProvider } from '../src/index.js'

export const ETHER = 1_000_000_000_000_000_000n // 18 decimals
export const MOCK_L1_ID = 1652
export const MOCK_L2_ID = 1654
export const ZERO_ADDRESS =
  '0x0000000000000000000000000000000000000000' as const

const MOCK_L1_CHAIN: Chain = {
  id: MOCK_L1_ID,
  name: 'Mock L1',
  nativeCurrency: { decimals: 18, name: 'Ether', symbol: 'ETH' },
  rpcUrls: {
    default: {
      http: ['http://localhost:8001'],
    },
  },
}

const MOCK_L2_CHAIN: Chain = {
  id: MOCK_L2_ID,
  name: 'Mock L2',
  nativeCurrency: { decimals: 18, name: 'Ether', symbol: 'ETH' },
  rpcUrls: {
    default: {
      http: ['http://localhost:8002'],
    },
  },
}

const MOCK_CHAINS: Record<number, Chain> = {
  [MOCK_L1_ID]: MOCK_L1_CHAIN,
  [MOCK_L2_ID]: MOCK_L2_CHAIN,
}

export const accounts = ['0xE0cF003AC27FaeC91f107E3834968A601842e9c6'] as const

const mockConnector = mock({ accounts })

const account = privateKeyToAccount(
  '0xbb119deceaff95378015e684292e91a37ef2ae1522f300a2cfdcb5b004bbf00d',
)

function createClient({ chain }: { chain: Chain }) {
  return createWalletClient({ account, chain, transport: http() })
}

export const testConnector: CreateConnectorFn = (config) => {
  const connector = mockConnector(config)
  connector.getClient = async ({ chainId } = {}) => {
    const chain = chainId ? MOCK_CHAINS[chainId] : MOCK_L1_CHAIN
    if (chain == null) {
      throw new Error(`Unsupported chain: ${chainId}`)
    }
    return createClient({ chain })
  }
  return connector
}

export function createWagmiConfig() {
  return createConfig({
    chains: [MOCK_L1_CHAIN, MOCK_L2_CHAIN],
    client: createClient,
  })
}

export function createQueryClient() {
  return new QueryClient({
    defaultOptions: {
      queries: {
        retry: false,
      },
    },
  })
}

export type TestConfig = {
  queryClient?: QueryClient
  wagmiConfig?: Config
}

export type ContextProviderProps = TestConfig & {
  children: React.ReactNode
}

export function ContextProvider(props: ContextProviderProps) {
  const client = props.queryClient ?? createQueryClient()
  const config = props.wagmiConfig ?? createWagmiConfig()

  return (
    <WagmiProvider config={config}>
      <QueryClientProvider client={client}>
        <OmniProvider env="devnet">{props.children}</OmniProvider>
      </QueryClientProvider>
    </WagmiProvider>
  )
}

export function createRenderHook(config: TestConfig) {
  return function customRenderHook<Result>(
    render: () => Result,
  ): RenderHookResult<Result, ContextProviderProps> {
    return renderHook(render, {
      initialProps: config,
      wrapper: ContextProvider,
    })
  }
}
