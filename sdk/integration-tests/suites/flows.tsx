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
import { useBalance } from 'wagmi'
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

  test('behaviour: fails with native deposit', async () => {
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

  const preDestBalance = renderHook(() => {
    return useBalance({
      address: testAccount.address,
      chainId: mockL2Id,
    })
  })

  await waitFor(
    () => {
      expect(preDestBalance.result.current.data).toBeDefined()
    },
    { timeout: 5_000 },
  )

  const quoteHook = renderHook(() => {
    return useQuote({
      enabled: true,
      mode: 'expense',
      srcChainId: mockL1Id,
      destChainId: mockL2Id,
      deposit: {
        amount: parseEther('2'),
        isNative: true,
      },
      expense: {
        amount: parseEther('1'),
        isNative: true,
      },
    })
  })

  await waitFor(() => {
    expect(quoteHook.result.current.query.error).toBeNull()
    expect(quoteHook.result.current.query.data).toBeDefined()
    expect(quoteHook.result.current.isSuccess).toBe(true)
    expect(quoteHook.result.current.isPending).toBe(false)
    expect(quoteHook.result.current.isError).toBe(false)
  })

  const quote = quoteHook.result.current.query.data as Quote

  expect(quote).toEqual({
    deposit: { token: zeroAddress, amount: parseEther('2') },
    expense: { token: zeroAddress, amount: expect.any(BigInt) },
  })
  expect(quote.expense.amount).toBeLessThan(parseEther('2'))

  const orderParams = {
    deposit: { token: zeroAddress, amount: parseEther('2') },
    expense: { token: zeroAddress, amount: parseEther('1') },
    calls: [{ target: testAccount.address, value: parseEther('1') }],
    srcChainId: mockL1Id,
    destChainId: mockL2Id,
    validateEnabled: false,
  }

  const validateHook = renderHook(() => {
    return useValidateOrder({ enabled: true, order: orderParams })
  })

  await waitFor(() => {
    expect(validateHook.result.current.status === 'accepted').toBe(true)
  })

  const orderRef = useOrderRef(testConnector, orderParams)
  await waitFor(() => expect(orderRef.current?.isReady).toBe(true))

  act(() => {
    orderRef.current?.open()
  })

  await waitFor(() => {
    expect(orderRef.current?.status).toBeOneOf(['open', 'opening'])
    expect(orderRef.current?.txHash).toBeDefined()
    expect(orderRef.current?.error).toBeUndefined()
    expect(orderRef.current?.isError).toBe(false)
    expect(orderRef.current?.isTxSubmitted).toBe(true)
    expect(orderRef.current?.isTxPending).toBe(false)
  })

  await waitFor(() => expect(orderRef.current?.status).toBe('filled'), {
    timeout: 10_000,
  })

  const postDestBalance = renderHook(() => {
    return useBalance({
      address: testAccount.address,
      chainId: mockL2Id,
    })
  })

  await waitFor(
    () => {
      expect(postDestBalance.result.current.data).toBeDefined()
      expect(postDestBalance.result.current.data?.value).toBe(
        // biome-ignore lint/style/noNonNullAssertion: safe due to throwing condition above
        preDestBalance.result.current.data?.value! + orderParams.calls[0].value,
      )
    },
    { timeout: 5_000 },
  )
})
