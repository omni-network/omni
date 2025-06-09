import { GetRejectionError, type Rejection } from '@omni-network/core'
import type { useQuery } from '@tanstack/react-query'
import { waitFor } from '@testing-library/react'
import { expect, test, vi } from 'vitest'
import { orderId, renderHook } from '../../test/index.js'
import { type UseRejectionParams, useRejection } from './useRejection.js'

const { getRejection, useQueryMock } = vi.hoisted(() => {
  return {
    getRejection: vi.fn().mockImplementation(() => {
      return Promise.reject(new Error('No mock'))
    }),
    useQueryMock: vi.fn(),
  }
})

vi.mock('@omni-network/core', async () => {
  const actual = await vi.importActual('@omni-network/core')
  return { ...actual, getRejection }
})

vi.mock('@tanstack/react-query', async () => {
  const actual = await vi.importActual('@tanstack/react-query')
  const actualUseQuery = actual.useQuery as typeof useQuery
  return {
    ...actual,
    useQuery: useQueryMock.mockImplementation((params) => {
      return actualUseQuery(params)
    }),
  }
})

const mockRejection: Rejection = {
  txHash: '0xabcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890',
  rejectReason: 'Insufficient deposit',
}

test('default: returns rejection when core api returns a rejection', async () => {
  const { result, rerender } = renderHook(
    (props: UseRejectionParams) => useRejection({ srcChainId: 1, ...props }),
    { mockContractsCall: true },
  )

  expect(result.current.data).toBeUndefined()

  getRejection.mockResolvedValue(mockRejection)
  rerender({ srcChainId: 1, orderId, fromBlock: 1n })

  await waitFor(() =>
    expect(result.current.data?.txHash).toBe(mockRejection.txHash),
  )
  await waitFor(() =>
    expect(result.current.data?.rejectReason).toBe(mockRejection.rejectReason),
  )
})

test('behaviour: handles error when no rejection found', async () => {
  const { result, rerender } = renderHook(
    (props: UseRejectionParams) => useRejection({ srcChainId: 1, ...props }),
    { mockContractsCall: true },
  )

  getRejection.mockRejectedValue(
    new GetRejectionError("Expected exactly one 'Rejected' event but found 0."),
  )
  rerender({ srcChainId: 1, orderId, fromBlock: 1n })

  await waitFor(() => expect(result.current.isFetched).toBe(true))
  await waitFor(() => expect(result.current.status).toBe('error'))
  await waitFor(() =>
    expect(result.current.error?.message).toBe(
      "Expected exactly one 'Rejected' event but found 0.",
    ),
  )
})

test('behaviour: no core api call when orderId is undefined', () => {
  const { result } = renderHook(
    () => useRejection({ srcChainId: 1, fromBlock: 1n }),
    { mockContractsCall: true },
  )

  expect(result.current.data).toBeUndefined()
  expect(result.current.status).toBe('pending')
  expect(result.current.isFetched).toBeFalsy()
  expect(getRejection).not.toHaveBeenCalled()
})

test('behaviour: no core api call when fromBlock is undefined', () => {
  const { result } = renderHook(
    () => useRejection({ srcChainId: 1, orderId }),
    { mockContractsCall: true },
  )

  expect(result.current.data).toBeUndefined()
  expect(result.current.status).toBe('pending')
  expect(result.current.isFetched).toBeFalsy()
  expect(getRejection).not.toHaveBeenCalled()
})

test('behaviour: no core api call when srcChainId is undefined', () => {
  const { result } = renderHook(
    () => useRejection({ orderId, fromBlock: 1n }),
    { mockContractsCall: true },
  )

  expect(result.current.data).toBeUndefined()
  expect(result.current.status).toBe('pending')
  expect(result.current.isFetched).toBeFalsy()
  expect(getRejection).not.toHaveBeenCalled()
})

test('behaviour: no core api call when enabled is false', () => {
  const { result } = renderHook(
    () =>
      useRejection({
        srcChainId: 1,
        orderId,
        fromBlock: 1n,
        enabled: false,
      }),
    { mockContractsCall: true },
  )

  expect(result.current.data).toBeUndefined()
  expect(result.current.status).toBe('pending')
  expect(result.current.isFetched).toBeFalsy()
  expect(getRejection).not.toHaveBeenCalled()
})

test('behaviour: no core api call when all required inputs undefined', () => {
  const { result } = renderHook(() => useRejection({}), {
    mockContractsCall: true,
  })

  expect(result.current.data).toBeUndefined()
  expect(result.current.status).toBe('pending')
  expect(result.current.isFetched).toBeFalsy()
  expect(getRejection).not.toHaveBeenCalled()
})

test('parameters: calls getRejection with correct parameters', async () => {
  const fromBlock = 12345n
  const { rerender } = renderHook(
    (props: UseRejectionParams) => useRejection({ srcChainId: 1, ...props }),
    { mockContractsCall: true },
  )

  getRejection.mockResolvedValue(mockRejection)
  rerender({ srcChainId: 1, orderId, fromBlock })

  await waitFor(() => {
    expect(getRejection).toHaveBeenCalledWith({
      client: expect.any(Object),
      orderId,
      inboxAddress: '0x123',
      fromBlock,
    })
  })
})

test('parameters: queryOpts are passed react query', async () => {
  const fromBlock = 12345n
  const queryOpts = { staleTime: 5000 }
  const { rerender } = renderHook(
    (props: UseRejectionParams) =>
      useRejection({ srcChainId: 1, ...props, queryOpts }),
    { mockContractsCall: true },
  )

  getRejection.mockResolvedValue(mockRejection)
  rerender({ srcChainId: 1, orderId, fromBlock })

  await waitFor(() => {
    expect(useQueryMock).toHaveBeenCalledWith(
      expect.objectContaining(queryOpts),
    )
  })
})
