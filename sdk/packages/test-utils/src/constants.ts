import { readFileSync } from 'node:fs'
import type { Chain } from 'viem'

type RPCEndpoints = {
  omni_evm: string
  mock_l1: string
  mock_l2: string
}

let rpcEndpoints: RPCEndpoints = {
  omni_evm: 'http://127.0.0.1:8001',
  mock_l1: 'http://127.0.0.1:8003',
  mock_l2: 'http://127.0.0.1:8004',
}
const endpointsFilePath = process.env.E2E_RPC_ENDPOINTS
if (endpointsFilePath != null && endpointsFilePath.trim() !== '') {
  rpcEndpoints = JSON.parse(readFileSync(endpointsFilePath, 'utf-8'))
}

export const invalidChainId = 1234
export const omniDevnetId = 1651
export const mockL1Id = 1652
export const mockL2Id = 1654

// Addresses from lib/contracts/testdata/TestContractAddressReference.golden
export const inbox = '0x7c7759b801078ecb2c41c9caecc2db13c3079c76' as const
export const tokenAddress =
  '0x73cc960fb6705e9a6a3d9eaf4de94a828cfa6d2a' as const
export const nomAddress = '0x31eb5432c37540a12a3ed647777283c9f185770d' as const
export const invalidTokenAddress =
  '0x1234000000000000000000000000000000000000' as const
export const outbox = '0x29d9e8faa760841aacbe79a8632c1f42e0a858e6' as const
export const executor = '0xe73fbea025982dcb008c689669a60468d23b1eab' as const
export const vault = '0x81487c7b22a0babadC98D5cA1d7D21240beB14Cc' as const

export const mockL1Chain: Chain = {
  id: mockL1Id,
  name: 'Mock L1',
  nativeCurrency: { decimals: 18, name: 'Ether', symbol: 'ETH' },
  rpcUrls: {
    default: {
      http: [rpcEndpoints.mock_l1],
    },
  },
}

export const mockL2Chain: Chain = {
  id: mockL2Id,
  name: 'Mock L2',
  nativeCurrency: { decimals: 18, name: 'Ether', symbol: 'ETH' },
  rpcUrls: {
    default: {
      http: [rpcEndpoints.mock_l2],
    },
  },
}

export const omniDevnetChain: Chain = {
  id: omniDevnetId,
  name: 'Omni Devnet',
  nativeCurrency: { decimals: 18, name: 'Ether', symbol: 'ETH' },
  rpcUrls: {
    default: {
      http: [rpcEndpoints.omni_evm],
    },
  },
}

export const mockChains: Record<number, Chain> = {
  [mockL1Id]: mockL1Chain,
  [mockL2Id]: mockL2Chain,
  [omniDevnetId]: omniDevnetChain,
}

export const omniTokenAbi = [
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

export const nomTokenAbi = [
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
    name: 'convert',
    inputs: [
      { name: 'to', type: 'address', internalType: 'address' },
      { name: 'amount', type: 'uint256', internalType: 'uint256' },
    ],
    outputs: [],
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
