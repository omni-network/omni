import { waitFor } from '@testing-library/react'
import { beforeEach, expect, test, vi } from 'vitest'
import { orderId, renderHook, resolvedOrder } from '../../test/index.js'
import { createMockReadContractResult } from '../../test/mocks.js'
import { useGetOrder } from './useGetOrder.js'

const { useReadContract } = vi.hoisted(() => {
  return {
    useReadContract: vi.fn().mockImplementation(() => {
      return createMockReadContractResult()
    }),
  }
})

vi.mock('wagmi', async () => {
  const actual = await vi.importActual('wagmi')
  return {
    ...actual,
    useReadContract,
  }
})

beforeEach(() => {
  useReadContract.mockReturnValue(createMockReadContractResult())
})

test('default: returns order when contract read returns an order', async () => {
  const { result, rerender } = renderHook(
    () =>
      useGetOrder({
        chainId: 1,
      }),
    { mockContractsCall: true },
  )

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
  const { result } = renderHook(
    () =>
      useGetOrder({
        chainId: 1,
      }),
    { mockContractsCall: true },
  )

  expect(result.current.data).toBeUndefined()
  expect(result.current.status).toBe('pending')
  expect(result.current.isFetched).toBeFalsy()
  // once on mount
  expect(useReadContract).toHaveBeenCalledOnce()
})

test('behaviour: no contract read when chainId is undefined', () => {
  const { result } = renderHook(
    () =>
      useGetOrder({
        orderId,
      }),
    { mockContractsCall: true },
  )

  expect(result.current.data).toBeUndefined()
  expect(result.current.status).toBe('pending')
  expect(result.current.isFetched).toBeFalsy()
  // once on mount
  expect(useReadContract).toHaveBeenCalledOnce()
})

test('behaviour: no contract read when all inputs undefined', () => {
  const { result } = renderHook(() => useGetOrder({}), {
    mockContractsCall: true,
  })

  expect(result.current.data).toBeUndefined()
  expect(result.current.status).toBe('pending')
  expect(result.current.isFetched).toBeFalsy()
  // once on mount
  expect(useReadContract).toHaveBeenCalledOnce()
})
