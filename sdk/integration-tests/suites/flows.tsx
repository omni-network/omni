import { type Quote, useQuote, useValidateOrder } from '@omni-network/react'
import {
  invalidTokenAddress,
  mintOMNI,
  mockL1Id,
  mockL2Id,
  omniDevnetId,
  testAccount,
  tokenAddress,
} from '@omni-network/test-utils'
import { act, waitFor } from '@testing-library/react'
import { parseEther, zeroAddress } from 'viem'
import { describe, expect, test } from 'vitest'

import {
  type AnyOrder,
  createRenderHook,
  executeTestOrder,
  testConnector,
  useOrderRef,
} from '../test-utils.js'

describe('ERC20 OMNI to native OMNI transfer orders', () => {
  test('default: succeeds with valid expense', async () => {
    await mintOMNI()
    const amount = parseEther('10')
    const order: AnyOrder = {
      owner: testAccount.address,
      srcChainId: mockL1Id,
      destChainId: omniDevnetId,
      expense: { token: zeroAddress, amount },
      calls: [{ target: testAccount.address, value: amount }],
      deposit: { token: tokenAddress, amount },
    }
    await executeTestOrder(order)
  }, 30_000)

  test('behaviour: native ETH to native OMNI swap succeeds', async () => {
    const amount = parseEther('10')
    const order: AnyOrder = {
      owner: testAccount.address,
      srcChainId: mockL1Id,
      destChainId: omniDevnetId,
      expense: { token: zeroAddress, amount },
      calls: [{ target: testAccount.address, value: amount }],
      deposit: { token: zeroAddress, amount },
    }
    await executeTestOrder(order)
  })

  test('behaviour: fails with unsupported ERC20 deposit', async () => {
    const amount = parseEther('10')
    const order: AnyOrder = {
      owner: testAccount.address,
      srcChainId: mockL1Id,
      destChainId: omniDevnetId,
      expense: { token: zeroAddress, amount },
      calls: [{ target: testAccount.address, value: amount }],
      deposit: { token: invalidTokenAddress, amount },
    }
    await executeTestOrder(order, 'UnsupportedDeposit')
  })
})

describe('ETH transfer orders', () => {
  test('default: succeeds with valid expense', async () => {
    const account = testAccount
    const amount = parseEther('1')
    const order: AnyOrder = {
      owner: account.address,
      srcChainId: mockL1Id,
      destChainId: mockL2Id,
      expense: { token: zeroAddress, amount },
      calls: [{ target: account.address, value: amount }],
      deposit: { token: zeroAddress, amount: 2n * amount },
    }
    await executeTestOrder(order)
  })

  test('behaviour: fails with expense over max amount', async () => {
    const account = testAccount
    const amount = parseEther('1000')
    const order: AnyOrder = {
      owner: account.address,
      srcChainId: mockL1Id,
      destChainId: mockL2Id,
      expense: { token: zeroAddress, amount },
      calls: [{ target: account.address, value: amount }],
      deposit: { token: zeroAddress, amount: amount + parseEther('1') },
    }
    await executeTestOrder(order, 'ExpenseOverMax')
  })

  test('behaviour: fails with expense under min amount', async () => {
    const account = testAccount
    const amount = 1n
    const order: AnyOrder = {
      owner: account.address,
      srcChainId: mockL1Id,
      destChainId: mockL2Id,
      expense: { token: zeroAddress, amount },
      calls: [{ target: account.address, value: amount }],
      deposit: { token: zeroAddress, amount: 2n * amount },
    }
    await executeTestOrder(order, 'ExpenseUnderMin')
  })
})

test('default: successfully processes order from quote to filled', async () => {
  const renderHook = createRenderHook()
  const oneEth = parseEther('1')
  const twoEth = parseEther('2')

  const quoteHook = renderHook(() => {
    return useQuote({
      enabled: true,
      mode: 'expense',
      srcChainId: mockL1Id,
      destChainId: mockL2Id,
      deposit: {
        amount: twoEth,
        isNative: true,
      },
      expense: {
        amount: oneEth,
        isNative: true,
      },
    })
  })
  await waitFor(() => expect(quoteHook.result.current.isSuccess).toBe(true))

  const quote = quoteHook.result.current.query.data as Quote
  expect(quote).toEqual({
    deposit: { token: zeroAddress, amount: twoEth },
    expense: { token: zeroAddress, amount: expect.any(BigInt) },
  })
  expect(quote.expense.amount).toBeLessThan(twoEth)

  const orderParams = {
    deposit: { token: zeroAddress, amount: twoEth },
    expense: { token: zeroAddress, amount: oneEth },
    calls: [{ target: testAccount.address, value: oneEth }],
    srcChainId: mockL1Id,
    destChainId: mockL2Id,
    validateEnabled: false,
  }

  const validateHook = renderHook(() => {
    return useValidateOrder({ enabled: true, order: orderParams })
  })
  await waitFor(() => {
    return expect(validateHook.result.current.status === 'accepted').toBe(true)
  })

  const orderRef = useOrderRef(testConnector, orderParams)
  await waitFor(() => expect(orderRef.current?.isReady).toBe(true))

  act(() => {
    orderRef.current?.open()
  })
  await waitFor(() => expect(orderRef.current?.txHash).toBeDefined())
})
