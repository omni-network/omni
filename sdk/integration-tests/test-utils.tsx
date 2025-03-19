import { readFileSync } from 'node:fs'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { type RenderHookResult, renderHook } from '@testing-library/react'
import { http, type Chain, createWalletClient, publicActions } from 'viem'
import { privateKeyToAccount } from 'viem/accounts'
import { type Config, WagmiProvider, createConfig, mock } from 'wagmi'

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
// Addresses from lib/contracts/testdata/TestContractAddressReference.golden
export const SOLVERNET_INBOX_ADDRESS = '0x7c7759b801078ecb2c41c9caecc2db13c3079c76' as const
export const TOKEN_ADDRESS =
  '0x73cc960fb6705e9a6a3d9eaf4de94a828cfa6d2a' as const
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

const OMNI_TOKEN_ABI = [
  {
    type: 'function',
    name: 'approve',
    inputs: [
      { name: 'spender', type: 'address', internalType: 'address' },
      { name: 'value', type: 'uint256', internalType: 'uint256' },
    ],
    outputs: [{ name: '', type: 'bool', internalType: 'bool' }],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    name: 'mint',
    inputs: [
      { name: 'to', type: 'address', internalType: 'address' },
      { name: 'amount', type: 'uint256', internalType: 'uint256' },
    ],
    outputs: [],
    stateMutability: 'nonpayable',
  },
] as const

export const testAccount = privateKeyToAccount(
  '0xbb119deceaff95378015e684292e91a37ef2ae1522f300a2cfdcb5b004bbf00d',
)

const mockConnector = mock({ accounts: [testAccount.address] as const })

function createClient({ chain }: { chain: Chain }) {
  return createWalletClient({
    account: testAccount,
    chain,
    transport: http(),
  }).extend(publicActions)
}

const mockL1Client = createClient({ chain: MOCK_L1_CHAIN })

export async function mintOMNI() {
  // mint token
  const mintHash = await mockL1Client.writeContract({
    address: TOKEN_ADDRESS,
    abi: OMNI_TOKEN_ABI,
    functionName: 'mint',
    args: [testAccount.address, 100n * ETHER],
  })
  // wait for transaction to be mined
  await mockL1Client.waitForTransactionReceipt({ hash: mintHash })
  // approve transfers to inbox contract
  const approveHash = await mockL1Client
    .writeContract({
      address: TOKEN_ADDRESS,
      abi: OMNI_TOKEN_ABI,
      functionName: 'approve',
      args: [SOLVERNET_INBOX_ADDRESS, 100n * ETHER],
    })
  // wait for transaction to be mined
  await mockL1Client.waitForTransactionReceipt({ hash: approveHash })
}

export function testConnector(config) {
  const connector = mockConnector(config)
  connector.getClient = async ({ chainId } = {}) => {
    if (chainId === MOCK_L1_ID) {
      return mockL1Client
    }
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

export function createRenderHook(config: TestConfig = {}) {
  return function customRenderHook<Result>(
    render: () => Result,
  ): RenderHookResult<Result, ContextProviderProps> {
    return renderHook(render, {
      initialProps: config,
      wrapper: ContextProvider,
    })
  }
}
