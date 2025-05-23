import * as core from '@omni-network/core'
import type { useQuery } from '@tanstack/react-query'
import { waitFor } from '@testing-library/react'
import { expect, test, vi } from 'vitest'
import { order, renderHook } from '../../test/index.js'
import { useValidateOrder } from './useValidateOrder.js'

const { useQueryMock } = vi.hoisted(() => {
  return { useQueryMock: vi.fn() }
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

const renderValidateOrderHook = (
  params: Parameters<typeof useValidateOrder>[0],
  options?: Parameters<typeof renderHook>[1],
) => {
  return renderHook(() => useValidateOrder(params), {
    mockContractsCall: true,
    ...options,
  })
}

test('default: native transfer order', async () => {
  vi.spyOn(core, 'validateOrder').mockResolvedValue({
    accepted: true,
  })

  const { result, rerender } = renderHook(
    ({ enabled }: { enabled: boolean }) =>
      useValidateOrder({
        order,
        enabled,
      }),
    {
      initialProps: { enabled: false },
    },
  )

  expect(result.current.status).toBe('pending')

  rerender({
    enabled: true,
  })

  await waitFor(() => expect(result.current.status).toBe('accepted'))
})

test('parameters: passes debug to validateOrder', async () => {
  const spy = vi.spyOn(core, 'validateOrder').mockResolvedValue({
    accepted: true,
  })
  renderValidateOrderHook({ order, debug: true, enabled: true })
  expect(spy).toHaveBeenCalledWith(
    expect.objectContaining({ ...order, debug: true }),
  )
})

test('parameters: passes through queryOpts to useQuery', async () => {
  const queryOpts = {
    refetchInterval: 5000,
    staleTime: 10000,
  }
  renderValidateOrderHook({ order, enabled: true, queryOpts })
  expect(useQueryMock).toHaveBeenCalledWith(expect.objectContaining(queryOpts))
})

test('behaviour: pending if query not fired', async () => {
  const { result } = renderHook(() =>
    useValidateOrder({ order, enabled: false }),
  )

  await waitFor(() => expect(result.current.status).toBe('pending'))
})

test('behaviour: error if response is error', async () => {
  vi.spyOn(core, 'validateOrder').mockResolvedValue({
    error: {
      code: 1,
      message: 'an error',
    },
  })

  const { result } = renderValidateOrderHook(
    { order, enabled: true },
    {
      mockContractsCall: false,
    },
  )

  await waitFor(() => expect(result.current.status).toBe('error'))
  await waitFor(() =>
    expect(
      result.current.status === 'error' ? result.current.error.message : null,
    ).toBe('an error'),
  )
})

test('behaviour: rejected if response is rejected', async () => {
  vi.spyOn(core, 'validateOrder').mockResolvedValue({
    rejected: true,
    rejectReason: 'a reason',
    rejectDescription: 'a description',
  })

  const { result } = renderValidateOrderHook(
    { order, enabled: true },
    {
      mockContractsCall: false,
    },
  )

  await waitFor(() => expect(result.current.status).toBe('rejected'))
  await waitFor(() =>
    expect(
      result.current.status === 'rejected' ? result.current.rejectReason : null,
    ).toBe('a reason'),
  )
})

test('behaviour: error if call throws', async () => {
  const error = new Error('Unexpected validation response')
  vi.spyOn(core, 'validateOrder').mockRejectedValue(error)

  const { result } = renderValidateOrderHook({ order, enabled: true })

  await waitFor(() => {
    expect(result.current.status).toBe('error')
    if (result.current.status === 'error') {
      expect(result.current.error).toBe(error)
    }
  })
})

test('behaviour: returns an error instead of throwing when the order encoding throws', async () => {
  const error = new Error('Address "0xinvalid" is invalid')
  vi.spyOn(core, 'validateOrder').mockRejectedValue(error)

  const invalidOrder = {
    ...order,
    calls: [{ ...order.calls[0], args: ['0xinvalid', 0n] }],
  }
  const { result } = renderValidateOrderHook({
    order: invalidOrder,
    enabled: true,
  })

  await waitFor(() => {
    expect(result.current.status).toBe('error')
    if (result.current.status === 'error') {
      expect(result.current.error.message).toMatch(
        'Address "0xinvalid" is invalid',
      )
    }
  })
})
