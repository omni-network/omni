import * as core from '@omni-network/core'
import { waitFor } from '@testing-library/react'
import { beforeEach, expect, test, vi } from 'vitest'
import {
  contracts,
  createMockWaitForTransactionReceiptResult,
  orderRequest,
  renderHook,
  resolvedOrder,
} from '../../test/index.js'
import { useOrder } from './useOrder.js'

const {
  useValidateOrder,
  useGetOrderStatus,
  useOmniContracts,
  useParseOpenEvent,
  useWaitForTransactionReceipt,
  sendOrder,
  getConnectorClient,
} = vi.hoisted(() => {
  return {
    useValidateOrder: vi.fn(),
    useParseOpenEvent: vi.fn(),
    useGetOrderStatus: vi.fn(),
    useWaitForTransactionReceipt: vi.fn(),
    useOmniContracts: vi.fn().mockImplementation(() => {
      return {
        data: {
          ...contracts,
        },
      }
    }),
    getConnectorClient: vi.fn().mockImplementation(() => {
      return {
        account: '0xAccount',
        chain: '0xChain',
        connector: '0xConnector',
      }
    }),
    sendOrder: vi.fn().mockImplementation(() => {
      return Promise.resolve('0xTxHash')
    }),
  }
})

vi.mock('wagmi', async () => {
  const actual = await vi.importActual('wagmi')
  return {
    ...actual,
    useWaitForTransactionReceipt,
  }
})

vi.mock('wagmi/actions', async () => {
  const actual = await vi.importActual('wagmi/actions')
  return {
    ...actual,
    getConnectorClient,
  }
})

vi.mock('./useValidateOrder.js', async () => {
  return {
    useValidateOrder,
  }
})

vi.mock('./useOmniContracts.js', async () => {
  return {
    useOmniContracts,
  }
})

vi.mock('./useGetOrderStatus.js', async () => {
  return {
    useGetOrderStatus,
  }
})

vi.mock('./useParseOpenEvent.js', async () => {
  return {
    useParseOpenEvent,
  }
})

beforeEach(() => {
  vi.spyOn(core, 'sendOrder').mockResolvedValue('0xTxHash')
  useParseOpenEvent.mockReturnValue({
    resolvedOrder,
    error: null,
  })
  useGetOrderStatus.mockReturnValue({
    status: 'not-found',
  })
  useValidateOrder.mockReturnValue({
    status: 'pending',
  })
  useWaitForTransactionReceipt.mockReturnValue(
    createMockWaitForTransactionReceiptResult({
      isPending: true,
      isSuccess: false,
      data: undefined,
      status: 'pending',
    }),
  )
})

const renderOrderHook = (
  params: Parameters<typeof useOrder>[0],
  options?: Parameters<typeof renderHook>[1],
) => {
  return renderHook(() => useOrder({ ...params }), {
    mockContractsCall: true,
    ...options,
  })
}

test(`default: validates, opens, and transitions order through it's lifecycle`, async () => {
  vi.mock('@omni-network/core', async () => {
    const actual = await vi.importActual('@omni-network/core')
    return { ...actual, sendOrder }
  })

  const { result, rerender } = renderHook(
    ({ validateEnabled }: { validateEnabled: boolean }) =>
      useOrder({ ...orderRequest, validateEnabled }),
    { mockContractsCall: true, initialProps: { validateEnabled: true } },
  )

  await waitFor(() => {
    expect(result.current.isReady).toBe(true)
    expect(result.current.isValidated).toBe(false)
    expect(result.current.isError).toBe(false)
    expect(result.current.isTxPending).toBe(false)
    expect(result.current.status).toBe('ready')
    expect(result.current.validation?.status).toBe('pending')
    expect(result.current.isTxSubmitted).toBe(false)
    expect(result.current.txMutation.data).toBeUndefined()
    expect(result.current.isOpen).toBe(false)
    expect(result.current.txHash).toBeUndefined()
    expect(result.current.error).toBeUndefined()
  })

  useValidateOrder.mockReturnValue({
    status: 'accepted',
  })

  useWaitForTransactionReceipt.mockImplementation(() =>
    createMockWaitForTransactionReceiptResult({
      isPending: true,
      isSuccess: false,
      data: undefined,
      status: 'pending',
      fetchStatus: 'fetching',
    }),
  )

  rerender({ validateEnabled: true })

  await waitFor(() => {
    expect(result.current.isValidated).toBe(true)
    expect(result.current.validation?.status).toBe('accepted')
  })

  result.current.open()

  await waitFor(() => {
    expect(result.current.txHash).toBe('0xTxHash')
    expect(result.current.isTxPending).toBe(false)
    expect(result.current.isTxSubmitted).toBe(true)
    expect(result.current.txMutation.data).toBe('0xTxHash')
    expect(result.current.txMutation.isSuccess).toBe(true)
    expect(result.current.status).toBe('opening')
    expect(result.current.waitForTx.data).toBeUndefined()
    expect(result.current.waitForTx.isSuccess).toBe(false)
  })

  useWaitForTransactionReceipt.mockImplementation(() =>
    createMockWaitForTransactionReceiptResult(),
  )

  useGetOrderStatus.mockReturnValue({
    status: 'open',
  })

  rerender({ validateEnabled: true })

  await waitFor(() => {
    expect(result.current.waitForTx.data).toBe('0xTxHash')
    expect(result.current.waitForTx.isSuccess).toBe(true)
    expect(result.current.isOpen).toBe(true)
  })

  useGetOrderStatus.mockReturnValue({
    status: 'filled',
    destTxHash: '0x123',
  })

  rerender({ validateEnabled: true })

  await waitFor(() => {
    expect(result.current.status).toBe('filled')
    expect(result.current.destTxHash).toBe('0x123')
  })
})

test('parameters: debugValidation is passed to useValidateOrder', async () => {
  renderOrderHook({
    ...orderRequest,
    validateEnabled: true,
    debugValidation: true,
  })
  expect(useValidateOrder).toHaveBeenCalledWith(
    expect.objectContaining({ order: orderRequest, debug: true }),
  )
})

test('parameters: omniContractsQueryOpts, getOrderQueryOpts and didFillQueryOpts are passed to the relevant hooks', async () => {
  const omniContractsQueryOpts = { staleTime: 5000 }
  const getOrderQueryOpts = { staleTime: 1000 }
  const didFillQueryOpts = { staleTime: 2000 }

  const { result } = renderOrderHook({
    ...orderRequest,
    validateEnabled: false,
    omniContractsQueryOpts,
    getOrderQueryOpts,
    didFillQueryOpts,
  })

  await waitFor(() => {
    expect(result.current.isReady).toBe(true)
  })
  result.current.open()

  expect(useOmniContracts).toHaveBeenCalledWith(
    expect.objectContaining({ queryOpts: omniContractsQueryOpts }),
  )
  expect(useGetOrderStatus).toHaveBeenCalledWith(
    expect.objectContaining({ getOrderQueryOpts, didFillQueryOpts }),
  )
})

test('behaviour: handles order rejection', async () => {
  const { result, rerender } = renderOrderHook({
    ...orderRequest,
    validateEnabled: false,
  })

  useWaitForTransactionReceipt.mockImplementation(() =>
    createMockWaitForTransactionReceiptResult(),
  )

  result.current.open()

  useGetOrderStatus.mockReturnValue({
    status: 'rejected',
  })

  rerender()

  await waitFor(() => {
    expect(result.current.status).toBe('rejected')
    expect(result.current.isOpen).toBe(false)
  })
})

test('behaviour: closed order is handled', async () => {
  const { result, rerender } = renderOrderHook({
    ...orderRequest,
    validateEnabled: false,
  })

  useWaitForTransactionReceipt.mockImplementation(() =>
    createMockWaitForTransactionReceiptResult(),
  )

  result.current.open()

  useGetOrderStatus.mockReturnValue({
    status: 'closed',
  })

  rerender()

  await waitFor(() => {
    expect(result.current.status).toBe('closed')
    expect(result.current.isOpen).toBe(false)
  })
})

test('behaviour: handles sendOrder error', async () => {
  vi.spyOn(core, 'sendOrder').mockRejectedValue(new Error('Tx mutation error'))

  const { result } = renderOrderHook({
    ...orderRequest,
    validateEnabled: false,
  })

  result.current.open()

  await waitFor(() => {
    expect(result.current.isError).toBe(true)
    expect(result.current.error).toBeInstanceOf(core.OpenError)
  })
})

test('behaviour: validate false/true: toggles validation behavior', async () => {
  useValidateOrder.mockReturnValue({
    status: 'pending',
  })

  const { result, rerender } = renderOrderHook({
    ...orderRequest,
    validateEnabled: false,
  })

  await waitFor(() => {
    expect(result.current.isValidated).toBe(false)
  })

  useValidateOrder.mockReturnValue({
    status: 'accepted',
  })

  rerender({ validateEnabled: true })

  await waitFor(() => {
    expect(result.current.isValidated).toBe(true)
  })
})

test('behaviour:  handles validation error', async () => {
  useValidateOrder.mockReturnValue({
    status: 'error',
    error: new Error('Validation failed'),
  })

  const { result } = renderOrderHook({
    ...orderRequest,
    validateEnabled: true,
  })

  await waitFor(() => {
    expect(result.current.isError).toBe(true)
    expect(result.current.error).toBeInstanceOf(core.ValidateOrderError)
  })
})

test('behaviour: handles transaction receipt error', async () => {
  useWaitForTransactionReceipt.mockImplementation(() =>
    createMockWaitForTransactionReceiptResult({
      isError: true,
      error: new Error('Receipt fetch failed'),
      status: 'error',
    }),
  )

  const { result } = renderOrderHook({
    ...orderRequest,
    validateEnabled: false,
  })

  await waitFor(() => {
    expect(result.current.isError).toBe(true)
    expect(result.current.error).toBeInstanceOf(core.TxReceiptError)
    expect(result.current.status).toBe('error')
  })
})

test('behaviour: handles parse open event error', async () => {
  useParseOpenEvent.mockReturnValue({
    status: 'error',
    error: new core.ParseOpenEventError('Failed to parse open event'),
  })

  useWaitForTransactionReceipt.mockImplementation(() =>
    createMockWaitForTransactionReceiptResult({
      error: null,
    }),
  )

  const { result } = renderOrderHook({
    ...orderRequest,
    validateEnabled: false,
  })

  await waitFor(() => {
    expect(result.current.isError).toBe(true)
    expect(result.current.error).toBeInstanceOf(core.ParseOpenEventError)
  })
})

test('behaviour: handles order status error', async () => {
  useGetOrderStatus.mockReturnValue({
    status: 'error',
    error: new core.WatchDidFillError('Failed to get order status'),
  })

  useWaitForTransactionReceipt.mockImplementation(() =>
    createMockWaitForTransactionReceiptResult({
      isSuccess: true,
      status: 'success',
    }),
  )

  const { result } = renderOrderHook({
    ...orderRequest,
    validateEnabled: false,
  })

  await waitFor(() => {
    expect(result.current.isError).toBe(true)
    expect(result.current.error).toBeInstanceOf(core.WatchDidFillError)
  })
})

test('behaviour: handles order status error', async () => {
  useGetOrderStatus.mockReturnValue({
    status: 'error',
    error: new core.DidFillError('Failed to get order status'),
  })

  useWaitForTransactionReceipt.mockImplementation(() =>
    createMockWaitForTransactionReceiptResult({
      isSuccess: true,
      status: 'success',
    }),
  )

  const { result } = renderOrderHook({
    ...orderRequest,
    validateEnabled: false,
  })

  await waitFor(() => {
    expect(result.current.isError).toBe(true)
    expect(result.current.error).toBeInstanceOf(core.DidFillError)
  })
})

test('behaviour: handles wait success but order not found', async () => {
  useWaitForTransactionReceipt.mockImplementation(() =>
    createMockWaitForTransactionReceiptResult({
      isSuccess: true,
      status: 'success',
    }),
  )

  const { result } = renderOrderHook({
    ...orderRequest,
    validateEnabled: true,
  })

  await waitFor(() => {
    expect(result.current.isError).toBe(true)
    expect(result.current.error).toBeInstanceOf(core.GetOrderError)
  })
})

test('behaviour: handles contracts load error', async () => {
  useOmniContracts.mockImplementation(() => {
    return {
      data: {
        inbox: undefined,
      },
      error: new Error('Failed to load contracts'),
      isError: true,
    }
  })

  const { result } = renderOrderHook(
    { ...orderRequest, validateEnabled: true },
    { mockContractsCallFailure: true },
  )

  await waitFor(() => {
    expect(result.current.isError).toBe(true)
    expect(result.current.error).toBeInstanceOf(core.LoadContractsError)
    expect(result.current.status).toBe('error')
  })
})

test('behaviour: open rejects when inbox contract not loaded', async () => {
  const { result, rerender } = renderOrderHook(
    { ...orderRequest, validateEnabled: false },
    { mockContractsCall: false },
  )

  useOmniContracts.mockImplementation(() => {
    return {
      data: {
        inbox: undefined,
      },
    }
  })

  rerender()

  result.current.open()

  await waitFor(() => {
    expect(result.current.txMutation.isError).toBe(true)
  })
})
