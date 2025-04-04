import { waitFor } from '@testing-library/react'
import { beforeEach, expect, test, vi } from 'vitest'
import { orderId, renderHook, resolvedOrder } from '../../test/index.js'
import { createMockReadContractResult } from '../../test/mocks.js'
import { useGetOrder } from './useGetOrder.js'
import { useInboxStatus } from './useInboxStatus.js'

const data = {
  status: 1,
  updatedBy: '0x123',
  timestamp: 1,
  rejectReason: 0,
} as const

const { mockUseGetOrder } = vi.hoisted(() => {
  return {
    mockUseGetOrder: vi.fn().mockImplementation(() => {
      return createMockReadContractResult()
    }),
  }
})

vi.mock('./useGetOrder.js', async () => {
  const actual = await vi.importActual('./useGetOrder.js')
  return {
    ...actual,
    useGetOrder: mockUseGetOrder,
  }
})

beforeEach(() => {
  mockUseGetOrder.mockReturnValue(createMockReadContractResult())
})

test('default: returns appropriate inbox status when order is resolved', async () => {
  const { result, rerender } = renderHook(
    () => useInboxStatus({ chainId: 1 }),
    {
      mockContractsCall: true,
    },
  )

  // once on mount
  expect(useGetOrder).toHaveBeenCalledOnce()
  expect(result.current).toBe('unknown')

  mockUseGetOrder.mockReturnValue(
    createMockReadContractResult<ReturnType<typeof useGetOrder>>({
      data: [resolvedOrder, data, 0n],
      isSuccess: true,
      status: 'success',
    }),
  )

  rerender({
    chainId: 1,
    orderId,
  })

  await waitFor(() => expect(result.current).toBe('open'))
})

test('parameters: status unknown', () => {
  const { result } = renderHook(() => useInboxStatus({ chainId: 1 }), {
    mockContractsCall: true,
  })

  expect(result.current).toBe('unknown')
})

test('parameters: status open', () => {
  mockUseGetOrder.mockReturnValue(
    createMockReadContractResult<ReturnType<typeof useGetOrder>>({
      data: [resolvedOrder, data, 0n],
    }),
  )

  const { result } = renderHook(() => useInboxStatus({ chainId: 1 }), {
    mockContractsCall: true,
  })

  expect(result.current).toBe('open')
})

test('parameters: status rejected', () => {
  mockUseGetOrder.mockReturnValue(
    createMockReadContractResult<ReturnType<typeof useGetOrder>>({
      data: [resolvedOrder, { ...data, status: 2 }, 0n],
    }),
  )

  const { result } = renderHook(() => useInboxStatus({ chainId: 1 }), {
    mockContractsCall: true,
  })

  expect(result.current).toBe('rejected')
})

test('parameters: status closed', () => {
  mockUseGetOrder.mockReturnValue(
    createMockReadContractResult<ReturnType<typeof useGetOrder>>({
      data: [resolvedOrder, { ...data, status: 3 }, 0n],
    }),
  )

  const { result } = renderHook(() => useInboxStatus({ chainId: 1 }), {
    mockContractsCall: true,
  })

  expect(result.current).toBe('closed')
})

test('parameters: status filled', () => {
  mockUseGetOrder.mockReturnValue(
    createMockReadContractResult<ReturnType<typeof useGetOrder>>({
      data: [resolvedOrder, { ...data, status: 4 }, 0n],
    }),
  )

  const { result } = renderHook(() => useInboxStatus({ chainId: 1 }), {
    mockContractsCall: true,
  })

  expect(result.current).toBe('filled')
})

test('parameters: status claimed', () => {
  mockUseGetOrder.mockReturnValue(
    createMockReadContractResult<ReturnType<typeof useGetOrder>>({
      data: [resolvedOrder, { ...data, status: 5 }, 0n],
    }),
  )

  const { result } = renderHook(() => useInboxStatus({ chainId: 1 }), {
    mockContractsCall: true,
  })

  expect(result.current).toBe('filled')
})
