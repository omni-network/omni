import { waitFor } from '@testing-library/react'
import { erc20Abi } from 'viem'
import { expect, test, vi } from 'vitest'
import { accounts, renderHook } from '../../test/index.js'
import { useValidateOrder } from './useValidateOrder.js'

// TODO calls as empty array should not be allowed // throw error

const order = {
  owner: accounts[0],
  srcChainId: 1,
  destChainId: 2,
  calls: [
    {
      abi: erc20Abi,
      functionName: 'transfer',
      target: '0x23e98253f372ee29910e22986fe75bb287b011fc',
      value: BigInt(0),
      args: [accounts[0], 0n],
    },
  ],
  deposit: {
    token: '0x123',
    amount: 0n,
  },
  expense: {
    token: '0x123',
    amount: 0n,
  },
} as const

const { fetchJSON } = vi.hoisted(() => {
  return {
    fetchJSON: vi.fn(),
  }
})

vi.mock('../internal/api.js', async () => {
  const actual = await vi.importActual('../internal/api.js')
  return {
    ...actual,
    fetchJSON,
  }
})

test('default: native transfer order', async () => {
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

  fetchJSON.mockResolvedValue({
    accepted: true,
  })

  rerender({
    enabled: true,
  })
  await waitFor(() => expect(result.current.status).toBe('accepted'))
})

test('default: order', async () => {
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

  fetchJSON.mockResolvedValue({
    accepted: true,
  })

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
  fetchJSON.mockResolvedValue({
    error: {
      code: 1,
      message: 'an error',
    },
  })

  const { result } = renderHook(() =>
    useValidateOrder({ order, enabled: true }),
  )

  await waitFor(() => expect(result.current.status).toBe('error'))
  await waitFor(() =>
    expect(
      result.current.status === 'error' ? result.current.error.message : null,
    ).toBe('an error'),
  )
})

test('behaviour: rejected if response is rejected', async () => {
  fetchJSON.mockResolvedValue({
    rejected: true,
    rejectReason: 'a reason',
    rejectDescription: 'a description',
  })

  const { result } = renderHook(() =>
    useValidateOrder({ order, enabled: true }),
  )

  await waitFor(() => expect(result.current.status).toBe('rejected'))
  await waitFor(() =>
    expect(
      result.current.status === 'rejected' ? result.current.rejectReason : null,
    ).toBe('a reason'),
  )
})

test.each([
  'test',
  {},
  { rejected: true },
  { rejected: true, rejectReason: 'a reason' },
  { rejecetd: true, rejectDescription: 'a description' },
])('behaviour: error if response is not valid: %s', async (mockReturn) => {
  const { result } = renderHook(() =>
    useValidateOrder({ order, enabled: true }),
  )

  fetchJSON.mockReturnValue(mockReturn)

  await waitFor(() => result.current.status === 'error')
})

test('behaviour: returns an error instead of throwing when the order encoding throws', async () => {
  const invalidOrder = {
    ...order,
    calls: [{ ...order.calls[0], args: ['0xinvalid', 0n] }],
  }
  const { result } = renderHook(() => {
    // @ts-expect-error: invalid order
    return useValidateOrder({ order: invalidOrder, enabled: true })
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
