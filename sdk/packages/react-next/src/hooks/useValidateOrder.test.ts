import * as core from '@omni-network/core'
import { waitFor } from '@testing-library/react'
import { expect, test, vi } from 'vitest'
import { accounts, renderHook } from '../../test/index.js'
import { order } from '../../test/shared.js'
import { useValidateOrder } from './useValidateOrder.js'

// TODO calls as empty array should not be allowed // throw error

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
  vi.spyOn(core, 'validateOrderEncoded').mockResolvedValue({
    accepted: true,
  })

  const { result, rerender } = renderHook(
    ({ enabled }: { enabled: boolean }) =>
      useValidateOrder({
        order: {
          owner: accounts[0],
          srcChainId: 1,
          destChainId: 2,
          calls: [
            {
              target: accounts[0],
              value: 0n,
            },
          ],
          deposit: {
            amount: 0n,
          },
          expense: {
            amount: 0n,
          },
        },
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

test('default: order', async () => {
  vi.spyOn(core, 'validateOrderEncoded').mockResolvedValue({
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

test('behaviour: pending if query not fired', async () => {
  const { result } = renderHook(() =>
    useValidateOrder({ order, enabled: false }),
  )

  await waitFor(() => expect(result.current.status).toBe('pending'))
})

test('behaviour: error if response is error', async () => {
  vi.spyOn(core, 'validateOrderEncoded').mockResolvedValue({
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
  vi.spyOn(core, 'validateOrderEncoded').mockResolvedValue({
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
  vi.spyOn(core, 'validateOrderEncoded').mockRejectedValue(error)

  const { result } = renderValidateOrderHook({ order, enabled: true })

  await waitFor(() => {
    expect(result.current.status).toBe('error')
    if (result.current.status === 'error') {
      expect(result.current.error).toBe(error)
    }
  })
})

test('behaviour: returns an error instead of throwing when the order encoding throws', async () => {
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
