import type { Quoteable } from '@omni-network/core'
import * as core from '@omni-network/core'
import { waitFor } from '@testing-library/react'
import { beforeEach, expect, test, vi } from 'vitest'
import { quote, renderHook } from '../../test/index.js'
import { useQuote } from './useQuote.js'

const token = '0x123'

beforeEach(() => {
  vi.spyOn(core, 'getQuote').mockResolvedValue(quote)
})

const params = {
  srcChainId: 1,
  destChainId: 2,
  mode: 'expense',
  deposit: { token, isNative: false } satisfies Quoteable,
  expense: { isNative: true } satisfies Quoteable,
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

test('behaviour: quote is error if call throws', async () => {
  const error = new Error('Unexpected quote response')
  vi.spyOn(core, 'getQuote').mockRejectedValue(error)

  const { result } = renderQuoteHook({ ...params, enabled: true })

  await waitFor(() => {
    expect(result.current.query.isLoading).toBe(false)
    expect(result.current.isError).toBe(true)
    if (result.current.isError) {
      expect(result.current.error).toBe(error)
    }
  })
})
