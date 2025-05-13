import {
  useGetOrder,
  useOmniAssets,
  useOmniContracts,
  useParseOpenEvent,
  useQuote,
  useValidateOrder,
} from '@omni-network/react'
import {
  createAnvilClient,
  invalidChainId,
  invalidTokenAddress,
  mockL1Chain,
  mockL1Id,
  mockL2Chain,
  mockL2Id,
  testAccount,
  tokenAddress,
} from '@omni-network/test-utils'
import { act, renderHook, waitFor } from '@testing-library/react'
import { pad, parseEther, zeroAddress } from 'viem'
import { beforeAll, describe, expect, test } from 'vitest'
import {
  type AnyOrder,
  ContextProvider,
  assertResolvedOrder,
  createRenderHook,
  executeTestOrderUsingReact,
  useOrderRef,
} from '../test-utils.js'

beforeAll(async () => {
  const value = parseEther('1000')
  await Promise.all([
    createAnvilClient(mockL1Chain).setBalance({
      address: testAccount.address,
      value,
    }),
    createAnvilClient(mockL2Chain).setBalance({
      address: testAccount.address,
      value,
    }),
  ])
})

async function execOrder() {
  const orderParams = {
    deposit: { token: zeroAddress, amount: parseEther('2') },
    expense: { token: zeroAddress, amount: parseEther('1') },
    calls: [{ target: testAccount.address, value: parseEther('1') }],
    srcChainId: mockL2Id,
    destChainId: mockL1Id,
    validateEnabled: false,
  }

  const orderRef = useOrderRef(orderParams)

  await waitFor(() => expect(orderRef.current?.isReady).toBe(true))

  act(() => {
    orderRef.current?.open()
  })

  return orderRef
}

describe.concurrent('useQuote()', () => {
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

    await waitFor(() => {
      expect(result.current.isSuccess).toBe(true)
      expect(result.current.isError).toBe(false)
      expect(result.current.isPending).toBe(false)
      expect(result.current.query.data).toEqual({
        deposit: { token: zeroAddress, amount: 1n },
        expense: { token: zeroAddress, amount: 0n },
      })
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

    await waitFor(() => {
      expect(result.current.isSuccess).toBe(true)
      expect(result.current.isError).toBe(false)
      expect(result.current.isPending).toBe(false)
      expect(result.current.query.data).toEqual({
        deposit: { token: zeroAddress, amount: 2n },
        expense: { token: zeroAddress, amount: 1n },
      })
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

    if (!result.current.isError) throw new Error('We expect an error')

    expect(result.current.error).toEqual({
      code: 400,
      status: 'Bad Request',
      message:
        'InvalidDeposit: deposit and expense must be of the same chain class (e.g. mainnet, testnet) [deposit=mainnet, expense=testnet]',
    })
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

    if (!result.current.isError) throw new Error('We expect an error')

    expect(result.current.error).toEqual({
      code: 400,
      status: 'Bad Request',
      message:
        'deposit and expense amount cannot be both zero or both non-zero',
    })
  })
})

describe.concurrent('useValidateOrder()', () => {
  test('default: returns the "accepted" status if the validation is successful', async () => {
    const amount = parseEther('1') / 2n
    const order: AnyOrder = {
      srcChainId: mockL1Id,
      destChainId: mockL2Id,
      expense: { token: zeroAddress, amount },
      deposit: { token: zeroAddress, amount: parseEther('1') },
      calls: [{ target: testAccount.address, value: amount }],
    }

    const { result } = renderHook(
      () => useValidateOrder({ enabled: true, order }),
      { wrapper: ContextProvider },
    )

    await waitFor(() => expect(result.current.status).toBe('accepted'))
  })

  test('behaviour: returns the "rejected" status with a rejection reason and description', async () => {
    const amount = parseEther('1') / 2n
    const order: AnyOrder = {
      srcChainId: invalidChainId,
      destChainId: mockL2Id,
      expense: { token: zeroAddress, amount },
      deposit: { token: zeroAddress, amount: parseEther('1') },
      calls: [{ target: testAccount.address, value: amount }],
    }

    const { result } = renderHook(
      () => useValidateOrder({ enabled: true, order }),
      { wrapper: ContextProvider },
    )

    await waitFor(() => expect(result.current.status).toBe('rejected'))

    if (result.current.status !== 'rejected')
      throw new Error('We expect an error')
    expect(result.current.rejectReason).toBe('UnsupportedSrcChain')
    expect(result.current.rejectDescription).toBe(
      'unsupported source chain [chain_id=1234]',
    )
  })
})

describe('useOrder()', () => {
  test('default: succeeds with valid order', async () => {
    const amount = parseEther('1') / 2n
    const order: AnyOrder = {
      srcChainId: mockL1Id,
      destChainId: mockL2Id,
      expense: { token: zeroAddress, amount },
      deposit: { token: zeroAddress, amount: parseEther('1') },
      calls: [{ target: testAccount.address, value: amount }],
    }
    await executeTestOrderUsingReact({ order })
  })

  test('behaviour: rejects when using invalid source chain', async () => {
    const order: AnyOrder = {
      srcChainId: invalidChainId,
      destChainId: mockL1Id,
      expense: { token: zeroAddress, amount: 1n },
      deposit: { token: zeroAddress, amount: 1n },
      calls: [{ target: testAccount.address, value: 1n }],
    }
    await executeTestOrderUsingReact({
      order,
      rejectReason: 'UnsupportedSrcChain',
    })
  })

  test('behaviour: rejects when using invalid destination chain', async () => {
    const order: AnyOrder = {
      srcChainId: mockL1Id,
      destChainId: invalidChainId,
      expense: { token: zeroAddress, amount: 1n },
      deposit: { token: zeroAddress, amount: 1n },
      calls: [{ target: testAccount.address, value: 1n }],
    }
    await executeTestOrderUsingReact({
      order,
      rejectReason: 'UnsupportedDestChain',
    })
  })

  test('behaviour: rejects when source and destination chains are the same', async () => {
    const order: AnyOrder = {
      srcChainId: mockL1Id,
      destChainId: mockL1Id,
      expense: { token: zeroAddress, amount: 1n },
      deposit: { token: zeroAddress, amount: 1n },
      calls: [{ target: testAccount.address, value: 1n }],
    }
    await executeTestOrderUsingReact({ order, rejectReason: 'SameChain' })
  })

  test('behaviour: rejects when using an unsupported expense token', async () => {
    const order: AnyOrder = {
      srcChainId: mockL1Id,
      destChainId: mockL2Id,
      expense: { token: invalidTokenAddress, amount: 1n },
      deposit: { token: tokenAddress, amount: 1n },
      calls: [{ target: testAccount.address, value: 1n }],
    }
    await executeTestOrderUsingReact({
      order,
      rejectReason: 'UnsupportedExpense',
    })
  })
})

describe('useParseOpenEvent()', () => {
  test('default: returns order details from the open event logs', async () => {
    const renderHook = createRenderHook()

    const orderRef = await execOrder()

    await waitFor(
      () => expect(orderRef.current?.waitForTx.status).toBe('success'),
      { timeout: 20_000 },
    )

    const parseOpenEventHook = renderHook(() => {
      return useParseOpenEvent({
        // type assertion is safe due to throwing condition above
        status: orderRef.current?.waitForTx.status as 'success',
        logs: orderRef.current?.waitForTx.data?.logs,
      })
    })

    await waitFor(() => {
      expect(parseOpenEventHook.result.current.resolvedOrder).toBeDefined()
      expect(parseOpenEventHook.result.current.error).toBeUndefined()
    })

    // biome-ignore lint/style/noNonNullAssertion: safe due throwing condition above
    const resolvedOrder = parseOpenEventHook.result.current.resolvedOrder!

    // assert shape of return
    assertResolvedOrder(resolvedOrder)
  })
})

describe('useGetOrder()', () => {
  test('default: returns expected order data from the getOrder inbox contract method', async () => {
    const renderHook = createRenderHook()

    const orderRef = await execOrder()

    await waitFor(
      () => expect(orderRef.current?.waitForTx.status).toBe('success'),
      { timeout: 20_000 },
    )

    const parseOpenEventHook = renderHook(() => {
      return useParseOpenEvent({
        // type assertion is safe due to throwing condition above
        status: orderRef.current?.waitForTx.status as 'success',
        logs: orderRef.current?.waitForTx.data?.logs,
      })
    })

    await waitFor(() => {
      expect(parseOpenEventHook.result.current.resolvedOrder).toBeDefined()
      expect(parseOpenEventHook.result.current.error).toBeUndefined()
    })

    // biome-ignore lint/style/noNonNullAssertion: safe due throwing condition above
    const orderId = parseOpenEventHook.result.current.resolvedOrder?.orderId!

    const getOrderHook = renderHook(() => {
      return useGetOrder({
        chainId: mockL2Id,
        orderId: orderId,
      })
    })

    await waitFor(() => {
      expect(getOrderHook.result.current.data).toBeDefined()
    })

    // biome-ignore lint/style/noNonNullAssertion: safe due to throwing condition above
    const orderDetails = getOrderHook.result.current.data!

    // assert shape of return
    expect(orderDetails).toBeInstanceOf(Array)
    expect(orderDetails).toHaveLength(3)

    const resolvedOrder = orderDetails[0]
    assertResolvedOrder(resolvedOrder)

    // order state
    expect(orderDetails[1]).toBeInstanceOf(Object)
    expect(orderDetails[1].status).toBeOneOf([1, 4, 5]) // open / filled / claimed
    expect(orderDetails[1].rejectReason).toBe(0)
    expect(orderDetails[1].timestamp).toBeTypeOf('number')
    expect(orderDetails[1].updatedBy).toBeTypeOf('string')
    expect(orderDetails[1].updatedBy).toContain('0x')

    // offset
    expect(orderDetails[2]).toBeTypeOf('bigint')
  })

  test('behaviour: returns not found when an invalid order id is provided', async () => {
    const renderHook = createRenderHook()

    const getOrderHook = renderHook(() => {
      return useGetOrder({
        chainId: mockL2Id,
        orderId: pad('0x', { size: 32, dir: 'left' }),
      })
    })

    await waitFor(() => {
      expect(getOrderHook.result.current.isFetched).toBe(true)
      expect(getOrderHook.result.current.isError).toBe(false)
      expect(getOrderHook.result.current.data).toBeDefined()
    })

    // biome-ignore lint/style/noNonNullAssertion: safe due to throwing condition above
    const orderDetails = getOrderHook.result.current.data!

    expect(orderDetails).toBeInstanceOf(Array)
    expect(orderDetails).toHaveLength(3)

    const resolvedOrder = orderDetails[0]
    expect(resolvedOrder.orderId).toBe(pad('0x', { size: 32, dir: 'left' }))
    expect(resolvedOrder.user).toBe(zeroAddress)

    const orderState = orderDetails[1]
    expect(orderState.status).toBe(0) // not-found
  })
})

describe.concurrent('useOmniContracts()', () => {
  test('default: returns the expected contract addresses', async () => {
    const renderHook = createRenderHook()

    const omniContractsHook = renderHook(() => {
      return useOmniContracts()
    })

    await waitFor(() => {
      expect(omniContractsHook.result.current.data).toBeDefined()
    })

    // biome-ignore lint/style/noNonNullAssertion: safe due to throwing condition above
    const omniContracts = omniContractsHook.result.current.data!

    const expectedKeys = ['portal', 'inbox', 'outbox', 'middleman', 'executor']
    expect(Object.keys(omniContracts).sort()).toEqual(expectedKeys.sort())
  })
})

describe.concurrent('useOmniAssets()', () => {
  test('default: returns expected asset shape', async () => {
    const renderHook = createRenderHook()

    const omniAssetsHook = renderHook(() => {
      return useOmniAssets()
    })

    await waitFor(() => {
      expect(omniAssetsHook.result.current.data).toBeDefined()
    })

    // biome-ignore lint/style/noNonNullAssertion: safe due to throwing condition above
    const omniAssets = omniAssetsHook.result.current.data!

    expect(omniAssets).toBeInstanceOf(Array)

    const asset = omniAssets[0]
    expect(asset).toBeInstanceOf(Object)
    const expectedKeys = [
      'address',
      'chainId',
      'decimals',
      'expenseMin',
      'expenseMax',
      'enabled',
      'name',
      'symbol',
    ]
    expect(Object.keys(asset).sort()).toEqual(expectedKeys.sort())
    expect(asset.enabled).toBeTypeOf('boolean')
    expect(asset.name).toBeTypeOf('string')
    expect(asset.symbol).toBeTypeOf('string')
    expect(asset.address).toBeTypeOf('string')
    expect(asset.address.startsWith('0x')).toBe(true)
    expect(asset.decimals).toBeTypeOf('number')
    expect(asset.chainId).toBeTypeOf('number')
    expect(asset.expenseMin).toBeTypeOf('bigint')
    expect(asset.expenseMax).toBeTypeOf('bigint')
    expect(asset.expenseMin).toBeGreaterThan(0n)
    expect(asset.expenseMax).toBeGreaterThan(0n)
  })
})
