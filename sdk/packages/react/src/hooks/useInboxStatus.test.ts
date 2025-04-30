import { waitFor } from '@testing-library/react'
import { expect, test, vi } from 'vitest'
import {
  createMockQueryResult,
  orderId,
  orderStatusData,
  renderHook,
  resolvedOrder,
} from '../../test/index.js'
import { useGetOrder } from './useGetOrder.js'
import { useInboxStatus } from './useInboxStatus.js'

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

test('default: returns appropriate inbox status when order is resolved', async () => {
  const { result, rerender } = renderInboxStatusHook()

  // once on mount
  expect(useGetOrder).toHaveBeenCalledOnce()
  expect(result.current).toBe('unknown')

  mockUseGetOrder.mockReturnValue(
    createMockQueryResult({
      data: [resolvedOrder, orderStatusData, 0n],
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

test.each([
  ['not-found', 0],
  ['open', 1],
  ['rejected', 2],
  ['closed', 3],
  ['filled', 4],
  ['claimed', 5],
])('parameters: status %s if order status is %s', async (statusStr, status) => {
  mockUseGetOrder.mockReturnValue(
    createMockQueryResult({
      data: [resolvedOrder, { ...orderStatusData, status }, 0n],
    }),
  )

  const { result } = renderInboxStatusHook()

  expect(result.current).toBe(statusStr === 'claimed' ? 'filled' : statusStr)
})
