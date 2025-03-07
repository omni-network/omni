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

test('default', async () => {
  const { result, rerender } = renderHook(
    () =>
      useGetOrder({
        chainId: 1,
      }),
    { mockContractsCall: true },
  )

  expect(result.current.data).toBeUndefined()

  useReadContract.mockReturnValue(
    createMockReadContractResult({
      data: [
        resolvedOrder,
        { status: 1, claimant: '0x123', timestamp: 1 } as const,
      ],
      isSuccess: true,
      status: 'success',
    }),
  )

  rerender({
    chainId: 1,
    orderId,
  })

  await waitFor(() => result.current.data?.[0].orderId === orderId)
  await waitFor(() => result.current.data?.[1].status === 1)
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
