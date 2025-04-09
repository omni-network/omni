import { waitFor } from '@testing-library/react'
import { beforeEach, expect, test, vi } from 'vitest'
import {
  createMockWaitForTransactionReceiptResult,
  createMockWriteContractResult,
  mockWagmiHooks,
} from '../../test/mocks.js'
import { renderHook } from '../../test/react.js'
import { contracts, orderRequest, resolvedOrder } from '../../test/shared.js'
import {
  DidFillError,
  GetOrderError,
  LoadContractsError,
  OpenError,
  ParseOpenEventError,
  TxReceiptError,
  ValidateOrderError,
} from '../errors/base.js'
import { useOrder } from './useOrder.js'

const { useWriteContract, useWaitForTransactionReceipt } = mockWagmiHooks()

const {
  useValidateOrder,
  useGetOrderStatus,
  useOmniContracts,
  useParseOpenEvent,
} = vi.hoisted(() => {
  return {
    useValidateOrder: vi.fn(),
    useParseOpenEvent: vi.fn(),
    useGetOrderStatus: vi.fn(),
    useOmniContracts: vi.fn().mockImplementation(() => {
      return {
        data: {
          ...contracts,
        },
      }
    }),
  }
})

vi.mock('./useValidateOrder.js', async () => {
  return {
    useValidateOrder: useValidateOrder,
  }
})

vi.mock('./useOmniContracts.js', async () => {
  return {
    useOmniContracts: useOmniContracts,
  }
})

vi.mock('./useGetOrderStatus.js', async () => {
  return {
    useGetOrderStatus: useGetOrderStatus,
  }
})

vi.mock('./useParseOpenEvent.js', async () => {
  return {
    useParseOpenEvent: useParseOpenEvent,
  }
})

beforeEach(() => {
  useParseOpenEvent.mockReturnValue({
    resolvedOrder,
    error: null,
  })
  useGetOrderStatus.mockReturnValue({
    status: 'not-found',
  })
  useValidateOrder.mockReturnValue({
    status: 'accepted',
  })
  useWriteContract.mockReturnValue(
    createMockWriteContractResult({
      isPending: true,
      isSuccess: false,
      data: undefined,
      isIdle: false,
      status: 'pending',
    }),
  )
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
  const { result, rerender } = renderHook(
    ({ validateEnabled }: { validateEnabled: boolean }) =>
      useOrder({ ...orderRequest, validateEnabled }),
    { mockContractsCall: true, initialProps: { validateEnabled: true } },
  )

  await waitFor(() => {
    expect(result.current.isReady).toBe(true)
    expect(result.current.isValidated).toBe(true)
    expect(result.current.isError).toBe(false)
    expect(result.current.isTxPending).toBe(true)
    expect(result.current.isTxSubmitted).toBe(false)
    expect(result.current.txMutation.data).toBeUndefined()
    expect(result.current.isOpen).toBe(false)
    expect(result.current.txHash).toBeUndefined()
    expect(result.current.error).toBeUndefined()
  })

  useValidateOrder.mockReturnValue({
    status: 'accepted',
  })

  useGetOrderStatus.mockReturnValue({
    status: 'open',
  })

  useWriteContract.mockImplementation(() => createMockWriteContractResult())

  useWaitForTransactionReceipt.mockImplementation(() =>
    createMockWaitForTransactionReceiptResult(),
  )

  rerender({ validateEnabled: true })

  const res = await result.current.open()

  await waitFor(() => {
    expect(result.current.isOpen).toBe(true)
    expect(result.current.isTxPending).toBe(false)
    expect(result.current.isTxSubmitted).toBe(true)
    expect(result.current.txMutation.data).toBe('0xTxHash')
    expect(result.current.txMutation.isSuccess).toBe(true)
    expect(res).toBe('0xTxHash')
  })

  useGetOrderStatus.mockReturnValue({
    status: 'filled',
  })

  rerender({ validateEnabled: true })

  await waitFor(() => expect(result.current.status).toBe('filled'))
})

test('behaviour: handles order rejection', async () => {
  const { result, rerender } = renderOrderHook({
    ...orderRequest,
    validateEnabled: false,
  })

  useWriteContract.mockImplementation(() => createMockWriteContractResult())

  useWaitForTransactionReceipt.mockImplementation(() =>
    createMockWaitForTransactionReceiptResult(),
  )

  await result.current.open()

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

  useWriteContract.mockImplementation(() => createMockWriteContractResult())

  useWaitForTransactionReceipt.mockImplementation(() =>
    createMockWaitForTransactionReceiptResult(),
  )

  await result.current.open()

  useGetOrderStatus.mockReturnValue({
    status: 'closed',
  })

  rerender()

  await waitFor(() => {
    expect(result.current.status).toBe('closed')
    expect(result.current.isOpen).toBe(false)
  })
})

test('behaviour: handles tx mutation error', async () => {
  useWriteContract.mockImplementation(() =>
    createMockWriteContractResult({
      isError: true,
      error: new Error('Transaction failed'),
      status: 'error',
    }),
  )

  const { result } = renderOrderHook({
    ...orderRequest,
    validateEnabled: false,
  })

  await waitFor(() => {
    expect(result.current.isError).toBe(true)
    expect(result.current.error).toBeInstanceOf(OpenError)
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

  const { result } = renderOrderHook({ ...orderRequest, validateEnabled: true })

  await waitFor(() => {
    expect(result.current.isError).toBe(true)
    expect(result.current.error).toBeInstanceOf(ValidateOrderError)
  })
})

test('behaviour: handles transaction receipt error', async () => {
  useWriteContract.mockImplementation(() =>
    createMockWriteContractResult({
      error: null,
    }),
  )

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
    expect(result.current.error).toBeInstanceOf(TxReceiptError)
    expect(result.current.status).toBe('error')
  })
})

test('behaviour: handles parse open event error', async () => {
  useParseOpenEvent.mockReturnValue({
    status: 'error',
    error: new ParseOpenEventError('Failed to parse open event'),
  })

  useWaitForTransactionReceipt.mockImplementation(() =>
    createMockWaitForTransactionReceiptResult({
      error: null,
    }),
  )

  useWriteContract.mockReset().mockImplementation(() =>
    createMockWriteContractResult({
      isSuccess: true,
      status: 'success',
      error: null,
    }),
  )

  const { result } = renderOrderHook({
    ...orderRequest,
    validateEnabled: false,
  })

  await waitFor(() => {
    expect(result.current.isError).toBe(true)
    expect(result.current.error).toBeInstanceOf(ParseOpenEventError)
  })
})

test('behaviour: handles order status error', async () => {
  useGetOrderStatus.mockReturnValue({
    status: 'error',
    error: new DidFillError('Failed to get order status'),
  })

  useWriteContract.mockReset().mockImplementation(() =>
    createMockWriteContractResult({
      isSuccess: true,
      status: 'success',
    }),
  )

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
    expect(result.current.error).toBeInstanceOf(DidFillError)
  })
})

test('behaviour: handles wait success but order not found', async () => {
  useWriteContract.mockReset().mockImplementation(() =>
    createMockWriteContractResult({
      isSuccess: true,
      status: 'success',
    }),
  )

  useWaitForTransactionReceipt.mockImplementation(() =>
    createMockWaitForTransactionReceiptResult({
      isSuccess: true,
      status: 'success',
    }),
  )

  const { result } = renderOrderHook({ ...orderRequest, validateEnabled: true })

  await waitFor(() => {
    expect(result.current.isError).toBe(true)
    expect(result.current.error).toBeInstanceOf(GetOrderError)
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
    expect(result.current.error).toBeInstanceOf(LoadContractsError)
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

  await expect(
    result.current.open(),
  ).rejects.toThrowErrorMatchingInlineSnapshot(
    '[LoadContractsError: Inbox contract address needs to be loaded]',
  )
})
