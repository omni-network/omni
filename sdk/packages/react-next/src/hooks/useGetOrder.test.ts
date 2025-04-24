import { waitFor } from '@testing-library/react'
import { expect, test, vi } from 'vitest'
import { orderId, renderHook, resolvedOrder } from '../../test/index.js'
import { type UseGetOrderParameters, useGetOrder } from './useGetOrder.js'

const { getOrder } = vi.hoisted(() => {
  return {
    getOrder: vi.fn().mockImplementation(() => {
      return Promise.reject(new Error('No mock'))
    }),
  }
})

vi.mock('@omni-network/core', async () => {
  const actual = await vi.importActual('@omni-network/core')
  return { ...actual, getOrder }
})

test('default: returns order when core api returns an order', async () => {
  const { result, rerender } = renderHook(
    (props: UseGetOrderParameters) => useGetOrder({ chainId: 1, ...props }),
    { mockContractsCall: true },
  )

  expect(result.current.data).toBeUndefined()

  getOrder.mockReturnValue(
    Promise.resolve([
      resolvedOrder,
      {
        status: 1,
        updatedBy: '0x123',
        timestamp: 1,
        rejectReason: 0,
      } as const,
      0n,
    ]),
  )
  rerender({ chainId: 1, orderId })
  await waitFor(() => expect(result.current.data?.[0].orderId).toBe(orderId))
  await waitFor(() => expect(result.current.data?.[1].status).toBe(1))
})

test('behaviour: no core api call when orderId is undefined', () => {
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
  expect(getOrder).not.toHaveBeenCalled()
})

test('behaviour: no core api call when chainId is undefined', () => {
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
  expect(getOrder).not.toHaveBeenCalled()
})

test('behaviour: no core api call when all inputs undefined', () => {
  const { result } = renderHook(() => useGetOrder({}), {
    mockContractsCall: true,
  })

  expect(result.current.data).toBeUndefined()
  expect(result.current.status).toBe('pending')
  expect(result.current.isFetched).toBeFalsy()
  expect(getOrder).not.toHaveBeenCalled()
})
