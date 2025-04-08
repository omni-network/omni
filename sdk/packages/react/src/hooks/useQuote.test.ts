import { waitFor } from '@testing-library/react'
import { beforeEach, expect, test, vi } from 'vitest'
import { renderHook } from '../../test/react.js'
import { quote } from '../../test/shared.js'
import * as api from '../internal/api.js'
import type { Quoteable } from '../types/quote.js'
import { useQuote } from './useQuote.js'

const token = '0x123'
const deposit = { token, isNative: false } satisfies Quoteable
const nativeExpense = { isNative: true } satisfies Quoteable

beforeEach(() => {
  vi.spyOn(api, 'fetchJSON').mockResolvedValue(quote)
})

const params = {
  srcChainId: 1,
  destChainId: 2,
  mode: 'expense',
  deposit: deposit,
  expense: nativeExpense,
  enabled: true,
} as const

const renderQuoteHook = (params: Parameters<typeof useQuote>[0]) => {
  return renderHook(() => useQuote(params), {
    mockContractsCall: true,
  })
}

test('default: fetches a quote once enabled', async () => {
  const { result, rerender } = renderHook(
    ({ enabled }: { enabled: boolean }) => useQuote({ ...params, enabled }),
    { initialProps: { enabled: false } },
  )

  expect(result.current.isPending).toBe(true)
  expect(result.current.query.data).toBeUndefined()
  expect(result.current.query.isFetched).toBe(false)

  rerender({ enabled: true })

  await waitFor(() => {
    expect(result.current.query.isFetched).toBe(true)
    expect(result.current.isPending).toBe(false)
    expect(result.current.query.data).toBeDefined()
  })
})

test('parameters: expense', () => {
  const { result, rerender } = renderQuoteHook({
    ...params,
    expense: { token, isNative: false },
  })

  expect(result.current).toBeDefined()

  rerender({ ...params, expense: { token, isNative: true } })

  expect(result.current).toBeDefined()
})

test('parameters: deposit', () => {
  const { result, rerender } = renderQuoteHook({
    ...params,
    deposit: { token, isNative: false },
  })

  expect(result.current).toBeDefined()

  rerender({ ...params, expense: { token, isNative: true } })

  expect(result.current).toBeDefined()
})

test('parameters: mode', () => {
  const { result, rerender } = renderQuoteHook({
    ...params,
    mode: 'expense',
    deposit: { isNative: true, amount: 100n },
    // TODO expense amount shouldn't be allowed if mode === 'expense'
    expense: { isNative: true, amount: 100n },
  })

  expect(result.current).toBeDefined()

  rerender({
    ...params,
    mode: 'deposit',
    deposit: { isNative: true, amount: 100n },
    expense: { isNative: true, amount: 100n },
  })

  expect(result.current).toBeDefined()
})

test('behaviour: quote does not fire when enabled is false', () => {
  const { result } = renderQuoteHook({ ...params, enabled: false })

  expect(result.current.isPending).toBe(true)
  expect(result.current.query.data).toBeUndefined()
  expect(result.current.query.isFetched).toBe(false)
})

test.each([
  'test',
  {},
  { deposit: { token, amount: '100' } },
  { expense: { token, amount: '100' } },
  { deposit: { token }, expense: { token } },
  { deposit: { amount: '100' }, expense: { amount: '99' } },
])(
  'behaviour: quote is error if response is not a quote: %s',
  async (mockReturn) => {
    vi.spyOn(api, 'fetchJSON').mockResolvedValue(mockReturn)

    const { result } = renderQuoteHook({ ...params, enabled: true })

    await waitFor(() => {
      expect(result.current.query.isLoading).toBe(false)
      expect(result.current.isError).toBe(true)
    })
  },
)
