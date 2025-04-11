import { useQuote, useValidateOrder } from '@omni-network/react'
import {
  invalidChainId,
  invalidTokenAddress,
  mockL1Id,
  mockL2Id,
  testAccount,
  tokenAddress,
} from '@omni-network/test-utils'
import { renderHook, waitFor } from '@testing-library/react'
import { parseEther, zeroAddress } from 'viem'
import { describe, expect, test } from 'vitest'

import {
  type AnyOrder,
  ContextProvider,
  executeTestOrder,
} from '../test-utils.js'

describe('useQuote()', () => {
  test('parameters: gets a quote in expense mode', async () => {
    const { result } = renderHook(
      () => {
        return useQuote({
          enabled: true,
          mode: 'expense',
          srcChainId: mockL1Id,
          destChainId: mockL2Id,
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
      deposit: { token: zeroAddress, amount: 1n },
      expense: { token: zeroAddress, amount: 0n },
    })
  })

  test('parameters: gets a quote in deposit mode', async () => {
    const { result } = renderHook(
      () => {
        return useQuote({
          enabled: true,
          mode: 'deposit',
          srcChainId: mockL1Id,
          destChainId: mockL2Id,
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
      deposit: { token: zeroAddress, amount: 2n },
      expense: { token: zeroAddress, amount: 1n },
    })
  })

  // Test vector folder: solver/app/testdata/TestQuote/invalid_deposit_(chain_mismatch)
  test('behaviour: handles chain mismatch error', async () => {
    const { result } = renderHook(
      () => {
        return useQuote({
          enabled: true,
          mode: 'expense',
          srcChainId: 1,
          destChainId: 17000,
          deposit: { isNative: true, amount: parseEther('1') },
          expense: { isNative: true },
        })
      },
      { wrapper: ContextProvider },
    )

    await waitFor(() => expect(result.current.isError).toBe(true))
    if (result.current.isError) {
      expect(result.current.error).toEqual({
        code: 400,
        status: 'Bad Request',
        message:
          'InvalidDeposit: deposit and expense must be of the same chain class (e.g. mainnet, testnet) [deposit=1 ETH, expense=ETH]',
      })
    }
  })

  // Test vector folder: solver/app/testdata/TestQuote/no_deposit_of_expense_amount_specified
  test('behaviour: handles invalid deposit or expense amount error', async () => {
    const { result } = renderHook(
      () => {
        return useQuote({
          enabled: true,
          mode: 'expense',
          srcChainId: 1,
          destChainId: 42161,
          deposit: { isNative: true },
          expense: { isNative: true },
        })
      },
      { wrapper: ContextProvider },
    )

    await waitFor(() => expect(result.current.isError).toBe(true))
    if (result.current.isError) {
      expect(result.current.error).toEqual({
        code: 400,
        status: 'Bad Request',
        message:
          'deposit and expense amount cannot be both zero or both non-zero',
      })
    }
  })
})

describe('useValidateOrder()', () => {
  test('default: returns the "accepted" status if the validation is successful', async () => {
    const oneEth = parseEther('1')
    const amount = oneEth / 2n
    const order: AnyOrder = {
      srcChainId: mockL1Id,
      destChainId: mockL2Id,
      expense: { token: zeroAddress, amount },
      deposit: { token: zeroAddress, amount: oneEth },
      calls: [{ target: testAccount.address, value: amount }],
    }

    const { result } = renderHook(
      () => useValidateOrder({ enabled: true, order }),
      { wrapper: ContextProvider },
    )

    await waitFor(() => expect(result.current.status).toBe('accepted'))
  })

  test('behaviour: returns the "rejected" status with a rejection reason and description', async () => {
    const oneEth = parseEther('1')
    const amount = oneEth / 2n
    const order: AnyOrder = {
      srcChainId: invalidChainId,
      destChainId: mockL2Id,
      expense: { token: zeroAddress, amount },
      deposit: { token: zeroAddress, amount: oneEth },
      calls: [{ target: testAccount.address, value: amount }],
    }

    const { result } = renderHook(
      () => useValidateOrder({ enabled: true, order }),
      { wrapper: ContextProvider },
    )

    await waitFor(() => expect(result.current.status).toBe('rejected'))
    if (result.current.status === 'rejected') {
      expect(result.current.rejectReason).toBe('UnsupportedSrcChain')
      expect(result.current.rejectDescription).toBe(
        'unsupported source chain [chain_id=1234]',
      )
    }
  })
})

describe('useOrder()', () => {
  test('default: succeeds with valid order', async () => {
    const oneEth = parseEther('1')
    const amount = oneEth / 2n
    const order: AnyOrder = {
      srcChainId: mockL1Id,
      destChainId: mockL2Id,
      expense: { token: zeroAddress, amount },
      deposit: { token: zeroAddress, amount: oneEth },
      calls: [{ target: testAccount.address, value: amount }],
    }
    await executeTestOrder(order)
  })

  test('behaviour: rejects when using invalid source chain', async () => {
    const order: AnyOrder = {
      srcChainId: invalidChainId,
      destChainId: mockL1Id,
      expense: { token: zeroAddress, amount: 1n },
      deposit: { token: zeroAddress, amount: 1n },
      calls: [{ target: testAccount.address, value: 1n }],
    }
    await executeTestOrder(order, 'UnsupportedSrcChain')
  })

  test('behaviour: rejects when using invalid destination chain', async () => {
    const order: AnyOrder = {
      srcChainId: mockL1Id,
      destChainId: invalidChainId,
      expense: { token: zeroAddress, amount: 1n },
      deposit: { token: zeroAddress, amount: 1n },
      calls: [{ target: testAccount.address, value: 1n }],
    }
    await executeTestOrder(order, 'UnsupportedDestChain')
  })

  test('behaviour: rejects when source and destination chains are the same', async () => {
    const order: AnyOrder = {
      srcChainId: mockL1Id,
      destChainId: mockL1Id,
      expense: { token: zeroAddress, amount: 1n },
      deposit: { token: zeroAddress, amount: 1n },
      calls: [{ target: testAccount.address, value: 1n }],
    }
    await executeTestOrder(order, 'SameChain')
  })

  test('behaviour: rejects when using an unsupported expense token', async () => {
    const order: AnyOrder = {
      srcChainId: mockL1Id,
      destChainId: mockL2Id,
      expense: { token: invalidTokenAddress, amount: 1n },
      deposit: { token: tokenAddress, amount: 1n },
      calls: [{ target: testAccount.address, value: 1n }],
    }
    await executeTestOrder(order, 'UnsupportedExpense')
  })
})
