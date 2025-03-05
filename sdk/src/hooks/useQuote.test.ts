import { waitFor } from '@testing-library/react'
import { zeroAddress } from 'viem'
import { expect, test, vi } from 'vitest'
import { renderHook } from '../../test/react.js'
import type { Quoteable } from '../types/quote.js'
import { useQuote } from './useQuote.js'

const token = '0x123'
const deposit = { token, isNative: false } satisfies Quoteable
const nativeExpense = { isNative: true } satisfies Quoteable

const params = {
  srcChainId: 1,
  destChainId: 2,
  mode: 'expense',
  deposit: deposit,
  expense: nativeExpense,
  enabled: true,
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

test('default', async () => {
  const { result, rerender } = renderHook(
    () => useQuote({ ...params, enabled: false }),
    {
      mockContractsCall: true,
    },
  )

  expect(result.current.isPending).toBe(true)
  expect(result.current.query.data).toBeUndefined()
  expect(result.current.query.isFetched).toBeFalsy()

  fetchJSON.mockResolvedValue({
    deposit: { token, amount: '100' },
    expense: { token: zeroAddress, amount: '99' },
  })

  rerender({ ...params, enabled: true })

  await Promise.all([
    waitFor(() => result.current.isPending === false),
    waitFor(() => result.current.isError === false),
    waitFor(() => result.current.isSuccess === true),
    waitFor(() => result.current.query.data?.deposit.token === token),
    waitFor(() => result.current.query.data?.deposit.amount === BigInt(100)),
    waitFor(() => result.current.query.data?.expense.token === zeroAddress),
    waitFor(() => result.current.query.data?.expense.amount === BigInt(99)),
  ])
})

test('parameters: expense', () => {
  const { result, rerender } = renderHook(
    () => useQuote({ ...params, expense: { token, isNative: false } }),
    {
      mockContractsCall: true,
    },
  )

  expect(result.current).not.toBeUndefined()

  // TODO token shouldn't be allowed if isNative === true
  rerender({ ...params, expense: { token, isNative: true } })

  expect(result.current).not.toBeUndefined()
})

test('parameters: deposit', () => {
  const { result, rerender } = renderHook(
    () =>
      useQuote({
        ...params,
        deposit: { token, isNative: false },
      }),
    {
      mockContractsCall: true,
    },
  )

  expect(result.current).not.toBeUndefined()

  // TODO token shouldn't be allowed if isNative === true
  rerender({ ...params, expense: { token, isNative: true } })

  expect(result.current).not.toBeUndefined()
})

test('parameters: mpde', () => {
  const { result, rerender } = renderHook(
    () =>
      useQuote({
        ...params,
        mode: 'expense',
        deposit: { isNative: true, amount: 100n },
        // TODO expense amount shouldn't be allowed if mode === 'expense'
        expense: { isNative: true, amount: 100n },
      }),
    {
      mockContractsCall: true,
    },
  )

  expect(result.current).not.toBeUndefined()

  rerender({
    ...params,
    mode: 'deposit',
    // TODO deposit amount shouldn't be allowed if mode === 'deposit'
    deposit: { isNative: true, amount: 100n },
    expense: { isNative: true, amount: 100n },
  })

  expect(result.current).not.toBeUndefined()
})

test('behaviour: quote does not fire when enabled is false', () => {
  const { result } = renderHook(() => useQuote({ ...params, enabled: false }), {
    mockContractsCall: true,
  })

  expect(result.current.isPending).toBe(true)
  expect(result.current.query.data).toBeUndefined()
  expect(result.current.query.isFetched).toBeFalsy()
})

test('behaviour: quote is error if response is not a quote', async () => {
  const { result } = renderHook(() => useQuote({ ...params, enabled: false }), {
    mockContractsCall: true,
  })

  fetchJSON.mockReturnValue('test')

  await waitFor(() => result.current.isPending === false)
  await waitFor(() => result.current.isError === true)

  fetchJSON.mockReturnValue({
    deposit: { token, amount: '100' },
  })

  await waitFor(() => result.current.isPending === false)
  await waitFor(() => result.current.isError === true)

  fetchJSON.mockReturnValue({
    expense: { token, amount: '100' },
  })

  await waitFor(() => result.current.isPending === false)
  await waitFor(() => result.current.isError === true)

  fetchJSON.mockReturnValue({
    deposit: { token },
    expense: { token },
  })

  await waitFor(() => result.current.isPending === false)
  await waitFor(() => result.current.isError === true)

  fetchJSON.mockReturnValue({
    deposit: { amount: '100' },
    expense: { amount: '99' },
  })

  await waitFor(() => result.current.isPending === false)
  await waitFor(() => result.current.isError === true)
})
