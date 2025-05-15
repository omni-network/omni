import type { useQuery } from '@tanstack/react-query'
import { waitFor } from '@testing-library/react'
import { expect, test, vi } from 'vitest'
import { renderHook, resolvedOrder } from '../../test/index.js'
import { type UseDidFillParams, useDidFill } from './useDidFill.js'

const { didFill, useQueryMock } = vi.hoisted(() => {
  return {
    didFill: vi.fn().mockImplementation(() => {
      return Promise.reject(new Error('No mock'))
    }),
    useQueryMock: vi.fn(),
  }
})

vi.mock('@omni-network/core', async () => {
  const actual = await vi.importActual('@omni-network/core')
  return { ...actual, didFill }
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

const renderDidFillHook = (withResolvedOrder = false, queryOpts = {}) => {
  return renderHook(
    (props: Partial<UseDidFillParams>) =>
      useDidFill({
        destChainId: 1,
        resolvedOrder: withResolvedOrder ? resolvedOrder : undefined,
        queryOpts,
        ...props,
      }),
    { mockContractsCall: true },
  )
}

test('default: returns true when core api returns truthy', async () => {
  const { result, rerender } = renderDidFillHook()

  expect(result.current.data).toBeUndefined()
  expect(didFill).not.toHaveBeenCalled()

  didFill.mockResolvedValue(true)

  rerender({
    destChainId: 1,
    resolvedOrder,
  })

  await waitFor(() => expect(result.current.data).toBe(true))
  expect(didFill).toHaveBeenCalled()
})

test('parameters: passes through queryOpts to useQuery', async () => {
  const queryOpts = {
    refetchInterval: 5000,
    staleTime: 10000,
  }
  renderDidFillHook(true, queryOpts)
  expect(useQueryMock).toHaveBeenCalledWith(expect.objectContaining(queryOpts))
})

test('behaviour: no exception if core api throws', async () => {
  didFill.mockRejectedValue(new Error('Contract read failed'))

  const { result } = renderDidFillHook(true)

  await waitFor(() => expect(result.current.status).toBe('error'))
  expect(result.current.isError).toBe(true)
  expect(result.current.data).toBeUndefined()
  expect(didFill).toHaveBeenCalled()
})

test('behaviour: no core api call when resolvedOrder is undefined', async () => {
  const { result } = renderDidFillHook()

  expect(result.current.data).toBeUndefined()
  expect(result.current.status).toBe('pending')
  expect(result.current.isFetched).toBe(false)
  expect(didFill).not.toHaveBeenCalledOnce()
})
