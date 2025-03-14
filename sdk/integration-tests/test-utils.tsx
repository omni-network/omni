import { readFileSync } from 'node:fs'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { type RenderHookResult, renderHook } from '@testing-library/react'
import { http, type Address, type Chain, createWalletClient } from 'viem'
import { type PrivateKeyAccount, privateKeyToAccount } from 'viem/accounts'
import {
  type Config,
  type CreateConnectorFn,
  WagmiProvider,
  createConfig,
  mock,
} from 'wagmi'

import { OmniProvider } from '../src/index.js'

type RPCEndpoints = {
  mock_l1: string
  mock_l2: string
}

let RPC_ENDPOINTS: RPCEndpoints = {
  mock_l1: 'http://127.0.0.1:8003',
  mock_l2: 'http://127.0.0.1:8004',
}
const endpointsFilePath = process.env.E2E_RPC_ENDPOINTS
if (endpointsFilePath != null && endpointsFilePath.trim() !== '') {
  RPC_ENDPOINTS = JSON.parse(readFileSync(endpointsFilePath, 'utf-8'))
}

export const ETHER = 1_000_000_000_000_000_000n // 18 decimals
export const OMNI_DEVNET_ID = 1651
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
      http: [RPC_ENDPOINTS.mock_l1],
    },
  },
}

const MOCK_L2_CHAIN: Chain = {
  id: MOCK_L2_ID,
  name: 'Mock L2',
  nativeCurrency: { decimals: 18, name: 'Ether', symbol: 'ETH' },
  rpcUrls: {
    default: {
      http: [RPC_ENDPOINTS.mock_l2],
    },
  },
}

const MOCK_CHAINS: Record<number, Chain> = {
  [MOCK_L1_ID]: MOCK_L1_CHAIN,
  [MOCK_L2_ID]: MOCK_L2_CHAIN,
}

export const ACCOUNTS_RECORD = {
  '0xE0cF003AC27FaeC91f107E3834968A601842e9c6': privateKeyToAccount(
    '0xbb119deceaff95378015e684292e91a37ef2ae1522f300a2cfdcb5b004bbf00d',
  ),
  '0x3C298a8fAb961CC155F48557872D37D39015c5bc': privateKeyToAccount(
    '0xed3dccb053880be5b681f6f0256fc18410f99bd69fedfc80f6b37e7930d1c526',
  ),
  '0xD1862026335f1cfD7c5DB13e58bA4C97247A7998': privateKeyToAccount(
    '0xe0cff2136e89d72576e1e8a3af640af8509fa07191429766d89814923b9ddbc2',
  ),
  '0xC38Ef10ecC4aD9c24554cACd67cA4896B4fb2F9C': privateKeyToAccount(
    '0x7b63e2b097b37620054a0c3ba9e2c0253751b26f62f866dcbde64508071c0048',
  ),
  '0x8696e3d6cAD982C32F86320a6f0E1AB8aB3Db3b9': privateKeyToAccount(
    '0x2b6c35e1914655810e6471aea16358c335995985eb5d1a8e2a49ee5dca6779c1',
  ),
  '0xB5775c7e3796822dA059B73218E8a9033dA9bC67': privateKeyToAccount(
    '0x49f5dabfb06f9febf27b3b07dc68f5c3a022edf8ac56631e92ba8879d7d7f44e',
  ),
  '0x9637Bc245647B4cdD85e3bB092c672e2ddD28539': privateKeyToAccount(
    '0x07fe974baf69d3d4d93438698155adffae10e24ab14eaa468e69a17c2a1295d3',
  ),
  '0x23a4523A3EE6220fB2CdDc5Ab94A2780D6493230': privateKeyToAccount(
    '0x8b8638c2c593903c63f096200dad748af3e36ca4dbcfd11f57cfa83e98cfdbc8',
  ),
  '0xACb32F1b31F818511139a67b79010fA011960764': privateKeyToAccount(
    '0xdf6cd4fcdba1068873acef34a91f41c9af581dafd52fbc2908add6fb213f5e02',
  ),
  '0x3c56a0cDB54D07A91791b698d8B390aB53208E92': privateKeyToAccount(
    '0xefebe60ea08dbc53c3e59612380501d4cbe7309548d84692ec978e06e3c61137',
  ),
} as const satisfies Record<Address, PrivateKeyAccount>

export const accounts = Object.keys(ACCOUNTS_RECORD) as [Address, ...Address[]]

const mockConnector = mock({ accounts })

export function createClientFactory(account: PrivateKeyAccount) {
  return function createClient({ chain }: { chain: Chain }) {
    return createWalletClient({ account, chain, transport: http() })
  }
}

export function createTestConnector(createClient): CreateConnectorFn {
  return function testConnector(config) {
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
}

const createClient = createClientFactory(ACCOUNTS_RECORD[accounts[0]])

export const testConnector = createTestConnector(createClient)

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
