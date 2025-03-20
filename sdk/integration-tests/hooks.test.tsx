import { renderHook, waitFor } from '@testing-library/react'
import { describe, expect, test } from 'vitest'

import { useQuote } from '../src/index.js'

import {
  type AnyOrder,
  ContextProvider,
  ETHER,
  INVALID_CHAIN_ID,
  INVALID_TOKEN_ADDRESS,
  MOCK_L1_ID,
  MOCK_L2_ID,
  TOKEN_ADDRESS,
  ZERO_ADDRESS,
  executeTestOrder,
  testAccount,
} from './test-utils.js'

describe('useQuote()', () => {
  test('parameters: gets a quote in expense mode', async () => {
    const { result } = renderHook(
      () => {
        return useQuote({
          enabled: true,
          mode: 'expense',
          srcChainId: MOCK_L1_ID,
          destChainId: MOCK_L2_ID,
          deposit: {
            amount: 1n,
            isNative: true,
          },
          expense: {
            amount: 1n,
            isNative: true,
          },
        })
      },
      { wrapper: ContextProvider },
    )

    await waitFor(() => expect(result.current.isSuccess).toBe(true))

    expect(result.current.query.data).toEqual({
      deposit: { token: ZERO_ADDRESS, amount: 1n },
      expense: { token: ZERO_ADDRESS, amount: 0n },
    })
  })

  test('parameters: gets a quote in deposit mode', async () => {
    const { result } = renderHook(
      () => {
        return useQuote({
          enabled: true,
          mode: 'deposit',
          srcChainId: MOCK_L1_ID,
          destChainId: MOCK_L2_ID,
          deposit: {
            amount: 1n,
            isNative: true,
          },
          expense: {
            amount: 1n,
            isNative: true,
          },
        })
      },
      { wrapper: ContextProvider },
    )

    await waitFor(() => expect(result.current.isSuccess).toBe(true))

    expect(result.current.query.data).toEqual({
      deposit: { token: ZERO_ADDRESS, amount: 2n },
      expense: { token: ZERO_ADDRESS, amount: 1n },
    })
  })
})

describe('useOrder()', () => {
  test('default: succeeds with valid order', async () => {
    const amount = ETHER / 2n
    const order: AnyOrder = {
      srcChainId: MOCK_L1_ID,
      destChainId: MOCK_L2_ID,
      expense: { token: ZERO_ADDRESS, amount },
      deposit: { token: ZERO_ADDRESS, amount: ETHER },
      calls: [{ target: testAccount.address, value: amount }],
    }
    await executeTestOrder(order)
  })

  test('behaviour: rejects when using invalid source chain', async () => {
    const order: AnyOrder = {
      srcChainId: INVALID_CHAIN_ID,
      destChainId: MOCK_L1_ID,
      expense: { token: ZERO_ADDRESS, amount: 1n },
      deposit: { token: ZERO_ADDRESS, amount: 1n },
      calls: [{ target: testAccount.address, value: 1n }],
    }
    await executeTestOrder(order, 'UnsupportedSrcChain')
  })

  test('behaviour: rejects when using invalid destination chain', async () => {
    const order: AnyOrder = {
      srcChainId: MOCK_L1_ID,
      destChainId: INVALID_CHAIN_ID,
      expense: { token: ZERO_ADDRESS, amount: 1n },
      deposit: { token: ZERO_ADDRESS, amount: 1n },
      calls: [{ target: testAccount.address, value: 1n }],
    }
    await executeTestOrder(order, 'UnsupportedDestChain')
  })

  test('behaviour: rejects when source and destination chains are the same', async () => {
    const order: AnyOrder = {
      srcChainId: MOCK_L1_ID,
      destChainId: MOCK_L1_ID,
      expense: { token: ZERO_ADDRESS, amount: 1n },
      deposit: { token: ZERO_ADDRESS, amount: 1n },
      calls: [{ target: testAccount.address, value: 1n }],
    }
    await executeTestOrder(order, 'SameChain')
  })

  test('behaviour: rejects when using an unsupported expense token', async () => {
    const order: AnyOrder = {
      srcChainId: MOCK_L1_ID,
      destChainId: MOCK_L2_ID,
      expense: { token: INVALID_TOKEN_ADDRESS, amount: 1n },
      deposit: { token: TOKEN_ADDRESS, amount: 1n },
      calls: [{ target: testAccount.address, value: 1n }],
    }
    await executeTestOrder(order, 'UnsupportedExpense')
  })
})
