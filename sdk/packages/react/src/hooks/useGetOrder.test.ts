import { waitFor } from '@testing-library/react'
import { expect, test } from 'vitest'
import { orderId, renderHook, resolvedOrder } from '../../test/index.js'
import {
  createMockReadContractResult,
  mockWagmiHooks,
} from '../../test/mocks.js'
import { useGetOrder } from './useGetOrder.js'

const { useReadContract } = mockWagmiHooks()

const renderGetOrderHook = (params: Parameters<typeof useGetOrder>[0]) => {
  return renderHook(() => useGetOrder(params), { mockContractsCall: true })
}

test('default: returns order when contract read returns an order', async () => {
  const { result, rerender } = renderGetOrderHook({ chainId: 1 })

  expect(result.current.data).toBeUndefined()

  useReadContract.mockReturnValue(
    createMockReadContractResult<ReturnType<typeof useGetOrder>>({
      data: [
        resolvedOrder,
        {
          status: 1,
          updatedBy: '0x123',
          timestamp: 1,
          rejectReason: 0,
        } as const,
        0n,
      ],
      isSuccess: true,
      status: 'success',
    }),
  )

  rerender({
    chainId: 1,
    orderId,
  })

  await waitFor(() => expect(result.current.data?.[0].orderId).toBe(orderId))
  await waitFor(() => expect(result.current.data?.[1].status).toBe(1))
})

test('behaviour: no contract read when orderId is undefined', () => {
  const { result } = renderGetOrderHook({ chainId: 1 })

  expect(result.current.data).toBeUndefined()
  expect(result.current.status).toBe('pending')
  expect(result.current.isFetched).toBe(false)
  // once on mount
  expect(useReadContract).toHaveBeenCalledOnce()
})

test('behaviour: no contract read when chainId is undefined', () => {
  const { result } = renderGetOrderHook({ orderId })

  expect(result.current.data).toBeUndefined()
  expect(result.current.status).toBe('pending')
  expect(result.current.isFetched).toBe(false)
  // once on mount
  expect(useReadContract).toHaveBeenCalledOnce()
})

test('behaviour: no contract read when all inputs undefined', () => {
  const { result } = renderGetOrderHook({})

  expect(result.current.data).toBeUndefined()
  expect(result.current.status).toBe('pending')
  expect(result.current.isFetched).toBe(false)
  // once on mount
  expect(useReadContract).toHaveBeenCalledOnce()
})
