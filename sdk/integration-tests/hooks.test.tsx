import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { renderHook, waitFor } from '@testing-library/react'
import { baseSepolia, holesky } from 'viem/chains'
import { describe, expect, test } from 'vitest'

import { OmniProvider, useQuote } from '../src/index.js'

const collateralAbi = [
  {
    type: 'function',
    inputs: [
      { name: 'owner', internalType: 'address', type: 'address' },
      { name: 'spender', internalType: 'address', type: 'address' },
    ],
    name: 'allowance',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [
      { name: 'spender', internalType: 'address', type: 'address' },
      { name: 'value', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'approve',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [],
    name: 'asset',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: 'account', internalType: 'address', type: 'address' }],
    name: 'balanceOf',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [
      { name: 'issuer', internalType: 'address', type: 'address' },
      { name: 'recipient', internalType: 'address', type: 'address' },
    ],
    name: 'debt',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [
      { name: 'recipient', internalType: 'address', type: 'address' },
      { name: 'amount', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'deposit',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [
      { name: 'recipient', internalType: 'address', type: 'address' },
      { name: 'amount', internalType: 'uint256', type: 'uint256' },
      { name: 'deadline', internalType: 'uint256', type: 'uint256' },
      { name: 'v', internalType: 'uint8', type: 'uint8' },
      { name: 'r', internalType: 'bytes32', type: 'bytes32' },
      { name: 's', internalType: 'bytes32', type: 'bytes32' },
    ],
    name: 'deposit',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [{ name: 'amount', internalType: 'uint256', type: 'uint256' }],
    name: 'increaseLimit',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [
      { name: 'recipient', internalType: 'address', type: 'address' },
      { name: 'amount', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'issueDebt',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [{ name: 'issuer', internalType: 'address', type: 'address' }],
    name: 'issuerDebt',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: 'issuer', internalType: 'address', type: 'address' }],
    name: 'issuerRepaidDebt',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'limit',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'limitIncreaser',
    outputs: [{ name: '', internalType: 'address', type: 'address' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: 'recipient', internalType: 'address', type: 'address' }],
    name: 'recipientDebt',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [{ name: 'recipient', internalType: 'address', type: 'address' }],
    name: 'recipientRepaidDebt',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [
      { name: 'issuer', internalType: 'address', type: 'address' },
      { name: 'recipient', internalType: 'address', type: 'address' },
    ],
    name: 'repaidDebt',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [
      { name: 'limitIncreaser', internalType: 'address', type: 'address' },
    ],
    name: 'setLimitIncreaser',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [],
    name: 'totalDebt',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'totalRepaidDebt',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [],
    name: 'totalSupply',
    outputs: [{ name: '', internalType: 'uint256', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    inputs: [
      { name: 'to', internalType: 'address', type: 'address' },
      { name: 'value', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'transfer',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [
      { name: 'from', internalType: 'address', type: 'address' },
      { name: 'to', internalType: 'address', type: 'address' },
      { name: 'value', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'transferFrom',
    outputs: [{ name: '', internalType: 'bool', type: 'bool' }],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    inputs: [
      { name: 'recipient', internalType: 'address', type: 'address' },
      { name: 'amount', internalType: 'uint256', type: 'uint256' },
    ],
    name: 'withdraw',
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'owner',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'spender',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'value',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'Approval',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'depositor',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'recipient',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'Deposit',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'IncreaseLimit',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'issuer',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'recipient',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'debtIssued',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'IssueDebt',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'issuer',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'recipient',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'debtRepaid',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'RepayDebt',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'limitIncreaser',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
    ],
    name: 'SetLimitIncreaser',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      { name: 'from', internalType: 'address', type: 'address', indexed: true },
      { name: 'to', internalType: 'address', type: 'address', indexed: true },
      {
        name: 'value',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'Transfer',
  },
  {
    type: 'event',
    anonymous: false,
    inputs: [
      {
        name: 'withdrawer',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'recipient',
        internalType: 'address',
        type: 'address',
        indexed: true,
      },
      {
        name: 'amount',
        internalType: 'uint256',
        type: 'uint256',
        indexed: false,
      },
    ],
    name: 'Withdraw',
  },
  { type: 'error', inputs: [], name: 'ExceedsLimit' },
  { type: 'error', inputs: [], name: 'InsufficientDeposit' },
  { type: 'error', inputs: [], name: 'InsufficientIssueDebt' },
  { type: 'error', inputs: [], name: 'InsufficientWithdraw' },
  { type: 'error', inputs: [], name: 'NotLimitIncreaser' },
] as const

const testCall = {
  operation: 'Deposit',
  destChain: holesky,
  destChainId: holesky.id,
  supportedL2s: [baseSepolia.id],
  target: '0x23e98253f372ee29910e22986fe75bb287b011fc',
  abi: collateralAbi,
  deposit: {
    name: 'wstETH',
    symbol: 'wstETH',
    contract: '0x6319df7c227e34B967C1903A08a698A3cC43492B',
    imgPath: '/tokens/steth.png',
  },
  receive: {
    name: 'DC_wstETH',
    symbol: 'DC_wstETH',
    contract: '0x23e98253f372ee29910e22986fe75bb287b011fc',
    imgPath: '/protocols/symbiotic.png',
  },
  expense: {
    name: 'wstETH',
    symbol: 'wstETH',
    contract: '0x8d09a4502Cc8Cf1547aD300E066060D043f6982D',
    imgPath: '/tokens/steth.png',
  },
  maxAmount: 0.1,
  minAmount: 0.002,
} as const

function ContextProvider({ children }: { children: React.ReactNode }) {
  return (
    <QueryClientProvider client={new QueryClient()}>
      <OmniProvider env="devnet">{children}</OmniProvider>
    </QueryClientProvider>
  )
}

describe('useQuote()', () => {
  test('should get a quote', async () => {
    const user = '0x0123456789abcdef0123456789abcdef01234567'
    const order = {
      owner: user,
      srcChainId: 1,
      destChainId: testCall.destChainId,
      calls: [
        {
          target: testCall.target,
          value: BigInt(0),
          abi: testCall.abi,
          functionName: 'deposit',
          args: [user, 0n],
        },
      ],
      deposit: {
        token: testCall.deposit.contract,
        amount: 1n,
      },
      expense: {
        token: testCall.expense.contract,
        spender: testCall.target,
        amount: 0n,
      },
    }

    const { result } = renderHook(
      () => {
        return useQuote({
          enabled: true,
          mode: 'expense',
          ...order,
          deposit: {
            ...order.deposit,
            isNative: false,
          },
          expense: {
            ...order.expense,
            isNative: false,
          },
        })
      },
      { wrapper: ContextProvider },
    )

    console.log('result', result.current)

    await waitFor(() => expect(result.current.isSuccess).toBe(true))
  })
})
