import { readFileSync } from 'node:fs'
import type { Chain } from 'viem'

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

export const INVALID_CHAIN_ID = 1234
export const OMNI_DEVNET_ID = 1651
export const MOCK_L1_ID = 1652
export const MOCK_L2_ID = 1654

// Addresses from lib/contracts/testdata/TestContractAddressReference.golden
export const SOLVERNET_INBOX_ADDRESS =
  '0x7c7759b801078ecb2c41c9caecc2db13c3079c76' as const
export const TOKEN_ADDRESS =
  '0x73cc960fb6705e9a6a3d9eaf4de94a828cfa6d2a' as const
export const INVALID_TOKEN_ADDRESS =
  '0x1234000000000000000000000000000000000000' as const
export const ZERO_ADDRESS =
  '0x0000000000000000000000000000000000000000' as const

export const MOCK_L1_CHAIN: Chain = {
  id: MOCK_L1_ID,
  name: 'Mock L1',
  nativeCurrency: { decimals: 18, name: 'Ether', symbol: 'ETH' },
  rpcUrls: {
    default: {
      http: [RPC_ENDPOINTS.mock_l1],
    },
  },
}

export const MOCK_L2_CHAIN: Chain = {
  id: MOCK_L2_ID,
  name: 'Mock L2',
  nativeCurrency: { decimals: 18, name: 'Ether', symbol: 'ETH' },
  rpcUrls: {
    default: {
      http: [RPC_ENDPOINTS.mock_l2],
    },
  },
}

export const MOCK_CHAINS: Record<number, Chain> = {
  [MOCK_L1_ID]: MOCK_L1_CHAIN,
  [MOCK_L2_ID]: MOCK_L2_CHAIN,
}

export const OMNI_TOKEN_ABI = [
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
