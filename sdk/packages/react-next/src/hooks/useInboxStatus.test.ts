import { waitFor } from '@testing-library/react'
import { beforeEach, expect, test, vi } from 'vitest'
import { orderId, renderHook, resolvedOrder } from '../../test/index.js'
import { createMockQueryResult } from '../../test/mocks.js'
import { useGetOrder } from './useGetOrder.js'
import { useInboxStatus } from './useInboxStatus.js'

const data = {
  status: 1,
  updatedBy: '0x123',
  timestamp: 1,
  rejectReason: 0,
} as const

const renderInboxStatusHook = () => {
  return renderHook(
    () =>
      useInboxStatus({
        chainId: 1,
      }),
    { mockContractsCall: true },
  )
}

const { mockUseGetOrder } = vi.hoisted(() => {
  return {
    mockUseGetOrder: vi.fn().mockImplementation(() => {
      return createMockQueryResult()
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
  mockUseGetOrder.mockReturnValue(createMockQueryResult())
})

test('default: returns appropriate inbox status when order is resolved', async () => {
  const { result, rerender } = renderInboxStatusHook()

  // once on mount
  expect(useGetOrder).toHaveBeenCalledOnce()
  expect(result.current).toBe('unknown')

  mockUseGetOrder.mockReturnValue(
    createMockQueryResult({
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
  const { result } = renderInboxStatusHook()

  expect(result.current).toBe('unknown')
})

test('parameters: status open', () => {
  mockUseGetOrder.mockReturnValue(
    createMockQueryResult({
      data: [resolvedOrder, data, 0n],
    }),
  )

  const { result } = renderInboxStatusHook()

  expect(result.current).toBe('open')
})

test('parameters: status rejected', () => {
  mockUseGetOrder.mockReturnValue(
    createMockQueryResult({
      data: [resolvedOrder, { ...data, status: 2 }, 0n],
    }),
  )

  const { result } = renderInboxStatusHook()

  expect(result.current).toBe('rejected')
})

test('parameters: status closed', () => {
  mockUseGetOrder.mockReturnValue(
    createMockQueryResult({
      data: [resolvedOrder, { ...data, status: 3 }, 0n],
    }),
  )

  const { result } = renderInboxStatusHook()

  expect(result.current).toBe('closed')
})

test('parameters: status filled', () => {
  mockUseGetOrder.mockReturnValue(
    createMockQueryResult({
      data: [resolvedOrder, { ...data, status: 4 }, 0n],
    }),
  )

  const { result } = renderInboxStatusHook()

  expect(result.current).toBe('filled')
})

test('parameters: status claimed', () => {
  mockUseGetOrder.mockReturnValue(
    createMockQueryResult({
      data: [resolvedOrder, { ...data, status: 5 }, 0n],
    }),
  )

  const { result } = renderInboxStatusHook()

  expect(result.current).toBe('filled')
})
